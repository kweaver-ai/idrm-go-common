package v1

import (
	"net/url"
	"strconv"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// 指标维度规则
type IndicatorDimensionalRule struct {
	// 元数据
	meta_v1.Metadata `json:"metadata,omitempty"`
	// 指标维度规则的内容
	Spec IndicatorDimensionalRuleSpec `json:"spec,omitempty"`
}

// 指标维度规则，复用 Rule
type IndicatorDimensionalRuleSpec struct {
	//上一级的范围ID
	AuthScopeID int `json:"auth_scope_id,omitempty,string"`
	// 维度规则所属指标的 ID
	IndicatorID int `json:"indicator_id,omitempty,string"`
	//当前用户是否有该指标维度的授权权限
	CanAuth bool `json:"can_auth"`
	// 规则
	Rule
}

// 指标维度规则列表
type IndicatorDimensionalRuleList meta_v1.List[IndicatorDimensionalRule]

// 获取指标维度规则列表的选项
type IndicatorDimensionalRuleListOptions struct {
	meta_v1.ListOptions
	// 未指定且非空时返回属于指定指标的维度规则
	IndicatorID int `json:"indicator_id,omitempty"`
	// 排序的依据,为空时不排序
	Sort IndicatorDimensionalRuleSort `form:"sort" json:"sort,omitempty"`
	// 排序的方向
	Direction meta_v1.Direction `form:"direction,default=asc" json:"direction,omitempty"`
}

func (opts *IndicatorDimensionalRuleListOptions) MarshalQuery() (url.Values, error) {
	q, err := opts.ListOptions.MarshalQuery()
	if err != nil {
		return nil, err
	}

	if opts.IndicatorID != 0 {
		q.Set("indicator_id", strconv.Itoa(opts.IndicatorID))
	}
	if opts.Sort != "" {
		q.Set("sort", string(opts.Sort))
	}
	if opts.Direction != "" {
		q.Set("direction", string(opts.Direction))
	}

	return q, nil
}

func (opts *IndicatorDimensionalRuleListOptions) UnmarshalQuery(data url.Values) (err error) {
	if err = opts.ListOptions.UnmarshalQuery(data); err != nil {
		return
	}
	for k, values := range data {
		for _, v := range values {
			switch k {
			// 指标维度规则所属的指标 ID
			case "indicator_id":
				if opts.IndicatorID, err = strconv.Atoi(v); err != nil {
					return
				}
			// 指标维度规则的排序依据
			case "sort":
				opts.Sort = IndicatorDimensionalRuleSort(v)
			// 排序的方向
			case "direction":
				opts.Direction = meta_v1.Direction(v)
			default:
				continue
			}
		}
	}
	return
}

// IndicatorDimensionalRuleSort 定义指标维度规则的排序依据
type IndicatorDimensionalRuleSort string

const (
	// 根据当前用户指标维度规则是否拥有权限排序。升序：无权限在前，有权限
	// 在后。降序：有权限在前，无权限在后。拥有 read 权限即为有权
	// 限
	IndicatorDimensionalRuleSortIsAuthorized IndicatorDimensionalRuleSort = "is_authorized"
)

// IndicatorDimensionalRuleListArgs  获取指标维度规则列表的选项, 逗号拼接，支持多个
type IndicatorDimensionalRuleListArgs struct {
	IndicatorRuleID string `json:"indicator_rule_id" form:"indicator_rule_id" binding:"required"`
}

// IndicatorDimensionalByIndicatorRulesReq  获取指标维度规则列表的选项,指标ID, 逗号拼接，支持多个
type IndicatorDimensionalByIndicatorRulesReq struct {
	IndicatorID string `json:"indicator_id" form:"indicator_id" binding:"required"`
}
