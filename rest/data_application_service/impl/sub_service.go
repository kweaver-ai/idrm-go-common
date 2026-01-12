package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_application_service"
)

func (d *DrivenImpl) GetSubServiceSimple(ctx context.Context, id string) (*driven.SubService, error) {
	urlStr := d.baseURL + fmt.Sprintf("/api/internal/data-application-service/v1/sub-service/%s", id)
	return base.GET[*driven.SubService](ctx, d.httpClient, urlStr, nil)
}

func (d *DrivenImpl) GeSubServiceByServices(ctx context.Context, ids []string) (map[string][]string, error) {
	urlStr := d.baseURL + fmt.Sprintf("/api/data-application-service/internal/v1/services/sub-service/batch?service_id=%s", strings.Join(ids, ","))
	resp, err := base.GET[map[string][]string](ctx, d.httpClient, urlStr, nil)
	return resp, err
}
