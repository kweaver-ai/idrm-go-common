package v1

import meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"

// Policy 定义一条权限策略
type Policy struct {
	// 访问者
	Subject
	// 资源
	Object
	// 动作
	Action Action `json:"action,omitempty"`
	// 效果
	PolicyEffect PolicyEffect `json:"policy_effect,omitempty"`
	// 过期时间
	ExpiredAt *meta_v1.Time `json:"expired_at,omitempty"`
}

// PolicyListOptions 定义获取 Policy 列表的参数
type PolicyListOptions struct {
	// 访问者，非空时返回访问者属于指定值的 Policy 列表
	Subjects []Subject `json:"subjects,omitempty"`
	// 资源，非空时返回指定资源的 Policy 列表
	Objects []Object `json:"objects,omitempty"`
}
