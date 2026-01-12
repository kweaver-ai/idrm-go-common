package task_center

import (
	"context"

	task_center_v1 "github.com/kweaver-ai/idrm-go-common/api/task_center/v1"
)

type Driven interface {
	GetTaskDetailById(ctx context.Context, id string) (*GetTaskDetailByIdRes, error)
	GetStandardFormAggregationInfo(ctx context.Context, formID []string) ([]*task_center_v1.BusinessFormDataTableItem, error)
	GetWorkOrderList(ctx context.Context, req *GetWorkOrderListReq) (*GetWorkOrderListResp, error)
	GetProjectModels(ctx context.Context, id string) (*ProjectModelInfo, error)
	GetCatalogTaskStatus(ctx context.Context, formId, formName, catalogId string) (*CatalogTaskStatusResp, error)
	GetCatalogTask(ctx context.Context, formId, formName, catalogId string) (*CatalogTaskResp, error)
	GetDataAggregationTask(ctx context.Context, formNames string) (*DataAggregationTaskResp, error)
	SandboxDriven
	WorkOrderGetter
	WorkOrderTaskGetter
	WorkOrderTaskInternalGetter
}

// region GetTaskDetailById

type GetTaskDetailByIdRes struct {
	Id                          string   `json:"id"`                             // 任务id
	Name                        string   `json:"name"`                           // 任务名称
	ProjectId                   string   `json:"project_id"`                     // 项目id
	ProjectName                 string   `json:"project_name"`                   // 项目名称
	Image                       string   `json:"image"`                          // 项目图片
	StageId                     string   `json:"stage_id"`                       // 阶段id
	StageName                   string   `json:"stage_name"`                     // 阶段名称
	NodeId                      string   `json:"node_id"`                        // 节点id
	NodeName                    string   `json:"node_name"`                      // 节点名称
	Status                      string   `json:"status"`                         // 任务状态
	ConfigStatus                string   `json:"config_status"`                  // 任务配置状态，标记缺失的依赖，业务域或者主干业务
	ExecutableStatus            string   `json:"executable_status"`              // 任务的可执行状态
	Deadline                    int64    `json:"deadline"`                       // 截止日期
	Overdue                     string   `json:"overdue"`                        // 是否逾期
	Priority                    string   `json:"priority"`                       // 任务优先级
	ExecutorId                  string   `json:"executor_id"`                    // 任务执行人id
	ExecutorName                string   `json:"executor_name"`                  // 任务执行人
	Description                 string   `json:"description"`                    // 任务描述
	CreatedBy                   string   `json:"created_by"`                     // 创建人
	CreatedAt                   int64    `json:"created_at"`                     // 创建时间
	UpdatedBy                   string   `json:"updated_by"`                     // 修改人
	UpdatedAt                   int64    `json:"updated_at"`                     // 修改时间
	OrgType                     *int     `json:"org_type,omitempty"`             // 标准分类
	TaskType                    string   `json:"task_type"`                      // 任务类型
	DomainId                    string   `json:"domain_id,omitempty"`            // 业务流程id
	DomainName                  string   `json:"domain_name,omitempty"`          // 业务流程名字
	BusinessModelID             string   `json:"business_model_id,omitempty"`    // 主干业务id
	BusinessModelName           string   `json:"business_model_name,omitempty"`  // 主干业务id
	ParentTaskId                string   `json:"parent_task_id,omitempty"`       // 父任务的Id
	NewMainBusinessId           string   `json:"new_main_business_id"`           // 新建主干业务任务的主干业务ID,只有新建主干业务有
	DataCatalogID               []string `json:"data_catalog_id"`                // 关联数据资源目录
	DataComprehensionTemplateID string   `json:"data_comprehension_template_id"` // 关联数据理解模板
}

//endregion

// region GetWorkOrderList

type GetWorkOrderListReq struct {
	Type         string   `json:"type" form:"type" binding:"omitempty,oneof=data_comprehension data_aggregation data_standardization data_fusion"` // 工单类型: 数据理解data_comprehension 数据归集data_aggregation 数据标准化data_standardization 数据融合data_fusion
	SourceType   string   `json:"source_type" form:"source_type" binding:"omitempty,oneof=standalone plan business_form data_analysis"`            // 来源类型: 无standalone 计划plan 业务表business_form 数据分析data_analysis
	SourceIds    []string `json:"source_ids" form:"source_ids" binding:"omitempty"`                                                                // 来源id列表
	WorkOrderIds []string `json:"work_order_ids" form:"work_order_ids" binding:"omitempty"`                                                        // 工单id列表
}

type GetWorkOrderListResp struct {
	Entries []*WorkOrderInfo `json:"entries" binding:"required"` // 对象列表
}

type WorkOrderInfo struct {
	WorkOrderId        string      `json:"work_order_id"`        // 工单id
	Name               string      `json:"name"`                 // 名称
	Code               string      `json:"code"`                 // 工单编号
	AuditStatus        string      `json:"audit_status"`         // 审核状态
	AuditDescription   string      `json:"audit_description"`    // 审核描述
	Status             string      `json:"status"`               // 工单状态
	Draft              bool        `json:"draft,omitempty"`      //是否是草稿
	Type               string      `json:"type"`                 // 工单类型
	Priority           string      `json:"priority"`             // 优先级
	ResponsibleUID     string      `json:"responsible_uid"`      // 责任人
	ResponsibleUName   string      `json:"responsible_uname"`    // 责任人名称
	SourceId           string      `json:"source_id"`            // 来源id
	SourceType         string      `json:"source_type"`          // 来源类型
	CreatedAt          int64       `json:"created_at"`           // 创建时间
	FinishedAt         int64       `json:"finished_at"`          // 截止日期
	TaskInfo           []*TaskInfo `json:"tasks"`                // 工单任务信息
	CompletedTaskCount int64       `json:"completed_task_count"` // 已完成任务数量
	FusionTableName    string      `json:"fusion_table_name"`    // 融合表名称，融合工单使用
}

