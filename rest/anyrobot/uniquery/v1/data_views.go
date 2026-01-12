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
	"strconv"
	"strings"

	v1 "github.com/kweaver-ai/idrm-go-common/api/anyrobot/uniquery/v1"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
)

type DataViewsGetter interface {
	DataViews() DataViewInterface
}

type DataViewInterface interface {
	Get(ctx context.Context, ids []string, query *v1.DataViewQuery, opts *v1.DataViewQueryOptions) (*v1.ViewUniResponse, error)
}

type DataViewClient struct {
	// anyrobot data-view api endpoint, such as https://anyrobot.example.org/api/uniquery/v1/data-views
	base *url.URL
	// http client
	client *http.Client
}

func NewDataViewClient(base *url.URL, client *http.Client) *DataViewClient {
	return &DataViewClient{base: base, client: client}
}

// Get implements DataViewInterface.
//
// TODO: 识别 AnyRobot 返回的结构化错误
func (c *DataViewClient) Get(ctx context.Context, ids []string, query *v1.DataViewQuery, opts *v1.DataViewQueryOptions) (result *v1.ViewUniResponse, err error) {
	base := *c.base
	base.Path = path.Join(base.Path, strings.Join(ids, ","))
	// query
	base.RawQuery = newQueryForOptions(opts).Encode()

	var body []byte

	if body, err = json.Marshal(query); err != nil {
		return nil, fmt.Errorf("marshal query fail: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base.String(), bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create http request fail: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Http-Method-Override", "GET")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errorcode.NewDoRequestError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.NewAnyRobotUniqueryServerError(resp)
	}

	if body, err = io.ReadAll(resp.Body); err != nil {
		return nil, fmt.Errorf("read http response body fail: %w", err)
	}

	result = &v1.ViewUniResponse{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, fmt.Errorf("unmarshal response body fail: %w", err)
	}

	return result, nil
}

var _ DataViewInterface = &DataViewClient{}

// 生成 v1.DataViewQueryOptions 对应的 query 参数
func newQueryForOptions(opts *v1.DataViewQueryOptions) url.Values {
	if opts == nil {
		return nil
	}

	q := make(url.Values)

	if opts.AllowNonExistField {
		q.Set("allow_non_exist_field", strconv.FormatBool(opts.AllowNonExistField))
	}

	if len(q) == 0 {
		return nil
	}
	return q
}
