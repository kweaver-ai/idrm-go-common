package impl

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/indicator_management"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

// NewDrivenImpl  指标管理
func NewDrivenImpl(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.IndicatorManagement,
		httpClient: httpClient,
	}
}

// GetIndicator implements indicator_management.Driven.
func (d *DrivenImpl) GetIndicator(ctx context.Context, id string) (*driven.Indicator, error) {
	const path = "/api/indicator-management/v1/indicator/:id"

	args := &driven.GetIndicatorArgs{ID: id}

	resp, err := base.Call[driven.Indicator](ctx, d.httpClient, http.MethodGet, d.baseURL+path, args)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d *DrivenImpl) QueryDomainIndicators(ctx context.Context, flag string, id ...string) (*driven.QueryDomainIndicatorsResp, error) {
	path := "/api/internal/indicator-management/v1/indicator"
	args := driven.QueryDomainIndicatorsArgs{
		Flag: flag,
		ID:   id,
	}
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.QueryDomainIndicatorsResp](ctx, d.httpClient, http.MethodGet, url, args)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d *DrivenImpl) QueryDomainIndicatorCountMap(ctx context.Context, flag string, id ...string) (map[string]int64, error) {
	result := make(map[string]int64)
	resp, err := d.QueryDomainIndicators(ctx, flag, id...)
	if err != nil {
		return result, err
	}
	for _, node := range resp.RelationNum {
		if node.RelationCount <= 0 {
			continue
		}
		result[node.SubjectDomainID] = node.RelationCount
	}
	return result, nil
}

func (d *DrivenImpl) UserIndicatorAuth(ctx context.Context, userID string, indicatorID ...string) ([]string, error) {
	urlStr := fmt.Sprintf("%s/api/internal/indicator-management/v1/indicator/authed", d.baseURL)
	args := struct {
		UserID      string `query:"user_id"`
		IndicatorID string `query:"indicator_id"`
	}{
		UserID:      userID,
		IndicatorID: strings.Join(indicatorID, ","),
	}
	resp, err := base.GET[[]string](ctx, d.httpClient, urlStr, args)
	return resp, err
}
