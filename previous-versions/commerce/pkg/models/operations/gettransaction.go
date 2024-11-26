// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/models/shared"
	"net/http"
)

type GetTransactionRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for a transaction.
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

func (o *GetTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetTransactionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetTransactionRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

type GetTransactionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Transaction *shared.Transaction
}

func (o *GetTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransactionResponse) GetTransaction() *shared.Transaction {
	if o == nil {
		return nil
	}
	return o.Transaction
}
