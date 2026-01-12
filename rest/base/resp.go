package base

import (
	"errors"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
)

const (
	SUCCESS_CODE = "0"
	SUCCESS_TAG  = "success"
)

type IntIDResp struct {
	ID int `json:"id" binding:"required,uuid" example:"3ccd8d5a76b711edb78d00505697bd0b"` // 资源对象ID
}

type IDNameResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CommonResponse[T any] struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Solution    string `json:"solution"`
	Data        T      `json:"data"`
}

func (c CommonResponse[T]) Error() error {
	if c.Code == SUCCESS_CODE || c.Code == SUCCESS_TAG {
		return nil
	}
	return errors.New(c.Description)
}

func (c CommonResponse[T]) ErrorCode() error {
	if c.Code == SUCCESS_CODE || c.Code == SUCCESS_TAG {
		return nil
	}
	return errorcode.ManualNew(c.Code, errorcode.ErrorCodeInfo{
		Description: c.Description,
		Solution:    c.Solution,
		Cause:       "",
	})
}

type PageResult[T any] struct {
	Entries    []*T  `json:"entries" binding:"required"`                       // 对象列表
	TotalCount int64 `json:"total_count" binding:"required,gte=0" example:"3"` // 当前筛选条件下的对象数量
}
