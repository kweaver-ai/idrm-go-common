package access_control

import (
	"reflect"
	"strconv"
)

//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate $GOPATH/bin/stringer -type=Scope
//go:generate go mod tidy

/**
增加前端资源，用于界面的访问控制
步骤：
1. 增加前端资源枚举值
2. 增加前端资源转换结构体映射属性
3. 增加前端资源返回结构体映射属性
4. 在configuration-center服务中给内置角色增加该权限值，通过INSERT SQL
*/

// Scope 1.前端资源枚举值，必须为负数,且不能与已有资源值重复
type Scope int32

const (
	BusinessDomainScope           Scope = -1  //业务域
	BusinessStructureScope        Scope = -2  //业务架构
	BusinessModelScope            Scope = -3  //主干业务
	BusinessFormScope             Scope = -4  //业务表
	BusinessFlowchartScope        Scope = -5  //业务流程图
	BusinessIndicatorScope        Scope = -6  //指标
	ProjectScope                  Scope = -8  //项目列表
	PipelineKanbanScope           Scope = -9  //流水线看板
	TaskKanbanScope               Scope = -10 //任务看板
	TaskScope                     Scope = -11 //任务列表
	PipelineScope                 Scope = -12 //流水线
	RoleScope                     Scope = -13 //角色
	BusinessStandardScope         Scope = -14 //业务标准
	BusinessKnowledgeNetworkScope Scope = -15 //业务知识网络
	DataConnectScope              Scope = -17 //数据连接
	MetadataScope                 Scope = -18 //元数据管理
	DataSecurityScope             Scope = -19 //数据安全
	DataQualityScope              Scope = -20 //数据质量
	DataUnderstandScope           Scope = -22 //数据理解
	TaskBusinessModel             Scope = -26 //任务下主干业务
	TaskBusinessForm              Scope = -27 //任务下业务表
	TaskBusinessFlowchart         Scope = -28 //任务下业务流程图
	TaskBusinessIndicator         Scope = -29 //任务下指标
	NewStandardScope              Scope = -30 //新建标准
	TaskNewStandardScope          Scope = -31 //任务下新建标准
	DataAcquisitionScope          Scope = -32 //数据采集
	TaskDataAcquisitionScope      Scope = -33 //任务下数据采集
	DataProcessingScope           Scope = -34 //数据加工
	TaskDataProcessingScope       Scope = -35 //任务下数据加工
	TaskStdCreateScope            Scope = -36 //标准化-标准创建任务表
	DataSourceScope               Scope = -37 //数据源

	//GlossaryAndBusinessDomainDefinitionScope Scope = -38 //术语表和业务域定义
	CatalogCategoryScope Scope = -39 //目录分类

	DataAssetOverviewScope                 Scope = -40 //数据资产全景 无对应后端接口
	DataResourceCatalogBusinessObjectScope Scope = -41 //数据资产目录-业务对象 22
	DataResourceCatalogDataCatalogScope    Scope = -42 //数据目录(前台) 23
	DataResourceCatalogDataFeatureScope    Scope = -43 //数据资产目录-需求申请(前台) 24
	DataResourceCatalogApplyListScope      Scope = -44 //数据资产目录-申请清单 25

	DataResourceCatalogScope Scope = -45 //数据资源目录(后台) 26   数据运营管理-数据目录管理-目录分类-39  21
	DataFeatureScope         Scope = -46 //数据需求管理-数据需求申请(后台) 27
	DataFeatureAnalyzeScope  Scope = -47 //数据需求管理-数据需求分析 28

	GlossaryScope     Scope = -48 //业务标准-业务术语 29
	DataStandardScope Scope = -49 //业务标准-数据标准-数据元 30

	ServiceManagementScope Scope = -50 //接口服务管理(后台) 31

	DataExplorationUnderstandingDataResourceCatalogUnderstandingScope Scope = -51 //数据探查理解-数据资源目录理解

	IndependentDataAcquisitionTask Scope = -52 //独立任务 数据采集任务  TaskScope -11  任务子类
	IndependentDataProcessingTask  Scope = -53 //独立任务 数据加工任务  TaskScope -11  任务子类
	AuditProcessScope              Scope = -54 //审核流程
	AuditStrategyScope             Scope = -55 //审核策略
	AuditPendingScope              Scope = -56 //审核待办

	ServiceManagementFrontScope Scope = -57 //接口服务管理(前台) 33

	InfoSystemScope          Scope = -58 //信息系统 39
	SceneAnalysisScope       Scope = -59 //场景分析 40
	FormViewScope            Scope = -60 //数据表视图 41
	IndicatorManagementScope Scope = -61 // 指标管理42
	DataModelScope           Scope = -62 // 数据建模43
	SubjectDomainScope       Scope = -63 // 数据建模44
	BusinessDiagnosisScope   Scope = -64 // 业务诊断
	AuthServiceScope         Scope = -65 // 权限服务

	CodeGenerationRuleScope    Scope = -66 // 编码生成规则 47
	CodeScope                  Scope = -67 // 编码生成规则 48
	DataUsingTypeScope         Scope = -68 // 资源类型管理
	GetDataUsingTypeScope      Scope = -69 // 获取资源类型管理
	TimestampBlacklistScope    Scope = -70 // 业务更新时间戳黑名单
	DownloadTaskScope          Scope = -71 // 下载任务
	ExploreTaskScope           Scope = -72 // 探查任务
	CompletionScope            Scope = -73 // 业务名称补全
	ApplicationManagementScope Scope = -74 // 应用管理
	AppsScope                  Scope = -75 // 集成应用管理

	// 权限服务 - 子视图（行列规则）
	AuthServiceSubViewsScope Scope = -76
	// 权限服务 - 逻辑视图授权申请
	AuthServiceLogicViewAuthorizingRequestScope Scope = -77
	// 权限服务 - 指标授权申请
	AuthServiceIndicatorAuthorizingRequestScope Scope = -78
	// 权限服务 - 接口授权申请
	AuthServiceAPIAuthorizingRequestScope Scope = -79

	// 审计日志
	AuditScope Scope = -80

	DATA_DICT_Scope Scope = -81 //数据字典管理

	ProvinceAppsReportScope Scope = -82 // 集成应用上报管理
	AppsAuditScope          Scope = -83 // 集成应用审核管理
	SharedDeclarationScope  Scope = -84 // 共享申请
	CategoryLabelScope      Scope = -85 //分类标签
	CategoryAppsAuthScope   Scope = -86 //分类标签应用授权
	WorkOrderScope          Scope = -87 // 工单

	FirmsScope Scope = -88 // 厂商管理

	// 前置机概览
	FrontEndProcessorsOverviewScope Scope = -89
	// 前置机申请
	FrontEndProcessorRequestScope Scope = -90
	// 前置机签收
	FrontEndProcessorReceiptScope Scope = -91
	// 前置机审核
	FrontEndProcessorAuditScope Scope = -92
	// 前置机分配
	FrontEndProcessorAllocationScope Scope = -93
	HomePage                         Scope = -94 // 首页

	AddressBookScope Scope = -95 // 通讯录管理

	SecurityManagementScope Scope = -96 //

	// 供需对接 - 申请清单
	FrontEndRequireListScope Scope = -97
	// 供需对接 - 分析
	FrontEndRequirAnalysisScope Scope = -98
	// 供需对接 - 资源确认
	FrontEndRequirConfirmScope Scope = -99
	// 供需对接 - 实施
	FrontEndRequirImplementScope Scope = -100
	// 供需对接 - 审核
	FrontEndRequirAuditScope Scope = -101

	// 业务认知分析平台 - 数据查询
	CognitionAnalysisDataQueryScope Scope = -102

	// 共享申请 - 分析完善
	ShareApplyAnalysisScope Scope = -103
	// 共享申请 - 数据资源实施
	ShareApplyImplementScope Scope = -104
	// 数据处理 - 租户申请管理
	TenantApplyManagementScope Scope = -105
	// 工作专区
	PlatformZoneScope Scope = -106
	// 工作专区管理
	PlatformZoneManagementScope Scope = -107
	// 平台服务
	PlatformServiceScope Scope = -108
	// 平台服务管理
	PlatformServiceManagementScope Scope = -109
	// 数据反馈管理
	FeedbackManagementScope Scope = -110
)

