package impl

import (
	"context"

	driven "github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/samber/lo"
)

// MenuResourceEnforce 对功能鉴权, 菜单功能鉴权，主体只能是人
func (d drivenImpl) MenuResourceEnforce(ctx context.Context, userID string, operation string, keys ...string) (policyEnforceEffect *driven.MenuResourceEnforceResult, err error) {
	if len(keys) == 0 {
		return &driven.MenuResourceEnforceResult{
			Scope:  "all",
			Effect: true,
		}, nil
	}
	if len(keys) == 1 {
		return d.menuResourceEnforce(ctx, userID, keys[0], operation)
	}
	return d.menuResourcesEnforce(ctx, userID, keys, operation)
}

// MenuResourceEnforce  单个鉴权
func (d drivenImpl) menuResourceEnforce(ctx context.Context, userID string, key string, operation string) (*driven.MenuResourceEnforceResult, error) {
	arg := &driven.OperationCheckArgs{
		Accessor: driven.Accessor{
			ID:   userID,
			Type: driven.ACCESSOR_TYPE_USER,
		},
		Resource: driven.ResourceObject{
			ID:   key,
			Type: driven.RESOURCE_TYPE_MENUS,
		},
		Operation: []string{operation},
		Include:   []string{"operation_obligations"},
		Method:    "GET",
	}
	result, err := d.OperationCheck(ctx, arg)
	if err != nil {
		log.Errorf("CheckUserPermission Error %v", err.Error())
		return nil, err
	}
	return &driven.MenuResourceEnforceResult{
		Scope:  result.OperationScope(),
		Effect: result.Result,
	}, nil
}

// menuResourcesEnforce 对功能鉴权, 菜单功能鉴权，主体只能是人
func (d drivenImpl) menuResourcesEnforce(ctx context.Context, userID string, keys []string, operation string) (*driven.MenuResourceEnforceResult, error) {
	arg := &driven.ResourceFilterArgs{
		Accessor: driven.Accessor{
			ID:   userID,
			Type: driven.ACCESSOR_TYPE_USER,
		},
		Method: "GET",
		Resources: lo.Times(len(keys), func(index int) driven.ResourceObject {
			return driven.ResourceObject{
				ID:   keys[index],
				Type: driven.RESOURCE_TYPE_MENUS,
			}
		}),
		Operation:      []string{operation},
		AllowOperation: true,
		Include:        []string{driven.INCLUDE_OPERATION_OBLIGATIONS},
	}
	resources, err := d.ResourceFilter(ctx, arg)
	if err != nil {
		log.Errorf("menuResourcesEnforce Error %v", err.Error())
		return nil, err
	}
	for _, resource := range resources {
		if resource.Id == "" || len(resource.AllowOperation) <= 0 {
			continue
		}
		if resource.AllowOperation[0] == operation {
			return &driven.MenuResourceEnforceResult{
				Scope:  resource.OperationScope(),
				Effect: true,
			}, nil
		}
	}
	return &driven.MenuResourceEnforceResult{}, nil
}
