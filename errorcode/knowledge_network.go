package errorcode

import (
	"github.com/kweaver-ai/idrm-go-frame/core/errorx"
)

var (
	KnowledgeNetworkModel = errorx.New(ServiceName + ".KnowledgeNetwork.")
)

var (
	GetKnowledgeNetworkDetailErr = KnowledgeNetworkModel.Description("GetKnowledgeNetworkDetail", "获取知识网络详情失败")
)
