package interception

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSeAuthorizationIfEmpty(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		auth string
		want string
	}{
		{
			name: "not empty",
			auth: "NOT_EMPTY",
			want: "NOT_EMPTY",
		},
		{
			name: "without authorization",
			ctx:  context.Background(),
		},
		{
			name: "NewContextWithBearerToken",
			ctx:  NewContextWithBearerToken(context.Background(), "BEARER_TOKEN_0000"),
			want: "Bearer BEARER_TOKEN_0000",
		},
		{
			name: "NewContextWithBearerToken",
			ctx: func() context.Context {
				c := &gin.Context{}
				SetGinContextWithBearerToken(c, "BEARER_TOKEN_1111")
				return c
			}(),
			want: "Bearer BEARER_TOKEN_1111",
		},
		{
			name: "NewContextWithAuth",
			ctx:  NewContextWithAuth(context.Background(), "HTTP_AUTH_0000"),
			want: "HTTP_AUTH_0000",
		},
		{
			name: "SetGinContextWithAuth",
			ctx: func() context.Context {
				c := &gin.Context{}
				SetGinContextWithAuth(c, "HTTP_AUTH_1111")
				return c
			}(),
			want: "HTTP_AUTH_1111",
		},
		{
			name: "context.WithValue",
			ctx:  context.WithValue(context.Background(), Token, "TOKEN_0000"),
			want: "TOKEN_0000",
		},
		{
			name: "gin.Context.Set",
			ctx: func() context.Context {
				c := &gin.Context{}
				c.Set(Token, "TOKEN_1111")
				return c
			}(),
			want: "TOKEN_1111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := make(http.Header)
			h.Set("Authorization", tt.auth)
			SeAuthorizationIfEmpty(tt.ctx, h)
			assert.Equal(t, tt.want, h.Get("Authorization"))
		})
	}
}
