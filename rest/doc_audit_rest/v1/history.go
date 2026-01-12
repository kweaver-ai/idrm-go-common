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

type HistoriesGetter interface {
	Histories() HistoryInterface
}

type HistoryInterface interface {
	List(ctx context.Context, opts *v1.HistoryListOptions) (*v1.HistoryList, error)
}

type history struct {
	base   *url.URL
	client *http.Client
}

func newHistory(c *docAudit) *history {
	b := *c.base
	b.Path = path.Join(b.Path, "historys")
	return &history{base: &b, client: c.client}
}

// List implements HistoryInterface.
func (c *history) List(ctx context.Context, opts *v1.HistoryListOptions) (*v1.HistoryList, error) {
	// create http request
	b := *c.base
	if opts != nil {
		b.RawQuery = queryForHistoryListOptions(opts).Encode()
	}
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
		return nil, fmt.Errorf("list histories failed: statusCode=%d, body=%s", resp.StatusCode, body)
	}
	result := &v1.HistoryList{}
	if err := json.Unmarshal(body, result); err != nil {
		// TODO: Define errorcode
		return nil, err
	}

	return result, nil
}
