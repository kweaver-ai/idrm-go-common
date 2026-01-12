package data_lineage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/kweaver-ai/idrm-go-common/database_callback"
	"github.com/samber/lo"
)

const (
	ContentTypeLineage          = "lineage"
	ContentTypeSQL              = "sql"
	LineageCacheRefreshInterval = int64(3600)
	LineageCommitWaitInterval   = 5 * time.Second
)

const (
	ChangeOptionCreate = "insert"
	ChangeOptionUpdate = "update"
	ChangeOptionDelete = "delete"
)

// Processor 数据查询补充，发送接口，外部需要实现的唯一接口
type Processor interface {
	//QueryByID  根据ID查询出需要发送的对象, opt 是需要的具体的操作
	QueryByID(ctx context.Context, model CallbackModel, opt string) (any, error)
	//GetByRawSQL 根据给定的SQL语句查询出具体的对象
	GetByRawSQL(ctx context.Context, modelName string, sql string) ([]any, error)
	//Send 发送对象将对象发送到指定地方，可以使用MQ，可以使用http
	Send(ctx context.Context, body any) error
}

// CallbackModel 血缘监控作识别的对象
type CallbackModel interface {
	//EntityName 返回实体名称l, 格式是图谱
	EntityName() string
	//UniqueID  返回能唯一标识该对象的ID
	UniqueID() string
	//TableName  返回对象所在的表的名称
	TableName() string
}

// LineageCallback 变更消息发送处理逻辑
type LineageCallback struct {
	processor       Processor
	stage           map[string]*CacheModel
	lock            *sync.RWMutex
	lastRefreshTime int64
}

func NewLineageCallback(p Processor) database_callback.DataLineageCallback {
	return &LineageCallback{
		processor:       p,
		stage:           make(map[string]*CacheModel),
		lock:            &sync.RWMutex{},
		lastRefreshTime: time.Now().Unix(),
	}
}

func (c *LineageCallback) SendModels(ctx context.Context, option string, models ...any) {
	time.Sleep(LineageCommitWaitInterval)
	for _, model := range models {
		c.send(ctx, option, model)
	}
}

func (c *LineageCallback) Send(ctx context.Context, option string, model any) {
	time.Sleep(LineageCommitWaitInterval)
	c.send(ctx, option, model)
}

// Send  发送血缘数据
func (c *LineageCallback) send(ctx context.Context, option string, model any) {
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

	data, err := c.content(ctx, model, option)
	if err != nil {
		log.Printf("query lineage error,%v\n", err)
		return
	}

	msg := &EntityMsgMessage{
		Header: EntityMsgMessageHeader{},
		Payload: EntityMsgMessagePayload{
			Type:    ContentTypeLineage,
			Content: *data,
		},
	}
	if err := c.processor.Send(ctx, msg); err != nil {
		fmt.Printf("lineage send %s\n error %v\n", string(lo.T2(json.Marshal(msg)).A), err.Error())
	} else {
		fmt.Printf("lineage send :%s\n", string(lo.T2(json.Marshal(msg)).A))
	}
}

func (c *LineageCallback) content(ctx context.Context, model any, option string) (*Content, error) {
	lineageOpt, ok := IsValidCallbackModel(model)
	if !ok {
		return nil, fmt.Errorf("this is not a lineage model")
	}
	d, err := c.processor.QueryByID(ctx, lineageOpt, option)
	if err != nil {
		fmt.Printf("queryByID: error %v\n", err.Error())
		return nil, err
	}
	return &Content{
		Type:      option,
		ClassName: lineageOpt.EntityName(),
		Entities:  []any{d},
	}, nil
}

func toQuerySQL(sql string) string {
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

func IsValidCallbackModel(model any) (CallbackModel, bool) {
	object, ok := isCallbackModel(model)
	if !ok {
		return object, ok
	}
	if object.UniqueID() == "" {
		fmt.Errorf("this is lineage model has a empty uniqueID: %v", printModel(model))
		return object, false
	}
	return object, ok
}

func isCallbackModel(model any) (CallbackModel, bool) {
	if model == nil {
		fmt.Printf("this is not a empty lineage model:%s\n", printModel(model))
		return nil, false
	}
	lineageOpt, ok := model.(CallbackModel)
	if ok {
		return lineageOpt, true
	}
	//如果不是，尝试下面的几种情况
	mv := reflect.ValueOf(model)
	mt := reflect.TypeOf(model)
	if mv.Kind() == reflect.Struct {
		emptyModel, ok := (reflect.New(mt).Interface()).(CallbackModel)
		if ok {
			return setPoniterValue(model, emptyModel)
		}
	}
	for mv.Kind() == reflect.Pointer {
		if mv.Kind() == reflect.Pointer {
			obj := mv.Interface()
			object, ok := obj.(CallbackModel)
			if ok {
				return object, true
			}
		}
		if mv.Kind() == reflect.Struct {
			emptyModel, ok := (reflect.New(mt).Interface()).(CallbackModel)
			if ok {
				return setPoniterValue(model, emptyModel)
			}
		}
		//如果到了struct，那就真的不是了，放弃吧
		if mv.Kind() == reflect.Struct {
			break
		}
		mv = mv.Elem()
	}
	fmt.Printf("this is not a lineage model:%s\n", printModel(model))
	return nil, false
}

func setPoniterValue(obj any, emptyModel CallbackModel) (CallbackModel, bool) {
	bts, _ := json.Marshal(obj)
	if err := json.Unmarshal(bts, &emptyModel); err != nil {
		return emptyModel, false
	}
	return emptyModel, true
}

func printModel(model any) string {
	return string(lo.T2(json.Marshal(model)).A)
}
