package errorcode

import (
	"os"

	"github.com/kweaver-ai/idrm-go-frame/core/errorx"
)

var ServiceName = func() string {
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName != "" {
		return serviceName
	}
	return "Common"
}()

var (
	publicModule   = errorx.New(ServiceName + ".Public.")
	workflowModule = errorx.New(ServiceName + ".Workflow.")
)

var (
	PublicDatabaseErr        = publicModule.Description("DatabaseError", "数据库异常")
	PublicQueryUserInfoError = publicModule.Description("QueryUserInfoError", "查询用户信息错误")
	PublicInternalErr        = publicModule.Description("InternalError", "内部错误")
	PublicDecodeJsonErr      = publicModule.Description("DecodeJsonErr", "解析json错误")
)

var (
	AuditProcessNotExistErr = workflowModule.Description("AuditProcessNotExistErr", "审核策略不存在或配置错误,请检查审核策略")
	AuditMsgSendErr         = workflowModule.Description("AuditMsgSendErr", "审核消息发送错误")
)
