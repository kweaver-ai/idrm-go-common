package errorcode

import (
	"github.com/kweaver-ai/idrm-go-frame/core/errorx"
)

var (
	APIAuthModel = errorx.New(ServiceName + ".APIAuthModel.")
)

var (
	RequesterWithoutInnerRoleErr = APIAuthModel.Description("RequesterWithoutInnerRole", "申请者缺少内置角色")
)
