package impl

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_sync"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

func (d DrivenImpl) Run(ctx context.Context, mid string) (err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/model/run/" + mid
	_, err = base.POST[driven.CsCommonResp[any]](ctx, d.httpClient, url, nil)
	return err
}
