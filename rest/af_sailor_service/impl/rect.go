package impl

import (
	"context"
	"fmt"

	driven "github.com/kweaver-ai/idrm-go-common/rest/af_sailor_service"
)

const (
	httpPathPrefix = "/api/internal/af-sailor-service/v1"
)

func (d DrivenImpl) RecTable(ctx context.Context, req *driven.RecTableReq, userId string) (*driven.RecTableResp, error) {
	url := fmt.Sprintf("%s%s/recommend/table", d.baseURL, httpPathPrefix)
	return httpPostDo[driven.RecTableResp](ctx, d.httpClient, url, req, map[string]string{"userId": userId})
}

func (d DrivenImpl) RecFlow(ctx context.Context, req *driven.RecFlowReq, userId string) (*driven.RecFlowResp, error) {
	url := fmt.Sprintf("%s%s/recommend/flow", d.baseURL, httpPathPrefix)
	return httpPostDo[driven.RecFlowResp](ctx, d.httpClient, url, req, map[string]string{"userId": userId})
}

func (d DrivenImpl) RecCode(ctx context.Context, req *driven.RecCodeReq) (*driven.RecCodeResp, error) {
	url := fmt.Sprintf("%s%s/recommend/code", d.baseURL, httpPathPrefix)
	return httpPostDo[driven.RecCodeResp](ctx, d.httpClient, url, req, nil)
}

func (d DrivenImpl) RecCheckCode(ctx context.Context, req *driven.CheckCodeReq) (*driven.CheckCodeResp, error) {
	url := fmt.Sprintf("%s%s/recommend/check/code", d.baseURL, httpPathPrefix)
	return httpPostDo[driven.CheckCodeResp](ctx, d.httpClient, url, req, nil)
}

func (d DrivenImpl) RecView(ctx context.Context, req *driven.RecViewReq) (*driven.RecViewResp, error) {
	url := fmt.Sprintf("%s%s/recommend/view", d.baseURL, httpPathPrefix)
	return httpPostDo[driven.RecViewResp](ctx, d.httpClient, url, req, nil)
}

func (d DrivenImpl) GraphNeighbors(ctx context.Context, req *driven.GraphNeighborsReq) (*driven.GraphNeighborsResp, error) {
	url := fmt.Sprintf("%s%s/tools/knowledge-network/alg-server/explore/kgs/neighbors", d.baseURL, httpPathPrefix)
	return httpPostDo[driven.GraphNeighborsResp](ctx, d.httpClient, url, req, nil)
}

func (d DrivenImpl) GraphFullText(ctx context.Context, req *driven.GraphFullTextReq) (*driven.GraphFullTextResp, error) {
	url := fmt.Sprintf("%s%s/tools/knowledge-network/alg-server/graph-search/kgs/full-text", d.baseURL, httpPathPrefix)
	return httpPostDo[driven.GraphFullTextResp](ctx, d.httpClient, url, req, nil)
}

func (d DrivenImpl) LogicalViewDataCategorize(ctx context.Context, req *driven.LogicalViewDatacategorizeReq) (*driven.LogicalViewDataCategorizeResp, error) {
	url := fmt.Sprintf("%s%s/logical-view/data-categorize", d.baseURL, httpPathPrefix)
	return httpPostDo[driven.LogicalViewDataCategorizeResp](ctx, d.httpClient, url, req, nil)
}

func (d DrivenImpl) TableCompletionTableInfo(ctx context.Context, req *driven.TableCompletionTableInfoReqBody, authorization string) (*driven.TableCompletionTableInfoResp, error) {
	url := fmt.Sprintf("%s%s/understanding/table/completion/table_info", d.baseURL, httpPathPrefix)

	headers := make(map[string]string)
	headers["Authorization"] = authorization
	return httpPostDo[driven.TableCompletionTableInfoResp](ctx, d.httpClient, url, req, headers)
}

func (d DrivenImpl) TableCompletionAll(ctx context.Context, req *driven.TableCompletionReqBody, authorization string) (*driven.TableCompletionResp, error) {
	url := fmt.Sprintf("%s%s/understanding/table/completion", d.baseURL, httpPathPrefix)

	headers := make(map[string]string)
	headers["Authorization"] = authorization
	return httpPostDo[driven.TableCompletionResp](ctx, d.httpClient, url, req, headers)
}
