package data_sync

import (
	"context"
	"encoding/json"
)

type Driven interface {
	CollectingModel
	ProcessingModel
	SchedulePlan
	ModelTask
	TaskLog
}

type TaskLog interface {
	QueryTaskHistory(ctx context.Context, req *TaskLogReq) (*TaskLogDetail, error)
}

type CollectingModel interface {
	CreateCollectingModel(ctx context.Context, id string, model *CollectModelReq) error
	UpdateCollectingModel(ctx context.Context, id string, model *CollectModelReq) error
	DeleteCollectingModel(ctx context.Context, id string) error
}

type ProcessingModel interface {
	CreateProcessingModel(ctx context.Context, id string, model *ProcessModelCUReq) (CsCommonResp[any], error)
	UpdateProcessingModel(ctx context.Context, id string, model *ProcessModelCUReq) (CsCommonResp[any], error)
	DeleteProcessingModel(ctx context.Context, id string) (CsCommonResp[any], error)
}

type SchedulePlan interface {
	CreateWorkflow(ctx context.Context, id string, plan *WorkflowReq) (CsCommonResp[any], error)
	UpdateWorkflow(ctx context.Context, id string, plan *WorkflowReq) (CsCommonResp[any], error)
	DeleteWorkflow(ctx context.Context, id string) (CsCommonResp[any], error)
	UpdateWorkflowOnline(ctx context.Context, workflowOnline *WorkflowOnline) (CsCommonResp[any], error)
	UpdateWorkflowTimer(ctx context.Context, workflowOnline *WorkflowTimer) (CsCommonResp[any], error)
}
type ModelTask interface {
	Run(ctx context.Context, mid string) error
}

type CsCommonResp[T any] struct {
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Detail      json.RawMessage `json:"detail"`
	Solution    string          `json:"solution"`
	Data        T               `json:"data"`
}

type CollectModelCreateReq struct {
	ID              string        `json:"model_uuid"`
	Name            string        `json:"model_name"`
	SourceDsId      string        `json:"source_ds_id"`
	TargetDsId      string        `json:"target_ds_id"`
	SourceTableName string        `json:"source_table_name"`
	TargetTableName string        `json:"target_table_name"`
	TargetTableSql  string        `json:"target_table_sql"`
	SourceFields    []*Field      `json:"source_fields"`
	TargetFields    []*Field      `json:"target_fields"`
	AdvancedParams  []interface{} `json:"advanced_params"`
}

type CollectModelReq struct {
	Name            string        `json:"model_name"`
	SourceDsId      string        `json:"source_ds_id"`
	TargetDsId      string        `json:"target_ds_id"`
	SourceTableName string        `json:"source_table_name"`
	TargetTableName string        `json:"target_table_name"`
	TargetTableSql  string        `json:"target_table_sql"`
	SourceFields    []*Field      `json:"source_fields"`
	TargetFields    []*Field      `json:"target_fields"`
	AdvancedParams  []interface{} `json:"advanced_params"`
}

type Field struct {
	Name           string `json:"field_name"`
	Type           string `json:"field_type"`
	Length         *int   `json:"field_length"`
	FieldPrecision *int   `json:"field_precision"`
	Description    string `json:"field_description"`
}

type WorkflowCreateReq struct {
	ID           string            `json:"process_uuid"`
	Crontab      string            `json:"crontab"`
	ProcessName  string            `json:"process_name"`
	StartTime    string            `json:"start_time"`
	EndTime      string            `json:"end_time"`
	OnlineStatus int               `json:"crontab_status"` //是否启用改表达式，0禁用，1启用
	Models       []CollectionModel `json:"models"`
}

type WorkflowReq struct {
	Crontab      string            `json:"crontab"`
	ProcessName  string            `json:"process_name"`
	StartTime    string            `json:"start_time"`
	EndTime      string            `json:"end_time"`
	OnlineStatus int               `json:"crontab_status"` //是否启用改表达式，0禁用，1启用
	Models       []CollectionModel `json:"models"`
}

type CollectionModel struct {
	Uuid       string `json:"model_uuid"`
	ModelType  int8   `json:"model_type"` //模型类型，默认1 同步，2 加工
	Dependency string `json:"dependency"`
}

type WorkflowOnline struct {
	ID           string `json:"process_uuid"`
	OnlineStatus int    `json:"crontab_status"`
}

type WorkflowTimer struct {
	ID           string `json:"process_uuid"`
	Crontab      string `json:"crontab"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	OnlineStatus int    `json:"crontab_status"`
}

type TaskLogReq struct {
	Offset          int    `query:"offset"`
	Limit           int    `query:"limit"`
	Direction       string `query:"direction"`
	Sort            string `query:"sort"`
	Step            string `query:"step"`
	Status          string `query:"status"`          //执行状态，SUCCESS FAILURE RUNNING_EXECUTION
	ScheduleExecute string `query:"scheduleExecute"` //执行方式，手动执行还是自动执行
	ModelUUID       string `query:"model_uuid"`
}

type TaskLogDetail struct {
	Total       int64          `json:"total"`        //总数
	TotalPage   int            `json:"total_page"`   //总页数
	PageSize    int            `json:"page_size"`    //每页数量
	CurrentPage int            `json:"current_page"` //当前页码
	ModelUuid   string         `json:"model_uuid"`   //
	ModelName   string         `json:"model_name"`
	TotalList   []*TaskLogInfo `json:"total_list"`
}

type TaskLogInfo struct {
	StartTime         string `json:"start_time"`          //请求时间
	EndTime           string `json:"end_time"`            //完成时间
	SyncCount         string `json:"sync_count"`          //推送总数
	SyncTime          string `json:"sync_time"`           //执行时间，单位，秒
	SyncMethod        string `json:"sync_method"`         //执行方式
	Status            string `json:"status"`              //状态
	StepName          string `json:"step_name"`           //步骤名称，uuid
	StepId            string `json:"step_id"`             //步骤ID，序号
	ProcessInstanceId string `json:"process_instance_id"` //处理实例ID
	ErrorMessage      string `json:"error_message"`       //错误信息
}

type ProcessModelCreateReq struct {
	ID                string        `json:"model_uuid"`
	Name              string        `json:"model_name"`
	TargetDsId        string        `json:"target_ds_id"`
	TargetTableName   string        `json:"target_table_name"`
	TargetTableSql    string        `json:"target_table_sql"`
	TargetTableInsert string        `json:"target_table_insert"`
	AdvancedParams    []interface{} `json:"advanced_params"`
}

type ProcessModelCUReq struct {
	Name              string        `json:"model_name"`
	TargetDsId        string        `json:"target_ds_id"`
	TargetTableName   string        `json:"target_table_name"`
	TargetTableSql    string        `json:"target_table_sql"`
	TargetTableInsert string        `json:"target_table_insert"`
	AdvancedParams    []interface{} `json:"advanced_params"`
}
