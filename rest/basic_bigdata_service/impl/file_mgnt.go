package impl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/basic_bigdata_service"
)

type driven struct {
	baseURL    string
	httpClient *http.Client
}

func NewDriven(client *http.Client) basic_bigdata_service.Driven {
	return &driven{
		baseURL:    base.Service.BasicBigDataHost,
		httpClient: client,
	}
}

// 上传文件
func (d *driven) UploadFile(ctx context.Context, req *basic_bigdata_service.UploadReq) (res *basic_bigdata_service.UploadResp, err error) {
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/upload", d.baseURL)
	res, err = base.Call[*basic_bigdata_service.UploadResp](ctx, d.httpClient, http.MethodPost, url, req)
	return
}

// 删除文件
func (d *driven) DeleteFile(ctx context.Context, req *basic_bigdata_service.DeleteReq) (err error) {
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/%s", d.baseURL, req.ID)
	_, err = base.Call[any](ctx, d.httpClient, http.MethodDelete, url, nil)
	return
}

// 根据ossId 删除文件
func (d *driven) DeleteOssFile(ctx context.Context, req *basic_bigdata_service.DeleteReq) (err error) {
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/fileByOssId/%s", d.baseURL, req.ID)
	_, err = base.Call[any](ctx, d.httpClient, http.MethodDelete, url, nil)
	return
}

// 根据关联对象批量删除文件
func (d *driven) DeleteFiles(ctx context.Context, req *basic_bigdata_service.DeleteReq) (err error) {
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/batch/%s", d.baseURL, req.ID)
	_, err = base.Call[any](ctx, d.httpClient, http.MethodDelete, url, nil)
	return
}

// 分页查询文件列表
func (d *driven) QueryPageFile(ctx context.Context, req *basic_bigdata_service.QueryPageReq) (res *basic_bigdata_service.QueryResp, err error) {
	str := ""
	for _, id := range req.IDs {
		str += "&id=" + id
	}
	for _, relatedObjectId := range req.RelatedObjectIDs {
		str += "&related_object_id=" + relatedObjectId
	}
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/page?offset=%d&limit=%d&keyword=%s&type=%s%s",
		d.baseURL, req.Offset, req.Limit, req.Keyword, req.Type, str)
	res, err = base.Call[*basic_bigdata_service.QueryResp](ctx, d.httpClient, http.MethodGet, url, req)
	return
}

// 下载文件
func (d *driven) DownloadFile(ctx context.Context, req *basic_bigdata_service.DownloadReq) (res *basic_bigdata_service.DownloadResp, err error) {
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/down?id=%s&oss_id=%s", d.baseURL, req.ID, req.OssID)
	res, err = base.Call[*basic_bigdata_service.DownloadResp](ctx, d.httpClient, http.MethodGet, url, req)
	return
}

func (d *driven) GetCountsByIds(ctx context.Context, req *basic_bigdata_service.GetCountsByIdsReq) (res *basic_bigdata_service.GetCountsByIdsResp, err error) {
	idStr := ""
	for _, id := range req.IDs {
		idStr += "&ids=" + id
	}
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/countsByIds?type=%s%s", d.baseURL, req.Type, idStr)
	res, err = base.Call[*basic_bigdata_service.GetCountsByIdsResp](ctx, d.httpClient, http.MethodGet, url, req)
	return
}

// 根据关联对象id批量删除不包括这些文件ID的文件
func (d *driven) DeleteExcludeFiles(ctx context.Context, req *basic_bigdata_service.DeleteExcludeOssIdsReq) (err error) {
	if len(req.ExcludeOssIDs) <= 0 {
		return
	}
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/batchExcludeOssIds?related_object_id=%s&exclude_oss_ids=%s", d.baseURL, req.RelatedObjectId, req.ExcludeOssIDs)
	_, err = base.Call[any](ctx, d.httpClient, http.MethodDelete, url, nil)
	return
}

// 下载文件
func (d *driven) PreviewPdfHref(ctx context.Context, req *basic_bigdata_service.PreviewPdfReq) (res *basic_bigdata_service.PreviewPdfResp, err error) {
	url := fmt.Sprintf("%s/api/internal/basic-bigdata-service/v1/file-management/preview-pdf-href?id=%s&preview_id=%s", d.baseURL, req.ID, req.PreviewID)
	res, err = base.Call[*basic_bigdata_service.PreviewPdfResp](ctx, d.httpClient, http.MethodGet, url, req)
	return
}
