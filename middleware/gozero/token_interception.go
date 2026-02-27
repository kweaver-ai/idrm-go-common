package gozero

import (
	"context"
	"net/http"
	"strings"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	af_trace "github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"go.uber.org/zap"
)

// TokenInterception Token 拦截中间件，验证 Bearer Token 并解析用户信息
// 验证失败会中断请求并返回错误
func (m *Middleware) TokenInterception() MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			newCtx, span := af_trace.StartInternalSpan(r.Context())
			defer func() { af_trace.TelemetrySpanEnd(span, err) }()

			tokenID := r.Header.Get("Authorization")
			token := strings.TrimPrefix(tokenID, "Bearer ")
			if tokenID == "" || token == "" {
				respondWithError(w, http.StatusUnauthorized, errorcode.Desc(errorcode.NotAuthentication))
				return
			}

			info, err := m.hydra.Introspect(newCtx, token)
			if err != nil {
				log.WithContext(r.Context()).Error("TokenInterception Introspect", zap.Error(err))
				respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.HydraException))
				return
			}
			if !info.Active {
				respondWithError(w, http.StatusUnauthorized, errorcode.Desc(errorcode.AuthenticationFailure))
				return
			}

			// 保存 Bearer token 用于身份认证
			newCtx = interception.NewContextWithBearerToken(newCtx, token)

			// 判断是应用还是用户
			if info.VisitorID == info.ClientID || info.VisitorTyp == hydra.App {
				// 根据名称判断是不是虚拟化引擎的内部账号
				clientName, err := m.hydra.GetClientNameById(newCtx, info.VisitorID)
				if err != nil {
					log.WithContext(r.Context()).Error("TokenInterception GetClientNameById", zap.Error(err))
					respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.HydraException))
					return
				}
				if clientName == middleware.VirtualEngineApp {
					// 虚拟化引擎内部账号，不检查权限
					newCtx = context.WithValue(newCtx, interception.Token, tokenID)
					newCtx = context.WithValue(newCtx, interception.TokenType, interception.TokenTypeClient)
					next(w, r.WithContext(newCtx))
					return
				}

				var id, name string
				// 查看是否是 AF 应用中创建的账号
				userInfoApp, err := m.userMgm.GetAppInfo(newCtx, info.VisitorID)
				if err != nil {
					log.WithContext(r.Context()).Error("TokenInterception userMgm GetAppInfo", zap.Error(err))
					respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.GetProtonAppInfoFailure))
					return
				}
				id = userInfoApp.ID
				name = userInfoApp.Name

				userInfo := &middleware.User{ID: id, Name: name, UserType: interception.TokenTypeClient}
				// 保存访问者到 Context 用于鉴权
				subject := &v1.Subject{Type: v1.SubjectAPP, ID: userInfo.ID}
				newCtx = interception.NewContextWithAuthServiceSubject(newCtx, subject)
				newCtx = context.WithValue(newCtx, interception.InfoName, userInfo)
				newCtx = context.WithValue(newCtx, interception.Token, tokenID)
				newCtx = context.WithValue(newCtx, interception.TokenType, interception.TokenTypeClient)
				next(w, r.WithContext(newCtx))
				return
			}

			// 用户类型
			name, _, depInfos, err := m.userMgm.GetUserNameByUserID(newCtx, info.VisitorID)
			if err != nil {
				log.WithContext(r.Context()).Error("userMgm GetUserNameByUserID err", zap.Error(err))
				respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.GetUserInfoFailure))
				return
			}
			userInfo := &middleware.User{ID: info.VisitorID, Name: name, OrgInfos: depInfos, UserType: interception.TokenTypeUser}
			subject := &v1.Subject{Type: v1.SubjectUser, ID: userInfo.ID}
			newCtx = interception.NewContextWithAuthServiceSubject(newCtx, subject)
			newCtx = context.WithValue(newCtx, interception.InfoName, userInfo)
			newCtx = context.WithValue(newCtx, interception.Token, tokenID)
			newCtx = context.WithValue(newCtx, interception.TokenType, interception.TokenTypeUser)

			// 设置 interception.User（简化版）
			newCtx = interception.NewContextWithUser(newCtx, &interception.User{ID: userInfo.ID, Name: userInfo.Name})

			next(w, r.WithContext(newCtx))
		}
	}
}

