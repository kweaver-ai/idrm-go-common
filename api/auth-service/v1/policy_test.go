package v1

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPolicy_MarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		policy Policy
	}{
		{
			name: "example",
			policy: Policy{
				Subject: Subject{
					Type: SubjectUser,
					ID:   "00000000-0000-0000-0000-000000000000",
				},
				Object: Object{
					Type: ObjectAPI,
					ID:   "00000000-0000-1111-0000-000000000000",
				},
				Action:       ActionRead,
				PolicyEffect: PolicyAllow,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policyJSON, err := os.ReadFile(path.Join("testdata", "policy_"+tt.name+".json"))
			require.NoError(t, err)

			got, err := json.Marshal(&tt.policy)
			require.NoError(t, err)

			assert.JSONEq(t, string(policyJSON), string(got))
		})
	}
}

func TestPolicy_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name   string
		policy Policy
	}{
		{
			name: "example",
			policy: Policy{
				Subject: Subject{
					Type: SubjectUser,
					ID:   "00000000-0000-0000-0000-000000000000",
				},
				Object: Object{
					Type: ObjectAPI,
					ID:   "00000000-0000-1111-0000-000000000000",
				},
				Action:       ActionRead,
				PolicyEffect: PolicyAllow,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			policyJSON, err := os.ReadFile(path.Join("testdata", "policy_"+tt.name+".json"))
			require.NoError(t, err)

			got := &Policy{}
			require.NoError(t, json.Unmarshal(policyJSON, got))
			assert.Equal(t, &tt.policy, got)
		})
	}
}
