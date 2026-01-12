package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/kweaver-ai/idrm-go-common/access_control"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-common/rest/workflow"
	"github.com/kweaver-ai/idrm-go-common/rest/workflow/impl"
	"github.com/kweaver-ai/idrm-go-common/util"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	af_trace "github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
)

type ConfigurationCenterDriven struct {
	baseURL  string
	client   *http.Client
	workflow workflow.WorkflowDriven
}

func (c *ConfigurationCenterDriven) QueryBaseUserByIds(ctx context.Context, ids []string) ([]*configuration_center.UserBase, error) {
	return make([]*configuration_center.UserBase, 0), nil
}

func NewConfigurationCenterDrivenByService(client *http.Client) configuration_center.Driven {
	return &ConfigurationCenterDriven{
		client:   client,
		baseURL:  base.Service.ConfigurationCenterHost,
		workflow: impl.NewWorkflowDriven(client),
	}
}

func NewConfigurationCenterDriven(client *http.Client, configurationCenterUrl, workflowRestUrl, docAuditRestUrl string) configuration_center.Driven {
	return &ConfigurationCenterDriven{
		client:   client,
		baseURL:  configurationCenterUrl,
		workflow: impl.NewWorkflowDriven(client),
	}
}

func (c *ConfigurationCenterDriven) HasAccessPermission(ctx context.Context, uid string, accessType access_control.AccessType, resource access_control.Resource) (bool, error) {
	var err error
	ctx, span := af_trace.StartInternalSpan(ctx)
	defer func() { af_trace.TelemetrySpanEnd(span, err) }()
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/access-control", c.baseURL)
	query := map[string]string{
		"access_type": strconv.Itoa(int(accessType.ToInt32())),
		"resource":    strconv.Itoa(int(resource.ToInt32())),
	}
	if uid != "" {
		query["user_id"] = uid
	}
	params := make([]string, 0, len(query))
	for k, v := range query {
		params = append(params, k+"="+v)
	}
	if len(params) > 0 {
		urlStr = urlStr + "?" + strings.Join(params, "&")
	}

	request, _ := http.NewRequest("GET", urlStr, nil)
	token, err := util.GetToken(ctx)
	if err != nil {
		return false, err
	}
	request.Header.Set("Authorization", token)
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.WithContext(ctx).Error("DrivenConfigurationCenter HasAccessPermission client.Do error", zap.Error(err))
		return false, errorcode.Detail(errorcode.GetAccessPermissionFailure, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error("DrivenConfigurationCenter HasAccessPermission io.ReadAll error", zap.Error(err))
		return false, errorcode.Detail(errorcode.GetAccessPermissionFailure, err)
	}
	var has bool
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &has)
		if err != nil {
			log.WithContext(ctx).Error("DrivenConfigurationCenter HasAccessPermission jsoniter.Unmarshal error", zap.Error(err))
			return false, errorcode.Detail(errorcode.GetAccessPermissionFailure, err)
		}
		return has, nil
	}
	return false, nil
}

func (c *ConfigurationCenterDriven) GetRolesInfo(ctx context.Context, rid, uid string) (bool, error) {
	var err error
	ctx, span := af_trace.StartInternalSpan(ctx)
	defer func() { af_trace.TelemetrySpanEnd(span, err) }()

	errorMsg := "DrivenConfigurationCenter GetRolesInfo "

	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/roles/%s/%s", c.baseURL, rid, uid)
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	token, err := util.GetToken(ctx)
	if err != nil {
		return false, err
	}
	request.Header.Set("Authorization", token)
	resp, err := c.client.Do(request)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"client.Do error", zap.Error(err))
		return false, errorcode.Detail(errorcode.GetRolesInfo, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"io.ReadAll", zap.Error(err))
		return false, errorcode.Detail(errorcode.GetRolesInfo, err.Error())
	}
	var has bool
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &has)
		if err != nil {
			log.WithContext(ctx).Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return false, errorcode.Detail(errorcode.GetRolesInfo, err)
		}
		return has, nil
	} else if resp.StatusCode == http.StatusUnauthorized {
		return false, errorcode.Detail(errorcode.GetRolesInfo, errors.New("401 UserNotLogin"))
	} else if resp.StatusCode == http.StatusForbidden {
		return false, errorcode.Detail(errorcode.GetRolesInfo, errors.New("403 UserNotHavePermission"))
	} else {
		log.WithContext(ctx).Error(errorMsg+"http status error", zap.String("status", resp.Status))
		return false, errorcode.Desc(errorcode.GetRolesInfo, errors.New("http status error: "+resp.Status))
	}
}

