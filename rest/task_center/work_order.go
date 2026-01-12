package task_center

import (
	"context"

	v1 "github.com/kweaver-ai/idrm-go-common/api/task_center/v1"
)

type WorkOrderGetter interface {
	WorkOrders() WorkOrderInterface
}

type WorkOrderInterface interface {
	Create(ctx context.Context, workOrder *v1.WorkOrder) (result *v1.WorkOrder, err error)
}
