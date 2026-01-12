package common

import (
	"github.com/kweaver-ai/idrm-go-common/workflow/common"
)

type MQInterface interface {
	RegistConusmeHandlers(
		auditType string,
		hAuditProcess common.Handler[common.AuditProcessMsg],
		hAuditResult common.Handler[common.AuditResultMsg],
		hAuditProcessDefDel common.Handler[common.AuditProcDefDelMsg])
	Produce(topic string, key []byte, msg []byte) error
	Start() error
	Stop()
}

type MsgConsumeHandlersInterface interface {
	RegistConusmeHandlers(
		auditType string,
		hAuditProcess common.Handler[common.AuditProcessMsg],
		hAuditResult common.Handler[common.AuditResultMsg],
		hAuditProcessDefDel common.Handler[common.AuditProcDefDelMsg])
	GetAuditProcessTopics() []string
	GetAuditProcessHandlers() map[string]common.Handler[common.AuditProcessMsg]
	GetAuditResultTopics() []string
	GetAuditResultHandlers() map[string]common.Handler[common.AuditResultMsg]
	GetAuditProcessDefDelTopics() []string
	GetAuditProcessDefDelHandlers() map[string]common.Handler[common.AuditProcDefDelMsg]
}

func NewMsgConsumeHandlers() MsgConsumeHandlersInterface {
	return &msgConsumeHandlers{
		auditResultTopics:       make([]string, 0),
		processDelTopics:        make([]string, 0),
		auditProcessMsgHandlers: make(map[string]common.Handler[common.AuditProcessMsg]),
		auditResultMsgHandlers:  make(map[string]common.Handler[common.AuditResultMsg]),
		processDelMsgHandlers:   make(map[string]common.Handler[common.AuditProcDefDelMsg]),
	}
}

type msgConsumeHandlers struct {
	auditProcessMsgHandlers map[string]common.Handler[common.AuditProcessMsg]
	auditResultTopics       []string
	auditResultMsgHandlers  map[string]common.Handler[common.AuditResultMsg]
	processDelTopics        []string
	processDelMsgHandlers   map[string]common.Handler[common.AuditProcDefDelMsg]
}

func (mh *msgConsumeHandlers) RegistConusmeHandlers(
	auditType string,
	hAuditProcess common.Handler[common.AuditProcessMsg],
	hAuditResult common.Handler[common.AuditResultMsg],
	hAuditProcessDefDel common.Handler[common.AuditProcDefDelMsg]) {
	if hAuditProcess != nil {
		mh.auditProcessMsgHandlers[auditType] = hAuditProcess
	}
	if hAuditResult != nil {
		topic := common.AUDIT_RESULT_TOPIC_PREFIX + auditType
		mh.auditResultTopics = append(mh.auditResultTopics, topic)
		mh.auditResultMsgHandlers[topic] = hAuditResult
	}
	if hAuditProcessDefDel != nil {
		topic := common.PROCESS_DEL_TOPIC_PREFIX + auditType
		mh.processDelTopics = append(mh.processDelTopics, topic)
		mh.processDelMsgHandlers[topic] = hAuditProcessDefDel
	}
}

func (mh *msgConsumeHandlers) GetAuditProcessTopics() []string {
	return []string{common.AUDIT_PROCESS_TOPIC}
}

func (mh *msgConsumeHandlers) GetAuditResultTopics() []string {
	return mh.auditResultTopics
}

func (mh *msgConsumeHandlers) GetAuditProcessDefDelTopics() []string {
	return mh.processDelTopics
}

func (mh *msgConsumeHandlers) GetAuditProcessHandlers() map[string]common.Handler[common.AuditProcessMsg] {
	return mh.auditProcessMsgHandlers
}

func (mh *msgConsumeHandlers) GetAuditResultHandlers() map[string]common.Handler[common.AuditResultMsg] {
	return mh.auditResultMsgHandlers
}

func (mh *msgConsumeHandlers) GetAuditProcessDefDelHandlers() map[string]common.Handler[common.AuditProcDefDelMsg] {
	return mh.processDelMsgHandlers
}
