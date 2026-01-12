package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/virtual_engine"
	"github.com/kweaver-ai/idrm-go-frame/core/errorx/agcodes"
	"github.com/kweaver-ai/idrm-go-frame/core/errorx/agerrors"
	"github.com/kweaver-ai/idrm-go-frame/core/transport/rest/ginx"

	"github.com/kweaver-ai/idrm-go-common/util"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_view"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

// NewDataViewDriven  初次迁移过来
func NewDataViewDriven(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.DataViewHost,
		httpClient: httpClient,
	}
}

func (d *DrivenImpl) DeleteRelated(ctx context.Context, req *driven.DeleteRelatedReq) error {
	errorMsg := "dataViewDriven DeleteRelated"
	log.WithContext(ctx).Infof(errorMsg+" req: %+v", *req)

	if req.Empty() {
		return nil
	}

	url := fmt.Sprintf("%s/api/internal/data-view/v1/subject-domain/logic-view/related", d.baseURL)
	request, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, req.Reader())
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"http.NewRequest", zap.Error(err))
		return errorcode.Detail(errorcode.CascadeDeleteSubjectDomainViewRelatedError, err.Error())
	}
	resp, err := d.httpClient.Do(request)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"client.Do error", zap.Error(err))
		return errorcode.Detail(errorcode.CascadeDeleteSubjectDomainViewRelatedError, err.Error())
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"io.ReadAll", zap.Error(err))
		return errorcode.Detail(errorcode.CascadeDeleteSubjectDomainViewRelatedError, err.Error())
	}

	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		res := new(errorcode.ErrorCodeFullInfo)
		if err = jsoniter.Unmarshal(body, res); err != nil {
			log.WithContext(ctx).Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
			return errorcode.Detail(errorcode.CascadeDeleteSubjectDomainViewRelatedError, err.Error())
		}
		log.WithContext(ctx).Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
		return errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
	} else {
		log.WithContext(ctx).Error(errorMsg+"http status error", zap.String("status", resp.Status))
		return errorcode.Desc(errorcode.CascadeDeleteSubjectDomainViewRelatedError, errors.New("http status error: "+resp.Status))
	}
}

func (d *DrivenImpl) QueryViewCount(ctx context.Context, flag string, isOperator bool, id ...string) (*driven.QueryViewDetailBySubjectIDResp, error) {
	path := "/api/internal/data-view/v1/subject-domain/logical-view/precision"
	args := driven.QueryViewDetailBySubjectIDReq{
		Flag:       flag,
		IsOperator: isOperator,
		ID:         id,
	}
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.QueryViewDetailBySubjectIDResp](ctx, d.httpClient, http.MethodGet, url, args)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d *DrivenImpl) QueryViewCountInMap(ctx context.Context, flag string, isOperator bool, id ...string) (map[string]int64, error) {
	results := make(map[string]int64)
	resp, err := d.QueryViewCount(ctx, flag, isOperator, id...)
	if err != nil {
		return results, err
	}
	for _, v := range resp.RelationNum {
		if v.RelationViewNum <= 0 {
			continue
		}
		results[v.SubjectDomainID] = v.RelationViewNum
	}
	return results, nil
}

func (d *DrivenImpl) QueryViewFieldInfo(ctx context.Context, isOperator bool, id ...string) ([]*driven.SubjectFormViewInfo, error) {
	if len(id) <= 0 {
		return make([]*driven.SubjectFormViewInfo, 0), nil
	}
	path := "/api/internal/data-view/v1/subject-domain/logic-view/fields"
	args := driven.GetRelatedFieldInfoReq{
		IsOperator: isOperator,
		IDs:        strings.Join(id, ","),
	}
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.GetRelatedFieldInfoResp](ctx, d.httpClient, http.MethodGet, url, args)
	if err != nil {
		return nil, err
	}
	return resp.Data, err
}

