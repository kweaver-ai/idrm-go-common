package audit

import (
	"context"
	"errors"
)

var (
	ErrNotFound       = errors.New("logger is not found")
	ErrUnexpectedType = errors.New("unexpected value type for context key")
)

// contextKey is how we find Loggers in a context.Context.
type contextKey struct{}

// String returns a string as the context value key for compatibility with
// frameworks such as github.com/gin-gonic/gin.
func (_ contextKey) String() string { return "audit-logger" }

func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// fromStdContext returns a Logger from context.Context or an error if no Logger
// is found.
func fromStdContext(ctx context.Context) (Logger, error) {
	v := ctx.Value(contextKey{})
	if v == nil {
		return Logger{}, ErrNotFound
	}

	l, ok := v.(Logger)
	if !ok {
		// not reached
		return Logger{}, ErrUnexpectedType
	}

	return l, nil
}

// CustomContext is an interface to an abstract non context.Context
// implementation.
type CustomContext interface {
	context.Context
	Set(key any, value any)
}

// SetCustomContext sets the logger in the context.
func SetCustomContext(ctx CustomContext, logger Logger) {
	ctx.Set(contextKey{}.String(), logger)
}

// FromContext returns a Logger from ctx or an error if no Logger is found.
func FromContext(ctx context.Context) (Logger, error) {
	for _, k := range []any{
		contextKey{},
		contextKey{}.String(),
	} {
		v := ctx.Value(k)
		if v == nil {
			continue
		}
		l, ok := v.(Logger)
		if !ok {
			return Logger{}, ErrUnexpectedType
		}
		return l, nil
	}

	return Logger{}, ErrNotFound
}

// FromContextOrDiscard returns a Logger from ctx. If no Logger is found, this
// returns a Logger that discards all audit records.
func FromContextOrDiscard(ctx context.Context) Logger {
	if l, err := FromContext(ctx); err == nil {
		return l
	}

	return Discard()
}
