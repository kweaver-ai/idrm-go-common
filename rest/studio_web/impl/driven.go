package impl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/studio_web"
)

type drivenImpl struct {
	baseURL string
	client  *http.Client
}

func NewDriven(client *http.Client) driven.Driven {
	return &drivenImpl{
		baseURL: base.Service.StudioWebServiceHost,
		client:  client,
	}
}

func (d drivenImpl) InsertWebApps(ctx context.Context, req []*driven.WebAppConfig) error {
	urlStr := fmt.Sprintf("%s/api/workstation/v1/webapp", d.baseURL)
	_, err := base.PUT[any](ctx, d.client, urlStr, req)
	return err
}
