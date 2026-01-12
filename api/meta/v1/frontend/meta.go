package frontend

import (
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// 元数据，包含创建、更新、删除这个数据的用户的 ID
type MetadataWithOperator struct {
	meta_v1.Metadata
	// 创建人
	CreatedBy ReferenceWithName `json:"created_by,omitempty"`
	// 更新人
	UpdatedBy ReferenceWithName `json:"updated_by,omitempty"`
	// 删除人。当用户请求优雅删除时，服务器会设置这个字段，并不是由用户直接设
	// 置。
	DeletedBy *ReferenceWithName `json:"deleted_by,omitempty"`
}

type ReferenceWithName struct {
	// ID
	ID string `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
}
