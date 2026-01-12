package util

import (
	"context"
	"runtime"
	"strconv"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/middleware"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func ObtainToken(c context.Context) string {
	value := c.Value(interception.Token)
	if value == nil {
		return ""
	}
	token, ok := value.(string)
	if !ok {
		return ""
	}
	return token
}

func ObtainUserInfo(ctx context.Context) *middleware.User {
	user := &middleware.User{}
	value := ctx.Value(interception.InfoName)
	if value == nil {
		return user
	}
	user, ok := value.(*middleware.User)
	if !ok {
		return user
	}
	return user
}

func GetToken(ctx context.Context) (string, error) {
	value := ctx.Value(interception.Token)
	if value == nil {
		log.WithContext(ctx).Error("GetToken Get interception.Token  not exist")
		return "", errorcode.Desc(errorcode.ContextNotHaveToken)
	}
	token, ok := value.(string)
	if !ok {
		pc, _, line, _ := runtime.Caller(1)
		log.WithContext(ctx).Error("GetToken transfer string  error" + runtime.FuncForPC(pc).Name() + " | " + strconv.Itoa(line))
		return "", errorcode.Desc(errorcode.ContextNotHaveToken)
	}
	if len(token) == 0 {
		log.WithContext(ctx).Error("Get token len 0")
		return "", errorcode.Desc(errorcode.ContextNotHaveToken)
	}
	return token, nil
}

func GetUserInfo(ctx context.Context) (*middleware.User, error) {
	//获取用户信息
	value := ctx.Value(interception.InfoName)
	if value == nil {
		log.WithContext(ctx).Error("GetUserInfo Get TokenIntrospectInfo not exist")
		return nil, errorcode.Desc(errorcode.ContextNotHaveUserInfo)
	}
	user, ok := value.(*middleware.User)
	if !ok {
		pc, _, line, _ := runtime.Caller(1)
		log.WithContext(ctx).Error("GetUserInfo transfer middleware User error" + runtime.FuncForPC(pc).Name() + " | " + strconv.Itoa(line))
		return nil, errorcode.Desc(errorcode.ContextNotHaveUserInfo)
	}
	return user, nil
}
