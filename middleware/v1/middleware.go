package v1

import (
	"github.com/kweaver-ai/idrm-go-common/audit"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	auth_service "github.com/kweaver-ai/idrm-go-common/rest/auth-service"
	"github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
	"github.com/kweaver-ai/idrm-go-common/rest/user_management"
)

type Middleware struct {
	hydra                     hydra.Hydra
	userMgm                   user_management.DrivenUserMgnt
	configurationCenterDriven configuration_center.Driven
	// 审计日志的日志器
	auditLogger               audit.Logger
	authorization             authorization.Driven
	authServiceInternalClient auth_service.AuthServiceInternalV1Interface
}

func NewMiddleware(
	hydra hydra.Hydra,
	userMgm user_management.DrivenUserMgnt,
	configurationCenterDriven configuration_center.Driven,
	// 审计日志的日志器
	auditLogger audit.Logger,
	authorization authorization.Driven,
	authServiceInternalClient auth_service.AuthServiceInternalV1Interface,
) middleware.Middleware {
	return &Middleware{
		hydra:                     hydra,
		userMgm:                   userMgm,
		configurationCenterDriven: configurationCenterDriven,
		auditLogger:               auditLogger,
		authorization:             authorization,
		authServiceInternalClient: authServiceInternalClient,
	}
}
