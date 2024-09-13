// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type GetAccountTransactionRequest struct {
	// Unique identifier for an account transaction.
	AccountTransactionID string `pathParam:"style=simple,explode=false,name=accountTransactionId"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *GetAccountTransactionRequest) GetAccountTransactionID() string {
	if o == nil {
		return ""
	}
	return o.AccountTransactionID
}

func (o *GetAccountTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountTransactionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetAccountTransactionResponse struct {
	// Success
	AccountTransaction *shared.AccountTransaction
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountTransactionResponse) GetAccountTransaction() *shared.AccountTransaction {
	if o == nil {
		return nil
	}
	return o.AccountTransaction
}

func (o *GetAccountTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
