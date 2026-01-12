package v1

import (
	"time"

	meta "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// TODO: 可能代表审核申请，需要与 audit-doc, authority, process-instance 等对象区分
type Apply struct {
	// ID
	ID string `json:"id,omitempty"`
	// TODO: biz 是什么？可能是申请（还是审核？）类型
	BizID string `json:"biz_id,omitempty"`
	// TODO: biz 是什么？可能是申请（还是审核？）类型
	BizType BizType `json:"biz_type,omitempty"`
	// TODO：添加描述信息
	ProcInstID string `json:"proc_inst_id,omitempty"`
	// TODO: 确认定义，可能代表 Apply 的创建时间
	ApplyTime time.Time `json:"apply_time,omitempty"`
	// 状态
	AuditStatus AuditStatus `json:"audit_status,omitempty"`
	// 审核消息
	AuditMsg string `json:"audit_msg,omitempty"`
	// TODO: 补充注释 TaskID
	TaskID string `json:"task_id,omitempty"`
}

// TODO: 明确定义，biz 是什么，有什么用
type BizType string

const (
	// 应用案例上报
	AF_SSZD_ApplicationExampleReport BizType = "af-sszd-application-example-report"
	// 应用案例下架
	AF_SSZD_ApplicationExampleWithdraw BizType = "af-sszd-application-example-withdraw"
)

// 审核状态（？）
type AuditStatus string

const (
	// 未审核
	Pending AuditStatus = "pending"
	// 审核中
	Auditing AuditStatus = "auditing"
	// 已拒绝
	Reject AuditStatus = "reject"
	// 已通过
	Pass AuditStatus = "pass"
	// 已撤销
	Undone AuditStatus = "undone"
)

// Apply 列表
type ApplyList meta.List[Apply]

// 获取 Apply 列表的参数
type ApplyListOptions struct {
	meta.ListOptions
	// 非空时返回属于指定 Type
	Type []BizType `json:"type,omitempty"`
	// 非空时返回指定状态
	Status AuditStatus `json:"status,omitempty"`
}

// 代表 GET /api/doc-audit-rest/v1/doc-audit/tasks 返回数据的 .apply_detail
//
// 可能代表申请或审核的详情
//
// TODO: 明确资源定义
type ApplyDetail struct {
	Process ApplyDetailProcess `json:"process,omitempty"`
}

// 代表 GET /api/doc-audit-rest/v1/doc-audit/tasks 返回数据的 .apply_detail.process
//
// TODO: 明确资源定义
type ApplyDetailProcess struct {
	ApplyID string `json:"apply_id,omitempty"`
}

// 代表 GET /api/doc-audit-rest/v1/doc-audit/tasks 返回的数据
//
// TODO: 与 Apply 相似，需要明确资源定义
type Task struct {
	// ID
	ID string `json:"id,omitempty"`
	// BizType
	BizType BizType `json:"biz_type,omitempty"`
	// Apply Detail
	ApplyDetail ApplyDetail `json:"apply_detail,omitempty"`
	// TODO：添加描述信息
	ProcInstID string `json:"proc_inst_id,omitempty"`
	// TODO: 确认定义，可能代表 Apply 的创建时间
	ApplyTime time.Time `json:"apply_time,omitempty"`
	// 状态
	AuditStatus AuditStatus `json:"audit_status,omitempty"`
}

// 代表 GET /api/doc-audit-rest/v1/doc-audit/tasks 返回的数据
//
// TODO: 与 Apply 相似，需要明确资源定义
type TaskList meta.List[Task]

// 获取 Task 列表的参数
type TaskListOptions struct {
	meta.ListOptions
	// 非空时返回属于指定 Type
	Type []BizType `json:"type,omitempty"`
	// 非空时返回指定状态
	Status AuditStatus `json:"status,omitempty"`
}

// 代表 GET /api/doc-audit-rest/v1/doc-audit/historys 返回的数据
//
// TODO: 与 Apply 相似，需要明确资源定义
type History struct {
	ID string `json:"id,omitempty"`
	// BizType
	BizType BizType `json:"biz_type,omitempty"`
	// Apply Detail
	ApplyDetail ApplyDetail `json:"apply_detail,omitempty"`
	// TODO：添加描述信息
	ProcInstID string `json:"proc_inst_id,omitempty"`
	// TODO: 确认定义，可能代表 Apply 的创建时间
	ApplyTime time.Time `json:"apply_time,omitempty"`
	// 状态
	AuditStatus AuditStatus `json:"audit_status,omitempty"`
}

// 代表 GET /api/doc-audit-rest/v1/doc-audit/historys 返回的数据
//
// TODO: 与 Apply 相似，需要明确资源定义
type HistoryList meta.List[History]

// 获取 History 列表的参数
type HistoryListOptions struct {
	meta.ListOptions
	// 非空时返回属于指定 Type
	Type []BizType `json:"type,omitempty"`
	// 非空时返回指定状态
	Status AuditStatus `json:"status,omitempty"`
}
