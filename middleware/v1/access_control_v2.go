package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	"github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/transport/rest/ginx"
)

func GetUserIDFromContext(c *gin.Context) string {
	if c.Value(interception.TokenTypeClient) != nil && c.Value(interception.TokenTypeClient).(int) == interception.TokenTypeClient {
		return c.Request.Header.Get("userId")
	}
	if user, exist := c.Value(interception.InfoName).(*middleware.User); exist {
		return user.ID
	}
	return ""
}

func (m *Middleware) MenuPermissionMarker() *middleware.MenuResourceMarkerGenerator {
	return &middleware.MenuResourceMarkerGenerator{
		HandlerGenerator: m.checkResourceAuth,
	}
}

func (m *Middleware) checkResourceAuth(action string, menuKeys ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestPath := c.FullPath()
		//内部接口不需要
		if strings.Contains(requestPath, "/internal/") {
			c.Next()
			return
		}
		//尝试从context里面获取下
		if len(menuKeys) <= 0 {
			//没有标注的，放过
			menuKeyStr, exists := c.Get(middleware.API_MENU_KEY)
			if !exists {
				c.Next()
				return
			}
			menuKeys = strings.Split(menuKeyStr.(string), ",")
			if len(menuKeys) <= 0 {
				c.Next()
				return
			}
		}
		//获取用户ID
		userID := GetUserIDFromContext(c)
		if userID == "" {
			ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
			return
		}
		//当前还不知道怎么处理结果
		for _, menuKey := range menuKeys {
			authReq := &authorization.OperationCheckArgs{
				Accessor: authorization.Accessor{
					ID:   userID,
					Type: authorization.ACCESSOR_TYPE_USER,
				},
				Resource: authorization.ResourceObject{
					ID:   menuKey,
					Type: authorization.RESOURCE_TYPE_MENUS,
				},
				Operation: []string{action},
				Include:   []string{authorization.INCLUDE_OPERATION_OBLIGATIONS},
				Method:    "GET",
			}
			result, err := m.authorization.OperationCheck(c, authReq)
			if err != nil {
				log.Errorf("CheckUserPermission Error %v", err.Error())
				ginx.AbortResponseWithCode(c, http.StatusBadRequest, errorcode.Desc(errorcode.AuthServiceException))
				return
			}
			//如果有一个允许，就可以访问
			if result.Result {
				//设置范围
				c.Set(interception.PermissionScope, result.OperationScope())
				c.Next()
				return
			}
		}
		//被拒绝了
		menuKeyCNNames := middleware.GetMenuComments(menuKeys)
		ginx.AbortResponseWithCode(c, http.StatusForbidden, errorcode.Desc(errorcode.PermissionCheckFailure, strings.Join(menuKeyCNNames, ",")))
	}
}
