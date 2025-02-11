// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/models/shared"
)

type CreateBillPaymentRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a bill.
	BillID string `pathParam:"style=simple,explode=false,name=billId"`
	// A unique identifier to ensure idempotent behaviour for subsequent requests.
	IdempotencyKey       *string                      `header:"style=simple,explode=false,name=Idempotency-Key"`
	BillPaymentPrototype *shared.BillPaymentPrototype `request:"mediaType=application/json"`
}

func (o *CreateBillPaymentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBillPaymentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateBillPaymentRequest) GetBillID() string {
	if o == nil {
		return ""
	}
	return o.BillID
}

func (o *CreateBillPaymentRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *CreateBillPaymentRequest) GetBillPaymentPrototype() *shared.BillPaymentPrototype {
	if o == nil {
		return nil
	}
	return o.BillPaymentPrototype
}

type CreateBillPaymentResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Created
	BillPayment *shared.BillPayment
}

func (o *CreateBillPaymentResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateBillPaymentResponse) GetBillPayment() *shared.BillPayment {
	if o == nil {
		return nil
	}
	return o.BillPayment
}
