package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Time is a wrapper around time.Time which supports correct marshaling to YAML
// and JSON.  Wrappers are provided for many of the factory methods that the
// time package offers.
type Time struct {
	time.Time `protobuf:"-"`
}

// DeepCopyInto creates a deep-copy of the Time value.  The underlying time.Time
// type is effectively immutable in the time API, so it is safe to
// copy-by-assign, despite the presence of (unexported) Pointer fields.
func (t *Time) DeepCopyInto(out *Time) {
	*out = *t
}

// DeepCopy is an deep-copy function , copying the receiver, creating a new
// Time.
func (t *Time) DeepCopy() *Time {
	if t == nil {
		return nil
	}
	out := new(Time)
	t.DeepCopyInto(out)
	return out
}

// NewTime returns a wrapped instance of the provided time
func NewTime(time time.Time) Time {
	return Time{time}
}

// Date returns the Time corresponding to the supplied parameters by wrapping
// time.Date.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{time.Date(year, month, day, hour, min, sec, nsec, loc)}
}

// Now returns the current local time.
func Now() Time {
	return Time{time.Now()}
}

// IsZero returns true if the value is nil or time is zero.
func (t *Time) IsZero() bool {
	if t == nil {
		return true
	}
	return t.Time.IsZero()
}

// Before reports whether the time instant t is before u.
func (t *Time) Before(u *Time) bool {
	if t != nil && u != nil {
		return t.Time.Before(u.Time)
	}
	return false
}

// Equal reports whether the time instant t is equal to u.
func (t *Time) Equal(u *Time) bool {
	if t == nil && u == nil {
		return true
	}
	if t != nil && u != nil {
		return t.Time.Equal(u.Time)
	}
	return false
}

// Unix returns the local time corresponding to the given Unix time by wrapping
// time.Unix.
func Unix(sec int64, nsec int64) Time {
	return Time{time.Unix(sec, nsec)}
}

// UnixMilli returns the local Time corresponding to the given Unix time by
// wrapping time.Unix.
func UnixMilli(msec int64) Time {
	return Time{time.UnixMilli(msec)}
}

// Rfc3339Copy returns a copy of the Time at second-level precision.
func (t Time) Rfc3339Copy() Time {
	copied, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return Time{copied}
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) == 4 && string(b) == "null" {
		t.Time = time.Time{}
		return nil
	}

	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	pt, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	t.Time = pt.Local()
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		// Encode unset/nil objects as JSON's "null".
		return []byte("null"), nil
	}
	buf := make([]byte, 0, len(time.RFC3339)+2)
	buf = append(buf, '"')
	// time cannot contain non escapable JSON characters
	buf = t.UTC().AppendFormat(buf, time.RFC3339)
	buf = append(buf, '"')
	return buf, nil
}

// ToUnstructured implements the value.UnstructuredConverter interface.
func (t Time) ToUnstructured() interface{} {
	if t.IsZero() {
		return nil
	}
	buf := make([]byte, 0, len(time.RFC3339))
	buf = t.UTC().AppendFormat(buf, time.RFC3339)
	return string(buf)
}

// TimestampUnixMilli is a wrapper around time.Time which supports correct
// marshaling to YAML and JSON. Wrappers are provided for many of factory
// methods that the time package offers.
type TimestampUnixMilli struct {
	time.Time
}

// NewTimestampUnixMilli returns a wrapped instance of the provided time
func NewTimestampUnixMilli(t time.Time) TimestampUnixMilli {
	return TimestampUnixMilli{t}
}

// IsZero returns true if the value is nil or time is zero
func (t *TimestampUnixMilli) IsZero() bool {
	if t == nil {
		return true
	}
	return t.Time.IsZero()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TimestampUnixMilli) UnmarshalJSON(b []byte) error {
	if len(b) == 4 && string(b) == "null" {
		t.Time = time.Time{}
		return nil
	}

	var s int64
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t.Time = time.UnixMilli(s).Local()
	return nil
}

// MarshalJSON  implements the json.Marshaler interface.
func (t TimestampUnixMilli) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		// Encode unset/nil objects as JSON's "null".
		return []byte("null"), nil
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d", t.Time.UnixMilli())
	return buf.Bytes(), nil
}
