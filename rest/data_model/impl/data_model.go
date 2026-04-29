package impl

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_model"
)

type drivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewDrivenImpl(httpClient *http.Client) driven.Driven {
	return &drivenImpl{
		baseURL:    base.Service.MDLDataModelServiceHost,
		httpClient: httpClient,
	}
}

func (d *drivenImpl) GetDataModelByID(ctx context.Context, ids ...string) ([]*driven.DataModel, error) {
	path := fmt.Sprintf("/api/mdl-data-model/v1/data-views/%s", strings.Join(ids, ","))
	url := d.baseURL + path
	resp, err := base.GET[[]*driven.DataModel](ctx, d.httpClient, url, ids)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *drivenImpl) GetDataModelByIDInternal(ctx context.Context, ids ...string) ([]*driven.DataModel, error) {
	path := fmt.Sprintf("/api/mdl-data-model/in/v1/data-views/%s", strings.Join(ids, ","))
	url := d.baseURL + path
	resp, err := base.GET[[]*driven.DataModel](ctx, d.httpClient, url, ids)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
