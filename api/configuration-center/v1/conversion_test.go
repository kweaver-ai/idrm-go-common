package v1

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

func TestConvert_v1_RoleListOptions_To_url_Values(t *testing.T) {
	tests := []struct {
		name string
		in   *RoleListOptions
		want *url.Values
	}{
		{
			name: "empty",
			in:   &RoleListOptions{},
			want: &url.Values{},
		},
		{
			name: "all",
			in: &RoleListOptions{
				ListOptions: meta_v1.ListOptions{
					Offset:    2,
					Limit:     1024,
					Sort:      "created_at",
					Direction: meta_v1.Descending,
				},
				Keyword:     "KEYWORD",
				Type:        RoleTypeCustom,
				RoleGroupID: "ROLE_GROUP_ID",
				UserIDs:     []string{"USER_ID_0", "USER_ID_1"},
			},
			want: &url.Values{
				"offset":        []string{"2"},
				"limit":         []string{"1024"},
				"sort":          []string{"created_at"},
				"direction":     []string{"desc"},
				"keyword":       []string{"KEYWORD"},
				"type":          []string{"Custom"},
				"role_group_id": []string{"ROLE_GROUP_ID"},
				"user_ids":      []string{"USER_ID_0,USER_ID_1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &url.Values{}
			require.NoError(t, Convert_V1_RoleListOptions_To_url_Values(tt.in, got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert_v1_RoleGroupListOptions_To_url_Values(t *testing.T) {
	tests := []struct {
		name string
		in   *RoleGroupListOptions
		want *url.Values
	}{
		{
			name: "empty",
			in:   &RoleGroupListOptions{},
			want: &url.Values{},
		},
		{
			name: "all",
			in: &RoleGroupListOptions{
				ListOptions: meta_v1.ListOptions{
					Offset:    2,
					Limit:     1024,
					Sort:      "created_at",
					Direction: meta_v1.Descending,
				},
				Keyword: "KEYWORD",
				UserIDs: []string{"USER_ID_0", "USER_ID_1"},
			},
			want: &url.Values{
				"offset":    []string{"2"},
				"limit":     []string{"1024"},
				"sort":      []string{"created_at"},
				"direction": []string{"desc"},
				"keyword":   []string{"KEYWORD"},
				"user_ids":  []string{"USER_ID_0,USER_ID_1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &url.Values{}
			require.NoError(t, Convert_V1_RoleGroupListOptions_To_url_Values(tt.in, got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert_v1_UserListOptions_To_url_Values(t *testing.T) {
	tests := []struct {
		name string
		in   *UserListOptions
		want *url.Values
	}{
		{
			name: "empty",
			in:   &UserListOptions{},
			want: &url.Values{},
		},
		{
			name: "all",
			in: &UserListOptions{
				ListOptions: meta_v1.ListOptions{
					Offset:    2,
					Limit:     1024,
					Sort:      "created_at",
					Direction: meta_v1.Descending,
				},
				Keyword:      "KEYWORD",
				DepartmentID: "DEPARTMENT_ID",
			},
			want: &url.Values{
				"offset":        []string{"2"},
				"limit":         []string{"1024"},
				"sort":          []string{"created_at"},
				"direction":     []string{"desc"},
				"keyword":       []string{"KEYWORD"},
				"department_id": []string{"DEPARTMENT_ID"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &url.Values{}
			require.NoError(t, Convert_V1_UserListOptions_To_url_Values(tt.in, got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert_url_Values_To_v1_RoleListOptions(t *testing.T) {
	tests := []struct {
		name string
		in   *url.Values
		want *RoleListOptions
	}{
		{
			name: "empty",
			in:   &url.Values{},
			want: &RoleListOptions{},
		},
		{
			name: "all",
			in: &url.Values{
				"offset":        []string{"2"},
				"limit":         []string{"1024"},
				"sort":          []string{"created_at"},
				"direction":     []string{"desc"},
				"keyword":       []string{"KEYWORD"},
				"type":          []string{"Custom"},
				"role_group_id": []string{"ROLE_GROUP_ID"},
				"user_ids":      []string{"USER_ID_0,USER_ID_1"},
			},
			want: &RoleListOptions{
				ListOptions: meta_v1.ListOptions{
					Offset:    2,
					Limit:     1024,
					Sort:      "created_at",
					Direction: meta_v1.Descending,
				},
				Keyword:     "KEYWORD",
				Type:        RoleTypeCustom,
				RoleGroupID: "ROLE_GROUP_ID",
				UserIDs:     []string{"USER_ID_0", "USER_ID_1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &RoleListOptions{}
			assert.NoError(t, Convert_url_Values_To_V1_RoleListOptions(tt.in, got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert_url_Values_To_v1_RoleGroupListOptions(t *testing.T) {
	tests := []struct {
		name string
		in   *url.Values
		want *RoleGroupListOptions
	}{
		{
			name: "empty",
			in:   &url.Values{},
			want: &RoleGroupListOptions{},
		},
		{
			name: "all",
			in: &url.Values{
				"offset":    []string{"2"},
				"limit":     []string{"1024"},
				"sort":      []string{"created_at"},
				"direction": []string{"desc"},
				"keyword":   []string{"KEYWORD"},
				"user_ids":  []string{"USER_ID_0,USER_ID_1"},
			},
			want: &RoleGroupListOptions{
				ListOptions: meta_v1.ListOptions{
					Offset:    2,
					Limit:     1024,
					Sort:      "created_at",
					Direction: meta_v1.Descending,
				},
				Keyword: "KEYWORD",
				UserIDs: []string{"USER_ID_0", "USER_ID_1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &RoleGroupListOptions{}
			assert.NoError(t, Convert_url_Values_To_V1_RoleGroupListOptions(tt.in, got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert_url_Values_To_v1_UserListOptions(t *testing.T) {
	tests := []struct {
		name string
		in   *url.Values
		want *UserListOptions
	}{
		{
			name: "empty",
			in:   &url.Values{},
			want: &UserListOptions{},
		},
		{
			name: "all",
			in: &url.Values{
				"offset":        []string{"2"},
				"limit":         []string{"1024"},
				"sort":          []string{"created_at"},
				"direction":     []string{"desc"},
				"keyword":       []string{"KEYWORD"},
				"department_id": []string{"DEPARTMENT_ID"},
			},
			want: &UserListOptions{
				ListOptions: meta_v1.ListOptions{
					Offset:    2,
					Limit:     1024,
					Sort:      "created_at",
					Direction: meta_v1.Descending,
				},
				Keyword:      "KEYWORD",
				DepartmentID: "DEPARTMENT_ID",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &UserListOptions{}
			assert.NoError(t, Convert_url_Values_To_V1_UserListOptions(tt.in, got))
			assert.Equal(t, tt.want, got)
		})
	}
}
