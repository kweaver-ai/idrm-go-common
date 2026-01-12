package v1

// Total 代表搜索结果的总数
type Total struct {
	// 匹配到的数据额数量
	Value int `json:"value,omitempty"`
	// Value 的类型
	Relation TotalRelation `json:"relation,omitempty"`
}

type TotalRelation string

const (
	// 精确计数
	TotalEqual TotalRelation = "eq"
	// 下限估算值
	TotalGreaterThanOrEqual TotalRelation = "gte"
)
