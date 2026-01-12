package impl

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"testing"

	"github.com/kweaver-ai/idrm-go-common/interception"
)

func TestConfigurationCenterDriven_GetApplication(t *testing.T) {
	// testdata
	var (
		baseURL       = loadRequiredEnv(t, "TEST_BASE_URL")
		bearerToken   = loadRequiredEnv(t, "TEST_BEARER_TOKEN")
		applicationID = loadRequiredEnv(t, "TEST_APPLICATION_ID")
	)

	// context with bearer token
	ctx := interception.NewContextWithBearerToken(context.Background(), bearerToken)
	// http client skipping verifying tls
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	// configuration-center driven
	driven := &ConfigurationCenterDriven{baseURL: baseURL, client: client}

	application, err := driven.GetApplication(ctx, applicationID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("application name: %s", application.Name)
}

func loadRequiredEnv(t *testing.T, key string) string {
	t.Helper()

	v, ok := os.LookupEnv(key)
	if !ok {
		t.Skipf("required env %q is missing", key)
	}
	return v
}
