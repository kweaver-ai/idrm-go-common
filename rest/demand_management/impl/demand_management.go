package impl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/demand_management"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

// NewDataViewDriven  初次迁移过来
func NewDemandManagementDriven(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.DemandManagementHost,
		httpClient: httpClient,
	}
}

func (d *DrivenImpl) GetShareApply(ctx context.Context, req *driven.GetShareApplyReq) (*driven.GetShareApplyResp, error) {
	path := "/api/demand-management/v1/share-apply"

	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.GetShareApplyResp](ctx, d.httpClient, http.MethodGet, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d *DrivenImpl) GetUserShareApplyResource(ctx context.Context, req *driven.UserShareApplyResourceReq) (*driven.UserShareApplyResourceResp, error) {
	path := fmt.Sprintf("/api/demand-management/v1/share-apply/user-resources?type=%d&user_id=%s", req.Type, req.UserId)

	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.UserShareApplyResourceResp](ctx, d.httpClient, http.MethodGet, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// 通过分析场景产物ID查询产物及需求名称
func (d *DrivenImpl) GetNameByAnalOutputItemID(ctx context.Context, analOutputItemID string) (*driven.NameGetResp, error) {
	urlStr := fmt.Sprintf("%s/api/demand-management/v2/data-anal-require/anal-output-item/%s", d.baseURL, analOutputItemID)
	//处理参数
	resp, err := base.Call[driven.NameGetResp](ctx, d.httpClient, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
