package interception

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

// contextKeyBearerToken 是从 context.Context 获取 BearerToken 的 key。
//
// 对于 HTTP Header "Authorization: Bearer ory_at_xxxx"，BearerToken 是其中的
// "ory_at_xxxx"
const contextKeyBearerToken = "GoCommon/interception.BearerToken"

// NewContextWithBearerToken 生成一个包含 BearerToken 的 context.Context
func NewContextWithBearerToken(parent context.Context, t string) context.Context {
	return context.WithValue(parent, contextKeyBearerToken, t)
}

// SetGinContextWithBearerToken 把 BearerToken 保存在 gin.Context
func SetGinContextWithBearerToken(c *gin.Context, t string) {
	c.Set(contextKeyBearerToken, t)
}

// BearerTokenFromContext 从 context.Context 获取 BearerToken，如果未找到或类型不符返回 error
func BearerTokenFromContext(ctx context.Context) (string, error) {
	v := ctx.Value(contextKeyBearerToken)
	if v == nil {
		return "", ErrNotExist
	}

	t, ok := v.(string)
	if !ok {
		return "", ErrUnexpectedType
	}

	return t, nil
}

// BearerTokenFromContextCompatible 以下列顺序依次尝试获取 BearerToken
//  1. BearerTokenFromContext: 通过 NewContextWithBearerToken, SetGinContextWithBearerToken 设置
//  2. AuthFromContext: 通过 NewContextWithAuth, SetGinContextWithAuth 设置
//  3. ctx.Value(Token): 通过 context.WithValue, gin.Context.Set 设置
func BearerTokenFromContextCompatible(ctx context.Context) (string, error) {
	if t, err := BearerTokenFromContext(ctx); err == nil {
		return t, nil
	}

	if auth, err := AuthFromContext(ctx); err == nil {
		scheme, parameters, found := strings.Cut(auth, " ")
		if found && scheme == "Bearer" {
			return parameters, nil
		}
	}

	if t, ok := ctx.Value(Token).(string); ok {
		if scheme, parameters, found := strings.Cut(t, " "); found {
			if found && scheme == "Bearer" {
				return parameters, nil
			}
		}
		return t, nil
	}

	return "", ErrNotExist
}
