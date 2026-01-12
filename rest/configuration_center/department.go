package configuration_center

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type DepartmentService interface {
	GetDepartments(ctx context.Context, orgCodes []string) ([]*DepartmentObject, error)  //获取部门列表，省市直达用，其他服务慎用
	GetDepartmentsByCode(ctx context.Context, orgCode string) (*DepartmentObject, error) //获取部门列表，省市直达用，其他服务慎用
	//下面是常规服务可用
	GetDepartmentsByCodeInternal(ctx context.Context, orgCode string) (*DepartmentObject, error)
	GetDepartmentList(ctx context.Context, reqParam QueryPageReqParam) (res *QueryPageReapParam, err error)
	GetDepartmentListInternal(ctx context.Context, reqParam QueryPageReqParam) (res *QueryPageReapParam, err error)
	GetDepartmentPrecision(ctx context.Context, ids []string) (res *GetDepartmentPrecisionRes, err error)
	GetDepartmentByPath(ctx context.Context, paths *GetDepartmentByPathReq) (res *GetDepartmentByPathRes, err error)
	DeleteFile(ctx context.Context, deptID string, ossID string, fileName string) (err error) // 删除部门关联文件
	GetDepartmentsByUserID(ctx context.Context, userID string) ([]*DepartmentObject, error)
	GetChildDepartments(ctx context.Context, orgCode string) (*base.PageResult[DepartmentObject], error)
	GetDepartAndSubDepartIds(ctx context.Context, userId string) ([]string, error)
	GetDepartmentsByIds(ctx context.Context, ids []string) ([]*DepartmentObject, error)
	GetMainDepartIdsByUserID(ctx context.Context, userID string) ([]string, error)
	DepartmentServiceExtend
}
type DepartmentServiceExtend interface {
	GetDepartmentInMap(ctx context.Context, ids []string) (res map[string]*DepartmentObject, err error)
}

//region GetDepartmentList

type QueryPageReqParam struct {
	Offset       int    `json:"offset" form:"offset,default=1" binding:"min=1" default:"1"`                           // 页码
	Limit        int    `json:"limit" form:"limit,default=10" binding:"min=0,max=100" default:"10"`                   // 每页大小，为0时不分页
	ID           string `json:"id" form:"id" binding:"omitempty,uuid" example:"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"` // 对象id
	IDsSubDepart string `json:"ids_sub_depart" form:"ids_sub_depart" `                                                // 多个对象id的子部门
	ThirdDeptID  string `json:"third_dept_id,omitempty" form:"third_dept_id"`                                         // 第三方部门 ID
}

type QueryPageReapParam struct {
	Entries    []*SummaryInfo `json:"entries" binding:"required"`                      // 对象列表
	TotalCount int64          `json:"total_count" binding:"required,ge=0" example:"3"` // 当前筛选条件下的对象数量
}
type SummaryInfo struct {
	ID   string `json:"id" `  // 对象ID
	Name string `json:"name"` // 对象名称
	Type string `json:"type"` // 对象类型
	Path string `json:"path"` // 对象路径

	PathID string `json:"path_id"` // 对象ID路径
	Expand bool   `json:"expand"`  // 是否能展开

	ThirdDeptID string `json:"third_dept_id,omitempty"` // 第三方部门 ID
}

//endregion

// region GetDepartments

type QueryDepartmentPageResp struct {
	Entries    []*DepartmentObject `json:"entries" binding:"required"`                      // 对象列表
	TotalCount int64               `json:"total_count" binding:"required,ge=0" example:"3"` // 当前筛选条件下的对象数量
}

type DepartmentObject struct {
	ID               string `json:"id"`      // 对象ID
	Name             string `json:"name"`    // 对象名称
	PathID           string `json:"path_id"` // 路径ID
	Path             string `json:"path"`    // 路径
	Type             string `json:"type"`    // 类型
	Expand           bool   `json:"expand"`  // 是否能展开
	ObjectAttributes `json:"attributes"`
	DeletedAt        int32  `json:"deleted_at"`    // 删除时间(逻辑删除)
	Subtype          int32  `json:"subtype"`       // 对象子类型，用于对象类型二次分类，有效值包括0-未分类 1-行政区 2-部门 3-处（科）室
	ThirdDeptId      string `json:"third_dept_id"` //第三方部门ID
}

type DepartmentsResp struct {
	Entries []*DepartmentObject `json:"entries" binding:"required"` // 对象列表
}

//endregion

//region GetDepartmentList

type GetDepartmentPrecisionRes struct {
	Departments []*DepartmentInternal `json:"departments"`
}
type DepartmentInternal struct {
	ID               string `gorm:"column:id;primaryKey" json:"id"`         // 对象ID
	Name             string `gorm:"column:name;not null" json:"name"`       // 对象名称
	PathID           string `gorm:"column:path_id;not null" json:"path_id"` // 路径ID
	Path             string `gorm:"column:path;not null" json:"path"`       // 路径
	Type             int32  `gorm:"column:type" json:"type"`                // 类型
	Expand           bool   `json:"expand"`                                 // 是否能展开
	ObjectAttributes `json:"attributes"`
	DeletedAt        int32  `gorm:"column:deleted_at;not null;softDelete:milli" json:"deleted_at"` // 删除时间(逻辑删除)
	ThirdDeptId      string `json:"third_dept_id"`                                                 //第三方部门ID
}

type ObjectAttributes struct {
	CreditCode string `json:"uniform_credit_code"` // 统一社会信用代码
}

//endregion

//region GetDepartmentByPath

type GetDepartmentByPathReq struct {
	Paths []string `json:"paths" binding:"required,gt=0,unique"`
}

type GetDepartmentByPathRes struct {
	Departments map[string]*DepartmentInternal `json:"departments"`
}

//endregion
