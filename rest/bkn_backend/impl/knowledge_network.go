package impl

import (
	"context"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/bkn_backend"
)

type drivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

type getKnowledgeNetworkDetailArgs struct {
	KNID              string `uri:"kn_id"`
	IncludeDetail     string `query:"include_detail"`
	IncludeStatistics string `query:"include_statistics"`
}

func NewDriven(httpClient *http.Client) bkn_backend.Driven {
	return &drivenImpl{
		baseURL:    base.Service.BKNBackendHost,
		httpClient: httpClient,
	}
}

// GetDetail GET /api/ontology-manager/in/v1/knowledge-networks/{kn_id}?include_detail=true&mode=export
func (d *drivenImpl) GetDetail(ctx context.Context, knID string) (*bkn_backend.KnowledgeNetworkDetail, error) {
	if knID == "" {
		return nil, errorcode.GetKnowledgeNetworkDetailErr.Detail("knID不能为空")
	}
	path := d.baseURL + "/api/ontology-manager/in/v1/knowledge-networks/:kn_id"
	args := getKnowledgeNetworkDetailArgs{
		KNID:              knID,
		IncludeDetail:     "false",
		IncludeStatistics: "false",
	}
	resp, err := base.GET[*bkn_backend.KnowledgeNetworkDetail](ctx, d.httpClient, path, args)
	if err != nil {
		return nil, errorcode.GetKnowledgeNetworkDetailErr.Detail(err.Error())
	}
	if resp == nil {
		return nil, errorcode.GetKnowledgeNetworkDetailErr.Detail("空响应")
	}
	return resp, nil
}
