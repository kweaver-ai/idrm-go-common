package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"

	v1 "github.com/kweaver-ai/idrm-go-common/api/data_application_service/v1"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_application_service"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

// NewDrivenImpl  指标管理
func NewDrivenImpl(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.DataApplicationServiceHost,
		httpClient: httpClient,
	}
}

// Service implements data_application_service.Driven.
func (d *DrivenImpl) Service(ctx context.Context, id string) (*v1.Service, error) {
	url := fmt.Sprintf("%s/api/internal/data-application-service/v1/services/%s", d.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	// authorization
	interception.SeAuthorizationIfEmpty(ctx, req.Header)

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.NewGetApplicationFailure(id, resp.StatusCode, body)
	}

	result := &v1.Service{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	return result, nil
}

func (d *DrivenImpl) QueryDomainServices(ctx context.Context, flag string, isOperator bool, id ...string) (*driven.QueryDomainServicesResp, error) {
	path := "/api/data-application-service/internal/v1/stats/subject-relation-count"
	args := driven.QueryDomainServicesArgs{
		Flag:       flag,
		IsOperator: isOperator,
		ID:         id,
	}
	url := d.baseURL + path
	//处理参数
	resp, err := base.Call[driven.QueryDomainServicesResp](ctx, d.httpClient, http.MethodPost, url, args)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d *DrivenImpl) QueryDomainApplicationServiceCountMap(ctx context.Context, flag string, isOperator bool, id ...string) (map[string]int64, error) {
	result := make(map[string]int64)
	resp, err := d.QueryDomainServices(ctx, flag, isOperator, id...)
	if err != nil {
		return result, err
	}
	for _, node := range resp.RelationNum {
		if node.RelationServiceNum <= 0 {
			continue
		}
		result[node.SubjectDomainID] = node.RelationServiceNum
	}
	return result, nil
}

// ServiceOwnerID implements data_application_service.Driven.
func (d *DrivenImpl) ServiceOwnerID(ctx context.Context, id string) (string, error) {
	// return base.Call[string](ctx, d.httpClient, http.MethodGet, fmt.Sprintf("%s/api/data-application-service/internal/v1/services/%s/owner_id", d.baseURL, id), nil)
	url := fmt.Sprintf("%s/api/data-application-service/internal/v1/services/%s/owner_id", d.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		log.Error("create http request fail", zap.Error(err))
		return "", errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return "", errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read http response body fail", zap.Error(err))
		return "", errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("get user roles fail", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))
		return "", errorcode.Detail(errorcode.UsersRolesFailure, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	var ownerID string
	if err := json.Unmarshal(body, &ownerID); err != nil {
		log.Error("decode response body fail", zap.Error(err), zap.ByteString("body", body))
		return "", errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	return ownerID, nil
}
func (d *DrivenImpl) GetServicesDataView(ctx context.Context, serviceId string) (*driven.GetServicesDataViewRes, error) {
	errorMsg := "ApplicationServiceDriven GetServicesDataView "
	url := d.baseURL + "/api/data-application-service/frontend/v1/services/" + serviceId + "/data-view"
	res := &driven.GetServicesDataViewRes{}
	log.Infof(errorMsg+" url:%s \n ", url)
	err := base.CallWithToken(ctx, d.httpClient, errorMsg, http.MethodGet, url, nil, res)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (d *DrivenImpl) GetDataViewServices(ctx context.Context, dataViewId string) (*driven.GetDataViewServicesRes, error) {
	errorMsg := "ApplicationServiceDriven GetDataViewServices "
	url := d.baseURL + "/api/data-application-service/frontend/v1/data-view/" + dataViewId + "/services"
	res := &driven.GetDataViewServicesRes{}
	log.Infof(errorMsg+" url:%s \n ", url)
	err := base.CallWithToken(ctx, d.httpClient, errorMsg, http.MethodGet, url, nil, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *DrivenImpl) InternalGetServiceDetail(ctx context.Context, id string) (*v1.Service, error) {
	errorMsg := "ApplicationServiceDriven GetDataViewServices "
	url := d.baseURL + "/api/internal/data-application-service/v1/services/" + id
	res := &v1.Service{}
	log.Infof(errorMsg+" url:%s \n ", url)
	err := base.CallInternal(ctx, d.httpClient, errorMsg, http.MethodGet, url, nil, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// InternalBatchPublishAndOnline 批量发布、上线接口服务
func (d *DrivenImpl) InternalBatchPublishAndOnline(ctx context.Context, batch *v1.BatchPublishAndOnline) (err error) {
	body, err := json.Marshal(&batch)
	if err != nil {
		return
	}

	url := d.baseURL + "/api/data-application-service/internal/v1/batch/services/publish-and-online"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = base.StatusCodeNotOK("data_application_service.InternalBatchPublishAndOnline", resp.StatusCode, got)
	}
	return
}

func (d *DrivenImpl) UserServiceAuth(ctx context.Context, userID string, serviceID ...string) ([]string, error) {
	urlStr := fmt.Sprintf("%s/api/data-application-service/internal/v1/services/authed", d.baseURL)
	args := struct {
		UserID    string `query:"user_id"`
		ServiceID string `query:"service_id"`
	}{
		UserID:    userID,
		ServiceID: strings.Join(serviceID, ","),
	}
	resp, err := base.GET[[]string](ctx, d.httpClient, urlStr, args)
	return resp, err
}

// InternalGetServicesByIDs 通过接口ID列表获取接口列表
func (d *DrivenImpl) InternalGetServicesByIDs(ctx context.Context, ids []string) (*driven.ArrayResult[v1.Service], error) {
	body, err := json.Marshal(map[string][]string{"ids": ids})
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	url := d.baseURL + "/api/data-application-service/internal/v1/batch/services"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	defer resp.Body.Close()

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, base.StatusCodeNotOK("data_application_service.InternalGetServicesByIDs", resp.StatusCode, got)
	}

	result := &driven.ArrayResult[v1.Service]{}
	if err := json.Unmarshal(got, &result); err != nil {
		return nil, errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(got)})
	}

	return result, nil
}

// InternalSyncServicesToGateway 通过接口ID与网关同步
func (d *DrivenImpl) InternalSyncServicesToGateway(ctx context.Context, id string) (driven.SyncServicesToGatewayRes, error) {
	url := d.baseURL + "/api/data-application-service/internal/v1/services/sync/" + id

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, http.NoBody)
	if err != nil {
		return driven.SyncServicesToGatewayRes{}, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return driven.SyncServicesToGatewayRes{}, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	defer resp.Body.Close()

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		return driven.SyncServicesToGatewayRes{}, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return driven.SyncServicesToGatewayRes{}, base.StatusCodeNotOK("data_application_service.InternalSyncServicesToGateway", resp.StatusCode, got)
	}

	result := driven.SyncServicesToGatewayRes{}
	if err := json.Unmarshal(got, &result); err != nil {
		return driven.SyncServicesToGatewayRes{}, errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(got)})
	}

	return result, nil
}