func (d *DrivenImpl) BatchQueryViewFieldInfo(ctx context.Context, ids ...string) ([]*driven.GetViewFieldsResp, error) {
	if len(ids) <= 0 {
		return make([]*driven.GetViewFieldsResp, 0), nil
	}
	path := "/api/internal/data-view/v1/form-view/fields"
	args := struct {
		ID []string `query:"id"`
	}{
		ID: ids,
	}
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[[]*driven.GetViewFieldsResp](ctx, d.httpClient, http.MethodGet, url, args)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (d *DrivenImpl) GetDataViewDetails(ctx context.Context, id string) (*driven.GetFormViewDetailsRes, error) {
	errorMsg := "dataViewDriven GetDataViewDetails"

	urlStr := fmt.Sprintf("%s/api/internal/data-view/v1/form-view/%s/details", d.baseURL, id)

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)

	resp, err := d.httpClient.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDataViewDetailsError, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDataViewDetailsError, err.Error())
	}
	var res *driven.GetFormViewDetailsRes
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetDataViewDetailsError, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetDataViewDetailsError, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetDataViewDetailsError, errors.New("http status error: "+resp.Status))
		}
	}
}

func (d *DrivenImpl) GetDataViewField(ctx context.Context, id string) (*driven.GetFieldsRes, error) {
	errorMsg := "dataViewDriven GetDataViewField"

	urlStr := fmt.Sprintf("%s/api/data-view/v1/form-view/%s", d.baseURL, id)

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	token, err := util.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", token)

	resp, err := d.httpClient.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDataViewFieldError, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDataViewFieldError, err.Error())
	}
	var res *driven.GetFieldsRes
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetDataViewFieldError, err.Error())
		}
		res.ID = id
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetDataViewFieldError, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetDataViewFieldError, errors.New("http status error: "+resp.Status))
		}
	}
}

func (d *DrivenImpl) GetDataViewFieldByInternal(ctx context.Context, id string) (*driven.GetFieldsRes, error) {
	errorMsg := "dataViewDriven GetDataViewFieldByInternal"

	urlStr := fmt.Sprintf("%s/api/internal/data-view/v1/form-view/%s", d.baseURL, id)

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	//token, err := util.GetToken(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//request.Header.Set("Authorization", token)

	resp, err := d.httpClient.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDataViewFieldError, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDataViewFieldError, err.Error())
	}
	var res *driven.GetFieldsRes
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetDataViewFieldError, err.Error())
		}
		res.ID = id
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetDataViewFieldError, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetDataViewFieldError, errors.New("http status error: "+resp.Status))
		}
	}
}

