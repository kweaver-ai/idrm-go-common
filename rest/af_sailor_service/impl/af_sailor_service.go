package impl

import (
	"context"
	"fmt"
	"net/http"

	driven "github.com/kweaver-ai/idrm-go-common/rest/af_sailor_service"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewSailorDriven(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.AfSailorServiceHost,
		httpClient: httpClient,
	}
}

func (d DrivenImpl) UpdateGraph(ctx context.Context, detail *driven.ModelDetail) (*base.IntIDResp, error) {
	urlStr := d.baseURL + "/api/internal/af-sailor-service/v1/knowledge-build/model/graph"
	resp, err := base.POST[base.IntIDResp](ctx, d.httpClient, urlStr, detail)
	return &resp, err
}

func (d DrivenImpl) DeleteGraph(ctx context.Context, graphID int) (*base.IntIDResp, error) {
	urlStr := d.baseURL + fmt.Sprintf("/api/internal/af-sailor-service/v1/knowledge-build/model/graph/%v", graphID)
	resp, err := base.DELETE[base.IntIDResp](ctx, d.httpClient, urlStr, nil)
	return &resp, err
}

func (d DrivenImpl) GraphBuildTask(ctx context.Context, req *driven.GraphBuildTaskReq) (*base.IntIDResp, error) {
	urlStr := d.baseURL + "/api/internal/af-sailor-service/v1/knowledge-build/model/graph/task"
	resp, err := base.POST[base.IntIDResp](ctx, d.httpClient, urlStr, req)
	return &resp, err
}
