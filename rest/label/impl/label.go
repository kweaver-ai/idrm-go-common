package impl

import (
	"context"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/label"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewBigDataDriven(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.BasicBigDataHost,
		httpClient: httpClient,
	}
}

func (d *DrivenImpl) GetLabelByIds(ctx context.Context, ids []string) (*driven.LabelListResp, error) {
	path := "/api/internal/basic-bigdata-service/v1/label/getByIds"
	url := d.baseURL + path
	//处理参数
	arg := struct {
		ID []string `query:"id"`
	}{
		ID: ids,
	}
	resp, err := base.Call[driven.LabelListResp](ctx, d.httpClient, http.MethodGet, url, arg)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d *DrivenImpl) GetRangeTypeByIds(ctx context.Context, rangeTypeKey string, ids []string) (*driven.LabelListResp, error) {
	path := "/api/internal/basic-bigdata-service/v1/label/getRangeTypeByIds"
	url := d.baseURL + path
	//处理参数
	arg := struct {
		ID           []string `query:"id"`
		RangeTypeKey string   `query:"range_type"`
	}{
		ID:           ids,
		RangeTypeKey: rangeTypeKey,
	}
	resp, err := base.Call[driven.LabelListResp](ctx, d.httpClient, http.MethodGet, url, arg)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
