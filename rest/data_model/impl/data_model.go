package impl

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_model"
)

const pathInDataViewRowColumnRules = "/api/mdl-data-model/v1/data-view-row-column-rules"

type drivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewDrivenImpl(httpClient *http.Client) driven.Driven {
	return &drivenImpl{
		baseURL:    base.Service.MDLDataModelServiceHost,
		httpClient: httpClient,
	}
}

func (d *drivenImpl) GetDataModelByID(ctx context.Context, ids ...string) ([]*driven.DataModel, error) {
	path := fmt.Sprintf("/api/mdl-data-model/v1/data-views/%s", strings.Join(ids, ","))
	url := d.baseURL + path
	resp, err := base.GET[[]*driven.DataModel](ctx, d.httpClient, url, ids)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *drivenImpl) GetDataModelByIDInternal(ctx context.Context, ids ...string) ([]*driven.DataModel, error) {
	path := fmt.Sprintf("/api/mdl-data-model/in/v1/data-views/%s", strings.Join(ids, ","))
	url := d.baseURL + path
	resp, err := base.GET[[]*driven.DataModel](ctx, d.httpClient, url, ids)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *drivenImpl) ListDataViewRowColumnRulesInternal(ctx context.Context, q driven.ListDataViewRowColumnRulesQuery) (*driven.ListDataViewRowColumnRulesResult, error) {
	url := d.baseURL + pathInDataViewRowColumnRules
	resp, err := base.GET[*driven.ListDataViewRowColumnRulesResult](ctx, d.httpClient, url, q)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type getDataViewRowColumnRulesPath struct {
	RuleIDs string `uri:"rule_ids"`
}

func (d *drivenImpl) GetDataViewRowColumnRulesInternal(ctx context.Context, ruleIDs []string) ([]*driven.DataViewRowColumnRule, error) {
	if len(ruleIDs) == 0 {
		return nil, errorcode.Detail(errorcode.BadRequestError, "ruleIDs 不能为空")
	}
	url := d.baseURL + pathInDataViewRowColumnRules + "/:rule_ids"
	resp, err := base.GET[[]*driven.DataViewRowColumnRule](ctx, d.httpClient, url, getDataViewRowColumnRulesPath{
		RuleIDs: strings.Join(ruleIDs, ","),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type createDataViewRowColumnRuleRef struct {
	ID string `json:"id"`
}

func (d *drivenImpl) CreateDataViewRowColumnRulesInternal(ctx context.Context, rules []driven.DataViewRowColumnRuleWrite) ([]string, error) {
	if len(rules) == 0 {
		return nil, errorcode.Detail(errorcode.BadRequestError, "rules 不能为空")
	}
	url := d.baseURL + pathInDataViewRowColumnRules
	refs, err := base.POST[[]createDataViewRowColumnRuleRef](ctx, d.httpClient, url, rules)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(refs))
	for _, r := range refs {
		out = append(out, r.ID)
	}
	return out, nil
}

type updateDataViewRowColumnRuleArgs struct {
	RuleID     string                   `uri:"rule_id"`
	Name       string                   `json:"name"`
	ViewID     string                   `json:"view_id"`
	Tags       []string                 `json:"tags,omitempty"`
	Comment    string                   `json:"comment,omitempty"`
	Fields     []string                 `json:"fields,omitempty"`
	RowFilters *driven.RowColumnCondCfg `json:"row_filters,omitempty"`
}

func (d *drivenImpl) UpdateDataViewRowColumnRuleInternal(ctx context.Context, ruleID string, rule *driven.DataViewRowColumnRuleWrite) error {
	if strings.TrimSpace(ruleID) == "" {
		return errorcode.Detail(errorcode.BadRequestError, "ruleID 不能为空")
	}
	if rule == nil {
		return errorcode.Detail(errorcode.BadRequestError, "rule 不能为空")
	}
	url := d.baseURL + pathInDataViewRowColumnRules + "/:rule_id"
	args := updateDataViewRowColumnRuleArgs{
		RuleID:     ruleID,
		Name:       rule.Name,
		ViewID:     rule.ViewID,
		Tags:       rule.Tags,
		Comment:    rule.Comment,
		Fields:     rule.Fields,
		RowFilters: rule.RowFilters,
	}
	_, err := base.PUT[struct{}](ctx, d.httpClient, url, args)
	return err
}

type deleteDataViewRowColumnRulesPath struct {
	RuleIDs string `uri:"rule_ids"`
}

func (d *drivenImpl) DeleteDataViewRowColumnRulesInternal(ctx context.Context, ruleIDs []string) error {
	if len(ruleIDs) == 0 {
		return errorcode.Detail(errorcode.BadRequestError, "ruleIDs 不能为空")
	}
	url := d.baseURL + pathInDataViewRowColumnRules + "/:rule_ids"
	_, err := base.DELETE[struct{}](ctx, d.httpClient, url, deleteDataViewRowColumnRulesPath{
		RuleIDs: strings.Join(ruleIDs, ","),
	})
	return err
}
