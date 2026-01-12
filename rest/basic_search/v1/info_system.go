package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	basic_search_v1 "github.com/kweaver-ai/idrm-go-common/api/basic_search/v1"
	"github.com/kweaver-ai/idrm-go-common/interception"
)

//go:generate mockgen -destination=mock/info_system.go -package=mock -typed -write_generate_directive -source info_system.go

type InfoSystemsGetter interface {
	InfoSystems() InfoSystemInterface
}

type InfoSystemInterface interface {
	// Search 搜索信息系统
	Search(ctx context.Context, query *basic_search_v1.InfoSystemSearchQuery, opts *basic_search_v1.InfoSystemSearchOptions) (*basic_search_v1.InfoSystemSearchResult, error)
}

type InfoSystemClient struct {
	Base   *url.URL
	Client *http.Client
}

func NewInfoSystemClient(base *url.URL, client *http.Client) *InfoSystemClient {
	return &InfoSystemClient{Base: base, Client: client}
}

// Search 搜索信息系统
func (c *InfoSystemClient) Search(ctx context.Context, query *basic_search_v1.InfoSystemSearchQuery, opts *basic_search_v1.InfoSystemSearchOptions) (*basic_search_v1.InfoSystemSearchResult, error) {
	var base url.URL = *c.Base
	// Path
	base.Path = path.Join(base.Path, "search")
	// Query
	q, err := opts.MarshalQuery()
	if err != nil {
		return nil, err
	}
	base.RawQuery = q.Encode()

	// Create Body
	reqBody, err := json.Marshal(&basic_search_v1.InfoSystemSearch{Query: query})
	if err != nil {
		// TODO: 返回结构化错误
		return nil, err
	}

	// Create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), bytes.NewReader(reqBody))
	if err != nil {
		// TODO: 返回结构化错误
		return nil, err
	}

	// Authorization
	if token, err := interception.BearerTokenFromContextCompatible(ctx); err == nil {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取 HTTP Response Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO: 返回结构化错误
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		// TODO: 返回结构化错误
		return nil, fmt.Errorf("search info system fail, %s: %s", resp.Status, body)
	}

	var result basic_search_v1.InfoSystemSearchResult
	if err := json.Unmarshal(body, &result); err != nil {
		// TODO: 返回结构化错误
		return nil, err
	}

	return &result, nil
}

var _ InfoSystemInterface = &InfoSystemClient{}
