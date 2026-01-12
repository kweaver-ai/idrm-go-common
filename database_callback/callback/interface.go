package callback

import (
	"context"
	"time"
)

type Transport interface {
	Process(ctx context.Context, model DataModel, tableName, operation string) (any, error) //处理下数据,也可能不处理
	Send(ctx context.Context, body any) error                                               //发送消息的方式
}

// DataModel 数据模型方便传播的
type DataModel map[string]any

type CallbackModel interface {
	TableName() string
	UniqueKey() string
}

type EntityMsgMessage struct {
	Header  EntityMsgMessageHeader  `json:"header"`
	Payload EntityMsgMessagePayload `json:"payload"`
}
type EntityMsgMessageHeader struct{}

type EntityMsgMessagePayload struct {
	Type    string `json:"type"`    //模块类型，什么模块的，lineage， 只能推荐
	Content any    `json:"content"` //具体的内容
}

type DefaultContent struct {
	Type      string    `json:"type"`       //实体变更的类型  ： create update delete
	TableName string    `json:"table_name"` //实体所在的表名称
	Entities  []any     `json:"entities"`   //携带的数据
	UpdatedAt time.Time `json:"updated_at"` //操作的时间
}

// DataLineageContent 数据血缘的结构
type DataLineageContent struct {
	Type      string `json:"type"`       //操作的类型：create update delete
	ClassName string `json:"class_name"` //血缘的实体种类：table field indicator
	Entities  []any  `json:"entities"`   //携带的数据
}
