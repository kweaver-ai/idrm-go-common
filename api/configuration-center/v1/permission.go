package v1

import (
	"github.com/google/uuid"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// Permission 代表权限
type Permission struct {
	// 元数据
	meta_v1.Metadata
	// 定义
	PermissionSpec
}

// PermissionSpec 定义一个权限
type PermissionSpec struct {
	// 名称，最大长度 128
	Name string `json:"name,omitempty"`
	// 分类，最大长度 128
	Category PermissionCategory `json:"category,omitempty"`
	// 描述，最大长度 300
	Description string `json:"description,omitempty"`
}

// PermissionCategory 代表权限分类
type PermissionCategory string

const (
	// 权限分类：基础权限
	PermissionCategoryBasicPermission PermissionCategory = "BasicPermission"
	// 权限分类：基础类
	PermissionCategoryBasic PermissionCategory = "Basic"
	// 权限分类：运营类
	PermissionCategoryOperation PermissionCategory = "Operation"
	// 权限分类：服务类
	PermissionCategoryService PermissionCategory = "Service"
)

// 权限范围
type Scope string

const (
	// 权限范围：全部
	ScopeAll Scope = "All"
	// 权限范围：当前部门
	ScopeCurrentDepartment Scope = "CurrentDepartment"
)

// PermissionList 代表权限列表
type PermissionList meta_v1.List[Permission]

// ScopeAndPermissions 代表权限范围和权限列表
type ScopeAndPermissions struct {
	// 权限范围
	Scope Scope `json:"scope,omitempty"`
	// 权限 ID 列表
	Permissions uuid.UUIDs `json:"permissions,omitempty"`
}

// ScopedPermission 代表指定影响范围的权限
type ScopedPermission struct {
	// 范围
	Scope Scope `json:"scope,omitempty"`
	// 权限
	Permission
}
