package impl

import (
	"context"

	driven "github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

func (d drivenImpl) SetResource(ctx context.Context, id string, req *driven.ResourceConfig) error {
	uri := "/api/authorization/v1/resource_type/" + id
	_, err := base.PUT[any](ctx, d.httpClient, d.privatePath(uri), req)
	return err
}

func (d drivenImpl) GetResource(ctx context.Context, id string) (*driven.ResourceConfig, error) {
	uri := "/api/authorization/v1/resource_type/" + id
	return base.GET[*driven.ResourceConfig](ctx, d.httpClient, d.publicPath(uri), nil)
}

func (d drivenImpl) DeleteResource(ctx context.Context, id string) error {
	uri := "/api/authorization/v1/resource_type/" + id
	_, err := base.DELETE[any](ctx, d.httpClient, d.publicPath(uri), nil)
	return err
}
