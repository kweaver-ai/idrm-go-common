package frontend

import (
	configuration_center_v1 "github.com/kweaver-ai/idrm-go-common/api/configuration-center/v1"
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// RoleGroup 代表角色组及其相关数据
type RoleGroup struct {
	// 元数据
	meta_v1.Metadata
	// 定义
	configuration_center_v1.RoleGroupSpec
	// 关联的角色
	Roles []configuration_center_v1.Role `json:"roles,omitempty"`
}

// RoleGroupList 代表角色组列表，及其关联的资源
type RoleGroupList meta_v1.List[RoleGroup]
