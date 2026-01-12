package v1

// Operator 定义了操作者，操作者可以是用户或应用账号
type Operator struct {
	// 操作者类型
	Type OperatorType `json:"type,omitempty"`
	// 操作者 ID
	ID string `json:"id,omitempty"`
	// 操作者名称，如果是用户则显示为用户名
	Name string `json:"name,omitempty"`
	// 登录名称
	LoginName string `json:"login_name"`
	// 用户所属部门，仅当 Type 为 OperatorAuthenticatedUser 时需要
	Department []Department `json:"department,omitempty"`
	//直属部门编号, 如果直属部门的第三方ID不为空，则是直属部门的第三方ID，否则是AF自身的uuid
	DepartmentCode string `json:"department_code,omitempty"`
	// 用户代理，仅当 Type 为 OperatorAuthenticatedUser 时需要
	Agent Agent `json:"agent,omitempty"`
}

// OperatorType 定义操作者的类型
type OperatorType string

// OperatorType 操作者类型
const (
	// 未知
	OperatorUnknown OperatorType = "unknown"
	// 用户
	OperatorAuthenticatedUser OperatorType = "authenticated_user"
	// 应用
	OperatorAPP OperatorType = "app"
)

// Department 定义用户所属的部门
type Department struct {
	// 用户所属部门 ID 全路径
	ID string `json:"id,omitempty"`
	// 用户所属部门名称全路径
	Name string `json:"name,omitempty"`
}
