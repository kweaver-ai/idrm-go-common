package impl

import (
	"context"
	"net/http"
	"strings"

	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/dip_studio"
)

type drivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

type digitalHumanPathArgs struct {
	ID string `uri:"id"`
}

type updateDigitalHumanBknArgs struct {
	ID  string            `uri:"id"`
	Bkn []driven.BknEntry `json:"bkn"`
}

func NewDriven(httpClient *http.Client) driven.Driven {
	return &drivenImpl{
		baseURL:    base.Service.DipStudioHost,
		httpClient: httpClient,
	}
}

// ListDigitalHumans GET /api/dip-studio/v1/digital-human
func (d *drivenImpl) ListDigitalHumans(ctx context.Context) ([]*driven.DigitalHumanDetail, error) {
	path := d.baseURL + "/api/dip-studio/v1/digital-human"
	resp, err := base.GET[[]*driven.DigitalHumanDetail](ctx, d.httpClient, path, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetDigitalHumanDetail GET /api/dip-studio/v1/digital-human/{id}
func (d *drivenImpl) GetDigitalHumanDetail(ctx context.Context, digitalHumanID string) (*driven.DigitalHumanDetail, error) {
	if strings.TrimSpace(digitalHumanID) == "" {
		return nil, errorcode.Detail(errorcode.BadRequestError, "digitalHumanID 不能为空")
	}
	path := d.baseURL + "/api/dip-studio/v1/digital-human/:id"
	detail, err := base.GET[*driven.DigitalHumanDetail](ctx, d.httpClient, path, digitalHumanPathArgs{ID: digitalHumanID})
	if err != nil {
		return nil, err
	}
	if detail == nil {
		return nil, errorcode.Detail(errorcode.BadRequestError, "获取数字员工详情为空响应")
	}
	return detail, nil
}

func bknDedupeKey(e driven.BknEntry) string {
	u := strings.TrimSpace(e.URL)
	n := strings.TrimSpace(e.Name)
	if u != "" {
		return "u:" + u
	}
	if n != "" {
		return "n:" + n
	}
	return ""
}

func mergeBkn(existing, add []driven.BknEntry) []driven.BknEntry {
	seen := make(map[string]struct{}, len(existing)+len(add))
	out := make([]driven.BknEntry, 0, len(existing)+len(add))
	for _, e := range existing {
		k := bknDedupeKey(e)
		if k == "" {
			continue
		}
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		out = append(out, e)
	}
	for _, e := range add {
		k := bknDedupeKey(e)
		if k == "" {
			continue
		}
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		out = append(out, e)
	}
	return out
}

// AddDigitalHumanKnowledgeNetwork GET /api/dip-studio/v1/digital-human/{id} 后 PUT 合并后的 bkn
func (d *drivenImpl) AddDigitalHumanKnowledgeNetwork(ctx context.Context, digitalHumanID string, entries []driven.BknEntry) (*driven.DigitalHumanDetail, error) {
	if strings.TrimSpace(digitalHumanID) == "" {
		return nil, errorcode.Detail(errorcode.BadRequestError, "digitalHumanID 不能为空")
	}
	if len(entries) == 0 {
		return nil, errorcode.Detail(errorcode.BadRequestError, "entries 不能为空")
	}
	for _, e := range entries {
		if strings.TrimSpace(e.Name) == "" || strings.TrimSpace(e.URL) == "" {
			return nil, errorcode.Detail(errorcode.BadRequestError, "每条 BknEntry 的 name 与 url 均不能为空")
		}
	}

	detail, err := d.GetDigitalHumanDetail(ctx, digitalHumanID)
	if err != nil {
		return nil, err
	}

	merged := mergeBkn(detail.Bkn, entries)
	path := d.baseURL + "/api/dip-studio/v1/digital-human/:id"
	updated, err := base.PUT[*driven.DigitalHumanDetail](ctx, d.httpClient, path, updateDigitalHumanBknArgs{
		ID:  digitalHumanID,
		Bkn: merged,
	})
	if err != nil {
		return nil, err
	}
	return updated, nil
}
