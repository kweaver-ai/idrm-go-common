package errorcode

import "encoding/json"

func init() {
	RegisterErrorCode(authServiceErrorMap)
}

const authServicePreCoder = "ServiceCall.AuthService."

const (
	// 获取策略列表失败
	ListPoliciesFailure = authServicePreCoder + "ListPoliciesFailure"
	// 资源授权申请的申请者缺少角色：应用开发者
	RequesterWithoutRoleApplicationDeveloper = authServicePreCoder + "RequesterWithoutRoleApplicationDeveloper"
	// 资源授权申请的申请者不是应用的应用开发者
	RequesterIsNotApplicationDeveloper = authServicePreCoder + "RequesterIsNotApplicationDeveloper"
	// 获取指标维度规则列表失败
	ListIndicatorDimensionalRulesFailure = authServicePreCoder + "ListIndicatorDimensionalRulesFailure"
)

var authServiceErrorMap = ErrorCode{
	ListPoliciesFailure: {
		Description: "获取策略列表失败",
	},
	RequesterWithoutRoleApplicationDeveloper: {
		Description: "申请者缺少角色：应用开发者",
	},
	RequesterIsNotApplicationDeveloper: {
		Description: "申请者不是应用的开发者",
	},
	ListIndicatorDimensionalRulesFailure: {
		Description: "获取指标维度规则列表失败",
	},
}

type ListPoliciesFailureDetail struct {
	// http response status code
	StatusCode int `json:"status_code,omitempty"`
	// http response body
	Body json.RawMessage `json:"body,omitempty"`
}

func NewListPoliciesFailure(statusCode int, body []byte) error {
	return Detail(ListPoliciesFailure, &ListPoliciesFailureDetail{StatusCode: statusCode, Body: body})
}

type ListIndicatorDimensionalRulesFailureDetail struct {
	// http response status code
	StatusCode int `json:"status_code,omitempty"`
	// http response body
	Body json.RawMessage `json:"body,omitempty"`
}

func NewListIndicatorDimensionalRulesFailure(statusCode int, body []byte) error {
	return Detail(ListPoliciesFailure, &ListIndicatorDimensionalRulesFailureDetail{StatusCode: statusCode, Body: body})
}
