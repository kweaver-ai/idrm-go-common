package auth_service

import (
	"context"

	authServiceV1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	metaV1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

type AuthServiceV1Interface interface {
	// 策略验证
	Enforce(ctx context.Context, requests []authServiceV1.EnforceRequest) ([]bool, error)
	// 获取指定的指标维度规则
	GetIndicatorDimensionalRule(ctx context.Context, id string) (*authServiceV1.IndicatorDimensionalRule, error)
	// 获取访问者拥有的资源
	ListSubjectObjects(ctx context.Context, opts *authServiceV1.SubjectObjectsListOptions) (*metaV1.List[authServiceV1.ObjectWithPermissions], error)
}

// AuthServiceInternalV1Interface  auth-service 内部接口的客户端
type AuthServiceInternalV1Interface interface {
	Enforce(ctx context.Context, requests []authServiceV1.EnforceRequest) (responses []bool, err error) // 策略验证
	RuleEnforce(ctx context.Context, arg *authServiceV1.RulePolicyEnforce) (effectResp *authServiceV1.RulePolicyEnforceEffect, err error)
	MenuResourceCheck(ctx context.Context, requests *MenuResourceCheckRequest) (responses *MenuResourceCheckResponse, err error) // 接口权限校验
	MenuResourceActions(ctx context.Context, userID, resourceID string) (*MenuResourceActionsResp, error)
	ListPolicies(ctx context.Context, opts *authServiceV1.PolicyListOptions) ([]authServiceV1.Policy, error)                                                         // 获取权限策略列表
	ListIndicatorDimensionalRules(ctx context.Context, opts *authServiceV1.IndicatorDimensionalRuleListOptions) (*authServiceV1.IndicatorDimensionalRuleList, error) // 获取指标维度规则列表
	GetIndicatorRules(ctx context.Context, ruleID ...string) (ds []*authServiceV1.IndicatorDimensionalRule, err error)
	GetRulesByIndicators(ctx context.Context, indicators ...string) (ds map[string][]string, err error)
	FilterPolicyHasExpiredObjects(ctx context.Context, objectID ...string) (ds []string, err error)
	QueryViewHasDWHDataAuthRequestForm(ctx context.Context, uid string, dataViewIDSlice []string) (map[string]int, error)
}
