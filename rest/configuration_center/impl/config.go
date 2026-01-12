package impl

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"go.uber.org/zap"
)

func (c *ConfigurationCenterDriven) GetThirdPartyAddr(ctx context.Context, name string) ([]*driven.GetThirdPartyAddressRes, error) {
	var err error
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()

	errorMsg := "DrivenConfigurationCenter GetThirdPartyAddr "

	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/third_party_addr?name=%s", c.baseURL, name)
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	request.Header.Set("Authorization", ctx.Value(interception.Token).(string))
	resp, err := c.client.Do(request)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetThirdPartyAddr, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetThirdPartyAddr, err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		res := make([]*driven.GetThirdPartyAddressRes, 0)
		if err = jsoniter.Unmarshal(body, &res); err != nil {
			log.WithContext(ctx).Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetThirdPartyAddr, err.Error())
		}
		return res, nil
	} else if resp.StatusCode == http.StatusBadRequest {
		res := new(base.CommonResponse[any])
		if err = jsoniter.Unmarshal(body, res); err != nil {
			log.WithContext(ctx).Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetThirdPartyAddr, err.Error())
		}
		log.WithContext(ctx).Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
		return nil, errorcode.Detail(errorcode.GetThirdPartyAddr, res.Code)
	} else if resp.StatusCode == http.StatusUnauthorized {
		return nil, errorcode.Detail(errorcode.GetThirdPartyAddr, errors.New("401 UserNotLogin"))
	} else if resp.StatusCode == http.StatusForbidden {
		return nil, errorcode.Detail(errorcode.GetThirdPartyAddr, errors.New("401 UserNotHavePermission"))
	} else {
		log.WithContext(ctx).Error(errorMsg+"http status error", zap.String("status", resp.Status))
		return nil, errorcode.Desc(errorcode.GetThirdPartyAddr)
	}
}
