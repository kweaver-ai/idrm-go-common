package interception

import (
	"context"

	"github.com/gin-gonic/gin"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
)

// contextKeyAuthServiceSubject 是从 context.Context 获取 auth-service/v1.Subject
// 的 key。
type contextKeyAuthServiceSubject struct{}

func (contextKeyAuthServiceSubject) String() string { return contextKeyAuthServiceSubjectString }

// contextKeyAuthServiceSubjectString 是从 context.Context 获取
// auth-service/v1.Subject的 key。
const contextKeyAuthServiceSubjectString = "GoCommon/interception.AuthServiceSubject"

func NewContextWithAuthServiceSubject(ctx context.Context, s *v1.Subject) context.Context {
	return context.WithValue(ctx, contextKeyAuthServiceSubject{}, s)
}

// SetGinContextWithAuthServiceSubject 把 auth-service.Subject 保存在
// gin.Context
func SetGinContextWithAuthServiceSubject(c *gin.Context, s *v1.Subject) {
	c.Set(contextKeyAuthServiceSubject{}.String(), s)
}

// AuthServiceSubjectFromContext 从 context.Context 获取 auth-service.Subject，
// 如果未找到或类型不符返回 error
func AuthServiceSubjectFromContext(ctx context.Context) (*v1.Subject, error) {
	for _, k := range []any{
		contextKeyAuthServiceSubject{},
		contextKeyAuthServiceSubjectString,
	} {
		v := ctx.Value(k)
		if v == nil {
			continue
		}
		s, ok := v.(*v1.Subject)
		if !ok {
			return nil, ErrUnexpectedType
		}
		return s, nil
	}
	return nil, ErrNotExist
}
