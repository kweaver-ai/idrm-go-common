package af_sailor_service

import (
	"context"
	"fmt"
)

type TableCompletionTableInfoReqBody struct {
	Id            string `json:"id"`
	TechnicalName string `json:"technical_name"`
	BusinessName  string `json:"business_name"`
	Desc          string `json:"desc"`
	Database      string `json:"database"`
	Subject       string `json:"subject"`
	Columns       []struct {
		Id            string `json:"id"`
		TechnicalName string `json:"technical_name"`
		BusinessName  string `json:"business_name"`
		DataType      string `json:"data_type"`
		Comment       string `json:"comment"`
	} `json:"columns"`
	RequestType           int    `json:"request_type"`
	ViewSourceCatalogName string `json:"view_source_catalog_name"`
}

type TableCompletionTableInfoResp struct {
	Res struct {
		TaskId string `json:"task_id"`
	} `json:"res"`
}

type TableCompletionReqBody struct {
	Id            string `json:"id"`
	TechnicalName string `json:"technical_name"`
	BusinessName  string `json:"business_name"`
	Desc          string `json:"desc"`
	Database      string `json:"database"`
	Subject       string `json:"subject"`
	Columns       []struct {
		Id            string `json:"id"`
		TechnicalName string `json:"technical_name"`
		BusinessName  string `json:"business_name"`
		DataType      string `json:"data_type"`
		Comment       string `json:"comment"`
	} `json:"columns"`
	RequestType           int      `json:"request_type"`
	GenFieldIds           []string `json:"gen_field_ids"`
	ViewSourceCatalogName string   `json:"view_source_catalog_name"`
}

type TableCompletionResp struct {
	Res struct {
		TaskId string `json:"task_id"`
	} `json:"res"`
}

// 补全表格信心、字段信息

func (c *client) TableCompletionAll(ctx context.Context, req *TableCompletionReqBody, authorization string) (*TableCompletionResp, error) {
	url := fmt.Sprintf("%s%s/understanding/table/completion", c.baseUrl, httpPathPrefix)

	headers := make(map[string]string)
	headers["Authorization"] = authorization
	return httpPostDo[TableCompletionResp](ctx, c.httpClient, url, req, headers)
}
