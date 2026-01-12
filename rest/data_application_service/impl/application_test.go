package impl

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"testing"

	"github.com/kweaver-ai/idrm-go-common/interception"
)

func TestDrivenImpl_Service(t *testing.T) {
	// testdata
	var (
		baseURL     = loadRequiredEnv(t, "TEST_BASE_URL")
		bearerToken = loadRequiredEnv(t, "TEST_BEARER_TOKEN")
		serviceID   = loadRequiredEnv(t, "TEST_SERVICE_ID")
	)

	// context
	ctx := context.Background()
	ctx = interception.NewContextWithBearerToken(ctx, bearerToken)
	// http client
	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	// data-application-service client
	d := &DrivenImpl{baseURL: baseURL, httpClient: c}
	// get service
	s, err := d.Service(ctx, serviceID)
	if err != nil {
		t.Errorf("get service fail: %v", err)
		return
	}

	t.Logf("service: %+v", s)
}

func loadRequiredEnv(t *testing.T, key string) string {
	t.Helper()

	v, ok := os.LookupEnv(key)
	if !ok {
		t.Skipf("required env %q is missing", key)
	}
	return v
}
