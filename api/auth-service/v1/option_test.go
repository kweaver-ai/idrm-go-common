package v1

import (
	"net/url"
	"testing"

	"github.com/kweaver-ai/idrm-go-frame/core/logx/zapx"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func TestOption(t *testing.T) {
	base, err := url.Parse("http://localhost:8080/api/auth-service/v1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("base: %+v", base)

	opts := &SubjectObjectsListOptions{
		SubjectType: SubjectUser,
		SubjectID:   "0190bb0c-66d1-75b5-b877-3f22c309b59e",
		ObjectTypes: []ObjectType{
			ObjectAPI,
			ObjectDataView,
		},
	}
	q, err := opts.MarshalQueryParameter()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("query: %s", q)

	base.RawQuery = q
	t.Logf("base: %s", base)

	log.InitLogger(zapx.LogConfigs{
		Logs: []zapx.Options{
			{
				Name:         t.Name(),
				Default:      true,
				EnableCaller: true,
				CoreConfigs: []zapx.CoreConfig{
					{
						Destination:  zapx.ConsoleDestination,
						CoreType:     zapx.StdoutCore,
						EnableColor:  true,
						OutputFormat: zapx.ConsoleFormat,
						LogLevel:     zapx.DebugLevel.String(),
					},
				},
			},
		},
	}, &telemetry.Config{})

	log.Debug("hello, debug")
	log.Info("hello, info")
}
