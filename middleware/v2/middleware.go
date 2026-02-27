package v2

import (
	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
	"github.com/kweaver-ai/idrm-go-common/rest/user_management"
)

// Middleware Token 鉴权中间件
type Middleware struct {
	hydra   hydra.Hydra
	userMgm user_management.DrivenUserMgnt
}

// NewMiddleware 创建中间件实例
func NewMiddleware(hydra hydra.Hydra, userMgm user_management.DrivenUserMgnt) *Middleware {
	return &Middleware{
		hydra:   hydra,
		userMgm: userMgm,
	}
}
