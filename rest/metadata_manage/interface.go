package metadata_manage

import (
	"context"
)

type Driven interface {
	SendLineage(ctx context.Context, data any) error
	//GetLineageRelationData  查询血缘的关系数据，返回的单纯是血缘的关系，只有ID和类型
	GetLineageRelationData(ctx context.Context, req *QueryLineageReqParams) (*QueryLineageResp, error)
	//GetLineageData  查询血缘实体信息，entity_type: table,column,indicator; id可以支持逗号分割的数组
	GetLineageData(ctx context.Context, entityType string, id ...string) (*QueryLineageTableResp, error)
	//GetLineageColumns 查询血缘表的字段ID，没有的返回空数组
	GetLineageColumns(ctx context.Context, tableID string) ([]string, error)

	SyncTableInfo(ctx context.Context, req *PayloadTableReq) (*SyncLineageResp, error)
	SyncColumnDetail(ctx context.Context, req *PayloadColumnReq) (*SyncLineageResp, error)
	SyncTaskTableInfo(ctx context.Context, req *PayloadTaskTableReq) (*SyncLineageResp, error)
	SyncTaskColumnDetail(ctx context.Context, req *PayloadTaskColumnReq) (*SyncLineageResp, error)
}

// QueryLineageNodes region

type QueryLineageReqParams struct {
	ID        string `query:"ids"`       //节点的ID
	Step      string `query:"step"`      //跳数，包含首尾节点
	Direction string `query:"direction"` //方向，forward: 查找子节点；backward查找父节点
}

type QueryLineageResp struct {
	Edges map[string][]*LineageEdge `json:"edges"` //边的数组
}

type LineageEdge struct {
	Id      string   `json:"id"`     //实体的unique_id
	Type    int      `json:"type"`   //实体类型，1:column, 2:indicator
	Parent  string   `json:"parent"` //父节点的unique_id数组
	Parents []string `json:"parents"`
	Child   string   `json:"child"` //子节点的unique_id数组
	Childs  []string `json:"childs"`
}

// Column 数据血缘的字段
type Column struct {
	UniqueId        string `json:"unique_id"`         //列的唯一id
	Uuid            string `json:"uuid"`              //字段的uuid
	TechnicalName   string `json:"technical_name"`    //列技术名称
	BusinessName    string `json:"business_name"`     //列业务名称
	Comment         string `json:"comment"`           //字段注释
	DataType        string `json:"data_type"`         //字段的数据类型
	PrimaryKey      string `json:"primary_key"`       //是否主键
	TableUniqueID   string `json:"table_unique_id"`   //属于血缘表的uuid
	ExpressionName  string `json:"expression_name"`   //column的生成表达式
	ColumnUniqueIDs string `json:"column_unique_ids"` //column的生成依赖的column的uid
	ActionType      string `json:"action_type"`       //操作类型:insert,update,delete
	CreatedAt       string `json:"created_at"`        //创建时间
	UpdatedAt       string `json:"updated_at"`        //更新时间
}

func (c *Column) ShellID() string {
	return c.TableUniqueID
}

// Table 业务表信息
type Table struct {
	UniqueId          string `json:"unique_id"`           //表的uuid
	Uuid              string `json:"uuid"`                //唯一id
	TechnicalName     string `json:"technical_name"`      //表技术名称
	BusinessName      string `json:"business_name"`       //表业务名称
	Comment           string `json:"comment"`             //表注释
	TableType         string `json:"table_type"`          //表类型
	TaskExecutionInfo string `json:"task_execution_info"` //任务执行信息
	DatasourceId      string `json:"datasource_id"`       //数据源id
	DatasourceName    string `json:"datasource_name"`     //数据源名称
	OwnerId           string `json:"owner_id"`            //数据Ownerid
	OwnerName         string `json:"owner_name"`          //数据OwnerName
	DepartmentId      string `json:"department_id"`       //所属部门id
	DepartmentName    string `json:"department_name"`     //所属部门mame
	InfoSystemId      string `json:"info_system_id"`      //信息系统id
	InfoSystemName    string `json:"info_system_name"`    //信息系统名称
	DatabaseName      string `json:"database_name"`       //数据库名称
	CatalogName       string `json:"catalog_name"`        //数据源catalog名称
	CatalogAddr       string `json:"catalog_addr"`        //数据源地址
	CatalogType       string `json:"catalog_type"`        //数据库类型名称
	ActionType        string `json:"action_type"`         //操作类型:insert,update,delete
	CreatedAt         string `json:"created_at"`          //创建时间
	UpdatedAt         string `json:"updated_at"`          //更新时间
}

