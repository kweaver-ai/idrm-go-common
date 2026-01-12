package impl

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/standardization"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"go.uber.org/zap"
)

type StandardizationDriven struct {
	baseURL    string
	httpClient *http.Client
}

func NewDriven(httpClient *http.Client) driven.Driven {
	return &StandardizationDriven{
		baseURL:    base.Service.StandardizationHost,
		httpClient: httpClient,
	}
}

func (s StandardizationDriven) GetDataElementDetailByCode(ctx context.Context, value ...string) ([]*driven.DataResp, error) {
	return s.getDataElementDetail(ctx, driven.CodeType, value...)
}

func (s StandardizationDriven) GetStandardMapByCode(ctx context.Context, value ...string) (map[string]*driven.DataResp, error) {
	resps, err := s.getDataElementDetail(ctx, driven.CodeType, value...)
	if err != nil {
		return nil, err
	}
	data := make(map[string]*driven.DataResp)
	for _, res := range resps {
		data[res.Code] = res
	}
	return data, nil
}

func (s StandardizationDriven) GetDataElementDetailByID(ctx context.Context, value ...string) ([]*driven.DataResp, error) {
	return s.getDataElementDetail(ctx, driven.IDType, value...)
}

func (s StandardizationDriven) GetStandardMapByID(ctx context.Context, value ...string) (map[string]*driven.DataResp, error) {
	resps, err := s.getDataElementDetail(ctx, driven.IDType, value...)
	if err != nil {
		return nil, err
	}
	data := make(map[string]*driven.DataResp)
	for _, res := range resps {
		data[res.ID] = res
	}
	return data, nil
}

// getDataElementDetail typeCode default：1，id匹配     2，code匹配
func (s StandardizationDriven) getDataElementDetail(ctx context.Context, typeCode int, value ...string) ([]*driven.DataResp, error) {
	path := "/api/standardization/v1/dataelement/internal/query/list"
	args := driven.GetDataElementDetailReq{
		IDS: strings.Join(value, ","),
	}
	if typeCode == driven.CodeType {
		args = driven.GetDataElementDetailReq{
			Codes: strings.Join(value, ","),
		}
	}
	url := s.baseURL + path
	//处理参数
	resp, err := base.Call[base.CommonResponse[[]*driven.DataResp]](ctx, s.httpClient, http.MethodGet, url, args)
	if err != nil {
		return nil, err
	}
	return resp.Data, err
}

func (s *StandardizationDriven) GetStandardDict(ctx context.Context, ids []string) (data map[string]driven.DictResp, err error) {
	errorMsg := "DrivenStandardizationRepo GetStandardDict "
	urlStr := fmt.Sprintf("%s/api/standardization/v1/dataelement/dict/internal/queryByIds", s.baseURL)
	bodyReq := map[string][]string{
		"ids": ids,
	}
	log.WithContext(ctx).Infof("errorMsg :%s,urlStr :%s,ids :%s,", errorMsg, urlStr, bodyReq)

	statusCode, body, err := base.DOWithOutToken(ctx, errorMsg, http.MethodPost, urlStr, s.httpClient, bodyReq)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetStandardDictError, err.Error())
	}
	if statusCode != http.StatusOK {
		return nil, base.StatusCodeNotOK(errorMsg, statusCode, body)
	}

	var standardDetail driven.StandardDictResp
	err = jsoniter.Unmarshal(body, &standardDetail)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return data, errorcode.Detail(errorcode.GetStandardDictError, err.Error())
	}
	data = make(map[string]driven.DictResp)
	// standardDetail to map[string]standardizationbackend.DictResp
	for _, preDict := range standardDetail.Data {
		data[preDict.ID] = preDict
	}
	return

}

func (s StandardizationDriven) DeleteStandFile(ctx context.Context, standFileID string) error {
	path := "/api/standardization/v1/std-file/internal/delete/" + standFileID
	url := s.baseURL + path
	//处理参数
	_, err := base.Call[any](ctx, s.httpClient, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *StandardizationDriven) GetStandardRule(ctx context.Context, ids []string) (data map[string]driven.RuleResp, err error) {
	errorMsg := "DrivenStandardizationRepo GetStandardRule "
	urlStr := fmt.Sprintf("%s/api/standardization/v1/rule/internal/queryByIds", s.baseURL)
	bodyReq := map[string][]string{
		"ids": ids,
	}
	log.WithContext(ctx).Infof("errorMsg :%s,urlStr :%s,ids :%s,", errorMsg, urlStr, bodyReq)

	statusCode, body, err := base.DOWithOutToken(ctx, errorMsg, http.MethodPost, urlStr, s.httpClient, bodyReq)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetStandardRuleError, err.Error())
	}
	if statusCode != http.StatusOK {
		return nil, base.StatusCodeNotOK(errorMsg, statusCode, body)
	}

	var standardDetail driven.StandardRuleResp
	err = jsoniter.Unmarshal(body, &standardDetail)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return data, errorcode.Detail(errorcode.GetStandardRuleError, err.Error())
	}
	data = make(map[string]driven.RuleResp)
	// standardDetail to map[string]standardizationbackend.RuleResp
	for _, preRule := range standardDetail.Data {
		data[preRule.ID] = preRule
	}
	return

}

func (s *StandardizationDriven) GetStandardFiles(ctx context.Context, ids ...string) (data map[string]driven.RuleResp, err error) {
	errorMsg := "DrivenStandardizationRepo GetStandardFiles "
	urlStr := fmt.Sprintf("%s/api/standardization/v1/std-file/internal/queryByIds", s.baseURL)
	bodyReq := map[string][]string{
		"ids": ids,
	}
	log.WithContext(ctx).Infof("errorMsg :%s,urlStr :%s,ids :%s,", errorMsg, urlStr, bodyReq)

	statusCode, body, err := base.DOWithOutToken(ctx, errorMsg, http.MethodPost, urlStr, s.httpClient, bodyReq)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetStandardRuleError, err.Error())
	}
	if statusCode != http.StatusOK {
		return nil, base.StatusCodeNotOK(errorMsg, statusCode, body)
	}

	var standardDetail driven.StandardRuleResp
	err = jsoniter.Unmarshal(body, &standardDetail)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return data, errorcode.Detail(errorcode.GetStandardRuleError, err.Error())
	}
	data = make(map[string]driven.RuleResp)
	for _, preRule := range standardDetail.Data {
		data[preRule.ID] = preRule
	}
	return

}

func (s *StandardizationDriven) GetStandardList(ctx context.Context, req driven.GetListReq) (*driven.GetStandardListRes, error) {
	urlStr := fmt.Sprintf("%s/api/standardization/v1/dataelement/internal/list?catalog_id=%s&keyword=%s", s.baseURL, req.CatalogID, req.Keyword)

	log.Infof("StandardizationDriven GetStandardList url:%s \n", urlStr)
	res := &driven.GetStandardListRes{}
	err := base.CallInternal(ctx, s.httpClient, "StandardizationDriven GetStandardList", http.MethodGet, urlStr, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *StandardizationDriven) GetCodeTableList(ctx context.Context, req driven.GetListReq) (*driven.GetCodeTableListRes, error) {
	urlStr := fmt.Sprintf("%s/api/standardization/v1/dataelement/dict/internal/list?catalog_id=%s&keyword=%s", s.baseURL, req.CatalogID, req.Keyword)

	log.Infof("StandardizationDriven GetCodeTableList url:%s \n", urlStr)
	res := &driven.GetCodeTableListRes{}
	err := base.CallInternal(ctx, s.httpClient, "StandardizationDriven GetCodeTableList", http.MethodGet, urlStr, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
