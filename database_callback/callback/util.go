package callback

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"strings"
)

func StatementID(s *gorm.Statement) string {
	return fmt.Sprintf("%p", s)
}

// 支持从update，delete 语句到select 语句
func toSelectSQL(sql string) string {
	op := sqlOperation(sql)
	if op == ChangeOptionDelete {
		return removeDeleteCondition(delete2select(sql))
	}
	if op == ChangeOptionUpdate {
		return removeDeleteCondition(update2select(sql))
	}
	return ""
}

// findDeleteKey 判断下当前SQL的软删除key
func findDeleteKey(sql string) string {
	deleteKey := "deleted_at"
	if strings.Contains(sql, "delete_time") {
		deleteKey = "delete_time"
	}
	return deleteKey
}

// removeDeleteCondition 这个是针对当前情况的j简单处理
func removeDeleteCondition(sql string) string {
	deleteKey := findDeleteKey(sql)

	lowerSQL := strings.ToLower(sql)
	deleteIndex := strings.LastIndex(lowerSQL, deleteKey)
	whereIndex := strings.LastIndex(lowerSQL, "where")
	if deleteIndex < 0 || whereIndex < 0 {
		return sql
	}
	//select * from xxx where deleted_at=0
	if !strings.Contains(lowerSQL, " and ") && (deleteIndex > whereIndex) {
		return strings.Split(lowerSQL, " where ")[0]
	}
	//select * from xxx where deleted_at=0 and name='xx'
	qs := strings.Split(lowerSQL, " where ")
	if len(qs) != 2 {
		return sql
	}
	conditions := strings.Split(qs[1], " and ")
	newConditions := make([]string, 0)
	for i := range conditions {
		if strings.Contains(conditions[i], deleteKey) {
			continue
		}
		newConditions = append(newConditions, conditions[i])
	}
	return qs[0] + " where " + strings.Join(newConditions, " and ")
}

// update2select 单纯的更新操作
func update2select(sql string) string {
	if sql == "" {
		return ""
	}
	updateReg1, _ := regexp.Compile("(?i)^UPDATE\\b") //匹配以UPDATE开头的SQL语句的UPDATE ，不区分大小写
	updateReg2, _ := regexp.Compile("(?i)UPDATE\\s+.*?(\\bSET\\s+.*?)(?:\\s+WHERE\\s+.*|$)")

	//替换掉UPDATE的条件
	sql = updateReg2.ReplaceAllStringFunc(sql, func(s string) string {
		subMatches := updateReg2.FindStringSubmatch(s)
		if len(subMatches) > 0 {
			return strings.ReplaceAll(sql, subMatches[1], "")
		}
		return s
	})
	return updateReg1.ReplaceAllString(sql, "SELECT * FROM ")
}

// delete2select 删除语句改查询语句，条件是where条件中不带 delete_at=0 的过滤条件
func delete2select(sql string) string {
	if sql == "" {
		return ""
	}
	deleteReg, _ := regexp.Compile("(?i)^DELETE\\b")  //匹配以DELETE开头的SQL语句的DELETE ，不区分大小写
	updateReg1, _ := regexp.Compile("(?i)^UPDATE\\b") //匹配以UPDATE开头的SQL语句的UPDATE ，不区分大小写
	updateReg2, _ := regexp.Compile("(?i)UPDATE\\s+.*?(\\bSET\\s+.*?)(?:\\s+WHERE\\s+.*|$)")

	//DELETE 正则替换
	if strings.HasPrefix(strings.ToLower(sql), "delete ") {
		return deleteReg.ReplaceAllString(sql, "SELECT  * ")
	}

	//UPDATE软删除
	if strings.HasPrefix(strings.ToLower(sql), "update ") {
		//替换掉UPDATE
		sql = updateReg2.ReplaceAllStringFunc(sql, func(s string) string {
			subMatches := updateReg2.FindStringSubmatch(s)
			if len(subMatches) > 0 {
				return strings.ReplaceAll(sql, subMatches[1], "")
			}
			return s
		})
		return updateReg1.ReplaceAllString(sql, "SELECT * FROM ")
	}
	return ""
}

func IsValidDataModel(model any, idKey string) (*DataModel, bool) {
	bts, err := json.Marshal(model)
	if err != nil {
		return nil, false
	}
	dataModel := DataModel(make(map[string]any))
	d := json.NewDecoder(bytes.NewReader(bts))
	d.UseNumber()

	if err := d.Decode(&dataModel); err != nil {
		return nil, false
	}
	value, ok := dataModel[idKey]
	if !ok {
		return nil, false
	}
	valueStr := fmt.Sprintf("%v", value)
	if valueStr == "" || valueStr == "0" {
		return nil, false
	}
	return &dataModel, true
}

func printModel(model any) string {
	return string(lo.T2(json.Marshal(model)).A)
}

