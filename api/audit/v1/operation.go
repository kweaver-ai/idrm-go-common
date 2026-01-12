package v1

import "github.com/kweaver-ai/idrm-go-common/util/sets"

// Operation 定义了可唯一标识的操作动作
type Operation string

// Operation 枚举名称使用 CamelCase，包括动作和资源两部分。例如 ModifyLogicView
//
// Operation 枚举值使用 snake_case，与名称对应。例如 modify_logic_view
//
// TODO: 补充逻辑视图、维度模型、指标相关 Operation
const (
	//用户登录
	OperationLogin Operation = "login"
	// 用户登出
	OperationLogout Operation = "logout"
	// 生成接口
	OperationGenerateAPI Operation = "generate_api"
	// 注册接口
	OperationRegisterAPI Operation = "register_api"
	// 修改接口
	OperationUpdateAPI Operation = "update_api"
	// 设置接口权限
	OperationSetAPIAuthorization Operation = "set_api_authorization"
	// 发布接口
	OperationPublicAPI Operation = "public_api"
	// 上线接口
	OperationUpAPI Operation = "up_api"
	// 下线接口
	OperationDownAPI Operation = "down_api"
	// 删除接口
	OperationDeleteAPI Operation = "delete_api"
	// 查看接口调用信息
	OperationGetAPIAuthInfo Operation = "get_api_auth_info"
	// 申请接口权限
	OperationRequestAPIAuthorization = "request_api_authorization"
	// 授权接口权限
	OperationAuthorizeAPI Operation = "authorize_api"
	// 创建数据下载任务
	OperationCreateDataDownloadTask Operation = "create_data_download_task"
	// 删除数据下载任务
	OperationDeleteDataDownloadTask Operation = "delete_data_download_task"
	// 申请逻辑视图权限
	OperationRequestDataViewAuthorization Operation = "request_data_view_authorization"
	// 授权逻辑视图权限
	OperationAuthorizeDataView Operation = "authorize_data_view"
	// 数据下载
	OperationDataDownload Operation = "data_download"
	// 预览逻辑视图数据
	OperationDataPreview Operation = "data_preview"
	// 创建维度模型
	OperationCreateDimensionModel Operation = "create_dimension_model"
	// 修改维度模型基本信息
	OperationUpdateDimensionModelBasicInfo Operation = "update_dimension_model_basic_info"
	// 修改维度模型配置信息
	OperationUpdateDimensionModelConfigInfo Operation = "update_dimension_model_config_info"
	// 删除维度模型
	OperationDeleteDimensionModel Operation = "delete_dimension_model"
	// 创建指标
	OperationCreateIndicator Operation = "create_indicator"
	// 编辑指标
	OperationUpdateIndicator Operation = "update_indicator"
	// 删除指标
	OperationDeleteIndicator Operation = "delete_indicator"
	// 查询指标结果
	OperationQueryIndicatorResult Operation = "query_indicator"
	// 申请指标权限
	OperationRequestIndicatorAuthorization Operation = "request_indicator_authorization"
	// 授权指标权限
	OperationAuthorizeIndicator Operation = "authorize_indicator"
	// 创建业务域分组
	OperationCreateBusinessDomainGroup Operation = "create_business_domain_group"
	// 编辑业务域分组
	OperationUpdateBusinessDomainGroup Operation = "update_business_domain_group"
	// 删除业务域分组
	OperationDeleteBusinessDomainGroup Operation = "delete_business_domain_group"
	// 创建业务域
	OperationCreateBusinessDomain Operation = "create_business_domain"
	// 编辑业务域
	OperationUpdateBusinessDomain Operation = "update_business_domain"
	// 删除业务域
	OperationDeleteBusinessDomain Operation = "delete_business_domain"
	// 创建主干业务
	OperationCreateBusinessProcess Operation = "create_business_process"
	// 编辑主干业务
	OperationUpdateBusinessProcess Operation = "update_business_process"
	// 删除主干业务
	OperationDeleteBusinessProcess Operation = "delete_business_process"
	// 创建业务模型
	OperationCreateBusinessModel Operation = "create_business_model"
	// 编辑业务模型
	OperationUpdateBusinessModel Operation = "update_business_model"
	// 删除业务模型
	OperationDeleteBusinessModel Operation = "delete_business_model"
	// 创建数据模型
	OperationCreateDataModel Operation = "create_data_model"
	// 编辑数据模型
	OperationUpdateDataModel Operation = "update_data_model"
	// 删除数据模型
	OperationDeleteDataModel Operation = "delete_data_model"
	// 创建业务表
	OperationCreateBusinessForms Operation = "create_business_form"
	// 编辑业务表
	OperationUpdateBusinessForms Operation = "update_business_form"
	// 删除业务表
	OperationDeleteBusinessForms Operation = "delete_business_form"
	// 修改业务表表结构
	OperationUpdateBusinessFormsContent Operation = "update_business_form_content"
	// 创建数据表
	OperationCreateDataForms Operation = "create_data_form"
	// 编辑数据表
	OperationUpdateDataForms Operation = "update_data_form"
	// 删除数据表
	OperationDeleteDataForms Operation = "delete_data_form"
	// 修改数据表表结构
	OperationUpdateDataFormsContent Operation = "update_data_form_content"
	// 创建业务流程图
	OperationCreateBusinessFlowcharts Operation = "create_business_flowcharts"
	// 编辑业务流程图
	OperationUpdateBusinessFlowcharts Operation = "update_business_flowcharts"
	// 删除业务流程图
	OperationDeleteBusinessFlowcharts Operation = "delete_business_flowcharts"
	// 修改业务流程图内容
	OperationUpdateBusinessFlowchartsContent Operation = "update_business_flowcharts_content"
	//导出业务流程图
	OperationExportBusinessFlowcharts Operation = "export_business_flowcharts"
	// 创建业务指标
	OperationCreateBusinessIndicator Operation = "create_business_indicator"
	// 编辑业务指标
	OperationUpdateBusinessIndicator Operation = "update_business_indicator"
	// 删除业务指标
	OperationDeleteBusinessIndicator Operation = "delete_business_indicator"
	// 创建数据指标
	OperationCreateDataIndicator Operation = "create_data_indicator"
	// 编辑数据指标
	OperationUpdateDataIndicator Operation = "update_data_indicator"
	// 删除数据指标
	OperationDeleteDataIndicator Operation = "delete_data_indicator"
	// 创建主题域分组
	OperationCreateSubjectDomainGroup Operation = "create_subject_domain_group"
	// 编辑主题域分组
	OperationUpdateSubjectDomainGroup Operation = "update_subject_domain_group"
	// 删除主题域分组
	OperationDeleteSubjectDomainGroup Operation = "delete_subject_domain_group"
	// 创建主题域
	OperationCreateSubjectDomain Operation = "create_subject_domain"
	// 编辑主题域
	OperationUpdateSubjectDomain Operation = "update_subject_domain"
	// 删除主题域
	OperationDeleteSubjectDomain Operation = "delete_subject_domain"
	// 创建业务对象
	OperationCreateBusinessObject Operation = "create_business_object"
	// 编辑业务对象基本信息
	OperationUpdateBusinessObject Operation = "update_business_object"
	// 删除业务对象
	OperationDeleteBusinessObject Operation = "delete_business_object"
	// 编辑业务对象内容
	OperationUpdateBusinessObjectContent Operation = "update_business_object_content"
	// 创建业务活动
	OperationCreateBusinessActivity Operation = "create_business_activity"
	// 编辑业务活动基本信息
	OperationUpdateBusinessActivity Operation = "update_business_activity"
	// 删除业务活动
	OperationDeleteBusinessActivity Operation = "delete_business_activity"
	// 编辑业务活动内容
	OperationUpdateBusinessActivityContent Operation = "update_business_activity_content"
	// 新建逻辑视图
	OperationCreateLogicView Operation = "create_logic_view"
	// 修改逻辑视图
	OperationUpdateLogicView Operation = "update_logic_view"
	// 删除逻辑视图
	OperationDeleteLogicView Operation = "delete_logic_view"
	// 扫描数据源
	OperationScanDataSource Operation = "scan_data_source"
	// 上线逻辑视图
	OperationOnlineLogicView Operation = "online_logic_view"
	// 下线逻辑视图
	OperationOfflineLogicView Operation = "offline_logic_view"
	// 发布逻辑视图
	OperationPublishLogicView Operation = "publish_logic_view"

	// 数据标准相关===开始====
	CREATE_DATAELEMENT_API Operation = "create_dataelement_api"
	//EXPORT_DATAELEMENT_API          Operation = "export_dataelement_api"
	//EXPORT_IDS_DATAELEMENT_API      Operation = "export_ids_dataelement_api"
	UPDATE_DATAELEMENT_API       Operation = "update_dataelement_api"
	BATCH_DELETE_DATAELEMENT_API Operation = "batch_delete_dataelement_api"
	//MOVE_CATALOG_DATAELEMENT_API    Operation = "move_catalog_dataelement_api"
	//DELETE_FJLABEL_DATAELEMENT_API  Operation = "delete_fjLabel_dataelement_api"
	//CREATE_DATA_CATALOG_API         Operation = "create_data_catalog_api"
	//UPDATE_DATA_CATALOG_API         Operation = "update_data_catalog_api"
	//DELETE_DATA_CATALOG_API         Operation = "delete_data_catalog_api"
	CREATE_DICT_API       Operation = "create_dict_api"
	UPDATE_DICT_API       Operation = "update_dict_api"
	DELETE_DICT_API       Operation = "delete_dict_api"
	BATCH_DELETE_DICT_API Operation = "batch_delete_dict_api"
	//EXPORT_DICT_API                 Operation = "export_dict_api"
	//MOVE_CATALOG_DICT_API           Operation = "move_catalog_dict_api"
	//STATE_DICT_API                  Operation = "state_dict_api"
	CREATE_RULE_API       Operation = "create_rule_api"
	UPDATE_RULE_API       Operation = "update_rule_api"
	BATCH_DELETE_RULE_API Operation = "batch_delete_rule_api"
	//STATE_RULE_API                  Operation = "state_rule_api"
	//MOVE_CATALOG_RULE_API           Operation = "move_catalog_rule_api"
	//STD_CREATE_STAGING_API          Operation = "std_create_staging_api"
	//STD_CREATE_SUBMIT_API           Operation = "std_create_submit_api"
	//STD_DELETE_TASK_API             Operation = "std_delete_task_api"
	//STD_CREATE_TASK_API             Operation = "std_create_task_api"
	//STD_CANCEL_TASK_API             Operation = "std_cancel_task_api"
	//STD_SUBMIT_TASK_API             Operation = "std_submit_task_api"
	//STD_FINISH_TASK_API             Operation = "std_finish_task_api"
	//STD_UPDATE_DESCRIPTION_API      Operation = "std_update_description_api"
	//STD_ACCEPT_API                  Operation = "std_accept_api"
	//STD_UPDATE_TABLE_API            Operation = "std_update_table_name_api"
	//STD_CREATE_PINGING_API          Operation = "std_create_pinging_api"
	STD_CREATE_FILE_API       Operation = "std_create_file_api"
	STD_UPDATE_FILE_API       Operation = "std_update_file_api"
	BATCH_STD_DELETE_FILE_API Operation = "batch_std_delete_file_api"
	//STATE_STD_FILE_API              Operation = "state_std_file_api"
	//MOVE_STD_CATALOG_FILE_API       Operation = "move_std_catalog_file_api"
	//STD_DOWN_FILE_API               Operation = "std_down_file_api"
	//STD_BATCH_DOWN_FILE_API         Operation = "std_batch_down_file_api"
	//STD_DICT_RULE_RELATION_FILE_API Operation = "std_dict_rule_relation_file_api"
	// 数据标准相关======结束=====
)

