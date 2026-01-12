package v1

import (
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// User 代表用户
type User struct {
	// 元数据。已经的删除用户通过 Status 标识，所以 DeletedAt 永远为 nil
	meta_v1.Metadata
	// 定义
	UserSpec
}

// UserSpec 定义一个用户
type UserSpec struct {
	// 显示名称
	//
	// Deprecated: 为避免歧义，请使用 DisplayName
	Name string `json:"name,omitempty"`
	// 显示名称
	DisplayName string `json:"display_name,omitempty"`
	// 登录名称
	LoginName string `json:"login_name,omitempty"`
	// 权限范围
	Scope Scope `json:"scope,omitempty"`
	// 用户类型
	UserType UserType `json:"user_type,omitempty"`
	// 手机号码
	PhoneNumber string `json:"phone_number,omitempty"`
	// 邮箱地址
	MailAddress string `json:"mail_address,omitempty"`
	// 用户状态
	Status UserStatus `json:"status,omitempty"`
	// 注册时间
	RegisteredAt string `json:"register_at,omitempty"`
	// 是否注册
	Registered int32 `json:"registered,omitempty"`
	// 第三方服务ID
	ThirdServiceID string `json:"gateway_id,omitempty"`
	// 第三方用户ID
	ThirdUserId string `json:"third_user_id,omitempty"` //第三方用户ID
	Sex         string `json:"sex,omitempty"`           //性别
}

// 用户类型
type UserType int

const (
	// 用户类型：普通账号
	UserTypeNormal UserType = iota + 1
	// 用户类型：应用账号
	UserTypeAPP
)

// 用户状态
type UserStatus int

const (
	// 用户状态：正常
	UserStatusNormal UserStatus = iota + 1
	// 用户状态：删除
	UserStatusDeleted
)

// UserList 代表用户列表
type UserList meta_v1.List[User]

// UserListOptions 代表获取用户列表接口的 query 参数
type UserListOptions struct {
	meta_v1.ListOptions
	// 关键字
	Keyword string `json:"keyword,omitempty" form:"keyword"`
	// 所属部门 ID
	DepartmentID string `json:"department_id,omitempty" form:"department_id"`
	// 是否注册
	Registered int32 `json:"registered,omitempty" form:"registered"`
}

type UserRoleOrRoleGroupBinding struct {
	// 用户 ID
	UserID string `json:"user_id,omitempty"`
	// 角色 ID，与角色组 ID 冲突
	RoleID string `json:"role_id,omitempty"`
	// 角色组 ID，与角色 ID 冲突
	RoleGroupID string `json:"role_group_id,omitempty"`
}

// UserRoleOrRoleGroupBindingProcessing 代表对用户角色、用户角色组绑定关系的处理
type UserRoleOrRoleGroupBindingProcessing struct {
	UserRoleOrRoleGroupBinding
	// 期望状态
	State meta_v1.ProcessingState `json:"state,omitempty"`
}

// UserRoleOrRoleGroupBindingBatchProcessing 代表对用户角色、用户角色组绑定关系的批处理
type UserRoleOrRoleGroupBindingBatchProcessing struct {
	// 用户 ID 列表，非空时作为默认值
	UserIDs []string `json:"user_ids,omitempty"`
	// 角色 ID 列表，非空时作为默认值
	RoleIDs []string `json:"role_ids,omitempty"`
	// 角色组 ID 列表，非空时作为默认值
	RoleGroupIDs []string `json:"role_group_ids,omitempty"`
	// 期望绑定关系是否存在，非空时作为默认值
	State meta_v1.ProcessingState `json:"state,omitempty"`
	// 用户角色、用户角色组绑定关系列表
	Bindings []UserRoleOrRoleGroupBindingProcessing `json:"bindings,omitempty"`
}