func sqlOperation(sql string) string {
	deleteKey := findDeleteKey(sql)
	lowerSQL := strings.ToLower(sql)
	if strings.HasPrefix(lowerSQL, "delete") {
		return ChangeOptionDelete
	}
	if strings.HasPrefix(lowerSQL, "update") {
		updateReg, _ := regexp.Compile("set\\s+.*?" + deleteKey)
		if updateReg.Match([]byte(sql)) {
			return ChangeOptionDelete
		}
		return ChangeOptionUpdate
	}
	if strings.HasPrefix(lowerSQL, "insert") {
		return ChangeOptionCreate
	}
	return ChangeOptionSelect
}

func extractTableNameFromSQL(sql string) string {
	sql = strings.ReplaceAll(sql, "`", "") //去掉`的影响
	// 它假设表名紧跟在"FROM"（对于DELETE）或操作之后（对于INSERT/UPDATE），并且没有别名或复杂的子查询
	re := regexp.MustCompile(`(?i)(?:FROM|INSERT INTO|UPDATE) (\w+)`)
	matches := re.FindAllStringSubmatch(sql, -1)

	var tableNames []string
	for _, match := range matches {
		if len(match) > 1 {
			tableNames = append(tableNames, match[1])
		}
	}
	return tableNames[0]
}

func extractTableName(sql string, db *gorm.DB) string {
	if db.Statement != nil && db.Statement.Schema != nil {
		if db.Statement.Schema.Table != "" {
			return db.Statement.Table
		}
	}
	return extractTableNameFromSQL(sql)
}

func queryFromSQL(ctx context.Context, db *gorm.DB, sql string) ([]DataModel, error) {
	ds := make([]map[string]any, 0)
	if err := db.WithContext(ctx).Raw(sql).Scan(&ds).Error; err != nil {
		return nil, err
	}
	results := make([]DataModel, 0)
	for _, d := range ds {
		results = append(results, d)
	}
	return results, nil
}

func IsTx(db *gorm.DB) bool {
	_, ok := db.Statement.ConnPool.(*sql.Tx)
	return ok
}

func HasCommited(connPool gorm.ConnPool) bool {
	if tx, ok := connPool.(*sql.Tx); ok && tx != nil && !reflect.ValueOf(tx).IsNil() {
		fmt.Printf("tx pointer: %p", tx)
		vtx := reflect.ValueOf(tx)
		if vtx.Kind() == reflect.Pointer {
			vtx = vtx.Elem()
		}
		if vtx.Kind() != reflect.Struct {
			return false
		}
		field := vtx.FieldByName("done")
		if !field.IsValid() {
			return false
		}
		if field.Kind() != reflect.Int32 {
			return false
		}
		return field.Int() == 1
	}
	return false
}

func HasNotProcessed(before DataModel, after any) bool {
	bts1, _ := json.Marshal(before)
	bts2, _ := json.Marshal(after)
	return len(bts1) == len(bts2) && string(bts1) == string(bts2)
}

func PrintModel(model any) string {
	return string(lo.T2(json.Marshal(model)).A)
}

type QueryFromRawType interface {
	CallbackModel
}

func QueryFromRaw[T QueryFromRawType](ctx context.Context, db *gorm.DB, idValue string) (dest T, err error) {
	var callbackModel CallbackModel = dest
	tableName := callbackModel.TableName()
	sql := fmt.Sprintf("select * from `%s` where %s='%s'", tableName, callbackModel.UniqueKey(), idValue)
	if strings.Contains("dm8", db.Name()) || db.Name() == "oracle" {
		sql = fmt.Sprintf(`select * from "%s" where %s='%s'`, tableName, callbackModel.UniqueKey(), idValue)
	}
	err = db.WithContext(ctx).Raw(sql).Unscoped().Take(&dest).Error
	return dest, err
}

func QueryFromDB[T QueryFromRawType](ctx context.Context, db *gorm.DB, data DataModel) (T, error) {
	emptyObj := *new(T)
	id, ok := data[emptyObj.UniqueKey()]
	if !ok {
		return emptyObj, fmt.Errorf("invalid DataModel %v", PrintModel(data))
	}
	return QueryFromRaw[T](ctx, db, fmt.Sprintf("%v", id))
}

func CheckTxFinished(connPool gorm.ConnPool) bool {
	if connPool == nil {
		return true
	}
	tx, ok := connPool.(*sql.Tx)
	if !ok || tx == nil {
		return true
	}
	if reflect.ValueOf(tx).IsNil() {
		return true
	}
	//rawSQL := `select 1 from dual`
	////如果事务已经关闭或者回退，则无法执行下面的语句，说明事务结束
	//stmt, err := connPool.PrepareContext(context.Background(), rawSQL)
	//if err != nil {
	//	return errors.Is(err, sql.ErrTxDone)
	//}
	//defer stmt.Close()
	return false
}
