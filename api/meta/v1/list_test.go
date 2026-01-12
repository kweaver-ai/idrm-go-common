package v1

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListOptions_MarshalQuery(t *testing.T) {
	tests := []struct {
		name string
		opts *ListOptions
		want string
	}{
		{
			name: "empty",
			opts: &ListOptions{},
			want: "",
		},
		{
			name: "offset",
			opts: &ListOptions{Offset: 2},
			want: "offset=2",
		},
		{
			name: "limit",
			opts: &ListOptions{Limit: 1024},
			want: "limit=1024",
		},
		{
			name: "sort",
			opts: &ListOptions{Sort: "created_at"},
			want: "sort=created_at",
		},
		{
			name: "ascending",
			opts: &ListOptions{Direction: Ascending},
			want: "direction=asc",
		},
		{
			name: "descending",
			opts: &ListOptions{Direction: Descending},
			want: "direction=desc",
		},
		{
			name: "all",
			opts: &ListOptions{Offset: 2, Limit: 1024, Sort: "created_at", Direction: Descending},
			want: "direction=desc&limit=1024&offset=2&sort=created_at",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, err := tt.opts.MarshalQuery()
			require.NoError(t, err)

			assert.Equal(t, tt.want, q.Encode())
		})
	}
}

func TestListOptions_UnmarshalQuery(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  *ListOptions
	}{
		{
			name:  "example",
			query: "offset=2&limit=1024&sort=created_at&direction=desc",
			want: &ListOptions{
				Offset:    2,
				Limit:     1024,
				Sort:      "created_at",
				Direction: Descending,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := url.ParseQuery(tt.query)
			require.NoError(t, err)

			opts := &ListOptions{}
			require.NoError(t, opts.UnmarshalQuery(data))

			assert.Equal(t, tt.want, opts)
		})
	}
}
