package v1

import (
	"context"
	"net/http"
	"net/url"
	"path"

	v1 "github.com/kweaver-ai/idrm-go-common/api/doc_audit_rest/v1"
)

type ApplysGetter interface {
	Applys() ApplyInterface
}

type ApplyInterface interface {
	List(ctx context.Context, opts *v1.ApplyListOptions) (*v1.ApplyList, error)
}

type applys struct {
	base   *url.URL
	client *http.Client
}

var _ ApplyInterface = &applys{}

func newApplys(c *DocAuditRestV1Client) *applys {
	b := *c.Base()
	b.Path = path.Join(b.Path, "applys")
	return &applys{
		base:   &b,
		client: c.HTTPClient(),
	}
}

// List implements ApplyInterface.
func (c *applys) List(ctx context.Context, opts *v1.ApplyListOptions) (*v1.ApplyList, error) {
	panic("unimplemented")
}
