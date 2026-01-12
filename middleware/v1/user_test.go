package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
)

func TestNewContextWithUser(t *testing.T) {
	user := &middleware.User{
		ID:       "00000000-0000-0000-0000-000000000000",
		Name:     "example-name",
		UserType: 0,
	}
	ctx := NewContextWithUser(context.Background(), user)

	assert.Equal(t, user, ctx.Value(interception.InfoName))
}

func TestUserFromContext(t *testing.T) {
	// testdata
	var (
		user0 = &middleware.User{}
	)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *middleware.User
		wantErr error
	}{
		{
			name: "ok",
			args: args{ctx: context.WithValue(context.Background(), interception.InfoName, user0)},
			want: user0,
		},
		{
			name:    "not exist",
			args:    args{ctx: context.Background()},
			wantErr: ErrNotExist,
		},
		{
			name:    "unexpected type",
			args:    args{ctx: context.WithValue(context.Background(), interception.InfoName, "user")},
			wantErr: ErrUnexpectedType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserFromContext(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
