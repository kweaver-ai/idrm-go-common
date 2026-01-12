package workflow

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kweaver-ai/idrm-go-common/workflow/common"
	"github.com/kweaver-ai/idrm-go-common/workflow/mq"
	mq_common "github.com/kweaver-ai/idrm-go-common/workflow/mq/common"

	"github.com/avast/retry-go"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

type workFlow struct {
	conf *common.MQConf
	mq   mq_common.MQInterface
}

var wf *workFlow = nil
var stopCh chan os.Signal

type WFStarter interface {
	Start() error
}

type WorkflowInterface interface {
	RegistConusmeHandlers(auditType string,
		hAuditProcess common.Handler[common.AuditProcessMsg],
		hAuditResult common.Handler[common.AuditResultMsg],
		hAuditProcessDefDel common.Handler[common.AuditProcDefDelMsg]) // 根据审核类型添加Consume Handler
	WFStarter
	Stop()
	AuditApply(msg *common.AuditApplyMsg) error   // 发起审核
	AuditCancel(msg *common.AuditCancelMsg) error // 撤销审核
}

func NewWorkflow(httpClient *http.Client, conf *common.MQConf) (WorkflowInterface, error) {
	if wf != nil {
		return wf, nil
	}

	var err error
	wf = &workFlow{conf: conf}
	if wf.mq, err = mq.NewMQ(httpClient, conf); err != nil {
		return nil, err
	}

	stopCh = make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		<-stopCh
		if wf.mq != nil {
			wf.mq.Stop()
		}
		wf.conf = nil
		wf.mq = nil
	}()

	return wf, err
}

func (w *workFlow) Stop() {
	stopCh <- syscall.SIGINT
	time.Sleep(1 * time.Second)
}

type Handler[T common.ValidMsg] func(ctx context.Context, auditType string, msg *T) error

func HandlerFunc[T common.ValidMsg](auditType string, handler Handler[T]) common.Handler[T] {
	return func(ctx context.Context, msg *T) error {
		return handler(ctx, auditType, msg)
	}
}

func (w *workFlow) RegistConusmeHandlers(
	auditType string,
	hAuditProcess common.Handler[common.AuditProcessMsg],
	hAuditResult common.Handler[common.AuditResultMsg],
	hAuditProcessDefDel common.Handler[common.AuditProcDefDelMsg],
) {
	w.mq.RegistConusmeHandlers(auditType, hAuditProcess, hAuditResult, hAuditProcessDefDel)
}

func (w *workFlow) Start() error {
	if err := w.mq.Start(); err != nil {
		w.mq.Stop()
		return err
	}

	return nil
}

func (w *workFlow) AuditApply(msg *common.AuditApplyMsg) error {
	buf, err := json.Marshal(msg)
	if err == nil {
		err = w.produce(common.TOPIC_PUB_NSQ_AUDIT_APPLY, buf)
	}
	return err
}

func (w *workFlow) AuditCancel(msg *common.AuditCancelMsg) error {
	buf, err := json.Marshal(msg)
	if err == nil {
		err = w.produce(common.TOPIC_PUB_NSQ_AUDIT_CANCEL, buf)
	}
	return err
}

func (w *workFlow) produce(topic string, msg []byte) error {
	return retry.Do(
		func() error {
			return w.mq.Produce(topic, nil, msg)
		},
		retry.Attempts(3),
		retry.Delay(500*time.Millisecond),
		retry.OnRetry(func(n uint, err error) {
			if n > 0 {
				log.Warnf("failed to publish msg - %v, retry %d times ...", err, n)
			}
		}),
		retry.RetryIf(func(err error) bool { return err != nil }),
		retry.MaxDelay(1*time.Second),
		retry.Context(context.Background()),
		retry.LastErrorOnly(true),
	)
}
