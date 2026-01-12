package v1

import "path"

// Subject 定义访问者
type Subject struct {
	// 访问者类型
	Type SubjectType `json:"subject_type,omitempty"`
	// 访问者 ID
	ID string `json:"subject_id,omitempty"`
}

func (s *Subject) Key() string {
	return path.Join(string(s.Type), s.ID)
}

// SubjectType 定义访问者的类型
type SubjectType string

const (
	// 用户
	SubjectUser SubjectType = "user"
	// 部门
	SubjectDepartment SubjectType = "department"
	// 角色
	SubjectRole SubjectType = "role"
	// 应用
	SubjectAPP SubjectType = "app"
)
