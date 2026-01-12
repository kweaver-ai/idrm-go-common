package v1

import meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"

// PermissionResource 代表权限资源
//
// 例如：
//   - 组织架构
//   - 用户信息
//   - 业务域
//

type PermissionResource struct {
	meta_v1.Metadata

	// 名称
	Name string `json:"name,omitempty"`
	// 范围
	Scope PermissionResourceScope `json:"scope,omitempty"`
}

// PermissionResourceScope 代表权限资源的生效范围，全局生效还是部门内生效
type PermissionResourceScope string

const (
	// 全局
	PermissionResourceAll PermissionResourceScope = "All"
	// 部门
	PermissionResourceDepartment PermissionResourceScope = "Department"
)

// 预定义的权限资源 ID
const (
	// 仅用于测试：组织架构
	PermissionResourceID_组织架构 = "01964c1d-fd4c-7711-b413-e7da5035ccdf"
	// 仅用于测试：角色
	PermissionResourceID_角色 = "01964c1d-fd4c-7717-a293-eede163b5b43"
	// 仅用于测试：角色组
	PermissionResourceID_角色组 = "01964c1d-fd4c-771b-a39d-a4b75a6b9417"
)
