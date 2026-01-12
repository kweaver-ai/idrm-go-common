package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	"github.com/kweaver-ai/idrm-go-common/util"
	"github.com/kweaver-ai/idrm-go-frame/core/transport/rest/ginx"
)

func (m *Middleware) PermissionControl(roleIDs []string, exclude bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		util.HandleReqWithTraceIncludingErrLog(c.Request.Context(), func(ctx context.Context) (_ any, err error) {
			// [来自APP Client的Token允许访问]
			if c.Value(interception.TokenType) == interception.TokenTypeClient {
				return
			} // [/]
			// [解析用户ID]
			var userID string
			if c.Value(interception.TokenTypeClient) == interception.TokenTypeClient {
				userID = c.Request.Header.Get("userId")
			} else if user, ok := c.Value(interception.InfoName).(*middleware.User); ok {
				userID = user.ID
			} // [/]
			// [用户ID为空不允许访问]
			if userID == "" {
				ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
				return
			} // [/]

			// TODO 20250603去掉信息资源编目接口（内部+外部），单独的一套验证角色权限
			//[检查用户是否拥有符合权限要求的角色身份]
			//userRoleIDs, err := m.configurationCenterDriven.GetRoleIDs(ctx, userID)
			//if err != nil {
			//	ginx.AbortResponseWithCode(c, http.StatusBadRequest, err)
			//	return
			//}
			//for _, id := range userRoleIDs {
			//	if util.Contains(roleIDs, id) == !exclude {
			//		return
			//	}
			//}
			//ginx.AbortResponseWithCode(c, http.StatusForbidden, errorcode.Desc(errorcode.AuthorizationFailure)) // [/]
			return
		})
	}
}
