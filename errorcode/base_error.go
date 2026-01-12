package errorcode

import (
	"encoding/json"
	"io"
	"net/http"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

func init() {
	RegisterErrorCode(baseErrorMap)
}

const baseErrorPreCode = "ServiceCall.Base."
const (
	BadRequestError        = baseErrorPreCode + "BadRequestError"
	DoRequestError         = baseErrorPreCode + "DoRequestError"
	ReadResponseError      = baseErrorPreCode + "ReadResponseError"
	RequestFailedError     = baseErrorPreCode + "RequestFailedErrors"
	UnmarshalResponseError = baseErrorPreCode + "UnmarshalResponseError"
)

var baseErrorMap = ErrorCode{

	BadRequestError: {
		Description: "调用服务参数或配置错误",
		Solution:    "请检查参数和调用服务的配置",
	},
	DoRequestError: {
		Description: "服务请求错误",
		Solution:    "请检查参数和调用服务的配置",
	},
	ReadResponseError: {
		Description: "读取服务返回数据错误",
		Solution:    "请联系技术人员",
	},
	RequestFailedError: {
		Description: "服务请求错误",
		Solution:    "请检查参数和调用服务的配置",
	},
	UnmarshalResponseError: {
		Description: "解析服务返回值错误",
		Solution:    "请检查服务和返回值",
	},
}

// err 是 http.Client.Do() 返回的 error
func NewDoRequestError(err error) error {
	return Detail(DoRequestError, err.Error())
}

type UnmarshalResponseErrorDetail struct {
	// error json.Unmarshal return
	Error string `json:"error,omitempty"`
	// response body was unmarshal by json
	Body json.RawMessage `json:"body,omitempty"`
}

// 用于在错误码中记录 Body
type stringOrJSONRawMessage []byte

// MarshalJSON implements json.Marshaler.
func (d stringOrJSONRawMessage) MarshalJSON() ([]byte, error) {
	if json.Valid(d) {
		return json.RawMessage(d).MarshalJSON()
	}
	return json.Marshal(string(d))
}

var _ json.Marshaler = &stringOrJSONRawMessage{}

func NewRequestFailedError(resp *http.Response, body []byte) error {
	detail := make(map[string]any)
	// HTTP Request
	if req := resp.Request; req != nil {
		detail["request"] = map[string]any{
			"method": req.Method,
			"url":    req.URL.String(),
			"header": req.Header,
		}
	}
	// HTTP Response
	detail["response"] = map[string]any{
		"status": resp.Status,
		"header": resp.Header,
		"body":   stringOrJSONRawMessage(body),
	}
	return WithDetail(RequestFailedError, detail)
}

func NewUnmarshalResponseError(err error, responseBody []byte) error {
	return Detail(UnmarshalResponseError, &UnmarshalResponseErrorDetail{Error: err.Error(), Body: responseBody})
}

func NewErrorForHTTPResponse(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Detail(ReadResponseError, err.Error())
	}

	var metaErr meta_v1.Error
	if err := json.Unmarshal(body, &metaErr); err != nil || metaErr.Code == "" {
		return NewRequestFailedError(resp, body)
	}

	return &metaErr
}
