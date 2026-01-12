package errorcode

// Model Name
const (
	drivenModelName = "Common.Driven"
)

// Driven error
const (
	drivenPreCoder                    = drivenModelName + "."
	GetRolesInfo                      = drivenPreCoder + "GetRolesInfo"
	GetInfoSystemDetail               = drivenPreCoder + "GetInfoSystemDetail"
	GetAccessPermissionFailure        = drivenPreCoder + "GetAccessPermissionFailure"
	GetAppListFailure                 = drivenPreCoder + "GetAppListFailure"
	GetDepartmentList                 = drivenPreCoder + "GetDepartmentList"
	GetDepartmentPrecision            = drivenPreCoder + "GetDepartmentPrecision"
	SceneAnalysisDrivenError          = drivenPreCoder + "SceneAnalysisDrivenError"
	SceneAnalysisDrivenGetSceneError  = drivenPreCoder + "SceneAnalysisDrivenGetSceneError"
	ConfigurationServiceInternalError = drivenPreCoder + "ConfigurationServiceInternalError"
	GetDepartmentByPathError          = drivenPreCoder + "GetDepartmentByPathError"
	GetDataSubjectByPathError         = drivenPreCoder + "GetDataSubjectByPathError"
	GetDataSubjectByIDError           = drivenPreCoder + "GetDataSubjectByIDError"
	GetAuditProcessDefinitionError    = drivenPreCoder + "GetAuditProcessDefinitionError"
	GetGradeLabelError                = drivenPreCoder + "GetGradeLabelError"
	GetProcessBindByAuditTypeError    = drivenPreCoder + "GetProcessBindByAuditTypeError"
	DeleteProcessBindByAuditTypeError = drivenPreCoder + "DeleteProcessBindByAuditTypeError"
	AuditProcessNotExist              = drivenPreCoder + "AuditProcessNotExist"
	GetAppsError                      = drivenPreCoder + "GetAppsError"
	UsersRolesFailure                 = drivenPreCoder + "UsersRolesFailure"
	GenerateCodeError                 = drivenPreCoder + "GenerateCodeError"
	GetStandardDictError              = drivenPreCoder + "GetStandardDictError"
	GetSubjectListError               = drivenPreCoder + "GetSubjectListError"
	GetAuditListError                 = drivenPreCoder + "GetAuditListError"
	GetAuditLogsError                 = drivenPreCoder + "GetAuditLogsError"
	GetStandardRuleError              = drivenPreCoder + "GetStandardRuleError"
	GetWorkOrderListError             = drivenPreCoder + "GetWorkOrderListError"
)

var drivenErrorMap = ErrorCode{

	GetAccessPermissionFailure: {
		Description: "配置中心获取访问权限失败",
		Cause:       "",
		Solution:    "",
	},
	GetAppListFailure: {
		Description: "配置中心获取应用列表失败",
		Cause:       "",
		Solution:    "",
	},
	GetRolesInfo: {
		Description: "配置中心获取访问权限失败",
		Cause:       "",
		Solution:    "",
	},
	GetInfoSystemDetail: {
		Description: "配置中心获取信息系统信息失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetDepartmentList: {
		Description: "配置中心获取部门列表失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetDepartmentPrecision: {
		Description: "配置中心获取部门失败",
		Cause:       "",
		Solution:    "请重试",
	},
	SceneAnalysisDrivenError: {
		Description: "场景分析解析错误码失败",
		Cause:       "",
		Solution:    "请重试",
	},
	SceneAnalysisDrivenGetSceneError: {
		Description: "场景分析获取详情失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetDepartmentByPathError: {
		Description: "内部获取部门信息失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetDataSubjectByPathError: {
		Description: "内部获取主题域信息失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetAuditProcessDefinitionError: {
		Description: "workflow获取规则失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetGradeLabelError: {
		Description: "获取分级标签信息错误",
		Solution:    "请重试",
	},
	ConfigurationServiceInternalError: {
		Description: "配置中心服务调用失败",
		Cause:       "",
		Solution:    "请重试或检查服务状态",
	},
	GetProcessBindByAuditTypeError: {
		Description: "调用配置中心服务获取绑定信息失败",
		Cause:       "",
		Solution:    "请重试或检查服务状态",
	},
	DeleteProcessBindByAuditTypeError: {
		Description: "调用配置中心服务删除绑定信息失败",
		Cause:       "",
		Solution:    "请重试或检查服务状态",
	},
	AuditProcessNotExist: {
		Description: "调用配置中心服务获取到的绑定信息与workflow不一致",
		Cause:       "",
		Solution:    "请检查检查服务状态以及确定绑定流程是否正确",
	},
	GetAppsError: {
		Description: "获取应用信息错误",
		Solution:    "请重试",
	},
	UsersRolesFailure: {
		Description: "获取当前用户所拥有的角色列表失败",
	},
	GenerateCodeError: {
		Description: "配置中心生成编码规则失败",
		Cause:       "",
		Solution:    "",
	},
	GetStandardDictError: {
		Description: "获取数据标准码表失败",
		Cause:       "",
		Solution:    "检查standardization-backend服务",
	},
	GetSubjectListError: {
		Description: "获取data-subject列表数据失败",
		Cause:       "",
		Solution:    "检查data-subject服务",
	},
	GetAuditListError: {
		Description: "获取审核列表失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetAuditLogsError: {
		Description: "获取审核日志失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetStandardRuleError: {
		Description: "获取数据标准编码规则失败",
		Cause:       "",
		Solution:    "检查standardization-backend服务",
	},
	GetWorkOrderListError: {
		Description: "获取工单失败",
		Cause:       "",
		Solution:    "检查task-center服务",
	},
}

func init() {
	RegisterErrorCode(drivenErrorMap)
}
