package frontend

import (
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
	task_center_v1 "github.com/kweaver-ai/idrm-go-common/api/task_center/v1"
)

type WorkOrderTaskList meta_v1.List[WorkOrderTaskListItem]

// WorkOrderTaskListItem 代表工单任务列表的其中一项
type WorkOrderTaskListItem struct {
	// ID，格式 UUID v7
	ID string `json:"id,omitempty"`
	// 创建时间
	CreatedAt meta_v1.Time `json:"created_at,omitempty"`
	// 名称，最大长度 128
	Name string `json:"name,omitempty"`
	// 所属工单
	WorkOrder meta_v1.ReferenceWithName `json:"work_order,omitempty"`
	// 工单任务状态
	Status task_center_v1.WorkOrderTaskStatus `json:"status,omitempty"`
	// 任务处于当前状态的原因，比如失败原因
	Reason string `json:"reason,omitempty"`
	// 任务失败处理 URL
	Link string `json:"link,omitempty"`
	// 根据工单类型对应不同任务详情
	WorkOrderTaskTypedDetail
}

// WorkOrderTaskTypedDetail 代表任务详情
type WorkOrderTaskTypedDetail struct {
	// 数据归集工单的任务详情
	DataAggregation []WorkOrderTaskDetailAggregation `json:"data_aggregation,omitempty"`
	// 数据理解工单的任务详情
	DataComprehension *task_center_v1.WorkOrderTaskDetailComprehensionDetail `json:"data_comprehension,omitempty"`
	// 数据融合工单的任务详情
	DataFusion *task_center_v1.WorkOrderTaskDetailFusionDetail `json:"data_fusion,omitempty"`
	// 数据质量工单的任务详情
	DataQuality *task_center_v1.WorkOrderTaskDetailQualityDetail `json:"data_quality,omitempty"`
	// 数据质量稽查工单的任务详情
	DataQualityAudit []*task_center_v1.WorkOrderTaskDetailQualityAuditDetail `json:"data_quality_audit,omitempty"`
}

// WorkOrderTaskDetailAggregation 数据归集工单的任务详情
type WorkOrderTaskDetailAggregation struct {
	// 部门
	Department DepartmentReference `json:"department,omitempty"`
	// 目标表名称
	TableName string `json:"table_name,omitempty"`
	// 归集数量，代表这个任务中归集的数据的数量
	Count int `json:"count,omitempty"`
}

// DepartmentReference 代表对部门的引用
type DepartmentReference struct {
	// 部门 ID
	ID string `json:"id,omitempty"`
	// 部门路径
	Path string `json:"path,omitempty"`
}
