package v1

import meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"

// AlarmRule 告警规则
type AlarmRule struct {
	// 告警规则ID
	ID string `json:"id,omitempty"`
	// 规则类型
	Type AlarmRuleType `json:"type,omitempty"`
	// 截止告警时间
	DeadlineTime meta_v1.Time `json:"deadline_time,omitempty"`
	// 截止告警内容
	DeadlineReminder string `json:"deadline_reminder,omitempty"`
	// 提前告警时间
	BeforehandTime meta_v1.Time `json:"beforehand_time,omitempty"`
	// 提前告警内容
	BeforehandReminder string `json:"beforehand_reminder,omitempty"`
	// 更新时间
	UpdatedAt meta_v1.Time `json:"updated_at,omitempty"`
	// 更新用户ID
	UpdatedBy string `json:"updated_by,omitempty"`
}

type AlarmRuleType string

const (
	// 数据质量
	AlarmRuleDataQuality AlarmRuleType = "data_quality"
)
