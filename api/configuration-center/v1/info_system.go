package v1

import (
	"github.com/google/uuid"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// 信息系统
type InfoSystem struct {
	// 元数据
	meta_v1.Metadata
	// 创建者 ID
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// 更新者 ID
	UpdatedBy uuid.UUID `json:"updated_by,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
	// 所属部门 ID，空值代表信息系统未属于任何部门
	DepartmentID uuid.UUID `json:"department_id,omitempty"`
}
