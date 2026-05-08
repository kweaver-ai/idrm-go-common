package bkn_backend

import "context"

// Driven 知识网络（ontology-manager 内部接口）HTTP 客户端
type Driven interface {
	// GetDetail 获取知识网络详情（export 模式，含 object_types、relation_types 等）
	GetDetail(ctx context.Context, knID string) (*KnowledgeNetworkDetail, error)
}

// KnowledgeNetworkDetail 知识网络详情（仅声明映射构建所需字段；其余 JSON 字段反序列化时会被忽略）
type KnowledgeNetworkDetail struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ModuleType     string `json:"module_type"`
	BusinessDomain string `json:"business_domain"`
}
