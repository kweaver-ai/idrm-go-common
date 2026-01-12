package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	v1 "github.com/kweaver-ai/idrm-go-common/api/anyrobot/data-model/v1"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
)

type DataViewsGetter interface {
	DataViews() DataViewInterface
}

type DataViewInterface interface {
	Get(ctx context.Context, ids []string) ([]v1.DataView, error)
}

type DataViewClient struct {
	// anyrobot data-view api endpoint, such as https://anyrobot.example.org/api/data-model/v1/data-views
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
func (c *DataViewClient) Get(ctx context.Context, ids []string) ([]v1.DataView, error) {
	base := *c.base
	base.Path = path.Join(base.Path, strings.Join(ids, ","))

	var body []byte

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create http request fail: %w", err)
	}
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

	var result []v1.DataView
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response body fail: %w", err)
	}

	return result, nil
}

var _ DataViewInterface = &DataViewClient{}
