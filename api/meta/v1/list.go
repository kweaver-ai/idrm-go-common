package v1

import (
	"net/url"
	"strconv"
)

type List[T any] struct {
	// 对象列表
	Entries []T `json:"entries"`
	// 总数量
	TotalCount int `json:"total_count"`
}

type ListOptions struct {
	// 页码，从 1 开始
	Offset int `json:"offset,omitempty" form:"offset"`
	// 每页显示的数量
	Limit int `json:"limit,omitempty" form:"limit"`
	// 排序所用字段
	Sort string `json:"sort,omitempty" form:"sort"`
	// 排序的方向
	Direction Direction `json:"direction,omitempty" form:"direction"`
}

// Deprecated: Use Convert_V1_ListOptions_To_url_Values instead.
func (opts *ListOptions) MarshalQuery() (url.Values, error) {
	q := make(url.Values)

	if opts.Offset != 0 {
		q.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit != 0 {
		q.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Sort != "" {
		q.Set("sort", opts.Sort)
	}
	if opts.Direction != "" {
		q.Set("direction", string(opts.Direction))
	}

	return q, nil
}

// Deprecated: Use Convert_url_Values_To_V1_ListOptions instead.
func (opts *ListOptions) UnmarshalQuery(data url.Values) (err error) {
	for k, values := range data {
		for _, v := range values {
			switch k {
			case "offset":
				if opts.Offset, err = strconv.Atoi(v); err != nil {
					return
				}
			case "limit":
				if opts.Limit, err = strconv.Atoi(v); err != nil {
					return
				}
			case "sort":
				opts.Sort = v
			case "direction":
				opts.Direction = Direction(v)
			default:
				continue
			}
		}
	}
	return nil
}

// 排序方向
type Direction string

const (
	// 升序
	Ascending Direction = "asc"
	// 降序
	Descending Direction = "desc"
)
