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

func TestUserRoleOrRoleGroupBindingBatchProcessing_UnmarshalJSON(t *testing.T) {
	// testdata
	var (
		UserID0      = "11111111-0000-0000-0000-000000000000"
		UserID1      = "11111111-0000-0000-1111-000000000000"
		RoleID0      = "11111111-0000-1111-0000-000000000000"
		RoleID1      = "11111111-0000-1111-1111-000000000000"
		RoleGroupID0 = "11111111-0000-2222-0000-000000000000"
		RoleGroupID1 = "11111111-0000-2222-1111-000000000000"
	)
	tests := []struct {
		name string
		want *UserRoleOrRoleGroupBindingBatchProcessing
	}{
		{
			name: "default",
			want: &UserRoleOrRoleGroupBindingBatchProcessing{
				UserIDs:      []string{UserID0, UserID1},
				RoleIDs:      []string{RoleID0, RoleID1},
				RoleGroupIDs: []string{RoleGroupID0, RoleGroupID1},
				State:        meta_v1.ProcessingStatePresent,
			},
		},
		{
			name: "explicitness",
			want: &UserRoleOrRoleGroupBindingBatchProcessing{
				Bindings: []UserRoleOrRoleGroupBindingProcessing{
					{UserRoleOrRoleGroupBinding: UserRoleOrRoleGroupBinding{UserID: UserID0, RoleID: RoleID0}, State: meta_v1.ProcessingStatePresent},
					{UserRoleOrRoleGroupBinding: UserRoleOrRoleGroupBinding{UserID: UserID0, RoleID: RoleID1}, State: meta_v1.ProcessingStatePresent},
					{UserRoleOrRoleGroupBinding: UserRoleOrRoleGroupBinding{UserID: UserID0, RoleGroupID: RoleGroupID0}, State: meta_v1.ProcessingStatePresent},
					{UserRoleOrRoleGroupBinding: UserRoleOrRoleGroupBinding{UserID: UserID0, RoleGroupID: RoleGroupID1}, State: meta_v1.ProcessingStatePresent},
					{UserRoleOrRoleGroupBinding: UserRoleOrRoleGroupBinding{UserID: UserID1, RoleID: RoleID0}, State: meta_v1.ProcessingStatePresent},
					{UserRoleOrRoleGroupBinding: UserRoleOrRoleGroupBinding{UserID: UserID1, RoleID: RoleID1}, State: meta_v1.ProcessingStatePresent},
					{UserRoleOrRoleGroupBinding: UserRoleOrRoleGroupBinding{UserID: UserID1, RoleGroupID: RoleGroupID0}, State: meta_v1.ProcessingStatePresent},
					{UserRoleOrRoleGroupBinding: UserRoleOrRoleGroupBinding{UserID: UserID1, RoleGroupID: RoleGroupID1}, State: meta_v1.ProcessingStatePresent},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("testdata", "user-role-or-role-group-binding-batch-processes", tt.name+".json"))
			require.NoError(t, err)

			got := &UserRoleOrRoleGroupBindingBatchProcessing{}
			require.NoError(t, json.Unmarshal(data, got))

			assert.Equal(t, tt.want, got)
		})
	}
}
