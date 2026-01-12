package v1

import (
	"context"
	"net/http"
	"strings"

	v1 "github.com/kweaver-ai/idrm-go-common/api/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
	af_trace "github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"github.com/kweaver-ai/idrm-go-frame/core/transport/rest/ginx"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func (m *Middleware) TokenInterception() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		newCtx, span := af_trace.StartInternalSpan(c.Request.Context())
		defer func() { af_trace.TelemetrySpanEnd(span, err) }()
		tokenID := c.GetHeader("Authorization")
		token := strings.TrimPrefix(tokenID, "Bearer ")
		if tokenID == "" || token == "" {
			ginx.AbortResponseWithCode(c, http.StatusUnauthorized, errorcode.Desc(errorcode.NotAuthentication))
			return
		}
		info, err := m.hydra.Introspect(newCtx, token)
		if err != nil {
			log.WithContext(c.Request.Context()).Error("TokenInterception Introspect", zap.Error(err))
			ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.HydraException))
			return
		}
		if !info.Active {
			ginx.AbortResponseWithCode(c, http.StatusUnauthorized, errorcode.Desc(errorcode.AuthenticationFailure))
			return
		}

		// 保存 Bearer token 用于身份认证
		interception.SetGinContextWithBearerToken(c, token)

		if info.VisitorID == info.ClientID || info.VisitorTyp == hydra.App {

			// 根据名称判断是不是虚拟化引擎的内部账号，如果是，后面不检查权限
			client_name, err := m.hydra.GetClientNameById(newCtx, info.VisitorID)
			if err != nil {
				log.WithContext(c.Request.Context()).Error("TokenInterception GetUserFromHydarWrong", zap.Error(err))
				ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.HydraException))
				return
			}
			if client_name == middleware.VirtualEngineApp {
				c.Set(interception.Token, tokenID)
				c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.Token, tokenID))
				c.Set(interception.TokenType, interception.TokenTypeClient)
				c.Next()
				return
			}

			var id, name string
			// 查看是否是AF应用中创建的账号
			//appinfo, err := m.configurationCenterDriven.GetAppsByAccountId(newCtx, info.VisitorID)
			//if err != nil {
			//	log.WithContext(c.Request.Context()).Error("configuration GetAppsByAccountId err", zap.Error(err))
			//	ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.GetAppInfoFailure))
			//	return
			//}
			//id = appinfo.ID
			//name = appinfo.Name

			// 如果应用没有查到就是proton部署控制台建的账号
			//if appinfo.ID == "" {
			userInfoApp, err := m.userMgm.GetAppInfo(c, info.VisitorID)
			if err != nil {
				log.WithContext(c.Request.Context()).Error("TokenInterception userMgm GetAppInfo", zap.Error(err))
				ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.GetProtonAppInfoFailure))
				return
			}
			id = userInfoApp.ID
			name = userInfoApp.Name
			//}

			userInfo := &middleware.User{ID: id, Name: name, UserType: interception.TokenTypeClient}
			// 保存访问者到 Context 用于鉴权
			subject := &v1.Subject{Type: v1.SubjectAPP, ID: userInfo.ID}
			interception.SetGinContextWithAuthServiceSubject(c, subject)
			c.Request = c.Request.WithContext(interception.NewContextWithAuthServiceSubject(c.Request.Context(), subject))
			c.Set(interception.InfoName, userInfo)
			c.Set(interception.Token, tokenID)
			c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.InfoName, userInfo))
			c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.Token, tokenID))
			c.Set(interception.TokenType, interception.TokenTypeClient)
			c.Next()
			return
		}
		name, _, depInfos, err := m.userMgm.GetUserNameByUserID(newCtx, info.VisitorID)
		if err != nil {
			log.WithContext(c.Request.Context()).Error("userMgm GetUserNameByUserID err", zap.Error(err))
			ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.GetUserInfoFailure))
			return
		}
		userInfo := &middleware.User{ID: info.VisitorID, Name: name, OrgInfos: depInfos, UserType: interception.TokenTypeUser}
		subject := &v1.Subject{Type: v1.SubjectUser, ID: userInfo.ID}
		interception.SetGinContextWithAuthServiceSubject(c, subject)
		c.Request = c.Request.WithContext(interception.NewContextWithAuthServiceSubject(c.Request.Context(), subject))
		c.Set(interception.InfoName, userInfo)
		c.Set(interception.Token, tokenID)
		c.Set(interception.TokenType, interception.TokenTypeUser)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.InfoName, userInfo))
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.Token, tokenID))
		interception.SetGinContextWithUser(c, &interception.User{ID: userInfo.ID, Name: userInfo.Name})
		c.Next()
	}
}

