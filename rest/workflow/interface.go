package workflow

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kweaver-ai/idrm-go-common/workflow/common"
)

type WorkflowDriven interface {
	GetAuditProcessDefinition(ctx context.Context, procDefKey string) (res *ProcessDefinitionGetRes, err error)
	GetAuditList(ctx context.Context, listType WorkflowListType, auditTypes []string, offset, limit int) (*AuditResponse, error)
	GetAuditLogsByProcInstID(ctx context.Context, procInstID string) ([]*AuditNodeLog, error)
}

type ProcessDefinitionGetRes struct {
	Id             string      `json:"id"`
	Key            string      `json:"key"`
	Name           string      `json:"name"`
	Type           string      `json:"type"`
	TypeName       string      `json:"type_name"`
	CreateTime     time.Time   `json:"create_time"`
	CreateUserName string      `json:"create_user_name"`
	TenantId       string      `json:"tenant_id"`
	Description    interface{} `json:"description"`
	Effectivity    int         `json:"effectivity"` // 0 有效  1 无效
}

type WorkflowListType string

const (
	WORKFLOW_LIST_TYPE_APPLY   WorkflowListType = "applys"   // 我的申请
	WORKFLOW_LIST_TYPE_TASK    WorkflowListType = "tasks"    // 我的待办
	WORKFLOW_LIST_TYPE_HISTORY WorkflowListType = "historys" // 我处理的
)

// AuditResponse 表示审核响应的顶层结构
type AuditResponse struct {
	Entries    []*AuditEntry `json:"entries"`     // 审核条目列表
	TotalCount int64         `json:"total_count"` // 总条目数
}

// AuditEntry 表示单个审核条目的详细信息
type AuditEntry struct {
	ID          string      `json:"id"`           // 流程实例ID
	BizType     string      `json:"biz_type"`     // 业务类型
	DocID       *string     `json:"doc_id"`       // 文档ID，可为空
	DocPath     *string     `json:"doc_path"`     // 文档路径，可为空
	DocType     *string     `json:"doc_type"`     // 文档类型，可为空
	DocLibType  *string     `json:"doc_lib_type"` // 文档库类型，可为空
	ProcInstID  string      `json:"proc_inst_id"` // 审核任务ID
	Auditors    []*Auditor  `json:"auditors"`     // 审核人列表
	ApplyTime   string      `json:"apply_time"`   // 申请时间
	AuditStatus string      `json:"audit_status"` // 审核状态
	DocNames    string      `json:"doc_names"`    // 文档名称
	ApplyDetail ApplyDetail `json:"apply_detail"` // 申请详情
	Workflow    Workflow    `json:"workflow"`     // 工作流信息
	Version     *string     `json:"version"`      // 版本，可为空
}

// Auditor 表示审核人信息
type Auditor struct {
	ID        string  `json:"id"`         // 审核人ID
	Name      string  `json:"name"`       // 审核人姓名
	Account   *string `json:"account"`    // 审核人账号，可为空
	Status    string  `json:"status"`     // 审核状态
	AuditDate string  `json:"audit_date"` // 审核日期
}

// ApplyDetail 表示申请详情
type ApplyDetail struct {
	Process  Process `json:"process"` // 流程信息
	Data     string  `json:"data"`    // 申请数据，JSON字符串
	dataDict map[string]any
	Workflow ApplyDetailWorkflow `json:"workflow"` // 工作流信息
}

func (a ApplyDetail) DecodeData() map[string]any {
	ds := make(map[string]any)
	json.Unmarshal([]byte(a.Data), &ds)
	return ds
}

func (a *ApplyDetail) StrValue(key string) string {
	if a.dataDict == nil {
		a.dataDict = a.DecodeData()
	}
	return fmt.Sprintf("%v", a.dataDict[key])
}

// Process 表示流程信息
type Process struct {
	ConflictApplyID string `json:"conflict_apply_id,omitempty"` // 冲突申请ID，可选
	UserID          string `json:"user_id"`                     // 发起人用户ID
	UserName        string `json:"user_name"`                   // 发起人用户名
	ApplyID         string `json:"apply_id"`                    // 申请ID
	ProcDefKey      string `json:"proc_def_key"`                // 流程定义键
	AuditType       string `json:"audit_type"`                  // 审核类型
}

