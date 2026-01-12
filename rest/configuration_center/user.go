package configuration_center

type UserInfo struct {
	ID          string `json:"id"`           // 主键，uuid
	Name        string `json:"name"`         // 显示名称
	Status      int32  `json:"status"`       // 用户状态,1正常,2删除
	UserType    int32  `json:"user_type"`    // 用户类型,1普通账号,2应用账号
	PhoneNumber string `json:"phone_number"` // 手机号码
	MailAddress string `json:"mail_address"` // 邮箱地址
	LoginName   string `json:"login_name"`   // 登录名称
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
}

type UserRespItem struct {
	UserInfo
	ParentDeps [][]DepartV1 `json:"parent_deps"`
	Roles      []*Role      `json:"roles"`
}

// GetFirstDepartCode 获取用户第一个第三方部门ID，没有就返回AF的部门ID
func (u *UserRespItem) GetFirstDepartCode() string {
	if len(u.ParentDeps) <= 0 || len(u.ParentDeps[0]) <= 0 {
		return ""
	}
	lastThirdDept := u.ParentDeps[0][len(u.ParentDeps[0])-1]
	if lastThirdDept.ThirdDeptId != "" {
		return lastThirdDept.ThirdDeptId
	}
	return lastThirdDept.ID
}

type DepartV1 struct {
	ID          string `json:"department_id"`   // 部门标识
	Name        string `json:"department_name"` // 部门名称
	ThirdDeptId string `json:"third_dept_id"`   // 第三方部门ID，部分接口可能没有，具体见接口注释！注意！
}

type GetUserListReq struct {
	DepartId                 string `query:"depart_id"`                   // 部门id
	IsDepartInNeed           string `query:"is_depart_in_need"`           // 是否返回用户部门信息
	IsIncludeUnassignedRoles string `query:"is_include_unassigned_roles"` // 是否返回未分配角色的用户
	Offset                   int    `query:"offset"`                      // 页码
	Limit                    int    `query:"limit"`                       // 每页大小，默认为0不分页
	Keyword                  string `query:"keyword"`                     // 关键字查询
	Direction                string `query:"direction"`                   // 排序方向，可选值：asc desc，默认desc
	Sort                     string `query:"sort"`                        // 排序类型，可选值：name updated_at，默认name
}

type GetDatasourcesOptions struct {
	// 第三方数据源 ID
	HuaAoID string `json:"hua_ao_id,omitempty" query:"hua_ao_id"`
}
