package impl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/basic_bigdata_service"
)

func (d *driven) GetLabelByIds(ctx context.Context, ids []string) (res *basic_bigdata_service.GetLabelByIdsRes, err error) {
	tmpIds := ""
	for i, id := range ids {
		if i == 0 {
			tmpIds = "id=" + id
		} else {
			tmpIds += "&id=" + id
		}
	}
	url := fmt.Sprintf("%s/api/basic-bigdata-service/v1/label/getByIds?%s", d.baseURL, tmpIds)
	res, err = base.Call[*basic_bigdata_service.GetLabelByIdsRes](ctx, d.httpClient, http.MethodGet, url, nil)
	return
}
