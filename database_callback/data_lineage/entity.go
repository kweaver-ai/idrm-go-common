package data_lineage

import (
	"fmt"

	"github.com/kweaver-ai/idrm-go-common/util"
)

type EntityMsgMessage struct {
	Header  EntityMsgMessageHeader  `json:"header"`
	Payload EntityMsgMessagePayload `json:"payload"`
}

type EntityMsgMessageHeader struct{}

type EntityMsgMessagePayload struct {
	Type    string  `json:"type" binding:"required,oneof="sql lineage"` //实体变更的类型
	Content Content `json:"content"`
}

type Content struct {
	Type      string `json:"type" binding:"required,oneof="create update delete"`        //实体变更的类型
	ClassName string `json:"class_name" binding:"required,oneof="table field indicator"` //实体类的类型
	Entities  []any  `json:"entities"`
}

// LineageProcessModel 数据加工采集模型
type LineageProcessModel struct {
	ID            string `json:"id"`
	SourceTableID string `json:"source_table_id"`
	TargetTableID string `json:"target_table_id"`
	CreateSQL     string `json:"create_sql"`
	InsertSQL     string `json:"insert_sql"`
}

// LineageTable   血缘图谱的表
type LineageTable struct {
	UniqueID          string `json:"unique_id"`
	UUID              string `json:"uuid"`
	BusinessName      string `json:"business_name"`
	TechnicalName     string `json:"technical_name"`
	Comment           string `json:"comment"`
	TableType         string `json:"table_type"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	TaskExecutionInfo string `json:"task_execution_info"`
	DatasourceID      string `json:"datasource_id"`
	DatasourceName    string `json:"datasource_name"`
	OwnerId           string `json:"owner_id"`
	OwnerName         string `json:"owner_name"`
	DepartmentID      string `json:"department_id"`
	DepartmentName    string `json:"department_name"`
	InfoSystemID      string `json:"info_system_id"`
	InfoSystemName    string `json:"info_system_name"`
	DatabaseName      string `json:"database_name"`
	CatalogName       string `json:"catalog_name"`
	CatalogAddr       string `json:"catalog_addr"`
	CatalogType       string `json:"catalog_type"`
	DataViewSource    string `json:"-"`
	SceneID           string `json:"-"`
}

func (t *LineageTable) GenUniqueID() string {
	return util.MD5(fmt.Sprintf("%s%s%s", t.CatalogName, t.DatabaseName, t.TechnicalName))
}

type LineageField struct {
	UniqueID        string `json:"unique_id"`
	UUID            string `json:"uuid"`
	BusinessName    string `json:"business_name"`
	TechnicalName   string `json:"technical_name"`
	Comment         string `json:"comment"`
	DataType        string `json:"data_type"`
	PrimaryKey      int8   `json:"primary_key"`
	TableUniqueID   string `json:"table_unique_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	ExpressionName  string `json:"expression_name"`
	Expression      string `json:"expression"`
	ColumnUniqueIDS string `json:"column_unique_ids"`
}

type LineageIndicator struct {
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Code            string `json:"code"`
	IndicatorType   string `json:"indicator_type"`
	Expression      string `json:"expression"`        //表达式
	ExecSQL         string `json:"exec_sql"`          //执行SQL
	IndicatorUUIDS  string `json:"indicator_uuids"`   //引用的指标ID
	ColumnUniqueIDs string `json:"column_unique_ids"` //依赖的字段的uuid
	OwnerUID        string `json:"owner_uid"`
	OwnerName       string `json:"owner_name"`
	DepartmentID    string `json:"department_id"`
	DepartmentName  string `json:"department_name"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// region fetcherModel

// DepartmentInfo 部门信息  来自DepartmentInternal
type DepartmentInfo struct {
	ID     string `json:"id" `     // 对象ID
	Name   string `json:"name"`    // 对象名称
	Type   string `json:"type"`    // 对象类型
	Path   string `json:"path"`    // 对象路径
	PathID string `json:"path_id"` // 对象ID路径
}

// InfoSystemInfo 信息系统
type InfoSystemInfo struct {
	ID          string `json:"id"`          // 信息系统业务id
	Name        string `json:"name"`        // 信息系统名称
	Description string `json:"description"` // 信息系统描述
}

// UserInfo 用户信息
type UserInfo struct {
	ID   string `json:"id"`   // 用户ID
	Name string `json:"name"` // 用户名称
}

// DataSource 数据源信息
type DataSource struct {
	DataSourceID   uint64 `json:"data_source_id"`   // 数据源雪花id
	ID             string `json:"id"`               // 数据源业务id
	InfoSystemID   string `json:"info_system_id"`   // 信息系统id
	Name           string `json:"name"`             // 数据源名称
	CatalogName    string `json:"catalog_name"`     // 数据源catalog名称
	Type           int32  `json:"type"`             // 数据库类型
	TypeName       string `json:"type_name"`        // 数据库类型名称
	Host           string `json:"host"`             // 连接地址
	Port           int32  `json:"port"`             // 端口
	Username       string `json:"username"`         // 用户名
	DatabaseName   string `json:"database_name"`    // 数据库名称
	Schema         string `json:"schema"`           // 数据库模式
	SourceType     int32  `json:"source_type"`      // 数据源类型 1:记录型、2:分析型
	DataViewSource string `json:"data_view_source"` // 数据视图源
}

// TableInfo 表的信息，应该兼容其他任何表和视图
type TableInfo struct {
	ID            string `json:"id"`
	TechnicalName string `json:"technical_name"` // 表技术名称
	BusinessName  string `json:"business_name"`  // 表业务名称
	Type          int32  `json:"type"`           // 逻辑视图来源 1：数据源、2：用户、3：应用
	DatasourceID  string `json:"datasource_id"`  // 数据源id
	DatabaseName  string `json:"database_name"`  //数据库名称
	CatalogName   string `json:"catalog_name"`   //数据catalog
	SceneID       string `json:"-"`              //场景分析ID
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (t *TableInfo) GenUniqueID() string {
	return util.MD5(fmt.Sprintf("%s%s%s", t.CatalogName, t.DatabaseName, t.TechnicalName))
}

const DataLineageTimeFormat = "2006-01-02 15:04:05.000"

//endregion
