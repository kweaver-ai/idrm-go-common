package impl

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/business_grooming"
	drivenLabel "github.com/kweaver-ai/idrm-go-common/rest/label"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewDriven(client *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.BusinessGroomingHost,
		httpClient: client,
	}
}

func (d DrivenImpl) GetBusinessFormSource(ctx context.Context, formID string) (res *driven.BusinessFormSourceRes, err error) {
	url := fmt.Sprintf("%s/api/internal/business-grooming/v1/business-model/form/%s/source", d.baseURL, formID)
	resp, err := base.Call[driven.BusinessFormSourceRes](ctx, d.httpClient, http.MethodGet, url, struct{}{})
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d DrivenImpl) GetBusinessFormDetails(ctx context.Context, formIDs, tableKinds []string, pageNumber, pageSize int) (res []*driven.BusinessFormDetail, err error) {
	url := fmt.Sprintf("%s/api/internal/business-grooming/v1/business-model/forms?form_ids=%s&table_kind=%s&offset=%d&limit=%d", d.baseURL, strings.Join(formIDs, ","), strings.Join(tableKinds, ","), pageNumber, pageSize)
	res, err = base.Call[[]*driven.BusinessFormDetail](ctx, d.httpClient, http.MethodGet, url, struct{}{})
	return
}

func (d DrivenImpl) GetBusinessNodesBrief(ctx context.Context, ids []string) (res []*driven.BusinessNode, err error) {
	url := fmt.Sprintf("%s/api/internal/business-grooming/v1/domain/nodes?id=%s", d.baseURL, strings.Join(ids, ","))
	res, err = base.Call[[]*driven.BusinessNode](ctx, d.httpClient, http.MethodGet, url, struct{}{})
	return
}

func (d DrivenImpl) GetFormViewTableDict(ctx context.Context, businessModelID string, ids []string) (res map[string]string, err error) {
	url := d.baseURL + "/api/internal/business-grooming/v1/business-model/forms/view"
	args := struct {
		BusinessModelID string   `query:"business_model_id"`
		ID              []string `query:"id"`
	}{
		BusinessModelID: businessModelID,
		ID:              ids,
	}
	return base.GET[map[string]string](ctx, d.httpClient, url, args)
}

func (d DrivenImpl) GetBusinessFormDetailsFilterLabel(ctx context.Context, rangeTypeKey string, formIDs, tableKinds []string, pageNumber, pageSize int) (res []*driven.BusinessFormAndLabelDetail, err error) {
	url := fmt.Sprintf("%s/api/internal/business-grooming/v1/business-model/forms?form_ids=%s&table_kind=%s&offset=%d&limit=%d", d.baseURL, strings.Join(formIDs, ","), strings.Join(tableKinds, ","), pageNumber, pageSize)
	formLists, formErr := base.Call[[]*driven.BusinessFormDetail](ctx, d.httpClient, http.MethodGet, url, struct{}{})
	if formErr != nil {
		return nil, err
	}
	detailsLists := make([]*driven.BusinessFormAndLabelDetail, 0)
	for _, detail := range formLists {
		t := &driven.BusinessFormAndLabelDetail{}
		copier.Copy(t, detail)
		if len(detail.LabelIds) <= 0 {
			detailsLists = append(detailsLists, t)
			continue
		}
		log.Infof("===标签ids==%v", detail.LabelIds)
		labelUrl := "http://basic-bigdata-service:8287/api/internal/basic-bigdata-service/v1/label/getRangeTypeByIds"
		//处理参数
		arg := struct {
			ID        []string `query:"id"`
			RangeType string   `query:"range_type"`
		}{
			ID:        detail.LabelIds,
			RangeType: rangeTypeKey,
		}
		labelList, labelErr := base.Call[drivenLabel.LabelListResp](ctx, d.httpClient, http.MethodGet, labelUrl, arg)
		if labelErr != nil {
			log.WithContext(ctx).Error("GetBusinessFormDetailsFilterLabel getRangeTypeByIds err", zap.Error(err))
		} else {
			t.LabelListResp = labelList.LabelResp
		}
		detailsLists = append(detailsLists, t)
	}
	return detailsLists, nil
}
