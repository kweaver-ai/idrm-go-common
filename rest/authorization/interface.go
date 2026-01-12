package authorization

import "github.com/samber/lo"

type Driven interface {
	Resource
	Role
	Authorization
	Obligation
}

const (
	ACCESSOR_TYPE_USER       = "user"
	ACCESSOR_TYPE_DEPARTMENT = "department"
	ACCESSOR_TYPE_GROUP      = "group"
	ACCESSOR_TYPE_ROLE       = "role"
	ACCESSOR_TYPE_APP        = "app"
)

const (
	OBLIGATION_TYPE_IDRM_DATA = "idrm_data_obligation_types" //义务类型名称
)

const (
	INCLUDE_OPERATION_OBLIGATIONS = "operation_obligations" //义务
	INCLUDE_OBLIGATION_TYPES      = "obligation_types"      //义务类型
)

// ResourceType 资源类型
const (
	RESOURCE_TYPE_MENUS       = "idrm_menus"
	DATA_VIEW_RESOURCE_NAME   = "data_view"
	SUB_VIEW_RESOURCE_NAME    = "data_view_row_column_rule"
	API_RESOURCE_NAME         = "idrm_api"
	SUB_SERVICE_RESOURCE_NAME = "idrm_api_row_rule"
)

// InnerSystemRoles ISF 内置系统角色
var InnerSystemRoles = []string{
	"7dcfcc9c-ad02-11e8-aa06-000c29358ad6",
	"d2bd2082-ad03-11e8-aa06-000c29358ad6",
	"d8998f72-ad03-11e8-aa06-000c29358ad6",
	"def246f2-ad03-11e8-aa06-000c29358ad6",
	"e63e1c88-ad03-11e8-aa06-000c29358ad6",
	"f06ac18e-ad03-11e8-aa06-000c29358ad6",
}

const (
	ROLE_DATA_MANAGER = "00990824-4bf7-11f0-8fa7-865d5643e61f" //数据管理员
	ROLE_AI_MANAGER   = "3fb94948-5169-11f0-b662-3a7bdba2913f" //AI管理员
	ROLE_APP_MANAGER  = "1572fb82-526f-11f0-bde6-e674ec8dde71" //应用管理员
)

// InnerBusinessRoles ISF 内置业务角色
var InnerBusinessRoles = []string{
	ROLE_DATA_MANAGER,
	ROLE_AI_MANAGER,
	ROLE_APP_MANAGER,
}

// innerBusinessRoleMap 内置角色映射
var innerBusinessRoleMap = lo.SliceToMap(InnerBusinessRoles, func(item string) (string, bool) {
	return item, true
})

func IsInnerBusinessRole(role string) bool {
	return innerBusinessRoleMap[role]
}

func HasInnerBusinessRole(roles ...string) bool {
	for _, role := range roles {
		if IsInnerBusinessRole(role) {
			return true
		}
	}
	return false
}
