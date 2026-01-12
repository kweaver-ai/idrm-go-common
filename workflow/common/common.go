package common

import (
	"context"
	"encoding/json"
)

const ( // 消费消息类型
	MSG_TYPE_AUDIT_PROCESS      = iota + 1 // 审核进度消息类型
	MSG_TYPE_AUDIT_REUSLT                  // 审核结果消息类型
	MSG_TYPE_AUDIT_PROC_DEF_DEL            // 审核流程定义删除消息
)

const (
	AUDIT_RESULT_PASS   = "pass"
	AUDIT_RESULT_REJECT = "reject"
	AUDIT_RESULT_UNDONE = "undone"
)

// 常见的审核状态，客户端保存的, 不是workflow定义的状态
const (
	AUDIT_PENDING     string = "pending"     // 待处理
	AUDIT_AUDITING    string = "auditing"    // 审批中
	AUDIT_REJECT      string = "reject"      // 申请被拒绝
	AUDIT_PASS        string = "pass"        // 申请被允许
	AUDIT_UNDONE      string = "undone"      // 申请被发起者撤回
	AUDIT_AUTHORIZING string = "authorizing" // 授权中
	AUDIT_FAILED      string = "failed"      // 失败。创建资源失败，或应用权限策略失败
	AUDIT_COMPLETED   string = "completed"   // 完成
)

const (
	AUDIT_PROCESS_TOPIC = "workflow.audit.msg" // 审核进度消息TOPIC（SUB）

	TOPIC_PUB_NSQ_AUDIT_APPLY  = "workflow.audit.apply"  // 发起审核TOPIC（PUB）
	TOPIC_PUB_NSQ_AUDIT_CANCEL = "workflow.audit.cancel" // 主动撤销审核TOPIC（PUB）
)

const (
	AUDIT_RESULT_TOPIC_PREFIX = "workflow.audit.result."      // 审核结果消息TOPIC前缀（SUB）
	PROCESS_DEL_TOPIC_PREFIX  = "workflow.audit.proc.delete." // 审核流程定义删除消息TOPIC前缀（SUB）
)

type ValidMsg interface {
	AuditProcessMsg | AuditResultMsg | AuditProcDefDelMsg
}

type Handler[T ValidMsg] func(ctx context.Context, msg *T) error

// ProcessDefiniton 流程定义信息
type ProcessDefiniton struct {
	TenantId   string `json:"tenantId"`   // 流程所属租户
	Category   string `json:"category"`   // 流程所属类型
	ProcDefKey string `json:"procDefKey"` // 审核流程key
}

// ActivityInfo 审批节点信息
type ActivityInfo struct {
	CreateTime    int64  `json:"createTime"`    // 流程开始时间
	FinishTime    int64  `json:"finishTime"`    // 流程完结时间
	Receiver      string `json:"receiver"`      // 审核员用户id
	ReceiverOrgId string `json:"receiverOrgId"` // 审核员用户名
	Sender        string `json:"sender"`        // 流程发起人用户id
	ProcInstId    string `json:"procInstId"`    // 流程实例id
	ActDefName    string `json:"actDefName"`    // 流程节点名称
	ActDefId      string `json:"actDefId"`      // 流程节点id
}

// ProcessResultFields 流程节点审核结果信息
type ProcessResultFields struct {
	AuditIdea   bool   `json:"-"`                // 审核结果 true 通过 false 拒绝
	AuditIdeaV2 *bool  `json:"auditIdea,string"` // 审核结果 true 通过 false 拒绝
	AuditMsg    string `json:"auditMsg"`         // 审批意见（超过200字超出部分会被截断替换为...）
	BizType     string `json:"bizType"`          // 业务类型
	ApplyID     string `json:"bizId"`            // 审核申请id
	FlowApplyID string `json:"applyId"`          // 审核流程ID
}

type AuditTypeInterface interface {
	GetAuditType() string
}

// AuditProcessMsg 审核进展消息
type AuditProcessMsg struct {
	ProcessDef        ProcessDefiniton `json:"processDefinition"`
	ProcInstId        string           `json:"procInstId"`
	CurrentActivity   *ActivityInfo    `json:"currentActivity"`
	NextActivity      []*ActivityInfo  `json:"nextActivity"`
	ProcessInputModel struct {
		WFProcInstId string              `json:"wf_procInstId"` // 审核实例ID
		WFCurComment string              `json:"wf_curComment"` // 完整审批意见
		Fields       ProcessResultFields `json:"fields"`
	} `json:"processInputModel"`
}

