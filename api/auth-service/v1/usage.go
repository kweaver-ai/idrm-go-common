package v1

// AuthorizingRequestUsage 定义授权申请的用途
type AuthorizingRequestUsage string

const (
	// 授权申请
	AuthorizingRequestAuthorizingRequest AuthorizingRequestUsage = "AuthorizingRequest"
	// 需求管理
	AuthorizingRequestDemandManagement AuthorizingRequestUsage = "DemandManagement"
)
