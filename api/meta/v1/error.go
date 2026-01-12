package v1

import (
	"encoding/json"
)

// Error 定义 API 返回的错误结构
type Error struct {
	// 错误码
	Code string `json:"code,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
	// 解决方案
	Solution string `json:"solution,omitempty"`
	// 详情
	Detail json.RawMessage `json:"detail,omitempty"`
}

var _ error = &Error{}

func (e *Error) Error() string {
	if e.Description == "" {
		return e.Code
	}

	return e.Code + " " + e.Description
}
