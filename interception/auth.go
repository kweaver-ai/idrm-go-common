package interception

import (
	"context"
	"github.com/gin-gonic/gin"
)

// contextKeyAuth 是从 context.Context 获取 HTTP Header Authorization 的
// key。
const contextKeyAuth = "GoCommon/interception.Auth"

// NewContextWithAuth 生成一个包含 BearerToken 的 context.Context
func NewContextWithAuth(parent context.Context, t string) context.Context {
	return context.WithValue(parent, contextKeyAuth, t)
}

// SetGinContextWithAuth 把 HTTP Header Authorization 保存在 gin.Context
func SetGinContextWithAuth(c *gin.Context, t string) {
	c.Set(contextKeyAuth, t)
}

// AuthFromContext 从 context.Context 获取 BearerToken，如果未找到或类型不符返回 error
func AuthFromContext(ctx context.Context) (string, error) {
	v := ctx.Value(contextKeyAuth)
	if v == nil {
		return "", ErrNotExist
	}

	t, ok := v.(string)
	if !ok {
		return "", ErrUnexpectedType
	}

	return t, nil
}

// AuthFromContextCompatible 从 context.Context 获取 BearerToken，如果未找到或类型不符返回 error
func AuthFromContextCompatible(ctx context.Context) (string, error) {
	return get(ctx, contextKeyAuth)
}
func get(ctx context.Context, key string) (string, error) {
	v := ctx.Value(key)
	if v == nil {
		if key != Token { //兼容token处理
			return get(ctx, Token)
		}
		return "", ErrNotExist
	}
	t, ok := v.(string)
	if !ok {
		return "", ErrUnexpectedType
	}

	return t, nil
}

func GetHeaderScope(c *gin.Context) string {
	return c.GetHeader(PermissionScope)
}
