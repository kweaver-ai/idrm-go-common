package impl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/workflow"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"go.uber.org/zap"
)

type WorkflowDriven struct {
	workflowRestURL string
	docAuditRestURL string
	client          *http.Client
}

func NewWorkflowDriven(client *http.Client) workflow.WorkflowDriven {
	return &WorkflowDriven{
		client:          client,
		workflowRestURL: base.Service.WorkFlowRestHost,
		docAuditRestURL: base.Service.DocAuditRestHost,
	}
}

func (c *WorkflowDriven) GetAuditProcessDefinition(ctx context.Context, procDefKey string) (res *workflow.ProcessDefinitionGetRes, err error) {
	errorMsg := "WorkFlowDriven GetAuditProcessDefinition "

	urlStr := fmt.Sprintf("%s/api/workflow-rest/v1/process-definition/%s", c.workflowRestURL, procDefKey)

	log.Infof(errorMsg+" url:%s \n", urlStr)

	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodGet, urlStr, c.client, nil)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetAuditProcessDefinitionError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetAuditProcessDefinitionError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}

	res = &workflow.ProcessDefinitionGetRes{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetAuditProcessDefinitionError, err.Error())
	}
	return res, nil
}

func (c *WorkflowDriven) GetAuditList(ctx context.Context, listType workflow.WorkflowListType,
	auditTypes []string, offset, limit int) (*workflow.AuditResponse, error) {
	errorMsg := "WorkFlowDriven GetAuditList "
	values := url.Values{
		"type":   auditTypes,
		"offset": []string{fmt.Sprint((offset - 1) * limit)},
		"limit":  []string{fmt.Sprint(limit)},
	}

	urlStr := fmt.Sprintf("%s/api/doc-audit-rest/v1/doc-audit/%s?%s", c.docAuditRestURL, listType, values.Encode())

	log.Infof(errorMsg+" url:%s \n", urlStr)

	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodGet, urlStr, c.client, nil)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetAuditListError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetAuditListError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}

	res := &workflow.AuditResponse{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetAuditListError, err.Error())
	}
	return res, nil
}

func (c *WorkflowDriven) GetAuditLogsByProcInstID(ctx context.Context, procInstID string) ([]*workflow.AuditNodeLog, error) {
	errorMsg := "WorkFlowDriven GetAuditLogsByProcInstID "
	urlStr := fmt.Sprintf("%s/api/workflow-rest/v1/process-instance/%s/logs", c.workflowRestURL, procInstID)

	log.Infof(errorMsg+" url:%s \n", urlStr)

	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodGet, urlStr, c.client, nil)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetAuditLogsError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetAuditLogsError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}

	var res []*workflow.AuditNodeLog
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetAuditLogsError, err.Error())
	}
	return res, nil
}
