package v1

import (
	"net/http"
	"net/url"
	"path"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type DocAuditRestV1Interface interface {
	Base() *url.URL
	HTTPClient() *http.Client

	ApplysGetter
	DocAuditGetter
}

type DocAuditRestV1Client struct {
	base   *url.URL
	client *http.Client
}

var _ DocAuditRestV1Interface = &DocAuditRestV1Client{}

// Applys implements DocAuditRestV1Interface.
func (c *DocAuditRestV1Client) Applys() ApplyInterface {
	return newApplys(c)
}

// DocAudit implements DocAuditRestV1Interface.
func (c *DocAuditRestV1Client) DocAudit() DocAuditInterface {
	return newDocAudit(c)
}

func New(base *url.URL, client *http.Client) *DocAuditRestV1Client {
	return &DocAuditRestV1Client{
		base:   base,
		client: client,
	}
}

// NewForHTTPClient creates a DocAuditRestV1Interface based on given http client
// and GoCommon/rest/base configuration
func NewForHTTPClient(client *http.Client) (DocAuditRestV1Interface, error) {
	b, err := url.Parse(base.Service.DocAuditRestHost)
	if err != nil {
		return nil, err
	}
	b.Path = path.Join(b.Path, "/api/doc-audit-rest/v1")
	return New(b, client), nil
}

// Base implements DocAuditRestV1Interface.
func (c *DocAuditRestV1Client) Base() *url.URL {
	if c == nil {
		return nil
	}
	return c.base
}

// HTTPClient implements DocAuditRestV1Interface.
func (c *DocAuditRestV1Client) HTTPClient() *http.Client {
	if c == nil {
		return nil
	}
	return c.client
}
