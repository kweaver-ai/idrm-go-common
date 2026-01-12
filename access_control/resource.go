package access_control

//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate $GOPATH/bin/stringer -type=Resource
//go:generate go mod tidy

/**
增加后端资源，用于接口权限控制
步骤：
1. 增加后端资源枚举值
2. 在相关服务对一组资源的接口的路由组中增加AccessControl中间件，即可进行访问控制
3. 在configuration-center服务中给内置角色增加该权限值，通过INSERT SQL
*/

// Resource 后端资源枚举值，必须为正数，且不能与已有资源值重复
type Resource int32

const (
	BusinessDomain    Resource = 1 //业务域 主干业务
	BusinessModel     Resource = 2 //业务模型
	BusinessForm      Resource = 3 //业务表
	BusinessFlowchart Resource = 4 //流程图
	BusinessIndicator Resource = 5 //指标

	UnpublishedBusinessModel     Resource = 6 //任务下主干业务 废弃
	UnpublishedBusinessForm      Resource = 7 //任务下业务表 废弃
	UnpublishedBusinessFlowchart Resource = 8 //任务下流程图 废弃
	UnpublishedBusinessIndicator Resource = 9 //任务下指标 废弃

	Project           Resource = 10 //项目
	Task              Resource = 11 //任务
	OperationLog      Resource = 12 //操作日志
	Flowchart         Resource = 13 //流水线
	Role              Resource = 14 //角色
	BusinessStructure Resource = 15 //业务架构

	NewStandard     Resource = 16 //新建标准
	DataAcquisition Resource = 17 //数据采集
	DataProcessing  Resource = 18 //数据加工
	DataSource      Resource = 19 //数据源

	//GlossaryAndBusinessDomainDefinition Resource = 20 //业务域定义

	CatalogCategory Resource = 21 //目录分类

	DataResourceCatalogBusinessObject Resource = 22 //数据资产目录-业务对象
	DataResourceCatalogDataCatalog    Resource = 23 //数据目录(前台) 我的资产中心、数据服务超市
	DataResourceCatalogDataFeature    Resource = 24 //数据资产目录-需求申请(前台)
	DataResourceCatalogApplyList      Resource = 25 //数据资产目录-申请清单

	DataResourceCatalog Resource = 26 //数据资源目录(后台)
	DataFeature         Resource = 27 //数据需求管理-数据需求申请(后台)
	DataFeatureAnalyze  Resource = 28 //数据需求管理-数据需求分析

	Glossary     Resource = 29 //业务标准-业务术语
	DataStandard Resource = 30 //业务标准-数据标准

	ServiceManagement Resource = 31 //接口服务管理(后台)

	DataExplorationUnderstandingDataResourceCatalogUnderstanding Resource = 32 //数据探查理解-数据资源目录理解

	DataAssetOverview Resource = 34 //数据资产全景

	AuditStrategy Resource = 33 //审核策略
	AuditProcess  Resource = 35 //审核流程
	AuditPending  Resource = 37 //审核待办

	ServiceManagementFront Resource = 38 //接口服务管理(前台)

	InfoSystem          Resource = 39 //信息系统
	SceneAnalysis       Resource = 40 //场景分析
	FormView            Resource = 41 //数据表视图
	IndicatorManagement Resource = 42 // 指标管理
	DataModel           Resource = 43 // 数据建模
	SubjectDomain       Resource = 44 // 主题域
	BusinessDiagnosis   Resource = 45 // 业务诊断
	AuthService         Resource = 46 // 权限服务

	CodeGenerationRule    Resource = 47 // 编码生成规则
	Code                  Resource = 48 // 编码
	DataUsingType         Resource = 49 // 资产或目录类型
	GetDataUsingType      Resource = 50 // 获取资产或目录类型
	TimestampBlacklist    Resource = 51 // 业务更新时间戳黑名单
	DownloadTask          Resource = 52 // 下载任务
	ExploreTask           Resource = 53 // 探查任务
	Completion            Resource = 54 // 业务名称补全
	ApplicationManagement Resource = 55 // 应用管理
	Apps                  Resource = 56 // 集成应用管理

	// 权限服务 - 子视图（行列规则）
	AuthServiceSubView Resource = 57
	// 权限服务 - 逻辑视图授权申请
	AuthServiceLogicViewAuthorizingRequest Resource = 58
	// 权限服务 - 指标授权申请
	AuthServiceIndicatorAuthorizingRequest Resource = 59
	// 授权服务 - 接口授权申请
	AuthServiceAPIAuthorizingRequest Resource = 60
	// 授权服务 - 策略详情
	AuthServicePolicy Resource = 61

	// 审计日志
	Audit Resource = 62

	// 用户
	User Resource = 63

	ExceptSystemMgm Resource = 64 // 泛指非系统管理资源

	DATA_DICT Resource = 65 //数据字典管理

	ProvinceAppsReport Resource = 66 // 省直达应用上报
	AppsAudit          Resource = 67 // 省直达应用上报
	SharedDeclaration  Resource = 68 // 共享申请

	CategoryLabel    Resource = 69 //分类标签
	CategoryAppsAuth Resource = 70 //分类标签应用授权
	WorkOrder        Resource = 71 // 工单

	Firms Resource = 72 // 厂商管理

	// 前置机
	FrontEndProcessor Resource = 73
	// 前置机 - 申请
	FrontEndProcessorRequest Resource = 74
	// 前置机 - 节点分配
	FrontEndProcessorNode Resource = 75
	// 前置机 - 签收
	FrontEndProcessorReceipt Resource = 76
	// 前置机 - 回收
	FrontEndProcessorReclaim Resource = 77
	// 前置机 - 概览
	FrontEndProcessorsOverview Resource = 78

	AddressBook Resource = 79 // 通讯录管理

	SecurityManagement Resource = 80 // 安全管理

	// 供需对接 - 申请清单
	FrontEndRequireList Resource = 81
	// 供需对接 - 分析
	FrontEndRequirAnalysis Resource = 82
	// 供需对接 - 资源确认
	FrontEndRequirConfirm Resource = 83
	// 供需对接 - 实施
	FrontEndRequirImplement Resource = 84
	// 供需对接 - 审核
	FrontEndRequirAudit Resource = 85

	ShareApplication Resource = 86 // 共享申请

	// 共享申请-分析完善
	ShareApplyAnalysis Resource = 87
	// 共享申请-数据资源实施
	ShareApplyImplement Resource = 88

	// 数据处理 - 租户申请管理
	TenantApplyManagement Resource = 89

	// 数据处理 - 数据集管理
	DataSetManagement Resource = 90
)

func (i Resource) ToInt32() int32 {
	return int32(i)
}
