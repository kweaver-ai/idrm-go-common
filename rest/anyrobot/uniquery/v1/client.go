package v1

import (
	"net/http"
	"net/url"
	"path"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type Client struct {
	// anyrobot uniquery v1 api endpoint, such as https://anyrobot.example.org/api/uniquery/v1
	base *url.URL
	// http client
	client *http.Client
}

// DataViews implements Interface.
func (c *Client) DataViews() DataViewInterface {
	base := *c.base
	base.Path = path.Join(c.base.Path, "data-views")
	return NewDataViewClient(&base, c.client)
}

// NewClient 创建客户端
func NewClient(base *url.URL, client *http.Client) *Client {
	return &Client{base: base, client: client}
}

// NewClientForBaseService 根据 rest/base.Service 创建客户端
func NewClientForBaseService(client *http.Client) (*Client, error) {
	// parse anyrobot host url
	u, err := url.Parse(base.Service.AnyRobotHost)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "api/uniquery/v1")

	return NewClient(u, client), nil
}

var _ Interface = &Client{}
