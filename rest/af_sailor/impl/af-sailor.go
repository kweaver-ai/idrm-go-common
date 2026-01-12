package impl

import (
	"context"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	driven "github.com/kweaver-ai/idrm-go-common/rest/af_sailor"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewSailorDriven(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.AfSailorHost,
		httpClient: httpClient,
	}
}

func (d *DrivenImpl) QueryRecLabelList(ctx context.Context, req *driven.QueryRecommendLabelReq) (*driven.SailorResp, error) {
	path := "/api/af-sailor/v1/internal/recommend/label"
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.SailorResp](ctx, d.httpClient, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d *DrivenImpl) StandardConsistency(ctx context.Context, tables []*driven.StandardConsistencyTableInfo) ([][]*driven.StandardConsistencyRespRec, error) {
	path := "/api/af-sailor/v1/internal/recommend/check/code"
	url := d.baseURL + path
	//处理参数
	args := &driven.CommonConsistencyReq{
		Query: tables,
	}
	resp, err := base.POST[driven.SailorCommonResp[driven.SailorCommonRecBody[[]*driven.StandardConsistencyRespRec]]](ctx, d.httpClient, url, args)
	if err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errorcode.Desc(errorcode.CallAfSailorError)
	}
	return resp.Data.Rec, nil
}

func (d *DrivenImpl) IndicatorConsistency(ctx context.Context, indicators []*driven.IndicatorConsistency) ([][]*driven.IndicatorConsistencyRespRec, error) {
	path := "/api/af-sailor/v1/internal/recommend/check/indicator"
	url := d.baseURL + path
	//处理参数
	args := &driven.CommonConsistencyReq{
		Query: indicators,
	}
	resp, err := base.POST[driven.SailorCommonResp[driven.SailorCommonRecBody[[]*driven.IndicatorConsistencyRespRec]]](ctx, d.httpClient, url, args)
	if err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errorcode.Desc(errorcode.CallAfSailorError)
	}
	return resp.Data.Rec, nil
}

func (d *DrivenImpl) QueryBusinessSubjectRec(ctx context.Context, recReq []*driven.SailorBusinessSubjectRecReq) (*driven.SailorSubjectRecResp, error) {
	path := "/api/af-sailor/v1/internal/recommend/field/subject"
	url := d.baseURL + path
	args := &driven.SailorSubjectRecReq{
		Query: recReq,
	}
	resp, err := base.POST[driven.SailorCommonResp[[]*driven.SailorBusinessSubjectRecResp]](ctx, d.httpClient, url, args)
	if err != nil {
		return nil, err
	}
	return &driven.SailorSubjectRecResp{SailorBusinessSubjectRecResp: resp.Data}, nil
}