// AccessControl 3.前端资源返回结构体，增加的属性名字需与 （2.前端资源转换结构体 ScopeTransfer）中对应
type AccessControl struct {
	Normal Normal     `json:"normal"`
	Task   TaskStruct `json:"task"`
}
type Normal struct {
	NormalBusinessDomain    int32 `json:"business_domain"`         //业务域
	NormalBusinessStructure int32 `json:"enterprise_architecture"` //业务架构
	NormalBusinessModel     int32 `json:"business_model"`          //主干业务
	NormalBusinessForm      int32 `json:"business_form"`           //业务表
	NormalBusinessFlowchart int32 `json:"business_flowchart"`      //业务流程图
	NormalBusinessIndicator int32 `json:"business_indicator"`      //指标
	//NormalBusinessReport           int32 `json:"business_report"`            //业务诊断
	NormalProject                  int32 `json:"project"`                    //项目列表
	NormalPipelineKanban           int32 `json:"pipeline_kanban"`            //流水线看板
	NormalTaskKanban               int32 `json:"task_kanban"`                //任务看板
	NormalTask                     int32 `json:"task"`                       //任务列表
	NormalPipeline                 int32 `json:"pipeline"`                   //流水线
	NormalRole                     int32 `json:"role"`                       //角色
	NormalBusinessStandard         int32 `json:"business_standard"`          //业务标准
	NormalBusinessKnowledgeNetwork int32 `json:"business_knowledge_network"` //业务知识网络
	NormalDataAcquisition          int32 `json:"data_acquisition"`           //数据采集
	NormalDataConnect              int32 `json:"data_connection"`            //数据连接
	NormalMetadata                 int32 `json:"metadata"`                   //元数据管理
	NormalDataSecurity             int32 `json:"data_security"`              //数据安全
	NormalDataQuality              int32 `json:"data_quality"`               //数据质量
	NormalDataProcessing           int32 `json:"data_processing"`            //数据加工
	NormalDataUnderstand           int32 `json:"data_understand"`            //数据理解
	NormalNewStandard              int32 `json:"new_standard"`               //新建标准
	NormalDataSource               int32 `json:"datasource"`                 //数据源

	//NormalGlossaryAndBusinessDomainDefinition int32 `json:"glossary_and_business_domain_definition"` //术语表和业务域定义
	NormalCatalogCategory int32 `json:"catalog_category"` //目录分类

	NormalDataAssetOverviewScope                 int32 `json:"data_asset_overview"`
	NormalDataResourceCatalogBusinessObjectScope int32 `json:"data_resource_catalog_business_object"`
	NormalDataResourceCatalogDataCatalogScope    int32 `json:"data_resource_catalog_data_catalog"`
	NormalDataResourceCatalogDataFeatureScope    int32 `json:"data_resource_catalog_data_feature"`
	NormalDataResourceCatalogApplyListScope      int32 `json:"data_resource_catalog_apply_list"`
	NormalDataResourceCatalogScope               int32 `json:"data_resource_catalog"`
	NormalDataFeatureScope                       int32 `json:"data_feature"`
	NormalDataFeatureAnalyzeScope                int32 `json:"data_feature_analyze"`

	NormalGlossaryScope     int32 `json:"glossary"`
	NormalDataStandardScope int32 `json:"data_standard"`

	NormalServiceManagementScope int32 `json:"service_management"`

	NormalDataExplorationUnderstandingDataResourceCatalogUnderstandingScope int32 `json:"data_exploration_understanding_data_resource_catalog_understanding"`

	NormalIndependentDataAcquisitionTask int32 `json:"independent_data_acquisition_task"`
	NormalIndependentDataProcessingTask  int32 `json:"independent_data_processing_task"`
	AuditProcessScope                    int32 `json:"audit_process"`
	AuditStrategyScope                   int32 `json:"audit_strategy"`
	AuditPendingScope                    int32 `json:"audit_pending"`

	NormalServiceManagementFrontScope int32 `json:"service_management_front"`

	InfoSystemScope          int32 `json:"info_system"`
	SceneAnalysisScope       int32 `json:"scene_analysis"`
	FormViewScope            int32 `json:"form_view"`
	IndicatorManagementScope int32 `json:"indicator_management"`
	DataModelScope           int32 `json:"data_model"`
	SubjectDomainScope       int32 `json:"subject_domain"`
	BusinessDiagnosisScope   int32 `json:"business_diagnosis_scope"` // 业务诊断
	AuthServiceScope         int32 `json:"auth_service"`             // 权限服务

	CodeGenerationRuleScope    int32 `json:"code_generation_rule"`
	DataUsingTypeScope         int32 `json:"general_config"`
	TimestampBlacklistScope    int32 `json:"timestamp_blacklist"`
	DownloadTaskScope          int32 `json:"download_task"`
	ExploreTaskScope           int32 `json:"explore_task"`
	CompletionScope            int32 `json:"completion"`
	ApplicationManagementScope int32 `json:"application_management"`
	AppsScope                  int32 `json:"apps"`
	ProvinceAppsReportScope    int32 `json:"province_apps_report"`
	AppsAuditScope             int32 `json:"apps_audit"`
	DATA_DICT_Scope            int32 `json:"dict"`

	// 权限服务 - 子视图（行列规则）
	AuthServiceSubViewScope int32 `json:"auth_service_sub_view"`
	// 权限服务 - 逻辑视图授权申请
	AuthServiceLogicViewAuthorizingRequestScope int32 `json:"auth_service_logic_view_authorizing_request"`
	// 权限服务 - 指标授权申请
	AuthServiceIndicatorAuthorizingRequestScope int32 `json:"auth_service_indicator_authorizing_request"`
	// 权限服务 - 接口授权申请
	AuthServiceAPIAuthorizingRequestScope int32 `json:"auth_service_api_authorizing_request"`

	// 审计日志
	AuditScope             int32 `json:"audit"`
	SharedDeclarationScope int32 `json:"shared_declaration"`
	// 数据反馈管理
	FeedbackManagementScope int32 `json:"data_feedback_management"`
}

