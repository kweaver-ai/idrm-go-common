package v1

import (
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// InfoSystemSearch 代表搜索信息系统的请求
type InfoSystemSearch struct {
	// 过滤器，未指定代表不过滤
	Filter *InfoSystemSearchFilter `json:"filter,omitempty"`
	// limit 时搜索返回的最大响应数量。如果存在更多结果，服务端将返回 continue字
	// 段，客户端应使用 continue 字段的存在与否判断是否有更多可用结果。如果指定
	// 了 limit 且返回的 continue 为空，客户端可以认为没有更多结果。
	Limit int `json:"limit,omitempty"`
	// 从服务端搜索更多结果时，应设置 continue 参数。此值由服务端定义，客户端只
	// 能使用先前查询结果中的 continue 值，并且服务器可能会拒绝其无法识别的
	// continue 值。
	Continue string `json:"continue,omitempty"`
}

// InfoSystemSearchFilter 代表搜索信息系统的过滤器
type InfoSystemSearchFilter struct {
	// 关键字，非空时根据信息系统的名称、描述过滤
	Keyword string `json:"keyword,omitempty"`
	// 部门 ID，过滤属于指定部门及其子部门的信息系统。未指定代表不过滤。空字符串
	// 代表过滤未属于任何部门的信息系统。
	DepartmentID *string `json:"department_id,omitempty"`
}

// InfoSystemSearchResult 代表搜索信息系统的响应
type InfoSystemSearchResult struct {
	// 匹配到的信息系统的总数
	Total Total `json:"total,omitempty"`
	// 匹配到的信息系统的列表
	Entries []InfoSystemWithHighlight `json:"entries,omitempty"`
	// 如果客户端设置了返回数量的限制，则可能设置 continue 值，这表示服务端有更
	// 多可用的数据。该值不透明，可用于向服务器发送另一个请求，以检索下一组可用
	// 数据。
	Continue string `json:"continue,omitempty"`
}

// InfoSystem 代表一条信息系统的搜索结果
type InfoSystem struct {
	// ID
	ID string `json:"id,omitempty"`
	// 更新时间
	UpdatedAt meta_v1.Time `json:"updated_at,omitempty"`
	// 名称，带有高亮标签的名称
	Name string `json:"name,omitempty"`
	// 描述，带有高亮标签的描述
	Description string `json:"description,omitempty"`
	// 所属部门的 ID，未指定、空字符串代表不属于任何部门。
	DepartmentID string `json:"department_id,omitempty"`
	// 所属部门的完整路径，未指定、空字符传代表不属于任何部门。
	DepartmentPath string `json:"department_path,omitempty"`
}

// InfoSystemWithHighlight 代表包含高亮标签的信息系统
type InfoSystemWithHighlight struct {
	InfoSystem
	// 带有高亮标签的名称
	NameHighlight string `json:"name_highlight,omitempty"`
	// 带有高亮标签的描述
	DescriptionHighlight string `json:"description_highlight,omitempty"`
}
