package reconcile

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/transport/mq/kafkax"
)

func NewKafkaMsgHandleFunc[T any](r Reconciler[T]) kafkax.MsgHandleFunc {
	return func(ctx context.Context, msg *kafkax.Message) bool {
		var event meta_v1.WatchEvent[T]
		// 反序列化 Kafka 消息
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Warn("unmarshal kafka message value fail", zap.Error(err), zap.Any("msg", msg))
			// 返回 nil 为了忽略这个消息继续消费
			return true
		}
		return r.Reconcile(ctx, &event) == nil
	}
}
