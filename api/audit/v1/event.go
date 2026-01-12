package v1

import (
	"encoding/json"
	"time"
)

type EventList struct {
	Entries    []Event `json:"entries,omitempty"`
	TotalCount int     `json:"total_count,omitempty"`
}

// Event 定义一次审计事件
type Event struct {
	// 审计事件发生的事件，精度：毫秒
	Timestamp time.Time `json:"timestamp,omitempty"`
	// 审计事件的级别
	Level Level `json:"level,omitempty"`
	// 操作描述
	Description string `json:"description,omitempty"`
	// Authenticated user or application information.
	Operator Operator `json:"operator,omitempty"`
	// 操作
	Operation Operation `json:"operation,omitempty"`
	// 被操作资源的详情
	Detail json.RawMessage `json:"detail,omitempty"`
}

// Event 记入消息队列 Kafka 时，使用这个 topic
const KafkaTopic = "af.audit-log"