func (c *Table) ShellID() string {
	return ""
}

type Indicator struct {
	Uuid             string `json:"uuid"`              //指标的uuid
	Name             string `json:"name"`              //指标名称
	Description      string `json:"description"`       //指标名称描述
	Code             string `json:"code"`              //指标编号
	IndicatorType    string `json:"indicator_type"`    //指标类型:atomic原子derived衍生composite复合
	Expression       string `json:"expression"`        //指标表达式，如果指标是原子或复合指标时
	IndicatorUuids   string `json:"indicator_uuids"`   //引用的指标uuid
	TimeRestrict     string `json:"time_restrict"`     //时间限定表达式，如果指标是衍生指标时
	ModifierRestrict string `json:"modifier_restrict"` //普通限定表达式，如果指标是衍生指标时
	OwnerUid         string `json:"owner_uid"`         //数据ownerID
	OwnerName        string `json:"owner_name"`        //数据owner名称
	DepartmentId     string `json:"department_id"`     //所属部门id
	DepartmentName   string `json:"department_name"`   //所属部门名称
	ColumnUniqueIDs  string `json:"column_unique_ids"` //依赖的字段的uuid
	ActionType       string `json:"action_type"`       //操作类型:insert,update,delete
	CreatedAt        string `json:"created_at"`        //创建时间
	UpdatedAt        string `json:"updated_at"`        //更新时间
}

func (c *Indicator) ShellID() string {
	return c.Uuid
}

//endregion

// QueryLineageTableReq region

type QueryLineageTableReq struct {
	ID         string `query:"ids"`  //表的UUID，逗号分割的字符串
	EntityType string `query:"type"` //实体的类型
}

type QueryLineageTableResp struct {
	Entries []any `json:"entries"`
}

//endregion

// region QueryLineageTableColumnReq

type QueryLineageTableColumnReq struct {
	ID string `uri:"id"`
}

type QueryLineageTableColumnResp struct {
	Data string `json:"data"`
}

//endregion

// 第三方血缘同步k开始

type PayloadTableReq struct {
	Type       string            `json:"type" binding:"required,max=10,oneof=batch stream" example:"stream"`                //离线或实时类型（batch离线初始化、stream实时）
	ActionType string            `json:"action_type"  binding:"required,max=6,oneof=insert delete update" example:"insert"` //操作类型（insert新增、delete删除、update更新）
	DbType     *int              `json:"db_type"  binding:"required,max=1,oneof=0 1" example:"0"`                           //数据库类型1私有（不与AnyFabric共用数仓）、0共用（与AnyFabric共用数仓）
	Entities   []*TableEntityReq `json:"entities"  binding:"gte=0,lte=100,required,dive"`                                   //表信息的数组
}

type TableEntityReq struct {
	DataSourceId string `json:"dataSourceId"  binding:"omitempty,max=36" example:"11112222222222rr"`      //第三方数据源ID 注：会根据dataSourceId和（type、host、port、database_name）不能同时为空
	Type         string `json:"type"  binding:"omitempty,max=36" example:"hive"`                          //table的数据库类型（hive、mysql、clickhouse等）
	Host         string `json:"host"  binding:"omitempty,max=15" example:"127.0.0.1"`                     //数据源ip
	Port         *int   `json:"port"  binding:"omitempty,max=65535,lte=65535" example:"3306"`             //数据源端口
	DatabaseName string `json:"database_name" binding:"omitempty,max=100" example:"dwd"`                  //数据库名(英文)
	SchemaName   string `json:"schema_name"  binding:"omitempty,max=100" example:"public"`                //schema名(英文)根据数据库类型传递,如：postgre有schema
	TableName    string `json:"table_name"  binding:"required,max=100" example:"t_user"`                  //数据表名(英文)
	Comment      string `json:"comment" binding:"omitempty,max=1000" example:"用户表"`                       //数据表描述
	CreatedAt    string `json:"created_at"  binding:"omitempty,max=25" example:"2024-05-29 13:15:12.149"` //创建时间（原表创建时间），为空时自动填充接入时间
	UpdatedAt    string `json:"updated_at"  binding:"omitempty,max=25" example:"2024-05-29 13:15:12.149"` //更新时间（原表更新时间），为空时自动填充接入时间
}

