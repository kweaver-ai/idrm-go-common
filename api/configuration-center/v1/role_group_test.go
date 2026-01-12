package v1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

func TestRoleGroupRoleBindingBatchProcessing_UnmarshalJSON(t *testing.T) {
	// testdata
	var (
		RoleGroupID0 = "11111111-0000-0000-0000-000000000000"
		RoleGroupID1 = "11111111-0000-0000-1111-000000000000"
		RoleID0      = "11111111-0000-1111-0000-000000000000"
		RoleID1      = "11111111-0000-1111-1111-000000000000"
	)
	tests := []struct {
		name string
		want *RoleGroupRoleBindingBatchProcessing
	}{
		{
			name: "default",
			want: &RoleGroupRoleBindingBatchProcessing{
				RoleGroupIDs: []string{RoleGroupID0, RoleGroupID1},
				RoleIDs:      []string{RoleID0, RoleID1},
				State:        meta_v1.ProcessingStatePresent,
			},
		},
		{
			name: "explicitness",
			want: &RoleGroupRoleBindingBatchProcessing{
				Bindings: []RoleGroupRoleBindingProcessing{
					{RoleGroupRoleBinding: RoleGroupRoleBinding{RoleGroupID: RoleGroupID0, RoleID: RoleID0}, State: meta_v1.ProcessingStatePresent},
					{RoleGroupRoleBinding: RoleGroupRoleBinding{RoleGroupID: RoleGroupID0, RoleID: RoleID1}, State: meta_v1.ProcessingStatePresent},
					{RoleGroupRoleBinding: RoleGroupRoleBinding{RoleGroupID: RoleGroupID1, RoleID: RoleID0}, State: meta_v1.ProcessingStatePresent},
					{RoleGroupRoleBinding: RoleGroupRoleBinding{RoleGroupID: RoleGroupID1, RoleID: RoleID1}, State: meta_v1.ProcessingStatePresent},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("testdata", "role-group-role-binding-batch-processes", tt.name+".json"))
			require.NoError(t, err)

			got := &RoleGroupRoleBindingBatchProcessing{}
			require.NoError(t, json.Unmarshal(data, got))

			assert.Equal(t, tt.want, got)
		})
	}
}