// ShouldTokenInterception 尝试 token interception，即使失败不终止
// 经过此中间件的请求支持：
//   - 通过 interception.AuthFromContext 获取 HTTP Header Authorization
//   - 通过 interception.BearerTokenFromContext 获取 Bearer Token
//   - 通过 interception.AuthServiceSubjectFromContext 获取 auth-service Subject
func (m *Middleware) ShouldTokenInterception() MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx, span := af_trace.StartInternalSpan(r.Context())
			defer span.End()

			auth := r.Header.Get("Authorization")
			if auth == "" {
				next(w, r)
				return
			}
			ctx = interception.NewContextWithAuth(ctx, auth)

			if !strings.HasPrefix(auth, "Bearer ") {
				next(w, r.WithContext(ctx))
				return
			}
			bearerToken := strings.TrimPrefix(auth, "Bearer ")
			ctx = interception.NewContextWithBearerToken(ctx, bearerToken)

			info, err := m.hydra.Introspect(ctx, bearerToken)
			if err != nil || !info.Active {
				span.RecordError(err)
				next(w, r.WithContext(ctx))
				return
			}

			switch {
			// 访问者是一个应用
			case info.VisitorID == info.ClientID || info.VisitorTyp == hydra.App:
				// 根据名称判断是不是虚拟化引擎的内部账号
				name, err := m.hydra.GetClientNameById(ctx, info.VisitorID)
				if err != nil {
					span.RecordError(err)
					next(w, r.WithContext(ctx))
					return
				}
				if name == middleware.VirtualEngineApp {
					next(w, r.WithContext(ctx))
					return
				}

				userInfoApp, err := m.userMgm.GetAppInfo(ctx, info.VisitorID)
				if err != nil {
					log.WithContext(r.Context()).Error("ShouldTokenInterception userMgm GetAppInfo", zap.Error(err))
					respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.GetProtonAppInfoFailure))
					return
				}
				// 保存访问者到 Context 用于鉴权
				ctx = interception.NewContextWithAuthServiceSubject(ctx, &v1.Subject{Type: v1.SubjectAPP, ID: userInfoApp.ID})

			// 访问者是一个用户
			default:
				// 保存访问者到 Context 用于鉴权
				ctx = interception.NewContextWithAuthServiceSubject(ctx, &v1.Subject{Type: v1.SubjectUser, ID: info.VisitorID})
			}

			next(w, r.WithContext(ctx))
		}
	}
}

// TokenPassThrough 内部接口自动注册 Token 透传中间件
// 不强制验证 Token，如果 Token 有效则解析用户信息
func (m *Middleware) TokenPassThrough() MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			newCtx, span := af_trace.StartInternalSpan(r.Context())
			defer func() { af_trace.TelemetrySpanEnd(span, err) }()

			tokenID := r.Header.Get("Authorization")
			token := strings.TrimPrefix(tokenID, "Bearer ")
			if token != "" {
				info, err := m.hydra.Introspect(newCtx, token)
				if err == nil {
					name, _, depInfos, err := m.userMgm.GetUserNameByUserID(newCtx, info.VisitorID)
					if err == nil {
						userInfo := &middleware.User{ID: info.VisitorID, Name: name, OrgInfos: depInfos, UserType: interception.TokenTypeUser}
						newCtx = context.WithValue(newCtx, interception.InfoName, userInfo)
						newCtx = context.WithValue(newCtx, interception.Token, tokenID)
					}
				}
			}
			next(w, r.WithContext(newCtx))
		}
	}
}
