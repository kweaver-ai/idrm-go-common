package business_grooming

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/label"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type Driven interface {
	BusinessDomain
	BusinessForm
	Processing
}

type BusinessDomain interface {
	GetBusinessNodesBrief(ctx context.Context, ids []string) (details []*BusinessNode, err error)
	GetNodeChild(ctx context.Context, nodeID string) ([]*BusinessNodeObject, error)
}

type BusinessForm interface {
	GetBusinessFormSource(ctx context.Context, formID string) (sourceFormIDs *BusinessFormSourceRes, err error)
	GetBusinessFormDetails(ctx context.Context, formIDs, tableKinds []string, pageNumber, pageSize int) (formDetails []*BusinessFormDetail, err error)
	GetFormViewTableDict(ctx context.Context, businessModelID string, ids []string) (res map[string]string, err error)
	GetBusinessFormDetailsFilterLabel(ctx context.Context, rangeTypeKey string, formIDs, tableKinds []string, pageNumber, pageSize int) (formDetails []*BusinessFormAndLabelDetail, err error)
}

type Processing interface {
	QueryDataTable(ctx context.Context, datasourceID string, tableName string) ([]*DataTableFieldInfo, error)
	CreateSyncTask(ctx context.Context, req *CollectingModelCreateReq) (*base.IDNameResp, error)
	CreateWorkflow(ctx context.Context, req *WorkflowCreateReq) (*base.IDNameResp, error)
	CheckTableName(ctx context.Context, datasourceID, tableName string) (bool, error)
	QueryFormPathInfo(ctx context.Context, formID string) (*FormPathInfoResp, error)
	// BatchQueryMainBusinessModel 批量查询主干业务及其关联的业务模型信息（内部函数）
	BatchQueryMainBusinessModel(ctx context.Context, mainBusinessIDs []string) ([]MainBusinessModelResp, error)
}

type CollectingModelCreateReq struct {
	Source      *DataTableCreateReq `json:"source" form:"source" binding:"required"`
	Target      *DataTableCreateReq `json:"target" form:"target" binding:"required"`
	Name        string              `json:"name" form:"name" example:"采集模型名" binding:"required,VerifyName128NoSpace"`                      // 采集模型名
	Description string              `json:"description" form:"description" binding:"omitempty,VerifyDescription300"`                       //描述
	TaskID      string              `json:"task_id" form:"task_id" binding:"required,uuid" example:"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"` // 任务id，uuid
}

type DataTableCreateReq struct {
	DatasourceID string            `json:"datasource_id" form:"datasource_id" uri:"datasource_id" binding:"required,uuid"` // 数据源id，uuid
	Name         string            `json:"name" form:"name" example:"表名" binding:"required"`                               // 表名
	Fields       []*FieldCreateReq `json:"fields" binding:"required,gte=1,dive"`
}

type FieldCreateReq struct {
	Name           string `json:"name" binding:"required" example:"字段1"`           // 字段名称
	Type           string `json:"type" binding:"required" example:"char"`          // 字段类型，对应虚拟化引擎数据源配置的 sourceType。
	Length         *int   `json:"length" example:"128" binding:"omitempty,min=0" ` // 字段长度
	FieldPrecision *int   `json:"field_precision" binding:"omitempty,min=0" `      // 字段精度
	Description    string `json:"description" binding:"omitempty" example:"字段注释"`  // 字段注释
	UnMapped       bool   `json:"unmapped"`                                        // 映射是否取消，true表示取消了映射。只有同步模型的target表需要该参数

	// 字段类型，用于通过虚拟化引擎的接口创建表，对应虚拟化引擎数据源配置的 olkSearchType。只有 target 表需要此参数。
	SearchType string `json:"searchType" example:"DECIMAL"`
}

//region CreateWorkflow

type WorkflowCreateReq struct {
	Name        string `json:"name" form:"name" example:"调度计划名" binding:"required,VerifyName128NoSpace"` // 调度计划名
	Description string `json:"description" form:"description"  binding:"omitempty,VerifyDescription300"` //描述
	WorkflowBase
}
type WorkflowBase struct {
	Nodes  []*WorkNode `json:"nodes" form:"canvas" binding:"required,min=1,dive"`
	Canvas string      `json:"canvas" form:"canvas" binding:"required"`                                                       // 位置等画布数据
	TaskID string      `json:"task_id" form:"task_id" binding:"required,uuid" example:"4a5a3cc0-0169-4d62-9442-62214d8fcd8d"` // 任务id，uuid（36）
}

