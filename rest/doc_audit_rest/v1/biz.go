package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	v1 "github.com/kweaver-ai/idrm-go-common/api/doc_audit_rest/v1"
	"github.com/kweaver-ai/idrm-go-common/interception"
)

type BizGetter interface {
	Biz() BizInterface
}

type BizInterface interface {
	Get(ctx context.Context, id string) (*v1.Apply, error)
}

type biz struct {
	base   *url.URL
	client *http.Client
}

func newBiz(c *docAudit) *biz {
	b := *c.base
	b.Path = path.Join(b.Path, "biz")
	return &biz{base: &b, client: c.client}
}

// Get implements BizInterface.
func (c *biz) Get(ctx context.Context, id string) (*v1.Apply, error) {
	// create http request
	b := *c.base
	b.Path = path.Join(b.Path, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.String(), http.NoBody)
	if err != nil {
		// TODO: Define errorcode
		return nil, err
	}
	if t, err := interception.BearerTokenFromContextCompatible(ctx); err == nil {
		req.Header.Add("Authorization", "Bearer "+t)
	}

	// send http request
	resp, err := c.client.Do(req)
	if err != nil {
		// TODO: Define errorcode
		return nil, err
	}
	defer resp.Body.Close()

	// process http response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO: Define errorcode
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		// TODO: Define errorcode
		return nil, fmt.Errorf("get applys/%s failed: statusCode=%d, body=%s", id, resp.StatusCode, body)
	}
	result := &v1.Apply{}
	if err := json.Unmarshal(body, result); err != nil {
		// TODO: Define errorcode
		return nil, err
	}

	return result, nil
}

var _ BizInterface = &biz{}
