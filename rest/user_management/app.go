package user_management

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"go.uber.org/zap"
)

// GetUserManagementAppList 获取应用账户列表
func (u *usermgntSvc) GetUserManagementAppList(ctx context.Context, args *AppListArgs) (result *base.PageResult[AppEntry], err error) {
	// 构建URL，添加查询参数
	target := fmt.Sprintf("%s/v1/apps?limit=%d&offset=%d&direction=%s&sort=%s", u.publicBaseURL, args.Limit, args.Offset, args.Direction, args.Sort)
	if args.Keyword != "" {
		target = fmt.Sprintf("%s&keyword=%s", target, args.Keyword)
	}
	// 发送GET请求
	respParam, err := u.httpClientEx.Get(ctx, target, interception.AuthorizationHeaderMap(ctx))
	if err != nil {
		log.WithContext(ctx).Error("GetUserManagementAppList failed", zap.Error(err), zap.String("url", target))
		return
	}
	payload, _ := json.Marshal(respParam)
	result = &base.PageResult[AppEntry]{}
	if err = json.Unmarshal(payload, result); err != nil {
		log.WithContext(ctx).Error("GetUserManagementAppList failed", zap.Error(err), zap.String("url", target))
		return nil, errorcode.PublicDecodeJsonErr.Detail(err.Error())
	}
	return result, nil
}

// GetUserManagementApp 获取应用账户信息
func (u *usermgntSvc) GetUserManagementApp(ctx context.Context, id string) (result *AppEntry, err error) {
	// 构建URL，添加查询参数
	target := fmt.Sprintf("%s/v1/apps/%s", u.publicBaseURL, id)
	// 发送GET请求
	respParam, err := u.httpClient.Get(ctx, target, nil)
	if err != nil {
		log.WithContext(ctx).Error("GetUserManagementApp failed", zap.Error(err), zap.String("url", target))
		return
	}
	payload, _ := json.Marshal(respParam)
	result = &AppEntry{}
	if err = json.Unmarshal(payload, result); err != nil {
		log.WithContext(ctx).Error("GetUserManagementApp failed", zap.Error(err), zap.String("url", target))
		return nil, errorcode.PublicDecodeJsonErr.Detail(err.Error())
	}
	return result, nil
}
