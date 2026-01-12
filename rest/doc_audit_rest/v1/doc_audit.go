package v1

import (
	"net/http"
	"net/url"
	"path"
)

type DocAuditGetter interface {
	DocAudit() DocAuditInterface
}

type DocAuditInterface interface {
	BizGetter
	HistoriesGetter
	TasksGetter
}

type docAudit struct {
	base   *url.URL
	client *http.Client
}

func newDocAudit(c *DocAuditRestV1Client) *docAudit {
	b := *c.base
	b.Path = path.Join(b.Path, "doc-audit")
	return &docAudit{base: &b, client: c.client}
}

// Biz implements DocAuditInterface.
func (c *docAudit) Biz() BizInterface {
	return newBiz(c)
}

// Histories implements DocAuditInterface.
func (c *docAudit) Histories() HistoryInterface {
	return newHistory(c)
}

// Tasks implements DocAuditInterface.
func (c *docAudit) Tasks() TaskInterface {
	return newTask(c)
}

var _ DocAuditInterface = &docAudit{}
