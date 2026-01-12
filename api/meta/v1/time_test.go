package v1

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

type TimeHolder struct {
	T Time `json:"t"`
}

func TestTimeMarshalYAML(t *testing.T) {
	cases := []struct {
		input  Time
		result string
	}{
		{Time{}, "t: null\n"},
		{Date(1998, time.May, 5, 1, 5, 5, 50, time.FixedZone("test", -4*60*60)), "t: \"1998-05-05T05:05:05Z\"\n"},
		{Date(1998, time.May, 5, 5, 5, 5, 0, time.UTC), "t: \"1998-05-05T05:05:05Z\"\n"},
	}

	for _, c := range cases {
		input := TimeHolder{c.input}
		result, err := yaml.Marshal(&input)
		if err != nil {
			t.Errorf("Failed to marshal input: '%v': %v", input, err)
		}
		if string(result) != c.result {
			t.Errorf("Failed to marshal input: '%v': expected %+v, got %q", input, c.result, string(result))
		}
	}
}

func TestTimeUnmarshalYAML(t *testing.T) {
	cases := []struct {
		input  string
		result Time
	}{
		{"t: null\n", Time{}},
		{"t: 1998-05-05T05:05:05Z\n", Time{Date(1998, time.May, 5, 5, 5, 5, 0, time.UTC).Local()}},
	}

	for _, c := range cases {
		var result TimeHolder
		if err := yaml.Unmarshal([]byte(c.input), &result); err != nil {
			t.Errorf("Failed to unmarshal input '%v': %v", c.input, err)
		}
		if result.T != c.result {
			t.Errorf("Failed to unmarshal input '%v': expected %+v, got %+v", c.input, c.result, result)
		}
	}
}

func TestTimeMarshalJSON(t *testing.T) {
	cases := []struct {
		input  Time
		result string
	}{
		{Time{}, "{\"t\":null}"},
		{Date(1998, time.May, 5, 5, 5, 5, 50, time.UTC), "{\"t\":\"1998-05-05T05:05:05Z\"}"},
		{Date(1998, time.May, 5, 5, 5, 5, 0, time.UTC), "{\"t\":\"1998-05-05T05:05:05Z\"}"},
	}

	for _, c := range cases {
		input := TimeHolder{c.input}
		result, err := json.Marshal(&input)
		if err != nil {
			t.Errorf("Failed to marshal input: '%v': %v", input, err)
		}
		if string(result) != c.result {
			t.Errorf("Failed to marshal input: '%v': expected %+v, got %q", input, c.result, string(result))
		}
	}
}

func TestTimeUnmarshalJSON(t *testing.T) {
	cases := []struct {
		input  string
		result Time
	}{
		{"{\"t\":null}", Time{}},
		{"{\"t\":\"1998-05-05T05:05:05Z\"}", Time{Date(1998, time.May, 5, 5, 5, 5, 0, time.UTC).Local()}},
		{"{\"t\":\"1998-05-05T05:05:05.123456789Z\"}", Time{Date(1998, time.May, 5, 5, 5, 5, 123456789, time.UTC).Local()}},
	}

	for _, c := range cases {
		var result TimeHolder
		if err := json.Unmarshal([]byte(c.input), &result); err != nil {
			t.Errorf("Failed to unmarshal input '%v': %v", c.input, err)
		}
		if result.T != c.result {
			t.Errorf("Failed to unmarshal input '%v': expected %+v, got %+v", c.input, c.result, result)
		}
	}
}

func TestTimeMarshalJSONUnmarshalYAML(t *testing.T) {
	cases := []struct {
		input Time
	}{
		{Time{}},
		{Date(1998, time.May, 5, 5, 5, 5, 50, time.Local).Rfc3339Copy()},
		{Date(1998, time.May, 5, 5, 5, 5, 0, time.Local).Rfc3339Copy()},
	}

	for i, c := range cases {
		input := TimeHolder{c.input}
		jsonMarshalled, err := json.Marshal(&input)
		if err != nil {
			t.Errorf("%d-1: Failed to marshal input: '%v': %v", i, input, err)
		}

		var result TimeHolder
		err = yaml.Unmarshal(jsonMarshalled, &result)
		if err != nil {
			t.Errorf("%d-2: Failed to unmarshal '%+v': %v", i, string(jsonMarshalled), err)
		}

		iN, iO := input.T.Zone()
		oN, oO := result.T.Zone()
		if iN != oN || iO != oO {
			t.Errorf("%d-3: Time zones differ before and after serialization %s:%d %s:%d", i, iN, iO, oN, oO)
		}

		if input.T.UnixNano() != result.T.UnixNano() {
			t.Errorf("%d-4: Failed to marshal input '%#v': got %#v", i, input, result)
		}
	}
}

