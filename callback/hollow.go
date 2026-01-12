package callback

import (
	asset_portal_v1 "github.com/kweaver-ai/idrm-go-common/callback/asset_portal/v1"
	data_application_service_register "github.com/kweaver-ai/idrm-go-common/callback/data_application_service/register"
	data_catalog_v1 "github.com/kweaver-ai/idrm-go-common/callback/data_catalog/v1"
	task_center_v1 "github.com/kweaver-ai/idrm-go-common/callback/task_center/v1"
)

// Hollow 代表空接口，无任何实际操作
type Hollow struct{}

// AssetPortalV1 implements Interface.
func (h *Hollow) AssetPortalV1() asset_portal_v1.Interface {
	return &asset_portal_v1.Hollow{}
}

// TaskCenterV1 implements Interface.
func (Hollow) TaskCenterV1() task_center_v1.Interface {
	return &task_center_v1.Hollow{}
}

// DataCatalogV1 implements Interface.
func (Hollow) DataCatalogV1() data_catalog_v1.Interface {
	return &data_catalog_v1.Hollow{}
}

// UserServiceV1 implements Interface.
func (Hollow) UserServiceV1() data_application_service_register.Interface {
	return &data_application_service_register.Hollow{}
}

var _ Interface = &Hollow{}
