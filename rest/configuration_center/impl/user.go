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

	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

// GetUser 返回指定的用户信息
func (c *ConfigurationCenterDriven) GetUser(ctx context.Context, id string, opts configuration_center.GetUserOptions) (*configuration_center.User, error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	log := log.WithContext(ctx)

	// create api endpoint
	base, err := url.Parse(c.baseURL)
	if err != nil {
		span.RecordError(err)
		log.Error("parse configuration-center base url fail", zap.Error(err), zap.String("baseURL", c.baseURL))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "parse configuration-center base url fail"))
	}
	base.Path = path.Join(base.Path, "api/configuration-center/v1/users", id)

	// create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		span.RecordError(err)
		log.Error("create http request fail", zap.Error(err), zap.String("method", http.MethodGet), zap.Stringer("url", base))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "create http request fail"))
	}

	// set header authorization
	if token, err := interception.BearerTokenFromContextCompatible(ctx); err == nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	// send request
	resp, err := c.client.Do(req)
	if err != nil {
		span.RecordError(err)
		log.Error("send http request fail", zap.Error(err), zap.Any("request", req))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "send http request fail"))
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(err)
		log.Error("read http response body fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "read http response body fail"))
	}

	// check response status code
	if resp.StatusCode != http.StatusOK {
		log.Error("call api fail", zap.String("method", req.Method), zap.Stringer("url", req.URL), zap.String("status", resp.Status), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(nil, "call api %s %s fail, status: %s, body: %s", req.Method, req.URL, resp.Status, body))
	}

	// decode response body
	var user configuration_center.User
	if err := json.Unmarshal(body, &user); err != nil {
		span.RecordError(err)
		log.Error("decode response body fail", zap.Error(err), zap.Any("target", user), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "decode response body to %T fail", user))
	}

	return &user, nil
}

func (c *ConfigurationCenterDriven) GetUsers(ctx context.Context, ids []string) ([]*configuration_center.User, error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer span.End()

	log := log.WithContext(ctx)

	// create api endpoint
	base, err := url.Parse(c.baseURL)
	if err != nil {
		span.RecordError(err)
		log.Error("parse configuration-center base url fail", zap.Error(err), zap.String("baseURL", c.baseURL))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "parse configuration-center base url fail"))
	}
	base.Path = path.Join(base.Path, "api/internal/configuration-center/v1/users", strings.Join(ids, ","))

	// create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		span.RecordError(err)
		log.Error("create http request fail", zap.Error(err), zap.String("method", http.MethodGet), zap.Stringer("url", base))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "create http request fail"))
	}

	// send request
	resp, err := c.client.Do(req)
	if err != nil {
		span.RecordError(err)
		log.Error("send http request fail", zap.Error(err), zap.Any("request", req))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "send http request fail"))
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(err)
		log.Error("read http response body fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "read http response body fail"))
	}

	// check response status code
	if resp.StatusCode != http.StatusOK {
		log.Error("call api fail", zap.String("method", req.Method), zap.Stringer("url", req.URL), zap.String("status", resp.Status), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(nil, "call api %s %s fail, status: %s, body: %s", req.Method, req.URL, resp.Status, body))
	}

	// decode response body
	var users []*configuration_center.User
	if err := json.Unmarshal(body, &users); err != nil {
		span.RecordError(err)
		log.Error("decode response body fail", zap.Error(err), zap.Any("target", users), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.PublicInternalError, newDetail(err, "decode response body to %T fail", users))
	}

	return users, nil
}

func (c *ConfigurationCenterDriven) GetBaseUserByIds(ctx context.Context, ids []string) ([]*configuration_center.UserBase, error) {
	urlStr := c.baseURL + "/api/internal/configuration-center/v1/users/batchByIds"
	args := struct {
		IDs []string `query:"ids"`
	}{
		IDs: ids,
	}
	return base.GET[[]*configuration_center.UserBase](ctx, c.client, urlStr, args)
}

type detailError struct {
	Err     string `json:"err,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *detailError) String() string {
	var s string
	if e.Message != "" {
		s += e.Message + ": "
	}
	s += e.Err
	return s
}

func newDetail(err error, msgAndArgs ...any) *detailError {
	de := &detailError{Message: messageFromMsgAndArgs(msgAndArgs)}
	if err != nil {
		de.Err = err.Error()
	}
	return de
}

func messageFromMsgAndArgs(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msgAndArgs)
	}
	return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
}

// GetUserInfo 获取用户信息，带第三方部门ID
// 国开分支还没有该接口
func (c *ConfigurationCenterDriven) GetUserInfo(ctx context.Context, userID string) (*configuration_center.UserRespItem, error) {
	urlStr := fmt.Sprintf("%v%v/%v", c.baseURL, "/api/internal/configuration-center/v1/user", userID)
	return base.GET[*configuration_center.UserRespItem](ctx, c.client, urlStr, nil)
}

func (c *ConfigurationCenterDriven) GetUserInfoSlice(ctx context.Context, userID ...string) (*base.PageResult[configuration_center.UserRespItem], error) {
	urlStr := c.baseURL + "/api/internal/configuration-center/v1/users"
	args := struct {
		UserID []string `query:"user_id"`
	}{
		UserID: userID,
	}
	return base.GET[*base.PageResult[configuration_center.UserRespItem]](ctx, c.client, urlStr, args)
}

func (c *ConfigurationCenterDriven) GetUsersByDeptRoleID(ctx context.Context, deptID, roleID string) ([]*configuration_center.UserRespItem, error) {
	urlStr := fmt.Sprintf("%v%v?role_id=%v&depart_id=%v", c.baseURL, "/api/internal/configuration-center/v1/users/filter", roleID, deptID)
	return base.GET[[]*configuration_center.UserRespItem](ctx, c.client, urlStr, nil)
}