type TaskInfo struct {
	TaskId       string `json:"task_id"`        // 任务id
	TaskName     string `json:"task_name"`      // 任务名称
	DataSourceId string `json:"data_source_id"` // 数据源id，融合工单任务使用
	// 数据表名称，融合工单使用
	DataTable       string                                                `json:"data_table,omitempty"`
	DataAggregation []task_center_v1.WorkOrderTaskDetailAggregationDetail `json:"data_aggregation,omitempty"`
}

//endregion

//region GetProjectModels

type ProjectModelInfo struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Status           string   `json:"status"`
	BusinessDomainID []string `json:"business_domain_id"` //任务关联的业务流程ID
	DataDomainID     []string `json:"data_domain_id"`     //任务关联的业务流程ID
}

//endregion

//region GetCatalogTaskStatus

type CatalogTaskStatusResp struct {
	DataAggregationStatus   string `json:"data_aggregation_status"`   // 归集状态
	DataProcessingStatus    string `json:"data_processing_status"`    // 加工状态
	DataComprehensionStatus string `json:"data_comprehension_status"` // 理解状态
}

//endregion

//region GetCatalogTask

type CatalogTaskResp struct {
	DataAggregation   *DataAggregation   `json:"data_aggregation"`   // 归集
	Processing        *Processing        `json:"processing"`         // 加工
	DataComprehension *DataComprehension `json:"data_comprehension"` // 理解
}

type DataAggregation struct {
	TotalCount                int64         `json:"total_count"`                  // 总任务数
	RunningCount              int64         `json:"running_count"`                // 进行中任务数
	CompletedCount            int64         `json:"completed_count"`              // 已完成任务数
	FailedCount               int64         `json:"failed_count"`                 // 异常任务数
	DataAggregationStatus     string        `json:"data_aggregation_status"`      // 归集任务状态
	DataAggregationSourceInfo []*SourceInfo `json:"data_aggregation_source_info"` // 归集任务来源表信息
}

type Processing struct {
	TotalCount          int64                `json:"total_count"`          // 总任务数
	DataStandardization *DataStandardization `json:"data_standardization"` // 标准检测任务
	DataQualityAudit    *DataQualityAudit    `json:"data_quality_audit"`   // 质量检测任务
	DataFusion          *DataFusion          `json:"data_fusion"`          // 数据融合任务
}

type DataStandardization struct {
	DataStandardizationStatus string `json:"data_standardization_status"` // 标准检测任务状态
	ReportUpdatedAt           int64  `json:"report_updated_at"`           // 标准检测报告更新时间
}

type DataQualityAudit struct {
	DataQualityAuditStatus string `json:"data_quality_audit_status"` // 质量检测任务状态
	ReportUpdatedAt        int64  `json:"report_updated_at"`         // 数据质量报告更新时间
}

type DataFusion struct {
	DataFusionStatus      string                  `json:"data_fusion_status"`       // 数据融合任务状态
	DataFusionSourceForm  []*DataFusionSourceForm `json:"data_fusion_source_form"`  // 融合任务来源表信息
	DataFusionSourceField []string                `json:"data_fusion_source_field"` // 融合任务来源字段信息
}

type DataFusionSourceForm struct {
	*SourceInfo
	DataAggregationStatus     string        `json:"data_aggregation_status"`      // 数据归集任务状态
	DataAggregationSourceInfo []*SourceInfo `json:"data_aggregation_source_info"` // 归集任务来源表信息
}

type SourceInfo struct {
	SourceFormID   string `json:"source_form_id"`   // 源表ID
	SourceFormName string `json:"source_form_name"` // 源表名称
	SourceType     string `json:"source_type"`      // 数据源来源
}

type DataComprehension struct {
	TotalCount                    int64  `json:"total_count"`                      // 数据理解任务数
	DataComprehensionStatus       string `json:"data_comprehension_status"`        // 数据理解任务状态
	DataComprehensionReportStatus int8   `json:"data_comprehension_report_status"` // 数据理解报告状态
	AuditAdvice                   string `json:"audit_advice"`                     //审核意见，仅驳回时有用
	ReportUpdatedAt               int64  `json:"report_updated_at"`                // 数据理解报告更新时间
}

//endregion

//region GetDataAggregationTask

type DataAggregationTaskResp struct {
	Entries []*DataAggregationTaskInfo `json:"entries"`
}
type DataAggregationTaskInfo struct {
	FormName    string `json:"form_name"`     // 表名称
	WorkOrderId string `json:"work_order_id"` // 工单id
	Status      string `json:"status"`        // 状态
	CreatedAt   int64  `json:"created_at"`    // 创建时间
	UpdatedAt   int64  `json:"updated_at"`    // 更新时间
	Count       int    `json:"count"`         // 归集数量
}

//endregion
