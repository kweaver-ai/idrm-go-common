package impl

import (
	"context"
	"fmt"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/task_center"
)

func (d *DrivenImpl) GetSandboxSampleInfo(ctx context.Context, sandboxID string) (*driven.SandboxSpaceDetail, error) {
	urlStr := fmt.Sprintf("%s/api/internal/task-center/v1/sandbox/%s", d.baseURL, sandboxID)
	return base.GET[*driven.SandboxSpaceDetail](ctx, d.httpClient, urlStr, nil)
}

func (d *DrivenImpl) GetUserSandboxSampleInfo(ctx context.Context) (*base.PageResult[driven.SandboxSpaceListItem], error) {
	args := struct {
		Offset int `query:"offset"`
		Limit  int `query:"limit"`
	}{
		Offset: 1,
		Limit:  2000,
	}
	urlStr := fmt.Sprintf("%s/api/task-center/v1/sandbox/space", d.baseURL)
	return base.GET[*base.PageResult[driven.SandboxSpaceListItem]](ctx, d.httpClient, urlStr, args)
}
