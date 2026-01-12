package task_center

import (
	"context"

	task_center_v1 "github.com/kweaver-ai/idrm-go-common/api/task_center/v1"
)

//go:generate mockgen -destination=mock/work_order_task.go -package=mock -typed -write_generate_directive -source work_order_task.go

type WorkOrderTaskGetter interface {
	WorkOrderTasks() WorkOrderTaskInterface
}

type WorkOrderTaskInterface interface {
	Get(ctx context.Context, id string) (result *task_center_v1.WorkOrderTask, err error)
}

type WorkOrderTaskInternalGetter interface {
	WorkOrderTasksInternal() WorkOrderTaskInternalInterface
}

type WorkOrderTaskInternalInterface interface {
	Get(ctx context.Context, id string) (result *task_center_v1.WorkOrderTaskInternal, err error)
}
