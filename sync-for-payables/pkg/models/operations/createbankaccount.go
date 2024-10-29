// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v3/pkg/models/shared"
)

type CreateBankAccountRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// A unique identifier to ensure idempotent behaviour for subsequent requests.
	IdempotencyKey       *string                      `header:"style=simple,explode=false,name=Idempotency-Key"`
	BankAccountPrototype *shared.BankAccountPrototype `request:"mediaType=application/json"`
}

func (o *CreateBankAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBankAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateBankAccountRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *CreateBankAccountRequest) GetBankAccountPrototype() *shared.BankAccountPrototype {
	if o == nil {
		return nil
	}
	return o.BankAccountPrototype
}

type CreateBankAccountResponse struct {
	HTTPMeta shared.HTTPMetadata `json:"-"`
	// Created
	BankAccount *shared.BankAccount
}

func (o *CreateBankAccountResponse) GetHTTPMeta() shared.HTTPMetadata {
	if o == nil {
		return shared.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateBankAccountResponse) GetBankAccount() *shared.BankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccount
}
