package gozero

import (
	"context"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/access_control"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	af_trace "github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

// AccessControl 访问控制中间件，根据资源和 HTTP 方法验证用户权限
func (m *Middleware) AccessControl(resource access_control.Resource) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			var userID string
			newCtx, span := af_trace.StartInternalSpan(r.Context())
			defer func() { af_trace.TelemetrySpanEnd(span, err) }()

			// 获取用户 ID
			tokenType, ok := r.Context().Value(interception.TokenType).(int)
			if ok && tokenType == interception.TokenTypeClient {
				// Client Token 需要从 Header 获取 userId
				userID = r.Header.Get("userId")
				if userID == "" {
					respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
					return
				}
			} else {
				// User Token 从 Context 获取
				if user, ok := r.Context().Value(interception.InfoName).(*middleware.User); ok {
					userID = user.ID
				}
			}

			// 如果 TokenType 是 Client，允许访问
			if ok && tokenType == interception.TokenTypeClient {
				next(w, r)
				return
			}

			// 检查权限
			pass, err := m.distributeAccessType(newCtx, userID, r.Method, resource)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, err)
				return
			}
			if !pass {
				respondWithError(w, http.StatusForbidden, errorcode.Desc(errorcode.AuthorizationFailure))
				return
			}

			next(w, r)
		}
	}
}

// MultipleAccessControl 多资源访问控制，只要满足任一资源权限即可通过
func (m *Middleware) MultipleAccessControl(resources ...access_control.Resource) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			var userID string
			newCtx, span := af_trace.StartInternalSpan(r.Context())
			defer func() { af_trace.TelemetrySpanEnd(span, err) }()

			// 获取用户 ID
			tokenType, ok := r.Context().Value(interception.TokenType).(int)
			if ok && tokenType == interception.TokenTypeClient {
				userID = r.Header.Get("userId")
				if userID == "" {
					respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
					return
				}
			} else {
				if user, ok := r.Context().Value(interception.InfoName).(*middleware.User); ok {
					userID = user.ID
				}
			}

			// Client Token 允许访问
			if ok && tokenType == interception.TokenTypeClient {
				next(w, r)
				return
			}

			// 检查多个资源，任一通过即可
			for _, resource := range resources {
				pass, err := m.distributeAccessType(newCtx, userID, r.Method, resource)
				if err != nil {
					respondWithError(w, http.StatusBadRequest, err)
					return
				}
				if pass {
					next(w, r)
					return
				}
			}

			respondWithError(w, http.StatusForbidden, errorcode.Desc(errorcode.AuthorizationFailure))
		}
	}
}

// AccessControlWithAccessType 指定访问类型的访问控制
func (m *Middleware) AccessControlWithAccessType(accessType access_control.AccessType, resource access_control.Resource) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			var userID string
			newCtx, span := af_trace.StartInternalSpan(r.Context())
			defer func() { af_trace.TelemetrySpanEnd(span, err) }()

			// 获取用户 ID
			tokenType, ok := r.Context().Value(interception.TokenType).(int)
			if ok && tokenType == interception.TokenTypeClient {
				userID = r.Header.Get("userId")
				if userID == "" {
					respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
					return
				}
			} else {
				if user, ok := r.Context().Value(interception.InfoName).(*middleware.User); ok {
					userID = user.ID
				}
			}

			// Client Token 允许访问
			if ok && tokenType == interception.TokenTypeClient {
				next(w, r)
				return
			}

			// 检查权限
			pass, err := m.configurationCenterDriven.HasAccessPermission(newCtx, userID, accessType, resource)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, err)
				return
			}
			if !pass {
				respondWithError(w, http.StatusForbidden, errorcode.Desc(errorcode.AuthorizationFailure))
				return
			}

			next(w, r)
		}
	}
}

// AccessControlSkipTypeClient 跳过 Client 类型检查的访问控制
func (m *Middleware) AccessControlSkipTypeClient(resource access_control.Resource) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 如果 TokenType 是 Client，允许访问
			if tokenType, ok := r.Context().Value(interception.TokenType).(int); ok {
				if tokenType == interception.TokenTypeClient {
					next(w, r)
					return
				}
			}

			// 否则执行标准访问控制
			handler := m.AccessControl(resource)(next)
			handler(w, r)
		}
	}
}

// distributeAccessType 根据 HTTP 方法分发访问权限检查
func (m *Middleware) distributeAccessType(ctx context.Context, userID string, method string, resource access_control.Resource) (bool, error) {
	switch method {
	case http.MethodGet:
		return m.configurationCenterDriven.HasAccessPermission(ctx, userID, access_control.GET_ACCESS, resource)
	case http.MethodPost:
		return m.configurationCenterDriven.HasAccessPermission(ctx, userID, access_control.POST_ACCESS, resource)
	case http.MethodPut:
		return m.configurationCenterDriven.HasAccessPermission(ctx, userID, access_control.PUT_ACCESS, resource)
	case http.MethodDelete:
		return m.configurationCenterDriven.HasAccessPermission(ctx, userID, access_control.DELETE_ACCESS, resource)
	default:
		return false, errorcode.Desc(errorcode.AccessTypeNotSupport)
	}
}
