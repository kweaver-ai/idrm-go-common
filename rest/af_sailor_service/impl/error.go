package impl

import (
	"encoding/json"

	"github.com/kweaver-ai/idrm-go-frame/core/errorx/agcodes"
	"github.com/kweaver-ai/idrm-go-frame/core/errorx/agerrors"
)

const (
	errPrefix            = "CognitiveAssistant.Client."
	errRequestParamError = errPrefix + "RequestParamError"
	errRequestError      = errPrefix + "RequestError"
)

type ErrorBody struct {
	ErrorCode   string `json:"errorcode"`
	Description string `json:"description"`
	Solution    string `json:"solution"`
}

func ReqParamError(err error) error {
	return agerrors.NewCode(agcodes.New(errRequestParamError, "请求参数错误", "", "请检查参数", err, ""))
}

func ReqError(err error) error {
	return agerrors.NewCode(agcodes.New(errRequestError, "请求错误", "", "请检查", err, ""))
}

func RespBodyError(body string) error {
	errObj := new(ErrorBody)
	if err := json.Unmarshal([]byte(body), errObj); err != nil {
		return ReqError(err)
	}
	return agerrors.NewCode(agcodes.New(errRequestError, errObj.Description, "", errObj.Solution, "", ""))
}
