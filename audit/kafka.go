package audit

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/IBM/sarama"

	v1 "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
)

// 向 Kafka 记录日志
type kafka struct {
	client sarama.SyncProducer
}

// NewKafkaAndTopic 创建基于 Kafka 的日志器，并创建记录日志所需的 Topic
//
// 无法单独关闭日志器所使用的 SyncProducer，只能关闭 Client，可能存在泄露。
func NewKafkaAndTopic(c sarama.Client) (logger Logger, err error) {
	// 创建 Kafka Topic
	a, err := sarama.NewClusterAdminFromClient(c)
	if err != nil {
		return
	}
	defer a.Close()
	if err = NewKafkaTopic(a); err != nil {
		return
	}

	// 创建基于 Kafka 的日志器
	p, err := sarama.NewSyncProducerFromClient(c)
	if err != nil {
		return
	}
	logger = NewKafka(p)

	return
}

func NewKafka(client sarama.SyncProducer) Logger {
	return New(&kafka{client: client})
}

// LogEvent implements LogSink.
func (k *kafka) LogEvent(event *v1.Event) {
	value, err := json.Marshal(event)
	if err != nil {
		value = []byte(fmt.Sprintf("%#v", event))
	}

	m := &sarama.ProducerMessage{
		Topic: v1.KafkaTopic,
		Value: sarama.ByteEncoder(value),
	}

	// TODO: Retry
	if _, _, err := k.client.SendMessage(m); err != nil {
		panic(fmt.Sprintf("send message %s to kafka fail: %v", value, err))
	}
}

var _ LogSink = &kafka{}

// NewKafkaTopic 创建 Kafka Topic 用于记录审计日志
func NewKafkaTopic(c sarama.ClusterAdmin) (err error) {
	// 获取已存在的 Kafka Topic
	topics, err := c.ListTopics()
	if err != nil {
		return
	}

	const topic = v1.KafkaTopic

	// 如果 Kafka Topic 已存在则退出
	if _, ok := topics[topic]; ok {
		return
	}

	// 创建 Kafka Topic，忽略 Topic 已存在错误
	if err = c.CreateTopic(topic, &sarama.TopicDetail{NumPartitions: 1, ReplicationFactor: 1}, false); errors.Is(err, sarama.ErrTopicAlreadyExists) {
		err = nil
	}

	return
}