// HasRoles 是否有其中角色之一，带token接口, 无需uid
func (c *ConfigurationCenterDriven) HasRoles(ctx context.Context, roles ...string) (bool, error) {
	var err error
	ctx, span := af_trace.StartInternalSpan(ctx)
	defer func() { af_trace.TelemetrySpanEnd(span, err) }()

	path := "/api/configuration-center/v1/users/roles"
	url := c.baseURL + path
	//处理参数
	resp, err := base.Call[[]configuration_center.Role](ctx, c.client, http.MethodGet, url, base.EmptyArgs{})
	if err != nil {
		return false, err
	}
	rolesIDStr := strings.Join(roles, ",")
	for _, role := range resp {
		if strings.Contains(rolesIDStr, role.ID) {
			return true, nil
		}
	}
	return false, err
}

// GetInfoSystemsPrecision 批量查询 ids 和 names 二选一，另外一个可传nil
func (c *ConfigurationCenterDriven) GetInfoSystemsPrecision(ctx context.Context, ids []string, names []string) ([]*configuration_center.GetInfoSystemByIdsRes, error) {
	errorMsg := "DrivenConfigurationCenter GetInfoSystemsPrecision "
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/info-system/precision", c.baseURL)
	if len(ids) > 0 {
		params := make([]string, 0, len(ids))
		for _, id := range ids {
			params = append(params, "ids="+id)
		}
		if len(params) > 0 {
			urlStr = urlStr + "?" + strings.Join(params, "&")
		}
	} else {
		params := make([]string, 0, len(names))
		for _, name := range names {
			params = append(params, "names="+name)
		}
		if len(params) > 0 {
			urlStr = urlStr + "?" + strings.Join(params, "&")
		}
	}

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
	}
	res := make([]*configuration_center.GetInfoSystemByIdsRes, 0)
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetInfoSystemDetail, errors.New("http status error: "+resp.Status))
		}
	}
}

// GetDataSourcePrecision 批量查询
func (c *ConfigurationCenterDriven) GetDataSourcePrecision(ctx context.Context, ids []string) ([]*configuration_center.DataSourcesPrecision, error) {
	errorMsg := "DrivenConfigurationCenter GetDataSourcePrecision "
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/datasource/precision", c.baseURL)

	params := make([]string, 0, len(ids))
	for _, id := range ids {
		params = append(params, "ids="+id)
	}
	if len(params) > 0 {
		urlStr = urlStr + "?" + strings.Join(params, "&")
	}

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
	}
	res := make([]*configuration_center.DataSourcesPrecision, 0)
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetInfoSystemDetail, errors.New("http status error: "+resp.Status))
		}
	}
}

// GetDataSourcesByType 查询数据源
func (c *ConfigurationCenterDriven) GetDataSourcesByType(ctx context.Context, dataBaseType string) ([]*configuration_center.DataSourcePage, error) {
	errorMsg := "DrivenConfigurationCenter GetDataSourcesByType "
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/datasource?type=%s&limit=2000", c.baseURL, dataBaseType)

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	token, err := util.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", token)
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
	}
	res := make([]*configuration_center.DataSourcePage, 0)
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetInfoSystemDetail, errors.New("http status error: "+resp.Status))
		}
	}
}

// GetDatasourcesByHuaAoID 根据第三方数据源 ID 查询数据源
func (c *ConfigurationCenterDriven) GetDatasourcesByHuaAoID(ctx context.Context, huaAoID string) (*configuration_center.DataSourcePage, error) {
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/datasource", c.baseURL)
	result, err := base.GET[base.PageResult[configuration_center.DataSourcePage]](ctx, c.client, urlStr, configuration_center.GetDatasourcesOptions{HuaAoID: huaAoID})
	if err != nil {
		return nil, err
	}
	for _, e := range result.Entries {
		if e.HuaAoID == huaAoID {
			return e, nil
		}
	}
	return nil, errorcode.Detail(errorcode.DatasourceHuaAoIDNotFound, map[string]string{"hua_ao_id": huaAoID})
}

