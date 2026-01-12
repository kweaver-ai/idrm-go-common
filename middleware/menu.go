package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	auth_service "github.com/kweaver-ai/idrm-go-common/rest/auth-service"
)

const API_MENU_KEY = "api_menu_keys"

const (
	SCOPE_ALL        = "all"        //全部
	SCOPE_DEPARTMENT = "department" //本部门
	SCOPE_SELF       = "self"       //仅自己
)
const (
	SCOPE_CN_ALL        = "全部"  //全部
	SCOPE_CN_DEPARTMENT = "本部门" //本部门
	SCOPE_CN_SELF       = "仅自己" //仅自己
)

// Set 设置菜单key，批量设置，适合gin.Group 内部使用
func Set(menus ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(API_MENU_KEY, strings.Join(menus, ","))
		c.Next()
	}
}

type MenuResourceMarkerGenerator struct {
	HandlerGenerator HandlerGenerator
}

type HandlerGenerator func(string, ...string) gin.HandlerFunc

func (m *MenuResourceMarkerGenerator) Set(menu ...string) *MenuResourceMarker {
	return &MenuResourceMarker{
		HandlerGenerator: m.HandlerGenerator,
		MenuKeys:         menu,
	}
}

type MenuResourceMarker struct {
	HandlerGenerator HandlerGenerator
	MenuKeys         []string `json:"menu_keys"`
	Action           string   `json:"action"`
}

