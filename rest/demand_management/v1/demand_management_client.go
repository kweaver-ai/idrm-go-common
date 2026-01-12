package v1

import (
	"net/http"
	"net/url"
)

type DemandManagementV1Interface interface {
	SharedDeclarationGetter
}

// DemandManagementV1Client is used to interact with features provided by the
// demand_management group.
type DemandManagementV1Client struct {
	base   *url.URL
	client *http.Client
}

var _ DemandManagementV1Interface = &DemandManagementV1Client{}

// New creates a new DemandManagementV1Client for the given url.URL and
// http.Client.
func New(base *url.URL, client *http.Client) *DemandManagementV1Client {
	return &DemandManagementV1Client{
		base:   base,
		client: client,
	}
}

func (c *DemandManagementV1Client) SharedDeclaration() SharedDeclarationInterface {
	return newSharedDeclaration(c)
}
