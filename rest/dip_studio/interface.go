package dip_studio

import "context"

// Driven DIP Studio（/api/dip-studio/v1）HTTP 客户端
type Driven interface {
	// ListDigitalHumans 获取数字员工列表。
	ListDigitalHumans(ctx context.Context) ([]*DigitalHumanDetail, error)
	// GetDigitalHumanDetail 获取单个数字员工详情。
	GetDigitalHumanDetail(ctx context.Context, digitalHumanID string) (*DigitalHumanDetail, error)
	// AddDigitalHumanKnowledgeNetwork 在已有数字员工上追加业务知识网络（BKN）条目。
	// 会先 GET 详情再 PUT 整组 bkn；与同名同 URL 的已有条目去重（以 URL 为主键，URL 为空时退化为 name）。
	AddDigitalHumanKnowledgeNetwork(ctx context.Context, digitalHumanID string, entries []BknEntry) (*DigitalHumanDetail, error)
}

// BknEntry 业务知识网络 / 知识源条目（与 DIP Studio OpenAPI 中 BknEntry 一致）
type BknEntry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// DigitalHumanDetail 数字员工详情（仅声明常用字段；其余 JSON 会被忽略）
type DigitalHumanDetail struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Creature string     `json:"creature,omitempty"`
	IconID   string     `json:"icon_id,omitempty"`
	Soul     string     `json:"soul,omitempty"`
	Skills   []string   `json:"skills,omitempty"`
	Bkn      []BknEntry `json:"bkn,omitempty"`
}
