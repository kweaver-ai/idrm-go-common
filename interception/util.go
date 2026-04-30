package interception

import (
	"context"
	"fmt"
	"net/http"
)

// SeAuthorizationIfEmpty 如果 HTTP Header: Authorization 为空，以下列顺序依次尝试设置 Authorization，直到设置成功
//  1. BearerTokenFromContext: 通过 NewContextWithBearerToken, SetGinContextWithBearerToken 设置
//  2. AuthFromContext: 通过 NewContextWithAuth, SetGinContextWithAuth 设置
//  3. ctx.Value(Token): 通过 context.WithValue, gin.Context.Set 设置
func SeAuthorizationIfEmpty(ctx context.Context, h http.Header) {
	// 如果已经设置 Authorization，不再重复设置
	if h.Get("Authorization") != "" {
		return
	}

	if t, err := BearerTokenFromContext(ctx); err == nil {
		h.Set("Authorization", fmt.Sprintf("Bearer %s", t))
		return
	}

	if auth, err := AuthFromContext(ctx); err == nil {
		h.Set("Authorization", auth)
		return
	}

	if t, ok := ctx.Value(Token).(string); ok {
		h.Set("Authorization", t)
		return
	}
}

func AuthorizationHeader(ctx context.Context) http.Header {
	h := http.Header{}
	SeAuthorizationIfEmpty(ctx, h)
	return h
}

func AuthorizationHeaderMap(ctx context.Context) map[string]string {
	h := AuthorizationHeader(ctx)
	result := map[string]string{}
	for k, v := range h {
		result[k] = v[0]
	}
	return result
}
