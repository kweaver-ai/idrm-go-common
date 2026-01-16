package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

func httpPostDo[T any](ctx context.Context, cli *http.Client, url string, req any, headers map[string]string) (*T, error) {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	b, err := json.Marshal(req)
	if err != nil {
		log.WithContext(ctx).Errorf("req param json.Marshal failed, err: %v", err)
		return nil, ReqParamError(err)
	}

	hReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		log.WithContext(ctx).Errorf("req param http.NewRequestWithContext failed, err: %v", err)
		return nil, ReqParamError(err)
	}

	for k, v := range headers {
		hReq.Header.Add(k, v)
	}

	log.WithContext(ctx).Infof("http req url: %s, body: %s", hReq.URL, b)
	hResp, err := cli.Do(hReq)
	if err != nil {
		log.WithContext(ctx).Errorf("req param http.Do failed, err: %v", err)
		return nil, ReqError(err)
	}
	defer func() {
		_ = hResp.Body.Close()
	}()

	respData, err := io.ReadAll(hResp.Body)
	if err != nil {
		log.WithContext(ctx).Errorf("resp data io.ReadAll failed, err: %v", err)
		err = nil
	}

	log.WithContext(ctx).Infof("http resp url: %s, code: %d, data: %s", hReq.URL, hResp.StatusCode, respData)

	if hResp.StatusCode < 200 || hResp.StatusCode >= 300 {
		log.WithContext(ctx).Errorf("req server failed, status code: %v, resp: %s", hResp.StatusCode, respData)
		return nil, RespBodyError(string(respData))
	}

	var resp T
	if err = json.Unmarshal(respData, &resp); err != nil {
		log.WithContext(ctx).Errorf("resp data json.Unmarshal failed, resp: %s, err: %v", respData, err)
		return nil, err
	}

	return &resp, nil
}
