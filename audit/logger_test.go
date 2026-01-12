package audit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"time"

	audit "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
	clock "github.com/kweaver-ai/idrm-go-common/util/clock/testing"
)

type fakeLogger struct {
	logged bool

	loggedEvent *audit.Event
}

// LogEvent implements LogSink.
func (f *fakeLogger) LogEvent(event *audit.Event) {
	f.logged = true
	f.loggedEvent = event
}

var _ LogSink = &fakeLogger{}

type fakeResourceObject struct {
	Name string `json:"name,omitempty"`

	Detail json.RawMessage
}

// GetName implements v1.ResourceObject.
func (o *fakeResourceObject) GetName() string {
	return o.Name
}

// GetDetail implements v1.ResourceObject.
func (o *fakeResourceObject) GetDetail() json.RawMessage {
	return o.Detail
}

var _ audit.ResourceObject = &fakeResourceObject{}

func TestNew(t *testing.T) {
	sink := &fakeLogger{}
	logger := New(sink)
	assert.Equal(t, sink, logger.sink)
}

func TestLogger_IsZero(t *testing.T) {
	var logger Logger
	// logger without sink
	assert.True(t, logger.IsZero())
	// logger with sink
	logger.sink = &fakeLogger{}
	assert.False(t, logger.IsZero())
}

func TestLogger_Info(t *testing.T) {
	var (
		operator  = audit.Operator{}
		operation = audit.Operation("example-operation")
		object    = &fakeResourceObject{Name: "example-object", Detail: json.RawMessage(`{"hello":"world"}`)}
		sink      = &fakeLogger{}
		now       = time.Now()
	)

	var logger = Logger{
		sink:     sink,
		operator: operator,
		clock:    clock.NewFakePassiveClock(now),
	}

	tests := []struct {
		level audit.Level
		log   func(operation audit.Operation, object audit.ResourceObject)
	}{
		{
			level: audit.LevelInfo,
			log:   logger.Info,
		},
		{
			level: audit.LevelWarn,
			log:   logger.Warn,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.level), func(t *testing.T) {
			want := eventFor(now, tt.level, operator, operation, object)

			// logger without sink, won't panic
			assert.NotPanics(t, func() { tt.log(operation, object) })

			logger.sink = &fakeLogger{}
			tt.log(operation, object)
			assert.True(t, sink.logged)
			assert.Equal(t, want, sink.loggedEvent)
		})
	}
}

func TestLogger_WithOperator(t *testing.T) {
	var (
		oldOperator = audit.Operator{Name: "Old Operator"}
		newOperator = audit.Operator{Name: "New Operator"}
	)

	// logger without sink, won't panic
	assert.NotPanics(t, func() { New(nil).WithOperator(oldOperator) })

	logger := New(nil).WithOperator(oldOperator)
	logger = logger.WithOperator(newOperator)
	assert.Equal(t, newOperator, logger.operator)
}
