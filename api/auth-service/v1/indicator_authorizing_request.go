package v1

import (
	"time"
)

// IndicatorAuthorizingRequest 定义指标授权申请
type IndicatorAuthorizingRequest struct {
	// ID
	ID string `json:"id,omitempty"`
	// 创建时间
	CreationTimestamp time.Time `json:"creation_timestamp,omitempty"`
	// 申请的配置、定义。描述了对哪个指标，为哪些用户申请哪些动作的授权。
	Spec IndicatorAuthorizingRequestSpec `json:"spec,omitempty"`
	// 申请的状态
	Status IndicatorAuthorizingRequestStatus `json:"status,omitempty"`
}

// IndicatorAuthorizingRequestWithReferences 定义指标授权申请及其引用的资源
type IndicatorAuthorizingRequestWithReferences struct {
	IndicatorAuthorizingRequest `json:",inline"`
	// 指标授权索引用的资源列表，包括用户、部门等
	References []ReferenceSource `json:"references,omitempty"`
}

// IndicatorAuthorizingRequestSpec 描述了逻辑视图授权申请是为哪些用户申请哪个逻
// 辑视图、行列规则，执行哪些动作的授权。
type IndicatorAuthorizingRequestSpec struct {
	// 指标的 ID
	ID string `json:"id,omitempty"`
	// 指标维度规则
	Rule *IndicatorDimensionalRuleAuthorizingRequestSpec `json:"rule,omitempty"`
	// 对指标、指标维度规则所申请的授权列表。如果只申请行列规则（子视图），则不需要配置此字段。
	Policies []SubjectPolicy `json:"policies,omitempty"`
	// 创建逻辑视图授权申请的用户的 ID
	RequesterID string `json:"requester_id,omitempty"`
	// 发起授权申请的原因
	Reason string `json:"reason,omitempty"`
}

// 指标维度规则授权申请
type IndicatorDimensionalRuleAuthorizingRequestSpec struct {
	// 指标维度规则 ID，申请被批准后将授权这个指标维度规则。与 Template 互斥
	ID string `json:"id,omitempty"`
	// 指标维度规则模板，申请被批准后将按照这个模板创建维度规则并授权。与 ID 互
	// 斥。
	Template *IndicatorDimensionalRuleSpec `json:"template,omitempty"`
}

// IndicatorAuthorizingRequestStatus 描述了逻辑视图授权申请当前的状态，处于当前
// 状态的原因等。
type IndicatorAuthorizingRequestStatus struct {
	// 逻辑视图授权申请所处的阶段
	Phase IndicatorAuthorizingRequestPhase `json:"phase,omitempty"`
	// 逻辑视图授权申请处于当前阶段的原因，人类可读
	Message string `json:"message,omitempty"`
}

// IndicatorAuthorizingRequestPhase 定义 IndicatorAuthorizingRequest 在其生命周
// 期中所处的阶段
type IndicatorAuthorizingRequestPhase string

const (
	// 审批中
	IndicatorAuthorizingRequestAuditing IndicatorAuthorizingRequestPhase = "Auditing"
	// 申请被拒绝
	IndicatorAuthorizingRequestRejected IndicatorAuthorizingRequestPhase = "Rejected"
	// 申请被允许
	IndicatorAuthorizingRequestApproved IndicatorAuthorizingRequestPhase = "Approved"
	// 申请被发起者撤回
	IndicatorAuthorizingRequestUndone IndicatorAuthorizingRequestPhase = "Undone"
	// 失败。创建资源失败，或应用权限策略失败
	IndicatorAuthorizingRequestFailed IndicatorAuthorizingRequestPhase = "Failed"
	// 完成
	IndicatorAuthorizingRequestCompleted IndicatorAuthorizingRequestPhase = "Completed"
)
