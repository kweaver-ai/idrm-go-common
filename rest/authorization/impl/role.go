package impl

import (
	"context"

	driven "github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/samber/lo"
)

func (d drivenImpl) GetRole(ctx context.Context, id string) (*driven.RoleDetail, error) {
	uri := "/api/authorization/v1/roles/" + id
	return base.GET[*driven.RoleDetail](ctx, d.httpClient, d.publicPath(uri), nil)
}

func (d drivenImpl) ListRoles(ctx context.Context, req *driven.RoleListArgs) (*base.PageResult[driven.RoleDetail], error) {
	uri := "/api/authorization/v1/roles"
	return base.GET[*base.PageResult[driven.RoleDetail]](ctx, d.httpClient, d.publicPath(uri), req)
}

func (d drivenImpl) ListRoleMembers(ctx context.Context, req *driven.RoleMemberArgs) (*base.PageResult[driven.MemberInfo], error) {
	uri := "/api/authorization/v1/role-members/:id"
	return base.GET[*base.PageResult[driven.MemberInfo]](ctx, d.httpClient, d.publicPath(uri), req)
}

func (d drivenImpl) UpdateRoleMembers(ctx context.Context, req *driven.UpdateRoleMemberArgs) error {
	uri := "/api/authorization/v1/role-members/:id"
	_, err := base.POST[any](ctx, d.httpClient, d.publicPath(uri), req)
	return err
}

func (d drivenImpl) ListRoleTotalMembers(ctx context.Context, roleID string) ([]*driven.MemberInfo, error) {
	args := &driven.RoleMemberArgs{
		ID:     roleID,
		Limit:  1000,
		Offset: 0,
	}
	loopCount := 0
	total := make([]*driven.MemberInfo, 0)
	for {
		memberResp, err := d.ListRoleMembers(ctx, args)
		if err != nil {
			return nil, err
		}
		loopCount += 1
		total = append(total, memberResp.Entries...)
		if memberResp.TotalCount < 1000 {
			return total, nil
		}
		if loopCount >= 10 {
			break
		}
	}
	return total, nil
}

// ListAccessorRoles 获取访问者的角色列表
func (d drivenImpl) ListAccessorRoles(ctx context.Context, req *driven.ListAccessorRolesArgs) ([]*driven.RoleMetaInfo, error) {
	uri := "/api/authorization/v1/accessor_roles"
	roles, err := base.GET[*base.PageResult[driven.RoleMetaInfo]](ctx, d.httpClient, d.privatePath(uri), req)
	if err != nil {
		return nil, err
	}
	return roles.Entries, nil
}

// ListUserRoles 获取用户的角色列表
func (d drivenImpl) ListUserRoles(ctx context.Context, uid string) ([]*driven.RoleMetaInfo, error) {
	args := &driven.ListAccessorRolesArgs{
		AccessorID:   uid,
		AccessorType: driven.ACCESSOR_TYPE_USER,
		Limit:        1000,
	}
	return d.ListAccessorRoles(ctx, args)
}

// ListUserRoleID 获取用户的角色ID 列表
func (d drivenImpl) ListUserRoleID(ctx context.Context, uid string) ([]string, error) {
	roles, err := d.ListUserRoles(ctx, uid)
	if err != nil {
		return nil, err
	}
	return lo.Map(roles, func(item *driven.RoleMetaInfo, index int) string {
		return item.ID
	}), nil
}

// HasRoles 判断用户是都有某个角色, 如果有多个角色，有一个就返回true
func (d drivenImpl) HasRoles(ctx context.Context, uid string, roleID ...string) (bool, error) {
	roleInfos, err := d.ListUserRoles(ctx, uid)
	if err != nil {
		return false, err
	}
	roleIDMap := lo.SliceToMap(roleID, func(item string) (string, bool) {
		return item, true
	})
	for _, item := range roleInfos {
		if _, ok := roleIDMap[item.ID]; ok {
			return true, nil
		}
	}
	return false, nil
}

// HasInnerBusinessRoles 判断用户是都有某个内置角色
func (d drivenImpl) HasInnerBusinessRoles(ctx context.Context, uid string) ([]string, error) {
	roles, err := d.ListUserRoles(ctx, uid)
	if err != nil {
		return nil, err
	}
	rs := make([]string, 0)
	for _, role := range roles {
		if driven.IsInnerBusinessRole(role.ID) {
			rs = append(rs, role.ID)
		}
	}
	return rs, nil
}
