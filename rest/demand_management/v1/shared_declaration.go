package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	demand_management_v1 "github.com/kweaver-ai/idrm-go-common/api/demand_management/v1"
	"github.com/kweaver-ai/idrm-go-common/errorcode"
	"github.com/kweaver-ai/idrm-go-common/interception"
)

// SharedDeclarationsGetter has a method to return a SharedDeclarationInterface.
// A group's client should implement this interface.
type SharedDeclarationGetter interface {
	SharedDeclaration() SharedDeclarationInterface
}

// SharedDeclarationInterface has methods to work with SharedDeclaration
// resources.
type SharedDeclarationInterface interface {
	// 查询目录共享申请状态
	Status(ctx context.Context, req *demand_management_v1.SharedDeclarationStatusReq) ([]demand_management_v1.SharedDeclarationStatusResp, error)
}

// sharedDeclaration implements SharedDeclarationInterface
type sharedDeclaration struct {
	base   *url.URL
	client *http.Client
}

var _ SharedDeclarationInterface = &sharedDeclaration{}

func newSharedDeclaration(c *DemandManagementV1Client) *sharedDeclaration {
	base := *c.base
	base.Path = path.Join(base.Path, "shared-declaration")
	return &sharedDeclaration{
		base:   &base,
		client: c.client,
	}
}

// 查询目录共享申请状态
func (c *sharedDeclaration) Status(ctx context.Context, r *demand_management_v1.SharedDeclarationStatusReq) ([]demand_management_v1.SharedDeclarationStatusResp, error) {
	// endpoint
	base := *c.base
	base.Path = path.Join(base.Path, "status")

	reqJSON, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	// create http request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base.String(), bytes.NewReader(reqJSON))
	if err != nil {
		return nil, err
	}

	// authorization
	if t, err := interception.BearerTokenFromContextCompatible(ctx); err == nil {
		req.Header.Set("Authorization", "Bearer "+t)
	}

	// send http request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errorcode.NewErrorForHTTPResponse(resp)
	}

	// decode http response body
	var result []demand_management_v1.SharedDeclarationStatusResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
