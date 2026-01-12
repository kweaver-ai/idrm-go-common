package impl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/workflow"
)

type DocAuditDriven struct {
	baseURL string
	client  *http.Client
}

func NewDocAuditDriven(client *http.Client) workflow.DocAuditDriven {
	return &DocAuditDriven{
		client:  client,
		baseURL: base.Service.DocAuditRestHost,
	}
}

func (c *DocAuditDriven) GetMyTodoList(ctx context.Context, req *workflow.GetMyTodoListReq) (res *workflow.GetMyTodoListRes, err error) {
	fullURL := fmt.Sprintf("%s/api/doc-audit-rest/v1/doc-audit/tasks?doc_name=%s&type=%s&abstracts=%s&limit=%d&offset=%d",
		c.baseURL,
		url.QueryEscape(req.DocName),
		url.QueryEscape(strings.Join(req.Type, ",")),
		url.QueryEscape(req.Abstracts),
		req.Limit,
		req.Offset,
	)
	res, err = base.Call[*workflow.GetMyTodoListRes](ctx, c.client, http.MethodGet, fullURL, struct{}{})
	return
}
