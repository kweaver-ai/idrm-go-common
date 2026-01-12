package task_center

import (
	"context"
	"time"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type SandboxDriven interface {
	GetSandboxSampleInfo(ctx context.Context, sandboxID string) (*SandboxSpaceDetail, error)
	GetUserSandboxSampleInfo(ctx context.Context) (*base.PageResult[SandboxSpaceListItem], error)
}

type SandboxSpaceDetail struct {
	ID             string     `json:"id"`              // 主键雪UUID
	DepartmentID   string     `json:"department_id"`   // 所属部门ID
	DepartmentName string     `json:"department_name"` // 所属部门名称
	ProjectID      string     `json:"project_id"`      // 项目ID
	Status         int32      `json:"status"`          // 状态，0不可用，1可用
	TotalSpace     int32      `json:"total_space"`     // 总的沙箱空间，单位GB
	ValidStart     int64      `json:"valid_start"`     // 有效期开始时间，单位毫秒
	ValidEnd       int64      `json:"valid_end"`       // 有效期结束时间，单位毫秒
	ApplicantID    string     `json:"applicant_id"`    // 申请人ID
	ApplicantName  string     `json:"applicant_name"`  // 申请人名称
	ApplicantPhone string     `json:"applicant_phone"` // 申请人手机号
	ExecutedTime   *time.Time `json:"executed_time"`   // 沙箱第一次申请实施时间
	//沙箱空间信息
	DatasourceID       string `json:"datasource_id"`        // 数据源UUID
	DatasourceName     string `json:"datasource_name"`      // 数据源名称,catalog
	DatasourceTypeName string `json:"datasource_type_name"` // 数据库类型名称
	DatabaseName       string `json:"database_name"`        // 数据库名称
}

type SandboxSpaceListItem struct {
	SandboxID      string  `gorm:"column:sandbox_id;not null" json:"sandbox_id"`           // 沙箱ID
	ApplicantID    string  `gorm:"column:applicant_id" json:"applicant_id"`                // 申请人ID
	ApplicantName  string  `gorm:"column:applicant_name" json:"applicant_name"`            // 申请人名称
	DepartmentID   string  `gorm:"column:department_id;not null" json:"department_id"`     // 所属部门ID
	DepartmentName string  `gorm:"column:department_name;not null" json:"department_name"` // 所属部门名称
	ProjectID      string  `gorm:"column:project_id;not null" json:"project_id"`           // 项目ID
	ProjectName    string  `gorm:"column:project_name;not null" json:"project_name"`       // 项目名称
	TotalSpace     int32   `gorm:"column:total_space" json:"total_space"`                  // 总的沙箱空间，单位GB
	UsedSpace      float64 `gorm:"-" json:"used_space"`                                    // 已用空间
	ValidStart     int64   `gorm:"column:valid_start" json:"valid_start"`                  // 有效期开始时间，单位毫秒
	ValidEnd       int64   `gorm:"column:valid_end" json:"valid_end"`                      // 有效期结束时间，单位毫秒
	DataSetNumber  int32   `gorm:"-" json:"data_set_number"`                               // 数据集数量
	RecentDataSet  string  `gorm:"-"  json:"recent_data_set"`                              // 最近的一个数据及名称
	UpdatedAt      string  `gorm:"column:updated_at" json:"updated_at"`                    // 数据集更新时间
}
