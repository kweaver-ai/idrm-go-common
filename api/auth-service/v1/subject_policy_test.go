package v1

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/subject_policy.json
var testSubjectPolicyJSON []byte

func TestSubjectPolicy(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		sp := &SubjectPolicy{
			Subject: Subject{
				Type: SubjectUser,
				ID:   "00000000-0000-0000-0000-000000000000",
			},
			Actions: []Action{
				ActionRead,
				ActionDownload,
			},
		}

		got, err := json.Marshal(sp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, testSubjectPolicyJSON, got)
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		want := SubjectPolicy{
			Subject: Subject{
				Type: SubjectUser,
				ID:   "00000000-0000-0000-0000-000000000000",
			},
			Actions: []Action{
				ActionRead,
				ActionDownload,
			},
		}

		var got SubjectPolicy
		if err := json.Unmarshal(testSubjectPolicyJSON, &got); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, want, got)
	})
}
