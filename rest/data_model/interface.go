package data_model

import "context"

type Driven interface {
	GetDataModelByID(ctx context.Context, ids ...string) ([]*DataModel, error)
	GetDataModelByIDInternal(ctx context.Context, ids ...string) ([]*DataModel, error)

	// 数据视图行列权限（mdl-data-model 内部接口 /api/mdl-data-model/in/v1/data-view-row-column-rules）
	ListDataViewRowColumnRulesInternal(ctx context.Context, q ListDataViewRowColumnRulesQuery) (*ListDataViewRowColumnRulesResult, error)
	GetDataViewRowColumnRulesInternal(ctx context.Context, ruleIDs []string) ([]*DataViewRowColumnRule, error)
	CreateDataViewRowColumnRulesInternal(ctx context.Context, rules []DataViewRowColumnRuleWrite) ([]string, error)
	UpdateDataViewRowColumnRuleInternal(ctx context.Context, ruleID string, rule *DataViewRowColumnRuleWrite) error
	DeleteDataViewRowColumnRulesInternal(ctx context.Context, ruleIDs []string) error
}

type DataModel struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	TechnicalName  string   `json:"technical_name"`
	GroupId        string   `json:"group_id"`
	GroupName      string   `json:"group_name"`
	Type           string   `json:"type"`
	QueryType      string   `json:"query_type"`
	Tags           []any    `json:"tags"`
	Comment        string   `json:"comment"`
	Builtin        bool     `json:"builtin"`
	CreateTime     int64    `json:"create_time"`
	UpdateTime     int64    `json:"update_time"`
	DataSourceType string   `json:"data_source_type"`
	DataSourceId   string   `json:"data_source_id"`
	DataSourceName string   `json:"data_source_name"`
	Status         string   `json:"status"`
	Operations     []string `json:"operations"`
	Fields         []Field  `json:"fields"`
	ModuleType     string   `json:"module_type"`
	MetadataFormId string   `json:"metadata_form_id"`
	PrimaryKeys    []any    `json:"primary_keys"`
	SqlStr         string   `json:"sql_str"`
	MetaTableName  string   `json:"meta_table_name"`
}

type Field struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	Comment           string `json:"comment"`
	DisplayName       string `json:"display_name"`
	OriginalName      string `json:"original_name"`
	DataLength        int    `json:"data_length"`
	DataAccuracy      int    `json:"data_accuracy"`
	Status            string `json:"status"`
	IsNullable        string `json:"is_nullable"`
	BusinessTimestamp bool   `json:"business_timestamp"`
}

// RowColumnCondCfg 与 mdl-data-model interfaces.CondCfg 的 JSON 对齐。
type RowColumnCondCfg struct {
	Field         string              `json:"field,omitempty"`
	Operation     string              `json:"operation,omitempty"`
	SubConditions []*RowColumnCondCfg `json:"sub_conditions,omitempty"`
	ValueFrom     string              `json:"value_from,omitempty"`
	Value         any                 `json:"value,omitempty"`
}

type DataViewRowColumnAccount struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name,omitempty"`
}

// DataViewRowColumnRule 数据视图行列权限规则（与 mdl-data-model 返回体一致）。
type DataViewRowColumnRule struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	ViewID      string                   `json:"view_id"`
	ViewName    string                   `json:"view_name,omitempty"`
	Tags        []string                 `json:"tags"`
	Comment     string                   `json:"comment"`
	CreateTime  int64                    `json:"create_time"`
	UpdateTime  int64                    `json:"update_time"`
	Creator     DataViewRowColumnAccount `json:"creator"`
	Updater     DataViewRowColumnAccount `json:"updater"`
	Fields      []string                 `json:"fields"`
	RowFilters  *RowColumnCondCfg        `json:"row_filters"`
	Operations  []string                 `json:"operations,omitempty"`
}

// DataViewRowColumnRuleWrite 创建或更新请求体字段（与 mdl-data-model 请求 JSON 对齐）。
type DataViewRowColumnRuleWrite struct {
	ID         string            `json:"id,omitempty"`
	Name       string            `json:"name"`
	ViewID     string            `json:"view_id"`
	Tags       []string          `json:"tags,omitempty"`
	Comment    string            `json:"comment,omitempty"`
	Fields     []string          `json:"fields,omitempty"`
	RowFilters *RowColumnCondCfg `json:"row_filters,omitempty"`
}

// ListDataViewRowColumnRulesQuery 内部分页列表查询参数（缺省与 mdl-data-model 服务端 default 一致）。
type ListDataViewRowColumnRulesQuery struct {
	Name        string `query:"name"`
	NamePattern string `query:"name_pattern"`
	ViewID      string `query:"view_id"`
	Tag         string `query:"tag"`
	Offset      int    `query:"offset"`
	Limit       int    `query:"limit"`
	Sort        string `query:"sort"`
	Direction   string `query:"direction"`
}

// DefaultListDataViewRowColumnRulesQuery 返回 offset=0, limit=10, sort=update_time, direction=desc。
func DefaultListDataViewRowColumnRulesQuery() ListDataViewRowColumnRulesQuery {
	return ListDataViewRowColumnRulesQuery{
		Offset:    0,
		Limit:     10,
		Sort:      "update_time",
		Direction: "desc",
	}
}

// ListDataViewRowColumnRulesResult 列表接口返回体。
type ListDataViewRowColumnRulesResult struct {
	Entries    []*DataViewRowColumnRule `json:"entries"`
	TotalCount int                      `json:"total_count"`
}
