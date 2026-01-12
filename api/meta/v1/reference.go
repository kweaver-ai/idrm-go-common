package v1

// 资源引用，包含资源的名称
type ReferenceWithName struct {
	// ID
	ID string `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
}
