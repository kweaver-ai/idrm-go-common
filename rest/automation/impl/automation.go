package impl

import (
	"context"
	"fmt"

	"net/http"

	driven "github.com/kweaver-ai/idrm-go-common/rest/automation"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type drivenImpl struct {
	publicBaseURL string
	httpClient    *http.Client
}

func NewDriven(client *http.Client) driven.Driven {
	return &drivenImpl{
		publicBaseURL: base.Service.AutomationPublicHost,
		httpClient:    client,
	}
}

func (d drivenImpl) publicPath(uri string) string {
	return d.publicBaseURL + uri
}

func (d drivenImpl) DagList(ctx context.Context, req *driven.DagListArgs) (*driven.DagListResp, error) {
	uri := "/api/automation/v1/dags"
	return base.GET[*driven.DagListResp](ctx, d.httpClient, d.publicPath(uri), req)
}

func (d drivenImpl) DagDetail(ctx context.Context, id string) (*driven.DagDetailResp, error) {
	uri := fmt.Sprintf("/api/automation/v1/dag/%s", id)
	return base.GET[*driven.DagDetailResp](ctx, d.httpClient, d.publicPath(uri), nil)
}

func (d drivenImpl) RunInstanceForm(ctx context.Context, id string, body map[string]any) error {
	uri := fmt.Sprintf("/api/automation/v1/run-instance-form/%s", id)
	args := driven.RunInstanceFormArgs{
		Data: body,
	}
	_, err := base.POST[any](ctx, d.httpClient, d.publicPath(uri), args)
	return err
}
