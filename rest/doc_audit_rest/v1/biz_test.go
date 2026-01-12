package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/kweaver-ai/idrm-go-common/interception"
)

func Test_biz_Get(t *testing.T) {
	// testdata
	var (
		endpoint    = loadRequiredEnv(t, "TEST_ENDPOINT")
		bearerToken = loadRequiredEnv(t, "TEST_BEARER_TOKEN")
		bizID       = loadRequiredEnv(t, "TEST_BIZ_ID")
	)

	b, err := url.Parse(endpoint)
	if err != nil {
		t.Skip(err)
		return
	}

	c := New(b, &http.Client{Transport: &debuggingRoundTripper{t: t}})

	ctx := interception.NewContextWithBearerToken(context.Background(), bearerToken)

	apply, err := c.DocAudit().Biz().Get(ctx, bizID)
	if err != nil {
		t.Skip(err)
		return
	}

	applyJSON, err := json.MarshalIndent(apply, "", "  ")
	if err != nil {
		t.Skip(err)
		return
	}

	t.Logf("apply: %s", applyJSON)
}
