package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/samber/lo"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
	"github.com/kweaver-ai/idrm-go-common/util"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

func (c *ConfigurationCenterDriven) GetDepartmentsByCode(ctx context.Context, orgCode string) (*configuration_center.DepartmentObject, error) {
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/objects/%s", c.baseURL, orgCode)
	request, _ := http.NewRequest("GET", urlStr, nil)
	resp, err := c.client.Do(request)
	if err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByCode client.Do error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByCode io.ReadAll error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	dept := &configuration_center.DepartmentObject{}
	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, util.BytesToString(body))
	}
	if err = json.Unmarshal(body, dept); err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByCode json.Unmarshal error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	return dept, nil
}

func (c *ConfigurationCenterDriven) GetDepartmentsByCodeInternal(ctx context.Context, orgCode string) (*configuration_center.DepartmentObject, error) {
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/objects/%s", c.baseURL, orgCode)
	request, _ := http.NewRequest("GET", urlStr, nil)
	resp, err := c.client.Do(request)
	if err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByCodeInternal client.Do error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByCodeInternal io.ReadAll error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	dept := &configuration_center.DepartmentObject{}
	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, util.BytesToString(body))
	}
	if err = json.Unmarshal(body, dept); err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByCodeInternal json.Unmarshal error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	return dept, nil
}

func (c *ConfigurationCenterDriven) GetDepartments(ctx context.Context, orgCodes []string) ([]*configuration_center.DepartmentObject, error) {
	val := url.Values{
		"type":             []string{"organization,department"},
		"limit":            []string{"0"},
		"is_attr_returned": []string{"true"},
	}
	if len(orgCodes) > 0 {
		val.Add("ids", strings.Join(orgCodes, ","))
		val.Add("is_all", "false")
	}
	urlStr := fmt.Sprintf("%s/api/configuration-center/v1/objects/internal?%s", c.baseURL, val.Encode())
	request, _ := http.NewRequest("GET", urlStr, nil)
	resp, err := c.client.Do(request)
	if err != nil {
		log.WithContext(ctx).Errorf("GetDepartments GetDepartments client.Do error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Errorf("DrivenConfigurationCenter GetGetDepartAndSubDepartIdsDepartments io.ReadAll error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	var depts configuration_center.QueryDepartmentPageResp
	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, util.BytesToString(body))
	}
	if err = json.Unmarshal(body, &depts); err != nil {
		log.WithContext(ctx).Errorf("DrivenConfigurationCenter GetDepartments json.Unmarshal error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	return depts.Entries, nil
}

func (c *ConfigurationCenterDriven) DeleteFile(ctx context.Context, deptID string, ossID string, fileName string) (err error) {
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/objects/file/%s", c.baseURL, deptID)
	req := &configuration_center.ObjectUpdateFileReq{
		FileId:   ossID,
		FileName: fileName,
	}
	if _, err := base.Call[any](ctx, c.client, http.MethodPut, urlStr, req); err != nil {
		return err
	}
	return nil
}

func (c *ConfigurationCenterDriven) GetDepartmentsByUserID(ctx context.Context, userID string) ([]*configuration_center.DepartmentObject, error) {
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/%s/depart", c.baseURL, userID)
	request, _ := http.NewRequest("GET", urlStr, nil)
	resp, err := c.client.Do(request)
	if err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByUserID client.Do error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByUserID io.ReadAll error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	var dept []*configuration_center.DepartmentObject
	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, util.BytesToString(body))
	}
	if err = json.Unmarshal(body, &dept); err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetDepartmentsByUserID json.Unmarshal error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	return dept, nil
}

// GetChildDepartments  查询某个部门的所有子部门信息
func (c *ConfigurationCenterDriven) GetChildDepartments(ctx context.Context, orgCode string) (*base.PageResult[configuration_center.DepartmentObject], error) {
	url := c.baseURL + "/api/configuration-center/v1/objects/internal"
	args := struct {
		ID    string `query:"id"`
		ISAll bool   `query:"is_all"`
		Limit int    `query:"limit"`
	}{
		ID:    orgCode,
		ISAll: true,
		Limit: 0,
	}
	return base.GET[*base.PageResult[configuration_center.DepartmentObject]](ctx, c.client, url, args)
}

// 获取用户部门及子部门ID集合
func (c *ConfigurationCenterDriven) GetDepartAndSubDepartIds(ctx context.Context, userId string) ([]string, error) {

	userDepartment, err := c.GetDepartmentsByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}
	var iDsSubDepart string
	for _, department := range userDepartment {
		iDsSubDepart = iDsSubDepart + department.ID + ","
	}
	iDsSubDepart = strings.TrimSuffix(iDsSubDepart, ",")
	departmentList, err := c.GetDepartmentList(ctx, configuration_center.QueryPageReqParam{Offset: 1, Limit: 0, IDsSubDepart: iDsSubDepart}) //limit 0 Offset 1 not available
	if err != nil {
		return nil, err
	}
	subDepartmentIDs := make([]string, 0)
	var set = map[string]bool{}
	for _, entry := range departmentList.Entries {
		if entry.ID != "" && !set[entry.ID] {
			subDepartmentIDs = append(subDepartmentIDs, entry.ID)
			set[entry.ID] = true
		}
	}
	return subDepartmentIDs, nil
}

func (c *ConfigurationCenterDriven) GetDepartmentsByIds(ctx context.Context, ids []string) ([]*configuration_center.DepartmentObject, error) {
	ids = lo.Filter(ids, func(item string, index int) bool {
		return item != ""
	})
	if len(ids) <= 0 {
		return nil, errorcode.Desc(errorcode.PublicInvalidParameter)
	}
	urlStr := c.baseURL + "/api/internal/configuration-center/v1/objects/departments/batchByIds"
	args := struct {
		IDs []string `query:"ids"`
	}{
		IDs: ids,
	}
	departmentInfoSlice, err := base.GET[[]*configuration_center.DepartmentObject](ctx, c.client, urlStr, args)
	if err != nil {
		return nil, err
	}
	if len(departmentInfoSlice) <= 0 {
		return nil, errorcode.Desc(errorcode.PublicResourceNotFoundError)
	}
	return departmentInfoSlice, nil
}

// 获取用户主部门及子部门ID集合
func (c *ConfigurationCenterDriven) GetMainDepartIdsByUserID(ctx context.Context, userID string) ([]string, error) {
	urlStr := fmt.Sprintf("%s/api/internal/configuration-center/v1/user/%s/main-depart-ids", c.baseURL, userID)
	request, _ := http.NewRequest("GET", urlStr, nil)
	resp, err := c.client.Do(request)
	if err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetMainDepartIdsByUserID client.Do error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetMainDepartIdsByUserID io.ReadAll error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	var dept []string
	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, util.BytesToString(body))
	}
	if err = json.Unmarshal(body, &dept); err != nil {
		log.WithContext(ctx).Errorf("ConfigurationCenterDriven GetMainDepartIdsByUserID json.Unmarshal error, %v", err)
		return nil, errorcode.Detail(errorcode.GetDepartmentPrecision, err)
	}
	return dept, nil
}

func (c *ConfigurationCenterDriven) GetDepartmentInMap(ctx context.Context, ids []string) (res map[string]*configuration_center.DepartmentObject, err error) {
	infoSlice, err := c.GetDepartmentsByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(infoSlice, func(item *configuration_center.DepartmentObject) (string, *configuration_center.DepartmentObject) {
		return item.ID, item
	}), nil
}
