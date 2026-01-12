package v1

import (
	"net/url"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// WorkOrderTask 代表工单任务
type WorkOrderTask struct {
	// ID，格式 UUID v7
	ID string `json:"id,omitempty"`
	// 第三方平台的 ID
	ThirdPartyID string `json:"third_party_id,omitempty"`
	// 创建时间
	CreatedAt meta_v1.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt meta_v1.Time `json:"updated_at,omitempty"`
	// 名称，最大长度 128
	Name string `json:"name,omitempty"`
	// 所属工单 ID，格式 UUID v7
	WorkOrderID string `json:"work_order_id,omitempty"`
	// 工单任务状态
	Status WorkOrderTaskStatus `json:"status,omitempty"`
	// 任务处于当前状态的原因，比如失败原因
	Reason string `json:"reason,omitempty"`
	// 任务失败处理 URL
	Link string `json:"link,omitempty"`

	// 根据工单类型对应不同任务详情
	WorkOrderTaskTypedDetail
}

// WorkOrderTaskInternal 与 WorkOrderTask 类似，代表工单任务。区别在于
//
//  1. 部门 ID 是 AnyFabric 的 ID，而不是第三方的 ID
//  2. 数据源 ID 是 AnyFabric 的 ID，而不是第三方的 ID
type WorkOrderTaskInternal struct {
	// ID，格式 UUID v7
	ID string `json:"id,omitempty"`
	// 第三方平台的 ID
	ThirdPartyID string `json:"third_party_id,omitempty"`
	// 创建时间
	CreatedAt meta_v1.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt meta_v1.Time `json:"updated_at,omitempty"`
	// 名称，最大长度 128
	Name string `json:"name,omitempty"`
	// 所属工单 ID，格式 UUID v7
	WorkOrderID string `json:"work_order_id,omitempty"`
	// 工单任务状态
	Status WorkOrderTaskStatus `json:"status,omitempty"`
	// 任务处于当前状态的原因，比如失败原因
	Reason string `json:"reason,omitempty"`
	// 任务失败处理 URL
	Link string `json:"link,omitempty"`

	// 根据工单类型对应不同任务详情
	WorkOrderTaskTypedDetailInternal
}

// WorkOrderTaskTypedDetail 代表各种类型工单的任务详情
//
// 有且只有一个字段有值
type WorkOrderTaskTypedDetail struct {
	// 数据归集工单的任务详情
	DataAggregation []WorkOrderTaskDetailAggregationDetail `json:"data_aggregation,omitempty"`
	// 数据理解工单的任务详情
	//
	// Deprecated: 理解工单的任务不由 WorkOrderTask 定义
	DataComprehension *WorkOrderTaskDetailComprehensionDetail `json:"data_comprehension,omitempty"`
	// 数据融合工单的任务详情
	DataFusion *WorkOrderTaskDetailFusionDetail `json:"data_fusion,omitempty"`
	// 数据质量工单的任务详情
	DataQuality *WorkOrderTaskDetailQualityDetail `json:"data_quality,omitempty"`
	// 数据质量稽查工单的任务详情
	DataQualityAudit []*WorkOrderTaskDetailQualityAuditDetail `json:"data_quality_audit,omitempty"`
}

// WorkOrderTaskTypedDetail 与 WorkOrderTaskTypedDetail 类似，代表各种类型工单的
// 任务详情。区别在于
//
//  1. 部门 ID 是 AnyFabric 的 ID，而不是第三方的 ID
//  2. 数据源 ID 是 AnyFabric 的 ID，而不是第三方的 ID
//
// 有且只有一个字段有值
type WorkOrderTaskTypedDetailInternal struct {
	// 数据归集工单的任务详情
	DataAggregation []WorkOrderTaskDetailAggregationDetailInternal `json:"data_aggregation,omitempty"`
	// 数据融合工单的任务详情
	DataFusion *WorkOrderTaskDetailFusionDetail `json:"data_fusion,omitempty"`
	// 数据质量工单的任务详情
	DataQuality *WorkOrderTaskDetailQualityDetail `json:"data_quality,omitempty"`
	// 数据质量稽查工单的任务详情
	DataQualityAudit []*WorkOrderTaskDetailQualityAuditDetail `json:"data_quality_audit,omitempty"`
}

// 数据归集工单的任务详情
type WorkOrderTaskDetailAggregationDetail struct {
	// 部门 ID
	DepartmentID string `json:"department_id,omitempty"`
	// 源表
	Source WorkOrderTaskDetailAggregationTableReference `json:"source,omitempty"`
	// 目标表
	Target WorkOrderTaskDetailAggregationTableReference `json:"target,omitempty"`
	// 归集数量，代表这个任务中归集的数据的数量
	Count int `json:"count,omitempty"`
}

// WorkOrderTaskDetailAggregationDetailInternal 与
// WorkOrderTaskDetailAggregationDetail 类似，代表数据归集工单的任务详情。区别在于
//
//  1. 部门 ID 是 AnyFabric 的 ID，而不是第三方的 ID
//  2. 数据源 ID 是 AnyFabric 的 ID，而不是第三方的 ID
type WorkOrderTaskDetailAggregationDetailInternal struct {
	// 部门 ID
	DepartmentID string `json:"department_id,omitempty"`
	// 源表
	Source WorkOrderTaskDetailAggregationTableReferenceInternal `json:"source,omitempty"`
	// 目标表
	Target WorkOrderTaskDetailAggregationTableReferenceInternal `json:"target,omitempty"`
	// 归集数量，代表这个任务中归集的数据的数量
	Count int `json:"count,omitempty"`
}

// 数据归集工单任务中对数据表的引用
type WorkOrderTaskDetailAggregationTableReference struct {
	// 数据源 ID。第三方数据源的 ID，而不是 AnyFabric 数据源的 ID。
	DatasourceID string `json:"datasource_id,omitempty"`
	// 表名称
	TableName string `json:"table_name,omitempty"`
}

// WorkOrderTaskDetailAggregationTableReferenceInternal 与
// WorkOrderTaskDetailAggregationTableReference 类似，代表数据归集工单任务中对数
// 据表的引用。区别在于
//
//  1. 部门 ID 是 AnyFabric 的 ID，而不是第三方的 ID
//  2. 数据源 ID 是 AnyFabric 的 ID，而不是第三方的 ID
type WorkOrderTaskDetailAggregationTableReferenceInternal struct {
	// 数据源 ID。AnyFabric 数据源的 ID 而不是 第三方数据源的 ID。
	DatasourceID string `json:"datasource_id,omitempty"`
	// 表名称
	TableName string `json:"table_name,omitempty"`
}

// 数据理解工单的任务详情
type WorkOrderTaskDetailComprehensionDetail struct{}

// 数据融合工单的任务详情
type WorkOrderTaskDetailFusionDetail struct {
	// 数据源 ID
	DatasourceID string `json:"datasource_id,omitempty"`
	// 数据源名称
	DatasourceName string `json:"datasource_name,omitempty"`
	// 数据表名称
	DataTable string `json:"data_table,omitempty"`
}

// 数据质量工单的任务详情
type WorkOrderTaskDetailQualityDetail struct{}

// 数据质量稽查工单的任务详情
type WorkOrderTaskDetailQualityAuditDetail struct {
	// ID，格式 UUID v7
	ID string `json:"id,omitempty"`
	// 所属工单 ID，格式 UUID v7
	WorkOrderID string `json:"work_order_id,omitempty"`
	// 数据源 ID
	DatasourceID string `json:"datasource_id,omitempty"`
	// 数据源名称
	DatasourceName string `json:"datasource_name,omitempty"`
	// 数据表名称
	DataTable string `json:"data_table,omitempty"`
	// 检测方案
	DetectionScheme string `json:"detection_scheme,omitempty"`
	// 工单任务状态
	Status WorkOrderTaskStatus `json:"status,omitempty"`
	// 任务处于当前状态的原因，比如失败原因
	Reason string `json:"reason,omitempty"`
	// 任务失败处理 URL
	Link string `json:"link,omitempty"`
}

// WorkOrderTaskStatus 代表工单任务状态
type WorkOrderTaskStatus string

// WorkOrderTaskStatus 代表工单任务状态
const (
	// 进行中
	WorkOrderTaskRunning WorkOrderTaskStatus = "Running"
	// 已完成
	WorkOrderTaskCompleted WorkOrderTaskStatus = "Completed"
	// 异常
	WorkOrderTaskFailed WorkOrderTaskStatus = "Failed"
)

// WorkOrderTaskList 代表工单任务列表
type WorkOrderTaskList meta_v1.List[WorkOrderTask]

// WorkOrderTaskListOptions 代表获取工单任务的选项
type WorkOrderTaskListOptions struct {
	meta_v1.ListOptions

	// 关键字，非空时对工单任务名称做模糊匹配
	Keyword string `json:"keyword,omitempty"`
	// 状态，非空时过滤过滤指定状态的工单任务
	Status WorkOrderTaskStatus `json:"status,omitempty"`
	// 工单类型，非空时过滤属于指定类型工单的工单任务
	WorkOrderType WorkOrderType `json:"work_order_type,omitempty"`
	// 工单 ID。非空时，返回属于这个工单的任务。
	WorkOrderID string `json:"work_order_id,omitempty" form:"work_order_id"`
}

func (opts *WorkOrderTaskListOptions) UnmarshalQuery(data url.Values) (err error) {
	if err = opts.ListOptions.UnmarshalQuery(data); err != nil {
		return
	}
	for k, values := range data {
		for _, v := range values {
			switch k {
			case "work_order_id":
				opts.WorkOrderID = v
			default:
				continue
			}
		}
	}
	return nil
}
