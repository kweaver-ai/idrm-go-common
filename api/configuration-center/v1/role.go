package v1

import (
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// Role 代表角色
//

type Role struct {
	// 元数据
	meta_v1.MetadataWithOperator
	// 角色的定义
	RoleSpec
}

// RoleSpec 代表一个角色的定义
type RoleSpec struct {
	// 名称，最大长度 128
	Name string `json:"name,omitempty"`
	// 类型：内置角色、自定义角色
	Type RoleType `json:"type,omitempty"`
	// 描述，最大长度 255
	Description string `json:"description,omitempty"`
	// 颜色
	//
	// TODO: 补充定义，如何定义一个颜色，ID or RGB
	Color string `json:"color,omitempty"`
	// 权限范围，默认为当前部门
	Scope Scope `json:"scope,omitempty"`
	// 图标
	//
	// Deprecated: 兼容旧版本 API
	Icon string `json:"icon,omitempty"`
	// 是否是系统默认的角色，1 代表是，0 代表不是
	//
	// Deprecated: 兼容旧版本 API
	System int `json:"System,omitempty"`
}

// RoleType 代表角色类型
type RoleType string

const (
	// RoleTypeInternal 代表内置角色。内置角色无法被创建、更新、删除。内置角色与
	// 权限的关联无法添加、删除。
	RoleTypeInternal RoleType = "Internal"
	// RoleTypeCustom 代表自定义角色
	RoleTypeCustom RoleType = "Custom"
)

// RoleList 代表角色列表
type RoleList meta_v1.List[Role]

// RoleListOptions 代表获取角色列表接口的 query 参数
type RoleListOptions struct {
	meta_v1.ListOptions
	// 关键字
	Keyword string `json:"keyword,omitempty" form:"keyword"`
	// 类型
	Type RoleType `json:"type,omitempty" form:"type"`
	// 角色组
	RoleGroupID string `json:"role_group_id,omitempty" form:"role_group_id"`
	// 用户 ID 列表，非空时过滤指定用户关联的角色
	UserIDs []string `json:"user_ids,omitempty" form:"user_ids"`
}

// RoleNameCheck 检查角色名称是否可以使用
type RoleNameCheck struct {
	// 检查指定角色是否可以使用这个名称
	// 未指定，检查新建角色是否可以使用指定名称
	Id string `json:"id,omitempty" form:"id"`
	// 名称
	Name string `json:"name,omitempty" form:"name"`
}
