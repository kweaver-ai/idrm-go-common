package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/kweaver-ai/idrm-go-common/access_control"
	"github.com/kweaver-ai/idrm-go-common/rest/user_management"
)

type Middleware interface {
	TokenInterception() gin.HandlerFunc
	// ShouldTokenInterception 尝试 token interception，即使失败不终止
	// gin.Context 处理 request
	//
	// 经过此中间件的 gin.Context 支持：
	//
	//  - 通过 interception.AuthFromContext 获取 HTTP Header Authorization 用于身份认证
	//  - 通过 interception.BearerTokenFromContext 获取 Bearer Token 用于身份认证
	//  - 通过 interception.AuthServiceSubjectFromContext 获取 auth-service Subject 用于鉴权
	ShouldTokenInterception() gin.HandlerFunc
	AccessControl(resource access_control.Resource) gin.HandlerFunc
	MultipleAccessControl(resources ...access_control.Resource) gin.HandlerFunc
	AccessControlWithAccessType(accessType access_control.AccessType, resource access_control.Resource) gin.HandlerFunc
	AccessControlSkipTypeClient(resource access_control.Resource) gin.HandlerFunc
	// 记录审计日志
	AuditLogger() gin.HandlerFunc
	// 根据候选角色身份校验用户权限，exclude为true时用户拥有候选列表以外身份时才有权限
	PermissionControl(roleIDs []string, exclude bool) gin.HandlerFunc
	TokenPassThrough() gin.HandlerFunc
	MenuPermissionMarker() *MenuResourceMarkerGenerator
}

type User struct {
	ID       string                     `json:"id"`
	Name     string                     `json:"name"`
	UserType int                        `json:"user_type"`
	OrgInfos []*user_management.DepInfo `json:"org_info"`
}

const VirtualEngineApp = "af-virtual-engine-gateway"

// var PathResourceMap = make(map[string]string)
// var PathDisableList []string = []string{"/api/configuration-center/v1/apps"}
// var PathEnableList []string = []string{
// 	"/api/demand-management",
// 	"/api/task-center/v1",
// 	"/api/business-grooming/v1",
// 	"/api/data-subject/v1",
// 	"/api/standardization/v1",
// 	"/api/data-view/v1",
// 	"/api/indicator-management/v1",
// 	"/api/data-application-service/v1",
// 	"/data-application-gateway",
// 	"/api/data-catalog/v1",
// 	"/api/configuration-center/v1",
// }

// func init() {
// 	PathResourceMap["/api/demand-management"] = "demand_task"
// 	PathResourceMap["/api/task-center/v1"] = "demand_task"
// 	PathResourceMap["/api/business-grooming/v1"] = "business_grooming"
// 	PathResourceMap["/api/standardization/v1"] = "standardization"
// 	PathResourceMap["/api/data-subject/v1"] = "standardization"
// 	PathResourceMap["/api/data-view/v1"] = "resource_management"
// 	PathResourceMap["/api/indicator-management/v1"] = "resource_management"
// 	PathResourceMap["/api/data-application-service/v1"] = "resource_management"
// 	PathResourceMap["/data-application-gateway"] = "resource_management"
// 	PathResourceMap["/api/data-catalog/v1"] = "resource_management"
// 	PathResourceMap["/api/configuration-center/v1"] = "configuration_center"
// }
