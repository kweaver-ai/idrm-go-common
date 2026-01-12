package v1

import (
	"os"
	"testing"
)

func loadRequiredEnv(t *testing.T, key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		t.Skipf("required env %q is missing", key)
	}
	return v
}