// ShouldTokenInterception implements middleware.Middleware.
func (m *Middleware) ShouldTokenInterception() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := af_trace.StartInternalSpan(c)
		defer span.End()

		auth := c.GetHeader("Authorization")
		if auth == "" {
			return
		}
		interception.SetGinContextWithAuth(c, auth)

		if !strings.HasPrefix(auth, "Bearer ") {
			return
		}
		bearerToken := strings.TrimPrefix(auth, "Bearer ")
		interception.SetGinContextWithBearerToken(c, bearerToken)

		info, err := m.hydra.Introspect(ctx, bearerToken)
		if err != nil || !info.Active {
			span.RecordError(err)
			return
		}

		switch {
		// 访问者是一个应用
		case info.VisitorID == info.ClientID || info.VisitorTyp == hydra.App:
			// 根据名称判断是不是虚拟化引擎的内部账号，如果是，后面不检查权限
			name, err := m.hydra.GetClientNameById(ctx, info.VisitorID)
			if err != nil {
				span.RecordError(err)
				return
			}
			if name == middleware.VirtualEngineApp {
				// TODO: 添加方法，根据 Context 判断请求是否来自虚拟化引擎
				return
			}
			// 获取 hydra 账号对应的 AnyFabric 应用
			//app, err := m.configurationCenterDriven.GetAppsByAccountId(ctx, info.VisitorID)
			//// 未找到对应的 AnyFabric 时，返回的 err 为 nil，app.ID 为 空
			//if err != nil || app.ID == "" {
			//	span.RecordError(err)
			//	return
			//}
			userInfoApp, err := m.userMgm.GetAppInfo(c, info.VisitorID)
			if err != nil {
				log.WithContext(c.Request.Context()).Error("TokenInterception userMgm ShouldTokenInterception", zap.Error(err))
				ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.GetProtonAppInfoFailure))
				return
			}
			// 保存访问者到 Context 用于鉴权
			interception.SetGinContextWithAuthServiceSubject(c, &v1.Subject{Type: v1.SubjectAPP, ID: userInfoApp.ID})

		// 访问者是一个用户
		default:
			// 保存访问者到 Context 用于鉴权
			interception.SetGinContextWithAuthServiceSubject(c, &v1.Subject{Type: v1.SubjectUser, ID: info.VisitorID})
		}
	}
}

// TokenPassThrough 内部接口自动注册Token透传中间件
func (m *Middleware) TokenPassThrough() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		newCtx, span := af_trace.StartInternalSpan(c.Request.Context())
		defer func() { af_trace.TelemetrySpanEnd(span, err) }()

		tokenID := c.GetHeader("Authorization")
		token := strings.TrimPrefix(tokenID, "Bearer ")
		if token != "" {
			info, err := m.hydra.Introspect(newCtx, token)
			if err == nil {
				name, _, depInfos, err := m.userMgm.GetUserNameByUserID(newCtx, info.VisitorID)
				if err == nil {
					userInfo := &middleware.User{ID: info.VisitorID, Name: name, OrgInfos: depInfos, UserType: interception.TokenTypeUser}
					c.Set(interception.InfoName, userInfo)
					c.Set(interception.Token, tokenID)
					c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.InfoName, userInfo))
					c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), interception.Token, tokenID))
				}
			}
		}
		c.Next()
	}
}
