package mq

import (
	"errors"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/workflow/common"
	mq_common "github.com/kweaver-ai/idrm-go-common/workflow/mq/common"
	"github.com/kweaver-ai/idrm-go-common/workflow/mq/kafka"
	"github.com/kweaver-ai/idrm-go-common/workflow/mq/nsq"
)

func NewMQ(httpClient *http.Client, conf *common.MQConf) (mq_common.MQInterface, error) {
	switch conf.MqType {
	case common.MQ_TYPE_NSQ:
		return nsq.NewNSQ(httpClient, conf)
	case common.MQ_TYPE_KAFKA:
		return kafka.NewKafka(conf)
	default:
		return nil, errors.New("unknown mq type")
	}
}
