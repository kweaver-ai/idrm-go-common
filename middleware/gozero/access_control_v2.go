package gozero

import (
	"context"
	"net/http"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	"github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

// GetUserIDFromContext 从请求上下文中获取用户 ID
func GetUserIDFromContext(r *http.Request) string {
	tokenType, ok := r.Context().Value(interception.TokenType).(int)
	if ok && tokenType == interception.TokenTypeClient {
		return r.Header.Get("userId")
	}
	if user, ok := r.Context().Value(interception.InfoName).(*middleware.User); ok {
		return user.ID
	}
	return ""
}

// MenuPermissionMarker 返回菜单权限标记生成器
func (m *Middleware) MenuPermissionMarker() *MenuResourceMarkerGenerator {
	return &MenuResourceMarkerGenerator{
		HandlerGenerator: m.checkResourceAuth,
	}
}

// MenuResourceMarkerGenerator 菜单资源标记生成器
type MenuResourceMarkerGenerator struct {
	HandlerGenerator HandlerGenerator
}

// HandlerGenerator 处理器生成器函数类型
type HandlerGenerator func(action string, menuKeys ...string) MiddlewareFunc

// Set 设置菜单 key
func (m *MenuResourceMarkerGenerator) Set(menu ...string) *MenuResourceMarker {
	return &MenuResourceMarker{
		HandlerGenerator: m.HandlerGenerator,
		MenuKeys:         menu,
	}
}

// MenuResourceMarker 菜单资源标记
type MenuResourceMarker struct {
	HandlerGenerator HandlerGenerator
	MenuKeys         []string `json:"menu_keys"`
	Action           string   `json:"action"`
}

// Create 创建资源中间件
func (m *MenuResourceMarker) Create() MiddlewareFunc {
	m.Action = "create"
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

// Update 更新资源中间件
func (m *MenuResourceMarker) Update() MiddlewareFunc {
	m.Action = "update"
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

// Read 读取资源中间件
func (m *MenuResourceMarker) Read() MiddlewareFunc {
	m.Action = "read"
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

// Delete 删除资源中间件
func (m *MenuResourceMarker) Delete() MiddlewareFunc {
	m.Action = "delete"
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

// Import 导入资源中间件
func (m *MenuResourceMarker) Import() MiddlewareFunc {
	m.Action = "import"
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

// Offline 下线资源中间件
func (m *MenuResourceMarker) Offline() MiddlewareFunc {
	m.Action = "offline"
	return m.HandlerGenerator(m.Action, m.MenuKeys...)
}

// checkResourceAuth 检查资源权限的核心逻辑
func (m *Middleware) checkResourceAuth(action string, menuKeys ...string) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			requestPath := r.URL.Path

			// 内部接口不需要检查
			if strings.Contains(requestPath, "/internal/") {
				next(w, r)
				return
			}

			// 尝试从 context 里面获取
			if len(menuKeys) <= 0 {
				// 没有标注的，放过
				menuKeyStr := r.Context().Value(middleware.API_MENU_KEY)
				if menuKeyStr == nil {
					next(w, r)
					return
				}
				menuKeysStr, ok := menuKeyStr.(string)
				if !ok {
					next(w, r)
					return
				}
				menuKeys = strings.Split(menuKeysStr, ",")
				if len(menuKeys) <= 0 {
					next(w, r)
					return
				}
			}

			// 获取用户 ID
			userID := GetUserIDFromContext(r)
			if userID == "" {
				respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
				return
			}

			// 检查菜单权限
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
				result, err := m.authorization.OperationCheck(r.Context(), authReq)
				if err != nil {
					log.Errorf("CheckUserPermission Error %v", err.Error())
					respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.AuthServiceException))
					return
				}

				// 如果有一个允许，就可以访问
				if result.Result {
					// 设置范围到 context
					ctx := context.WithValue(r.Context(), interception.PermissionScope, result.OperationScope())
					next(w, r.WithContext(ctx))
					return
				}
			}

			// 被拒绝了
			menuKeyCNNames := middleware.GetMenuComments(menuKeys)
			respondWithError(w, http.StatusForbidden, errorcode.Desc(errorcode.PermissionCheckFailure, strings.Join(menuKeyCNNames, ",")))
		}
	}
}

// SetMenuKeys 设置菜单 key 到 context 的中间件（用于 gin.Group 模式）
func SetMenuKeys(menus ...string) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), middleware.API_MENU_KEY, strings.Join(menus, ","))
			next(w, r.WithContext(ctx))
		}
	}
}
