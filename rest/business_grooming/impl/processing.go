package impl

import (
	"context"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/business_grooming"
)

func (d DrivenImpl) CheckTableName(ctx context.Context, datasourceID, tableName string) (bool, error) {
	url := d.baseURL + "/api/business-grooming/v1/data-tables/name-check"
	//处理参数
	args := struct {
		Name         string `query:"name"`
		DatasourceID string `query:"datasource_id"`
	}{
		Name:         tableName,
		DatasourceID: datasourceID,
	}
	resp, err := base.GET[driven.TableNameCheckResp](ctx, d.httpClient, url, args)
	return resp.Repeat, err
}

func (d DrivenImpl) CreateSyncTask(ctx context.Context, req *driven.CollectingModelCreateReq) (*base.IDNameResp, error) {
	url := d.baseURL + "/api/business-grooming/v1/collecting/models"
	//处理参数
	resp, err := base.Call[base.IDNameResp](ctx, d.httpClient, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d DrivenImpl) CreateWorkflow(ctx context.Context, req *driven.WorkflowCreateReq) (*base.IDNameResp, error) {
	url := d.baseURL + "/api/business-grooming/v1/workflows"
	//处理参数
	resp, err := base.Call[base.IDNameResp](ctx, d.httpClient, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d DrivenImpl) QueryDataTable(ctx context.Context, datasourceID string, tableName string) ([]*driven.DataTableFieldInfo, error) {
	url := d.baseURL + "/api/business-grooming/v1/business-model/forms/data-tables/:name"
	args := struct {
		Name         string `uri:"name"`
		DatasourceID string `query:"datasource_id"`
	}{
		Name:         tableName,
		DatasourceID: datasourceID,
	}
	return base.Call[[]*driven.DataTableFieldInfo](ctx, d.httpClient, http.MethodGet, url, args)
}

func (d DrivenImpl) QueryFormPathInfo(ctx context.Context, formID string) (*driven.FormPathInfoResp, error) {
	url := d.baseURL + "/api/internal/business-grooming/v1/business-model/form/source"
	args := struct {
		FormId string `query:"form_id"`
	}{
		FormId: formID,
	}
	//处理参数
	resp, err := base.Call[driven.FormPathInfoResp](ctx, d.httpClient, http.MethodGet, url, args)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d DrivenImpl) GetNodeChild(ctx context.Context, nodeID string) ([]*driven.BusinessNodeObject, error) {
	url := d.baseURL + "/api/internal/business-grooming/v1/domain/nodes/children?node_id=" + nodeID
	resp, err := base.Call[[]*driven.BusinessNodeObject](ctx, d.httpClient, http.MethodGet, url, nil)
	return resp, err
}

// BatchQueryMainBusinessModel 批量查询主干业务及其关联的业务模型信息（内部函数）
func (d DrivenImpl) BatchQueryMainBusinessModel(ctx context.Context, mainBusinessIDs []string) ([]driven.MainBusinessModelResp, error) {
	if len(mainBusinessIDs) == 0 {
		return []driven.MainBusinessModelResp{}, nil
	}
	url := d.baseURL + "/api/internal/business-grooming/v1/main-businesses/models"
	args := struct {
		MainBusinessIDs []string `query:"main_business_ids"`
	}{
		MainBusinessIDs: mainBusinessIDs,
	}
	// API返回的是PageResult结构，提取Entries数组并转换为值数组返回
	pageResult, err := base.GET[*base.PageResult[driven.MainBusinessModelResp]](ctx, d.httpClient, url, args)
	if err != nil {
		return nil, err
	}
	// 将指针数组转换为值数组，返回 [{}, {}] 格式
	result := make([]driven.MainBusinessModelResp, 0, len(pageResult.Entries))
	for _, item := range pageResult.Entries {
		if item != nil {
			result = append(result, *item)
		}
	}
	return result, nil
}
