package v1

import "net"

// Agent 定义用户代理
type Agent struct {
	// 客户端类型
	Type AgentType `json:"type,omitempty"`
	// 客户端 IP 地址
	IP net.IP `json:"ip,omitempty"`
}

// AgentType 定义客户端类型
type AgentType string

const (
	// 浏览器
	AgentWeb AgentType = "web"
	// 未知类型
	AgentUnknown AgentType = "unknown"
)