// Operation 的简体中文名称
const (
	// 未知操作
	OperationSimplifiedChineseNameUnknown string = "未知操作"

	// 用户登录
	OperationSimplifiedChineseNameLogin string = "登录"
	// 用户登出
	OperationSimplifiedChineseNameLogout string = "登出"
	// 生成接口
	OperationSimplifiedChineseNameGenerateAPI string = "生成接口"
	// 注册接口
	OperationSimplifiedChineseNameRegisterAPI string = "注册接口"
	// 修改接口
	OperationSimplifiedChineseNameUpdateAPI string = "修改接口"
	// 设置接口权限
	OperationSimplifiedChineseNameSetAPIAuthorization string = "设置接口权限"
	// 发布接口
	OperationSimplifiedChineseNamePublicAPI string = "发布接口"
	// 上线接口
	OperationSimplifiedChineseNameUpAPI string = "上线接口"
	// 下线接口
	OperationSimplifiedChineseNameDownAPI string = "下线接口"
	// 删除接口
	OperationSimplifiedChineseNameDeleteAPI string = "删除接口"
	// 查看接口调用信息
	OperationSimplifiedChineseNameGetAPIAuthInfo string = "查看接口调用信息"
	// 申请接口权限
	OperationSimplifiedChineseNameRequestAPIAuthorization string = "申请接口权限"
	// 授权接口权限
	OperationSimplifiedChineseNameAuthorizeAPI string = "授权接口"
	// 创建数据下载任务
	OperationSimplifiedChineseNameCreateDataDownloadTask string = "新建逻辑视图下载任务"
	// 删除数据下载任务
	OperationSimplifiedChineseNameDeleteDataDownloadTask string = "删除逻辑视图下载任务"
	// 下载数据
	OperationSimplifiedChineseNameDataDownload string = "下载逻辑视图"
	// 申请逻辑视图权限
	OperationSimplifiedChineseNameRequestDataViewAuthorization string = "申请逻辑视图权限"
	// 授权逻辑视图权限
	OperationSimplifiedChineseNameAuthorizeDataView string = "授权逻辑视图"
	// 预览逻辑视图数据
	OperationSimplifiedCDataPreview string = "预览逻辑视图数据"
	// 创建维度模型
	OperationSimplifiedChineseNameCreateDimensionModel string = "新建维度模型"
	// 修改维度模型基本信息
	OperationSimplifiedChineseNameUpdateDimensionModelBasicInfo string = "修改维度模型基本信息"
	// 修改维度模型配置信息
	OperationSimplifiedChineseNameUpdateDimensionModelConfigInfo string = "修改维度模型配置信息"
	// 删除维度模型
	OperationSimplifiedChineseNameDeleteDimensionModel string = "删除维度模型"
	// 创建指标
	OperationSimplifiedChineseNameCreateIndicator string = "新建指标"
	// 编辑指标
	OperationSimplifiedChineseNameUpdateIndicator string = "修改指标"
	// 删除指标
	OperationSimplifiedChineseNameDeleteIndicator string = "删除指标"
	// 查询指标结果
	OperationSimplifiedChineseNameQueryIndicatorResult string = "查看指标数据"
	// 申请指标权限
	OperationSimplifiedChineseNameRequestIndicatorAuthorization string = "申请指标权限"
	// 授权指标权限
	OperationSimplifiedChineseNameAuthorizeIndicator string = "授权指标"

	// 创建业务域分组
	OperationSimplifiedChineseNameCreateBusinessDomainGroup string = "新建业务域分组"
	// 编辑业务域分组
	OperationSimplifiedChineseNameUpdateBusinessDomainGroup string = "修改业务域分组"
	// 删除业务域分组
	OperationSimplifiedChineseNameDeleteBusinessDomainGroup string = "删除业务域分组"
	// 创建业务域
	OperationSimplifiedChineseNameCreateBusinessDomain string = "新建业务域"
	// 编辑业务域
	OperationSimplifiedChineseNameUpdateBusinessDomain string = "修改业务域"
	// 删除业务域
	OperationSimplifiedChineseNameDeleteBusinessDomain string = "删除业务域"
	// 创建主干业务
	OperationSimplifiedChineseNameCreateBusinessProcess string = "新建主干业务"
	// 编辑主干业务
	OperationSimplifiedChineseNameUpdateBusinessProcess string = "修改主干业务"
	// 删除主干业务
	OperationSimplifiedChineseNameDeleteBusinessProcess string = "删除主干业务"
	// 创建业务模型
	OperationSimplifiedChineseNameCreateBusinessModel string = "新建业务模型"
	// 编辑业务模型
	OperationSimplifiedChineseNameUpdateBusinessModel string = "修改业务模型"
	// 删除业务模型
	OperationSimplifiedChineseNameDeleteBusinessModel string = "删除业务模型"
	// 创建数据模型
	OperationSimplifiedChineseNameCreateDataModel string = "新建数据模型"
	// 编辑数据模型
	OperationSimplifiedChineseNameUpdateDataModel string = "修改数据模型"
	// 删除数据模型
	OperationSimplifiedChineseNameDeleteDataModel string = "删除数据模型"
	// 创建业务表
	OperationSimplifiedChineseNameCreateBusinessForms string = "新建业务表"
	// 编辑业务表
	OperationSimplifiedChineseNameUpdateBusinessForms string = "修改业务表基本信息"
	// 删除业务表
	OperationSimplifiedChineseNameDeleteBusinessForms string = "删除业务表"
	// 修改业务表表结构
	OperationSimplifiedChineseNameUpdateBusinessFormsContent string = "修改业务表表结构"
	//创建数据表
	OperationSimplifiedChineseNameCreateDataForms string = "新建数据表"
	// 编辑数据表
	OperationSimplifiedChineseNameUpdateDataForms string = "修改数据表基本信息"
	// 删除数据表
	OperationSimplifiedChineseNameDeleteDataForms string = "删除数据表"
	// 修改数据表结构
	OperationSimplifiedChineseNameUpdateDataFormsContent string = "修改业务表表结构"
	// 创建业务流程图
	OperationSimplifiedChineseNameCreateBusinessFlowcharts string = "新建流程图"
	// 编辑业务流程图
	OperationSimplifiedChineseNameUpdateBusinessFlowcharts string = "修改流程图基本信息"
	// 删除业务流程图
	OperationSimplifiedChineseNameDeleteBusinessFlowcharts string = "删除流程图"
	// 修改业务流程图内容
	OperationSimplifiedChineseNameUpdateBusinessFlowchartsContent string = "修改流程图内容"
	// 导出业务流程图
	OperationSimplifiedChineseNameExportBusinessFlowcharts string = "导出流程图"
	// 创建业务指标
	OperationSimplifiedChineseNameCreateBusinessIndicator string = "新建业务指标"
	// 编辑业务指标
	OperationSimplifiedChineseNameUpdateBusinessIndicator string = "修改业务指标"
	// 删除业务指标
	OperationSimplifiedChineseNameDeleteBusinessIndicator string = "删除业务指标"
	//创建数据指标
	OperationSimplifiedChineseNameCreateDataIndicator string = "新建数据指标"
	// 编辑数据指标
	OperationSimplifiedChineseNameUpdateDataIndicator string = "修改数据指标"
	// 删除数据指标
	OperationSimplifiedChineseNameDeleteDataIndicator string = "删除数据指标"
	// 创建主题域分组
	OperationSimplifiedChineseNameCreateSubjectDomainGroup string = "新建主题域分组"
	// 编辑主题域分组
	OperationSimplifiedChineseNameUpdateSubjectDomainGroup string = "修改主题域分组"
	// 删除主题域分组
	OperationSimplifiedChineseNameDeleteSubjectDomainGroup string = "删除主题域分组"
	// 创建主题域
	OperationSimplifiedChineseNameCreateSubjectDomain string = "新建主题域"
	// 编辑主题域
	OperationSimplifiedChineseNameUpdateSubjectDomain string = "修改主题域"
	// 删除主题域
	OperationSimplifiedChineseNameDeleteSubjectDomain string = "删除主题域"
	// 创建业务对象
	OperationSimplifiedChineseNameCreateBusinessObject string = "新建业务对象"
	// 编辑业务对象基本信息
	OperationSimplifiedChineseNameUpdateBusinessObject string = "修改业务对象基本信息"
	// 删除业务对象
	OperationSimplifiedChineseNameDeleteBusinessObject string = "删除业务对象"
	// 编辑业务对象内容
	OperationSimplifiedChineseNameUpdateBusinessObjectContent string = "修改业务对象内容"
	// 创建业务活动
	OperationSimplifiedChineseNameCreateBusinessActivity string = "新建业务活动"
	// 编辑业务活动基本信息
	OperationSimplifiedChineseNameUpdateBusinessActivity string = "修改业务活动基本信息"
	// 删除业务活动
	OperationSimplifiedChineseNameDeleteBusinessActivity string = "删除业务活动"
	// 编辑业务活动内容
	OperationSimplifiedChineseNameUpdateBusinessActivityContent string = "修改业务活动内容"

	// 新建逻辑视图
	OperationSimplifiedChineseNameCreateLogicView string = "新建逻辑视图"
	// 修改逻辑视图基本信息
	OperationSimplifiedChineseNameUpdateLogicView string = "修改逻辑视图"
	// 删除逻辑视图
	OperationSimplifiedChineseNameDeleteLogicView string = "删除逻辑视图"
	// 扫描数据源
	OperationSimplifiedChineseNameScanDataSource string = "扫描数据源"
	// 上线逻辑视图
	OperationSimplifiedChineseNameOnlineLogicView string = "上线逻辑视图"
	// 下线逻辑视图
	OperationSimplifiedChineseNameOfflineLogicView string = "下线逻辑视图"
	// 发布逻辑视图
	OperationSimplifiedChineseNamePublishLogicView string = "发布逻辑视图"
	// 数据标准相关===开始=====
	CREATE_DATAELEMENT_NAME_API string = "创建数据元"
	//EXPORT_DATAELEMENT_NAME_API          string = "通过查询结果批量导出数据元"
	//EXPORT_IDS_DATAELEMENT_NAME_API      string = "通过ID集合批量导出数据元"
	UPDATE_DATAELEMENT_NAME_API       string = "修改数据元"
	BATCH_DELETE_DATAELEMENT_NAME_API string = "批量删除数据元"
	//MOVE_CATALOG_DATAELEMENT_NAME_API    string = "移动数据元目录"
	//DELETE_FJLABEL_DATAELEMENT_NAME_API  string = "删除数据元分级标签"
	//CREATE_DATA_CATALOG_NAME_API         string = "创建数据标准目录"
	//UPDATE_DATA_CATALOG_NAME_API         string = "修改数据标准目录"
	//DELETE_DATA_CATALOG_NAME_API string = "删除数据标准目录"
	CREATE_DICT_NAME_API       string = "新增码表"
	UPDATE_DICT_NAME_API       string = "修改码表"
	DELETE_DICT_NAME_API       string = "删除码表"
	BATCH_DELETE_DICT_NAME_API string = "批量删除码表"
	//EXPORT_DICT_NAME_API                 string = "导出码表"
	//MOVE_CATALOG_DICT_NAME_API           string = "移动码表目录"
	//STATE_DICT_NAME_API                  string = "码表停用和启用"
	CREATE_RULE_NAME_API       string = "新增编码规则"
	UPDATE_RULE_NAME_API       string = "修改编码规则"
	BATCH_DELETE_RULE_NAME_API string = "批量删除编码规则"
	//STATE_RULE_NAME_API                  string = "编码规则停用和启用"
	//MOVE_CATALOG_RULE_NAME_API           string = "移动编码规则目录"
	//STD_CREATE_STAGING_NAME_API          string = "创建标准任务-关联标准-暂存"
	//STD_CREATE_SUBMIT_NAME_API           string = "创建标准任务-关联标准-提交"
	//STD_DELETE_TASK_NAME_API             string = "删除待新建标准"
	//STD_CREATE_TASK_NAME_API             string = "待新建标准-新建标准任务"
	//STD_CANCEL_TASK_NAME_API             string = "待新建标准-取消标准任务"
	//STD_SUBMIT_TASK_NAME_API             string = "标准任务-提交选定的数据元"
	//STD_FINISH_TASK_NAME_API             string = "标准任务-完成任务"
	//STD_UPDATE_DESCRIPTION_NAME_API      string = "待新建标准-修改待新建标准字段说明"
	//STD_ACCEPT_NAME_API                  string = "待新建标准-采纳"
	//STD_UPDATE_TABLE_NAME_API            string = "待新建标准-修改业务标准表名称"
	//STD_CREATE_PINGING_NAME_API          string = "创建标准任务-添加至待新建标准接口"
	STD_CREATE_FILE_NAME_API       string = "新增标准文件"
	STD_UPDATE_FILE_NAME_API       string = "修改标准文件"
	BATCH_STD_DELETE_FILE_NAME_API string = "批量删除标准文件"
	//STATE_STD_FILE_NAME_API              string = "标准文件停用和启用"
	//MOVE_STD_CATALOG_FILE_NAME_API       string = "移动标准文件目录"
	//STD_DOWN_FILE_NAME_API               string = "下载标准文件附件"
	//STD_BATCH_DOWN_FILE_NAME_API         string = "批量下载标准文件附件"
	//STD_DICT_RULE_RELATION_FILE_NAME_API string = "添加标准文件和数据元&码表&编码规则的关联"
	// 数据标准相关=====结束=====

)