// GetAllDataSources 查询所有数据源
func (c *ConfigurationCenterDriven) GetAllDataSources(ctx context.Context) ([]*configuration_center.DataSources, error) {
	errorMsg := "DrivenConfigurationCenter GetAllDataSources "
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/datasource/all", c.baseURL)

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
	}
	res := make([]*configuration_center.DataSources, 0)
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		if err != nil {
			log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetInfoSystemDetail, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetInfoSystemDetail, errors.New("http status error: "+resp.Status))
		}
	}
}

// GetDepartmentList 获取部门列表，
func (c *ConfigurationCenterDriven) GetDepartmentList(ctx context.Context, req configuration_center.QueryPageReqParam) (res *configuration_center.QueryPageReapParam, err error) {
	errorMsg := "DrivenConfigurationCenter GetDepartmentList "

	token, err := util.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/configuration-center/v1/objects?id=%s&offset=%d&limit=%d&ids=%s", c.baseURL, req.ID, req.Offset, req.Limit, req.IDsSubDepart), nil)
	request.Header.Set("Authorization", token)

	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetDepartmentList, errors.New("http status error: "+resp.Status))
		}
	}

	res = &configuration_center.QueryPageReapParam{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
	}
	return res, nil
}

// GetDepartmentList 获取部门列表，
func (c *ConfigurationCenterDriven) GetDepartmentListInternal(ctx context.Context, req configuration_center.QueryPageReqParam) (res *configuration_center.QueryPageReapParam, err error) {
	errorMsg := "DrivenConfigurationCenter GetDepartmentList "

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/configuration-center/v1/objects/internal?id=%s&offset=%d&limit=%d&ids=%s&third_dept_id=%s", c.baseURL, req.ID, req.Offset, req.Limit, req.IDsSubDepart, req.ThirdDeptID), nil)
	if err != nil {
		log.Error(errorMsg+"http.NewRequest error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
	}
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
	}
	log.Debug("http response", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetDepartmentList, errors.New("http status error: "+resp.Status))
		}
	}

	res = &configuration_center.QueryPageReapParam{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentList, err.Error())
	}
	return res, nil
}

// GetDepartmentPrecision 获取部门
func (c *ConfigurationCenterDriven) GetDepartmentPrecision(ctx context.Context, ids []string) (res *configuration_center.GetDepartmentPrecisionRes, err error) {
	errorMsg := "DrivenConfigurationCenter GetDepartmentPrecision "

	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/department/precision", c.baseURL)

	params := make([]string, 0, len(ids))
	for _, id := range ids {
		params = append(params, "ids="+id)
	}
	if len(params) > 0 {
		urlStr = urlStr + "?" + strings.Join(params, "&")
	}

	log.Infof(errorMsg+" url:%s \n", urlStr)

	request, _ := http.NewRequest(http.MethodGet, urlStr, nil)

	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		log.Error(errorMsg+"client.Do error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(errorMsg+"io.ReadAll", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			res := new(errorcode.ErrorCodeFullInfo)
			if err = jsoniter.Unmarshal(body, res); err != nil {
				log.Error(errorMsg+"400 error jsoniter.Unmarshal", zap.Error(err))
				return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err.Error())
			}
			log.Error(errorMsg+"400 error", zap.String("code", res.Code), zap.String("description", res.Description))
			return nil, errorcode.New(res.Code, res.Description, res.Cause, res.Solution, res.Detail, "")
		} else {
			log.Error(errorMsg+"http status error", zap.String("status", resp.Status))
			return nil, errorcode.Desc(errorcode.GetDepartmentPrecision, errors.New("http status error: "+resp.Status))
		}
	}

	res = &configuration_center.GetDepartmentPrecisionRes{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err.Error())
	}
	return res, nil
}

func (c *ConfigurationCenterDriven) GetDepartmentByPath(ctx context.Context, paths *configuration_center.GetDepartmentByPathReq) (res *configuration_center.GetDepartmentByPathRes, err error) {
	errorMsg := "DrivenConfigurationCenter GetDepartmentByPath "

	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/department/paths", c.baseURL)

	log.Infof(errorMsg+" url:%s \n req : %v", urlStr, paths)

	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodPost, urlStr, c.client, paths)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetDepartmentByPathError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDepartmentByPathError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}

	res = &configuration_center.GetDepartmentByPathRes{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDepartmentByPathError, err.Error())
	}
	return res, nil
}

