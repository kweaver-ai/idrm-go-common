package base

import (
	"context"
	"io"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	tlog "github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func POST[R any](ctx context.Context, client *http.Client, url string, args any) (R, error) {
	return Call[R](ctx, client, http.MethodPost, url, args)
}

func GET[R any](ctx context.Context, client *http.Client, url string, args any) (R, error) {
	return Call[R](ctx, client, http.MethodGet, url, args)
}

func PUT[R any](ctx context.Context, client *http.Client, url string, args any) (R, error) {
	return Call[R](ctx, client, http.MethodPut, url, args)
}

func DELETE[R any](ctx context.Context, client *http.Client, url string, args any) (R, error) {
	return Call[R](ctx, client, http.MethodDelete, url, args)
}

func Call[R any](ctx context.Context, client *http.Client, method, url string, args any) (R, error) {
	var respData R
	if args == nil {
		args = EmptyArgs{}
	}
	req, err := NewRequest(ctx, method, url, args)
	if err != nil {
		return respData, errorcode.Detail(errorcode.BadRequestError, err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		return respData, errorcode.Detail(errorcode.DoRequestError, err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf(err.Error()+"io.ReadAll", zap.Error(err))
		return respData, errorcode.Detail(errorcode.ReadResponseError, err.Error())
	}
	tlog.WithContext(ctx).Debug(
		"http response",
		zap.String("status", resp.Status),
		zap.String("method", req.Method),
		zap.Stringer("url", req.URL),
		zap.ByteString("body", body),
	)

	//几种错误处理
	if !statusOK(resp.StatusCode) {
		return respData, errorcode.Detail(errorcode.RequestFailedError, string(body))
	}
	//这种一般都是无返回内容，直接返回正确即可
	if len(body) <= 0 {
		return respData, nil
	}
	err = jsoniter.Unmarshal(body, &respData)
	if err != nil {
		log.Printf("json.Unmarshal error:%v", err.Error())
		return respData, errorcode.Detail(errorcode.UnmarshalResponseError, err.Error())
	}
	return respData, nil
}

func statusOK(code int) bool {
	return code-http.StatusOK < 100
}
