package impl

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_catalog"
)

func (d *DrivenImpl) CreateDataPush(ctx context.Context, req *driven.DataPushCreateReq) (string, error) {
	url := d.baseURL + "/api/data-catalog/v1/data-push"
	//处理参数
	resp, err := base.POST[base.IDNameResp](ctx, d.httpClient, url, req)
	if err != nil {
		return "", err
	}
	return resp.ID, err
}

func (d *DrivenImpl) UpdateDataPush(ctx context.Context, req *driven.UpdateReq) (string, error) {
	url := d.baseURL + "/api/data-catalog/v1/data-push"
	//处理参数
	resp, err := base.PUT[base.IDNameResp](ctx, d.httpClient, url, req)
	if err != nil {
		return "", err
	}
	return resp.ID, err
}

func (d *DrivenImpl) History(ctx context.Context, req *driven.TaskExecuteHistoryReq) (*base.PageResult[driven.TaskLogInfo], error) {
	url := d.baseURL + "/api/internal/data-catalog/v1/data-push/schedule/history"
	//处理参数
	resp, err := base.GET[base.PageResult[driven.TaskLogInfo]](ctx, d.httpClient, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// BatchUpdateStatus 批量更新草稿状态的推送
// 返回成功的参数，中途遇到错误不会中止，同步执行，需要等到执行完毕才返回
func (d *DrivenImpl) BatchUpdateStatus(ctx context.Context, req *driven.BatchUpdateStatusReq) ([]uint64, error) {
	url := d.baseURL + "/api/data-catalog/v1/data-push/statues"
	//处理参数
	resp, err := base.PUT[[]uint64](ctx, d.httpClient, url, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (d *DrivenImpl) SandboxPushCount(ctx context.Context, ids []string) (map[string]int, error) {
	url := d.baseURL + "/api/internal/data-catalog/v1/data-push/sandbox"
	//处理参数
	args := struct {
		AuthedSandboxID []string `query:"authed_sandbox_id"`
	}{
		AuthedSandboxID: ids,
	}
	resp, err := base.GET[map[string]int](ctx, d.httpClient, url, args)
	if err != nil {
		return nil, err
	}
	return resp, err
}
