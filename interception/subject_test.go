package interception

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
)

func TestNewContextWithAuthSubject(t *testing.T) {
	ctx := context.Background()

	s := &v1.Subject{Type: v1.SubjectUser, ID: "00000000-0000-0000-0000-000000000000"}

	ctx = NewContextWithAuthServiceSubject(ctx, s)

	got := ctx.Value(contextKeyAuthServiceSubject{})
	assert.Same(t, s, got)
}

func TestSetGinContextWithAuthServiceSubject(t *testing.T) {
	c := &gin.Context{}

	s := &v1.Subject{Type: v1.SubjectUser, ID: "00000000-0000-0000-0000-000000000000"}

	SetGinContextWithAuthServiceSubject(c, s)

	assert.Equal(t, s, c.Keys[contextKeyAuthServiceSubjectString], "incorrect: gin.Context.Keys[contextKeyAuthServiceSubject]")
}

func TestAuthServiceSubjectFromContext(t *testing.T) {
	var (
		subject0 = &v1.Subject{Type: v1.SubjectUser, ID: "00000000-0000-0000-0000-000000000000"}
		subject1 = &v1.Subject{Type: v1.SubjectUser, ID: "11111111-1111-1111-1111-111111111111"}
	)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.Subject
		wantErr error
	}{
		{
			name: "context.Context struct",
			args: args{ctx: context.WithValue(context.Background(), contextKeyAuthServiceSubject{}, subject0)},
			want: subject0,
		},
		{
			name: "context.Context string",
			args: args{ctx: context.WithValue(context.Background(), contextKeyAuthServiceSubjectString, subject1)},
			want: subject1,
		},
		{
			name:    "not exist",
			args:    args{ctx: context.Background()},
			wantErr: ErrNotExist,
		},
		{
			name:    "context.Context struct unexpected type",
			args:    args{ctx: context.WithValue(context.Background(), contextKeyAuthServiceSubject{}, "subject")},
			wantErr: ErrUnexpectedType,
		},
		{
			name:    "context.Context string unexpected type",
			args:    args{ctx: context.WithValue(context.Background(), contextKeyAuthServiceSubjectString, "subject")},
			wantErr: ErrUnexpectedType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AuthServiceSubjectFromContext(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
