package v1

import (
	"net/url"
	"strings"

	doc_audit_rest_v1 "github.com/kweaver-ai/idrm-go-common/api/doc_audit_rest/v1"
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
	"github.com/kweaver-ai/idrm-go-common/util/sets"
)

// 数据归集清单
type DataAggregationInventory struct {
	// ID
	ID string `json:"id,omitempty"`
	// 编号
	Code string `json:"code,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 创建方式
	CreationMethod DataAggregationInventoryCreationMethod `json:"creation_method,omitempty"`
	// 数源单位的 ID
	DepartmentID string `json:"department_id,omitempty"`
	// 归集资源列表
	Resources []DataAggregationResource `json:"resources,omitempty"`
	// Workflow 审核的 ApplyID
	ApplyID string `json:"apply_id,omitempty"`
	// 工单 ID。处理归集工单时，选择已存在的归集清单或新建归集清单
	//
	// Deprecated: Use WorkOrderIDs instead
	WorkOrderID string `json:"work_order_id,omitempty"`
	// 工单 ID 列表
	WorkOrderIDs []string `json:"work_order_i_ds,omitempty"`
	// 状态
	Status DataAggregationInventoryStatus `json:"status,omitempty"`
	// 创建时间
	CreatedAt meta_v1.Time `json:"created_at,omitempty"`
	// 创建人，创建清单的人的 ID
	CreatorID string `json:"creator_id,omitempty"`
	// 申请时间，因为需要根据申请时间排序，但 workflow 不支持，所以在本地记录
	RequestedAt *meta_v1.Time `json:"requested_at,omitempty"`
	// 申请人 ID，提交归集清单申请的人的 ID
	RequesterID string `json:"requester_id,omitempty"`
}

// 数据归集清单创建方式
type DataAggregationInventoryCreationMethod string

// 数据归集清单创建方式
const (
	// 直接创建
	DataAggregationInventoryCreationRaw DataAggregationInventoryCreationMethod = "Raw"
	// 通过工单创建
	DataAggregationInventoryCreationWorkOrder DataAggregationInventoryCreationMethod = "WorkOrder"
)

// 数据归集清单状态
type DataAggregationInventoryStatus string

// 数据归集清单状态
const (
	// 草稿，未发起审核
	DataAggregationInventoryDraft DataAggregationInventoryStatus = "Draft"
	// 审核中
	DataAggregationInventoryAuditing DataAggregationInventoryStatus = "Auditing"
	// 被拒绝
	DataAggregationInventoryReject DataAggregationInventoryStatus = "Reject"
	// 已完成，直接创建的数据归集清单被批准、或通过工单创建的数据归集清单。
	DataAggregationInventoryCompleted DataAggregationInventoryStatus = "Completed"
)

// 数据归集资源，一个数据归集清单包含多个数据归集资源
type DataAggregationResource struct {
	// 逻辑视图 ID
	DataViewID string `json:"data_view_id,omitempty"`
	// 采集方式
	CollectionMethod DataAggregationResourceCollectionMethod `json:"collection_method,omitempty"`
	// 同步频率
	SyncFrequency DataAggregationResourceSyncFrequency `json:"sync_frequency,omitempty"`
	// 关联业务表 ID
	BusinessFormID string `json:"business_form_id,omitempty"`
	// 目标数据源 ID
	TargetDatasourceID string `json:"target_datasource_id,omitempty"`
}

// 数据归集资源采集方式
type DataAggregationResourceCollectionMethod string

// 数据归集资源采集方式
const (
	// 全量
	DataAggregationResourceCollectionFull DataAggregationResourceCollectionMethod = "Full"
	// 增量
	DataAggregationResourceCollectionIncrement DataAggregationResourceCollectionMethod = "Increment"
)

// 数据归集资源同步频率
type DataAggregationResourceSyncFrequency string

// 数据归集资源同步频率
const (
	// 每分钟
	DataAggregationResourceSyncFrequencyPerMinute DataAggregationResourceSyncFrequency = "PerMinute"
	// 每小时
	DataAggregationResourceSyncFrequencyPerHour DataAggregationResourceSyncFrequency = "PerHour"
	// 每天
	DataAggregationResourceSyncFrequencyPerDay DataAggregationResourceSyncFrequency = "PerDay"
	// 每周
	DataAggregationResourceSyncFrequencyPerWeek DataAggregationResourceSyncFrequency = "PerWeek"
	// 每月
	DataAggregationResourceSyncFrequencyPerMonth DataAggregationResourceSyncFrequency = "PerMonth"
	// 每年
	DataAggregationResourceSyncFrequencyPerYear DataAggregationResourceSyncFrequency = "PerYear"
)

// 数据归集清单列表
type DataAggregationInventoryList meta_v1.List[DataAggregationInventory]

// 获取数据归集列表的选项
type DataAggregationInventoryListOptions struct {
	meta_v1.ListOptions
	// 非空时返回关键字匹配的结果
	Keyword string `json:"keyword,omitempty"`
	// 非空时关键字匹配这些字段
	Fields []DataAggregationInventoryListKeywordField `json:"fields,omitempty"`
	// 非空时过滤指定状态的数据归集列表
	Statuses []DataAggregationInventoryStatus `json:"statuses,omitempty"`
	// 非空时过滤指定数据源单位的数据归集列表
	DepartmentIDs []string `json:"department_ids,omitempty"`
}

func (opts *DataAggregationInventoryListOptions) UnmarshalQuery(data url.Values) (err error) {
	if err = opts.ListOptions.UnmarshalQuery(data); err != nil {
		return
	}

	var (
		fields        = sets.New[DataAggregationInventoryListKeywordField]()
		statuses      = sets.New[DataAggregationInventoryStatus]()
		departmentIDs = sets.New[string]()
	)
	for k, values := range data {
		for _, v := range values {
			switch k {
			case "keyword":
				opts.Keyword = v
			case "fields":
				for _, s := range strings.Split(v, ",") {
					fields.Insert(DataAggregationInventoryListKeywordField(s))
				}
			// 虽然是列表，但是与其他筛选条件保持一致使用单数
			case "status":
				for _, s := range strings.Split(v, ",") {
					statuses.Insert(DataAggregationInventoryStatus(s))
				}
			case "department_ids":
				for _, s := range strings.Split(v, ",") {
					departmentIDs.Insert(s)
				}
			default:
				continue
			}
		}
	}
	opts.Fields = sets.List(fields)
	opts.Statuses = sets.List(statuses)
	opts.DepartmentIDs = sets.List(departmentIDs)

	return
}

// DataAggregationInventory 用于 keyword 匹配的字段
type DataAggregationInventoryListKeywordField string

// DataAggregationInventory 用于 keyword 匹配的字段
const (
	// 编码
	DataAggregationInventoryListKeywordFieldCode DataAggregationInventoryListKeywordField = "code"
	// 名称
	DataAggregationInventoryListKeywordFieldName DataAggregationInventoryListKeywordField = "name"
)

// 聚合的数据归集清单，与 DataAggregationInventory 相比包含部门、申请者的名称
type AggregatedDataAggregationInventory struct {
	// ID
	ID string `json:"id,omitempty"`
	// 编号
	Code string `json:"code,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 创建方式
	CreationMethod DataAggregationInventoryCreationMethod `json:"creation_method,omitempty"`
	// 数源单位
	DepartmentPath string `json:"department_path,omitempty"`
	// 相关的业务表列表
	BusinessForms []AggregatedBusinessFormReference `json:"business_forms,omitempty"`
	// 数据归集资源列表
	Resources []AggregatedDataAggregationResource `json:"resources,omitempty"`
	//数据归集资源的数量，就是 Resources 的长度
	ResourcesCount int `json:"resources_count,omitempty"`
	// 工单
	//
	// Deprecated: Use WorkOrderNames instead.
	WorkOrderName string `json:"work_order_name,omitempty"`
	// 工单名称列表
	WorkOrderNames []string `json:"work_order_names,omitempty"`
	// Workflow 文档审核详情
	DocAudit *doc_audit_rest_v1.Apply `json:"doc_audit,omitempty"`
	// 状态
	Status DataAggregationInventoryStatus `json:"status,omitempty"`
	// 创建时间
	CreatedAt meta_v1.Time `json:"created_at,omitempty"`
	// 创建人的显示名称
	CreatorName string `json:"creator_name,omitempty"`
	// 申请时间，因为需要根据申请时间过滤，但 workflow 不支持，所以在本地记录
	RequestedAt *meta_v1.Time `json:"requested_at,omitempty"`
	// 申请人的显示名称，提交归集清单申请的人
	RequesterName string `json:"requester_name,omitempty"`

	AuditApplyID string `json:"audit_apply_id,omitempty"` // 审核申请ID
}

