package kafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/avast/retry-go"

	"github.com/kweaver-ai/idrm-go-common/workflow/common"
	mq_common "github.com/kweaver-ai/idrm-go-common/workflow/mq/common"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

type kafkaSt struct {
	addr            string
	channel         string
	config          *sarama.Config
	broker          *sarama.Broker
	producer        sarama.SyncProducer
	consumerGroups  []sarama.ConsumerGroup
	ctx             context.Context
	cancel          context.CancelFunc
	consumeHandlers mq_common.MsgConsumeHandlersInterface
}

func NewKafka(conf *common.MQConf) (mq_common.MQInterface, error) {
	k := &kafkaSt{
		addr:            conf.Host,
		channel:         conf.Channel,
		broker:          sarama.NewBroker(conf.Host),
		config:          sarama.NewConfig(),
		consumeHandlers: mq_common.NewMsgConsumeHandlers(),
	}

	var err error
	k.ctx, k.cancel = context.WithCancel(context.Background())
	if k.config.Version, err = sarama.ParseKafkaVersion(conf.Version); err != nil {
		return nil, err
	}
	k.config.Producer.Timeout = 100 * time.Millisecond
	k.config.Net.SASL.Enable = conf.Sasl.Enabled
	k.config.Net.SASL.Mechanism = sarama.SASLMechanism(conf.Sasl.Mechanism)
	k.config.Net.SASL.User = conf.Sasl.Username
	k.config.Net.SASL.Password = conf.Sasl.Password
	k.config.Net.SASL.Handshake = true
	k.config.Producer.Return.Successes = true
	k.config.Producer.Return.Errors = true
	k.config.Producer.RequiredAcks = sarama.WaitForAll
	k.config.Producer.Partitioner = sarama.NewRandomPartitioner
	k.config.Consumer.Offsets.AutoCommit.Enable = true
	k.consumerGroups = make([]sarama.ConsumerGroup, 0)
	err = k.broker.Open(k.config)
	if err == nil {
		var isConnected bool
		isConnected, err = k.broker.Connected()
		if err == nil && !isConnected {
			err = errors.New("kafka broker not connected")
		}
	}

	return k, err
}

func (k *kafkaSt) RegistConusmeHandlers(
	auditType string,
	hAuditProcess common.Handler[common.AuditProcessMsg],
	hAuditResult common.Handler[common.AuditResultMsg],
	hAuditProcessDefDel common.Handler[common.AuditProcDefDelMsg]) {
	k.consumeHandlers.RegistConusmeHandlers(auditType, hAuditProcess, hAuditResult, hAuditProcessDefDel)
}

func (k *kafkaSt) createTopic(topics []string) error {
	req := &sarama.CreateTopicsRequest{
		Timeout:      15 * time.Second,
		TopicDetails: make(map[string]*sarama.TopicDetail),
	}
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
		ReplicaAssignment: make(map[int32][]int32),
		ConfigEntries:     make(map[string]*string),
	}
	for i := range topics {
		req.TopicDetails[topics[i]] = topicDetail
	}

	// kafka controller，只有 controller 才能创建 topic。
	var controller *sarama.Broker

	// 从 metadata 获取 controller id 和 broker 列表
	metadataRequest := sarama.NewMetadataRequest(k.config.Version, nil)
	metadataRequest.AllowAutoTopicCreation = k.config.Metadata.AllowAutoTopicCreation
	metadataResponse, err := k.broker.GetMetadata(metadataRequest)
	if err != nil {
		return fmt.Errorf("get metadata fail: %w", err)
	}
	for _, b := range metadataResponse.Brokers {
		if b.ID() != metadataResponse.ControllerID {
			continue
		}
		controller = b
	}
	if controller == nil {
		return errors.New("kafka controller not found")
	}

	// 连接 kafka controller
	if err := controller.Open(k.config); err != nil {
		return fmt.Errorf("connect kafka controller fail: %w", err)
	}
	defer controller.Close()

	// 使用 controller 创建 topic
	resp, err := controller.CreateTopics(req)
	if err == nil {
		for topic, tErr := range resp.TopicErrors {
			if tErr != nil {
				if tErr.Err == sarama.ErrNoError || tErr.Err == sarama.ErrTopicAlreadyExists {
					continue
				}
				err = fmt.Errorf("topic %s create failed: %s", topic, tErr.Err.Error())
				break
			}
		}
	}
	return err
}

func (k *kafkaSt) newProducer() error {
	var err error
	if k.producer == nil {
		k.producer, err = sarama.NewSyncProducer([]string{k.addr}, k.config)
	}
	return err
}

func (k *kafkaSt) newConsumer(topics []string, handler sarama.ConsumerGroupHandler) error {
	if len(topics) == 0 || handler == nil {
		return nil
	}
	if err := retry.Do(
		func() error {
			return k.createTopic(topics)
		},
		retry.Attempts(3),
		retry.Delay(500*time.Millisecond),
		retry.OnRetry(func(n uint, err error) {
			if n > 0 {
				log.Warnf("failed to create kafka topic %#v- %#v, retry %d times ...", topics, err, n)
			}
		}),
		retry.RetryIf(func(err error) bool { return err != nil }),
		retry.MaxDelay(1*time.Second),
		retry.Context(context.Background()),
		retry.LastErrorOnly(true),
	); err != nil {
		return err
	}

	consumerGroup, err := sarama.NewConsumerGroup([]string{k.addr}, k.channel, k.config)
	if err != nil {
		return err
	}
	k.consumerGroups = append(k.consumerGroups, consumerGroup)
	ctx, _ := context.WithCancel(k.ctx)
	go func(topic []string) {
		for {
			err = consumerGroup.Consume(ctx, topic, handler) //创建消费者
			if err != nil {
				log.Warnf("consumer group consume topics %#v failed: %#v", topic, err)
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}(topics)

	return err
}

func (k *kafkaSt) Produce(topic string, key []byte, msg []byte) error {
	_, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(msg),
	})
	return err
}

func (k *kafkaSt) Start() error {
	err := k.newProducer()
	if err == nil {
		err = k.newConsumer(k.consumeHandlers.GetAuditResultTopics(),
			NewConsumerHandler(k.consumeHandlers.GetAuditResultHandlers()))
		if err == nil {
			err = k.newConsumer(k.consumeHandlers.GetAuditProcessTopics(),
				NewConsumerHandler(k.consumeHandlers.GetAuditProcessHandlers()))
			if err == nil {
				err = k.newConsumer(k.consumeHandlers.GetAuditProcessDefDelTopics(),
					NewConsumerHandler(k.consumeHandlers.GetAuditProcessDefDelHandlers()))
			}
		}
	}

	if err != nil {
		k.Stop()
	}
	return err
}

func (k *kafkaSt) Stop() {
	if k.cancel != nil {
		k.cancel()
	}
	k.cancel = nil
	k.ctx = nil

	if k.broker != nil {
		isConnected, err := k.broker.Connected()
		if err == nil && isConnected {
			k.broker.Close()
		}
	}
	k.broker = nil

	for i := range k.consumerGroups {
		k.consumerGroups[i].Close()
		k.consumerGroups[i] = nil
	}
	k.consumerGroups = nil

	if k.producer != nil {
		k.producer.Close()
		k.producer = nil
	}
}