func (c *ConfigurationCenterDriven) GetProcessBindByAuditType(ctx context.Context, auditType *configuration_center.GetProcessBindByAuditTypeReq) (res *configuration_center.GetProcessBindByAuditTypeRes, err error) {
	errorMsg := "DrivenConfigurationCenter GetProcessBindByAuditType "

	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/audit-process/%s", c.baseURL, auditType.AuditType)

	log.Infof(errorMsg+" url:%s \n req : %v", urlStr, auditType)

	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodGet, urlStr, c.client, auditType)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetProcessBindByAuditTypeError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetProcessBindByAuditTypeError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}
	res = &configuration_center.GetProcessBindByAuditTypeRes{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetProcessBindByAuditTypeError, err.Error())
	}
	// 封装查绑定表后，调用workflow校验process_key的合法性
	if res.ProcDefKey != "" {
		// 若该类型未绑定审核流程，则不用校验
		processDefinition, err := c.workflow.GetAuditProcessDefinition(ctx, res.ProcDefKey)
		if err != nil {
			return nil, err
		}

		if processDefinition == nil || processDefinition.Key != res.ProcDefKey || processDefinition.Type != res.AuditType || processDefinition.Effectivity > 0 {
			log.Error(errorMsg+" workflow GetAuditProcessDefinition error", zap.String("processDefinition.Key", processDefinition.Key), zap.String("res.ProcDefKey", res.ProcDefKey))
			return nil, errorcode.Desc(errorcode.AuditProcessNotExist)
		}
	}

	return res, nil

}

func (c *ConfigurationCenterDriven) GetProcessBindByResourceId(ctx context.Context, id string) (res *configuration_center.GetProcessBindByAuditTypeRes, err error) {
	errorMsg := "DrivenConfigurationCenter GetProcessBindByResourceId "

	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/audit_policy/resource/%s/audit-process", c.baseURL, id)

	log.Infof(errorMsg+" url:%s \n req : %v", urlStr, id)

	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodGet, urlStr, c.client, nil)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetProcessBindByAuditTypeError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetProcessBindByAuditTypeError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}
	res = &configuration_center.GetProcessBindByAuditTypeRes{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetProcessBindByAuditTypeError, err.Error())
	}
	// 封装查绑定表后，调用workflow校验process_key的合法性
	if res.ProcDefKey != "" {
		// 若该类型未绑定审核流程，则不用校验
		processDefinition, err := c.workflow.GetAuditProcessDefinition(ctx, res.ProcDefKey)
		if err != nil {
			return nil, err
		}
		if processDefinition == nil || processDefinition.Key != res.ProcDefKey || processDefinition.Type != res.AuditType || processDefinition.Effectivity > 0 {
			log.Error(errorMsg+" workflow GetAuditProcessDefinition error", zap.String("processDefinition.Key", processDefinition.Key), zap.String("res.ProcDefKey", res.ProcDefKey))
			return nil, errorcode.Desc(errorcode.AuditProcessNotExist)
		}
	}

	return res, nil

}

func (c *ConfigurationCenterDriven) DeleteProcessBindByAuditType(ctx context.Context, auditType *configuration_center.DeleteProcessBindByAuditTypeReq) (err error) {
	errorMsg := "DrivenConfigurationCenter GetProcessBindByAuditType "

	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/audit-process/%s", c.baseURL, auditType.AuditType)

	log.Infof(errorMsg+" url:%s \n req : %v", urlStr, auditType)

	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodDelete, urlStr, c.client, auditType)
	if err != nil {
		return errorcode.Detail(errorcode.DeleteProcessBindByAuditTypeError, err.Error())
	}

	if statusCode != http.StatusOK {
		return errorcode.Detail(errorcode.DeleteProcessBindByAuditTypeError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}
	return nil

}

