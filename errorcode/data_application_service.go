package errorcode

import "encoding/json"

func init() {
	RegisterErrorCode(dataApplicationServiceErrorMap)
}

const dataApplicationServicePreCoder = "ServiceCall.DataApplicationService."

const (
	// 获取接口服务失败
	GetServiceFailure = dataApplicationServicePreCoder + "GetServiceFailure"
)

var dataApplicationServiceErrorMap = ErrorCode{
	GetServiceFailure: {
		Description: "获取接口服务失败",
	},
}

type GetServiceFailureDetail struct {
	// ID
	ID string `json:"id,omitempty"`
	// http response status code
	StatusCode int `json:"status_code,omitempty"`
	// http response body
	Body json.RawMessage `json:"body,omitempty"`
}

func NewGetServiceFailure(id string, statusCode int, body []byte) error {
	return Detail(GetServiceFailure, &GetServiceFailureDetail{ID: id, StatusCode: statusCode, Body: body})
}