// GetSubViewIDs implements data_view.Driven.
func (d *DrivenImpl) GetSubViewIDs(ctx context.Context, opts *driven.GetSubViewIDsOptions) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/internal/data-view/v1/sub-view-ids", d.baseURL), http.NoBody)
	if err != nil {
		log.Error("create http request fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read http response body fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("get sub view ids fail", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.UsersRolesFailure, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	var result []string
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("decode response body fail", zap.Error(err), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	return result, nil
}

// GetSubViewLogicViewID implements data_view.Driven.
func (d *DrivenImpl) GetSubViewLogicViewID(ctx context.Context, id string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/internal/data-view/v1/sub-views/%s/logic_view_id", d.baseURL, id), http.NoBody)
	if err != nil {
		log.Error("create http request fail", zap.Error(err))
		return "", errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read http response body fail", zap.Error(err))
		return "", errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("get sub view's logic view id fail", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))
		return "", errorcode.Detail(errorcode.UsersRolesFailure, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	var result string
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("decode response body fail", zap.Error(err), zap.ByteString("body", body))
		return "", errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	return result, nil
}
func (d *DrivenImpl) GetLogicViewReportInfo(ctx context.Context, req *driven.GetLogicViewReportInfoBody) (*driven.GetLogicViewReportInfoRes, error) {
	errorMsg := "dataViewDriven GetLogicViewReportInfo "
	urlStr := d.baseURL + "/api/internal/data-view/v1/logic-view/report-info"
	res := &driven.GetLogicViewReportInfoRes{}
	u, err := url.Parse(urlStr)
	u.RawQuery = url.Values{
		"field_id": req.FieldIds,
	}.Encode()
	log.Infof(errorMsg+" url:%s \n ", u.String())
	err = base.CallInternal(ctx, d.httpClient, errorMsg, http.MethodPost, u.String(), req, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *DrivenImpl) GetDataViewBasic(ctx context.Context, ids []string) ([]*driven.ViewBasicInfo, error) {
	urlStr := d.baseURL + "/api/internal/data-view/v1/logic-view/simple"
	args := struct {
		ID []string `query:"id"`
	}{
		ID: ids,
	}
	resp, err := base.GET[[]*driven.ViewBasicInfo](ctx, d.httpClient, urlStr, args)
	return resp, err
}

func (d *DrivenImpl) GetViewByName(ctx context.Context, name string, datasourceID string) (*driven.GetViewFieldsResp, error) {
	urlStr := d.baseURL + "/api/internal/data-view/v1/form-view/simple"
	args := struct {
		Name         string `query:"name"`
		DatasourceID string `query:"datasource_id"`
	}{
		Name:         name,
		DatasourceID: datasourceID,
	}
	resp, err := base.GET[*driven.GetViewFieldsResp](ctx, d.httpClient, urlStr, args)
	return resp, err
}

func (d *DrivenImpl) GetTableCount(ctx context.Context, departmentID string) (int64, error) {
	urlStr := d.baseURL + "/api/internal/data-view/v1/form-view/count"
	args := struct {
		ID string `query:"department_id"` // 将id作为查询参数传递
	}{
		ID: departmentID,
	}
	resp, err := base.GET[int64](ctx, d.httpClient, urlStr, args)
	log.Infof("GetTableCount id:%s, count:%d", departmentID, resp)
	return resp, err
}

func (d *DrivenImpl) GetViewInfo(ctx context.Context, req *driven.GetViewInfoReq) (*driven.GetViewInfoResp, error) {
	urlStr := d.baseURL + "/api/data-view/v1/form-view/basic"
	params := make([]string, 0, len(req.IDs))
	for _, v := range req.IDs {
		params = append(params, "ids="+v)
	}
	if len(params) > 0 {
		urlStr = urlStr + "?" + strings.Join(params, "&")
	} else {
		return &driven.GetViewInfoResp{Entries: []*driven.ViewInfo{}}, nil
	}
	resp, err := base.Call[driven.GetViewInfoResp](ctx, d.httpClient, http.MethodGet, urlStr, req)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (d *DrivenImpl) GetDataPreview(ctx context.Context, req *driven.DataPreviewReq) (*driven.DataPreviewResp, error) {
	errorMsg := "dataViewDriven GetDataPreview "

	urlStr := fmt.Sprintf("%s/api/data-view/v1/form-view/GetTableCount", d.baseURL)

	log.Infof(errorMsg+" url:%s \n req : %v", urlStr, req)
	d.httpClient.Timeout = 50 * time.Second
	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodPost, urlStr, d.httpClient, req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDataViewDetailsError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}

	res := &driven.DataPreviewResp{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.UnmarshalResponseError, err.Error())
	}
	return res, nil
}

func (d *DrivenImpl) GetViewBasicInfoByName(ctx context.Context, req *driven.GetViewListByTechnicalNameInMultiDatasourceReq) (*driven.GetViewListByTechnicalNameInMultiDatasourceRes, error) {
	errorMsg := "dataViewDriven GetViewBasicInfoByName "
	urlStr := d.baseURL + driven.GetViewListByTechnicalNameInMultiDatasourceUrl
	res := &driven.GetViewListByTechnicalNameInMultiDatasourceRes{}
	log.Infof(errorMsg+" url:%s \n ", urlStr)
	err := base.CallInternal(ctx, d.httpClient, errorMsg, http.MethodPost, urlStr, req, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *DrivenImpl) GetViewField(ctx context.Context, id string) (*driven.GetFieldsRes, error) {
	errorMsg := "dataViewDriven GetViewField"
	urlStr := fmt.Sprintf("%s/api/internal/data-view/v1/form-view/:id", d.baseURL)
	args := struct {
		ID string `uri:"id"`
	}{
		ID: id,
	}
	log.Infof(errorMsg+" url:%s \n ", urlStr)
	resp, err := base.GET[*driven.GetFieldsRes](ctx, d.httpClient, urlStr, args)
	return resp, err
}

func (d *DrivenImpl) UserViewAuth(ctx context.Context, userID string, viewID ...string) ([]string, error) {
	urlStr := fmt.Sprintf("%s/api/internal/data-view/v1/form-view/authed", d.baseURL)
	args := struct {
		UserID string `query:"user_id"`
		ViewID string `query:"view_id"`
	}{
		UserID: userID,
		ViewID: strings.Join(viewID, ","),
	}
	resp, err := base.GET[[]string](ctx, d.httpClient, urlStr, args)
	return resp, err
}

func (d *DrivenImpl) GeSubViewByViews(ctx context.Context, ids []string) (map[string][]string, error) {
	urlStr := d.baseURL + fmt.Sprintf("/api/internal/data-view/v1/sub-view/batch?logic_view_id=%s", strings.Join(ids, ","))
	resp, err := base.GET[map[string][]string](ctx, d.httpClient, urlStr, nil)
	return resp, err
}

func (d *DrivenImpl) GetSyntheticDataCatalog(ctx context.Context, id string) (*virtual_engine.FetchDataRes, error) {
	errorMsg := "dataViewDriven GetSampleData"
	urlStr := fmt.Sprintf("%s/api/internal/data-view/v1/logic-view/%s/synthetic-data", d.baseURL, id)
	log.Infof(errorMsg+" url:%s \n ", urlStr)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"http.NewRequest error", zap.Error(err))
		return nil, err
	}
	token, err := util.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", token)
	resp, err := d.httpClient.Do(request)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+" io.ReadAll error", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		res := &virtual_engine.FetchDataRes{}
		if err = jsoniter.Unmarshal(body, &res); err != nil {
			log.WithContext(ctx).Error(errorMsg+" jsoniter.Unmarshal error", zap.Error(err))
			return nil, err
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			return nil, base.Unmarshal(ctx, body, errorMsg)
		} else {
			log.WithContext(ctx).Error(errorMsg+"http status error", zap.String("status", resp.Status), zap.String("body", string(body)))
			return nil, err
		}
	}
}

func (d *DrivenImpl) GetSampleData(ctx context.Context, id string) (*driven.GetSampleDataRes, error) {
	urlStr := fmt.Sprintf("%s/api/data-view/v1/logic-view/%s/sample-data", d.baseURL, id)
	resp, err := base.Call[driven.GetSampleDataRes](ctx, d.httpClient, http.MethodGet, urlStr, struct{}{})
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d *DrivenImpl) GetByAuditStatus(ctx context.Context, req *driven.GetByAuditStatusReq) (*driven.GetByAuditStatusResp, error) {
	errorMsg := "dataViewDriven GetByAuditStatus "
	urlStr := d.baseURL + "/api/data-view/v1/form-view/by-audit-status"
	args := struct {
		Offset         *int   `query:"offset"`          // 页码，默认1
		Limit          *int   `query:"limit"`           // 每页大小，默认10
		Keyword        string `query:"keyword"`         // 关键字查询，字符无限制
		DatasourceType string `query:"datasource_type"` // 数据源类型
		DatasourceId   string `query:"datasource_id"`   // 数据源id
		PublishStatus  string `query:"publish_status"`  // 发布状态
		IsAudited      *bool  `query:"is_audited"`      // 是否已稽核
	}{
		Offset:         req.Offset,
		Limit:          req.Limit,
		Keyword:        req.Keyword,
		DatasourceType: req.DatasourceType,
		DatasourceId:   req.DatasourceId,
		PublishStatus:  req.PublishStatus,
		IsAudited:      req.IsAudited,
	}
	resp, err := base.Call[driven.GetByAuditStatusResp](ctx, d.httpClient, http.MethodGet, urlStr, args)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"http.NewRequest error", zap.Error(err))
		return nil, err
	}
	return &resp, nil
}

