package authorization

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type Role interface {
	GetRole(ctx context.Context, id string) (*RoleDetail, error)
	ListAccessorRoles(ctx context.Context, req *ListAccessorRolesArgs) ([]*RoleMetaInfo, error)
	ListUserRoles(ctx context.Context, uid string) ([]*RoleMetaInfo, error)
	ListUserRoleID(ctx context.Context, uid string) ([]string, error)
	HasInnerBusinessRoles(ctx context.Context, uid string) ([]string, error)
	ListRoles(ctx context.Context, req *RoleListArgs) (*base.PageResult[RoleDetail], error)
	ListRoleMembers(ctx context.Context, req *RoleMemberArgs) (*base.PageResult[MemberInfo], error)
	ListRoleTotalMembers(ctx context.Context, roleID string) ([]*MemberInfo, error)
	UpdateRoleMembers(ctx context.Context, req *UpdateRoleMemberArgs) error
	HasRoles(ctx context.Context, uid string, roleID ...string) (bool, error)
}

//region  GetRole

type RoleMetaInfo struct {
	ID          string `json:"id"`   //角色唯一标识
	Name        string `json:"name"` //角色名称，唯一
	Description string `json:"description"`
}

type RoleDetail struct {
	RoleMetaInfo
	ResourceTypeScope *ResourceTypeScope `json:"resource_type_scope"` //资源类型范围
}

//endregion

//region ListRoles

type RoleListArgs struct {
	Offset  int    `query:"offset"`  //default:0,>=0,获取数据起始下标
	Limit   int    `query:"limit"`   //default:20,[1,1000],获取数据量
	Keyword string `query:"keyword"` //搜索内容, 即角色名
	//Source 角色来源:
	//- system 系统
	//- business 业务内置
	//- user 用户自定义
	//返回结果默认按照system、business、user 顺序排序
	//参数不传时默认返回business、user
	Source string `query:"source"`
}

//endregion

//region ListRoleMembers

type RoleMemberArgs struct {
	ID      string `uri:"id"`        //角色唯一标识
	Keyword string `query:"keyword"` //搜索内容, 即成员名称。
	Offset  int    `query:"offset"`  //default:0,>=0,获取数据起始下标
	Limit   int    `query:"limit"`   //default:20,[1,1000],获取数据量
}

type MemberInfo struct {
	ID         string                `json:"id"`          //成员id，可以是用户id，部门id
	Name       string                `json:"name"`        //成员显示名
	Type       string                `json:"type"`        //枚举：user用户，department部门，group用户组，app应用账户
	ParentDeps [][]*DepartmentObject `json:"parent_deps"` //父部门信息，描述多个父部门的层级关系信息，每个父部门层级数组内第一个对象是根部门，最后一个对象是直接父部门
}

type DepartmentObject struct {
	ID   string `json:"id"`   //部门ID
	Name string `json:"name"` //部门名称
	Type string `json:"type"` //固定为"department"
}

//endregion

//region UpdateRoleMembers

type UpdateRoleMemberArgs struct {
	ID      string              `uri:"id"`       //角色唯一标识
	Method  string              `json:"method"`  //方法名，目前支持添加和删除操作,POST添加，DELETE删除
	Members []*MemberSimpleInfo `json:"members"` //成员信息
}

type MemberSimpleInfo struct {
	ID   string `json:"id"`   //成员id
	Type string `json:"type"` //user用户,department部门,group用户组,app应用账户
}

//endregion

//region ListAccessorRoles

type ListAccessorRolesArgs struct {
	AccessorID   string `query:"accessor_id"`   //访问者唯一标识
	AccessorType string `query:"accessor_type"` //访问者类型
	Offset       int    `query:"offset"`        //default:0,>=0,获取数据起始下标
	Limit        int    `query:"limit"`         //default:20,[1,1000],获取数据量
	Keyword      string `query:"keyword"`       //搜索内容, 即角色名
	//Source 角色来源:
	//- system 系统
	//- business 业务内置
	//- user 用户自定义
	//返回结果默认按照system、business、user 顺序排序
	//参数不传时默认返回business、user
	//Source string `query:"source"`
}

//endregion
