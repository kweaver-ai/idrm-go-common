package v1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalEnforceRequest(t *testing.T) {
	want, err := os.ReadFile(filepath.Join("testdata", "enforce-request.json"))
	if err != nil {
		t.Fatal(err)
	}

	request := &EnforceRequest{
		Subject: Subject{
			Type: SubjectUser,
			ID:   "0192d78e-5dfa-7700-a672-e867e94eebc9",
		},
		Object: Object{
			Type: ObjectDataView,
			ID:   "0192d78e-952e-7233-b749-c43da962defd",
		},
		Action: ActionDownload,
	}
	got, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(t, string(want), string(got))
}

func TestUnmarshalEnforceRequest(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("testdata", "enforce-request.json"))
	if err != nil {
		t.Fatal(err)
	}

	got := &EnforceRequest{}
	if err := json.Unmarshal(data, got); err != nil {
		t.Fatal(err)
	}

	want := &EnforceRequest{
		Subject: Subject{
			Type: SubjectUser,
			ID:   "0192d78e-5dfa-7700-a672-e867e94eebc9",
		},
		Object: Object{
			Type: ObjectDataView,
			ID:   "0192d78e-952e-7233-b749-c43da962defd",
		},
		Action: ActionDownload,
	}

	assert.Equal(t, want, got)
}
