package v1

type Rule struct {
	// 名称
	Name string `json:"name,omitempty"`
	//固定范围的字段，
	ScopeFields []string `json:"scope_fields,omitempty"`
	// 列、字段列表
	Fields []Field `json:"fields,omitempty"`
	// 行过滤规则，nil 代表没有过滤规则
	RowFilters *RowFilters `json:"row_filters,omitempty"`
	// 固定的行过滤规则
	FixedRowFilters *RowFilters `json:"fixed_row_filters,omitempty"`
}

// 列、字段
type Field struct {
	// 字段 ID
	ID string `json:"id,omitempty"`
	// 中文名称、业务名称
	Name string `json:"name,omitempty"`
	// 英文名称、技术名称
	NameEn string `json:"name_en,omitempty"`
	// 数据类型,由逻辑视图定义。
	DataType string `json:"data_type,omitempty"`
}

// 行过滤条件
type RowFilters struct {
	// 条件组间关系
	WhereRelation string `json:"where_relation,omitempty"`
	// 条件组列表
	//
	// 前端要求空列表序列化 json 为 `[]` 所以不能使用 json tag omitempty
	Where []Where `json:"where"`
}

// 过滤条件
type Where struct {
	// 限定对象
	Member []Member `json:"member,omitempty"`
	// 限定关系
	Relation string `json:"relation,omitempty"`
}

type Member struct {
	Field `json:",inline"`
	// 限定条件
	Operator string `json:"operator,omitempty"`
	// 限定比较值
	Value string `json:"value,omitempty"`
}
