package v1

import (
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// RoleGroup 代表角色组
type RoleGroup struct {
	// 元数据
	meta_v1.Metadata
	// 定义
	RoleGroupSpec
}

// RoleGroupSpec 代表一个角色组的定义
type RoleGroupSpec struct {
	// 名称，最大长度 128
	Name string `json:"name,omitempty"`
	// 描述，最大长度 255
	Description string `json:"description,omitempty"`
}

// RoleGroupList 代表角色组列表
type RoleGroupList meta_v1.List[RoleGroup]

// RoleGroupListOptions 代表获取角色组列表接口的 query 参数
type RoleGroupListOptions struct {
	meta_v1.ListOptions
	// 关键字
	Keyword string `json:"keyword,omitempty" form:"keyword"`
	// 用户 ID 列表，非空时过滤指定用户关联的角色组
	UserIDs []string `json:"user_ids,omitempty" form:"user_ids"`
}

// RoleGroupRoleBinding 定义角色组、角色绑定关系
type RoleGroupRoleBinding struct {
	// 角色组 ID
	RoleGroupID string `json:"role_group_id,omitempty"`
	// 角色 ID
	RoleID string `json:"role_id,omitempty"`
}

// RoleGroupRoleBindingProcessing 代表对角色组、角色绑定关系的处理
type RoleGroupRoleBindingProcessing struct {
	RoleGroupRoleBinding
	// 期望状态
	State meta_v1.ProcessingState `json:"state,omitempty"`
}

// RoleGroupRoleBindingBatchProcessing 代表对角色组、角色绑定关系的批处理
type RoleGroupRoleBindingBatchProcessing struct {
	// 角色组 ID 列表，非空时作为默认值
	RoleGroupIDs []string `json:"role_group_ids,omitempty"`
	// 角色 ID 列表，非空时作为默认值
	RoleIDs []string `json:"role_ids,omitempty"`
	// 期望绑定关系是否存在，非空时作为默认值
	State meta_v1.ProcessingState `json:"state,omitempty"`
	// 角色组、角色绑定关系列表
	Bindings []RoleGroupRoleBindingProcessing `json:"bindings,omitempty"`
}

// RoleGroupNameCheck 检查角色组名称是否可以使用
type RoleGroupNameCheck struct {
	// 检查指定角色组是否可以使用这个名称
	// 未指定，检查新建角色组是否可以使用指定名称
	Id string `json:"id,omitempty" form:"id"`
	// 名称
	Name string `json:"name,omitempty" form:"name"`
}
