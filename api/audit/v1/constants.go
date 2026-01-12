package v1

import (
	"github.com/mileusna/useragent"

	"github.com/kweaver-ai/idrm-go-common/util/sets"
)

// BodyTypeResourceOperation 代表数据资源操作
const BodyTypeResourceOperation = "resource_operation"

// agentWebUserAgentNameSet 定义被视为 AgentWeb 的 UserAgent.Name。UserAgent 由
// github.com/mileusna/useragent 解析 HTTP Header User-Agent 得到。
var agentWebUserAgentNameSet = sets.New(
	useragent.Chrome,
	useragent.Edge,
	useragent.Firefox,
	useragent.HeadlessChrome,
	useragent.InternetExplorer,
	useragent.Opera,
	useragent.OperaMini,
	useragent.OperaTouch,
	useragent.Safari,
	useragent.Vivaldi,
)
