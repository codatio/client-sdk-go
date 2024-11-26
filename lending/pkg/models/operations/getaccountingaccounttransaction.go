// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"net/http"
)

type GetAccountingAccountTransactionRequest struct {
	// Unique identifier for an account transaction.
	AccountTransactionID string `pathParam:"style=simple,explode=false,name=accountTransactionId"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetAccountingAccountTransactionRequest) GetAccountTransactionID() string {
	if o == nil {
		return ""
	}
	return o.AccountTransactionID
}

func (o *GetAccountingAccountTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingAccountTransactionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetAccountingAccountTransactionResponse struct {
	// Success
	AccountingAccountTransaction *shared.AccountingAccountTransaction
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingAccountTransactionResponse) GetAccountingAccountTransaction() *shared.AccountingAccountTransaction {
	if o == nil {
		return nil
	}
	return o.AccountingAccountTransaction
}

func (o *GetAccountingAccountTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingAccountTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingAccountTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
