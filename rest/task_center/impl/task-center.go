package impl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	cc "github.com/kweaver-ai/idrm-go-common/rest/configuration_center/impl"
	driven "github.com/kweaver-ai/idrm-go-common/rest/task_center"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewDriven(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.TaskCenterHost,
		httpClient: httpClient,
	}
}

func (d *DrivenImpl) GetTaskDetailById(ctx context.Context, id string) (*driven.GetTaskDetailByIdRes, error) {
	urlStr := fmt.Sprintf("%s/api/task-center/v1/projects/tasks/%s", d.baseURL, id)
	log.Infof("task_center Driven GetTaskDetailById url:%s \n", urlStr)
	res := &driven.GetTaskDetailByIdRes{}
	err := base.CallWithToken(ctx, d.httpClient, "task_center Driven GetTaskDetailById", http.MethodGet, urlStr, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *DrivenImpl) GetWorkOrderList(ctx context.Context, req *driven.GetWorkOrderListReq) (data *driven.GetWorkOrderListResp, err error) {
	errorMsg := "task_center GetWorkOrderList "
	urlStr := fmt.Sprintf("%s/api/internal/task-center/v1/work-order/list", d.baseURL)
	log.WithContext(ctx).Infof("errorMsg :%s,urlStr :%s", errorMsg, urlStr)
	statusCode, body, err := base.DOWithOutToken(ctx, errorMsg, http.MethodPost, urlStr, d.httpClient, req)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetWorkOrderListError, err.Error())
	}
	if statusCode != http.StatusOK {
		return nil, base.StatusCodeNotOK(errorMsg, statusCode, body)
	}
	data = &driven.GetWorkOrderListResp{}
	err = jsoniter.Unmarshal(body, &data)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return data, errorcode.Detail(errorcode.GetWorkOrderListError, err.Error())
	}
	return
}

func (d *DrivenImpl) GetProjectModels(ctx context.Context, id string) (data *driven.ProjectModelInfo, err error) {
	urlStr := fmt.Sprintf("%s/api/internal/task-center/v1/projects/%s/process", d.baseURL, id)
	log.Infof("task_center Driven GetProjectModels url:%s \n", urlStr)
	data, err = base.GET[*driven.ProjectModelInfo](ctx, d.httpClient, urlStr, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DrivenImpl) WorkOrders() driven.WorkOrderInterface {
	return &workOrders{
		baseURL:    d.baseURL + "/api/task-center/v1/work-order",
		httpClient: d.httpClient,
	}
}

func (d *DrivenImpl) GetCatalogTaskStatus(ctx context.Context, formId, formName, catalogId string) (data *driven.CatalogTaskStatusResp, err error) {
	urlStr := fmt.Sprintf("%s/api/internal/task-center/v1/catalog-task-status?form_id=%s&form_name=%s&catalog_id=%s", d.baseURL, formId, formName, catalogId)
	log.Infof("task_center Driven GetCatalogTaskStatus url:%s \n", urlStr)
	data, err = base.GET[*driven.CatalogTaskStatusResp](ctx, d.httpClient, urlStr, struct{}{})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DrivenImpl) GetCatalogTask(ctx context.Context, formId, formName, catalogId string) (*driven.CatalogTaskResp, error) {
	urlStr := fmt.Sprintf("%s/api/task-center/v1/catalog-task?form_id=%s&form_name=%s&catalog_id=%s", d.baseURL, formId, formName, catalogId)
	log.Infof("task_center Driven GetCatalogTask url:%s \n", urlStr)
	res := &driven.CatalogTaskResp{}
	err := base.CallWithToken(ctx, d.httpClient, "task_center Driven GetCatalogTask", http.MethodGet, urlStr, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *DrivenImpl) WorkOrderTasks() driven.WorkOrderTaskInterface {
	u, _ := url.Parse(d.baseURL)
	u.Path = path.Join(u.Path, "api", "task-center", "v1", "work-order-tasks")
	return &workOrderTasks{
		base:   u,
		client: d.httpClient,
	}
}

// WorkOrderTasksInternal implements task_center.Driven.
func (d *DrivenImpl) WorkOrderTasksInternal() driven.WorkOrderTaskInternalInterface {
	cc := cc.NewConfigurationCenterDrivenByService(d.httpClient)
	return &workOrderTasksInternal{
		workOrderTasks: d.WorkOrderTasks(),
		dataSources:    cc,
		departments:    cc,
	}
}

func (d *DrivenImpl) GetDataAggregationTask(ctx context.Context, formNames string) (data *driven.DataAggregationTaskResp, err error) {
	urlStr := fmt.Sprintf("%s/api/internal/task-center/v1/work-order-tasks/data-aggregation?form_names=%s", d.baseURL, formNames)
	log.Infof("task_center Driven GetDataAggregationTask url:%s \n", urlStr)
	data, err = base.GET[*driven.DataAggregationTaskResp](ctx, d.httpClient, urlStr, struct{}{})
	if err != nil {
		return nil, err
	}
	return data, nil
}
