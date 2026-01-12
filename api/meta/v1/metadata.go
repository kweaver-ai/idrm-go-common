package v1

// 元数据
type Metadata struct {
	// ID
	ID string `json:"id,omitempty"`
	// 创建时间
	CreatedAt Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt Time `json:"updated_at,omitempty"`
	// 删除时间。当用户请求优雅删除时，服务器会设置这个字段，并不是由用户直接设
	// 置。
	DeletedAt *Time `json:"deleted_at,omitempty"`
}

// 元数据，包含创建、更新、删除这个数据的用户的 ID
type MetadataWithOperator struct {
	Metadata
	// 创建人 ID
	CreatedBy string `json:"created_by,omitempty"`
	// 更新人 ID
	UpdatedBy string `json:"updated_by,omitempty"`
	// 删除人 ID。当用户请求优雅删除时，服务器会设置这个字段，并不是由用户直接设
	// 置。
	DeletedBy *string `json:"deleted_by,omitempty"`
}
