package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kweaver-ai/idrm-go-common/access_control"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	af_trace "github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"github.com/kweaver-ai/idrm-go-frame/core/transport/rest/ginx"
)

func (m *Middleware) AccessControl(resource access_control.Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var userId string
		newCtx, span := af_trace.StartInternalSpan(c.Request.Context())
		defer func() { af_trace.TelemetrySpanEnd(span, err) }()
		if c.Value(interception.TokenTypeClient) != nil && c.Value(interception.TokenTypeClient).(int) == interception.TokenTypeClient {
			if c.Request.Header.Get("userId") == "" {
				ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
				return
			}
			userId = c.Request.Header.Get("userId")
		} else {
			if user, exist := c.Value(interception.InfoName).(*middleware.User); exist {
				userId = user.ID
			}
		}
		// 如果 TokenType 是 Client，即请求源自 APP 允许访问。
		v, exists := c.Get(interception.TokenType)
		if exists {
			vType, ok := v.(int)
			if ok && vType == interception.TokenTypeClient {
				c.Next()
				return
			}
		}

		// check := m.checkAppsAccessPermission(c, newCtx, userId)
		// if check {
		// 	c.Next()
		// 	return
		// }

		pass, err := m.distributeAccessType(newCtx, userId, c.Request.Method, resource)
		if err != nil {
			ginx.AbortResponseWithCode(c, http.StatusBadRequest, err)
			return
		}
		if !pass {
			ginx.AbortResponseWithCode(c, http.StatusForbidden, errorcode.Desc(errorcode.AuthorizationFailure))
			return
		}

		c.Next()
	}
}

func (m *Middleware) distributeAccessType(ctx context.Context, userid string, method string, resource access_control.Resource) (bool, error) {
	switch method {
	case http.MethodGet:
		return m.configurationCenterDriven.HasAccessPermission(ctx, userid, access_control.GET_ACCESS, resource)
	case http.MethodPost:
		return m.configurationCenterDriven.HasAccessPermission(ctx, userid, access_control.POST_ACCESS, resource)
	case http.MethodPut:
		return m.configurationCenterDriven.HasAccessPermission(ctx, userid, access_control.PUT_ACCESS, resource)
	case http.MethodDelete:
		return m.configurationCenterDriven.HasAccessPermission(ctx, userid, access_control.DELETE_ACCESS, resource)
	default:
		return false, errorcode.Desc(errorcode.AccessTypeNotSupport)
	}
}
func (m *Middleware) MultipleAccessControl(resources ...access_control.Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var userId string
		newCtx, span := af_trace.StartInternalSpan(c.Request.Context())
		defer func() { af_trace.TelemetrySpanEnd(span, err) }()
		if c.Value(interception.TokenTypeClient) != nil && c.Value(interception.TokenTypeClient).(int) == interception.TokenTypeClient {
			if c.Request.Header.Get("userId") == "" {
				ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
				return
			}
			userId = c.Request.Header.Get("userId")
		} else {
			if user, exist := c.Value(interception.InfoName).(*middleware.User); exist {
				userId = user.ID
			}
		}

		// 如果 TokenType 是 Client，即请求源自 APP 允许访问。
		v, exists := c.Get(interception.TokenType)
		if exists {
			vType, ok := v.(int)
			if ok && vType == interception.TokenTypeClient {
				c.Next()
				return
			}
		}

		// check := m.checkAppsAccessPermission(c, newCtx, userId)
		// if check {
		// 	c.Next()
		// 	return
		// }

		for _, resource := range resources {
			pass, err := m.distributeAccessType(newCtx, userId, c.Request.Method, resource)
			if err != nil {
				c.Writer.WriteHeader(http.StatusBadRequest)
				ginx.AbortResponse(c, err)
				return
			}
			if pass {
				c.Next()
				return
			}
		}
		c.Writer.WriteHeader(http.StatusForbidden)
		ginx.AbortResponse(c, errorcode.Desc(errorcode.AuthorizationFailure))
	}
}

func (m *Middleware) AccessControlWithAccessType(accessType access_control.AccessType, resource access_control.Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var userId string
		newCtx, span := af_trace.StartInternalSpan(c.Request.Context())
		defer func() { af_trace.TelemetrySpanEnd(span, err) }()
		if c.Value(interception.TokenTypeClient) != nil && c.Value(interception.TokenTypeClient).(int) == interception.TokenTypeClient {
			if c.Request.Header.Get("userId") == "" {
				ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
				return
			}
			userId = c.Request.Header.Get("userId")
		} else {
			if user, exist := c.Value(interception.InfoName).(*middleware.User); exist {
				userId = user.ID
			}
		}

		// 如果 TokenType 是 Client，即请求源自 APP 允许访问。
		v, exists := c.Get(interception.TokenType)
		if exists {
			vType, ok := v.(int)
			if ok && vType == interception.TokenTypeClient {
				c.Next()
				return
			}
		}

		// check := m.checkAppsAccessPermission(c, newCtx, userId)
		// if check {
		// 	c.Next()
		// 	return
		// }

		pass, err := m.configurationCenterDriven.HasAccessPermission(newCtx, userId, accessType, resource)
		if err != nil {
			c.Writer.WriteHeader(http.StatusBadRequest)
			ginx.AbortResponse(c, err)
			return
		}
		if !pass {
			c.Writer.WriteHeader(http.StatusForbidden)
			ginx.AbortResponse(c, errorcode.Desc(errorcode.AuthorizationFailure))
			return
		}

		c.Next()
	}
}

func (m *Middleware) AccessControlSkipTypeClient(resource access_control.Resource) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果 TokenType 是 Client，即请求源自 APP 允许访问。
		v, exists := c.Get(interception.TokenType)
		if exists {
			vType, ok := v.(int)
			if ok && vType == interception.TokenTypeClient {
				c.Next()
				return
			}
		}

		m.AccessControl(resource)(c)
		c.Next()
	}

}

// func (m *Middleware) checkAppsAccessPermission(c *gin.Context, newCtx context.Context, userId string) bool {
// 	// 如果 TokenType 是 Client，即请求源自 APP 允许访问。
// 	var err error
// 	v, exists := c.Get(interception.TokenType)
// 	if exists {
// 		vType, ok := v.(int)
// 		if ok && vType == interception.TokenTypeClient {
// 			var resource string
// 			for _, url := range middleware.PathDisableList {
// 				if strings.HasPrefix(c.Request.URL.Path, url) {
// 					ginx.AbortResponseWithCode(c, http.StatusBadRequest, err)
// 					return true
// 				}
// 			}
// 			for _, url := range middleware.PathEnableList {
// 				if strings.HasPrefix(c.Request.URL.Path, url) {
// 					resource = middleware.PathResourceMap[url]
// 				}
// 			}
// 			pass, err := m.distributeAppsAccessType(newCtx, userId, resource)
// 			if err != nil {
// 				ginx.AbortResponseWithCode(c, http.StatusBadRequest, err)
// 				return true
// 			}
// 			if !pass {
// 				ginx.AbortResponseWithCode(c, http.StatusForbidden, errorcode.Desc(errorcode.AuthorizationFailure))
// 				return true
// 			}
// 			return true
// 		}
// 	}
// 	return false
// }

// func (m *Middleware) distributeAppsAccessType(ctx context.Context, userid string, resource string) (bool, error) {
// 	return m.configurationCenterDriven.HasAppsAccessPermission(ctx, userid, resource)
// }