func TestTimeEqual(t *testing.T) {
	t1 := NewTime(time.Now())
	cases := []struct {
		name   string
		x      *Time
		y      *Time
		result bool
	}{
		{"nil =? nil", nil, nil, true},
		{"!nil =? !nil", &t1, &t1, true},
		{"nil =? !nil", nil, &t1, false},
		{"!nil =? nil", &t1, nil, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := c.x.Equal(c.y)
			if result != c.result {
				t.Errorf("Failed equality test for '%v', '%v': expected %+v, got %+v", c.x, c.y, c.result, result)
			}
		})
	}
}

func TestTimeBefore(t *testing.T) {
	t1 := NewTime(time.Now())
	cases := []struct {
		name string
		x    *Time
		y    *Time
	}{
		{"nil <? nil", nil, nil},
		{"!nil <? !nil", &t1, &t1},
		{"nil <? !nil", nil, &t1},
		{"!nil <? nil", &t1, nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := c.x.Before(c.y)
			if result {
				t.Errorf("Failed equality test for '%v', '%v': expected false, got %+v", c.x, c.y, result)
			}
		})
	}
}

func TestTimeIsZero(t *testing.T) {
	t1 := NewTime(time.Now())
	cases := []struct {
		name   string
		x      *Time
		result bool
	}{
		{"nil =? 0", nil, true},
		{"!nil =? 0", &t1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := c.x.IsZero()
			if result != c.result {
				t.Errorf("Failed equality test for '%v': expected %+v, got %+v", c.x, c.result, result)
			}
		})
	}
}

func TestNewTimestampUnixMilli(t *testing.T) {
	var (
		time_2024_0820_1734 = time.Date(2024, 8, 20, 17, 34, 0, 0, time.Local)
		time_1543_0529_0800 = time.Date(1543, 5, 29, 8, 0, 0, 0, time.Local)
	)
	type args struct {
		t time.Time
	}
	tests := []struct {
		time time.Time
	}{
		{time: time_2024_0820_1734},
		{time: time_1543_0529_0800},
	}
	for _, tt := range tests {
		t.Run(tt.time.Format(time.RFC3339), func(t *testing.T) {
			want := TimestampUnixMilli{Time: tt.time}
			got := NewTimestampUnixMilli(tt.time)
			assert.Equal(t, want, got)
		})
	}
}

func TestTimestampUnixMilli_IsZero(t *testing.T) {
	tests := []struct {
		name   string
		time   *TimestampUnixMilli
		assert assert.BoolAssertionFunc
	}{
		{
			name:   "unset",
			assert: assert.True,
		},
		{
			name:   "zero",
			time:   &TimestampUnixMilli{Time: time.Time{}},
			assert: assert.True,
		},
		{
			name:   "non zero",
			time:   &TimestampUnixMilli{Time: time.Date(1543, 5, 29, 8, 0, 0, 0, time.Local)},
			assert: assert.False,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assert(t, tt.time.IsZero())
		})
	}
}

func TestTimestampUnixMilli_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		data string
		time *TimestampUnixMilli
		want *TimestampUnixMilli
	}{
		{
			data: "null",
			time: &TimestampUnixMilli{Time: time.Date(1543, 5, 29, 8, 0, 0, 0, time.Local)},
			want: &TimestampUnixMilli{},
		},
		{
			data: "1724147261123",
			time: &TimestampUnixMilli{Time: time.Date(1543, 5, 29, 8, 0, 0, 0, time.Local)},
			want: &TimestampUnixMilli{Time: time.Date(2024, 8, 20, 17, 47, 41, 123000000, time.Local)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			if err := json.Unmarshal([]byte(tt.data), tt.time); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, tt.time)
		})
	}
}

func TestTimestampUnixMilli_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		time TimestampUnixMilli
		want []byte
	}{
		{
			name: "null",
			time: TimestampUnixMilli{},
			want: []byte(`null`),
		},
		{
			name: "1724147261123",
			time: TimestampUnixMilli{Time: time.Date(2024, 8, 20, 17, 47, 41, 123000000, time.Local)},
			want: []byte(`1724147261123`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.time.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