// operationSimplifiedChineseNameMap 定义 Operation 与中文简体名称的映射关系
var operationSimplifiedChineseNameMap = map[Operation]string{
	//用户登录
	OperationLogin: OperationSimplifiedChineseNameLogin,
	//用户登出
	OperationLogout: OperationSimplifiedChineseNameLogout,
	// 生成接口
	OperationGenerateAPI: OperationSimplifiedChineseNameGenerateAPI,
	// 注册接口
	OperationRegisterAPI: OperationSimplifiedChineseNameRegisterAPI,
	// 修改接口
	OperationUpdateAPI: OperationSimplifiedChineseNameUpdateAPI,
	// 设置接口权限
	OperationSetAPIAuthorization: OperationSimplifiedChineseNameSetAPIAuthorization,
	// 发布接口
	OperationPublicAPI: OperationSimplifiedChineseNamePublicAPI,
	// 上线接口
	OperationUpAPI: OperationSimplifiedChineseNameUpAPI,
	// 下线接口
	OperationDownAPI: OperationSimplifiedChineseNameDownAPI,
	// 删除接口
	OperationDeleteAPI: OperationSimplifiedChineseNameDeleteAPI,
	// 查看接口调用信息
	OperationGetAPIAuthInfo: OperationSimplifiedChineseNameGetAPIAuthInfo,
	// 申请接口权限
	OperationRequestAPIAuthorization: OperationSimplifiedChineseNameRequestAPIAuthorization,
	// 授权接口权限
	OperationAuthorizeAPI: OperationSimplifiedChineseNameAuthorizeAPI,
	// 创建数据下载任务
	OperationCreateDataDownloadTask: OperationSimplifiedChineseNameCreateDataDownloadTask,
	// 删除数据下载任务
	OperationDeleteDataDownloadTask: OperationSimplifiedChineseNameDeleteDataDownloadTask,
	// 下载数据
	OperationDataDownload: OperationSimplifiedChineseNameDataDownload,
	// 申请逻辑视图权限
	OperationRequestDataViewAuthorization: OperationSimplifiedChineseNameRequestDataViewAuthorization,
	// 授权逻辑视图权限
	OperationAuthorizeDataView: OperationSimplifiedChineseNameAuthorizeDataView,
	// 逻辑视图数据预览
	OperationDataPreview: OperationSimplifiedCDataPreview,
	// 创建维度模型
	OperationCreateDimensionModel: OperationSimplifiedChineseNameCreateDimensionModel,
	// 修改维度模型基本信息
	OperationUpdateDimensionModelBasicInfo: OperationSimplifiedChineseNameUpdateDimensionModelBasicInfo,
	// 修改维度模型配置信息
	OperationUpdateDimensionModelConfigInfo: OperationSimplifiedChineseNameUpdateDimensionModelConfigInfo,
	// 删除维度模型
	OperationDeleteDimensionModel: OperationSimplifiedChineseNameDeleteDimensionModel,
	// 创建指标
	OperationCreateIndicator: OperationSimplifiedChineseNameCreateIndicator,
	// 编辑指标
	OperationUpdateIndicator: OperationSimplifiedChineseNameUpdateIndicator,
	// 删除指标
	OperationDeleteIndicator: OperationSimplifiedChineseNameDeleteIndicator,
	// 查看指标数据
	OperationQueryIndicatorResult: OperationSimplifiedChineseNameQueryIndicatorResult,
	// 申请指标权限
	OperationRequestIndicatorAuthorization: OperationSimplifiedChineseNameRequestIndicatorAuthorization,
	// 授权指标权限
	OperationAuthorizeIndicator: OperationSimplifiedChineseNameAuthorizeIndicator,

	// 创建业务域分组
	OperationCreateBusinessDomainGroup: OperationSimplifiedChineseNameCreateBusinessDomainGroup,
	// 编辑业务域分组
	OperationUpdateBusinessDomainGroup: OperationSimplifiedChineseNameUpdateBusinessDomainGroup,
	// 删除业务域分组
	OperationDeleteBusinessDomainGroup: OperationSimplifiedChineseNameDeleteBusinessDomainGroup,
	// 创建业务域
	OperationCreateBusinessDomain: OperationSimplifiedChineseNameCreateBusinessDomain,
	// 编辑业务域
	OperationUpdateBusinessDomain: OperationSimplifiedChineseNameUpdateBusinessDomain,
	// 删除业务域
	OperationDeleteBusinessDomain: OperationSimplifiedChineseNameDeleteBusinessDomain,
	// 创建主干业务
	OperationCreateBusinessProcess: OperationSimplifiedChineseNameCreateBusinessProcess,
	// 编辑主干业务
	OperationUpdateBusinessProcess: OperationSimplifiedChineseNameUpdateBusinessProcess,
	// 删除主干业务
	OperationDeleteBusinessProcess: OperationSimplifiedChineseNameDeleteBusinessProcess,
	// 创建业务模型
	OperationCreateBusinessModel: OperationSimplifiedChineseNameCreateBusinessModel,
	// 编辑业务模型
	OperationUpdateBusinessModel: OperationSimplifiedChineseNameUpdateBusinessModel,
	// 删除业务模型
	OperationDeleteBusinessModel: OperationSimplifiedChineseNameDeleteBusinessModel,

	// 创建数据模型
	OperationCreateDataModel: OperationSimplifiedChineseNameCreateDataModel,
	// 编辑数据模型
	OperationUpdateDataModel: OperationSimplifiedChineseNameUpdateDataModel,
	// 删除数据模型
	OperationDeleteDataModel: OperationSimplifiedChineseNameDeleteDataModel,

	// 创建业务表
	OperationCreateBusinessForms: OperationSimplifiedChineseNameCreateBusinessForms,
	// 编辑业务表
	OperationUpdateBusinessForms: OperationSimplifiedChineseNameUpdateBusinessForms,
	// 删除业务表
	OperationDeleteBusinessForms: OperationSimplifiedChineseNameDeleteBusinessForms,

	// 修改业务表表结构
	OperationUpdateBusinessFormsContent: OperationSimplifiedChineseNameUpdateBusinessFormsContent,
	// 创建数据表
	OperationCreateDataForms: OperationSimplifiedChineseNameCreateDataForms,
	// 编辑数据表
	OperationUpdateDataForms: OperationSimplifiedChineseNameUpdateDataForms,
	// 删除数据表
	OperationDeleteDataForms: OperationSimplifiedChineseNameDeleteDataForms,
	//修改数据表结构
	OperationUpdateDataFormsContent: OperationSimplifiedChineseNameUpdateDataFormsContent,
	// 创建业务流程图
	OperationCreateBusinessFlowcharts: OperationSimplifiedChineseNameCreateBusinessFlowcharts,
	// 编辑业务流程图
	OperationUpdateBusinessFlowcharts: OperationSimplifiedChineseNameUpdateBusinessFlowcharts,
	// 删除业务流程图
	OperationDeleteBusinessFlowcharts: OperationSimplifiedChineseNameDeleteBusinessFlowcharts,
	// 修改业务流程图内容
	OperationUpdateBusinessFlowchartsContent: OperationSimplifiedChineseNameUpdateBusinessFlowchartsContent,
	// 导出业务流程图
	OperationExportBusinessFlowcharts: OperationSimplifiedChineseNameExportBusinessFlowcharts,
	// 创建业务指标
	OperationCreateBusinessIndicator: OperationSimplifiedChineseNameCreateBusinessIndicator,
	// 编辑业务指标
	OperationUpdateBusinessIndicator: OperationSimplifiedChineseNameUpdateBusinessIndicator,
	// 删除业务指标
	OperationDeleteBusinessIndicator: OperationSimplifiedChineseNameDeleteBusinessIndicator,
	//创建数据指标
	OperationCreateDataIndicator: OperationSimplifiedChineseNameCreateDataIndicator,
	// 编辑数据指标
	OperationUpdateDataIndicator: OperationSimplifiedChineseNameUpdateDataIndicator,
	// 删除数据指标
	OperationDeleteDataIndicator: OperationSimplifiedChineseNameDeleteDataIndicator,
	// 创建主题域分组
	OperationCreateSubjectDomainGroup: OperationSimplifiedChineseNameCreateSubjectDomainGroup,
	// 编辑主题域分组
	OperationUpdateSubjectDomainGroup: OperationSimplifiedChineseNameUpdateSubjectDomainGroup,
	// 删除主题域分组
	OperationDeleteSubjectDomainGroup: OperationSimplifiedChineseNameDeleteSubjectDomainGroup,
	// 创建主题域
	OperationCreateSubjectDomain: OperationSimplifiedChineseNameCreateSubjectDomain,
	// 编辑主题域
	OperationUpdateSubjectDomain: OperationSimplifiedChineseNameUpdateSubjectDomain,
	// 删除主题域
	OperationDeleteSubjectDomain: OperationSimplifiedChineseNameDeleteSubjectDomain,
	// 创建业务对象
	OperationCreateBusinessObject: OperationSimplifiedChineseNameCreateBusinessObject,
	// 编辑业务对象基本信息
	OperationUpdateBusinessObject: OperationSimplifiedChineseNameUpdateBusinessObject,
	// 删除业务对象
	OperationDeleteBusinessObject: OperationSimplifiedChineseNameDeleteBusinessObject,
	// 编辑业务对象内容
	OperationUpdateBusinessObjectContent: OperationSimplifiedChineseNameUpdateBusinessObjectContent,
	// 创建业务活动
	OperationCreateBusinessActivity: OperationSimplifiedChineseNameCreateBusinessActivity,
	// 编辑业务活动基本信息
	OperationUpdateBusinessActivity: OperationSimplifiedChineseNameUpdateBusinessActivity,
	// 删除业务活动
	OperationDeleteBusinessActivity: OperationSimplifiedChineseNameDeleteBusinessActivity,
	// 编辑业务活动内容
	OperationUpdateBusinessActivityContent: OperationSimplifiedChineseNameUpdateBusinessActivityContent,

	// 新建逻辑视图
	OperationCreateLogicView: OperationSimplifiedChineseNameCreateLogicView,
	// 修改逻辑视图
	OperationUpdateLogicView: OperationSimplifiedChineseNameUpdateLogicView,
	// 删除逻辑视图
	OperationDeleteLogicView: OperationSimplifiedChineseNameDeleteLogicView,
	// 扫描数据源
	OperationScanDataSource: OperationSimplifiedChineseNameScanDataSource,
	// 上线逻辑视图
	OperationOnlineLogicView: OperationSimplifiedChineseNameOnlineLogicView,
	// 下线逻辑视图
	OperationOfflineLogicView: OperationSimplifiedChineseNameOfflineLogicView,
	// 发布逻辑视图
	OperationPublishLogicView: OperationSimplifiedChineseNamePublishLogicView,
	// 数据标准相关===开始====
	CREATE_DATAELEMENT_API: CREATE_DATAELEMENT_NAME_API,
	//EXPORT_DATAELEMENT_API:          EXPORT_DATAELEMENT_NAME_API,
	//EXPORT_IDS_DATAELEMENT_API:      EXPORT_IDS_DATAELEMENT_NAME_API,
	UPDATE_DATAELEMENT_API:       UPDATE_DATAELEMENT_NAME_API,
	BATCH_DELETE_DATAELEMENT_API: BATCH_DELETE_DATAELEMENT_NAME_API,
	//MOVE_CATALOG_DATAELEMENT_API:    MOVE_CATALOG_DATAELEMENT_NAME_API,
	//DELETE_FJLABEL_DATAELEMENT_API:  DELETE_FJLABEL_DATAELEMENT_NAME_API,
	//CREATE_DATA_CATALOG_API:         CREATE_DATA_CATALOG_NAME_API,
	//UPDATE_DATA_CATALOG_API:         UPDATE_DATA_CATALOG_NAME_API,
	//DELETE_DATA_CATALOG_API:         DELETE_DATA_CATALOG_NAME_API,
	CREATE_DICT_API:       CREATE_DICT_NAME_API,
	UPDATE_DICT_API:       UPDATE_DICT_NAME_API,
	DELETE_DICT_API:       DELETE_DICT_NAME_API,
	BATCH_DELETE_DICT_API: BATCH_DELETE_DICT_NAME_API,
	//EXPORT_DICT_API:                 EXPORT_DICT_NAME_API,
	//MOVE_CATALOG_DICT_API:           MOVE_CATALOG_DICT_NAME_API,
	//STATE_DICT_API:                  STATE_DICT_NAME_API,
	CREATE_RULE_API:       CREATE_RULE_NAME_API,
	UPDATE_RULE_API:       UPDATE_RULE_NAME_API,
	BATCH_DELETE_RULE_API: BATCH_DELETE_RULE_NAME_API,
	//STATE_RULE_API:                  STATE_RULE_NAME_API,
	//MOVE_CATALOG_RULE_API:           MOVE_CATALOG_RULE_NAME_API,
	//STD_CREATE_STAGING_API:          STD_CREATE_STAGING_NAME_API,
	//STD_CREATE_SUBMIT_API:           STD_CREATE_SUBMIT_NAME_API,
	//STD_DELETE_TASK_API:             STD_DELETE_TASK_NAME_API,
	//STD_CREATE_TASK_API:             STD_CREATE_TASK_NAME_API,
	//STD_CANCEL_TASK_API:             STD_CANCEL_TASK_NAME_API,
	//STD_SUBMIT_TASK_API:             STD_SUBMIT_TASK_NAME_API,
	//STD_FINISH_TASK_API:             STD_FINISH_TASK_NAME_API,
	//STD_UPDATE_DESCRIPTION_API:      STD_UPDATE_DESCRIPTION_NAME_API,
	//STD_ACCEPT_API:                  STD_ACCEPT_NAME_API,
	//STD_UPDATE_TABLE_API:            STD_UPDATE_TABLE_NAME_API,
	//STD_CREATE_PINGING_API:          STD_CREATE_PINGING_NAME_API,
	STD_CREATE_FILE_API:       STD_CREATE_FILE_NAME_API,
	STD_UPDATE_FILE_API:       STD_UPDATE_FILE_NAME_API,
	BATCH_STD_DELETE_FILE_API: BATCH_STD_DELETE_FILE_NAME_API,
	//STATE_STD_FILE_API:              STATE_STD_FILE_NAME_API,
	//MOVE_STD_CATALOG_FILE_API:       MOVE_STD_CATALOG_FILE_NAME_API,
	//STD_DOWN_FILE_API:               STD_DOWN_FILE_NAME_API,
	//STD_BATCH_DOWN_FILE_API:         STD_BATCH_DOWN_FILE_NAME_API,
	//STD_DICT_RULE_RELATION_FILE_API: STD_DICT_RULE_RELATION_FILE_NAME_API,
	// 数据标准相关====结束====
}

