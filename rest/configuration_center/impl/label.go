package impl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"go.uber.org/zap"
)

func (c *ConfigurationCenterDriven) QueryDataGrade(ctx context.Context, id ...string) (map[string]*driven.HierarchyTag, error) {
	path := "/api/internal/configuration-center/v1/grade-label/list/:ids"
	args := driven.QueryDataGradeReq{
		Ids: strings.Join(id, ","),
	}
	url := c.baseURL + path
	//处理参数
	resp, err := base.Call[driven.QueryDataGradeResp](ctx, c.client, http.MethodGet, url, args)
	if err != nil {
		log.Errorf("QueryDataGrade error %v", err.Error())
		return nil, errorcode.Desc(errorcode.GetGradeLabelError)
	}
	result := make(map[string]*driven.HierarchyTag)
	for i := range resp.Entries {
		result[resp.Entries[i].ID] = &(resp.Entries[i])
	}
	return result, err
}

func (c *ConfigurationCenterDriven) GetLabelById(ctx context.Context, id string) (*driven.GetLabelByIdRes, error) {
	errorMsg := "DrivenConfigurationCenter GetLabelById"
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/grade-label/id/%s", c.baseURL, id)
	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	request.Header.Set("Authorization", ctx.Value(interception.Token).(string))

	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
	}
	var res *driven.GetLabelByIdRes
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetGradeLabelError)
		}
	}
}

func (c *ConfigurationCenterDriven) GetLabelByIds(ctx context.Context, ids string) (*driven.GetLabelByIdsRes, error) {
	errorMsg := "DrivenConfigurationCenter GetLabelByIds"
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/grade-label/list/%s", c.baseURL, ids)

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
	}
	var res *driven.GetLabelByIdsRes
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetGradeLabelError)
		}
	}
}

func (c *ConfigurationCenterDriven) GetLabelByName(ctx context.Context, name string) (*driven.GetLabelByIdRes, error) {
	errorMsg := "DrivenConfigurationCenter GetLabelByName"
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/grade-label/name/%s", c.baseURL, name)
	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	request.Header.Set("Authorization", ctx.Value(interception.Token).(string))
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
	}
	var res *driven.GetLabelByIdRes
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetGradeLabelError, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetGradeLabelError)
		}
	}
}
