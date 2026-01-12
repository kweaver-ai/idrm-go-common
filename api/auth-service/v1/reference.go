package v1

// ReferenceSource 定义引用的资源。有且只有一个属性会被设置。
type ReferenceSource struct {
	// application 代表引用的应用
	Application *ApplicationSource `json:"application,omitempty"`
	// department 代表引用的部门
	Department *DepartmentSource `json:"department,omitempty"`
	// user 代表引用的用户
	User *UserSource `json:"user,omitempty"`
	// IndicatorDimensionalRule 代表指标维度规则
	IndicatorDimensionalRule *IndicatorDimensionalRule `json:"indicator_dimensional_rule,omitempty"`
}

// ApplicationSource 定义引用的应用
type ApplicationSource struct {
	// ID
	ID string `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
}

// UserSource 定义引用的用户
type UserSource struct {
	// 用户 ID
	ID string `json:"id,omitempty"`
	// 显示名称
	Name string `json:"name,omitempty"`
	// 所属部门 ID 列表，如果为空，代表不属于任何部门
	DepartmentIDs []string `json:"department_ids,omitempty"`
}

// DepartmentSource 定义引用的部门
type DepartmentSource struct {
	// 部门 ID
	ID string `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 上级部门 ID，如果为空，代表不存在上级部门
	ParentID string `json:"parent_id,omitempty"`
}
