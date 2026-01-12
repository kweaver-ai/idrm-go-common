package v1

import (
	"net/url"
	"strconv"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
	"github.com/kweaver-ai/idrm-go-common/util/ptr"
)

// Notification 代表用户收到的一条消息通知
type Notification struct {
	meta_v1.Metadata `json:"metadata,omitempty"`
	// 消息通知的内容
	Spec NotificationSpec `json:"spec,omitempty"`
	// 消息通知的状态
	Status NotificationStatus `json:"status,omitempty"`
}

// NotificationSpec 代表消息通知的内容
type NotificationSpec struct {
	// 收件人 ID
	RecipientID string `json:"recipient_id,omitempty"`
	// 用户收到这个消息通知的理由，例如：数据质量工单告警
	Reason Reason `json:"reason,omitempty"`
	// 工单的 ID，通知的理由是数据质量工单告警时有值
	WorkOrderID string `json:"work_order_id,omitempty"`
	// 通知的内容
	Message string `json:"message,omitempty"`
}

// 用户收到消息通知的理由，例如：数据质量工单告警
type Reason string

const (
	// 数据质量工单告警
	NotificationReasonDataQualityWorkOrderAlarm Reason = "DataQualityWorkOrderAlarm"
)

// NotificationStatus 代表消息通知的状态
type NotificationStatus struct {
	// 是否已读
	Read bool `json:"read,omitempty"`
}

// NotificationList 代表通知列表
type NotificationList meta_v1.List[Notification]

// NotificationListOptions 代表获取通知列表的选项
type NotificationListOptions struct {
	meta_v1.ListOptions
	// 根据收件人 ID 过滤，空代表不过滤
	// RecipientID string `json:"recipient_id,omitempty"`
	// 根据通知是否已读过滤，空代表不过滤
	Read *bool `json:"read,omitempty"`
}

func (opts *NotificationListOptions) UnmarshalQuery(data url.Values) (err error) {
	if err = opts.ListOptions.UnmarshalQuery(data); err != nil {
		return
	}
	for k, values := range data {
		for _, v := range values {
			switch k {
			case "read":
				var r bool
				if r, err = strconv.ParseBool(v); err != nil {
					return
				}
				opts.Read = ptr.To(r)
			default:
				continue
			}
		}
	}
	return
}