func (apm *AuditProcessMsg) UnmarshalJSON(data []byte) error {
	type Alias AuditProcessMsg
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(apm),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.ProcessInputModel.Fields.AuditIdeaV2 != nil && *aux.ProcessInputModel.Fields.AuditIdeaV2 {
		aux.ProcessInputModel.Fields.AuditIdea = true
	}
	return nil
}

func (apm *AuditProcessMsg) GetAuditMsg() *string {
	if len(apm.ProcessInputModel.Fields.AuditMsg) > 0 {
		return &apm.ProcessInputModel.WFCurComment
	}
	auditAdvice := apm.ProcessInputModel.Fields.AuditMsg
	// workflow 里不填审核意见时默认是 default_comment, 排除这种情况
	if auditAdvice == "default_comment" {
		auditAdvice = ""
	}
	return &auditAdvice
}

func (apm *AuditProcessMsg) GetAuditType() string {
	return apm.ProcessDef.Category
}

// AuditProcDefDelMsg 审核流程删除消息
type AuditProcDefDelMsg struct {
	ProcDefKeys []string `json:"proc_def_keys"` // 被删除的审核流程定义key集合
}

// AuditResultMsg 流程审核最终结果
type AuditResultMsg struct {
	ApplyID string `json:"apply_id"` // 审核申请id
	Result  string `json:"result"`   // 审核结果 "pass": 通过  "reject": 拒绝  "undone": 撤销
}

// AuditApplyProcessInfo 审核流程信息
type AuditApplyProcessInfo struct {
	AuditType  string `json:"audit_type"`
	ApplyID    string `json:"apply_id"`
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	ProcDefKey string `json:"proc_def_key"`
}

// AuditApplyAbstractInfo
type AuditApplyAbstractInfo struct {
	Icon string `json:"icon"`
	Text string `json:"text"`
}

// Webhook
type Webhook struct {
	Webhook     string `json:"webhook"`
	StrategyTag string `json:"strategy_tag"`
}

// AuditApplyWorkflowInfo
type AuditApplyWorkflowInfo struct {
	TopCsf       int                    `json:"top_csf"`
	AbstractInfo AuditApplyAbstractInfo `json:"abstract_info"`
	Webhooks     []Webhook              `json:"webhooks"`
}

// AuditApplyMsg 发起流程消息
type AuditApplyMsg struct {
	Process  AuditApplyProcessInfo  `json:"process"`
	Data     map[string]any         `json:"data"`
	Workflow AuditApplyWorkflowInfo `json:"workflow"`
}

// AuditCancelMsg 撤销流程消息
type AuditCancelMsg struct {
	ApplyIDs []string `json:"apply_ids"` // 申请ID数组
	Cause    struct {
		ZHCN string `json:"zh-cn"` // 中文
		ZHTW string `json:"zh-tw"` // 繁体
		ENUS string `json:"en-us"` // 英文
	} `json:"cause"` // 撤销原因
}

func GenNormalCancelMsg(id ...string) *AuditCancelMsg {
	return &AuditCancelMsg{
		ApplyIDs: id,
		Cause: struct {
			ZHCN string "json:\"zh-cn\""
			ZHTW string "json:\"zh-tw\""
			ENUS string "json:\"en-us\""
		}{
			ZHCN: "revocation",
			ZHTW: "revocation",
			ENUS: "revocation",
		},
	}
}

type MQConf struct {
	MqType      string    `json:"mqType,omitempty"` // MQ类型 nsq kafka
	Host        string    `json:"host,omitempty"`
	HttpHost    string    `json:"httpHost,omitempty"`   // http host，NSQ必传
	LookupdHost string    `json:"ookupdHost,omitempty"` // lookupd host, NSQ必传
	Channel     string    `json:"channel,omitempty"`
	Sasl        *Sasl     `json:"sasl,omitempty"`     // KAFKA必传
	Producer    *Producer `json:"producer,omitempty"` // KAFKA必传
	Version     string    `json:"version"`            // KAFKA必传
}

type Producer struct {
	SendBufSize int32 `json:"sendBufSize,omitempty"`
	RecvBufSize int32 `json:"recvBufSize,omitempty"`
}

type Sasl struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Mechanism string `json:"mechanism,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

const (
	MQ_TYPE_NSQ   = "nsq"
	MQ_TYPE_KAFKA = "kafka"
)

func GetAuditMsg(curComment, auditMsg *string) *string {
	if len(*auditMsg) > 0 {
		return curComment
	}
	return auditMsg
}
