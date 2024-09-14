// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/pkg/models/shared"
)

type CreateConnectionRequestBody struct {
	// A unique 4-letter key to represent a platform in each integration. View [accounting](https://docs.codat.io/integrations/accounting/overview#platform-keys), [banking](https://docs.codat.io/integrations/banking/overview#platform-keys), and [commerce](https://docs.codat.io/integrations/commerce/overview#platform-keys) platform keys.
	PlatformKey *string `json:"platformKey,omitempty"`
}

func (o *CreateConnectionRequestBody) GetPlatformKey() *string {
	if o == nil {
		return nil
	}
	return o.PlatformKey
}

type CreateConnectionRequest struct {
	// Unique identifier for a company.
	CompanyID   string                       `pathParam:"style=simple,explode=false,name=companyId"`
	RequestBody *CreateConnectionRequestBody `request:"mediaType=application/json"`
}

func (o *CreateConnectionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateConnectionRequest) GetRequestBody() *CreateConnectionRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type CreateConnectionResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// OK
	Connection *shared.Connection
}

func (o *CreateConnectionResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateConnectionResponse) GetConnection() *shared.Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}