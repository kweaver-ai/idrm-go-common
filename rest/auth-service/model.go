package auth_service

import (
	"github.com/kweaver-ai/idrm-go-frame/core/enum"
	"github.com/kweaver-ai/idrm-go-frame/core/transport/rest/ginx"
	"github.com/gin-gonic/gin"
)

type EnumAction enum.Object

var (
	ActionCreate  = enum.New[EnumAction](1, "create", "创建")
	ActionUpdate  = enum.New[EnumAction](2, "update", "编辑")
	ActionRead    = enum.New[EnumAction](3, "read", "查询")
	ActionDelete  = enum.New[EnumAction](4, "delete", "删除")
	ActionImport  = enum.New[EnumAction](5, "import", "导入")
	ActionOffline = enum.New[EnumAction](6, "offline", "下线")
	ActionOnline  = enum.New[EnumAction](7, "online", "上线")
)

var ActionValueList = enum.Values("EnumAction")

var manageActionDict = map[string]bool{
	ActionCreate.String:  true,
	ActionUpdate.String:  true,
	ActionDelete.String:  true,
	ActionImport.String:  true,
	ActionOffline.String: true,
}

type MenuResource struct {
	ServiceName string `json:"service_name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Action      string `json:"action"`
	Resource    string `json:"resource"`
}

// MenuResourceCheckRequest 权限请求参数
type MenuResourceCheckRequest struct {
	UserID string `json:"user_id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

// MenuResourceCheckResponse 权限请求返回结果
type MenuResourceCheckResponse struct {
	Path     string `json:"path"`
	Method   string `json:"method"`
	Action   string `json:"action"`
	Scope    string `json:"scope"`
	Resource string `json:"resource"`
	Effect   string `json:"effect"`
}

func ServiceAC(rs []*MenuResource) gin.HandlerFunc {
	return func(c *gin.Context) {
		ginx.ResOKJson(c, rs)
	}
}
