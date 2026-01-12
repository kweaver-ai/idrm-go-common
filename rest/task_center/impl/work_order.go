package impl

import (
	"context"
	"net/http"

	v1 "github.com/kweaver-ai/idrm-go-common/api/task_center/v1"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/task_center"
)

type workOrders struct {
	baseURL    string
	httpClient *http.Client
}

// Create implements task_center.WorkOrderInterface.
func (c *workOrders) Create(ctx context.Context, workOrder *v1.WorkOrder) (result *v1.WorkOrder, err error) {
	result = &v1.WorkOrder{}
	err = base.CallWithToken(ctx, c.httpClient, "create task_center/v1.WorkOrder", http.MethodPost, c.baseURL, workOrder, result)
	return
}

var _ task_center.WorkOrderInterface = &workOrders{}
