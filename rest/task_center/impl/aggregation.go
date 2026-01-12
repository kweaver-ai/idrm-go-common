package impl

import (
	"context"

	task_center_v1 "github.com/kweaver-ai/idrm-go-common/api/task_center/v1"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

// GetStandardFormAggregationInfo  获取业务标准表的归集信息,  待对接好再补上这个方法
func (d *DrivenImpl) GetStandardFormAggregationInfo(ctx context.Context, formID []string) ([]*task_center_v1.BusinessFormDataTableItem, error) {
	urlStr := d.baseURL + "/api/internal/task-center/v1/data-aggregation-inventories/data-tables"
	args := struct {
		ID []string `query:"id"`
	}{
		ID: formID,
	}
	resp, err := base.GET[[]*task_center_v1.BusinessFormDataTableItem](ctx, d.httpClient, urlStr, args)
	return resp, err
}
