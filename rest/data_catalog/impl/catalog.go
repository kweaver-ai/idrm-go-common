package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_catalog"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewDrivenImpl(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.DataCatalogHost,
		httpClient: httpClient,
	}
}

func (d *DrivenImpl) GetDataCatalogDetail(ctx context.Context, catalogID string) (*driven.GetDataCatalogDetailResp, error) {
	const path = "/api/internal/data-catalog/v1/data-catalog/:catalog_id"

	args := &driven.QueryCatalogIDReq{
		CatalogID: catalogID,
	}

	resp, err := base.Call[driven.GetDataCatalogDetailResp](ctx, d.httpClient, http.MethodGet, d.baseURL+path, args)
	if err != nil {
		return nil, err
	}
	resp.ID = catalogID
	return &resp, nil
}

func (d *DrivenImpl) GetDataCatalogColumnList(ctx context.Context, catalogID string) (*driven.GetDataCatalogColumnsRes, error) {
	const path = "/api/internal/data-catalog/v1/data-catalog/:catalog_id/column"

	args := &driven.QueryCatalogIDReq{
		CatalogID:  catalogID,
		Limit:      0,
		ReportShow: true,
	}

	resp, err := base.Call[driven.GetDataCatalogColumnsRes](ctx, d.httpClient, http.MethodGet, d.baseURL+path, args)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d *DrivenImpl) GetDataCatalogMountList(ctx context.Context, catalogID string) (*driven.GetDataCatalogMountListRes, error) {
	const path = "/api/internal/data-catalog/v1/data-catalog/:catalog_id/mount"

	args := &driven.QueryCatalogIDReq{
		CatalogID: catalogID,
	}
	resp, err := base.Call[driven.GetDataCatalogMountListRes](ctx, d.httpClient, http.MethodGet, d.baseURL+path, args)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d *DrivenImpl) GetInfoCatalogByStandardForm(ctx context.Context, formID []string) ([]*driven.GetCatalogByStandardFormItem, error) {
	const path = "/api/internal/data-catalog/v1/data-catalog/standard/catalog"

	args := &driven.GetCatalogByStandardForm{
		StandardFormID: formID,
	}
	resp, err := base.GET[[]*driven.GetCatalogByStandardFormItem](ctx, d.httpClient, d.baseURL+path, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DrivenImpl) GetTemplateDetail(ctx context.Context, id string) (*driven.TemplateReq, error) {
	errorMsg := "CatalogDrivenImpl GetTemplateDetail "
	url := d.baseURL + "/api/data-catalog/frontend/v1/data-comprehension/template/detail?id=" + id
	res := &driven.TemplateReq{}
	log.Infof(errorMsg+" url:%s \n ", url)
	err := base.CallWithToken(ctx, d.httpClient, errorMsg, http.MethodGet, url, nil, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *DrivenImpl) GetColumnMapByIds(ctx context.Context, req *driven.GetColumnListByIdsReq) (map[uint64]*driven.ColumnNameInfo, error) {
	const path = "/api/internal/data-catalog/v1/data-catalog/column"

	resp, err := base.Call[driven.GetColumnListByIdsResp](ctx, d.httpClient, http.MethodPost, d.baseURL+path, req)
	if err != nil {
		return nil, err
	}
	data := make(map[uint64]*driven.ColumnNameInfo)
	for _, res := range resp.Columns {
		data[res.ID] = res
	}
	return data, nil
}

func (d *DrivenImpl) GetCatalogColumnByViewID(ctx context.Context, id string) ([]*driven.ColumnInfo, error) {
	path := d.baseURL + "/api/internal/data-catalog/v1/data-catalog/resource/:id/column"
	args := &struct {
		ID string `uri:"id"`
	}{ID: id}

	resp, err := base.GET[[]*driven.ColumnInfo](ctx, d.httpClient, path, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DrivenImpl) GetComprehensionDetail(ctx context.Context, catalogId, templateID string) (*driven.ComprehensionDetail, error) {
	errorMsg := "CatalogDrivenImpl GetComprehensionDetail "
	path := d.baseURL + "/api/data-catalog/frontend/v1/data-comprehension/" + catalogId
	if templateID != "" {
		path = path + "?template_id=" + templateID
	}
	res := &driven.ComprehensionDetail{}
	err := base.CallWithToken(ctx, d.httpClient, errorMsg, http.MethodGet, path, nil, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *DrivenImpl) GetResourceFavoriteByID(ctx context.Context, req *driven.CheckV1Req) (map[uint64]*driven.CheckV1Resp, error) {
	const path = "/api/internal/data-catalog/v1/data-catalog/favorite"

	// 将请求参数序列化为 JSON
	reqBody, err := json.Marshal(req)
	log.Infof("CatalogDrivenImpl GetResourceFavoriteByID request: %s \n", reqBody)
	if err != nil {
		return nil, err
	}

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, d.baseURL+path, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	httpReq.Header.Set("Content-Type", "application/json")

	// 发送请求
	httpResp, err := d.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	// 检查响应状态
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %d", httpResp.StatusCode)
	}

	// 解析响应
	var resp driven.CheckV1Resp
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return nil, err
	}

	log.Infof("CatalogDrivenImpl GetResourceFavoriteByID response: %+v \n", resp)

	// 构造 map 返回值
	result := make(map[uint64]*driven.CheckV1Resp)
	result[resp.FavorID] = &resp
	return result, nil
}
