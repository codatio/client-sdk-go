// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
)

type GetConnectionRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetConnectionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetConnectionResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// OK
	Connection *shared.Connection
}

func (o *GetConnectionResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}
