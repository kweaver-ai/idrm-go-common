package errorcode

import (
	"encoding/json"

	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func init() {
	RegisterErrorCode(configurationCenterErrorMap)
}

const configurationCenterPreCoder = "ServiceCall.ConfigurationCenter."
const (
	// configuration-center 未对接 AnyRobot
	ConfigurationCenterAnyRobotDisabled = configurationCenterPreCoder + "AnyRobotDisabled"

	GetThirdPartyAddr = configurationCenterPreCoder + "GetThirdPartyAddr"
	// 获取应用失败
	GetApplicationFailure = configurationCenterPreCoder + "GetApplicationFailure"
	// 指定的第三方数据源 ID 未找到
	DatasourceHuaAoIDNotFound = configurationCenterPreCoder + "DatasourceHuaAoIDNotFound"
)

var configurationCenterErrorMap = ErrorCode{
	ConfigurationCenterAnyRobotDisabled: {
		Description: "未接入 AnyRobot",
	},
	GetThirdPartyAddr: {
		Description: "获取配置中心第三方地址失败",
		Cause:       "",
		Solution:    "请重试",
	},
	GetApplicationFailure: {
		Description: "获取应用失败",
	},
	DatasourceHuaAoIDNotFound: {
		Description: "第三方数据源 ID 未找到 ",
	},
}

type GetApplicationFailureDetail struct {
	// 应用 ID
	ID string `json:"id,omitempty"`
	// http response status code
	StatusCode int `json:"status_code,omitempty"`
	// http response body
	Body json.RawMessage `json:"body,omitempty"`
}

func NewGetApplicationFailure(id string, statusCode int, body []byte) error {
	log.Error("NewGetApplicationFailure http status error", zap.String("id", id), zap.Int("status", statusCode), zap.String("body", string(body)))
	return Detail(GetApplicationFailure, &GetApplicationFailureDetail{ID: id, StatusCode: statusCode, Body: body})
}
