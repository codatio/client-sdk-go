// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v4/pkg/models/shared"
)

type CreateBillRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// A unique identifier to ensure idempotent behaviour for subsequent requests.
	IdempotencyKey *string               `header:"style=simple,explode=false,name=Idempotency-Key"`
	BillPrototype  *shared.BillPrototype `request:"mediaType=application/json"`
}

func (o *CreateBillRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBillRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateBillRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *CreateBillRequest) GetBillPrototype() *shared.BillPrototype {
	if o == nil {
		return nil
	}
	return o.BillPrototype
}

type CreateBillResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Created
	Bill *shared.Bill
}

func (o *CreateBillResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateBillResponse) GetBill() *shared.Bill {
	if o == nil {
		return nil
	}
	return o.Bill
}