// 聚合的数据归集资源，与 DataAggregationResource 相比包含其引用的其他资源
type AggregatedDataAggregationResource struct {
	// 逻辑视图 ID
	DataViewID string `json:"data_view_id,omitempty"`
	// 资源名称，即逻辑视图的业务名称
	BusinessName string `json:"business_name,omitempty"`
	// 技术名称，即逻辑视图的技术名称
	TechnicalName string `json:"technical_name,omitempty"`
	// 数据来源ID，即资源所属数据源的ID
	DatasourceID string `json:"datasource_id,omitempty"`
	// 数据来源，即资源所属数据源的名称
	DatasourceName string `json:"datasource_name,omitempty"`
	// 数据来源类型，即资源所属数据源的类型
	DatasourceType string `json:"datasource_type,omitempty"`
	// 数源单位，即资源所属部门的路径
	DepartmentPath string `json:"department_path,omitempty"`
	// 采集方式
	CollectionMethod DataAggregationResourceCollectionMethod `json:"collection_method,omitempty"`
	// 同步频率
	SyncFrequency DataAggregationResourceSyncFrequency `json:"sync_frequency,omitempty"`
	// 关联业务表 ID
	BusinessFormID string `json:"business_form_id,omitempty"`
	// 目标数据源 ID
	TargetDatasourceID string `json:"target_datasource_id,omitempty"`
	// 目标数据源，即数据源名称，数据资源被归集到这个数据源
	TargetDatasourceName string `json:"target_datasource_name,omitempty"`
	// 数据库，即数据源的数据库名称
	DatabaseName string `json:"database_name,omitempty"`
	// 价值评估状态
	ValueAssessmentStatus bool `json:"value_assessment_status,omitempty"`
}

