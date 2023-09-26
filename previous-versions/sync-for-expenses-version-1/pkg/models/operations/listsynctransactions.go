// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/utils"
	"net/http"
)

type ListSyncTransactionsRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `default:"100" queryParam:"style=form,explode=true,name=pageSize"`
	// Unique identifier for a sync.
	SyncID string `pathParam:"style=simple,explode=false,name=syncId"`
}

func (l ListSyncTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListSyncTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListSyncTransactionsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListSyncTransactionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListSyncTransactionsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListSyncTransactionsRequest) GetSyncID() string {
	if o == nil {
		return ""
	}
	return o.SyncID
}

type ListSyncTransactionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	TransactionMetadataList *shared.TransactionMetadataList
}

func (o *ListSyncTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSyncTransactionsResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ListSyncTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSyncTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSyncTransactionsResponse) GetTransactionMetadataList() *shared.TransactionMetadataList {
	if o == nil {
		return nil
	}
	return o.TransactionMetadataList
}
