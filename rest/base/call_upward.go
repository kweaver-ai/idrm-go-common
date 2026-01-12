package base

import (
	"bytes"
	"context"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"go.uber.org/zap"
)

func CallWithTokenUpward(ctx context.Context, client *http.Client, errorMsg string, method string, urlStr string, postReq any, res any, code string) (err error) {
	auth, err := interception.AuthFromContextCompatible(ctx)
	if err != nil {
		return
	}
	return callUpward(ctx, client, errorMsg, method, urlStr, postReq, res, code, auth)
}
func CallInternalUpward(ctx context.Context, client *http.Client, errorMsg string, method string, urlStr string, postReq any, res string, code string) (err error) {
	return callUpward(ctx, client, errorMsg, method, urlStr, postReq, res, code, "")
}
func callUpward(ctx context.Context, client *http.Client, errorMsg string, method string, urlStr string, postReq any, res any, code string, auth string) error {
	head := http.Header{}
	if auth != "" {
		head.Set("Authorization", auth)
	}
	var reqBody io.Reader
	switch method {
	case http.MethodGet:
		reqBody = nil
	case http.MethodDelete:
		reqBody = nil
	case http.MethodPost:
		jsonReq, err := jsoniter.Marshal(postReq)
		if err != nil {
			log.WithContext(ctx).Error(errorMsg+" json.Marshal error", zap.Error(err))
			return errorcode.Detail(code, err.Error())
		}
		reqBody = bytes.NewReader(jsonReq)
		head.Set("Content-Type", "application/json")
	}

	request, _ := http.NewRequest(method, urlStr, reqBody)
	request.Header = head
	resp, err := client.Do(request.WithContext(ctx))
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"client.Do error", zap.Error(err))
		return errorcode.Detail(code, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"io.ReadAll", zap.Error(err))
		return errorcode.Detail(code, err.Error())
	}
	log.Infof(errorMsg+" body:%s \n ", body)
	if resp.StatusCode != http.StatusOK {
		return StatusCodeNotOK(errorMsg, resp.StatusCode, body, code)
	}

	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.WithContext(ctx).Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return errorcode.Detail(code, err.Error())
	}
	return nil
}
