package impl

import (
	"context"
	"errors"
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

func (d drivenImpl) DagByName(ctx context.Context, name string) (*driven.DagMeta, error) {
	args := driven.DagListArgs{
		Keyword: name,
		Limit:   50,
		Page:    0,
		SortBy:  "updated_at",
		Order:   "desc",
	}
	dags, err := d.DagList(ctx, &args)
	if err != nil {
		return nil, err
	}
	for _, dag := range dags.Dags {
		if dag.Title == name {
			return dag, nil
		}
	}
	return nil, errors.New("流程不存在")
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
