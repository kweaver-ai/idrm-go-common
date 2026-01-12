package impl

import (
	"context"
	"net/http"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/metadata_manage"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

// NewDrivenImpl 元数据服务的调用
func NewDrivenImpl(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.MetaDataHost,
		httpClient: httpClient,
	}
}

// SendLineage 发送血缘到元数据平台
func (d DrivenImpl) SendLineage(ctx context.Context, data any) error {
	path := "/api/metadata-manage/v1/kafka/produce"
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[base.CommonResponse[string]](ctx, d.httpClient, http.MethodPost, url, data)
	if err != nil {
		return err
	}
	return resp.Error()
}

func (d DrivenImpl) GetLineageRelationData(ctx context.Context, req *driven.QueryLineageReqParams) (*driven.QueryLineageResp, error) {
	path := "/api/metadata-manage/v1/queryService/getLineageRelationData"
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[base.CommonResponse[map[string][]*driven.LineageEdge]](ctx, d.httpClient, http.MethodGet, url, req)
	if err != nil {
		return nil, err
	}
	data := &driven.QueryLineageResp{
		Edges: resp.Data,
	}
	for id := range data.Edges {
		relations := data.Edges[id]
		for _, relation := range relations {
			if len(relation.Parent) > 0 {
				relation.Parents = strings.Split(relation.Parent, ",")
			}
			if len(relation.Child) > 0 {
				relation.Childs = strings.Split(relation.Child, ",")
			}
		}
		data.Edges[id] = relations
	}
	return data, nil
}

func (d DrivenImpl) GetLineageData(ctx context.Context, entityType string, id ...string) (*driven.QueryLineageTableResp, error) {
	path := "/api/metadata-manage/v1/queryService/getLineageData"
	url := d.baseURL + path
	//处理参数
	req := &driven.QueryLineageTableReq{
		EntityType: entityType,
		ID:         strings.Join(id, ","),
	}
	resp, err := base.Call[base.CommonResponse[[]any]](ctx, d.httpClient, http.MethodGet, url, req)
	if err != nil {
		return nil, err
	}
	data := &driven.QueryLineageTableResp{
		Entries: resp.Data,
	}
	return data, nil
}

func (d DrivenImpl) GetLineageColumns(ctx context.Context, tableID string) ([]string, error) {
	path := "/api/metadata-manage/v1/queryService/getLineageColumns/:id"
	url := d.baseURL + path
	//处理参数
	req := &driven.QueryLineageTableColumnReq{
		ID: tableID,
	}
	resp, err := base.Call[base.CommonResponse[string]](ctx, d.httpClient, http.MethodGet, url, req)
	if err != nil {
		return nil, err
	}
	if resp.Data == "" {
		return []string{}, nil
	}
	return strings.Split(resp.Data, ","), nil
}

func (d DrivenImpl) SyncTableInfo(ctx context.Context, req *driven.PayloadTableReq) (*driven.SyncLineageResp, error) {
	path := "/api/metadata-manage/v1/kafka/data-lineage/table-info"
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.SyncLineageResp](ctx, d.httpClient, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d DrivenImpl) SyncColumnDetail(ctx context.Context, req *driven.PayloadColumnReq) (*driven.SyncLineageResp, error) {
	path := "/api/metadata-manage/v1/kafka/data-lineage/column-detail"
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.SyncLineageResp](ctx, d.httpClient, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d DrivenImpl) SyncTaskTableInfo(ctx context.Context, req *driven.PayloadTaskTableReq) (*driven.SyncLineageResp, error) {
	path := "/api/metadata-manage/v1/kafka/data-lineage/task-table-info"
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.SyncLineageResp](ctx, d.httpClient, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d DrivenImpl) SyncTaskColumnDetail(ctx context.Context, req *driven.PayloadTaskColumnReq) (*driven.SyncLineageResp, error) {
	path := "/api/metadata-manage/v1/kafka/data-lineage/task-column-detail"
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.SyncLineageResp](ctx, d.httpClient, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
