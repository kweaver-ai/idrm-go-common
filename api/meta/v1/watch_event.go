package v1

// WatchEvent 代表观察资源时接收到的创建、更新、删除事件。
type WatchEvent[T any] struct {
	Type WatchEventType `json:"type,omitempty"`

	// Type
	//  - Added, Modified：创建、更新后的资源
	//  - Deleted: 被删除前的资源
	Resource T `json:"resource,omitempty"`
}

// WatchEventType 代表观察事件类型
type WatchEventType string

const (
	Added    WatchEventType = "ADDED"
	Modified WatchEventType = "MODIFIED"
	Deleted  WatchEventType = "DELETED"
)
