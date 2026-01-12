package frontend

import (
	asset_portal_v1 "github.com/kweaver-ai/idrm-go-common/api/asset_portal/v1"
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// Notification 代表用户收到的一条通知
type Notification struct {
	ID string `json:"id,omitempty"`
	// 发送时间
	Time meta_v1.Time `json:"time,omitempty"`
	// 收件人 ID
	RecipientID string `json:"recipient_id,omitempty"`
	// 通知的内容
	Message string `json:"message,omitempty"`
	// 是否已读
	Read bool `json:"read,omitempty"`
	// 收到这个通知的理由
	Reason asset_portal_v1.Reason `json:"reason,omitempty"`
	// 工单，如果 reason 是 DataQualityWorkOrder
	WorkOrder *WorkOrder `json:"work_order,omitempty"`
}

// WorkOrder 代表通知中引用的工单
type WorkOrder struct {
	ID string `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 编号
	Code string `json:"code,omitempty"`
	// 截止日期
	Deadline meta_v1.Time `json:"deadline,omitempty"`
}

// NotificationList 代表通知列表
type NotificationList meta_v1.List[Notification]
