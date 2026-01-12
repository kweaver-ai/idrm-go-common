package callback

import (
	"google.golang.org/grpc"

	asset_portal_v1 "github.com/kweaver-ai/idrm-go-common/callback/asset_portal/v1"
	data_application_service_register "github.com/kweaver-ai/idrm-go-common/callback/data_application_service/register"
	data_catalog_v1 "github.com/kweaver-ai/idrm-go-common/callback/data_catalog/v1"
	task_center_v1 "github.com/kweaver-ai/idrm-go-common/callback/task_center/v1"
)

type Client struct {
	conn grpc.ClientConnInterface
}

// New 创建回调客户端
func New(conn grpc.ClientConnInterface) *Client {
	return &Client{conn: conn}
}

// AssetPortalV1 implements Interface.
func (c *Client) AssetPortalV1() asset_portal_v1.Interface {
	return asset_portal_v1.New(c.conn)
}

// TaskCenterV1 implements Interface.
func (c *Client) TaskCenterV1() task_center_v1.Interface {
	return task_center_v1.New(c.conn)
}

// DataCatalogV1 implements Interface.
func (c *Client) DataCatalogV1() data_catalog_v1.Interface {
	return data_catalog_v1.New(c.conn)
}

// UserServiceV1 implements Interface.
func (c *Client) UserServiceV1() data_application_service_register.Interface {
	return data_application_service_register.New(c.conn)
}

var _ Interface = &Client{}
