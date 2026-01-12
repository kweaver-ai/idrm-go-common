package nsq

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/nsqio/go-nsq"

	"github.com/kweaver-ai/idrm-go-common/workflow/common"
	mq_common "github.com/kweaver-ai/idrm-go-common/workflow/mq/common"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

type nsqSt struct {
	httpClient      *http.Client
	addr            string
	lookupdURL      string
	httpURL         string
	channel         string
	config          *nsq.Config
	producer        *nsq.Producer
	consumers       []*nsq.Consumer
	consumeHandlers mq_common.MsgConsumeHandlersInterface
}

func NewNSQ(httpClient *http.Client, conf *common.MQConf) (mq_common.MQInterface, error) {
	return &nsqSt{
		httpClient:      httpClient,
		addr:            conf.Host,
		lookupdURL:      conf.LookupdHost,
		httpURL:         conf.HttpHost,
		channel:         conf.Channel,
		config:          nsq.NewConfig(),
		consumeHandlers: mq_common.NewMsgConsumeHandlers(),
	}, nil
}

func (m *nsqSt) RegistConusmeHandlers(
	auditType string,
	hAuditProcess common.Handler[common.AuditProcessMsg],
	hAuditResult common.Handler[common.AuditResultMsg],
	hAuditProcessDefDel common.Handler[common.AuditProcDefDelMsg]) {
	m.consumeHandlers.RegistConusmeHandlers(auditType, hAuditProcess, hAuditResult, hAuditProcessDefDel)
}

func (m *nsqSt) createTopic(topic string) error {
	addr := fmt.Sprintf("http://%s/topic/create?topic=%s", m.httpURL, topic)

	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, addr, nil)
	if err != nil {
		return err
	}
	resp, err := m.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("create topic error, status: %s, header: %v, read response body fail: %w", resp.Status, resp.Header, err)
		}
		return fmt.Errorf("create topic error, status: %s, header: %v, body: %s", resp.Status, resp.Header, body)
	}
	return nil
}

func (m *nsqSt) newProducer() error {
	var err error
	if m.producer == nil {
		m.producer, err = nsq.NewProducer(m.addr, m.config)
	}
	return err
}

func handlerFunc[T common.ValidMsg](handler common.Handler[T]) nsq.HandlerFunc {
	return func(message *nsq.Message) error {
		var msg T
		if err := json.Unmarshal(message.Body, &msg); err != nil {
			log.Warnf("failed to parse audit process msg: %#v, err: %#v", msg, err)
			return err
		}
		return handler(context.Background(), &msg)
	}
}

func (m *nsqSt) newConsumer(topic string, handler nsq.HandlerFunc) error {
	if len(topic) == 0 || handler == nil {
		return nil
	}
	if err := retry.Do(
		func() error {
			return m.createTopic(topic)
		},
		retry.Attempts(3),
		retry.Delay(500*time.Millisecond),
		retry.OnRetry(func(n uint, err error) {
			if n > 0 {
				log.Warnf("failed to create nsq topic %s- %v, retry %d times ...", topic, err, n)
			}
		}),
		retry.RetryIf(func(err error) bool { return err != nil }),
		retry.MaxDelay(1*time.Second),
		retry.Context(context.Background()),
		retry.LastErrorOnly(true),
	); err != nil {
		return err
	}

	consumer, err := nsq.NewConsumer(topic, m.channel, m.config) //创建消费者
	if err == nil {
		// 定义 nsq 处理器
		consumer.AddHandler(handler)
		// 连接 lookupd->nsqd
		err = consumer.ConnectToNSQLookupd(m.lookupdURL)
		if err == nil {
			m.consumers = append(m.consumers, consumer)
		}
	}
	return err
}

func (m *nsqSt) Produce(topic string, _ []byte, msg []byte) error {
	return m.producer.Publish(topic, msg)
}

func (w *nsqSt) auditProcessMsgHandler(ctx context.Context, msg *common.AuditProcessMsg) error {
	h, ok := w.consumeHandlers.GetAuditProcessHandlers()[msg.GetAuditType()]
	if !ok {
		return nil
	}
	return h(ctx, msg)
}

func (m *nsqSt) Start() error {
	var (
		rhMap   map[string]common.Handler[common.AuditResultMsg]
		pddhMap map[string]common.Handler[common.AuditProcDefDelMsg]
	)

	err := m.newProducer()
	if err != nil {
		goto EXIT
	}

	rhMap = m.consumeHandlers.GetAuditResultHandlers()
	for k, v := range rhMap {
		err = m.newConsumer(k, handlerFunc(v))
		if err != nil {
			goto EXIT
		}
	}

	if len(m.consumeHandlers.GetAuditProcessHandlers()) > 0 {
		err = m.newConsumer(m.consumeHandlers.GetAuditProcessTopics()[0],
			handlerFunc(m.auditProcessMsgHandler))
		if err != nil {
			goto EXIT
		}
	}

	pddhMap = m.consumeHandlers.GetAuditProcessDefDelHandlers()
	for k, v := range pddhMap {
		err = m.newConsumer(k, handlerFunc(v))
		if err != nil {
			goto EXIT
		}
	}
EXIT:
	if err != nil {
		m.Stop()
	}
	return err
}

func (m *nsqSt) Stop() {
	for i := range m.consumers {
		m.consumers[i].Stop()
	}
	m.consumers = nil
	if m.producer != nil {
		m.producer.Stop()
	}
	m.producer = nil
}
