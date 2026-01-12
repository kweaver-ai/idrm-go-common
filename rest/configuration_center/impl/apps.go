package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/rest/base"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	driven "github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	af_trace "github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

func (c *ConfigurationCenterDriven) GetAppsByAccountId(ctx context.Context, id string) (*driven.Apps, error) {
	errorMsg := "DrivenConfigurationCenter GetAppsByAccountId "
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/apps/account/%s", c.baseURL, id)
	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	// request.Header.Set("Authorization", ctx.Value(interception.Token).(string))

	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetAppsError, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetAppsError, err.Error())
	}
	var res *driven.Apps
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetAppsError, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetAppsError, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetAppsError)
		}
	}
}

func (c *ConfigurationCenterDriven) GetAppSimpleInfo(ctx context.Context, ids []string) ([]*driven.AppSimpleInfo, error) {
	errorMsg := "DrivenConfigurationCenter GetAppsByAccountIDs"
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/apps/simple", c.baseURL)
	log.Infof(errorMsg+"%s", urlStr)

	args := struct {
		ID string `query:"id"`
	}{
		ID: strings.Join(ids, ","),
	}

	return base.GET[[]*driven.AppSimpleInfo](ctx, c.client, urlStr, args)
}

func (c *ConfigurationCenterDriven) GetAppsByAccountIDs(ctx context.Context, ids []string) ([]*driven.Apps, error) {
	errorMsg := "DrivenConfigurationCenter GetAppsByAccountIDs "
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/apps/accounts/%s", c.baseURL, strings.Join(ids, ","))
	log.Infof(errorMsg+"%s", urlStr)

	return base.GET[[]*driven.Apps](ctx, c.client, urlStr, nil)
}

func (c *ConfigurationCenterDriven) HasAppsAccessPermission(ctx context.Context, appsId string, resource string) (bool, error) {
	var err error
	ctx, span := af_trace.StartInternalSpan(ctx)
	defer func() { af_trace.TelemetrySpanEnd(span, err) }()
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/apps/access-control", c.baseURL)
	query := map[string]string{
		"user_id":  appsId,
		"resource": resource,
	}

	params := make([]string, 0, len(query))
	for k, v := range query {
		params = append(params, k+"="+v)
	}
	if len(params) > 0 {
		urlStr = urlStr + "?" + strings.Join(params, "&")
	}

	fmt.Println(urlStr)
	request, _ := http.NewRequest("GET", urlStr, nil)

	// request.Header.Set("Authorization", ctx.Value(interception.Token).(string))
	resp, err := c.client.Do(request.WithContext(ctx))
	fmt.Println(resp.StatusCode)
	if err != nil {
		log.WithContext(ctx).Error("DrivenConfigurationCenter HasAppsAccessPermission client.Do error", zap.Error(err))
		return false, errorcode.Detail(errorcode.GetAccessPermissionFailure, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error("DrivenConfigurationCenter HasAppsAccessPermission io.ReadAll error", zap.Error(err))
		return false, errorcode.Detail(errorcode.GetAccessPermissionFailure, err)
	}
	var has bool
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &has)
		if err != nil {
			log.WithContext(ctx).Error("DrivenConfigurationCenter HasAppsAccessPermission jsoniter.Unmarshal error", zap.Error(err))
			return false, errorcode.Detail(errorcode.GetAccessPermissionFailure, err)
		}
		fmt.Println(has)
		return has, nil
	}
	return false, nil
}

func (c *ConfigurationCenterDriven) HasAccessPermissionApps(ctx context.Context) (*driven.AppList, error) {
	var err error
	ctx, span := af_trace.StartInternalSpan(ctx)
	defer func() { af_trace.TelemetrySpanEnd(span, err) }()
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/apps?limit=999&offset=1&only_developer=true", c.baseURL)

	fmt.Println(urlStr)
	request, _ := http.NewRequest("GET", urlStr, nil)

	request.Header.Set("Authorization", ctx.Value(interception.Token).(string))
	resp, err := c.client.Do(request.WithContext(ctx))
	fmt.Println(resp.StatusCode)
	if err != nil {
		log.WithContext(ctx).Error("DrivenConfigurationCenter HasAccessPermissionApps client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetAppListFailure, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error("DrivenConfigurationCenter HasAccessPermissionApps io.ReadAll error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetAppListFailure, err)
	}
	var has *driven.AppList
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &has)
		if err != nil {
			log.WithContext(ctx).Error("DrivenConfigurationCenter HasAccessPermissionApps jsoniter.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetAppListFailure, err)
		}
		fmt.Println(has)
		return has, nil
	}
	return nil, errorcode.Detail(errorcode.GetAppListFailure, err)
}

// GetApplication implements configuration_center.Driven.
func (c *ConfigurationCenterDriven) GetApplication(ctx context.Context, id string) (*driven.Apps, error) {
	base, err := url.Parse(c.baseURL)
	if err != nil {
		log.Error("parse configuration-center url fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	// url path
	base.Path = path.Join(base.Path, "/api/configuration-center/v1/apps", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		log.Error("create http request fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	// authorization
	interception.SeAuthorizationIfEmpty(ctx, req.Header)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.DoRequestError, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read http response body fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.ReadResponseError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("get application fail", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))
		return nil, errorcode.NewGetApplicationFailure(id, resp.StatusCode, body)
	}

	result := &driven.Apps{}
	if err := json.Unmarshal(body, result); err != nil {
		log.Error("decode response body fail", zap.Error(err), zap.ByteString("body", body))
		return nil, errorcode.NewUnmarshalResponseError(err, body)
	}

	return result, nil
}

// GetApplication implements configuration_center.Driven.
func (c *ConfigurationCenterDriven) GetApplicationInternal(ctx context.Context, id string) (*driven.Apps, error) {
	base, err := url.Parse(c.baseURL)
	if err != nil {
		log.Error("parse configuration-center url fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	// url path
	base.Path = path.Join(base.Path, "/api/internal/configuration-center/v1/apps", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		log.Error("create http request fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.DoRequestError, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read http response body fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.ReadResponseError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("get application fail", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))
		return nil, errorcode.NewGetApplicationFailure(id, resp.StatusCode, body)
	}

	result := &driven.Apps{}
	if err := json.Unmarshal(body, result); err != nil {
		log.Error("decode response body fail", zap.Error(err), zap.ByteString("body", body))
		return nil, errorcode.NewUnmarshalResponseError(err, body)
	}

	return result, nil
}
