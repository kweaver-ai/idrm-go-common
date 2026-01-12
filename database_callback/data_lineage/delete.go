package data_lineage

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/kweaver-ai/idrm-go-common/util"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func (c *LineageCallback) StageDeleteEntity(ctx context.Context, db *gorm.DB) {
	// gorm的使用主要有下面的几种方式：
	//1. model, models
	if db.Statement.ReflectValue.Kind() == reflect.Slice {
		for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
			value := db.Statement.ReflectValue.Index(0)
			c.stageModel(value.Interface())
		}
	} else {
		c.stageModel(db.Statement.Model)
	}

	//2. raw SQL
	sql := db.Statement.ToSQL(func(tx *gorm.DB) *gorm.DB { return db })
	if sql == "" {
		return
	}
	tableName := ""
	if db.Statement != nil && db.Statement.Schema != nil {
		tableName = db.Statement.Schema.Table
	}
	if tableName == "" {
		fmt.Printf("invalid schema in statement, can not find table name")
		return
	}
	models, err := c.processor.GetByRawSQL(ctx, tableName, toQuerySQL(sql))
	if err != nil {
		fmt.Printf("query error %v", err.Error())
		return
	}
	stageObj := &CacheModel{
		ID:         fmt.Sprintf("%p", db.Statement),
		CreateTime: time.Now().Unix(),
		Models:     models,
	}
	c.stage[stageObj.ID] = stageObj
}

func (c *LineageCallback) SendDeleteEntity(ctx context.Context, db *gorm.DB) {
	if db.Statement.ReflectValue.Kind() == reflect.Slice {
		for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
			value := db.Statement.ReflectValue.Index(0)
			c.sendStage(ctx, value.Interface())
		}
	} else {
		c.sendStage(ctx, db.Statement.Model)
	}
	//使用statement的指针取出model
	models := c.Cut(fmt.Sprintf("%p", db.Statement))
	if len(models) <= 0 {
		return
	}
	c.SendModels(ctx, ChangeOptionDelete, models...)
}

func (c *LineageCallback) sendStage(ctx context.Context, model any) {
	lineageOpt, ok := IsValidCallbackModel(model)
	if !ok {
		fmt.Printf("this is not a lineage model:%s\n", lo.T2(json.Marshal(model)).A)
		return
	}
	key := util.MD5(lineageOpt.UniqueID())
	models := c.Cut(key)
	if len(models) <= 0 {
		fmt.Printf("this is not a valid lineage model key :%s\n", lo.T2(json.Marshal(model)).A)
		return
	}
	c.SendModels(ctx, ChangeOptionDelete, models...)
}

func (c *LineageCallback) stageModel(model any) {
	lineageOpt, ok := IsValidCallbackModel(model)
	if !ok {
		fmt.Printf("this is not a lineage model:%s\n", lo.T2(json.Marshal(model)).A)
		return
	}
	stageObj := &CacheModel{
		ID:         lineageOpt.UniqueID(),
		CreateTime: time.Now().Unix(),
		Models:     []any{model},
	}
	c.stage[stageObj.ID] = stageObj
}
