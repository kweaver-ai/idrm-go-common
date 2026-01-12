package data_lineage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"reflect"
)

// SendLineageCreateMsg 发送实体新建消息
func (c *LineageCallback) SendLineageCreateMsg(db *gorm.DB) {
	if db.Statement != nil && db.Statement.Schema != nil {
		if db.Statement.ReflectValue.Kind() == reflect.Slice {
			var models []any
			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				value := db.Statement.ReflectValue.Index(i)
				models = append(models, value.Interface())
			}
			go c.SendModels(context.Background(), ChangeOptionCreate, models...)
			return

		}
		//如果是struct，处理成指针
		if db.Statement.ReflectValue.Kind() == reflect.Struct {
			go c.Send(context.Background(), ChangeOptionCreate, db.Statement.ReflectValue.Interface())
		}
		//按照gorm的create源代码，走到这里应该是错的类型，不存在指针类型的
		fmt.Printf("%v, is not a valid callback model", string(lo.T2(json.Marshal(db.Statement.ReflectValue.Interface())).A))
	}
}

// SendLineageUpdateMsg 发送实体更新消息
func (c *LineageCallback) SendLineageUpdateMsg(db *gorm.DB) {
	if db.Statement != nil && db.Statement.Schema != nil {
		if db.Statement.ReflectValue.Kind() == reflect.Slice {
			var models []any
			for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
				value := db.Statement.ReflectValue.Index(i)
				models = append(models, value.Interface())
			}
			go c.SendModels(context.Background(), ChangeOptionUpdate, models...)
			return
		}
		//如果是struct，处理成指针
		if db.Statement.ReflectValue.Kind() == reflect.Struct {
			go c.Send(context.Background(), ChangeOptionUpdate, db.Statement.ReflectValue.Interface())
		}
		//按照gorm的create源代码，走到这里应该是错的类型，不存在指针类型的
		fmt.Printf("%v, is not a valid callback model", string(lo.T2(json.Marshal(db.Statement.ReflectValue.Interface())).A))
	}
}

// StageLineageDeleteMsg 获取实体删除消息
func (c *LineageCallback) StageLineageDeleteMsg(db *gorm.DB) {
	if db.Statement != nil && db.Statement.Model != nil {
		if _, ok := db.Statement.Model.(CallbackModel); !ok {
			return
		}
		go c.StageDeleteEntity(context.Background(), db)
	}
}

// SendLineageDeleteMsg 发送实体删除消息
func (c *LineageCallback) SendLineageDeleteMsg(db *gorm.DB) {
	if db.Statement != nil && db.Statement.Model != nil {
		if _, ok := db.Statement.Model.(CallbackModel); !ok {
			return
		}
		go c.SendDeleteEntity(context.Background(), db)
	}
}

// RegisterCallback 注册回调
func (c *LineageCallback) RegisterCallback(client *gorm.DB) {
	client.Callback().Create().After("gorm:create").Register("create_lineage_msg:after_save", c.SendLineageCreateMsg)
	client.Callback().Update().After("gorm:update").Register("update_lineage_msg:after_update", c.SendLineageUpdateMsg)
	client.Callback().Delete().Before("gorm:delete").Register("delete_lineage_msg:before_delete", c.StageLineageDeleteMsg) //删除消息，删除之前需要缓存下消息，然后在成功删除之后再发送消息
	client.Callback().Delete().After("gorm:delete").Register("delete_lineage_msg:after_delete", c.SendLineageDeleteMsg)
}

func (c *LineageCallback) IsEmpty() bool {
	return false
}
