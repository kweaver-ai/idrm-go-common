package authorization

import "context"

// Obligation 义务接口
type Obligation interface {
	CreateObligationType(ctx context.Context, typeName string, req *ObligationTypeReq) error
}

// ObligationTypeReq 义务类型详情
type ObligationTypeReq struct {
	Name                    string                  `json:"name"`
	Schema                  Schema                  `json:"schema"`
	ApplicableResourceTypes ApplicableResourceTypes `json:"applicable_resource_types"`
	Description             string                  `json:"description"`
	DefaultValue            DefaultValue            `json:"default_value"`
}

// ScopeSchema 作用域 schema 定义
type ScopeSchema struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Enum        []string `json:"enum"`
	EnumNames   []string `json:"enumNames"`
	Default     string   `json:"default"`
}

// SchemaProperties schema 属性定义
type SchemaProperties struct {
	Scope ScopeSchema `json:"scope"`
}

// Schema 模式定义
type Schema struct {
	Type       string           `json:"type"`
	Properties SchemaProperties `json:"properties"`
	Required   []string         `json:"required"`
}

// ApplicableOperations 适用操作定义
type ApplicableOperations struct {
	Unlimited  bool  `json:"unlimited"`
	Operations []any `json:"operations"`
}

// ObligationResourceType 义务资源类型定义
type ObligationResourceType struct {
	ID                   string               `json:"id"`
	ApplicableOperations ApplicableOperations `json:"applicable_operations"`
}

// ApplicableResourceTypes 适用资源类型定义
type ApplicableResourceTypes struct {
	Unlimited     bool                     `json:"unlimited"`
	ResourceTypes []ObligationResourceType `json:"resource_types"`
}

// DefaultValue 默认值定义
type DefaultValue struct {
	Scope string `json:"scope"`
}
