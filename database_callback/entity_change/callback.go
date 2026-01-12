package entity_change

import (
	"gorm.io/gorm"
	"time"
)

// SendEntityCreateMsg 发送实体创建消息
func (c *ChangeSender) SendEntityCreateMsg(db *gorm.DB) {
	c.callbackEntry(db, ChangeOptionCreate)
}

// SendEntityUpdateMsg 发送实体更新消息
func (c *ChangeSender) SendEntityUpdateMsg(db *gorm.DB) {
	c.callbackEntry(db, ChangeOptionUpdate)
}

// SendEntityDeleteMsg 发送实体删除消息
func (c *ChangeSender) SendEntityDeleteMsg(db *gorm.DB) {
	c.callbackEntry(db, ChangeOptionDelete)
}

func (c *ChangeSender) callbackEntry(db *gorm.DB, op string) {
	if db.Statement == nil {
		return
	}
	if len(c.sender.DestGraph(db.Statement.Table)) <= 0 {
		return
	}
	if db.Statement != nil && db.Statement.Schema != nil {
		go c.Remind(db.Statement.Table, time.Now(), op)
	}
}

// RegisterCallback 注册回调
func (c *ChangeSender) RegisterCallback(client *gorm.DB) {
	client.Callback().Create().Register("create_msg:after_save", c.SendEntityCreateMsg)
	client.Callback().Update().Register("update_msg:after_update", c.SendEntityUpdateMsg)
	client.Callback().Delete().Register("delete_msg:after_delete", c.SendEntityDeleteMsg)
	client.Callback().Raw().Register("create_msg:after_save", c.SendEntityUpdateMsg)
}
