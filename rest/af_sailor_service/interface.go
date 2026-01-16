package af_sailor_service

import (
	"context"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type Driven interface {
	UpdateGraph(ctx context.Context, detail *ModelDetail) (*base.IntIDResp, error)
	DeleteGraph(ctx context.Context, graphID int) (*base.IntIDResp, error)
	GraphBuildTask(ctx context.Context, req *GraphBuildTaskReq) (*base.IntIDResp, error)

	RecTable(ctx context.Context, req *RecTableReq, userId string) (*RecTableResp, error)
	RecFlow(ctx context.Context, req *RecFlowReq, userId string) (*RecFlowResp, error)
	RecCode(ctx context.Context, req *RecCodeReq) (*RecCodeResp, error)
	RecCheckCode(ctx context.Context, req *CheckCodeReq) (*CheckCodeResp, error)
	RecView(ctx context.Context, req *RecViewReq) (*RecViewResp, error)
	GraphNeighbors(ctx context.Context, req *GraphNeighborsReq) (*GraphNeighborsResp, error)
	GraphFullText(ctx context.Context, req *GraphFullTextReq) (*GraphFullTextResp, error)
	LogicalViewDataCategorize(ctx context.Context, req *LogicalViewDatacategorizeReq) (*LogicalViewDataCategorizeResp, error)
	TableCompletionTableInfo(ctx context.Context, req *TableCompletionTableInfoReqBody, authorization string) (*TableCompletionTableInfoResp, error)
	TableCompletionAll(ctx context.Context, req *TableCompletionReqBody, authorization string) (*TableCompletionResp, error)
}

type ModelDetail struct {
	ID              string         `json:"id"`                // 主键ID，uuid
	BusinessName    string         `json:"business_name"`     // 模型名称，业务名称
	TechnicalName   string         `json:"technical_name"`    // 模型技术名称
	CatalogID       string         `json:"catalog_id"`        // 目录的主键ID
	CatalogName     string         `json:"catalog_name"`      //目录的名称
	DataViewID      string         `json:"data_view_id"`      // 目录带的元数据视图ID
	DataViewName    string         `json:"data_view_name"`    // 视图名称
	GraphID         int            `json:"graph_id"`          //当前模型在图谱ID，整数
	SubjectID       string         `json:"subject_id"`        // 业务对象ID
	SubjectName     string         `json:"subject_name"`      // 业务对象名称
	Description     string         `json:"description"`       // 描述
	CreatedAt       int64          `json:"created_at"`        // 创建时间
	UpdatedAt       int64          `json:"updated_at"`        // 更新时间
	HasGraph        bool           `json:"has_graph"`         // 是否构建过图谱
	DisplayFieldKey string         `json:"display_field_key"` // 显示字段的key
	MetaModelSlice  []*ModelDetail `json:"meta_model_slice"`  // 元模型的信息
	Fields          []*TModelField `json:"fields"`            // 元模型字段
	Relations       []*Relation    `json:"relations"`         // 复合模型的关系
}

type Relation struct {
	ID                  string          `json:"id" binding:"omitempty,uuid"`                        // 关系ID
	BusinessName        string          `json:"business_name" binding:"required,VerifyName"`        // 业务名称
	TechnicalName       string          `json:"technical_name" binding:"required,VerifyNameEN"`     // 模型技术名称
	StartDisplayFieldID string          `json:"start_display_field_id" binding:"omitempty,uuid"`    // 起点显示属性
	EndDisplayFieldID   string          `json:"end_display_field_id" binding:"omitempty,uuid"`      // 终点显示属性
	Description         string          `json:"description"  binding:"TrimSpace,omitempty,lte=255"` // 描述
	Links               []*RelationLink `json:"links" binding:"required,gt=0,dive"`                 // 关联关系匹配的字段
}

type RelationLink struct {
	RelationID         string `json:"relation_id"  binding:"omitempty,uuid"`   // 模型关系ID
	StartModelID       string `json:"start_model_id"  binding:"required,uuid"` // 起点元模型ID
	StartModelName     string `json:"start_model_name"`                        // 起点模型名称
	StartModelTechName string `json:"start_model_tech_name"`                   // 起点模型技术名称
	StartFieldID       string `json:"start_field_id"  binding:"required,uuid"` // 起点字段ID
	StartFieldName     string `json:"start_field_name"`                        // 开始字段名称
	StartFieldTechName string `json:"start_field_tech_name"`                   // 开始字段名称
	EndModelID         string `json:"end_model_id"  binding:"required,uuid"`   // 终点元模型ID
	EndModelName       string `json:"end_model_name"`                          // 终点模型名称
	EndModelTechName   string `json:"end_model_tech_name"`                     // 终点模型技术名称
	EndFieldID         string `json:"end_field_id"  binding:"required,uuid"`   // 终点字段ID
	EndFieldName       string `json:"end_field_name"`                          // 终点字段名称
	EndFieldTechName   string `json:"end_field_tech_name"`                     // 终点字段名称
}

// TModelField 元模型字段表
type TModelField struct {
	ID            uint64 `json:"id"`             // 主键ID
	FieldID       string `json:"field_id"`       // 视图字段ID
	ModelID       string `json:"model_id"`       // 元模型ID
	TechnicalName string `json:"technical_name"` // 列技术名称
	BusinessName  string `json:"business_name"`  // 列业务名称
	DataType      string `json:"data_type"`      // 数据类型
	DataLength    int32  `json:"data_length"`    // 数据长度
	DataAccuracy  int32  `json:"data_accuracy"`  // 数据精度
	PrimaryKey    bool   `json:"primary_key"`    // 是否是主键,0不是，1是
	IsNullable    string `json:"is_nullable"`    // 是否为空
	Comment       string `json:"comment"`        // 字段注释
}

type GraphBuildTaskReq struct {
	GraphID  int    `json:"graph_id"`
	TaskType string `json:"task_type"`
}
