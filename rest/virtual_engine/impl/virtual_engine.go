package impl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/kweaver-ai/idrm-go-common/rest/base"
	"github.com/kweaver-ai/idrm-go-common/rest/virtual_engine"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"go.uber.org/zap"
)

type DrivenImpl struct {
	baseURL    string
	httpClient *http.Client
}

func NewDrivenImpl(client *http.Client) virtual_engine.Driven {
	return &DrivenImpl{
		baseURL:    base.Service.VirtualEngineHost,
		httpClient: client,
	}
}

func (d DrivenImpl) DatabaseTypeMapping(ctx context.Context, req *virtual_engine.DatabaseTypeMappingReq) (*virtual_engine.DatabaseTypeMappingResp, error) {
	path := "/api/virtual_engine_service/v1/connectors/type/mapping"
	url := d.baseURL + path
	//处理参数
	resp, err := base.POST[virtual_engine.DatabaseTypeMappingResp](ctx, d.httpClient, url, req)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (d DrivenImpl) queryTableDetail(ctx context.Context, catalog, schema, table string) ([]virtual_engine.TableWithFieldsResp, error) {
	path := "/api/virtual_engine_service/v1/metadata/columns/" + catalog + "/" + strings.ToLower(schema) + "/" + table
	url0 := d.baseURL + path
	resp, err := base.GET[base.CommonResponse[[]virtual_engine.TableWithFieldsResp]](ctx, d.httpClient, url0, nil)
	if err != nil {
		return nil, err
	}
	return resp.Data, err
}

func (d DrivenImpl) TableFieldsDetail(ctx context.Context, catalog, schema, table string) ([]virtual_engine.DataTableFieldResp, error) {
	details, err := d.queryTableDetail(ctx, catalog, schema, table)
	if err != nil {
		return nil, err
	}
	res := make([]virtual_engine.DataTableFieldResp, 0)
	for _, field := range details {
		length, precision := parsePrecision(field.OrigType)
		res = append(res, virtual_engine.DataTableFieldResp{
			Name:           field.Name,
			Type:           field.Type,
			Description:    field.Description,
			RawType:        field.TypeSignature.RawType,
			OrigType:       field.OrigType,
			Length:         length,
			FieldPrecision: precision,
		})
	}
	return res, err
}

func (d DrivenImpl) DBConnectorConfig(ctx context.Context, connector string) (*virtual_engine.ConnectorConfigResp, error) {
	path := "/api/virtual_engine_service/v1/connectors/config/" + connector
	url0 := d.baseURL + path
	resp, err := base.GET[virtual_engine.ConnectorConfigResp](ctx, d.httpClient, url0, nil)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (v *DrivenImpl) GetPreview(ctx context.Context, req *virtual_engine.ViewEntries) (*virtual_engine.FetchDataRes, error) {
	drivenMsg := "DrivenVirtualizationEngine GetPreview "
	log.Infof(drivenMsg+"%+v", *req)
	urlStr := fmt.Sprintf("%s/api/virtual_engine_service/v1/preview/%s/%s/%s?limit=%d&user_id=%s", v.baseURL, req.CatalogName, req.Schema, req.ViewName, req.Limit, req.UserId)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		log.WithContext(ctx).Error(drivenMsg+"http.NewRequest error", zap.Error(err))
		return nil, err
	}
	request.Header.Add("X-Presto-User", "admin")

	resp, err := v.httpClient.Do(request)
	if err != nil {
		log.WithContext(ctx).Error(drivenMsg+"client.Do error", zap.Error(err))
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithContext(ctx).Error(drivenMsg+" io.ReadAll error", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		res := &virtual_engine.FetchDataRes{}
		if err = jsoniter.Unmarshal(body, &res); err != nil {
			log.WithContext(ctx).Error(drivenMsg+" jsoniter.Unmarshal error", zap.Error(err))
			return nil, err
		}
		return res, nil
	} else {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusInternalServerError || resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			return nil, base.Unmarshal(ctx, body, drivenMsg)
		} else {
			log.WithContext(ctx).Error(drivenMsg+"http status error", zap.String("status", resp.Status), zap.String("body", string(body)))
			return nil, err
		}
	}
}
