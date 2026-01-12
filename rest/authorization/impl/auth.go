package impl

import (
	"context"
	"fmt"

	driven "github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

func (d drivenImpl) GetAccessorPolicy(ctx context.Context, req *driven.AccessorPolicyArgs) (*base.PageResult[driven.AccessorPolicy], error) {
	uri := "/api/authorization/v1/accessor-policy"
	return base.GET[*base.PageResult[driven.AccessorPolicy]](ctx, d.httpClient, d.publicPath(uri), req)
}

func (d drivenImpl) OperationCheck(ctx context.Context, req *driven.OperationCheckArgs) (*driven.OperationCheckResult, error) {
	uri := "/api/authorization/v1/operation-check"
	return base.POST[*driven.OperationCheckResult](ctx, d.httpClient, d.privatePath(uri), req)
}

func (d drivenImpl) ResourceFilter(ctx context.Context, req *driven.ResourceFilterArgs) ([]*driven.ResourceFilterResp, error) {
	uri := "/api/authorization/v1/resource-filter"
	return base.POST[[]*driven.ResourceFilterResp](ctx, d.httpClient, d.privatePath(uri), req)
}

func (d drivenImpl) GetResourceOperations(ctx context.Context, req *driven.GetResourceOperationsArgs) ([]*driven.ResourceOperations, error) {
	uri := "/api/authorization/v1/resource-operation"
	return base.POST[[]*driven.ResourceOperations](ctx, d.httpClient, d.privatePath(uri), req)
}

func (d drivenImpl) GetResourceTypeOperations(ctx context.Context, req *driven.GetResourceTypeOperationsArgs) ([]*driven.ResourceOperations, error) {
	uri := "/api/authorization/v1/resource-type-operation"
	return base.POST[[]*driven.ResourceOperations](ctx, d.httpClient, d.privatePath(uri), req)
}

func (d drivenImpl) CreatePolicy(ctx context.Context, req []*driven.CreatePolicyReq) (*driven.CreatePolicyResp, error) {
	uri := "/api/authorization/v1/policy"
	return base.POST[*driven.CreatePolicyResp](ctx, d.httpClient, d.privatePath(uri), req)
}

func (d drivenImpl) UpdatePolicy(ctx context.Context, ids string, req []*driven.UpdatePolicyReq) error {
	uri := fmt.Sprintf("/api/authorization/v1/policy/%s", ids)
	_, err := base.PUT[any](ctx, d.httpClient, d.publicPath(uri), req)
	return err
}

func (d drivenImpl) DeletePolicy(ctx context.Context, ids string) error {
	uri := "/api/authorization/v1/policy/:ids"
	args := driven.PathIDReq{
		Ids: ids,
	}
	_, err := base.DELETE[any](ctx, d.httpClient, d.publicPath(uri), args)
	return err
}

func (d drivenImpl) GetResourcePolicy(ctx context.Context, req *driven.GetResourcePolicyReq) (*base.PageResult[driven.ResourcePolicy], error) {
	uri := "/api/authorization/v1/resource-policy"
	return base.GET[*base.PageResult[driven.ResourcePolicy]](ctx, d.httpClient, d.publicPath(uri), req)
}

// ResourceList  资源列举
func (d drivenImpl) ResourceList(ctx context.Context, req *driven.ResourceListArgs) ([]*driven.ResourceListRespItem, error) {
	uri := "/api/authorization/v1/resource-list"
	resp, err := base.POST[[]*driven.ResourceListRespItem](ctx, d.httpClient, d.privatePath(uri), req)
	return resp, err
}
