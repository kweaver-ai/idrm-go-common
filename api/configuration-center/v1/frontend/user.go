package frontend

import (
	configuration_center_v1 "github.com/kweaver-ai/idrm-go-common/api/configuration-center/v1"
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
	meta_v1_frontend "github.com/kweaver-ai/idrm-go-common/api/meta/v1/frontend"
)

// User 代表用户及其关联的资源
type User struct {
	// 元数据
	meta_v1_frontend.MetadataWithOperator
	// 定义
	configuration_center_v1.UserSpec
	// 所属部门路径
	ParentDeps []DepartmentInfo `json:"parent_deps,omitempty"`
	// 关联的角色
	Roles []configuration_center_v1.Role `json:"roles,omitempty"`
	// 关联的角色组
	RoleGroups []RoleGroup `json:"role_groups,omitempty"`
	// 计算得到的权限，与用户权限、用户角色、用户角色组绑定有关
	Permissions []configuration_center_v1.ScopedPermission `json:"permissions,omitempty"`
}

// UserList 代表角色列表，及其关联的资源
type UserList meta_v1.List[User]

type DepartmentInfo struct {
	// 路径ID
	PathID string `json:"path_id"`
	// 路径
	Path string `json:"path"`
}
