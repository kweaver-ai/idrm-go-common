package errorcode

import (
	"encoding/json"
	"io"
	"net/http"

	v1 "github.com/kweaver-ai/idrm-go-common/api/anyrobot/uniquery/v1"
)

func init() {
	RegisterErrorCode(anyRobotErrorMap)
}

const anyRobotUniqueryPreCoder = "ServiceCall.AnyRobot.Uniquery."

const (
	// AnyRobot uniquery 返回的服务端错误
	AnyRobotUniqueryServerError = anyRobotUniqueryPreCoder + "ServerError"
)

var anyRobotErrorMap = ErrorCode{
	AnyRobotUniqueryServerError: {
		Description: "AnyRobot uniquery 服务端错误",
	},
}

func NewAnyRobotUniqueryServerError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Detail(ReadResponseError, err.Error())
	}

	var status v1.Status
	if err := json.Unmarshal(body, &status); err != nil || status.ErrorCode == "" {
		return NewRequestFailedError(resp, body)
	}

	return Detail(AnyRobotUniqueryServerError, status)
}