type PayloadColumnReq struct {
	Type       string             `json:"type" binding:"required,max=10,oneof=batch stream" example:"stream"`                //离线或实时类型（batch离线初始化、stream实时）
	ActionType string             `json:"action_type"  binding:"required,max=6,oneof=insert delete update" example:"insert"` //操作类型（insert新增、delete删除、update更新）
	DbType     *int               `json:"db_type"  binding:"required,max=1,oneof=0 1" example:"0"`                           //数据库类型1私有（不与AnyFabric共用数仓）、0共用（与AnyFabric共用数仓）
	Entities   []*ColumnEntityReq `json:"entities"  binding:"gte=0,lte=500,required,dive"`                                   //字段信息的数组
}
type ColumnEntityReq struct {
	DataSourceId string `json:"dataSourceId"  binding:"omitempty,max=36" example:"11112222222222rr"`      //第三方数据源ID 注：会根据dataSourceId和（type、host、port、database_name）不能同时为空
	Type         string `json:"type"  binding:"omitempty,max=36" example:"hive"`                          //table的数据库类型（hive、mysql、clickhouse等）
	Host         string `json:"host"  binding:"omitempty,max=15" example:"127.0.0.1"`                     //数据源ip
	Port         *int   `json:"port"  binding:"omitempty,max=65535,lte=65535" example:"3306"`             //数据源端口
	DatabaseName string `json:"database_name" binding:"omitempty,max=100" example:"dwd"`                  //数据库名(英文)
	SchemaName   string `json:"schema_name"  binding:"omitempty,max=100" example:"public"`                //schema名(英文)根据数据库类型传递,如：postgre有schema
	TableName    string `json:"table_name"  binding:"required,max=100" example:"t_user"`                  //数据表名(英文)
	PrimaryKey   int    `json:"primary_key"  binding:"omitempty,max=1,oneof=0 1" default:"0" example:"1"` //是否主键,1是0否
	ColumnName   string `json:"column_name"  binding:"required,max=100" example:"f_name"`                 //字段名(英文)
	DataType     string `json:"data_type"  binding:"required,max=36" example:"varchar(10)"`               //字段类型varchar(10)、int(10)等
	Comment      string `json:"comment"  binding:"omitempty,max=1000" example:"名称"`                       //字段描述
	CreatedAt    string `json:"created_at"  binding:"omitempty,max=25" example:"2024-05-29 13:15:12.149"` //创建时间（原表创建时间），为空时自动填充接入时间
	UpdatedAt    string `json:"updated_at"  binding:"omitempty,max=25" example:"2024-05-29 13:15:12.149"` //更新时间（原表更新时间），为空时自动填充接入时间
}

type TaskTableEntityReq struct {
	DataSourceId      string `json:"dataSourceId"  binding:"omitempty,max=36" example:"11112222222222rr"`         //第三方数据源ID 注：会根据dataSourceId和（type、host、port、database_name）不能同时为空
	Type              string `json:"type"  binding:"omitempty,max=36" example:"hive"`                             //table的数据库类型（hive、mysql、clickhouse等）
	Host              string `json:"host"  binding:"omitempty,max=15" example:"127.0.0.1"`                        //目标数据源ip
	Port              *int   `json:"port"  binding:"omitempty,max=65535,lte=65535" example:"3306"`                //目标数据源端口
	DatabaseName      string `json:"database_name" binding:"omitempty,max=100" example:"dwd"`                     //目标数据库名(英文)
	SchemaName        string `json:"schema_name"  binding:"omitempty,max=100" example:"public"`                   //schema名(英文)根据数据库类型传递,如：postgre有schema
	TableName         string `json:"table_name"  binding:"required,max=100" example:"t_user"`                     //目标数据表名(英文)
	Comment           string `json:"comment" binding:"omitempty,max=1000" example:"用户表"`                          //目标数据表描述
	CreatedAt         string `json:"created_at"  binding:"omitempty,max=25" example:"2024-05-29 13:15:12.149"`    //创建时间（原表创建时间），为空时自动填充接入时间
	UpdatedAt         string `json:"updated_at"  binding:"omitempty,max=25" example:"2024-05-29 13:15:12.149"`    //更新时间（原表更新时间），为空时自动填充接入时间
	TaskExecutionInfo string `json:"task_execution_info"  binding:"required,max=128" example:"dwd_insert_user加工"` //表加工任务的相关名称
}

type PayloadTaskTableReq struct {
	Type       string                `json:"type" binding:"required,max=10,oneof=batch stream" example:"stream"`                //离线或实时类型（batch离线初始化、stream实时）
	ActionType string                `json:"action_type"  binding:"required,max=6,oneof=insert delete update" example:"insert"` //操作类型（insert新增、delete删除、update更新）
	DbType     *int                  `json:"db_type"  binding:"required,max=1,oneof=0 1" example:"0"`                           //数据库类型1私有（不与AnyFabric共用数仓）、0共用（与AnyFabric共用数仓）
	Entities   []*TaskTableEntityReq `json:"entities"  binding:"gte=0,lte=100,required,dive"`                                   //任务的表信息数组
}

