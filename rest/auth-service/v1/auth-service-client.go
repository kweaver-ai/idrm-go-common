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

	authServiceV1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	metaV1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
	"github.com/kweaver-ai/idrm-go-common/interception"
	auth_service "github.com/kweaver-ai/idrm-go-common/rest/auth-service"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-frame/core/logx/zapx"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"go.uber.org/zap"
)

type AuthServiceV1Client struct {
	base   *url.URL
	client *http.Client
}

// NewForBaseAndClient 根据 GoCommon/rest/base 的配置和指定的 HTTP 客户端创建 AuthServiceV1Interface
func NewForBaseAndClient(client *http.Client) (auth_service.AuthServiceV1Interface, error) {
	return NewForServiceConfigAndClient(base.Service, client)
}

func New(base *url.URL, client *http.Client) *AuthServiceV1Client {
	return &AuthServiceV1Client{
		base:   base,
		client: client,
	}
}

func NewBaseClient(client *http.Client) auth_service.AuthServiceV1Interface {
	realURL, _ := url.Parse(base.Service.AuthServiceHost)
	realURL.Path = path.Join(realURL.Path, "api", "auth-service", "v1")
	return &AuthServiceV1Client{
		base:   realURL,
		client: client,
	}
}

func NewForServiceConfigAndClient(config *base.ServiceConfig, client *http.Client) (*AuthServiceV1Client, error) {
	base, err := url.Parse(config.AuthServiceHost)
	if err != nil {
		return nil, err
	}

	// add api path
	base.Path = path.Join(base.Path, "api", "auth-service", "v1")

	return New(base, client), nil
}

// Enforce implements AuthServiceV1Interface.
func (c *AuthServiceV1Client) Enforce(ctx context.Context, requests []authServiceV1.EnforceRequest) (responses []bool, err error) {
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
	base.Path = path.Join(base.Path, "enforce")

	log.Debug("create http request")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base.String(), bytes.NewReader(body))
	if err != nil {
		span.RecordError(err)
		return
	}

	setHeaderAuthorization(ctx, req)

	log.Debug(
		"http request",
		zap.String("method", req.Method),
		zap.Stringer("url", req.URL),
		zap.Any("header", req.Header),
		zap.ByteString("body", body),
	)

	log.Debug("send http request")
	resp, err := c.client.Do(req)
	if err != nil {
		span.RecordError(err)
		return
	}
	defer resp.Body.Close()

	log.Debug("read http response body")
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

// 获取指定的指标维度规则
func (c *AuthServiceV1Client) GetIndicatorDimensionalRule(ctx context.Context, id string) (result *authServiceV1.IndicatorDimensionalRule, err error) {
	log := log.WithContext(ctx)

	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	// endpoint
	base := *c.base
	base.Path = path.Join(base.Path, "indicator-dimensional-rules", id)

	log.Debug("create http request")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		span.RecordError(err)
		return
	}

	setHeaderAuthorization(ctx, req)

	log.Debug(
		"http request",
		zap.String("method", req.Method),
		zap.Stringer("url", req.URL),
		zap.Any("header", req.Header),
	)

	log.Debug("send http request")
	resp, err := c.client.Do(req)
	if err != nil {
		span.RecordError(err)
		return
	}
	defer resp.Body.Close()

	log.Debug("read http response body")
	var body []byte
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
	if err = json.Unmarshal(body, &result); err != nil {
		span.RecordError(err)
		return
	}

	return
}

// ListSubjectObjects implements AuthServiceV1Interface.
func (c *AuthServiceV1Client) ListSubjectObjects(ctx context.Context, opts *authServiceV1.SubjectObjectsListOptions) (list *metaV1.List[authServiceV1.ObjectWithPermissions], err error) {
	log := log.WithContext(ctx)

	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	// endpoint
	base := *c.base
	base.Path = path.Join(base.Path, "subject", "objects")

	log.Debug("marshal http request query parameters")
	q, err := opts.MarshalQueryParameter()
	if err != nil {
		span.RecordError(err)
		return
	}
	base.RawQuery = q

	log.Debug("create http request")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		span.RecordError(err)
		return
	}

	setHeaderAuthorization(ctx, req)

	log.Debug(
		"http request",
		zap.String("method", req.Method),
		zap.Stringer("url", req.URL),
		zap.Any("header", req.Header),
	)

	log.Debug("send http request")
	resp, err := c.client.Do(req)
	if err != nil {
		span.RecordError(err)
		return
	}
	defer resp.Body.Close()

	log.Debug("read http response body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
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
	list = new(metaV1.List[authServiceV1.ObjectWithPermissions])
	if err = json.Unmarshal(body, list); err != nil {
		span.RecordError(err)
		return
	}

	return
}

func setHeaderAuthorization(ctx context.Context, req *http.Request) {
	if len(req.Header.Get("Authorization")) != 0 {
		return
	}

	if t, err := interception.BearerTokenFromContextCompatible(ctx); err == nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t))
		return
	}

	if auth, err := interception.AuthFromContext(ctx); err == nil {
		req.Header.Set("Authorization", auth)
		return
	}
}
