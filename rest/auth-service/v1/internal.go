package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	authServiceV1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	auth_service "github.com/kweaver-ai/idrm-go-common/rest/auth-service"
	driven "github.com/kweaver-ai/idrm-go-common/rest/auth-service"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-frame/core/logx/zapx"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type AuthServiceInternalV1Client struct {
	base   *url.URL
	client *http.Client
}

func NewInternalForBase(client *http.Client) (auth_service.AuthServiceInternalV1Interface, error) {
	base, err := url.Parse(base.Service.AuthServiceHost)
	if err != nil {
		return nil, err
	}
	base.Path = path.Join(base.Path, "/api/internal/auth-service/v1")

	return &AuthServiceInternalV1Client{base: base, client: client}, nil
}

func (c *AuthServiceInternalV1Client) path(s string) string {
	base := *c.base
	base.Path = path.Join(base.Path, s)
	return base.String()
}

// Enforce  策略验证
func (c *AuthServiceInternalV1Client) Enforce(ctx context.Context, requests []authServiceV1.EnforceRequest) (responses []bool, err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	// endpoint
	base := *c.base
	base.Path = path.Join(base.Path, "enforce")
	// body
	reqBody, err := json.Marshal(requests)
	if err != nil {
		// return nil, errorcode.NewErrorForHTTPResponse()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.DoRequestError, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.NewErrorForHTTPResponse(resp)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorcode.Detail(errorcode.ReadResponseError, err.Error())
	}

	if err = json.Unmarshal(body, &responses); err != nil {
		return nil, errorcode.NewUnmarshalResponseError(err, body)
	}
	return
}

// MenuResourceCheck implements AuthServiceV1Interface
func (c *AuthServiceInternalV1Client) MenuResourceCheck(ctx context.Context, requests *auth_service.MenuResourceCheckRequest) (responses *auth_service.MenuResourceCheckResponse, err error) {
	log := log.WithContext(ctx)

	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	log.Debug("marshal request body", zapx.Any("requests", requests))
	body, err := json.Marshal(requests)
	if err != nil {
		span.RecordError(err)
		return
	}

	// endpoint
	base := *c.base
	base.Path = path.Join(base.Path, "/menu-resource/enforce")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base.String(), bytes.NewReader(body))
	if err != nil {
		span.RecordError(err)
		return
	}
	log.Debug(
		"http request",
		zap.String("method", req.Method),
		zap.Stringer("url", req.URL),
		zap.Any("header", req.Header),
		zap.ByteString("body", body),
	)
	resp, err := c.client.Do(req)
	if err != nil {
		span.RecordError(err)
		return
	}
	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		span.RecordError(err)
		body = []byte(fmt.Sprintf("read http response body fail: %v", err))
	}
	log.Debug(
		"http response",
		zap.String("method", resp.Request.Method),
		zap.Stringer("url", resp.Request.URL),
		zap.String("status", resp.Status),
		zap.Any("header", resp.Header),
		zap.ByteString("body", body),
	)

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("server side error, status: %s, body: %s", resp.Status, body)
		span.RecordError(err)
		return
	}

	log.Debug("unmarshal http response body")
	if err = json.Unmarshal(body, &responses); err != nil {
		span.RecordError(err)
		return
	}
	return
}

// ListPolicies implements AuthServiceInternalV1Interface.
func (c *AuthServiceInternalV1Client) ListPolicies(ctx context.Context, opts *v1.PolicyListOptions) ([]v1.Policy, error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	// endpoint
	base := *c.base
	// path
	base.Path = path.Join(base.Path, "policies")
	// query
	query := base.Query()
	for _, s := range opts.Subjects {
		query.Add("subject", fmt.Sprintf("%s:%s", s.Type, s.ID))
	}
	for _, o := range opts.Objects {
		query.Add("object", fmt.Sprintf("%s:%s", o.Type, o.ID))
	}
	base.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.DoRequestError, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorcode.Detail(errorcode.ReadResponseError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.NewListPoliciesFailure(resp.StatusCode, body)
	}

	var result []v1.Policy
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errorcode.NewUnmarshalResponseError(err, body)
	}

	return result, nil
}

