package audit

import (
	"time"

	audit "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
	"github.com/kweaver-ai/idrm-go-common/util/clock"
)

// Logger is an interface to an abstract logging implementation. This is a
// concrete type for performance reasons, but all the real work is passed on to
// a LogSink. Implementations of LogSink should provide their own constructors
// that return Logger, not LogSink.
//
// Normally the sink should be used only indirectly.
type Logger struct {
	sink LogSink

	operator audit.Operator

	clock clock.PassiveClock
}

// New returns a new Logger instance. This is primarily used by libraries
// implementing LogSink, rather than end users. Passing a nil sink will create a
// Logger which discards all log records.
func New(sink LogSink) Logger {
	return Logger{
		sink:  sink,
		clock: clock.RealClock{},
	}
}

// IsZero returns true if this logger is an uninitialized zero value
func (l Logger) IsZero() bool {
	return l.sink == nil
}

// Info logs at INFO level.
func (l Logger) Info(operation audit.Operation, object audit.ResourceObject) {
	l.output(audit.LevelInfo, operation, object)
}

// Warn logs at WARN level.
func (l Logger) Warn(operation audit.Operation, object audit.ResourceObject) {
	l.output(audit.LevelWarn, operation, object)
}

// WithOperator returns a new Logger instance with the specified Operator.
func (l Logger) WithOperator(operator audit.Operator) Logger {
	return Logger{
		sink:     l.sink,
		operator: operator,
		clock:    l.clock,
	}
}

func (l Logger) output(level audit.Level, operation audit.Operation, object audit.ResourceObject) {
	if l.sink == nil {
		return
	}

	l.sink.LogEvent(eventFor(l.clock.Now(), level, l.operator, operation, object))
}

func eventFor(time time.Time, level audit.Level, operator audit.Operator, operation audit.Operation, object audit.ResourceObject) *audit.Event {
	return &audit.Event{
		Timestamp:   time,
		Level:       level,
		Description: generateSimplifiedChineseDescription(&operator, operation, object),
		Operator:    operator,
		Operation:   operation,
		Detail:      object.GetDetail(),
	}
}

// LogSink represents a logging implementation. End-users will generally not
// interact with this type.
type LogSink interface {
	// LogEvent logs an audit event.
	LogEvent(event *audit.Event)
}
