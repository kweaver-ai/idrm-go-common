package v1

// 审核类型
type AuditType string

// 审核类型
const (
	AuditTypeAppApplyEscalate AuditType = "af-sszd-app-apply-escalate"

	AuditTypeAppApplyReport AuditType = "af-sszd-app-report-escalate"

	AuditTypeAuthServicePermissionRequest AuditType = "af-data-permission-request"
	// 业务标签授权审核
	AuditTypeBigdataAuthCategoryLabel AuditType = "af-basic-bigdata-auth-category-label"
	// 业务标签分类发布审核
	AuditTypeBigdataCreateCategoryLabel AuditType = "af-basic-bigdata-create-category-label"
	// 业务标签分类删除审核
	AuditTypeBigdataDeleteCategoryLabel AuditType = "af-basic-bigdata-delete-category-label"
	// 业务标签分类变更审核
	AuditTypeBigdataUpdateCategoryLabel AuditType = "af-basic-bigdata-update-category-label"

	AuditTypeDataCatalogChange AuditType = "af-data-catalog-change"

	AuditTypeDataCatalogDownload AuditType = "af-data-catalog-download"

	AuditTypeDataCatalogOffline AuditType = "af-data-catalog-offline"

	AuditTypeDataCatalogOnline AuditType = "af-data-catalog-online"

	AuditTypeDataCatalogOpen AuditType = "af-data-catalog-open"

	AuditTypeDataCatalogPublish AuditType = "af-data-catalog-publish"

	AuditTypeDataViewAuditTypeOffline AuditType = "af-data-view-offline"

	AuditTypeDataViewAuditTypeOnline AuditType = "af-data-view-online"

	AuditTypeDataViewAuditTypePublish AuditType = "af-data-view-publish"

	AuditTypeFrontEndProcessorRequest AuditType = "af-front-end-processor-request"

	AuditTypeInfoCatalogOffline AuditType = "af-info-catalog-offline"

	AuditTypeInfoCatalogOnline AuditType = "af-info-catalog-online"

	AuditTypeInfoCatalogPublish AuditType = "af-info-catalog-publish"
	// 数据归集清单审核
	AuditTypeTasksDataAggregationInventoryRequest AuditType = "af-data-aggregation-inventory"

	AuditTypeTasksDataAggregationPlan AuditType = "af-task-center-data-aggregation-plan"

	AuditTypeTasksDataProcessingPlan AuditType = "af-task-center-data-processing-plan"

	AuditTypeTasksDataSearchReport AuditType = "af-task-center-data-search-report"
)

// 服务类型，一个服务类型对应多个审核类型。
type ServiceType string

// 返回服务类型 st 是否包含的审核类型 at
func (st ServiceType) Contains(at AuditType) bool {
	for _, tt := range serviceTypeAuditTypeBindings[st] {
		if tt == at {
			return true
		}
	}
	return false
}

// 服务类型，一个服务类型对应多个审核类型。
const (
	// 授权服务
	ServiceTypeAuthService ServiceType = "auth-service"

	ServiceTypeBasicBigdataService ServiceType = "basic-bigdata-service"

	ServiceTypeConfigurationCenter ServiceType = "configuration-center"

	ServiceTypeDataCatalog ServiceType = "data-catalog"

	ServiceTypeDataView ServiceType = "data-view"

	ServiceTypeOpenCatalog ServiceType = "open-catalog"
	// 任务中心
	ServiceTypeTasks ServiceType = "task-center"
)

// 服务类型和审核类型的绑定关系
var serviceTypeAuditTypeBindings = map[ServiceType][]AuditType{
	ServiceTypeDataView: {
		AuditTypeDataViewAuditTypePublish,
		AuditTypeDataViewAuditTypeOnline,
		AuditTypeDataViewAuditTypeOffline,
	},
	ServiceTypeAuthService: {
		AuditTypeAuthServicePermissionRequest,
	},
	ServiceTypeDataCatalog: {
		AuditTypeDataCatalogPublish,
		AuditTypeDataCatalogOnline,
		AuditTypeDataCatalogOffline,
		AuditTypeInfoCatalogPublish,
		AuditTypeInfoCatalogOnline,
		AuditTypeInfoCatalogOffline,
	},
	ServiceTypeConfigurationCenter: {
		AuditTypeAppApplyEscalate,
		AuditTypeAppApplyReport,
		AuditTypeFrontEndProcessorRequest,
	},
	ServiceTypeTasks: {
		AuditTypeTasksDataAggregationInventoryRequest,
		AuditTypeTasksDataAggregationPlan,
		AuditTypeTasksDataProcessingPlan,
		AuditTypeTasksDataSearchReport,
	},
	ServiceTypeBasicBigdataService: {
		AuditTypeBigdataCreateCategoryLabel,
		AuditTypeBigdataUpdateCategoryLabel,
		AuditTypeBigdataDeleteCategoryLabel,
		AuditTypeBigdataAuthCategoryLabel,
	},
	ServiceTypeOpenCatalog: {
		AuditTypeDataCatalogOpen,
	},
}
