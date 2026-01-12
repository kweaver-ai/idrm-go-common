package impl

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_sync"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

func (d DrivenImpl) CreateWorkflow(ctx context.Context, id string, plan *driven.WorkflowReq) (resp driven.CsCommonResp[any], err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	args := &driven.WorkflowCreateReq{
		ID: id,
	}
	copier.Copy(args, plan)
	url := d.baseURL + "/api/data-sync/v1/process"
	resp, err = base.POST[driven.CsCommonResp[any]](ctx, d.httpClient, url, args)
	return resp, err
}

func (d DrivenImpl) UpdateWorkflow(ctx context.Context, id string, plan *driven.WorkflowReq) (resp driven.CsCommonResp[any], err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/process/" + id
	resp, err = base.PUT[driven.CsCommonResp[any]](ctx, d.httpClient, url, plan)
	return resp, err
}

func (d DrivenImpl) DeleteWorkflow(ctx context.Context, id string) (resp driven.CsCommonResp[any], err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/process/" + id
	resp, err = base.DELETE[driven.CsCommonResp[any]](ctx, d.httpClient, url, nil)
	return resp, err
}

func (d DrivenImpl) UpdateWorkflowOnline(ctx context.Context, workflowOnline *driven.WorkflowOnline) (resp driven.CsCommonResp[any], err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/process/cron/online"
	resp, err = base.POST[driven.CsCommonResp[any]](ctx, d.httpClient, url, workflowOnline)
	return resp, err
}

func (d DrivenImpl) UpdateWorkflowTimer(ctx context.Context, workflowOnline *driven.WorkflowTimer) (resp driven.CsCommonResp[any], err error) {
	ctx, _ = trace.StartInternalSpan(ctx)
	defer trace.EndSpan(ctx, err)

	url := d.baseURL + "/api/data-sync/v1/process/cron"
	resp, err = base.POST[driven.CsCommonResp[any]](ctx, d.httpClient, url, workflowOnline)
	return resp, err
}
