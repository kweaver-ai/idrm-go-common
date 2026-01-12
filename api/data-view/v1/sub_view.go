package v1

// Detail 定义行列规则（子视图）的具体定义
type Detail struct {
	// 字段列表
	Fields []Field `json:"fields,omitempty"`
	// 行过滤规则
	RowFilters *RowFilters `json:"row_filters,omitempty"`
}

// Field 定义字段
type Field struct {
	// 字段 ID
	ID string `json:"id,omitempty"`
	// 业务名称
	Name string `json:"name,omitempty"`
	// 技术名称
	NameEn string `json:"name_en,omitempty"`
	// 数据类型
	DataType DataType `json:"data_type,omitempty"`
}

// DataType 定义数据类型
type DataType string

const (
	// 数字型
	DataTypeNumber DataType = "number"
	// 字符型
	DataTypeChar DataType = "char"
	// 日期型
	DataTypeDate DataType = "date"
	// 日期时间型
	DataTypeDateTime DataType = "datetime"
	// 时间戳型
	DataTypeTimestamp DataType = "timestamp"
	// 布尔型
	DataTypeBool DataType = "bool"
	// 二进制
	DataTypeBinary DataType = "binary"
)

// RowFilters 定义行过滤规则
type RowFilters struct {
	// 条件组间关系
	WhereRelation Relation `json:"where_relation,omitempty"`
	// 条件组列表
	Where []Where `json:"where,omitempty"`
}

// Relation 定义关系
type Relation string

const (
	// And
	RelationAnd Relation = "and"
	// Or
	RelationOr Relation = "or"
)

type Where struct {
	// 限定对象列表
	Member []Member `json:"member,omitempty"`
	// 限定关系
	Relation Relation `json:"relation,omitempty"`
}

// Member 定义限定对象
type Member struct {
	// 字段
	Field
	// 限定条件
	Operator Operator
	// 限定比较值
	Value string
}

// Operator 定义限定条件
type Operator string

const ()
