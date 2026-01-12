package v1

import (
	"net/http"
	"net/url"
	"path"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type Client struct {
	Base   *url.URL
	Client *http.Client
}

// InfoSystems implements Interface.
func (c *Client) InfoSystems() InfoSystemInterface {
	base := *c.Base
	// go1.19 才有 url.JoinPath，升级到 go1.19 之前先这么写
	base.Path = path.Join(base.Path, "info-systems")
	return NewInfoSystemClient(&base, c.Client)
}

func New(base *url.URL, client *http.Client) *Client {
	return &Client{Base: base, Client: client}
}

func NewForServiceConfigAndClient(config *base.ServiceConfig, client *http.Client) (*Client, error) {
	base, err := url.Parse(config.BasicSearchHost)
	if err != nil {
		return nil, err
	}

	// add api path
	//
	// go1.19 才有 url.JoinPath，升级到 go1.19 之前先这么写
	base.Path = path.Join(base.Path, "api", "basic-search", "v1")

	return New(base, client), nil
}

// 根据 GoCommon/rest/base 的配置和指定的 HTTP 客户端创建 Client
func NewForBaseAndClient(client *http.Client) (*Client, error) {
	return NewForServiceConfigAndClient(base.Service, client)
}

var _ Interface = &Client{}
