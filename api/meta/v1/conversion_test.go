package v1

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvert_v1_ListOptions_To_url_Values(t *testing.T) {
	tests := []struct {
		name string
		in   *ListOptions
		want *url.Values
	}{
		{
			name: "empty",
			in:   &ListOptions{},
			want: &url.Values{},
		},
		{
			name: "all",
			in: &ListOptions{
				Offset:    2,
				Limit:     1024,
				Sort:      "created_at",
				Direction: Descending,
			},
			want: &url.Values{
				"offset":    []string{"2"},
				"limit":     []string{"1024"},
				"sort":      []string{"created_at"},
				"direction": []string{"desc"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &url.Values{}
			require.NoError(t, Convert_v1_ListOptions_To_url_Values(tt.in, got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert_url_Values_To_v1_ListOptions(t *testing.T) {
	tests := []struct {
		name string
		in   *url.Values
		want *ListOptions
	}{
		{
			name: "empty",
			in:   &url.Values{},
			want: &ListOptions{},
		},
		{
			name: "all",
			in: &url.Values{
				"offset":    []string{"2"},
				"limit":     []string{"1024"},
				"sort":      []string{"created_at"},
				"direction": []string{"desc"},
			},
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
			got := &ListOptions{}
			assert.NoError(t, Convert_url_Values_To_v1_ListOptions(tt.in, got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert_Slice_string_To_string(t *testing.T) {
	type args struct {
		in  *[]string
		out *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert_Slice_string_To_string(tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Convert_Slice_string_To_string() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConvert_Slice_string_To_int(t *testing.T) {
	type args struct {
		in  *[]string
		out *int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert_Slice_string_To_int(tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Convert_Slice_string_To_int() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConvert_Slice_string_To_bool(t *testing.T) {
	type args struct {
		in  *[]string
		out *bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert_Slice_string_To_bool(tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Convert_Slice_string_To_bool() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
