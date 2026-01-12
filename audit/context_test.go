package audit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	want := New(&fakeLogger{})

	ctx := NewContext(context.Background(), want)

	v := ctx.Value(contextKey{})

	if !assert.NotNil(t, v) {
		return
	}

	got, ok := v.(Logger)
	if !assert.True(t, ok) {
		return
	}

	assert.Equal(t, want, got)
}

func Test_fromStdContext(t *testing.T) {
	logger := New(&fakeLogger{})

	tests := []struct {
		name    string
		ctx     context.Context
		want    Logger
		wantErr error
	}{
		{
			name: "成功",
			ctx:  context.WithValue(context.Background(), contextKey{}, logger),
			want: logger,
		},
		{
			name:    "未找到",
			ctx:     context.Background(),
			wantErr: ErrNotFound,
		},
		{
			name:    "类型错误",
			ctx:     context.WithValue(context.Background(), contextKey{}, "logger"),
			wantErr: ErrUnexpectedType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fromStdContext(tt.ctx)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

type fakeCustomContext struct {
	context.Context

	Data map[string]any
}

var _ CustomContext = &fakeCustomContext{}

// Set implements CustomContext.
func (f *fakeCustomContext) Set(key string, value any) {
	f.Data[key] = value
}

// Value implements context.Context.
func (f *fakeCustomContext) Value(key any) any {
	if k, ok := key.(string); ok {
		v, ok := f.Data[k]
		if ok {
			return v
		}
	}

	return f.Context.Value(key)
}

func TestSetCustomContext(t *testing.T) {
	logger := New(&fakeLogger{})

	cc := &fakeCustomContext{Data: make(map[string]any)}

	SetCustomContext(cc, logger)

	assert.Equal(t, logger, cc.Data[contextKey{}.String()])
}

func TestFromContext(t *testing.T) {
	logger := New(&fakeLogger{})

	tests := []struct {
		name    string
		ctx     context.Context
		want    Logger
		wantErr error
	}{
		{
			name: "context.Context struct",
			ctx:  context.WithValue(context.Background(), contextKey{}, logger),
			want: logger,
		},
		{
			name:    "context.Context string",
			ctx:     context.WithValue(context.Background(), contextKey{}.String(), "logger"),
			wantErr: ErrUnexpectedType,
		},
		{
			name: "CustomContext",
			ctx:  &fakeCustomContext{Data: map[string]any{contextKey{}.String(): logger}, Context: context.Background()},
			want: logger,
		},
		{
			name:    "context.Context struct unexpected type",
			ctx:     context.WithValue(context.Background(), contextKey{}, "logger"),
			wantErr: ErrUnexpectedType,
		},
		{
			name:    "context.Context string unexpected type",
			ctx:     context.WithValue(context.Background(), contextKey{}.String(), "logger"),
			wantErr: ErrUnexpectedType,
		},
		{
			name:    "CustomContext unexpected type",
			ctx:     &fakeCustomContext{Data: map[string]any{contextKey{}.String(): "logger"}, Context: context.Background()},
			wantErr: ErrUnexpectedType,
		},
		{
			name:    "not found",
			ctx:     context.Background(),
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromContext(tt.ctx)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestFromContextOrDiscard(t *testing.T) {
	logger := New(&fakeLogger{})

	tests := []struct {
		name string
		ctx  context.Context
		want Logger
	}{
		{
			name: "成功",
			ctx:  NewContext(context.Background(), logger),
			want: logger,
		},
		{
			name: "未找到",
			ctx:  context.Background(),
			want: Discard(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromContextOrDiscard(tt.ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}
