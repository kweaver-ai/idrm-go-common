package entity_change

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/kweaver-ai/idrm-go-common/database_callback"
)

/*
--------说明：
AD的知识图谱需要实时增量构建，下面利用了gorm的callback，在新增，删除，更新的时候，发送提示到MQ
认知助手收到消息后，消费消息，通知AD开始增量构建
*/

const (
	ChangeOptionCreate = "create"
	ChangeOptionUpdate = "update"
	ChangeOptionDelete = "delete"
)

type EntityMsgMessage struct {
	Header  EntityMsgMessageHeader  `json:"header"`
	Payload EntityMsgMessagePayload `json:"payload"`
}
type EntityMsgMessageHeader struct{}
type EntityMsgMessagePayload struct {
	Type      string    `json:"type" binding:"required,oneof="create update delete"` //实体变更的类型
	Graph     string    `json:"graph"`                                               //图谱的名称
	Name      string    `json:"name" binding:"required"`                             //实体的名称
	Model     any       `json:"model" binding:"model"`                               //实体的所有信息
	UpdatedAt time.Time `json:"updated_at"`
}

// ChangeSender 变更消息发送处理逻辑
type ChangeSender struct {
	sender MessageSender
}

func NewChangeSender(s MessageSender) database_callback.Register {
	return &ChangeSender{
		sender: s,
	}
}

func NewEntityChangeCallback(s MessageSender) database_callback.EntityChangeCallback {
	return &ChangeSender{
		sender: s,
	}
}

// Remind 发送，这里的发送很简单，实际当数据很多的时候，比如批量插入，其实也只会一条消息。也就是和数据库的压力是一致的
// 所有也就去掉了使用channel逻辑
func (c *ChangeSender) Remind(tableName string, updateTime time.Time, option string) {
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

	graphs := c.sender.DestGraph(tableName)
	for i := range graphs {
		msg := &EntityMsgMessage{
			Header: EntityMsgMessageHeader{},
			Payload: EntityMsgMessagePayload{
				Type:      option,
				Graph:     graphs[i],
				Name:      tableName,
				Model:     struct{}{},
				UpdatedAt: updateTime,
			},
		}
		bs, _ := json.Marshal(msg)
		if err := c.sender.Send(bs); err != nil {
			fmt.Printf("send entity change %v to kafka error %v", msg, err.Error())
		}
	}
}

func (c *ChangeSender) IsEmpty() bool {
	return false
}