type ColumnDependencyReq struct {
	DataSourceId string `json:"dataSourceId"  binding:"omitempty,max=36" example:"11112222222222rr"` //第三方数据源ID 注：会根据dataSourceId和（type、host、port、database_name）不能同时为空
	Type         string `json:"type"  binding:"omitempty,max=36" example:"hive"`                     //table的数据库类型（hive、mysql、clickhouse等）
	Host         string `json:"host"  binding:"omitempty,max=15" example:"127.0.0.1"`                //数据源ip
	Port         *int   `json:"port"  binding:"omitempty,max=65535,lte=65535" example:"3306"`        //数据源端口
	DatabaseName string `json:"database_name" binding:"omitempty,max=100" example:"dwd"`             //数据库名(英文)
	SchemaName   string `json:"schema_name"  binding:"omitempty,max=100" example:"public"`           //schema名(英文)根据数据库类型传递,如：postgre有schema
	TableName    string `json:"table_name"  binding:"required,max=100" example:"t_user"`             //数据表名(英文)
	ColumnName   string `json:"column_name"  binding:"required,max=100" example:"f_name"`            //字段名(英文)
}

type TaskColumnEntityReq struct {
	DataSourceId   string                 `json:"dataSourceId"  binding:"omitempty,max=36" example:"11112222222222rr"`                                                               //第三方数据源ID 注：会根据dataSourceId和（type、host、port、database_name）不能同时为空
	Type           string                 `json:"type"  binding:"omitempty,max=36" example:"hive"`                                                                                   //table的数据库类型（hive、mysql、clickhouse等）
	Host           string                 `json:"host"  binding:"omitempty,max=15" example:"127.0.0.1"`                                                                              //数据源ip
	Port           *int                   `json:"port"  binding:"omitempty,max=65535,lte=65535" example:"3306"`                                                                      //数据源端口
	DatabaseName   string                 `json:"database_name" binding:"omitempty,max=100" example:"dwd"`                                                                           //数据库名(英文)
	SchemaName     string                 `json:"schema_name"  binding:"omitempty,max=100" example:"public"`                                                                         //schema名(英文)根据数据库类型传递,如：postgre有schema
	TableName      string                 `json:"table_name"  binding:"required,max=100" example:"t_user"`                                                                           //数据表名(英文)
	ColumnName     string                 `json:"column_name"  binding:"required,max=100" example:"f_name"`                                                                          //字段名(英文)
	CreatedAt      string                 `json:"created_at"  binding:"omitempty,max=25" example:"2024-05-29 13:15:12.149"`                                                          //创建时间（原表创建时间），为空时自动填充接入时间
	UpdatedAt      string                 `json:"updated_at"  binding:"omitempty,max=25" example:"2024-05-29 13:15:12.149"`                                                          //更新时间（原表更新时间），为空时自动填充接入时间
	ExpressionName string                 `json:"expression_name" binding:"omitempty,max=65535" example:"CASE WHEN salary BETWEEN 1000 AND 5000 THEN A ELSE D END AS salary_grade;"` //任务脚本（生成表达式或加工口径或sql）
	ColumnEtlDeps  []*ColumnDependencyReq `json:"column_etl_deps"  binding:"gte=0,lte=100,required,dive"`                                                                            //来源表字段合并运行变量（依赖的列）
}

type PayloadTaskColumnReq struct {
	Type       string                 `json:"type" binding:"required,max=10,oneof=batch stream" example:"stream"`                //离线或实时类型（batch离线初始化、stream实时）
	ActionType string                 `json:"action_type"  binding:"required,max=6,oneof=insert delete update" example:"insert"` //操作类型（insert新增、delete删除、update更新）
	DbType     *int                   `json:"db_type"  binding:"required,max=1,oneof=0 1" example:"0"`                           //数据库类型1私有（不与AnyFabric共用数仓）、0共用（与AnyFabric共用数仓）
	Entities   []*TaskColumnEntityReq `json:"entities"  binding:"gte=0,lte=500,required,dive"`                                   //任务的字段信息的数组
}

type SyncLineageResp struct {
	Code        string `json:"code"  binding:"required" example:"0"` //成功为0，否则失败
	Description string `json:"description"`                          //描述
	Detail      string `json:"detail"`                               //详情
	Solution    string `json:"solution"`                             //错误原因
}

// 第三方血缘同步k开始
