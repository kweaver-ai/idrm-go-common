package v1

// 批处理：发布、上线接口服务
type BatchPublishAndOnline struct {
	// 需要被发布、上线的接口服务 ID 列表
	IDs []string `json:"ids,omitempty"`
}
