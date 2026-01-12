package impl

import (
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/data_sync"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

// NewDataSyncDriven    数据同步
func NewDataSyncDriven(httpClient *http.Client) driven.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.DataSyncHost,
		httpClient: httpClient,
	}
}
