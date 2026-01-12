package v2

import (
	configuration_center_v1 "github.com/kweaver-ai/idrm-go-common/api/configuration-center/v1"
	configuration_center_v1_frontend "github.com/kweaver-ai/idrm-go-common/api/configuration-center/v1/frontend"
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
	ParentDeps []configuration_center_v1_frontend.DepartmentInfo `json:"parent_deps,omitempty"`
	// 关联的角色
	Roles []Role `json:"roles,omitempty"`
}

// UserList 代表角色列表，及其关联的资源
type UserList meta_v1.List[User]

// Role 代表一个角色的定义
type Role struct {
	//角色ID uuid
	ID string `json:"id"`
	// 名称，最大长度 128
	Name string `json:"name"`
	// 描述，最大长度 255
	Description string `json:"description,omitempty"`
}
