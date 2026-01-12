package interception

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewContextWithBearerToken(t *testing.T) {
	ctx := NewContextWithBearerToken(context.Background(), "TEST_BEARER_TOKEN")

	v := ctx.Value(contextKeyBearerToken)
	assert.Equal(t, "TEST_BEARER_TOKEN", v)
}

func TestSetGinContextWithBearerToken(t *testing.T) {
	c := new(gin.Context)

	SetGinContextWithBearerToken(c, "TEST_BEARER_TOKEN")

	v, exists := c.Get(contextKeyBearerToken)
	assert.Equal(t, "TEST_BEARER_TOKEN", v)
	assert.True(t, exists)

	auth, err := BearerTokenFromContext(c)
	assert.Equal(t, "TEST_BEARER_TOKEN", auth)
	assert.NoError(t, err)
}

func TestBearerTokenFromContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "Success",
			args: args{
				ctx: NewContextWithBearerToken(context.Background(), "TEST_BEARER_TOKEN"),
			},
			want: "TEST_BEARER_TOKEN",
		},
		{
			name: "NotExist",
			args: args{
				ctx: context.Background(),
			},
			wantErr: ErrNotExist,
		},
		{
			name: "UnexpectedType",
			args: args{
				ctx: context.WithValue(context.Background(), contextKeyBearerToken, 12450),
			},
			wantErr: ErrUnexpectedType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BearerTokenFromContext(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestBearerTokenFromContextCompatible(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "set by NewContextWithBearerToken",
			args: args{ctx: NewContextWithBearerToken(context.Background(), "BEARER_TOKEN_NewContextWithBearerToken")},
			want: "BEARER_TOKEN_NewContextWithBearerToken",
		},
		{
			name: "set by SetGinContextWithBearerToken",
			args: args{ctx: newGinContextWith(SetGinContextWithBearerToken, "BEARER_TOKEN_SetGinContextWithBearerToken")},
			want: "BEARER_TOKEN_SetGinContextWithBearerToken",
		},
		{
			name: "set by NewContextWithAuth",
			args: args{ctx: NewContextWithAuth(context.Background(), "Bearer BEARER_TOKEN_NewContextWithAuth")},
			want: "BEARER_TOKEN_NewContextWithAuth",
		},
		{
			name: "set by SetGinContextWithAuth",
			args: args{ctx: newGinContextWith(SetGinContextWithAuth, "Bearer BEARER_TOKEN_SetGinContextWithAuth")},
			want: "BEARER_TOKEN_SetGinContextWithAuth",
		},
		{
			name: "set by context.WithValue BearerToken",
			args: args{ctx: context.WithValue(context.Background(), Token, "BEARER_TOKEN_WITH_VALUE_BEARER_TOKEN")},
			want: "BEARER_TOKEN_WITH_VALUE_BEARER_TOKEN",
		},
		{
			name: "set by context.WithValue Authorization",
			args: args{ctx: context.WithValue(context.Background(), Token, "Bearer BEARER_TOKEN_WITH_VALUE_AUTHORIZATION")},
			want: "BEARER_TOKEN_WITH_VALUE_AUTHORIZATION",
		},
		{
			name:    "not found",
			args:    args{ctx: context.Background()},
			wantErr: ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BearerTokenFromContextCompatible(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func newGinContextWith[T any](setter func(*gin.Context, T), value T) *gin.Context {
	c, _ := gin.CreateTestContext(nil)
	setter(c, value)
	return c
}
