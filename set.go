package GoCommon

import (
	"github.com/google/wire"
	middlewareImpl "github.com/kweaver-ai/idrm-go-common/middleware/v1"
	"github.com/kweaver-ai/idrm-go-common/rest"
	ccImpl "github.com/kweaver-ai/idrm-go-common/rest/configuration_center/impl"
	hydraImpl "github.com/kweaver-ai/idrm-go-common/rest/hydra/impl"
	"github.com/kweaver-ai/idrm-go-common/trace"
	"github.com/kweaver-ai/idrm-go-common/workflow"
)

// Set 所有的依赖全部放一起，最省心的
var Set = wire.NewSet(
	workflow.NewWorkflow,
	middlewareImpl.NewMiddleware,
	rest.Set,
	trace.NewOtelHttpClient,
)

// Middleware 中间件的Set
var Middleware = wire.NewSet(
	hydraImpl.NewHydraByService,
	ccImpl.NewConfigurationCenterDrivenByService,
	middlewareImpl.NewMiddleware,
)