// ListIndicatorDimensionalRules 获取指标维度规则列表
func (c *AuthServiceInternalV1Client) ListIndicatorDimensionalRules(ctx context.Context, opts *v1.IndicatorDimensionalRuleListOptions) (*v1.IndicatorDimensionalRuleList, error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	// endpoint
	base := *c.base
	// path
	base.Path = path.Join(base.Path, "indicator-dimensional-rules")
	// query
	query, err := opts.MarshalQuery()
	if err != nil {
		return nil, err
	}
	base.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.DoRequestError, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorcode.Detail(errorcode.ReadResponseError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.NewListPoliciesFailure(resp.StatusCode, body)
	}

	var result v1.IndicatorDimensionalRuleList
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errorcode.NewUnmarshalResponseError(err, body)
	}

	return &result, nil
}

// RuleEnforce  策略验证
func (c *AuthServiceInternalV1Client) RuleEnforce(ctx context.Context, arg *authServiceV1.RulePolicyEnforce) (effectResp *v1.RulePolicyEnforceEffect, err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	// endpoint
	base := *c.base
	base.Path = path.Join(base.Path, "rule/enforce")
	// body
	reqBody, _ := json.Marshal(arg)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base.String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	token, err := interception.BearerTokenFromContextCompatible(ctx)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.DoRequestError, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.NewErrorForHTTPResponse(resp)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorcode.Detail(errorcode.ReadResponseError, err.Error())
	}

	if err = json.Unmarshal(body, &effectResp); err != nil {
		return nil, errorcode.NewUnmarshalResponseError(err, body)
	}
	return
}

// GetIndicatorRules  查询指标规则
func (c *AuthServiceInternalV1Client) GetIndicatorRules(ctx context.Context, ruleID ...string) (ds []*v1.IndicatorDimensionalRule, err error) {
	args := &struct {
		IndicatorRuleID string `query:"indicator_rule_id"`
	}{
		IndicatorRuleID: strings.Join(ruleID, ","),
	}

	realURL := c.path("indicator-dimensional-rules/indicators")
	ds, err = base.GET[[]*v1.IndicatorDimensionalRule](ctx, c.client, realURL, args)
	return ds, err
}

// GetRulesByIndicators  查询指标规则
func (c *AuthServiceInternalV1Client) GetRulesByIndicators(ctx context.Context, indicators ...string) (ds map[string][]string, err error) {
	realURL := c.path("indicator-dimensional-rules/batch")
	realURL = fmt.Sprintf("%s?indicator_id=%s", realURL, strings.Join(indicators, ","))
	ds, err = base.GET[map[string][]string](ctx, c.client, realURL, nil)
	return ds, err
}

// FilterPolicyHasExpiredObjects 过滤有过期的object
func (c *AuthServiceInternalV1Client) FilterPolicyHasExpiredObjects(ctx context.Context, objectID ...string) (ds []string, err error) {
	args := &struct {
		ObjectID []string `query:"object_id"`
	}{
		ObjectID: objectID,
	}

	realURL := c.path("/objects/policy/expired")
	ds, err = base.GET[[]string](ctx, c.client, realURL, args)
	return ds, err
}

// QueryViewHasDWHDataAuthRequestForm 查询某个视图用户是否有数仓数据申请单
func (c *AuthServiceInternalV1Client) QueryViewHasDWHDataAuthRequestForm(ctx context.Context, uid string, dataViewIDSlice []string) (map[string]int, error) {
	args := &struct {
		Applicant string `query:"applicant"` // 申请人
		DataID    string `query:"data_id"`   //数据ID，支持多个
	}{
		Applicant: uid,
		DataID:    strings.Join(dataViewIDSlice, ","),
	}
	type DWHDataAuthRequestInfoResp struct {
		DataID string `json:"data_id"` //库表ID
	}
	realURL := c.path("dwh-data-auth-request")
	ds, err := base.GET[[]*DWHDataAuthRequestInfoResp](ctx, c.client, realURL, args)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(ds, func(item *DWHDataAuthRequestInfoResp) (string, int) {
		return item.DataID, 1
	}), err
}

func (c *AuthServiceInternalV1Client) MenuResourceActions(ctx context.Context, userID, resourceID string) (*driven.MenuResourceActionsResp, error) {
	args := struct {
		UserID     string `query:"user_id"`
		ResourceID string `query:"resource_id"`
	}{
		UserID:     userID,
		ResourceID: resourceID,
	}
	realURL := c.path("menu-resource/actions")
	resp, err := base.GET[*driven.MenuResourceActionsResp](ctx, c.client, realURL, args)
	return resp, err
}