type WorkNode struct {
	NodeId    string   `json:"node_id" form:"node_id" binding:"required,uuid"`                              //节点UUID
	ModelId   string   `json:"model_id" form:"model_id" binding:"required,uuid"`                            //节点对应的模型UUID
	ModelType string   `json:"model_type" form:"model_type" binding:"required,oneof=collecting processing"` //模型类型：采集或加工，collecting processing
	PreNodeId []string `json:"pre_node_id" form:"pre_node_id" binding:"required"`                           //前序节点ID数组
}

//endregion

type BusinessFormSourceRes struct {
	Forms []string `json:"forms"`
}

type BusinessFormDetail struct {
	ID                 string            `json:"id"`                          //业务表ID
	Name               string            `json:"name"`                        //业务表名称
	ModelID            string            `json:"business_model_id,omitempty"` //业务模型ID
	Description        string            `json:"description"`                 //业务表描述
	DepartmentID       string            `json:"department_id"`               //所属部门ID
	DepartmentName     string            `json:"department_name"`             //所属部门名称
	DepartmentPath     string            `json:"department_path"`             //所属部门路径
	RelatedInfoSystems []base.IDNameResp `json:"related_info_systems"`        //关联信息系统列表
	BusinessDomainID   string            `json:"business_domain_id"`          //业务流程ID
	BusinessDomainName string            `json:"business_domain_name"`        //业务流程名称
	DomainID           string            `json:"domain_id"`                   //业务域ID
	DomainName         string            `json:"domain_name"`                 //业务域名称
	DomainGroupID      string            `json:"domain_group_id"`             //业务域分组ID
	DomainGroupName    string            `json:"domain_group_name"`           //业务域分组名称
	UpdateAt           int64             `json:"update_at"`                   //更新时间
	UpdateBy           string            `json:"update_by"`                   //更新者
	UpdateByName       string            `json:"update_by_name"`
	LabelIds           []string          `json:"label_ids"` //资源标签：数组，标签ID
}

type BusinessFormAndLabelDetail struct {
	ID                 string             `json:"id"`                          //业务表ID
	Name               string             `json:"name"`                        //业务表名称
	ModelID            string             `json:"business_model_id,omitempty"` //业务模型ID
	Description        string             `json:"description"`                 //业务表描述
	DepartmentID       string             `json:"department_id"`               //所属部门ID
	DepartmentName     string             `json:"department_name"`             //所属部门名称
	DepartmentPath     string             `json:"department_path"`             //所属部门路径
	RelatedInfoSystems []base.IDNameResp  `json:"related_info_systems"`        //关联信息系统列表
	BusinessDomainID   string             `json:"business_domain_id"`          //业务流程ID
	BusinessDomainName string             `json:"business_domain_name"`        //业务流程名称
	DomainID           string             `json:"domain_id"`                   //业务域ID
	DomainName         string             `json:"domain_name"`                 //业务域名称
	DomainGroupID      string             `json:"domain_group_id"`             //业务域分组ID
	DomainGroupName    string             `json:"domain_group_name"`           //业务域分组名称
	UpdateAt           int64              `json:"update_at"`                   //更新时间
	UpdateBy           string             `json:"update_by"`                   //更新者
	UpdateByName       string             `json:"update_by_name"`
	LabelIds           []string           `json:"label_ids"`       //资源标签：数组，标签ID
	LabelListResp      []*label.LabelResp `json:"label_list_resp"` //关联标签列表
}

const (
	TableKindBusinessNode     = 1 // 业务节点表
	TableKindBusinessStandard = 2 // 业务标准表
	TableKindDataStandard     = 3 // 数据标准表
	TableKindDataMerge        = 4 // 数据融合表
)

type BusinessNode struct {
	ID          string `json:"id"`          // 业务节点ID
	Name        string `json:"name"`        // 业务节点名称
	Description string `json:"description"` // 业务节点描述
	Type        string `json:"type"`        // 业务节点类型
}

type BusinessNodeObject struct {
	ID   string `json:"id"`   //业务节点ID
	Name string `json:"name"` // 业务节点名称
}
