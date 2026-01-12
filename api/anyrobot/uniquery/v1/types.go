package v1

import (
	"encoding/json"
	"time"

	v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// DataViewQuery 定义一次对 DataView 的查询
type DataViewQuery struct {
	// 过滤条件
	Filters *DataViewQueryFilters `json:"filters,omitempty"`

	// 过滤不早于这个时间戳，单位：毫秒
	Start *v1.TimestampUnixMilli `json:"start,omitempty"`
	// 过滤不晚于这个时间戳，单位：毫秒
	End *v1.TimestampUnixMilli `json:"end,omitempty"`

	// 排序字段，默认是 @timestamp
	Sort Sort `json:"sort,omitempty"`
	// 排序结果方向，可选 asc、desc。默认 desc
	Direction Direction `json:"direction,omitempty"`

	// 偏移量
	Offset int `json:"offset,omitempty"`
	// 单页返回的数量
	Limit int `json:"limit,omitempty"`
	// 返回数据的格式
	Format DataViewQueryFormat `json:"format,omitempty"`
}

// DataViewQueryFilters 定义对 DataView 查询的过滤条件
type DataViewQueryFilters struct {
	// 操作符
	Operation Operation `json:"operation,omitempty"`
	// 字段
	Field string `json:"field,omitempty"`
	// 值
	Value any `json:"value,omitempty"`
	// 值的来源
	ValueFrom ValueFrom `json:"value_from,omitempty"`
	// 子条件，仅当 Operation 是逻辑操作符是有值
	SubConditions []DataViewQueryFilters `json:"sub_conditions,omitempty"`
}

// Operation 定义过滤条件的操作符
type Operation string

// TODO: 补充其他操作符
const (
	// 逻辑操作符：AND
	OperationAnd Operation = "and"
	// 逻辑操作符：OR
	OperationOr Operation = "or"

	// 单目运算符：字段不存在
	OperationNotExist Operation = "not_exist"

	// 比较操作符：等于
	OperationEqual Operation = "=="
	// 比较操作符：相似
	OperationLike Operation = "like"
	// 比较操作符：属于
	OperationIn Operation = "in"
)

// ValueFrom 定义值的来源
type ValueFrom string

const (
	// 常量
	ValueFromConst ValueFrom = "const"
	// 引用自字段
	ValueFromField ValueFrom = "field"
	// 引用自用户
	ValueFromUser ValueFrom = "user"
)

// 排序字段
type Sort string

const (
	SortTimestamp Sort = "@timestamp"
)

// 排序结果方向
type Direction string

const (
	Ascending  Direction = "asc"
	Descending Direction = "desc"
)

// DataViewQueryFormat 定义对 DataView 查询的格式
type DataViewQueryFormat string

const (
	// 展开为键值对
	DataViewQueryFlat DataViewQueryFormat = "flat"
	// 原始 json 结构
	DataViewQueryOriginal DataViewQueryFormat = "original"
)

// DataViewQueryOptions 定义对 DataView 查询的 query 参数
type DataViewQueryOptions struct {
	AllowNonExistField bool `json:"allow_non_exist_field,omitempty"`
}

// 视图查询外部接口统一返回结构
//
// 仅包含使用到的部分字段，其他字段在需要时再补充。
type ViewUniResponse struct {
	Datas []ViewData `json:"datas,omitempty"`
}

type ViewData struct {
	Total  string  `json:"total,omitempty"`
	Values []Value `json:"values,omitempty"`
}

type Value struct {
	Timestamp    time.Time       `json:"@timestamp,omitempty"`
	Body         json.RawMessage `json:"Body,omitempty"`
	SeverityText string          `json:"SeverityText,omitempty"`
}

// Status 代表 AnyRobot 返回正常数据的结构
type Status struct {
	ErrorCode string `json:"error_code,omitempty"`

	Description string `json:"description,omitempty"`

	Solution string `json:"solution,omitempty"`

	ErrorLink string `json:"error_link,omitempty"`

	ErrorDetails string `json:"error_details,omitempty"`
}
