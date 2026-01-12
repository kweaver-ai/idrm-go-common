package impl

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/araddon/dateparse"
	"github.com/jinzhu/copier"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_view"
)

func (d *DrivenImpl) GetExploreReport(ctx context.Context, id string, thirdParty bool) (*driven.ExploreReportResp, error) {
	urlStr := d.baseURL + "/api/data-view/v1/form-view/explore-report"
	args := struct {
		ID         string `query:"id"`
		ThirdParty bool   `query:"third_party"`
	}{
		ID:         id,
		ThirdParty: thirdParty,
	}
	resp, err := base.GET[*driven.ExploreReportResp](ctx, d.httpClient, urlStr, args)
	return resp, err
}

func (d *DrivenImpl) GetViewExploreReport(ctx context.Context, id string, version *int32) (*driven.ExploreReportResp, error) {
	urlStr := d.baseURL + "/api/internal/data-view/v1/form-view/explore-report"
	args := struct {
		ID      string `query:"id"`
		version *int32 `query:"version"`
	}{
		ID:      id,
		version: version,
	}
	resp, err := base.GET[*driven.ExploreReportResp](ctx, d.httpClient, urlStr, args)
	return resp, err
}

func (d *DrivenImpl) GetViewBusinessUpdateTime(ctx context.Context, id string) (*driven.GetBusinessUpdateTimeResp, error) {
	urlStr := d.baseURL + "/api/data-view/v1/form-view/:id/business-update-time"
	args := struct {
		ID string `uri:"id"`
	}{
		ID: id,
	}
	resp, err := base.GET[*driven.GetBusinessUpdateTimeResp](ctx, d.httpClient, urlStr, args)
	return resp, err
}

// QueryViewExplore 查询整理探查结果
func (d *DrivenImpl) QueryViewExplore(ctx context.Context, formViewID string, thirdParty bool) (*driven.ViewExploreDetail, error) {
	exploreResult := &driven.ViewExploreDetail{}
	//step1:查询探查结果
	exportDetail, err := d.GetExploreReport(ctx, formViewID, thirdParty)
	if err != nil {
		return exploreResult, err
	}
	copier.Copy(exploreResult, exportDetail.Overview)
	if exploreResult.ConsistencyScore != nil {
		*exploreResult.ConsistencyScore *= float64(100)
	}
	if exploreResult.CompletenessScore != nil {
		*exploreResult.CompletenessScore *= float64(100)
	}
	if exploreResult.AccuracyScore != nil {
		*exploreResult.AccuracyScore *= float64(100)
	}
	if exploreResult.StandardizationScore != nil {
		*exploreResult.StandardizationScore *= float64(100)
	}
	if exploreResult.UniquenessScore != nil {
		*exploreResult.UniquenessScore *= float64(100)
	}
	//step2: 接下来解析及时性
	if len(exportDetail.ExploreViewDetails) <= 0 {
		return exploreResult, nil
	}
	var ruleResult *driven.RuleResult
	for _, ruleInfo := range exportDetail.ExploreViewDetails {
		if ruleInfo.Dimension != "timeliness" {
			continue
		}
		ruleResult = ruleInfo
	}
	if ruleResult == nil {
		return exploreResult, nil
	}
	//查询最后更修时间，确定及时性
	lastUpdateTimeInfo, err := d.GetViewBusinessUpdateTime(ctx, formViewID)
	if err != nil {
		return exploreResult, nil
	}
	ruleConfig := &driven.RuleConfig{}
	if err = json.Unmarshal([]byte(ruleResult.Result), &ruleConfig); err != nil {
		return exploreResult, nil
	}
	if ruleConfig.UpdatePeriod == nil {
		return exploreResult, nil
	}
	updateTime, err := dateparse.ParseLocal(lastUpdateTimeInfo.BusinessUpdateTime)
	if err != nil {
		return exploreResult, nil
	}
	var latestTime time.Time
	switch *ruleConfig.UpdatePeriod {
	case "day":
		latestTime = time.Now().AddDate(0, 0, -1)
	case "week":
		latestTime = time.Now().AddDate(0, 0, -7)
	case "month":
		latestTime = time.Now().AddDate(0, -1, 0)
	case "quarter":
		latestTime = time.Now().AddDate(0, -3, 0)
	case "half_a_year":
		latestTime = time.Now().AddDate(0, -6, 0)
	case "year":
		latestTime = time.Now().AddDate(1, 0, 0)
	}
	if updateTime.After(latestTime) {
		score := float64(100)
		exploreResult.TimelinessScore = &score
	}
	return exploreResult, nil
}

func (d *DrivenImpl) GetExploreRule(ctx context.Context, req *driven.GetRuleListReq) ([]*driven.GetRuleResp, error) {
	urlStr := d.baseURL + "/api/internal/data-view/v1/explore-config/rule"
	args := struct {
		FormViewId string `query:"form_view_id"` // 视图id
		RuleLevel  string `query:"rule_level"`   // 规则级别，元数据级 字段级 行级 视图级
		Dimension  string `query:"dimension"`    // 维度，完整性 规范性 唯一性 准确性 一致性 及时性 数据统计
		FieldId    string `query:"field_id"`     // 字段id
		Keyword    string `query:"keyword"`      // 关键字查询
		Enable     bool   `query:"enable"`       // 启用状态，true为已启用，false为未启用，不传该参数则不跟据启用状态筛选
	}{
		FormViewId: req.FormViewId,
		RuleLevel:  req.RuleLevel,
		Dimension:  req.Dimension,
		FieldId:    req.FieldId,
		Keyword:    req.Keyword,
		Enable:     req.Enable,
	}
	resp, err := base.GET[[]*driven.GetRuleResp](ctx, d.httpClient, urlStr, args)
	return resp, err
}

func (d *DrivenImpl) BatchGetExploreReport(ctx context.Context, req *driven.BatchGetExploreReportReq) (*driven.BatchGetExploreReportResp, error) {
	errorMsg := "dataViewDriven BatchGetExploreReport "
	urlStr := d.baseURL + "/api/internal/data-view/v1/form-view/explore-report/batch"
	res := &driven.BatchGetExploreReportResp{}
	err := base.CallWithToken(ctx, d.httpClient, errorMsg, http.MethodPost, urlStr, req, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *DrivenImpl) GetExploreTaskList(ctx context.Context, workOrderId string) (*driven.ListExploreTaskResp, error) {
	urlStr := d.baseURL + "/api/internal/data-view/v1/explore-task"
	args := struct {
		WorkOrderId string `query:"work_order_id"`
	}{
		WorkOrderId: workOrderId,
	}
	resp, err := base.GET[*driven.ListExploreTaskResp](ctx, d.httpClient, urlStr, args)
	return resp, err
}
