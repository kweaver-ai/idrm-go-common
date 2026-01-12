package v1

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfoSystemSearchFilter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want InfoSystemSearchQuery
	}{
		{
			name: "unspecified",
			data: []byte(`{}`),
		},
		{
			name: "zero length",
			data: []byte(`{"department_ids":[]}`),
			want: InfoSystemSearchQuery{DepartmentIDs: uuid.UUIDs{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got InfoSystemSearchQuery
			require.NoError(t, json.Unmarshal(tt.data, &got))

			assert.Equal(t, tt.want, got)
		})
	}
}
