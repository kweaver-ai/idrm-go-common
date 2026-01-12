package data_subject_impl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/data_subject"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"go.uber.org/zap"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewDataViewDriven(httpClient *http.Client) data_subject.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.DataSubjectHost,
		httpClient: httpClient,
	}
}

func (d *DrivenImpl) GetDataSubjectByPath(ctx context.Context, paths *data_subject.GetDataSubjectByPathReq) (res *data_subject.GetDataSubjectByPathRes, err error) {
	errorMsg := "DataViewDriven GetDataSubjectByPath "

	urlStr := fmt.Sprintf("%s/api/internal/data-subject/v1/subject-domain/paths", d.baseURL)

	log.Infof(errorMsg+" url:%s \n req : %v", urlStr, paths)

	statusCode, body, err := base.DOWithToken(ctx, errorMsg, http.MethodPost, urlStr, d.httpClient, paths)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetDataSubjectByPathError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDataSubjectByPathError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}

	res = &data_subject.GetDataSubjectByPathRes{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDataSubjectByPathError, err.Error())
	}
	return res, nil
}

func (d *DrivenImpl) GetDataSubjectByID(ctx context.Context, ids []string) (res *data_subject.DataSubjectListObject, err error) {
	var u *url.URL
	errorMsg := "DataViewDriven GetDataSubjectByID "

	urlStr := fmt.Sprintf("%s/api/internal/data-subject/v1/subject-domain/precision", d.baseURL)

	u, err = url.Parse(urlStr)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetDataSubjectByIDError, err.Error())
	}
	u.RawQuery = url.Values{
		"object_id": ids,
	}.Encode()
	log.Infof(errorMsg+" url:%s \n req : %v", u.String())
	statusCode, body, err := base.DOWithOutToken(ctx, errorMsg, http.MethodGet, u.String(), d.httpClient, nil)
	if err != nil {
		return nil, errorcode.Detail(errorcode.GetDataSubjectByIDError, err.Error())
	}

	if statusCode != http.StatusOK {
		return nil, errorcode.Detail(errorcode.GetDataSubjectByIDError, base.StatusCodeNotOK(errorMsg, statusCode, body).Error())
	}

	res = &data_subject.DataSubjectListObject{}
	if err = jsoniter.Unmarshal(body, &res); err != nil {
		log.Error(errorMsg+" json.Unmarshal error", zap.Error(err))
		return nil, errorcode.Detail(errorcode.GetDataSubjectByIDError, err.Error())
	}
	return res, nil
}

func (d *DrivenImpl) GetSubjectList(ctx context.Context, parentId, subjectType string) (*data_subject.DataSubjectListRes, error) {
	var res = &data_subject.DataSubjectListRes{}

	// 至少获取一次，如果已经获取到的主题域数量小于总数，则继续获取
	for offset := 1; offset == 1 || len(res.Entries) < res.TotalCount; offset++ {
		pageRes, err := d.GetSubjectListWithOffset(ctx, parentId, subjectType, offset)
		if err != nil {
			return nil, err
		}
		res.Entries, res.TotalCount = append(res.Entries, pageRes.Entries...), pageRes.TotalCount
	}

	return res, nil
}

func (d *DrivenImpl) GetSubjectListWithOffset(ctx context.Context, parentId, subjectType string, offset int) (*data_subject.DataSubjectListRes, error) {
	errorMsg := "DrivenDataSubject GetSubjectList "
	urlStr := fmt.Sprintf("%s/api/data-subject/v1/subject-domains?parent_id=%s&type=%s&is_all=true&need_count=true&offset=%d", d.baseURL, parentId, subjectType, offset)
	log.Infof(errorMsg+"%s", urlStr)

	res := &data_subject.DataSubjectListRes{}

	err := base.CallWithToken(ctx, d.httpClient, errorMsg, http.MethodGet, urlStr, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
