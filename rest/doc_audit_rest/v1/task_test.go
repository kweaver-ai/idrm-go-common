package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	v1 "github.com/kweaver-ai/idrm-go-common/api/doc_audit_rest/v1"
	meta "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
	"github.com/kweaver-ai/idrm-go-common/interception"
)

func Test_tasks_List(t *testing.T) {
	// testdata
	var (
		endpoint    = loadRequiredEnv(t, "TEST_ENDPOINT")
		bearerToken = loadRequiredEnv(t, "TEST_BEARER_TOKEN")
	)

	b, err := url.Parse(endpoint)
	if err != nil {
		t.Skip(err)
		return
	}

	c := New(b, &http.Client{Transport: &debuggingRoundTripper{t: t}})

	ctx := interception.NewContextWithBearerToken(context.Background(), bearerToken)

	opts := &v1.TaskListOptions{
		ListOptions: meta.ListOptions{
			// Offset: 50,
			// Limit: 50,
		},
		Type: []v1.BizType{
			"af-data-catalog-publish",
			"af-data-catalog-online",
			"af-data-catalog-offline",
			"af-data-view-online",
			"af-data-view-offline",
			"af-data-application-publish",
			"af-data-application-change",
			"af-data-application-online",
			"af-data-application-offline",
			"af_demand_analysis_confirm",
			"af-data-permission-request",
		},
	}

	result, err := c.DocAudit().Tasks().List(ctx, opts)
	if err != nil {
		t.Skip(err)
		return
	}

	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("result: %s", resultJSON)
}
