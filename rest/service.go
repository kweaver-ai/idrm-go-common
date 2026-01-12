package rest

import (
	"github.com/google/wire"
	"github.com/kweaver-ai/idrm-go-common/rest/af_sailor"
	sailor_impl "github.com/kweaver-ai/idrm-go-common/rest/af_sailor/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/af_sailor_service"
	af_sailor_service_impl "github.com/kweaver-ai/idrm-go-common/rest/af_sailor_service/impl"
	auth_service_v1 "github.com/kweaver-ai/idrm-go-common/rest/auth-service/v1"
	"github.com/kweaver-ai/idrm-go-common/rest/authorization"
	authorization_impl "github.com/kweaver-ai/idrm-go-common/rest/authorization/impl"
	basic_bigdata_service_impl "github.com/kweaver-ai/idrm-go-common/rest/basic_bigdata_service/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/business_grooming"
	business_grooming_impl "github.com/kweaver-ai/idrm-go-common/rest/business_grooming/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	configuration_center_impl "github.com/kweaver-ai/idrm-go-common/rest/configuration_center/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/data_application_service"
	data_application_service_impl "github.com/kweaver-ai/idrm-go-common/rest/data_application_service/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/data_catalog"
	data_catalog_impl "github.com/kweaver-ai/idrm-go-common/rest/data_catalog/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/data_subject"
	"github.com/kweaver-ai/idrm-go-common/rest/data_subject/data_subject_impl"
	"github.com/kweaver-ai/idrm-go-common/rest/data_view"
	data_view_impl "github.com/kweaver-ai/idrm-go-common/rest/data_view/impl"
	demand_management_impl "github.com/kweaver-ai/idrm-go-common/rest/demand_management/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/hydra"
	hydra_impl "github.com/kweaver-ai/idrm-go-common/rest/hydra/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/indicator_management"
	indicator_management_impl "github.com/kweaver-ai/idrm-go-common/rest/indicator_management/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/label"
	label_impl "github.com/kweaver-ai/idrm-go-common/rest/label/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/metadata_manage"
	metadata_manage_impl "github.com/kweaver-ai/idrm-go-common/rest/metadata_manage/impl"
	points_management_impl "github.com/kweaver-ai/idrm-go-common/rest/points_management/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/scene_analysis"
	scene_analysis_impl "github.com/kweaver-ai/idrm-go-common/rest/scene_analysis/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/standardization"
	standardization_impl "github.com/kweaver-ai/idrm-go-common/rest/standardization/impl"
	task_center "github.com/kweaver-ai/idrm-go-common/rest/task_center"
	task_center_impl "github.com/kweaver-ai/idrm-go-common/rest/task_center/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/user_management"
	"github.com/kweaver-ai/idrm-go-common/rest/virtual_engine"
	virtual_engine_impl "github.com/kweaver-ai/idrm-go-common/rest/virtual_engine/impl"
	"github.com/kweaver-ai/idrm-go-common/rest/workflow"
	workflow_impl "github.com/kweaver-ai/idrm-go-common/rest/workflow/impl"
)

type Service struct {
	BusinessGrooming       business_grooming.Driven
	TaskCenter             task_center.Driven
	ConfigurationCenter    configuration_center.Driven
	DataApplicationService data_application_service.Driven
	DataCatalog            data_catalog.Driven
	DataSubject            data_subject.Driven
	DataView               data_view.Driven
	IndicatorManagement    indicator_management.Driven
	MetadataManagement     metadata_manage.Driven
	SceneAnalysis          scene_analysis.SceneAnalysisDriven
	Standardization        standardization.Driven
	Workflow               workflow.WorkflowDriven
	VirtualEngine          virtual_engine.Driven
	Hydra                  hydra.Hydra
	BasicBigDataService    label.Driven
	Sail                   af_sailor.Driven
	SailService            af_sailor_service.Driven
	UserManagement         user_management.DrivenUserMgnt
	Authorization          authorization.Driven
}

// Set 所有rest服务的注入集合
var Set = wire.NewSet(
	wire.NewSet(wire.Struct(new(Service), "*")),
	business_grooming_impl.NewDriven,
	task_center_impl.NewDriven,
	configuration_center_impl.NewConfigurationCenterDrivenByService,
	data_application_service_impl.NewDrivenImpl,
	data_catalog_impl.NewDrivenImpl,
	data_subject_impl.NewDataViewDriven,
	data_view_impl.NewDataViewDriven,
	indicator_management_impl.NewDrivenImpl,
	metadata_manage_impl.NewDrivenImpl,
	scene_analysis_impl.NewSceneAnalysisDrivenByServiceName,
	standardization_impl.NewDriven,
	workflow_impl.NewWorkflowDriven,
	workflow_impl.NewDocAuditDriven,
	virtual_engine_impl.NewDrivenImpl,
	hydra_impl.NewHydraByService,
	basic_bigdata_service_impl.NewDriven,
	sailor_impl.NewSailorDriven,
	af_sailor_service_impl.NewSailorDriven,
	auth_service_v1.NewBaseClient,
	auth_service_v1.NewInternalForBase,
	user_management.NewUserMgntByService,
	authorization_impl.NewDriven,

	label_impl.NewBigDataDriven,
	points_management_impl.NewPointsEventPubRepoImpl,
	demand_management_impl.NewDemandManagementDriven,
)
