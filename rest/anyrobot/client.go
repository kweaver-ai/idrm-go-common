package anyrobot

import (
	"net/http"
	"net/url"
	"path"

	dataModelV1 "github.com/kweaver-ai/idrm-go-common/rest/anyrobot/data-model/v1"
	uniqueryV1 "github.com/kweaver-ai/idrm-go-common/rest/anyrobot/uniquery/v1"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

// AnyRobot client set
type ClientSet struct {
	dataModelV1 *dataModelV1.Client
	uniqueryV1  *uniqueryV1.Client
}

// New 返回 AnyRobot 客户端
func New(base *url.URL, client *http.Client) *ClientSet {
	return &ClientSet{
		dataModelV1: dataModelV1.NewClient(urlJoinPath(base, "api/data-model/v1"), client),
		uniqueryV1:  uniqueryV1.NewClient(urlJoinPath(base, "api/uniquery/v1"), client),
	}
}

// NewForConfig creates a new ClientSEet for the given config.
func NewForConfig(config *Config) (*ClientSet, error) {
	return NewForConfigAndClient(config, http.DefaultClient)
}

// NewForConfigAndClient creates a new ClientSet for the given config and http
// client.
func NewForConfigAndClient(config *Config, client *http.Client) (*ClientSet, error) {
	// parse anyrobot server address
	u, err := url.Parse(config.Server)
	if err != nil {
		return nil, err
	}
	return New(u, client), nil
}

// NewClientForBaseService 根据 rest/base.Service 创建客户端
func NewClientForBaseService(client *http.Client) (*ClientSet, error) {
	// parse anyrobot host url
	u, err := url.Parse(base.Service.AnyRobotHost)
	if err != nil {
		return nil, err
	}
	return New(u, client), nil
}

// DataModelV1 implements Interface.
func (cs *ClientSet) DataModelV1() dataModelV1.Interface {
	return cs.dataModelV1
}

// UniqueryV1 implements Interface.
func (cs *ClientSet) UniqueryV1() uniqueryV1.Interface {
	return cs.uniqueryV1
}

var _ Interface = &ClientSet{}

func urlJoinPath(in *url.URL, elem ...string) *url.URL {
	out := *in
	p := []string{in.Path}
	p = append(p, elem...)
	out.Path = path.Join(p...)
	return &out
}
