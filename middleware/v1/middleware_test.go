package v1

import (
	"testing"

	"github.com/kweaver-ai/idrm-go-frame/core/logx/zapx"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/common"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func TestMain(m *testing.M) {
	// 初始化日志器，避免记录日志时 panic
	log.InitLogger(zapx.LogConfigs{}, &common.TelemetryConf{})

	m.Run()
}
