package impl

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"path"

	task_center_v1 "github.com/kweaver-ai/idrm-go-common/api/task_center/v1"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-common/rest/task_center"
)

type workOrderTasks struct {
	base   *url.URL
	client *http.Client
}

// Get implements task_center.WorkOrderTaskInterface.
func (c *workOrderTasks) Get(ctx context.Context, id string) (result *task_center_v1.WorkOrderTask, err error) {
	u := *c.base
	u.Path = path.Join(u.Path, id)
	return base.GET[*task_center_v1.WorkOrderTask](ctx, c.client, u.String(), nil)
}

var _ task_center.WorkOrderTaskInterface = &workOrderTasks{}

type workOrderTasksInternal struct {
	workOrderTasks task_center.WorkOrderTaskInterface
	dataSources    configuration_center.DataSourceService
	departments    configuration_center.DepartmentService
}

// Get implements task_center.WorkOrderTaskInternalInterface.
func (c *workOrderTasksInternal) Get(ctx context.Context, id string) (result *task_center_v1.WorkOrderTaskInternal, err error) {
	task, err := c.workOrderTasks.Get(ctx, id)
	if err != nil {
		return
	}

	result = new(task_center_v1.WorkOrderTaskInternal)
	err = Aggregate_task_center_v1_WorkOrderTask_To_task_center_v1_WorkOrderTaskInternal(ctx, c.dataSources, c.departments, task, result)
	return
}

var _ task_center.WorkOrderTaskInternalInterface = &workOrderTasksInternal{}

func Aggregate_task_center_v1_WorkOrderTask_To_task_center_v1_WorkOrderTaskInternal(
	ctx context.Context,
	dataSources configuration_center.DataSourceService,
	departments configuration_center.DepartmentService,
	in *task_center_v1.WorkOrderTask,
	out *task_center_v1.WorkOrderTaskInternal,
) (err error) {
	out.ID = in.ID
	out.ThirdPartyID = in.ThirdPartyID
	out.CreatedAt = in.CreatedAt
	out.UpdatedAt = in.UpdatedAt
	out.Name = in.Name
	out.WorkOrderID = in.WorkOrderID
	out.Status = in.Status
	out.Reason = in.Reason
	out.Link = in.Link

	if err = Aggregate_task_center_v1_WorkOrderTaskTypedDetail_To_WorkOrderTaskTypedDetailInternal(
		ctx,
		dataSources,
		departments,
		&in.WorkOrderTaskTypedDetail,
		&out.WorkOrderTaskTypedDetailInternal,
	); err != nil {
		return
	}

	return
}

func Aggregate_task_center_v1_WorkOrderTaskTypedDetail_To_WorkOrderTaskTypedDetailInternal(
	ctx context.Context,
	dataSources configuration_center.DataSourceService,
	departments configuration_center.DepartmentService,
	in *task_center_v1.WorkOrderTaskTypedDetail,
	out *task_center_v1.WorkOrderTaskTypedDetailInternal,
) (err error) {
	// 归集
	if in.DataAggregation != nil {
		out.DataAggregation = make([]task_center_v1.WorkOrderTaskDetailAggregationDetailInternal, len(in.DataAggregation))
		for i := range in.DataAggregation {
			if err = Aggregate_task_center_v1_WorkOrderTaskDetailAggregationDetail_To_task_center_v1_WorkOrderTaskDetailAggregationDetailInternal(
				ctx,
				dataSources,
				departments,
				&in.DataAggregation[i],
				&out.DataAggregation[i],
			); err != nil {
				return
			}
		}
	}
	// 融合
	if in.DataFusion != nil {
		out.DataFusion = new(task_center_v1.WorkOrderTaskDetailFusionDetail)
		if err = Aggregate_task_center_v1_WorkOrderTaskDetailFusionDetail_To_task_center_v1_WorkOrderTaskDetailFusionDetail(
			ctx,
			dataSources,
			in.DataFusion,
			out.DataFusion,
		); err != nil {
			return
		}
	}
	// 质量
	if in.DataQuality != nil {
		err = errors.New("unimplemented")
		return
	}
	// 质量稽查
	if in.DataQualityAudit != nil {
		out.DataQualityAudit = make([]*task_center_v1.WorkOrderTaskDetailQualityAuditDetail, len(in.DataQualityAudit))
		for i := range in.DataQualityAudit {
			out.DataQualityAudit[i] = new(task_center_v1.WorkOrderTaskDetailQualityAuditDetail)
			if err = Aggregate_task_center_v1_WorkOrderTaskDetailQualityAuditDetail_To_task_center_v1_WorkOrderTaskDetailQualityAuditDetail(
				ctx,
				dataSources,
				in.DataQualityAudit[i],
				out.DataQualityAudit[i],
			); err != nil {
				return
			}
		}
	}

	return
}

