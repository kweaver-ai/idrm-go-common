package impl

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
)

func TestGetUser(t *testing.T) {
	// testdata
	var (
		ctx         = context.Background()
		baseURL     = loadRequiredEnv(t, "TEST_BASE_URL")
		bearerToken = loadRequiredEnv(t, "TEST_BEARER_TOKEN")
		userID      = loadRequiredEnv(t, "TEST_USER_ID")
	)

	ctx = interception.NewContextWithBearerToken(ctx, bearerToken)

	d := &ConfigurationCenterDriven{baseURL: baseURL, client: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}}

	u, err := d.GetUser(ctx, userID, configuration_center.GetUserOptions{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("user: %#v", u)
}
