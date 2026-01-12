package interception

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewContextWithAuth(t *testing.T) {
	ctx := NewContextWithAuth(context.Background(), "TEST_AUTH")

	v := ctx.Value(contextKeyAuth)
	assert.Equal(t, "TEST_AUTH", v)
}

func TestSetGinContextWithAuth(t *testing.T) {
	c := new(gin.Context)

	SetGinContextWithAuth(c, "TEST_AUTH")

	v, exists := c.Get(contextKeyAuth)
	assert.Equal(t, "TEST_AUTH", v)
	assert.True(t, exists)

	auth, err := AuthFromContext(c)
	assert.Equal(t, "TEST_AUTH", auth)
	assert.NoError(t, err)
}

func TestAuthFromContext(t *testing.T) {
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
				ctx: NewContextWithAuth(context.Background(), "TEST_AUTH"),
			},
			want: "TEST_AUTH",
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
				ctx: context.WithValue(context.Background(), contextKeyAuth, 12450),
			},
			wantErr: ErrUnexpectedType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AuthFromContext(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