// UsersRoles implements configuration_center.Driven.
func (c *ConfigurationCenterDriven) UsersRoles(ctx context.Context) ([]configuration_center.Role, error) {
	url := fmt.Sprintf("%s/api/configuration-center/v1/users/roles", c.baseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		log.Error("create http request fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	interception.SeAuthorizationIfEmpty(ctx, req.Header)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read http response body fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("get user roles fail", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.UsersRolesFailure, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	var roles []configuration_center.Role
	if err := json.Unmarshal(body, &roles); err != nil {
		log.Error("decode response body fail", zap.Error(err), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	return roles, nil
}

// GetRoleIDs implements configuration_center.Driven.
func (c *ConfigurationCenterDriven) GetRoleIDs(ctx context.Context, userID string) (roleIDs []string, err error) {
	// url := fmt.Sprintf("%s/api/internal/configuration-center/v1/role-ids", c.baseURL)
	base, err := url.Parse(c.baseURL)
	if err != nil {
		log.Error("parse configuration-center url fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	// url path
	base.Path = path.Join(base.Path, "/api/internal/configuration-center/v1/role-ids")
	// url query
	query := make(url.Values)
	query.Add("user_id", userID)
	base.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		log.Error("create http request fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read http response body fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("get role ids fail", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.UsersRolesFailure, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	if err := json.Unmarshal(body, &roleIDs); err != nil {
		log.Error("decode response body fail", zap.Error(err), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	return roleIDs, nil
}

// ListApplicationsByDeveloperID implements configuration_center.Driven.
func (c *ConfigurationCenterDriven) ListApplicationsByDeveloperID(ctx context.Context, id string) ([]configuration_center.Application, error) {
	base, err := url.Parse(c.baseURL)
	if err != nil {
		log.Error("parse configuration-center url fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}
	// url path
	base.Path = path.Join(base.Path, "/api/internal/configuration-center/v1/apps/application-developer", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		log.Error("create http request fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read http response body fail", zap.Error(err))
		return nil, errorcode.Detail(errorcode.PublicInternalError, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Error("list application  fail", zap.Int("statusCode", resp.StatusCode), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.UsersRolesFailure, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	var result []configuration_center.Application
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error("decode response body fail", zap.Error(err), zap.ByteString("body", body))
		return nil, errorcode.Detail(errorcode.UnmarshalResponseError, map[string]any{"statusCode": resp.StatusCode, "body": json.RawMessage(body)})
	}

	return result, nil
}

type GetDataUsingTypeRes struct {
	Using int `json:"using" example:"1,2"`
}

func (c *ConfigurationCenterDriven) GetUsingType(ctx context.Context) (t int, err error) {
	errorMsg := "DrivenConfigurationCenter GetUsingType "

	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/data/using", c.baseURL)

	log.Infof(errorMsg+" url:%s \n ", urlStr)

	statusCode, body, err := base.DOWithOutToken(ctx, errorMsg, http.MethodGet, urlStr, c.client, nil)
	if err != nil {
		return 0, err
	}

	if statusCode != http.StatusOK {
		return 0, err
	}

	res := &GetDataUsingTypeRes{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return 0, err
	}
	return res.Using, nil

}

// GetCssjjSwitch 获取长沙数据局开关配置
func (c *ConfigurationCenterDriven) GetCssjjSwitch(ctx context.Context) (bool, error) {
	return c.GetGlobalSwitch(ctx, configuration_center.CSSJJ_SWITCH_KEY)
}

func (c *ConfigurationCenterDriven) GetThirdPartySwitch(ctx context.Context) (bool, error) {
	return c.GetGlobalSwitch(ctx, configuration_center.THIRD_PARTY_KEY)
}

func (c *ConfigurationCenterDriven) Generate(ctx context.Context, id string, count int) (*configuration_center.CodeList, error) {
	msg := "DrivenConfigurationCenter Generate "
	if count == 0 {
		return &configuration_center.CodeList{TotalCount: 0}, nil
	}
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/code-generation-rules/%s/generation", c.baseURL, id)
	log.Info("ConfigurationCenterDriven Generate request", zap.Int("count", count), zap.String("url", urlStr))

	statusCode, body, err := base.DOWithToken(ctx, msg, http.MethodPost, urlStr, c.client, &GenerateRequest{Count: count})
	if err != nil {
		return nil, errorcode.Detail(errorcode.GenerateCodeError, err.Error())
	}

	log.WithContext(ctx).Info("response", zap.Int("statusCode", statusCode), zap.ByteString("body", body))

	if statusCode != http.StatusOK {
		return nil, base.StatusCodeNotOK(msg, statusCode, body)
	}

	list := &configuration_center.CodeList{}
	if err := json.Unmarshal(body, list); err != nil {
		return nil, err
	}
	return list, nil
}

type GenerateRequest struct {
	Count int `json:"count,omitempty"`
}

// GetGlobalSwitch  获取配置中心全局配开关
func (c *ConfigurationCenterDriven) GetGlobalSwitch(ctx context.Context, key string) (bool, error) {
	v, err := c.GetGlobalConfig(ctx, key)
	return strings.ToLower(v) == "yes" || strings.ToLower(v) == "true", err
}

// GetGlobalConfig  获取配置中心全局配置通用接口
func (c *ConfigurationCenterDriven) GetGlobalConfig(ctx context.Context, key string) (string, error) {
	const path = "/api/internal/configuration-center/v1/config-value"

	args := &configuration_center.GetConfigValueReq{
		Key: key,
	}

	resp, err := base.Call[configuration_center.GetConfigValueRes](ctx, c.client, http.MethodGet, c.baseURL+path, args)
	if err != nil {
		return "", err
	}
	return resp.Value, nil
}

// PutGlobalConfig 修改配置中心全局配置通用接口
func (c *ConfigurationCenterDriven) PutGlobalConfig(ctx context.Context, key, value string) error {
	const path = "/api/configuration-center/v1/config-value"

	args := &configuration_center.GetConfigValueRes{
		Key:   key,
		Value: value,
	}

	if _, err := base.Call[any](ctx, c.client, http.MethodPut, c.baseURL+path, args); err != nil {
		return err
	}
	return nil
}

func (c *ConfigurationCenterDriven) GetCheckUserPermission(ctx context.Context, permissionId, uid string) (bool, error) {
	var err error
	ctx, span := af_trace.StartInternalSpan(ctx)
	defer func() { af_trace.TelemetrySpanEnd(span, err) }()

	errorMsg := "DrivenConfigurationCenter GetCheckUserPermission "

	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/permission/check/%s/%s", c.baseURL, permissionId, uid)
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	//token, err := util.GetToken(ctx)
	//if err != nil {
	//	return false, err
	//}
	//request.Header.Set("Authorization", token)
	resp, err := c.client.Do(request)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"client.Do error", zap.Error(err))
		return false, errorcode.Detail(errorcode.GetRolesInfo, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error(errorMsg+"io.ReadAll", zap.Error(err))
		return false, errorcode.Detail(errorcode.GetRolesInfo, err.Error())
	}
	var has bool
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &has)
		if err != nil {
			log.WithContext(ctx).Error(errorMsg+" json.Unmarshal error", zap.Error(err))
			return false, errorcode.Detail(errorcode.GetRolesInfo, err)
		}
		return has, nil
	} else if resp.StatusCode == http.StatusUnauthorized {
		return false, errorcode.Detail(errorcode.GetRolesInfo, errors.New("401 UserNotLogin"))
	} else if resp.StatusCode == http.StatusForbidden {
		return false, errorcode.Detail(errorcode.GetRolesInfo, errors.New("403 UserNotHavePermission"))
	} else {
		log.WithContext(ctx).Error(errorMsg+"http status error", zap.String("status", resp.Status))
		return false, errorcode.Desc(errorcode.GetRolesInfo, errors.New("http status error: "+resp.Status))
	}
}

func (c *ConfigurationCenterDriven) GetBusinessMatters(ctx context.Context, ids []string) (res []*configuration_center.BusinessMattersObject, err error) {
	if len(ids) == 0 {
		return res, nil
	}
	var idString string
	for i, id := range ids {
		if i == 0 {
			idString += id
		} else {
			idString += "," + id
		}
	}
	errorMsg := "DrivenConfigurationCenter GetBusinessMatters "
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/business_matters/list/%s", c.baseURL, idString)
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)

	resp, err := c.client.Do(request)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		return
	} else {
		return res, errors.New(errorMsg + string(body))
	}
}

func (c *ConfigurationCenterDriven) GetBusinessMatterPage(ctx context.Context, req *configuration_center.GetBusinessMatterPageReq) (res *configuration_center.GetBusinessMatterPageRes, err error) {
	errorMsg := "DrivenConfigurationCenter GetBusinessMatterPage "
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/business_matters?limit=%d", c.baseURL, req.Limit)
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)

	resp, err := c.client.Do(request)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	if resp.StatusCode == http.StatusOK {
		err = jsoniter.Unmarshal(body, &res)
		return
	} else {
		return res, errors.New(errorMsg + string(body))
	}
}
