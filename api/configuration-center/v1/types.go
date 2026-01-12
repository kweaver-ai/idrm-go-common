package v1

import (
	audit "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
	doc_audit_rest_v1 "github.com/kweaver-ai/idrm-go-common/api/doc_audit_rest/v1"
	meta "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// AuditEventQuery 定义对 AuditEvent 的查询。
type AuditEventQuery struct {
	// 过滤条件
	Filter *QueryFilter `json:"filter,omitempty"`
	// 排序条件
	Sort QuerySort `json:"sort,omitempty"`
	// 排序方向
	Direction SortDirection `json:"direction,omitempty"`
	// 偏移量，默认为 0。
	Offset int `json:"offset,omitempty"`
	// 期望返回的数量上限，默认为 10。
	Limit int `json:"limit,omitempty"`
}

// QueryFilter 定义过滤审计日志的条件，同时指定多个条件时，过滤同时满足多个条件
// 的审计日志。
type QueryFilter struct {
	// 非空时，过滤指定类型的审计日志
	Type audit.AuditType `json:"type,omitempty"`
	// 非空时，过滤指定级别的审计日志。
	//
	// 指定多个时，过滤任意指定级别的审计日志。
	Levels []audit.Level `json:"levels,omitempty"`
	// 非空时，过滤指定操作类型的审计日志指定多个时。
	//
	// 过滤任意指定操作类型的审计日志。
	Operations []audit.Operation `json:"operations,omitempty"`
	// 非空时，过滤在指定时间段内生成的审计日志
	TimeRange TimeRange `json:"time_range,omitempty"`
	// 非空时，过滤由用户名包含指定字符串的用户产生的审计日志。指定多个用户时，
	//
	// 过滤由用户名包含任意指定字符串的用户产生的审计日志。
	OperatorNames []string `json:"operator_names,omitempty"`
	// 非空时，过滤由用户所属部门完整路径包含指定字符串的用户产生的审计日志。指
	//
	// 定多个用户时，过滤由用户所属部门完整路径包含任意指定字符串的用户产生的审
	// 计日志。
	OperatorDepartments []string `json:"operator_departments,omitempty"`
	// 非空时，过滤用户代理 IP 包含指定字符串的审计日志
	//
	// 指定多个字符串时，过滤用户代理 IP 包含任意指定字符串的审计日志
	OperatorAgentIPs []string `json:"operator_agent_ips,omitempty"`
	// 非空时，过滤描述包含指定字符串的审计日志
	//
	// 指定多个字符串时，过滤描述包含任意指定字符串的审计日志
	Descriptions []string `json:"descriptions,omitempty"`
	// 非空时，过滤详情包含指定字符串的审计日志
	//
	// 指定多个字符串时，过滤详情包含任意指定字符串的审计日志
	Details []string `json:"details,omitempty"`
}

// TimeRange 定义一段时间，闭区间
type TimeRange struct {
	// Unix 时间戳，单位：毫秒。
	//
	// 非空时，不早于指定时间
	Start *meta.TimestampUnixMilli `json:"start,omitempty"`
	// Unix 时间戳，单位：毫秒。
	//
	// 非空时，不晚于指定时间
	End *meta.TimestampUnixMilli `json:"end,omitempty"`
}

// 排序条件
type QuerySort string

const (
	// 根据审计时间发生的时间排序
	QuerySortTimestamp QuerySort = "timestamp"
)

// 排序方向
type SortDirection string

const (
	// 升序
	Ascending SortDirection = "asc"
	// 降序
	Descending SortDirection = "desc"
)

// AuditEventList 定义审计日志列表
type AuditEventList struct {
	// 审计日志列表
	Entries []AuditEvent `json:"entries,omitempty"`
	// 满足过滤条件的审计日志的总数
	TotalCount int `json:"total_count,omitempty"`
}

// AuditEvent 定义一个审计事件，对应一条审计日志。
type AuditEvent struct {
	// 审计日志发生的时间
	Timestamp meta.Time `json:"timestamp,omitempty"`
	// 日志级别
	Level audit.Level `json:"level,omitempty"`
	// 操作者的显示名称
	OperatorName string `json:"operator_name,omitempty"`
	// 操作者所属的部门的完整路径列表
	OperatorDepartments []string `json:"operator_departments,omitempty"`
	// 操作者代理的 IP
	OperatorAgentIP string `json:"operator_agent_ip,omitempty"`
	// 操作类型
	Operation audit.Operation `json:"operation,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
	// 详情，具体定义，每种资源不同
	//  - 接口：DetailAPI
	Detail any `json:"detail,omitempty"`
}

// DetailAPI 定义审计日志：接口详情
type DetailAPI struct {
	// 接口 ID
	ID string `json:"id,omitempty"`
	// 接口名称
	Name string `json:"name,omitempty"`
	// 接口 Owner 的显示名称
	OwnerName string `json:"owner_name,omitempty"`
	// 接口所属主题域的名称路径
	Subject string `json:"subject,omitempty"`
	// 接口所属部门的名称路径
	Department string `json:"department,omitempty"`
}

// 前置机
type FrontEndProcessor struct {
	// 元数据
	FrontEndProcessorMetadata
	// 申请
	Request FrontEndProcessorRequest `json:"request,omitempty"`
	// 节点
	Node *FrontEndProcessorNode `json:"node,omitempty"`
	// 状态
	Status FrontEndProcessorStatus `json:"status,omitempty"`
}

// 前置机
type FrontEndProcessorDetail struct {

	// 元数据
	FrontEndProcessorMetadataResponse
	// 申请
	Request FrontEndProcessorRequest `json:"request,omitempty"`
	// 部署信息
	Deployment FrontEndProcessorDeployment `json:"deployment,omitempty"`
	// 前置机信息
	Processor []FrontEndProcessorInfo `json:"processor,omitempty"`
	// 申报意见
	Comment string `json:"comment,omitempty"`
	// 状态
	Status AggregatedFrontEndProcessorStatus `json:"status,omitempty"`
}

type FrontEndProcessorMetadataResponse struct {
	ID string `json:"id,omitempty"`
	// 订单 ID
	OrderID string `json:"order_id,omitempty"`
	// 姓名
	Name string `json:"name,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty"`
	// 手机
	Mobile string `json:"mobile,omitempty"`
	// 邮件
	Mail string `json:"mail,omitempty"`
	//技术联系人
	AdministratorName string `json:"administrator_name,omitempty"`
	// 技术联系人手机
	AdministratorPhone string `json:"administrator_phone,omitempty"`
	// 技术联系人邮箱
	AdministratorMail string `json:"administrator_mail,omitempty"`
	// 技术联系人传真
	AdministratorFax string `json:"administrator_fax,omitempty"`
}

// 前置机元数据
type FrontEndProcessorMetadata struct {
	ID string `json:"id,omitempty"`
	// 订单 ID
	OrderID string `json:"order_id,omitempty"`
	// 创建者 ID
	CreatorID string `json:"creator_id,omitempty"`
	// 更新者 ID
	UpdaterID string `json:"updater_id,omitempty"`
	// 申请者 ID
	RequesterID string `json:"requester_id,omitempty"`
	// 签收者 ID
	RecipientID string `json:"recipient_id,omitempty"`
	// 创建时间
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
	// 更新时间
	UpdateTimestamp string `json:"update_timestamp,omitempty"`
	// 申请时间
	RequestTimestamp string `json:"request_timestamp,omitempty"`
	// 分配时间
	AllocationTimestamp string `json:"allocation_timestamp,omitempty"`
	// 签收时间
	ReceiptTimestamp string `json:"receipt_timestamp,omitempty"`
	// 回收时间
	ReclaimTimestamp string `json:"reclaim_timestamp,omitempty"`
}

// 前置机申请
type FrontEndProcessorRequest struct {
	// 所属部门
	Department FrontEndProcessorDepartment `json:"department,omitempty"`
	// 联系人
	Contact FrontEndProcessorContact `json:"contact,omitempty"`
	// 部署信息
	Deployment FrontEndProcessorDeployment `json:"deployment,omitempty"`
	// 前置机信息
	Processor []FrontEndProcessorInfo `json:"processor,omitempty"`
	// 申报意见
	Comment string `json:"comment,omitempty"`
	// 是否为草稿
	IsDraft bool `json:"is_draft,omitempty"`
	// 申请类型
	ApplyType string `json:"apply_type,omitempty"`
	// 驳回理由
	RejectReason string `json:"reject_reason,omitempty"`
}

// 前置机所属部门
type FrontEndProcessorDepartment struct {
	// 部门 ID
	ID string `json:"id,omitempty"`
	// 单位地址
	Address string `json:"address,omitempty"`
}

// 前置机联系人
type FrontEndProcessorContact struct {
	// 姓名
	Name string `json:"name,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty"`
	// 手机
	Mobile string `json:"mobile,omitempty"`
	// 邮件
	Mail string `json:"mail,omitempty"`
	//技术联系人
	AdministratorName string `json:"administrator_name,omitempty"`
	// 技术联系人手机
	AdministratorPhone string `json:"administrator_phone,omitempty"`
	// 技术联系人邮箱
	AdministratorMail string `json:"administrator_mail,omitempty"`
	// 技术联系人传真
	AdministratorFax string `json:"administrator_fax,omitempty"`
}

// 部署信息
type FrontEndProcessorDeployment struct {
	// 部署区域
	DeployArea string `json:"deployment_area,omitempty"`
	// 运行业务系统
	RunBusinessSystem string `json:"deployment_system,omitempty"`
	// 业务系统保护级别
	BusinessSystemLevel string `json:"protection_level,omitempty"`
}

// 前置机信息
type FrontEndProcessorInfo struct {
	// 前置机id
	ID string `json:"front_end_item_id,omitempty"`
	//关联前置机表front_end_id
	FrontEndID string `json:"front_end_id,omitempty"`
	//操作系统类型/版本
	OS string `json:"os,omitempty"`
	// 技术资源规格
	Spec string `json:"spec,omitempty"`
	// IP
	IP string `json:"node_ip,omitempty"`
	// 端口
	Port int32 `json:"node_port,omitempty"`
	// 节点名称
	Name string `json:"node_name,omitempty"`
	// 技术负责人名称
	AdministratorName string `json:"administrator_name,omitempty"`
	//技术负责人电话
	AdministratorPhone string `json:"administrator_phone,omitempty"`
	//业务磁盘空间
	BusinessDiskSpace string `json:"business_disk_space,omitempty"`
	//前置库数量
	LibraryCount string `json:"library_number,omitempty"`
	//更新日期
	UpdateTime string `json:"update_time,omitempty"`
	//前置库列表，可以添加多条，前置库类型与说明字段
	LibraryList []FrontEndLibrary `json:"library_list,omitempty"`
}

// FrontEndProcessorAllocationRequest 前置机分配请求
type FrontEndProcessorAllocationRequest struct {
	// 前置机分配信息
	Allocation []FrontEndProcessorAllocation `json:"allocation,omitempty"`
}

// FrontEndProcessorAllocation 前置机分配信息
type FrontEndProcessorAllocation struct {
	// 分配者 ID
	ID string `json:"front_end_item_id,omitempty"`
	//关联前置机表front_end_id
	FrontEndID string `json:"front_end_id,omitempty"`
	// 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// node信息
	//Node *FrontEndProcessorNode `json:"node,omitempty"`
	IP string `gorm:"column:node_ip" json:"ip,omitempty"`
	// 端口
	Port int32 `gorm:"column:node_port" json:"port,omitempty"`
	// 节点名称
	Name string `gorm:"column:node_name" json:"name,omitempty"`
	//技术负责人
	AdministratorName string `json:"administrator_name,omitempty"`
	//技术负责人手机
	AdministratorPhone string `json:"administrator_phone,omitempty"`
	//前置库列表
	LibraryList []FrontEndAllocationLibrary `json:"library_list,omitempty" gorm:"-"`
	// 前置机状态
	Status string `json:"status,omitempty"`
}

type FrontEndProcessorItemList struct {
	Items []*FrontEndProcessorItem `json:"items"`
	Total int                      `json:"total"`
}

type FrontEndProcessorItem struct {
	// 分配者 ID
	ID string `json:"front_end_item_id,omitempty"`
	//关联前置机表front_end_id
	FrontEndID string `json:"front_end_id,omitempty"`
	// 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// node信息
	IP string `gorm:"column:node_ip" json:"ip"`
	// 端口
	Port int32 `gorm:"column:node_port" json:"port"`
	// 节点名称
	Name string `gorm:"column:node_name" json:"name"`
	//技术负责人
	AdministratorName string `json:"administrator_name"`
	//技术负责人手机
	AdministratorPhone string `json:"administrator_phone"`
	// 前置机状态
	Status string `json:"status" `
	// 前置库数量
	LibraryCount string `gorm:"column:library_number" json:"library_number"`
	// 前置库类型
	LibraryType string `gorm:"column:type" json:"library_type"`
	// 前置库时间
	BusinessTime string `gorm:"column:update_time" json:"business_time"`
	// 所属单位
	Address string `gorm:"column:department_address" json:"department_address"`
	// 部门编号
	DepartmentID string `json:"department_id"`
	// 创建人
	CreatorId string `json:"creator_id"`
	// 所属单位
	DepartmentName string `gorm:"-" json:"department_name"`
	// 操作系统类型
	OS string `gorm:"column:operator_system" json:"os"`
	// 技术资源规格
	Spec string `gorm:"column:computer_resource" json:"spec"`
	// 业务磁盘空间
	BusinessDiskSpace string `gorm:"column:disk_space" json:"business_disk_space"`
	// 分配时间
	AllocationTimestamp string `gorm:"column:allocation_timestamp"  json:"allocation_timestamp"`
	// 签收时间
	ReceiptTimestamp string `gorm:"column:receipt_timestamp"  json:"receipt_timestamp"`
	// 回收时间
	ReclaimTimestamp string `gorm:"column:reclaim_timestamp"  json:"reclaim_timestamp"`
}

type FrontEndProcessorItemListOptions struct {
	//meta.ListOptions
	// 节点 IP，非空时模糊匹配
	NodeIP  string `json:"node_ip"  form:"node_ip"  binding:"omitempty"`
	Keyword string `json:"keyword"  form:"keyword"  binding:"omitempty"`
	// 状态
	Status string `json:"status"  form:"status"  binding:"omitempty"`
	// 节点名称
	NodeName string `json:"node_name"  form:"node_name"  binding:"omitempty"`
	// 前置库数量
	LibraryCount string `json:"library_number"  form:"library_number"  binding:"omitempty"`
	// 技术负责人
	AdministratorName string `json:"administrator_name"  form:"administrator_name"  binding:"omitempty"`
	//技术负责人电话
	AdministratorPhone string `json:"administrator_phone"  form:"administrator_phone"  binding:"omitempty"`
	// 部门编号
	DepartmentID string `json:"department_ids"  form:"department_ids"  binding:"omitempty"`
	// 单位名称
	DepartmentName string `json:"department_address"  form:"department_address"  binding:"omitempty"`
	// limit
	Limit int `json:"limit,omitempty"`
	// offset
	Offset int `json:"offset,omitempty"`
	// 排序字段
	Sort string `json:"sort,omitempty"`
	// 排序方向
	Direction SortDirection `json:"direction,omitempty"`
}

func (FrontEndProcessorAllocation) TableName() string {
	return "front_end_item"
}

type FrontEndProcessorLibrary struct {
	//表id
	ID string `json:"id,omitempty"`
	//前置库id
	FrontEndID string `json:"front_end_id,omitempty"`
	//前置库类型
	LibraryType string `json:"library_type,omitempty"`
	//前置库说明
	LibraryExplain string `json:"library_explain,omitempty"`
	//更新时间
	UpdateTime string `json:"update_time,omitempty"`
}

type FrontEndAllocationLibrary struct {
	//表id
	ID string `json:"library_id,omitempty"`
	//前置库id
	FrontEndID string `json:"front_end_id,omitempty"`
	//前置库类型
	LibraryType string `gorm:"column:type" json:"library_type,omitempty"`
	//前置库版本
	LibraryVersion string `gorm:"column:version" json:"library_version,omitempty"`
	// 前置库名称
	Name string `json:"name,omitempty"`
	//前置库用户名
	Username string `json:"username,omitempty"`
	//前置库密码
	Password string `json:"password,omitempty"`
	//前置库对接业务
	BusinessName string `json:"business_name,omitempty"`
	//前置库时间
	UpdateTime string `json:"business_time,omitempty"`
	//更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	//创建日期
	CreatedAt string `json:"created_at,omitempty"`
	//前置库关联前置机表front_end_item_id
	FrontEndItemID string `json:"front_end_item_id,omitempty"`
}

func (FrontEndAllocationLibrary) TableName() string {
	return "front_end_library"
}

type FrontEndAllocation struct {
	//表id
	ID string `json:"front_end_item_id,omitempty"`
	//前置库id
	FrontEndID string `json:"front_end_id,omitempty"`
	//前置库类型
	LibraryType string `json:"library_type,omitempty"`
	//前置库版本
	LibraryVersion string `json:"library_version,omitempty"`
	//前置库用户名
	Username string `json:"username,omitempty"`
	//前置库密码
	Password string `json:"password,omitempty"`
	//前置库对接业务
	BusinessName string `json:"business_name,omitempty"`
	//前置库时间
	UpdateTime string `json:"update_time,omitempty"`
	//更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	//创建日期
	CreatedAt string `json:"created_at,omitempty"`
}

func (FrontEndAllocation) TableName() string {
	return "front_end_library"
}

// 前置机节点信息
type FrontEndProcessorNode struct {
	IP string `json:"ip,omitempty"`
	// 端口
	Port int `json:"port,omitempty"`
	// 节点名称
	Name string `json:"name,omitempty"`
	// 技术负责人
	Administrator FrontEndProcessorNodeAdministrator `json:"administrator,omitempty"`
}

type FrontEnd struct {
	ID                 string            `json:"front_end_item_id" gorm:"primaryKey"`
	FrontEndID         string            `json:"front_end_id"`
	OperatorSystem     string            `json:"operator_system"`
	ComputerResource   string            `json:"computer_resource"`
	DiskSpace          string            `json:"disk_space"`
	IP                 string            `json:"node_ip" gorm:"column:node_ip"`
	Port               int32             `json:"node_port" gorm:"column:node_port"`
	Node               string            `json:"node_name,omitempty" gorm:"column:node_name"`
	AdministratorName  string            `json:"administrator_name,omitempty" gorm:"column:administrator_name"`
	AdministratorPhone string            `json:"administrator_phone,omitempty" gorm:"column:administrator_phone"`
	LibraryNumber      string            `json:"library_number"`
	LibraryList        []FrontEndLibrary `json:"library_list"`
	UpdatedAt          string            `json:"updated_at,omitempty"`
	CreatedAt          string            `json:"created_at,omitempty"`
	Status             string            `json:"status,omitempty"`
}

// TableName 方法指定表名
func (FrontEnd) TableName() string {
	return "front_end_item"
}

type FrontEndLibrary struct {
	ID             string `json:"library_id" gorm:"column:id;primary_key"`
	FrontEndID     string `json:"front_end_id" gorm:"column:front_end_id"`
	Type           string `json:"library_type" gorm:"column:type"`
	Name           string `json:"name" gorm:"column:name"`
	Username       string `json:"username" gorm:"column:username"`
	Password       string `json:"password" gorm:"column:password"`
	BusinessName   string `json:"business_name" gorm:"column:business_name"`
	Comment        string `json:"comment" gorm:"column:comment"`
	UpdatedAt      string `json:"updated_at" json:"updated_at,omitempty"`
	CreatedAt      string `json:"created_at" json:"created_at,omitempty"`
	FrontEndItemID string `json:"front_end_item_id" gorm:"column:front_end_item_id"`
	Version        string `json:"library_version" gorm:"column:version"`
	BusinessTime   string `json:"business_time" gorm:"column:update_time"`
}

// tableName 方法指定表名
func (FrontEndLibrary) TableName() string {
	return "front_end_library"
}

// 前置机节点的技术负责人
type FrontEndProcessorNodeAdministrator struct {
	// 姓名
	Name string `json:"name,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty"`
	// 邮箱
	Mail string `json:"mail,omitempty"`
	// 传真
	Fax string `json:"fax,omitempty"`
}

// 前置机状态
type FrontEndProcessorStatus struct {
	// 所处生命周期的阶段
	Phase FrontEndProcessorPhase
	// workflow 审核的 ApplyID
	ApplyID string
	// 驳回理由
	RejectReason string
}

// 前置机在生命周期中所处的阶段
type FrontEndProcessorPhase string

// 前置机在生命周期中所处的阶段
const (
	// 待处理，用户创建了前置机申请，但未上报。
	FrontEndProcessorPending FrontEndProcessorPhase = "Pending"
	// 审核中，用户的前置机申请正在被审核。
	FrontEndProcessorAuditing FrontEndProcessorPhase = "Auditing"
	// 分配中，用户的前置机申请已经被批准，正在分配前置机。
	FrontEndProcessorAllocating FrontEndProcessorPhase = "Allocating"
	// 已分配，已经分配前置机，等待用户签收。
	FrontEndProcessorAllocated FrontEndProcessorPhase = "Allocated"
	//分配驳回，用户 FrontEndProcessor 申请被驳回。
	FrontEndProcessorRejected FrontEndProcessorPhase = "Rejected"
	// 使用中，用户已经签收前置机，前置机在使用中。
	FrontEndProcessorInCompleted FrontEndProcessorPhase = "InCompleted"
	// 已回收，前置机已经被回收。
	FrontEndProcessorReclaimed FrontEndProcessorPhase = "Reclaimed"
)

// 聚合的前置机，与 FrontEndProcessor 相比包括所引用资源的属性
type AggregatedFrontEndProcessor struct {
	// 元数据
	AggregatedFrontEndProcessorMetadata
	// 申请
	Request AggregatedFrontEndProcessorRequest `json:"request,omitempty"`
	// 节点
	Node *FrontEndProcessorNode `json:"node,omitempty"`
	// 状态
	Status AggregatedFrontEndProcessorStatus `json:"status,omitempty"`
	// 部署信息
	Deployment FrontEndProcessorDeployment `json:"deployment,omitempty"`
	// 前置机信息
	Info []FrontEndProcessorInfo `json:"info,omitempty"`
}

// 聚合的前置机元数据，与 FrontEndProcessorMetadata 相比包括所引用用户的姓名
type AggregatedFrontEndProcessorMetadata struct {
	FrontEndProcessorMetadata

	// 创建者姓名
	CreatorName string `json:"creator_name,omitempty"`
	// 更新者姓名
	UpdaterName string `json:"updater_name,omitempty"`
	// 申请者姓名
	RequesterName string `json:"requester_name,omitempty"`
	// 签收者姓名
	RecipientName string `json:"recipient_name,omitempty"`
}

// 聚合的前置机申请，与 FrontEndProcessorRequest 相比包括所属部门的属性
type AggregatedFrontEndProcessorRequest struct {
	// 所属部门
	Department AggregatedFrontEndProcessorDepartment `json:"department,omitempty"`
	// 联系人
	Contact FrontEndProcessorContact `json:"contact,omitempty"`
	// 部署信息
	Deployment FrontEndProcessorDeployment `json:"deployment,omitempty"`
	// 前置机信息
	Info []FrontEndProcessorInfo `json:"processor,omitempty"`
	// 申报意见
	Comment string `json:"comment,omitempty"`
	// 是否为草稿
	IsDraft bool `json:"is_draft,omitempty"`
	// 申请类型
	ApplyType string `json:"apply_type,omitempty"`
}

// 聚合的前置机所属部门，与 FrontEndProcessorDepartment 相比包括部门及其上级部门的路径
type AggregatedFrontEndProcessorDepartment struct {
	FrontEndProcessorDepartment

	// 部门及其上级部门的路径
	Path string `json:"path,omitempty"`
}

// 聚合的前置机状态，与 FrontEndProcessorStatus 相比包括其他资源的属性
type AggregatedFrontEndProcessorStatus struct {
	// 所处生命周期的阶段
	Phase FrontEndProcessorPhase `json:"phase,omitempty"`
	// workflow 审核状态
	Audit *AggregatedWorkflowAudit `json:"audit,omitempty"`
	// 驳回理由
	RejectReason string `json:"reject_reason,omitempty"`
}

// 聚合的 workflow 审核状态，包括结果、意见
type AggregatedWorkflowAudit struct {
	// doc_audit_rest_v1.Apply.ID
	ID string `json:"id,omitempty"`
	// doc_audit_rest_v1.Apply.BizID
	BizID string `json:"biz_id,omitempty"`
	// doc_audit_rest_v1.Apply.AuditStatus
	AuditStatus doc_audit_rest_v1.AuditStatus `json:"audit_status,omitempty"`
	// doc_audit_rest_v1.Apply.AuditMsg
	AuditMsg string `json:"audit_msg,omitempty"`
	// doc_audit_rest_v1.Apply.AuditStatus
	TaskID string `json:"task_id,omitempty"`
}

// 前置机列表
type FrontEndProcessorList meta.List[FrontEndProcessor]

// 聚合的前置机列表
type AggregatedFrontEndProcessorList meta.List[AggregatedFrontEndProcessor]

// 获取前置机列表的选项
type FrontEndProcessorListOptions struct {
	meta.ListOptions
	// 单号，非空时模糊匹配
	OrderID string `json:"order_id,omitempty"`
	// 节点 IP，非空时模糊匹配
	NodeIP string `json:"node_ip,omitempty"`
	// Phase 列表，非空时过滤指定 phase 的前置机
	Phases []FrontEndProcessorPhase `json:"phases,omitempty"`
	// 部门 ID 列表，非空时过滤属于指定部门的前置机
	DepartmentIDs []string `json:"department_i_ds,omitempty"`
	// 申请时间，起始值，非空时返回不早于此时申请的前置机
	RequestTimestampStart meta.Time `json:"request_timestamp_start,omitempty"`
	// 申请时间，终止值，非空时返回不晚于此时申请的前置机
	RequestTimestampEnd meta.Time `json:"request_timestamp_end,omitempty"`
	//申请类型
	ApplyType string `json:"apply_type,omitempty"`
}

// 概览
type FrontEndProcessorsOverview struct {
	// 已分配的数量
	AllocatedCount int `json:"allocated_count"`
	// 使用中的数量
	InUseCount int `json:"in_use_count"`
	// 已回收的数量
	ReclaimedCount int `json:"reclaimed_count"`
	// 总数量
	TotalCount int `json:"total_count"`

	// 近一年，每个月使用中的前置机新增数量
	LastYearInUse []int `json:"last_year_in_use"`
	// 近一年，每个月新增回收的前置机数量
	LastYearReclaimed []int `json:"last_year_reclaimed"`

	// 部门的前置机数量 TOP 15
	DepartmentsTOP15 []DepartmentNameFrontEndProcessorCount `json:"departments_top_15"`
}

// 部门及属其前置机数量
type DepartmentNameFrontEndProcessorCount struct {
	// 部门名称
	Name string `json:"name,omitempty"`
	// 前置机数量
	Count int `json:"count,omitempty"`
}

type FrontEndProcessorsOverviewGetOptions struct {
	// 开始时间
	Start meta.Time `json:"start,omitempty" form:"start"`
	// 结束时间
	End meta.Time `json:"end,omitempty" form:"end"`
}

// 签收驳回
type FrontEndProcessorReject struct {
	// 驳回意见
	Comment string `json:"comment,omitempty"`
}

// AuditList
type AuditListGetReq struct {
	Target    string `form:"target" binding:"required,oneof=tasks historys"`                                      // 审核列表类型 tasks 待审核 historys 已审核
	Offset    int    `json:"offset" form:"offset,default=1" binding:"min=1"`                                      // 页码, 默认1
	Limit     int    `json:"limit" form:"limit,default=10" binding:"min=1,max=2000"`                              // 每页大小, 默认1
	Direction string `json:"direction" form:"direction,default=desc" binding:"oneof=asc desc" default:"desc"`     // 排序方向，枚举：asc：正序；desc：倒序。默认倒序
	Sort      string `json:"sort" form:"sort,default=apply_time" binding:"oneof=apply_time" default:"apply_time"` // 排序类型，枚举：apply_at：按申请时间排序
	Keyword   string `json:"keyword" form:"keyword"`
}

type FrontEndData struct {
	Id               string `json:"id"`
	OrderId          string `json:"order_id"`
	AppleType        string `json:"apply_type"`
	RequestTimestamp string `json:"request_timestamp"`
}

type AuditListResp struct {
	PageResults[AuditListItem]
}

type PageResults[T any] struct {
	Entries    []*T  `json:"entries" binding:"required"`                       // 对象列表
	TotalCount int64 `json:"total_count" binding:"required,gte=0" example:"3"` // 当前筛选条件下的对象数量
}

type AuditListItem struct {
	ID          string `json:"id"`                              // 主键，uuid
	OrderId     string `json:"order_id"`                        // 单号
	ApplyType   string `json:"apply_type"`                      //申请类型
	Applyer     string `json:"applyer"`                         // 申请人
	ApplyTime   string `json:"apply_time"`                      // 申请时间
	AuditStatus string `json:"audit_status" example:"auditing"` // 审核状态
	AuditTime   string `json:"audit_time"`                      // 审核时间
	ProcInstID  string `json:"proc_inst_id"`                    // 审核实例ID
	TaskID      string `json:"task_id"`                         // 审核任务ID

}
