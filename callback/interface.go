package callback

import (
	grpc_asset_portal_v1 "github.com/kweaver-ai/idrm-go-common/callback/asset_portal/v1"
	grpc_data_application_service_register "github.com/kweaver-ai/idrm-go-common/callback/data_application_service/register"
	grpc_data_catalog_v1 "github.com/kweaver-ai/idrm-go-common/callback/data_catalog/v1"
	grpc_task_center_v1 "github.com/kweaver-ai/idrm-go-common/callback/task_center/v1"
)

type Interface interface {
	AssetPortalV1() grpc_asset_portal_v1.Interface
	TaskCenterV1() grpc_task_center_v1.Interface
	DataCatalogV1() grpc_data_catalog_v1.Interface
	UserServiceV1() grpc_data_application_service_register.Interface
}
