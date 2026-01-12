package v1

// Service 定义接口服务
//
// Service 接口服务详情，字段基本都在了
type Service struct {
	ServiceInfo     ServiceInfo      `json:"service_info"`               // 基本信息
	ServiceParam    ServiceParamRead `json:"service_param"`              // 参数配置
	ServiceResponse *ServiceResponse `json:"service_response,omitempty"` // 返回结果
	ServiceTest     ServiceTest      `json:"service_test"`               // 接口测试

}

// ServiceInfo 定义接口服务的基本信息
type ServiceInfo struct {
	ServiceID         string     `json:"service_id,omitempty"`
	ServiceCode       string     `json:"service_code,omitempty"`           // 编码
	ServiceName       string     `json:"service_name,omitempty"`           // 接口名称
	Department        Department `json:"department"`                       // 所属部门
	PublishStatus     string     `json:"publish_status"`                   // 发布状态
	PublishTime       string     `json:"publish_time"`                     // 发布时间
	OnlineStatus      string     `json:"status"`                           // 上线状态
	ApplyNum          uint64     `json:"apply_num,omitempty"`              // 申请数
	PreviewNum        uint64     `json:"preview_num,omitempty"`            // 预览数
	Status            string     `json:"status" form:"status"`             // 状态 draft 草稿 publish 已发布
	AuditType         string     `json:"audit_type" form:"audit_type"`     // 审核类型 unpublished 未发布 af-data-application-publish 发布审核
	AuditStatus       string     `json:"audit_status" form:"audit_status"` // 审核状态 unpublished 未发布 auditing 审核中 pass 通过 reject 驳回
	AuditAdvice       string     `json:"audit_advice"`                     // 审核意见，仅驳回时有用
	OnlineAuditAdvice string     `json:"online_audit_advice"`
	SubjectDomainId   string     `json:"subject_domain_id" form:"subject_domain_id"`     // 主题域id
	SubjectDomainName string     `json:"subject_domain_name" form:"subject_domain_name"` // 主题域名称
	ServiceType       string     `json:"service_type" form:"service_type"`               // 接口类型 service_generate 接口生成 service_register 接口注册
	// 同步标识(success、fail)
	SyncFlag string `json:"sync_flag" form:"sync_flag" binding:"omitempty" example:"success"`
	// 同步消息
	SyncMsg string `json:"sync_msg" form:"sync_msg" binding:"omitempty" example:"同步消息"`
	// 更新标识(success、fail)
	UpdateFlag string `json:"update_flag" form:"update_flag" binding:"omitempty" example:"success"`
	// 更新消息
	UpdateMsg string `json:"update_msg" form:"update_msg" binding:"omitempty" example:"更新消息"`
	// 授权id
	PaasID string `json:"paas_id" form:"paas_id" binding:"omitempty" example:"授权id"`
	// 网关路径前缀
	PrePath string `json:"pre_path" form:"pre_path" binding:"omitempty" example:"网关路径前缀"`

	// Deprecated: use Owners.OwnerID
	OwnerId string `json:"owner_id"`
	// Deprecated: use Owners.OwnerName``
	OwnerName          string                        `json:"owner_name"` // 数据owner用户名
	Owners             []DataApplicationServiceOwner `json:"owners,omitempty"`
	GatewayUrl         string                        `json:"gateway_url"`          // 网关地址
	ServicePath        string                        `json:"service_path"`         // 接口路径
	BackendServiceHost string                        `json:"backend_service_host"` // 后台服务域名/IP
	BackendServicePath string                        `json:"backend_service_path"` // 后台服务路径
	HTTPMethod         string                        `json:"http_method"`          // 请求方式 post get
	ReturnType         string                        `json:"return_type"`          // 返回类型 json
	Protocol           string                        `json:"protocol"`             // 协议 http
	File               File                          `json:"file"`                 // 接口文档
	Description        string                        `json:"description"`          // 接口说明
	RateLimiting       int64                         `json:"rate_limiting"`        // 调用频次 次/秒
	Timeout            int64                         `json:"timeout"`              // 超时时间
	OnlineTime         string                        `json:"online_time,omitempty"`
	ChangedServiceId   string                        `json:"changed_service_id" form:"changed_service_id"` //发起变更的接口的service_id
	IsChanged          string                        `json:"is_changed" form:"is_changed"`                 //已变更1，未变更0，默认0,用于标记service是否已变更
	CreateTime         string                        `json:"create_time,omitempty"`                        // 创建时间
	UpdateTime         string                        `json:"update_time,omitempty"`                        // 更新时间
	CreatedBy          string                        `json:"created_by,omitempty"`
	UpdateBy           string                        `json:"update_by,omitempty"`
}

