// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
)

type CreateSupplierRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// A unique identifier to ensure idempotent behaviour for subsequent requests.
	IdempotencyKey    *string                   `header:"style=simple,explode=false,name=Idempotency-Key"`
	SupplierPrototype *shared.SupplierPrototype `request:"mediaType=application/json"`
}

func (o *CreateSupplierRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateSupplierRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateSupplierRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *CreateSupplierRequest) GetSupplierPrototype() *shared.SupplierPrototype {
	if o == nil {
		return nil
	}
	return o.SupplierPrototype
}

type CreateSupplierResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Success
	Supplier *shared.Supplier
}

func (o *CreateSupplierResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateSupplierResponse) GetSupplier() *shared.Supplier {
	if o == nil {
		return nil
	}
	return o.Supplier
}
