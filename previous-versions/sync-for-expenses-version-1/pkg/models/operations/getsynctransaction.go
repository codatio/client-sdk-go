// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"net/http"
)

type GetSyncTransactionRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a sync.
	SyncID string `pathParam:"style=simple,explode=false,name=syncId"`
	// The unique identifier for your SMB's transaction.
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionId"`
}

func (o *GetSyncTransactionRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetSyncTransactionRequest) GetSyncID() string {
	if o == nil {
		return ""
	}
	return o.SyncID
}

func (o *GetSyncTransactionRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

type GetSyncTransactionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Classes []shared.TransactionMetadata
}

func (o *GetSyncTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSyncTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSyncTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSyncTransactionResponse) GetClasses() []shared.TransactionMetadata {
	if o == nil {
		return nil
	}
	return o.Classes
}
