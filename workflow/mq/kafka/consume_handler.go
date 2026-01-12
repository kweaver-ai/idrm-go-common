package kafka

import (
	"encoding/json"

	"github.com/kweaver-ai/idrm-go-common/workflow/common"

	"github.com/IBM/sarama"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

type consumerHandler[T common.ValidMsg] struct {
	hMap map[string]common.Handler[T]
}

func NewConsumerHandler[T common.ValidMsg](hMap map[string]common.Handler[T]) sarama.ConsumerGroupHandler {
	return &consumerHandler[T]{hMap: hMap}
}

func (c *consumerHandler[T]) Setup(session sarama.ConsumerGroupSession) error {
	log.Infof("start recv msg from topic: %v", session.Claims())
	return nil
}

func (c *consumerHandler[T]) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Infof("end recv msg from topic: %v", session.Claims())
	return nil
}

func (c *consumerHandler[T]) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			var msg T
			if err := json.Unmarshal(message.Value, &msg); err == nil {
				key := message.Topic
				if atInterface, ok := any(&msg).(common.AuditTypeInterface); ok {
					key = atInterface.GetAuditType()
				}

				handler, exists := c.hMap[key]
				if exists {
					if err = handler(session.Context(), &msg); err != nil {
						log.Warnf("handle msg failed: %s, reserve it", err.Error())
						break
					}
				} else {
					log.Infof("no handler found for %s, discard it", key)
				}
			} else {
				log.Warnf("json.Unmarshal msg failed: %s, discard it", err.Error())
			}
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
