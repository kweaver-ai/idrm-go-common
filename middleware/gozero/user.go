package gozero

import (
	"context"
	"errors"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
)

var (
	// ErrNotExist 表示上下文中不存在该值
	ErrNotExist = errors.New("value does not exist")
	// ErrUnexpectedType 表示上下文中的值类型不符合预期
	ErrUnexpectedType = errors.New("unexpected value type for context key")
)

// NewContextWithUser 返回一个新的、包含指定用户的 Context
func NewContextWithUser(parent context.Context, user *middleware.User) context.Context {
	return context.WithValue(parent, interception.InfoName, user)
}

// UserFromContext 从 context 中获取当前用户
// 因为 User 在 middleware 中定义，所以 UserFromContext 不能在 interception 中定义
func UserFromContext(ctx context.Context) (*middleware.User, error) {
	v := ctx.Value(interception.InfoName)
	if v == nil {
		return nil, ErrNotExist
	}

	u, ok := v.(*middleware.User)
	if !ok {
		return nil, ErrUnexpectedType
	}

	return u, nil
}