func (d *DrivenImpl) CreateWorkOrderTask(ctx context.Context, req *driven.CreateWorkOrderTaskReq) (*driven.CreateWorkOrderTaskResp, error) {
	errorMsg := "dataViewDriven CreateWorkOrderTask "
	urlStr := d.baseURL + "/api/internal/data-view/v1/explore-task/work-order"
	res := &driven.CreateWorkOrderTaskResp{}
	log.Infof(errorMsg+" url:%s \n ", urlStr)
	err := base.CallInternal(ctx, d.httpClient, errorMsg, http.MethodPost, urlStr, req, res)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"http.NewRequest error", zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (d *DrivenImpl) CreateExploreTask(ctx context.Context, req *driven.CreateExploreTaskReq) (*driven.CreateExploreTaskResp, error) {
	errorMsg := "dataViewDriven CreateExploreTask "
	urlStr := d.baseURL + "/api/internal/data-view/v1/explore-task"
	res := &driven.CreateExploreTaskResp{}
	log.Infof(errorMsg+" url:%s \n ", urlStr)
	err := base.CallInternal(ctx, d.httpClient, errorMsg, http.MethodPost, urlStr, req, res)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"http.NewRequest error", zap.Error(err))
		return nil, err
	}
	return res, nil
}

// GetSubView implements DataViewRepo.
func (d *DrivenImpl) GetSubView(ctx context.Context, id string) (*driven.SubView, error) {
	urlStr := fmt.Sprintf("%s/api/internal/data-view/v1/sub-views/%s", d.baseURL, id)

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, http.NoBody)
	if err != nil {
		log.WithContext(ctx).Error("create http request fail", zap.Error(err), zap.String("method", http.MethodGet), zap.String("url", urlStr))
		return nil, errorcode.Detail(errorcode.PublicInternalError, fmt.Sprintf("create http request fail: %v", err))
	}
	// Set authorization
	if t, err := interception.BearerTokenFromContextCompatible(ctx); err == nil {
		req.Header.Set("Authorization", "Bearer "+t)
	}

	// send http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.WithContext(ctx).Error("send http request fail", zap.Error(err), zap.String("method", http.MethodGet), zap.String("url", urlStr))
		return nil, errorcode.Detail(errorcode.PublicInternalError, fmt.Sprintf("send http request fail: %v", err))
	}
	defer resp.Body.Close()

	// status code other than 200 is considered a failure
	if resp.StatusCode != http.StatusOK {
		var body []byte
		if body, err = io.ReadAll(resp.Body); err != nil {
			log.WithContext(ctx).Error("read response body fail", zap.Error(err))
		}
		log.WithContext(ctx).Error("invoke API GetSubView fail", zap.Error(err), zap.String("method", http.MethodGet), zap.String("url", urlStr), zap.ByteString("response.body", body))

		// 结构化错误
		var errH ginx.HttpError
		if err = json.Unmarshal(body, &errH); err != nil || errH.Code == "" {
			return nil, errorcode.Detail(errorcode.PublicInternalError, map[string]any{
				"method": req.Method,
				"url":    req.URL.String(),
				"status": resp.Status,
				"body":   body,
			})
		}
		return nil, agerrors.NewCode(agcodes.New(errH.Code, errH.Description, errH.Cause, errH.Solution, errH.Detail, ""))
	}

	var result driven.SubView
	// decode response body as json
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.WithContext(ctx).Error("decode response body of API GetSubView fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, fmt.Sprintf("decode response body of API dto.SubView fail: %v", err))
	}

	return &result, nil
}
