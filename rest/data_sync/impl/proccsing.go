package impl

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_sync"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

func (d DrivenImpl) CreateProcessingModel(ctx context.Context, id string, model *driven.ProcessModelCUReq) (resp driven.CsCommonResp[any], err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	args := &driven.ProcessModelCreateReq{
		ID: id,
	}
	copier.Copy(args, model)

	url := d.baseURL + "/api/data-sync/v1/compose/model"
	resp, err = base.POST[driven.CsCommonResp[any]](ctx, d.httpClient, url, args)
	return resp, err
}

func (d DrivenImpl) UpdateProcessingModel(ctx context.Context, id string, model *driven.ProcessModelCUReq) (resp driven.CsCommonResp[any], err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/compose/model/" + id
	resp, err = base.PUT[driven.CsCommonResp[any]](ctx, d.httpClient, url, model)
	return resp, err
}

func (d DrivenImpl) DeleteProcessingModel(ctx context.Context, id string) (resp driven.CsCommonResp[any], err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/compose/model/" + id
	resp, err = base.DELETE[driven.CsCommonResp[any]](ctx, d.httpClient, url, nil)
	return resp, err
}
