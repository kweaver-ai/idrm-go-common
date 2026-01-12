package frontend

import (
	configuration_center_v1 "github.com/kweaver-ai/idrm-go-common/api/configuration-center/v1"
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
	meta_v1_front "github.com/kweaver-ai/idrm-go-common/api/meta/v1/frontend"
)

// Role 代表角色，及其关联的资源
//

type Role struct {
	// 元数据
	meta_v1_front.MetadataWithOperator
	// 定义
	configuration_center_v1.RoleSpec
	// 关联的权限
	Permissions []configuration_center_v1.Permission `json:"permissions,omitempty"`
}

// RoleList 代表角色列表，及其关联的资源
type RoleList meta_v1.List[Role]
