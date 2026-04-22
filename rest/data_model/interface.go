package data_model

import "context"

type Driven interface {
	GetDataModelByID(ctx context.Context, ids ...string) ([]*DataModel, error)
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
	Creator        string   `json:"creator"`
	Updater        string   `json:"updater"`
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
