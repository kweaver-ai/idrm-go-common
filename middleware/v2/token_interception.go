package v2

import (
	"context"
	"net/http"
	"strings"

	common_middleware "github.com/kweaver-ai/idrm-go-common/middleware"
	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	// "github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"github.com/kweaver-ai/idrm-go-frame/core/transport/rest/gozerox"
	"go.uber.org/zap"
)

// TokenInterception Token 拦截中间件
func (m *Middleware) TokenInterception() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			// newCtx, span := trace.StartInternalSpan(r.Context())
			// defer func() { trace.TelemetrySpanEnd(span, err) }()
			newCtx := r.Context()

			tokenID := r.Header.Get("Authorization")
			token := strings.TrimPrefix(tokenID, "Bearer ")
			if tokenID == "" || token == "" {
				gozerox.AbortResponseWithCode(w, http.StatusUnauthorized, errorcode.Desc(errorcode.NotAuthentication))
				return
			}

			info, err := m.hydra.IntrospectV2(newCtx, token)
			if err != nil {
				log.WithContext(r.Context()).Error("Token Introspect", zap.Error(err))
				gozerox.AbortResponseWithCode(w, http.StatusBadRequest, errorcode.Desc(errorcode.HydraException))
				return
			}
			if !info.Active {
				gozerox.AbortResponseWithCode(w, http.StatusUnauthorized, errorcode.Desc(errorcode.AuthenticationFailure))
				return
			}

			// 保存 Bearer token 用于身份认证
			newCtx = interception.NewContextWithBearerToken(newCtx, token)

			// 区分应用 Token 和用户 Token
			if info.VisitorID == info.ClientID || info.VisitorTyp == hydra.App {
				// 应用 Token 处理逻辑
				newCtx, err = m.handleAppToken(newCtx, r, &info, tokenID)
				if err != nil {
					log.WithContext(r.Context()).Error("handleAppToken", zap.Error(err))
					gozerox.AbortResponseWithCode(w, http.StatusBadRequest, err)
					return
				}
			} else {
				// 用户 Token 处理逻辑
				newCtx, err = m.handleUserToken(newCtx, &info, tokenID)
				if err != nil {
					log.WithContext(r.Context()).Error("handleUserToken", zap.Error(err))
					gozerox.AbortResponseWithCode(w, http.StatusBadRequest, err)
					return
				}
			}

			next(w, r.WithContext(newCtx))
		}
	}
}

// handleAppToken 处理应用 Token
func (m *Middleware) handleAppToken(ctx context.Context, r *http.Request, info *hydra.TokenIntrospectInfo, tokenID string) (context.Context, error) {
	// 获取应用信息
	userInfoApp, err := m.userMgm.GetAppInfo(r.Context(), info.VisitorID)
	if err != nil {
		return ctx, errorcode.Desc(errorcode.GetProtonAppInfoFailure)
	}

	userInfo := &common_middleware.User{ID: userInfoApp.ID, Name: userInfoApp.Name, UserType: interception.TokenTypeClient}
	subject := &v1.Subject{Type: v1.SubjectAPP, ID: userInfo.ID}

	// 保存到 Context
	ctx = interception.NewContextWithAuthServiceSubject(ctx, subject)
	ctx = context.WithValue(ctx, interception.InfoName, userInfo)
	ctx = context.WithValue(ctx, interception.Token, tokenID)
	ctx = context.WithValue(ctx, interception.TokenType, interception.TokenTypeClient)

	return ctx, nil
}

// handleUserToken 处理用户 Token
func (m *Middleware) handleUserToken(ctx context.Context, info *hydra.TokenIntrospectInfo, tokenID string) (context.Context, error) {
	name, _, depInfos, err := m.userMgm.GetUserNameByUserID(ctx, info.VisitorID)
	if err != nil {
		return ctx, errorcode.Desc(errorcode.GetUserInfoFailure)
	}

	userInfo := &common_middleware.User{ID: info.VisitorID, Name: name, OrgInfos: depInfos, UserType: interception.TokenTypeUser}
	subject := &v1.Subject{Type: v1.SubjectUser, ID: userInfo.ID}

	// 保存到 Context
	ctx = interception.NewContextWithAuthServiceSubject(ctx, subject)
	ctx = context.WithValue(ctx, interception.InfoName, userInfo)
	ctx = context.WithValue(ctx, interception.Token, tokenID)
	ctx = context.WithValue(ctx, interception.TokenType, interception.TokenTypeUser)

	return ctx, nil
}
