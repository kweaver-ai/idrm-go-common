package callback

import (
	"context"
	"fmt"
	"time"
)

type Transporter struct {
	stageCache *StageCache
	rs         map[string][]TransportModule
	ts         map[string]CallbackModel
}

type TransportModule struct {
	Name string
	Transport
}

func NewTransporter() Transporter {
	return Transporter{
		rs:         make(map[string][]TransportModule),
		stageCache: NewStageCache(),
		ts:         make(map[string]CallbackModel),
	}
}

func (t *Transporter) HasRegistered(name string) bool {
	_, ok := t.rs[name]
	return ok
}

func (t *Transporter) UniqueKey(name string) string {
	callbackModel, ok := t.ts[name]
	if ok {
		return callbackModel.UniqueKey()
	}
	fmt.Printf("invalid callback table name %v\n", name)
	return "id"
}

func (t *Transporter) Register(callbackModel CallbackModel, rs ...TransportModule) {
	name := callbackModel.TableName()
	t.rs[name] = append(t.rs[name], rs...)
	t.ts[name] = callbackModel
}

// GetCache 从缓存中获取
func (t *Transporter) GetCache(statementID string) []DataModel {
	models := t.stageCache.Cut(statementID)
	results := make([]DataModel, 0)
	for _, model := range models {
		data, ok := model.(DataModel)
		if ok {
			results = append(results, data)
		}
	}
	return results
}

// SetCache 缓存数据
func (t *Transporter) SetCache(statementID string, dataModels []DataModel) {
	ds := make([]any, 0)
	for _, model := range dataModels {
		ds = append(ds, model)
	}
	t.stageCache.Set(statementID, ds)
}

// Handle 处理下数据，发送
func (t *Transporter) Handle(r *Record) {
	models := make([]DataModel, 0)
	//获取model
	models = r.DataModels
	if len(models) == 0 && r.hasStage() {
		models = t.GetCache(r.StatementID)
	}
	if len(models) <= 0 {
		return
	}
	//如果是删除操作和raw操作，先stage再send
	if r.needStage() {
		t.SetCache(r.StatementID, models)
		return
	}
	t.Send(context.Background(), r, models...)
}

func (t *Transporter) Send(ctx context.Context, r *Record, models ...DataModel) {
	ts := t.rs[r.TableName]
	for _, transport := range ts {
		for _, model := range models {
			//send on model
			msg, err := content(ctx, transport, r, model)
			if err != nil {
				fmt.Printf("query callback model error,%v\n", err)
				return
			}
			nowTimeFmt := time.Now().Format("2006-01-02 15:04:05")
			//fmt.Printf("%v: %s: callback mode send :%s\n", nowTimeFmt, transport.Name, printModel(msg))
			err = transport.Send(ctx, msg)
			if err != nil {
				fmt.Printf("%v:  %s:  callback mode send error %v\n", nowTimeFmt, transport.Name, err.Error())
			}
		}
	}
}

func content(ctx context.Context, transport TransportModule, r *Record, model DataModel) (c *EntityMsgMessage, err error) {
	c = &EntityMsgMessage{
		Header: EntityMsgMessageHeader{},
		Payload: EntityMsgMessagePayload{
			Type: transport.Name,
		},
	}
	processedModel, err := transport.Process(ctx, model, r.TableName, r.Operation)
	if err != nil {
		return nil, fmt.Errorf("this is not a callback model: %v", printModel(model))
	}
	//如果处理了，使用处理后的结构，如果没有处理，使用默认结构
	if HasNotProcessed(model, processedModel) {
		c.Payload.Content = DefaultContent{
			Type:      r.Operation,
			TableName: r.TableName,
			Entities:  []any{model},
			UpdatedAt: r.UpdatedAt,
		}
		return c, nil
	}

	c.Payload.Content = processedModel
	return c, nil
}