func (m *MenuResourceMarker) Create() gin.HandlerFunc {
	m.Action = auth_service.ActionCreate.String
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

func (m *MenuResourceMarker) Update() gin.HandlerFunc {
	m.Action = auth_service.ActionUpdate.String
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

func (m *MenuResourceMarker) Read() gin.HandlerFunc {
	m.Action = auth_service.ActionRead.String
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

func (m *MenuResourceMarker) Delete() gin.HandlerFunc {
	m.Action = auth_service.ActionDelete.String
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

func (m *MenuResourceMarker) Import() gin.HandlerFunc {
	m.Action = auth_service.ActionImport.String
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

func (m *MenuResourceMarker) Offline() gin.HandlerFunc {
	m.Action = auth_service.ActionOffline.String
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

// MenuConstants 菜单常量定义
const (
	MessageSet                     = "messageSet"                     // 消息设置
	GeneralConfig                  = "GeneralConfig"                  // 通用配置
	BusinessSystem                 = "businSystem"                    // 信息系统
	DataQualityWorkOrder           = "dataQualityWorkOrder"           // 质量整改单
	DefineObj                      = "defineObj"                      // 编辑定义
	DataViewDetail                 = "datatViewDetail"                // 库表详情
	ResourceDirAudit               = "resourceDirAudit"               // 目录审核
	DataCatalogFeedback            = "dataCatalogFeedback"            // 目录反馈
	DirContent                     = "dirContent"                     // 目录内容
	AssetAccess                    = "assetAccess"                    // 资源授权
	CodeTable                      = "codetable"                      // 码表
	WorkflowManage                 = "workflowManage"                 // 审核流程模板
	DataQualityWorkOrderAudit      = "dataQualityWorkOrderAudit"      // 整改审核
	PrivacyDataProtection          = "privacyDataProtection"          // 隐私数据保护
	DesensitizationRules           = "desensitizationRules"           // 脱敏算法
	DataClassificationTag          = "dataClassificationTag"          // 数据密级
	AddResourcesDirList            = "addResourcesDirList"            // 新增目录
	InterfaceServiceFeedback       = "interfaceServiceFeedback"       // 接口服务反馈
	CategoryManage                 = "categoryManage"                 // 类目管理
	DataServiceOverview            = "dataServiceOverview"            // 接口服务概览
	DataContent                    = "dataContent"                    // 数据资源目录
	RecognitionAlgorithmConfig     = "recognitionAlgorithmConfig"     // 算法模版
	InterfaceServiceMgt            = "interfaceServiceMgt"            // 接口服务
	DataQualityOverview            = "dataQualityOverview"            // 质量概览
	QualityExamineWorkOrder        = "qualityExamineWorkOrder"        // 质量检测工单
	DataComprehension              = "dataComprehension"              // 数据理解
	CreateDataService              = "createDataService"              // 服务生成
	DataComprehensionContent       = "dataComprehensionContent"       // 理解报告详情
	DataQualityImprovement         = "dataQualityImprovement"         // 质量整改
	DataElement                    = "dataelement"                    // 数据元
	DatasheetView                  = "datasheetView"                  // 库表
	QualityExamineWorkOrderAudit   = "qualityExamineWorkOrderAudit"   // 检测工单审核
	DataDictionary                 = "dataDictionary"                 // 字典管理
	WorkflowManageAuditor          = "workflowManageAuditor"          // 审核员匹配规则模板
	File                           = "file"                           // 标准文件
	DataStandards                  = "dataStandards"                  // 数据标准
	ResourceDirList                = "resourceDirList"                // 目录管理
	BusinessArchitecture           = "businessArchitecture"           // 部门职责管理
	FirmList                       = "firmList"                       // 厂商名录管理
	DataSource                     = "DataSource"                     // 数据源
	ApiServiceDetail               = "apiServiceDetail"               // 服务详情
	CodeRules                      = "coderules"                      // 编码规则
	DataQualityExamine             = "dataQualityExamine"             // 质量检测
	DataQualityWorkOrderProcessing = "dataQualityWorkOrderProcessing" // 整改处理
	FlowCenter                     = "flowCenter"                     // 审核管理
	ImportResourcesDir             = "importResourcesDir"             // 导入
	DataServiceList                = "dataServiceList"                // 接口服务开发
	CreateDataServiceRegistry      = "createDataServiceRegistry"      // 服务注册
	DataResourceCatalogOverview    = "dataResourceCatalogOverview"    // 目录概览
	DataQualityReport              = "dataQualityReport"              // 质量报告
	DataAssets                     = "data-assets"                    // 数据服务超市
	BusinessMatters                = "businessMatters"                // 业务事项
	ComprehensionReportAudit       = "comprehensionReportAudit"       // 理解报告审核
	DataQualityTemplate            = "dataQualityTemplate"            // 质量规则
	InterfaceServiceLog            = "interfaceServiceLog"            // 接口服务日志
	DataComprehensionReport        = "dataComprehensionReport"        // 数据理解报告
	CodingRuleConfig               = "codingRuleConfig"               // 编码生成规则
	Policy                         = "policy"                         // 审核策略
	BusinessDomain                 = "businessDomain"                 // 业务标准
)

// menuKeyToComment 菜单key到中文注释的映射
var menuKeyToComment = map[string]string{
	MessageSet:                     "消息设置",
	GeneralConfig:                  "通用配置",
	BusinessSystem:                 "信息系统",
	DataQualityWorkOrder:           "质量整改单",
	DefineObj:                      "编辑定义",
	DataViewDetail:                 "库表详情",
	ResourceDirAudit:               "目录审核",
	DataCatalogFeedback:            "目录反馈",
	DirContent:                     "目录内容",
	AssetAccess:                    "资源授权",
	CodeTable:                      "码表",
	WorkflowManage:                 "审核流程模板",
	DataQualityWorkOrderAudit:      "整改审核",
	PrivacyDataProtection:          "隐私数据保护",
	DesensitizationRules:           "脱敏算法",
	DataClassificationTag:          "数据密级",
	AddResourcesDirList:            "新增目录",
	InterfaceServiceFeedback:       "接口服务反馈",
	CategoryManage:                 "类目管理",
	DataServiceOverview:            "接口服务概览",
	DataContent:                    "数据资源目录",
	RecognitionAlgorithmConfig:     "算法模版",
	InterfaceServiceMgt:            "接口服务",
	DataQualityOverview:            "质量概览",
	QualityExamineWorkOrder:        "质量检测工单",
	DataComprehension:              "数据理解",
	CreateDataService:              "服务生成",
	DataComprehensionContent:       "理解报告详情",
	DataQualityImprovement:         "质量整改",
	DataElement:                    "数据元",
	DatasheetView:                  "库表",
	QualityExamineWorkOrderAudit:   "检测工单审核",
	DataDictionary:                 "字典管理",
	WorkflowManageAuditor:          "审核员匹配规则模板",
	File:                           "标准文件",
	DataStandards:                  "数据标准",
	ResourceDirList:                "目录管理",
	BusinessArchitecture:           "部门职责管理",
	FirmList:                       "厂商名录管理",
	DataSource:                     "数据源",
	ApiServiceDetail:               "服务详情",
	CodeRules:                      "编码规则",
	DataQualityExamine:             "质量检测",
	DataQualityWorkOrderProcessing: "整改处理",
	FlowCenter:                     "审核管理",
	ImportResourcesDir:             "导入",
	DataServiceList:                "接口服务开发",
	CreateDataServiceRegistry:      "服务注册",
	DataResourceCatalogOverview:    "目录概览",
	DataQualityReport:              "质量报告",
	DataAssets:                     "数据服务超市",
	BusinessMatters:                "业务事项",
	ComprehensionReportAudit:       "理解报告审核",
	DataQualityTemplate:            "质量规则",
	InterfaceServiceLog:            "接口服务日志",
	DataComprehensionReport:        "数据理解报告",
	CodingRuleConfig:               "编码生成规则",
	Policy:                         "审核策略",
	BusinessDomain:                 "业务标准",
}

// GetMenuComments 根据menuKey数组获取中文注释数组
func GetMenuComments(menuKeys []string) []string {
	comments := make([]string, 0, len(menuKeys))
	for _, key := range menuKeys {
		if comment, ok := menuKeyToComment[key]; ok {
			comments = append(comments, comment)
		} else {
			comments = append(comments, key) // 如果没有找到对应注释，返回原key
		}
	}
	return comments
}
