package impl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-common/util/validation"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func (c *ConfigurationCenterDriven) GetDictItemType(ctx context.Context, queryType string, dictTypes ...string) (*driven.GetDictItemTypeResp, error) {
	url := c.baseURL + "/api/internal/configuration-center/v1/dict/get-dict-item-type"
	args := driven.GetDictItemTypeReq{
		DictType:  dictTypes,
		QueryType: queryType,
	}
	//处理参数
	resp, err := base.Call[driven.GetDictItemTypeResp](ctx, c.client, http.MethodGet, url, args)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *ConfigurationCenterDriven) BatchCheckNotExistTypeKey(ctx context.Context, req driven.CheckDictTypeKeyReq) (*[]string, error) {
	url := c.baseURL + "/api/internal/configuration-center/v1/dict/batch-check-type-key"
	items, err := base.Call[[]string](ctx, c.client, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	resp := validation.FindDifference(req.DictTypeKey, items)
	return &resp, err
}

func (c *ConfigurationCenterDriven) GetDictItemPage(ctx context.Context, req *driven.GetDictItemPageReq) (*driven.GetDictItemPageRes, error) {
	errorMsg := "ConfigurationCenterDriven GetDictItemPage "
	url := c.baseURL + "/api/configuration-center/v1/dict/dict-item-page"
	url = fmt.Sprintf("%s?dict_id=%s&limit=%d&name=%s&offset=%d", url, req.DictId, req.Limit, req.Name, req.Offset)
	res := &driven.GetDictItemPageRes{}
	log.Infof(errorMsg+" url:%s \n ", url)
	err := base.CallWithToken(ctx, c.client, errorMsg, http.MethodGet, url, nil, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *ConfigurationCenterDriven) GetGradeLabel(ctx context.Context, req *driven.GetGradeLabelReq) (*driven.GetGradeLabelRes, error) {
	errorMsg := "ConfigurationCenterDriven GetGradeLabel "
	url := c.baseURL + "/api/internal/configuration-center/v1/grade-label"
	res := &driven.GetGradeLabelRes{}
	log.Infof(errorMsg+" url:%s \n ", url)
	err := base.CallInternal(ctx, c.client, errorMsg, http.MethodGet, url, nil, res)
	if err != nil {
		return res, err
	}
	return res, nil
}