type DataApplicationServiceOwner struct {
	OwnerID   string `json:"owner_id"`
	OwnerName string `json:"owner_name"`
}

type File struct {
	FileID   string `json:"file_id"`   // 文件id
	FileName string `json:"file_name"` // 文件名称
}

type Department struct {
	Id   string `json:"id"`   // 部门ID
	Name string `json:"name"` // 部门名称
}

type ServiceParamRead struct {
	CreateModel             string                   `json:"create_model,omitempty"`     // 创建模式 wizard 向导模式 script 脚本模式
	DatasourceId            string                   `json:"datasource_id"`              // 数据源id
	DatasourceName          string                   `json:"datasource_name"`            // 数据源名称
	DataViewId              string                   `json:"data_view_id"`               // 数据视图Id
	DataViewName            string                   `json:"data_view_name"`             // 数据视图名称
	Script                  string                   `json:"script,omitempty"`           // 脚本（仅脚本模式需要此参数）
	DataTableRequestParams  []DataTableRequestParam  `json:"data_table_request_params"`  // 请求参数
	DataTableResponseParams []DataTableResponseParam `json:"data_table_response_params"` // 返回参数
}

type DataTableRequestParam struct {
	CNName       string `json:"cn_name"`       // 中文名称
	EnName       string `json:"en_name"`       // 英文名称
	DataType     string `json:"data_type"`     // 字段类型
	Required     string `json:"required"`      // 是否必填 yes 必填 no非必填
	Operator     string `json:"operator"`      // 运算逻辑 = 等于, != 不等于, > 大于, >= 大于等于, < 小于, <= 小于等于, like 模糊匹配, in 包含, not in 不包含
	DefaultValue string `json:"default_value"` // 默认值
	Description  string `json:"description"`   // 描述
}

type DataTableResponseParam struct {
	CNName      string `json:"cn_name"`     // 中文名称
	EnName      string `json:"en_name"`     // 英文名称
	DataType    string `json:"data_type"`   // 字段类型
	Description string `json:"description"` // 描述
	Sort        string `json:"sort"`        // 排序方式 unsorted 不排序 asc 升序 desc 降序 默认 unsorted
	Masking     string `json:"masking"`     // 脱敏规则 plaintext 不脱敏 hash 哈希 override 覆盖 replace 替换 默认 plaintext
	Sequence    int64  `json:"sequence"`    // 序号
}

type ServiceResponse struct {
	Rules    []Rule `json:"rules"`     // 过滤规则
	Page     string `json:"page"`      // 是否分页 yes 是 no 否
	PageSize int64  `json:"page_size"` // 分页大小
}

type Rule struct {
	Param    string `json:"param"`    // 过滤字段
	Operator string `json:"operator"` // 运算逻辑 = 等于, != 不等于, > 大于, >= 大于等于, < 小于, <= 小于等于, like 模糊匹配, in 包含, not in 不包含
	Value    string `json:"value"`    // 过滤值
}

type ServiceTest struct {
	RequestExample  string `json:"request_example"`  // 请求示例
	ResponseExample string `json:"response_example"` // 返回示例
}
