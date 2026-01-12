package callback

import (
	"fmt"
	"gorm.io/gorm"
	"runtime"
	"time"
)

const (
	ChangePositionBefore = "before"
	ChangePositionAfter  = "after"
)

const (
	ChangeOptionSelect = "select"
	ChangeOptionCreate = "insert"
	ChangeOptionUpdate = "update"
	ChangeOptionDelete = "delete"
	ChangeOptionRaw    = "raw"
)

const (
	ChanSize                   = 10000 //暂存的记录总数
	MaxSubGoRoutineSize        = 10
	MaxCommitWaitCheckInterval = 15 * time.Second
	CommitCheckInterval        = 5 * time.Millisecond
	TxWaitInternal             = 3 * time.Second //事务提交等待时间，出现过事务已经提交，但是结果还是查不到的现象
)

type DatabaseCallback struct {
	db          *gorm.DB
	txGroups    txGroups
	transporter Transporter
	workerSize  int
	rc          chan *Record
	gc          chan int
}

type RegisterItem struct {
	Position string
	Name     string
	Fn       func(db *gorm.DB)
}

func NewDatabaseCallback(db *gorm.DB) *DatabaseCallback {
	callback := &DatabaseCallback{
		db:          db,
		txGroups:    NewTxGroups(),
		transporter: NewTransporter(),
		rc:          make(chan *Record, ChanSize),
		workerSize:  MaxSubGoRoutineSize,
		gc:          make(chan int, MaxSubGoRoutineSize),
	}
	for i := 0; i < MaxSubGoRoutineSize; i++ {
		callback.gc <- i
	}
	return callback
}

// RegisterCallback 注册回调
func (d *DatabaseCallback) RegisterCallback() {
	RegisterCallback(d.db)
}

// RegisterCallback 注册回调
func RegisterCallback(db *gorm.DB) {
	db.Callback().Create().After("gorm:create").Register("create_record:after_save", defaultDatabaseCallback.afterSave)
	db.Callback().Update().After("gorm:update").Register("update_record:after_update", defaultDatabaseCallback.afterUpdate)
	db.Callback().Delete().Before("gorm:delete").Register("delete_record:before_delete", defaultDatabaseCallback.beforeDelete)
	db.Callback().Delete().After("gorm:delete").Register("delete_record:after_delete", defaultDatabaseCallback.afterDelete)
	db.Callback().Raw().After("gorm:raw").Register("raw_record:before_raw", defaultDatabaseCallback.beforeRaw)
	db.Callback().Raw().After("gorm:raw").Register("raw_record:after_raw", defaultDatabaseCallback.afterRaw)
}
func (d *DatabaseCallback) afterSave(client *gorm.DB) {
	d.Entry(client, ChangeOptionCreate, ChangePositionAfter)
}
func (d *DatabaseCallback) afterUpdate(client *gorm.DB) {
	d.Entry(client, ChangeOptionUpdate, ChangePositionAfter)
}
func (d *DatabaseCallback) beforeDelete(client *gorm.DB) {
	d.Entry(client, ChangeOptionDelete, ChangePositionBefore)
}
func (d *DatabaseCallback) afterDelete(client *gorm.DB) {
	d.Entry(client, ChangeOptionDelete, ChangePositionAfter)
}
func (d *DatabaseCallback) beforeRaw(client *gorm.DB) {
	d.Entry(client, ChangeOptionRaw, ChangePositionBefore)
}
func (d *DatabaseCallback) afterRaw(client *gorm.DB) {
	d.Entry(client, ChangeOptionRaw, ChangePositionAfter)
}

// Entry 所有的回调的入口，处理各种操作
func (d *DatabaseCallback) Entry(db *gorm.DB, ormMethod, position string) {
	defer func() {
		if e := recover(); e != nil {
			if v, ok := e.(error); ok {
				fmt.Printf("remind error,%v\n", v)
			}
			buf := make([]byte, 1024*1024)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("%s", buf)
		}
	}()
	sql := db.Statement.ToSQL(func(tx *gorm.DB) *gorm.DB { return db })
	//如果使用的是raw  SQL，判断下是否是查询语句
	operation := ormMethod
	if ormMethod == ChangeOptionRaw {
		operation = sqlOperation(sql)
		if operation == ChangeOptionSelect {
			return
		}
	}
	//有错误的，也不必继续了
	if db.Statement.Error != nil {
		return
	}
	tableName := extractTableName(sql, db)
	//如果没有注册的，放弃
	if !d.transporter.HasRegistered(tableName) {
		fmt.Printf("not valid callback model %v \n", sql)
		return
	}
	record := &Record{
		ID:          fmt.Sprintf("%p", d.db),
		db:          d.db,
		StatementID: StatementID(db.Statement),
		TableName:   tableName,
		UniqueKey:   d.transporter.UniqueKey(tableName),
		ConnPool:    db.Statement.ConnPool,
		IsTx:        IsTx(db),
		SQL:         sql,
		Dest:        db.Statement.Dest,
		Model:       db.Statement.Model,
		Value:       db.Statement.ReflectValue,
		OrmMethod:   ormMethod,
		Operation:   operation,
		Position:    position,
		UpdatedAt:   time.Now(),
	}
	record.ParseDataModels()
	go d.input(record) // 不堵塞主进程执行
}

func (d *DatabaseCallback) input(record *Record) {
	d.rc <- record
}

func (d *DatabaseCallback) Run() {
	for {
		select {
		case data, ok := <-d.rc:
			if !ok {
				fmt.Printf("close callback msg chan")
				return
			}
			d.handle(data)
		}
	}
}

// handle 顺序执行
func (d *DatabaseCallback) handle(record *Record) {
	i := <-d.gc
	go func() {
		//如果事务失败了，那就不发送了,  TODO 无法识别事务回滚失败或者成功
		//暂时不处理事务事务成功或者失败的回滚，失败也发，接收端使用ID查询真实数据
		//d.WaitTx(record)
		time.Sleep(TxWaitInternal)
		d.transporter.Handle(record)
		d.gc <- i
	}()
}

// WaitTx 等待事务结束，如果true代表可以继续执行下去，如果不是，代表不可以执行下去
func (d *DatabaseCallback) WaitTx(r *Record) bool {
	if !r.IsTx {
		return true
	}
	timeout := time.NewTimer(MaxCommitWaitCheckInterval)
	ti := time.NewTimer(CommitCheckInterval)
	defer func() {
		timeout.Stop()
		ti.Stop()
	}()

	sch := make(chan any, 500)
	lock := d.InGroup(r.ID, sch)

	defer func() {
		d.CloseGroup(r.ID)
	}()

	for {
		select {
		case _, ok := <-sch:
			if !ok {
				return false
			}
			return true
		case <-timeout.C:
			fmt.Printf("record not commitedstate %v;\n", printModel(r.DataModels))
			d.CloseGroup(r.ID)
			return false
		case <-ti.C:
			lock.Lock()
			hasFinished := CheckTxFinished(r.ConnPool)
			lock.Unlock()

			if hasFinished {
				d.ReleaseGroup(r.ID)
				return true
			}
		}
	}
}

func (d *DatabaseCallback) Close() {
	close(d.rc)
}
