package v1

// DataView 代表 AnyRobot 的数据视图
//
// 仅包含 AnyFabric 用到的字段，缺少的字段在需要时再补充。
type DataView struct {
	// ID
	ID string `json:"id,omitempty"`
	// 逻辑视图拥有的字段。逻辑视图的字段范围是“选择全部字段”时，这些字段来自于
	// 已经记录的数据。
	Fields []Field `json:"fields,omitempty"`
}

// Field 代表 AnyRobot 数据视图的字段
type Field struct {
	// 名称
	Name string
}
