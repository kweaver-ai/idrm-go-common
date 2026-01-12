package impl

import (
	"context"

	driven "github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

func (d drivenImpl) CreateObligationType(ctx context.Context, typeName string, req *driven.ObligationTypeReq) error {
	uri := "/api/authorization/v1/obligation-types/" + typeName
	_, err := base.PUT[any](ctx, d.httpClient, d.privatePath(uri), req)
	return err
}