func Aggregate_task_center_v1_WorkOrderTaskDetailAggregationDetail_To_task_center_v1_WorkOrderTaskDetailAggregationDetailInternal(
	ctx context.Context,
	dataSources configuration_center.DataSourceService,
	departments configuration_center.DepartmentService,
	in *task_center_v1.WorkOrderTaskDetailAggregationDetail,
	out *task_center_v1.WorkOrderTaskDetailAggregationDetailInternal,
) (err error) {
	// 部门 ID，第三方 ID -> AnyFabric ID
	{
		list, err := departments.GetDepartmentListInternal(ctx, configuration_center.QueryPageReqParam{Offset: 1, ThirdDeptID: in.DepartmentID})
		if err != nil {
			return err
		}
		for _, d := range list.Entries {
			if d.ThirdDeptID != in.DepartmentID {
				continue
			}
			out.DepartmentID = d.ID
			break
		}
	}
	// 源表
	if err = Aggregate_task_center_v1_WorkOrderTaskDetailAggregationTableReference_To_task_center_v1_WorkOrderTaskDetailAggregationTableReferenceInternal(
		ctx,
		dataSources,
		&in.Source,
		&out.Source,
	); err != nil {
		return
	}
	// 目标表
	if err = Aggregate_task_center_v1_WorkOrderTaskDetailAggregationTableReference_To_task_center_v1_WorkOrderTaskDetailAggregationTableReferenceInternal(
		ctx,
		dataSources,
		&in.Target,
		&out.Target,
	); err != nil {
		return
	}

	return
}

func Aggregate_task_center_v1_WorkOrderTaskDetailAggregationTableReference_To_task_center_v1_WorkOrderTaskDetailAggregationTableReferenceInternal(
	ctx context.Context,
	dataSources configuration_center.DataSourceService,
	in *task_center_v1.WorkOrderTaskDetailAggregationTableReference,
	out *task_center_v1.WorkOrderTaskDetailAggregationTableReferenceInternal,
) (err error) {
	out.TableName = in.TableName

	page, err := dataSources.GetDatasourcesByHuaAoID(ctx, in.DatasourceID)
	if err != nil {
		return
	}
	out.DatasourceID = page.ID
	return
}

func Aggregate_task_center_v1_WorkOrderTaskDetailFusionDetail_To_task_center_v1_WorkOrderTaskDetailFusionDetail(
	ctx context.Context,
	dataSources configuration_center.DataSourceService,
	in, out *task_center_v1.WorkOrderTaskDetailFusionDetail,
) (err error) {
	out.DatasourceName = in.DatasourceName
	out.DataTable = in.DataTable

	// 数据源 ID
	s, err := dataSources.GetDatasourcesByHuaAoID(ctx, in.DatasourceID)
	if err != nil {
		return
	}
	out.DatasourceID = s.ID

	return
}

func Aggregate_task_center_v1_WorkOrderTaskDetailQualityAuditDetail_To_task_center_v1_WorkOrderTaskDetailQualityAuditDetail(
	ctx context.Context,
	dataSources configuration_center.DataSourceService,
	in, out *task_center_v1.WorkOrderTaskDetailQualityAuditDetail,
) (err error) {
	out.ID = in.ID
	out.WorkOrderID = in.WorkOrderID
	out.DatasourceName = in.DatasourceName
	out.DataTable = in.DataTable
	out.DetectionScheme = in.DetectionScheme
	out.Status = in.Status
	out.Reason = in.Reason
	out.Link = in.Link

	// 数据源 ID
	s, err := dataSources.GetDatasourcesByHuaAoID(ctx, in.DatasourceID)
	if err != nil {
		return
	}
	out.DatasourceID = s.ID

	return
}
