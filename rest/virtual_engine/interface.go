package virtual_engine

import (
	"context"
)

type Driven interface {
	DatabaseTypeMapping(ctx context.Context, req *DatabaseTypeMappingReq) (*DatabaseTypeMappingResp, error)
	TableFieldsDetail(ctx context.Context, catalog, schema, table string) ([]DataTableFieldResp, error)
	DBConnectorConfig(ctx context.Context, connector string) (*ConnectorConfigResp, error)
	GetPreview(ctx context.Context, req *ViewEntries) (*FetchDataRes, error)
}

type DatabaseTypeMappingReq struct {
	SourceConnectorName string              `json:"sourceConnectorName"`
	TargetConnectorName string              `json:"targetConnectorName"`
	Type                []SourceFieldObject `json:"type"`
}

type DatabaseTypeMappingResp struct {
	TargetConnectorName string              `json:"targetConnectorName"`
	Type                []TargetFieldObject `json:"type"`
}

type SourceFieldObject struct {
	Index          int    `json:"index"`
	SourceTypeName string `json:"sourceTypeName"`
	Precision      int32  `json:"precision,omitempty"`
	DecimalDigits  *int32 `json:"decimalDigits,omitempty"`
}

type TargetFieldObject struct {
	Index          int    `json:"index"`
	TargetTypeName string `json:"targetTypeName"`
	Precision      int    `json:"precision,omitempty"`
	DecimalDigits  *int   `json:"decimalDigits,omitempty"`
}

type TableWithFieldsResp struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	OrigType      string `json:"origType"`
	Description   string `json:"comment"`
	PrimaryKey    bool   `json:"primaryKey"`
	NullAble      bool   `json:"nullAble"`
	TypeSignature `json:"typeSignature"`
}

type TypeSignature struct {
	RawType   string `json:"rawType"`
	Arguments []struct {
		Kind  string      `json:"kind"`
		Value interface{} `json:"value"`
	} `json:"arguments"`
}

type DataTableFieldResp struct {
	Name           string `json:"name"`            //字段名
	Type           string `json:"type"`            //字段类型
	Description    string `json:"description"`     //描述
	RawType        string `json:"rawType"`         //原生类型
	OrigType       string `json:"origType"`        //来源类型
	Length         *int   `json:"length"`          // 字段长度
	FieldPrecision *int   `json:"field_precision"` // 字段精度
}

type ConnectorConfigResp struct {
	ConnectorName string                   `json:"connectorName"`
	SchemaExist   bool                     `json:"schemaExist"`
	URL           string                   `json:"url"`
	Type          []*ConnectorConfigColumn `json:"type"`
}

type ConnectorConfigColumn struct {
	SourceType    string `json:"sourceType"`
	OlkSearchType string `json:"olkSearchType"`
	OlkWriteType  string `json:"olkWriteType"`
	MinTypeLength int    `json:"minTypeLength"`
	MaxTypeLength int    `json:"maxTypeLength"`
}

type ViewEntries struct {
	CatalogName string `json:"catalogName"`
	ViewName    string `json:"viewName"`
	Schema      string `json:"schema"`
	Limit       int    `json:"limit"`
	UserId      string `json:"user_id"`
}

type FetchDataRes struct {
	TotalCount int       `json:"total_count"`
	Columns    []*Column `json:"columns"`
	Data       [][]any   `json:"data"`
}
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