func (o Operation) SimplifiedChineseName() string {
	n, ok := operationSimplifiedChineseNameMap[o]
	if !ok {
		return OperationSimplifiedChineseNameUnknown
	}
	return n
}

// AuditType 代表审计日志的类型
type AuditType string

const (
	// 管理日志
	AuditTypeManagement AuditType = "Management"
	// 操作日志
	AuditTypeOperation AuditType = "Operation"
	// 登录日志
	AuditTypeLogin AuditType = "Login"
)

// 定义审计日志类型 AuditType 与操作类型 Operation 之间的映射关系
type operationAndAuditTypeBinding struct {
	operation Operation
	auditType AuditType
}

// 定义审计日志类型 AuditType 与操作类型 Operation 之间的映射关系
var operationAndAuditTypeBindings = []operationAndAuditTypeBinding{
	// 用户登录
	{operation: OperationLogin, auditType: AuditTypeLogin},
	// 用户登出
	{operation: OperationLogout, auditType: AuditTypeLogin},
	// 生成接口
	{operation: OperationGenerateAPI, auditType: AuditTypeManagement},
	// 注册接口
	{operation: OperationRegisterAPI, auditType: AuditTypeManagement},
	// 修改接口
	{operation: OperationUpdateAPI, auditType: AuditTypeManagement},
	// 设置接口权限
	{operation: OperationSetAPIAuthorization, auditType: AuditTypeOperation},
	// 发布接口
	{operation: OperationPublicAPI, auditType: AuditTypeManagement},
	// 上线接口
	{operation: OperationUpAPI, auditType: AuditTypeManagement},
	// 下线接口
	{operation: OperationDownAPI, auditType: AuditTypeManagement},
	// 删除接口
	{operation: OperationDeleteAPI, auditType: AuditTypeManagement},
	// 查看接口调用信息
	{operation: OperationGetAPIAuthInfo, auditType: AuditTypeOperation},
	// 申请接口权限
	{operation: OperationRequestAPIAuthorization, auditType: AuditTypeOperation},
	// 授权接口权限
	{operation: OperationAuthorizeAPI, auditType: AuditTypeOperation},

	// 创建数据下载任务
	{operation: OperationCreateDataDownloadTask, auditType: AuditTypeOperation},
	// 删除数据下载任务
	{operation: OperationDeleteDataDownloadTask, auditType: AuditTypeOperation},
	// 下载数据
	{operation: OperationDataDownload, auditType: AuditTypeOperation},
	// 下载数据
	{operation: OperationDataPreview, auditType: AuditTypeOperation},
	// 申请逻辑视图权限
	{operation: OperationRequestDataViewAuthorization, auditType: AuditTypeOperation},
	// 授权逻辑视图权限
	{operation: OperationAuthorizeDataView, auditType: AuditTypeOperation},

	// 创建维度模型
	{operation: OperationCreateDimensionModel, auditType: AuditTypeManagement},
	// 编辑维度模型基本信息
	{operation: OperationUpdateDimensionModelBasicInfo, auditType: AuditTypeManagement},
	// 编辑维度模型配置信息
	{operation: OperationUpdateDimensionModelConfigInfo, auditType: AuditTypeManagement},
	// 删除维度模型
	{operation: OperationDeleteDimensionModel, auditType: AuditTypeManagement},

	// 创建指标
	{operation: OperationCreateIndicator, auditType: AuditTypeManagement},
	// 编辑指标
	{operation: OperationUpdateIndicator, auditType: AuditTypeManagement},
	// 删除指标
	{operation: OperationDeleteIndicator, auditType: AuditTypeManagement},
	// 查询指标结果
	{operation: OperationQueryIndicatorResult, auditType: AuditTypeOperation},
	// 申请指标权限
	{operation: OperationRequestIndicatorAuthorization, auditType: AuditTypeOperation},
	// 授权指标权限
	{operation: OperationAuthorizeIndicator, auditType: AuditTypeOperation},

	// 创建业务域分组
	{operation: OperationCreateBusinessDomainGroup, auditType: AuditTypeManagement},
	// 编辑业务域分组
	{operation: OperationUpdateBusinessDomainGroup, auditType: AuditTypeManagement},
	// 删除业务域分组
	{operation: OperationDeleteBusinessDomainGroup, auditType: AuditTypeManagement},
	// 创建业务域
	{operation: OperationCreateBusinessDomain, auditType: AuditTypeManagement},
	// 编辑业务域
	{operation: OperationUpdateBusinessDomain, auditType: AuditTypeManagement},
	// 删除业务域
	{operation: OperationDeleteBusinessDomain, auditType: AuditTypeManagement},
	// 创建主干业务
	{operation: OperationCreateBusinessProcess, auditType: AuditTypeManagement},
	// 编辑主干业务
	{operation: OperationUpdateBusinessProcess, auditType: AuditTypeManagement},
	// 删除主干业务
	{operation: OperationDeleteBusinessProcess, auditType: AuditTypeManagement},
	// 创建业务模型
	{operation: OperationCreateBusinessModel, auditType: AuditTypeManagement},
	// 编辑业务模型
	{operation: OperationUpdateBusinessModel, auditType: AuditTypeManagement},
	// 删除业务模型
	{operation: OperationDeleteBusinessModel, auditType: AuditTypeManagement},
	// 创建数据模型
	{operation: OperationCreateDataModel, auditType: AuditTypeManagement},
	// 编辑数据模型
	{operation: OperationUpdateDataModel, auditType: AuditTypeManagement},
	// 删除数据模型
	{operation: OperationDeleteDataModel, auditType: AuditTypeManagement},
	// 创建业务表
	{operation: OperationCreateBusinessForms, auditType: AuditTypeManagement},
	// 编辑业务表
	{operation: OperationUpdateBusinessForms, auditType: AuditTypeManagement},
	// 删除业务表
	{operation: OperationDeleteBusinessForms, auditType: AuditTypeManagement},
	// 修改业务表表结构
	{operation: OperationUpdateBusinessFormsContent, auditType: AuditTypeManagement},
	// 创建数据表
	{operation: OperationCreateDataForms, auditType: AuditTypeManagement},
	// 编辑数据表
	{operation: OperationUpdateDataForms, auditType: AuditTypeManagement},
	// 删除数据表
	{operation: OperationDeleteDataForms, auditType: AuditTypeManagement},
	// 修改数据表表结构
	{operation: OperationUpdateBusinessFormsContent, auditType: AuditTypeManagement},
	// 创建业务流程图
	{operation: OperationCreateBusinessFlowcharts, auditType: AuditTypeManagement},
	// 编辑业务流程图
	{operation: OperationUpdateBusinessFlowcharts, auditType: AuditTypeManagement},
	// 删除业务流程图
	{operation: OperationDeleteBusinessFlowcharts, auditType: AuditTypeManagement},
	// 修改业务流程图内容
	{operation: OperationUpdateBusinessFlowchartsContent, auditType: AuditTypeManagement},
	//导出业务流程图
	{operation: OperationExportBusinessFlowcharts, auditType: AuditTypeOperation},
	// 创建业务指标
	{operation: OperationCreateBusinessIndicator, auditType: AuditTypeManagement},
	// 编辑业务指标
	{operation: OperationUpdateBusinessIndicator, auditType: AuditTypeManagement},
	// 删除业务指标
	{operation: OperationDeleteBusinessIndicator, auditType: AuditTypeManagement},
	// 创建数据指标
	{operation: OperationCreateDataIndicator, auditType: AuditTypeManagement},
	// 编辑数据指标
	{operation: OperationUpdateDataIndicator, auditType: AuditTypeManagement},
	// 删除数据指标
	{operation: OperationDeleteDataIndicator, auditType: AuditTypeManagement},

	// 创建主题域分组
	{operation: OperationCreateSubjectDomainGroup, auditType: AuditTypeManagement},
	// 编辑主题域分组
	{operation: OperationUpdateSubjectDomainGroup, auditType: AuditTypeManagement},
	// 删除主题域分组
	{operation: OperationDeleteSubjectDomainGroup, auditType: AuditTypeManagement},
	// 创建主题域
	{operation: OperationCreateSubjectDomain, auditType: AuditTypeManagement},
	// 编辑主题域
	{operation: OperationUpdateSubjectDomain, auditType: AuditTypeManagement},
	// 删除主题域
	{operation: OperationDeleteSubjectDomain, auditType: AuditTypeManagement},
	// 创建业务对象
	{operation: OperationCreateBusinessObject, auditType: AuditTypeManagement},
	// 编辑业务对象基本信息
	{operation: OperationUpdateBusinessObject, auditType: AuditTypeManagement},
	// 删除业务对象
	{operation: OperationDeleteBusinessObject, auditType: AuditTypeManagement},
	// 编辑业务对象内容
	{operation: OperationUpdateBusinessObjectContent, auditType: AuditTypeManagement},
	// 创建业务活动
	{operation: OperationCreateBusinessActivity, auditType: AuditTypeManagement},
	// 编辑业务活动基本信息
	{operation: OperationUpdateBusinessActivity, auditType: AuditTypeManagement},
	// 删除业务活动
	{operation: OperationDeleteBusinessActivity, auditType: AuditTypeManagement},
	// 编辑业务活动内容
	{operation: OperationUpdateBusinessActivityContent, auditType: AuditTypeManagement},
	// 新建逻辑视图
	{operation: OperationCreateLogicView, auditType: AuditTypeManagement},
	// 修改逻辑视图
	{operation: OperationUpdateLogicView, auditType: AuditTypeManagement},
	// 删除逻辑视图
	{operation: OperationDeleteLogicView, auditType: AuditTypeManagement},
	// 扫描数据源
	{operation: OperationScanDataSource, auditType: AuditTypeManagement},
	// 上线逻辑视图
	{operation: OperationOnlineLogicView, auditType: AuditTypeManagement},
	// 下线逻辑视图
	{operation: OperationOfflineLogicView, auditType: AuditTypeManagement},
	// 发布逻辑视图
	{operation: OperationPublishLogicView, auditType: AuditTypeManagement},
	// 数据标准相关===开始====
	{operation: CREATE_DATAELEMENT_API, auditType: AuditTypeManagement},
	//{operation: EXPORT_DATAELEMENT_API, auditType: AuditTypeOperation},
	//{operation: EXPORT_IDS_DATAELEMENT_API, auditType: AuditTypeOperation},
	{operation: UPDATE_DATAELEMENT_API, auditType: AuditTypeManagement},
	{operation: BATCH_DELETE_DATAELEMENT_API, auditType: AuditTypeManagement},
	//{operation: MOVE_CATALOG_DATAELEMENT_API, auditType: AuditTypeManagement},
	//{operation: DELETE_FJLABEL_DATAELEMENT_API, auditType: AuditTypeManagement},
	//{operation: CREATE_DATA_CATALOG_API, auditType: AuditTypeManagement},
	//{operation: UPDATE_DATA_CATALOG_API, auditType: AuditTypeManagement},
	//{operation: DELETE_DATA_CATALOG_API, auditType: AuditTypeManagement},
	{operation: CREATE_DICT_API, auditType: AuditTypeManagement},
	{operation: UPDATE_DICT_API, auditType: AuditTypeManagement},
	{operation: DELETE_DICT_API, auditType: AuditTypeManagement},
	{operation: BATCH_DELETE_DICT_API, auditType: AuditTypeManagement},
	//{operation: EXPORT_DICT_API, auditType: AuditTypeOperation},
	//{operation: MOVE_CATALOG_DICT_API, auditType: AuditTypeManagement},
	//{operation: STATE_DICT_API, auditType: AuditTypeManagement},
	{operation: CREATE_RULE_API, auditType: AuditTypeManagement},
	{operation: UPDATE_RULE_API, auditType: AuditTypeManagement},
	{operation: BATCH_DELETE_RULE_API, auditType: AuditTypeManagement},
	//{operation: STATE_RULE_API, auditType: AuditTypeManagement},
	//{operation: MOVE_CATALOG_RULE_API, auditType: AuditTypeManagement},
	//{operation: STD_CREATE_STAGING_API, auditType: AuditTypeManagement},
	//{operation: STD_CREATE_SUBMIT_API, auditType: AuditTypeManagement},
	//{operation: STD_DELETE_TASK_API, auditType: AuditTypeManagement},
	//{operation: STD_CREATE_TASK_API, auditType: AuditTypeManagement},
	//{operation: STD_SUBMIT_TASK_API, auditType: AuditTypeManagement},
	//{operation: STD_CANCEL_TASK_API, auditType: AuditTypeManagement},
	//{operation: STD_FINISH_TASK_API, auditType: AuditTypeManagement},
	//{operation: STD_UPDATE_DESCRIPTION_API, auditType: AuditTypeManagement},
	//{operation: STD_ACCEPT_API, auditType: AuditTypeManagement},
	//{operation: STD_UPDATE_TABLE_API, auditType: AuditTypeManagement},
	//{operation: STD_CREATE_PINGING_API, auditType: AuditTypeManagement},
	{operation: STD_CREATE_FILE_API, auditType: AuditTypeManagement},
	{operation: STD_UPDATE_FILE_API, auditType: AuditTypeManagement},
	{operation: BATCH_STD_DELETE_FILE_API, auditType: AuditTypeManagement},
	//{operation: STATE_STD_FILE_API, auditType: AuditTypeManagement},
	//{operation: MOVE_STD_CATALOG_FILE_API, auditType: AuditTypeManagement},
	//{operation: STD_DOWN_FILE_API, auditType: AuditTypeOperation},
	//{operation: STD_BATCH_DOWN_FILE_API, auditType: AuditTypeOperation},
	//{operation: STD_DICT_RULE_RELATION_FILE_API, auditType: AuditTypeManagement},
	// 数据标准相关======结束=====
}

// 过滤属于指定审计日志类型 AuditType 的操作 Operation 列表
func FilterOperationsByAuditType(in []Operation, t AuditType) (out []Operation) {
	return sets.List(sets.New(in...).Intersection(sets.New(OperationsForAuditType(t)...)))
}

// 返回属于指定审计日志类型 AuditType 的操作 Operation 列表
func OperationsForAuditType(t AuditType) (operations []Operation) {
	for _, b := range operationAndAuditTypeBindings {
		if b.auditType != t {
			continue
		}
		operations = append(operations, b.operation)
	}
	return
}
