package callback

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type Record struct {
	ID          string //DB的ID，同一个事务内部，ID一致
	db          *gorm.DB
	IsTx        bool          //是否是事务
	Value       reflect.Value //gorm.statement.ReflectValue
	Dest        any           //gorm.statement.Dest
	Model       any           //gorm.statement.Model
	DataModels  []DataModel
	ConnPool    gorm.ConnPool
	StatementID string
	TableName   string
	UniqueKey   string
	SQL         string
	OrmMethod   string // insert,update,delete,raw
	Operation   string // insert,update,delete
	Position    string // before, after
	UpdatedAt   time.Time
}

func (r *Record) needStage() bool {
	if r.OrmMethod == ChangeOptionDelete || r.OrmMethod == ChangeOptionRaw {
		return r.Position == ChangePositionBefore
	}
	return false
}

func (r *Record) hasStage() bool {
	if r.OrmMethod == ChangeOptionDelete || r.OrmMethod == ChangeOptionRaw {
		return r.OrmMethod == ChangePositionAfter
	}
	return false
}

func (r *Record) ParseDataModels() {
	if r.OrmMethod == ChangeOptionRaw {
		r.DataModels = r.modelsFormSQL()
	}
	r.DataModels = r.modelsFormBean()
}

func (r *Record) modelsFormBean() []DataModel {
	var models []DataModel
	//如果是一个slice，那应该是一个合法的
	if r.Value.Kind() == reflect.Slice {
		for i := 0; i < r.Value.Len(); i++ {
			value := r.Value.Index(i)
			dataModel, ok := IsValidDataModel(value.Interface(), r.UniqueKey)
			if ok {
				models = append(models, *dataModel)
			}
		}
		if len(models) > 0 {
			return models
		}
	}
	//尝试从reflect_value寻找
	if r.Value.IsValid() {
		if dataModel, ok := IsValidDataModel(r.Value.Interface(), r.UniqueKey); ok {
			models = append(models, *dataModel)
			return models
		}
	}
	//尝试从dest寻找
	if dataModel, ok := IsValidDataModel(r.Dest, r.UniqueKey); ok {
		models = append(models, *dataModel)
		return models
	}
	//尝试从model寻找
	if dataModel, ok := IsValidDataModel(r.Dest, r.UniqueKey); ok {
		models = append(models, *dataModel)
		return models
	}
	//从SQl中获取
	if r.Operation != ChangeOptionCreate {
		return r.modelsFormSQL()
	}
	//按照gorm的create源代码，走到这里应该是错的类型，不存在指针类型的
	fmt.Printf("%v, is not a valid callback model", string(lo.T2(json.Marshal(r.Value.Interface())).A))
	return models
}

func (r *Record) modelsFormSQL() []DataModel {
	newSQL := toSelectSQL(r.SQL)
	if newSQL == "" || r.TableName == "" {
		fmt.Printf("can not find table name or invalid SQL from %v to %v,record: %v\n", r.SQL, newSQL, printModel(r))
		return nil
	}
	data, err := queryFromSQL(context.Background(), r.db, newSQL)
	if err != nil {
		fmt.Printf("query error %v,SQL from %v to %v", err.Error(), r.SQL, newSQL)
		return nil
	}
	return data
}
