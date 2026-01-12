package impl

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_sync"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

func (d DrivenImpl) CreateCollectingModel(ctx context.Context, id string, model *driven.CollectModelReq) (err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	args := &driven.CollectModelCreateReq{
		ID: id,
	}
	copier.Copy(args, model)

	url := d.baseURL + "/api/data-sync/v1/model"
	_, err = base.POST[driven.CsCommonResp[any]](ctx, d.httpClient, url, args)
	return err
}

func (d DrivenImpl) UpdateCollectingModel(ctx context.Context, id string, model *driven.CollectModelReq) (err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/model/" + id
	_, err = base.PUT[driven.CsCommonResp[any]](ctx, d.httpClient, url, model)
	return err
}

func (d DrivenImpl) DeleteCollectingModel(ctx context.Context, id string) (err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/model/" + id
	_, err = base.DELETE[driven.CsCommonResp[any]](ctx, d.httpClient, url, nil)
	return err
}