// ApplyDetailWorkflow 表示工作流信息
type ApplyDetailWorkflow struct {
	TopCSF          int             `json:"top_csf"`            // 顶级CSF值
	MsgForEmail     *string         `json:"msg_for_email"`      // 邮件消息，可为空
	MsgForLog       *string         `json:"msg_for_log"`        // 日志消息，可为空
	Content         *string         `json:"content"`            // 内容，可为空
	AbstractInfo    string          `json:"abstract_info"`      // 摘要信息
	FrontPluginInfo FrontPluginInfo `json:"front_plugin_info"`  // 前端插件信息
	Webhooks        []Webhook       `json:"webhooks,omitempty"` // Webhook列表，可选
}

// Workflow 表示工作流信息
type Workflow struct {
	TopCSF          int             `json:"top_csf"`            // 顶级CSF值
	MsgForEmail     *string         `json:"msg_for_email"`      // 邮件消息，可为空
	MsgForLog       *string         `json:"msg_for_log"`        // 日志消息，可为空
	Content         *string         `json:"content"`            // 内容，可为空
	AbstractInfo    AbstractInfo    `json:"abstract_info"`      // 摘要信息
	FrontPluginInfo FrontPluginInfo `json:"front_plugin_info"`  // 前端插件信息
	Webhooks        []Webhook       `json:"webhooks,omitempty"` // Webhook列表，可选
}

// AbstractInfo 表示摘要信息
type AbstractInfo struct {
	Icon string `json:"icon"` // 图标（Base64编码）
	Text string `json:"text"` // 文本描述
}

// FrontPluginInfo 表示前端插件信息
type FrontPluginInfo struct {
	TenantID       string            `json:"tenant_id"`       // 租户ID
	Entry          string            `json:"entry"`           // 入口URL
	Name           string            `json:"name"`            // 插件名称
	CategoryBelong string            `json:"category_belong"` // 所属类别
	Label          map[string]string `json:"label"`           // 多语言标签
	AuditType      string            `json:"audit_type"`      // 审核类型
}

// Webhook 表示webhook信息
type Webhook struct {
	Webhook     string `json:"webhook"`      // Webhook URL
	StrategyTag string `json:"strategy_tag"` // 策略标签
}

// 审核日志
type AuditNodeLog struct {
	ActStatus string           `json:"act_status"`   // 操作状态 1 审核中 2 审核完成(仅为2时有act_type)
	ActType   string           `json:"act_type"`     // 操作类型 autoPass 自动通过 autoReject 自动拒绝 userTask 用户审核task（仅类型为userTask时auditor_logs不为空）
	Logs      [][]*AuditorLogs `json:"auditor_logs"` // 审核员记录列表
}

// 审核日志项
type AuditorLogs struct {
	AuditorID   *string `json:"auditor"`      // 审核员ID
	AuditorName *string `json:"auditor_name"` // 审核员名称
	AuditIdea   *string `json:"audit_idea"`   // 审核意见 revocation 且 audit_status为reject时表示撤销
	AuditStatus *string `json:"audit_status"` // 审核结果(审核中没有) reject 驳回 pass 通过
}

type DocAuditDriven interface {
	GetMyTodoList(ctx context.Context, req *GetMyTodoListReq) (res *GetMyTodoListRes, err error)
}

type GetMyTodoListReq struct {
	DocName   string   `query:"doc_name"`
	Type      []string `query:"type"`
	Abstracts string   `query:"abstracts"`
	Limit     int      `query:"limit"`
	Offset    int      `query:"offset"`
}

type GetMyTodoListRes struct {
	TotalCount int              `json:"total_count"`
	Entries    []*TodoAuditItem `json:"entries"`
}

type TodoAuditItem struct {
	ID            string          `json:"id"`              // 申请ID
	ApplyUserName string          `json:"apply_user_name"` // 申请人用户名
	ApplyTime     string          `json:"apply_time"`      // 发起申请时间
	ApplyDetail   *ApplyDetailMsg `json:"apply_detail"`    // 申请详情
}

type ApplyDetailMsg struct {
	Process *common.AuditApplyProcessInfo `json:"process"` // 流程信息
	Data    string                        `json:"data"`    // 附加数据
}
