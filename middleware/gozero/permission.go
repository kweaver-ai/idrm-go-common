package gozero

import (
	"context"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	"github.com/kweaver-ai/idrm-go-common/util"
)

// PermissionControl 根据候选角色身份校验用户权限
// roleIDs: 候选角色 ID 列表
// exclude: 为 true 时，用户拥有候选列表以外身份时才有权限；为 false 时，用户拥有候选列表内身份时才有权限
func (m *Middleware) PermissionControl(roleIDs []string, exclude bool) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			_, _ = util.HandleReqWithTraceIncludingErrLog(r.Context(), func(ctx context.Context) (_ any, err error) {
				// 来自 APP Client 的 Token 允许访问
				if r.Context().Value(interception.TokenType) == interception.TokenTypeClient {
					next(w, r)
					return nil, nil
				}

				// 解析用户 ID
				var userID string
				tokenType, ok := r.Context().Value(interception.TokenType).(int)
				if ok && tokenType == interception.TokenTypeClient {
					userID = r.Header.Get("userId")
				} else if user, ok := r.Context().Value(interception.InfoName).(*middleware.User); ok {
					userID = user.ID
				}

				// 用户 ID 为空不允许访问
				if userID == "" {
					respondWithError(w, http.StatusBadRequest, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId))
					return nil, errorcode.Desc(errorcode.AccessControlClientTokenMustHasUserId)
				}

				// TODO 20250603 去掉信息资源编目接口，单独的一套验证角色权限
				// 检查用户是否拥有符合权限要求的角色身份
				// userRoleIDs, err := m.configurationCenterDriven.GetRoleIDs(ctx, userID)
				// if err != nil {
				// 	respondWithError(w, http.StatusBadRequest, err)
				// 	return nil, err
				// }
				// for _, id := range userRoleIDs {
				// 	if util.Contains(roleIDs, id) == !exclude {
				// 		next(w, r)
				// 		return nil, nil
				// 	}
				// }
				// respondWithError(w, http.StatusForbidden, errorcode.Desc(errorcode.AuthorizationFailure))
				// return nil, errorcode.Desc(errorcode.AuthorizationFailure)

				// 当前所有请求通过
				next(w, r)
				return nil, nil
			})
		}
	}
}
