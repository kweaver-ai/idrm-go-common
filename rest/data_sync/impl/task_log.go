package impl

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_sync"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

func (d DrivenImpl) QueryTaskHistory(ctx context.Context, args *driven.TaskLogReq) (*driven.TaskLogDetail, error) {
	var err error
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/model/history"
	resp, err := base.GET[driven.CsCommonResp[*driven.TaskLogDetail]](ctx, d.httpClient, url, args)
	if err != nil {
		return nil, err
	}
	return resp.Data, err
}
