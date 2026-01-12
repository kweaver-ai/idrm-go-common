package anyrobot

import (
	dataModelV1 "github.com/kweaver-ai/idrm-go-common/rest/anyrobot/data-model/v1"
	uniqueryV1 "github.com/kweaver-ai/idrm-go-common/rest/anyrobot/uniquery/v1"
)

type Interface interface {
	DataModelV1() dataModelV1.Interface
	UniqueryV1() uniqueryV1.Interface
}