type TaskStruct struct {
	TaskBusinessModel     int32 `json:"business_model"`     //主干业务
	TaskBusinessForm      int32 `json:"business_form"`      //业务表
	TaskBusinessFlowchart int32 `json:"business_flowchart"` //业务流程图
	TaskBusinessIndicator int32 `json:"business_indicator"` //指标
	TaskNewStandard       int32 `json:"new_standard"`       //新建标准
	TaskDataAcquisition   int32 `json:"data_acquisition"`   //数据采集
	TaskDataProcessing    int32 `json:"data_processing"`    //数据加工
	TaskStdCreate         int32 `json:"std_create"`         //标准化-标准创建任务表

	TaskDataExplorationUnderstandingDataResourceCatalogUnderstanding int32 `json:"data_exploration_understanding_data_resource_catalog_understanding"`
}

// ScopeTransfer 2.前端资源转换结构体
type ScopeTransfer struct {
	NormalBusinessDomain    int32 `type:"-1"` //业务域
	NormalBusinessStructure int32 `type:"-2"` //业务架构
	NormalBusinessModel     int32 `type:"-3"` //主干业务
	NormalBusinessForm      int32 `type:"-4"` //业务表
	NormalBusinessFlowchart int32 `type:"-5"` //业务流程图
	NormalBusinessIndicator int32 `type:"-6"` //指标
	//NormalBusinessReport           int32 `type:"-7"`  //业务诊断
	NormalProject                  int32 `type:"-8"`  //项目列表
	NormalPipelineKanban           int32 `type:"-9"`  //流水线看板
	NormalTaskKanban               int32 `type:"-10"` //任务看板
	NormalTask                     int32 `type:"-11"` //任务列表
	NormalPipeline                 int32 `type:"-12"` //流水线
	NormalRole                     int32 `type:"-13"` //角色
	NormalBusinessStandard         int32 `type:"-14"` //业务标准
	NormalBusinessKnowledgeNetwork int32 `type:"-15"` //业务知识网络
	NormalDataConnect              int32 `type:"-17"` //数据连接
	NormalMetadata                 int32 `type:"-18"` //元数据管理
	NormalDataSecurity             int32 `type:"-19"` //数据安全
	NormalDataQuality              int32 `type:"-20"` //数据质量
	NormalDataUnderstand           int32 `type:"-22"` //数据理解
	BusinessModelingTask           int32 `type:"-23"` //业务建模任务
	BusinessStandardizationTask    int32 `type:"-24"` //业务标准化任务
	BusinessIndicatorTask          int32 `type:"-25"` //业务指标梳理任务
	TaskBusinessModel              int32 `type:"-26"` //任务下主干业务
	TaskBusinessForm               int32 `type:"-27"` //任务下业务表
	TaskBusinessFlowchart          int32 `type:"-28"` //任务下业务流程图
	TaskBusinessIndicator          int32 `type:"-29"` //任务下指标
	NormalNewStandard              int32 `type:"-30"` //新建标准
	TaskNewStandard                int32 `type:"-31"` //任务下新建标准
	NormalDataAcquisition          int32 `type:"-32"` //数据采集
	TaskDataAcquisition            int32 `type:"-33"` //任务下数据采集
	NormalDataProcessing           int32 `type:"-34"` //数据加工
	TaskDataProcessing             int32 `type:"-35"` //任务下数据加工
	TaskStdCreate                  int32 `type:"-36"` //标准化-标准创建任务表
	NormalDataSource               int32 `type:"-37"` //数据源

	//NormalGlossaryAndBusinessDomainDefinition int32 `type:"-38"` //术语表和业务域定义
	NormalCatalogCategory int32 `type:"-39"` //目录分类

	NormalDataAssetOverviewScope                 int32 `type:"-40"`
	NormalDataResourceCatalogBusinessObjectScope int32 `type:"-41"`
	NormalDataResourceCatalogDataCatalogScope    int32 `type:"-42"`
	NormalDataResourceCatalogDataFeatureScope    int32 `type:"-43"`
	NormalDataResourceCatalogApplyListScope      int32 `type:"-44"`

	NormalDataResourceCatalogScope int32 `type:"-45"`
	NormalDataFeatureScope         int32 `type:"-46"`
	NormalDataFeatureAnalyzeScope  int32 `type:"-47"`

	NormalGlossaryScope     int32 `type:"-48"`
	NormalDataStandardScope int32 `type:"-49"`

	NormalServiceManagementScope int32 `type:"-50"`

	NormalDataExplorationUnderstandingDataResourceCatalogUnderstandingScope int32 `type:"-51"`
	TaskDataExplorationUnderstandingDataResourceCatalogUnderstanding        int32 `type:"-51"`

	NormalIndependentDataAcquisitionTask int32 `type:"-52"`
	NormalIndependentDataProcessingTask  int32 `type:"-53"`
	AuditProcessScope                    int32 `type:"-54"`
	AuditStrategyScope                   int32 `type:"-55"`
	AuditPendingScope                    int32 `type:"-56"`

	NormalServiceManagementFrontScope int32 `type:"-57"`

	InfoSystemScope          int32 `type:"-58"`
	SceneAnalysisScope       int32 `type:"-59"`
	FormViewScope            int32 `type:"-60"`
	IndicatorManagementScope int32 `type:"-61"`
	DataModelScope           int32 `type:"-62"`
	SubjectDomainScope       int32 `type:"-63"`
	BusinessDiagnosisScope   int32 `type:"-64"` // 业务诊断
	AuthServiceScope         int32 `type:"-65"` // 权限服务

	CodeGenerationRuleScope    int32 `type:"-66"`
	DataUsingTypeScope         int32 `type:"-68"`
	TimestampBlacklistScope    int32 `type:"-70"`
	DownloadTaskScope          int32 `type:"-71"`
	ExploreTaskScope           int32 `type:"-72"`
	CompletionScope            int32 `type:"-73"`
	ApplicationManagementScope int32 `type:"-74"`
	AppsScope                  int32 `type:"-75"`

	// 权限服务 - 子视图（行列规则）
	AuthServiceSubViewScope int32 `type:"-76"`
	// 权限服务 - 逻辑视图授权申请
	AuthServiceLogicViewAuthorizingRequestScope int32 `type:"-77"`
	// 权限服务 - 指标授权申请
	AuthServiceIndicatorAuthorizingRequestScope int32 `type:"-78"`
	// 权限服务 - 接口授权申请
	AuthServiceAPIAuthorizingRequestScope int32 `type:"-79"`

	// 审计日志
	AuditScope int32 `type:"-80"`

	DATA_DICT_Scope int32 `type:"-81"`

	ProvinceAppsReportScope int32 `type:"-82"`
	AppsAuditScope          int32 `type:"-83"`
	SharedDeclarationScope  int32 `type:"-84"`
	// 数据反馈管理
	FeedbackManagementScope int32 `type:"-110"`
}

//————————————————————————————————————————————————————分割线——————————————————————————————————————————————————

func (r *ScopeTransfer) SetValue(tag Scope, value int32) {
	rt := reflect.TypeOf(*r)
	for k := 0; k < rt.NumField(); k++ {
		if rt.Field(k).Tag.Get("type") == strconv.Itoa(int(tag)) {
			reflect.ValueOf(r).Elem().FieldByName(rt.Field(k).Name).SetInt(int64(value))
			//return
		}
	}
}
func (i Scope) ToInt32() int32 {
	return int32(i)
}

// ManageScopeSet 管理员权限，拥有以下资源即拥有管理员权限
var ManageScopeSet = map[Scope]struct{}{
	BusinessDomainScope:    {},
	BusinessStructureScope: {},
	BusinessModelScope:     {},
	BusinessFormScope:      {},
	BusinessFlowchartScope: {},
	BusinessIndicatorScope: {},
	ProjectScope:           {},
	PipelineKanbanScope:    {},
	TaskKanbanScope:        {},
	TaskScope:              {},
	PipelineScope:          {},
}

func (i Scope) IsSubResource() bool {
	if _, exist := ManageScopeSet[i]; exist {
		return true
	}
	return false
}
