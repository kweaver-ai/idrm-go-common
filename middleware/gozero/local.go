package gozero

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
)

// LocalToken 本地开发用 Token 中间件
// 从 HTTP Header 中获取用户信息并设置到 context
func LocalToken() MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenID := r.Header.Get("Authorization")
			userInfo := &middleware.User{
				ID:   r.Header.Get("user_id"),
				Name: r.Header.Get("user_name"),
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, interception.InfoName, userInfo)
			ctx = context.WithValue(ctx, interception.Token, tokenID)
			ctx = context.WithValue(ctx, interception.TokenType, interception.TokenTypeUser)
			ctx = interception.NewContextWithAuth(ctx, tokenID)

			next(w, r.WithContext(ctx))
		}
	}
}

// LocalAppToken 本地开发用应用 Token 中间件
// 从 HTTP Header 中获取应用信息并设置到 context
func LocalAppToken() MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenID := r.Header.Get("Authorization")
			token := strings.TrimPrefix(tokenID, "Bearer ")
			tokenType, _ := strconv.Atoi(r.Header.Get("token_type"))

			userInfo := &middleware.User{
				ID:       r.Header.Get("id"),
				Name:     r.Header.Get("name"),
				UserType: tokenType,
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, interception.InfoName, userInfo)
			ctx = context.WithValue(ctx, interception.Token, token)
			ctx = context.WithValue(ctx, interception.TokenType, tokenType)

			next(w, r.WithContext(ctx))
		}
	}
}
