package v1

import (
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// SubjectPolicy 描述了哪个操作者，执行哪些动作
type SubjectPolicy struct {
	// 操作者
	Subject `json:",inline"`
	// 动作列表
	Actions []Action `json:"actions,omitempty"`
	// 过期时间
	ExpiredAt *meta_v1.Time `json:"expired_at,omitempty"`
}
