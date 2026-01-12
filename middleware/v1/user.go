package v1

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
)

// NewContextWithUser 返回一个新的、包含指定用户的 Context
func NewContextWithUser(parent context.Context, user *middleware.User) context.Context {
	return context.WithValue(parent, interception.InfoName, user)
}

// UserFromContext 从 context 中获取当前用户
//
// 因为 User 在 middleware 中定义，所以 UserFromContext 不能在 interception 中定
// 义，否则会产生循环 import
//   - interception <- middleware
//   - middleware <- user_management
//   - rest/user_management <- rest/base
//   - rest/base <- interception
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
