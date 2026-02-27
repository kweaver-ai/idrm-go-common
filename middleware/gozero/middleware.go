package gozero

import (
	"github.com/kweaver-ai/idrm-go-common/audit"
	auth_service "github.com/kweaver-ai/idrm-go-common/rest/auth-service"
	"github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
	"github.com/kweaver-ai/idrm-go-common/rest/user_management"
)

// Middleware 中间件结构体，复用相同的依赖组件
type Middleware struct {
	hydra                     hydra.Hydra
	userMgm                   user_management.DrivenUserMgnt
	configurationCenterDriven configuration_center.Driven
	// 审计日志的日志器
	auditLogger               audit.Logger
	authorization             authorization.Driven
	authServiceInternalClient auth_service.AuthServiceInternalV1Interface
}

// NewMiddleware 创建新的 Middleware 实例
func NewMiddleware(
	hydra hydra.Hydra,
	userMgm user_management.DrivenUserMgnt,
	configurationCenterDriven configuration_center.Driven,
	// 审计日志的日志器
	auditLogger audit.Logger,
	authorization authorization.Driven,
	authServiceInternalClient auth_service.AuthServiceInternalV1Interface,
) *Middleware {
	return &Middleware{
		hydra:                     hydra,
		userMgm:                   userMgm,
		configurationCenterDriven: configurationCenterDriven,
		auditLogger:               auditLogger,
		authorization:             authorization,
		authServiceInternalClient: authServiceInternalClient,
	}
}
