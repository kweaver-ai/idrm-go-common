package impl

import (
	"net/http"

	driven "github.com/kweaver-ai/idrm-go-common/rest/authorization"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
)

type drivenImpl struct {
	privateBaseURL string
	publicBaseURL  string
	httpClient     *http.Client
}

func NewDriven(client *http.Client) driven.Driven {
	return &drivenImpl{
		privateBaseURL: base.Service.AuthorizationPrivateHost,
		publicBaseURL:  base.Service.AuthorizationPublicHost,
		httpClient:     client,
	}
}

func (d drivenImpl) privatePath(uri string) string {
	return d.privateBaseURL + uri
}

func (d drivenImpl) publicPath(uri string) string {
	return d.publicBaseURL + uri
}