// 数据归集资源价值评估状态
type DataAggregationResourceValueAssessmentStatus string

// 聚合的数据归集清单列表
type AggregatedDataAggregationInventoryList meta_v1.List[AggregatedDataAggregationInventory]

// AggregatedBusinessFormReference 代表包含其他资源的业务表引用
type AggregatedBusinessFormReference struct {
	// 业务表 - ID
	ID string `json:"id,omitempty"`
	// 业务表 - 名称
	Name string `json:"name,omitempty"`
	// 业务表 - 描述
	Description string `json:"description,omitempty"`
	// 业务表 - 更新时间
	UpdatedAt meta_v1.Time `json:"updated_at,omitempty"`

	// 业务表 - 业务模型 - 名称
	BusinessModelName string `json:"business_model_name,omitempty"`
	// 业务表 - 业务模型 - 业务域 - 部门 - 路径
	DepartmentPath string `json:"department_path,omitempty"`
	// 业务表 - 信息系统 - 名称
	InfoSystemNames []string `json:"info_system_name,omitempty"`
	// 业务表 - 更新人 - 名称
	UpdaterName string `json:"updater_name,omitempty"`
}

type BusinessFormDataTableItem struct {
	DataTableName      string `json:"data_table_name,omitempty"`       //业务标准表物化的数据表，由第三方加工平台返回
	TargetDataSourceID string `json:"target_data_source_id,omitempty"` //目标数据源 ID
	TargetCatalogName  string `json:"target_catalog_name,omitempty"`   //目标数据源 catalog
	TargetSchema       string `json:"target_schema,omitempty"`         //目标数据源 schema
	BusinessFormID     string `json:"business_form_id,omitempty"`      //关联业务表ID
}
