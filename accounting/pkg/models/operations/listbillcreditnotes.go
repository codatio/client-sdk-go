// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type ListBillCreditNotesRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *ListBillCreditNotesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListBillCreditNotesRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListBillCreditNotesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListBillCreditNotesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListBillCreditNotesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

// ListBillCreditNotes409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type ListBillCreditNotes409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

func (o *ListBillCreditNotes409ApplicationJSON) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *ListBillCreditNotes409ApplicationJSON) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *ListBillCreditNotes409ApplicationJSON) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *ListBillCreditNotes409ApplicationJSON) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListBillCreditNotes409ApplicationJSON) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ListBillCreditNotes409ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type ListBillCreditNotesResponse struct {
	// Success
	BillCreditNotes *shared.BillCreditNotes
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
	// The data type's dataset has not been requested or is still syncing.
	ListBillCreditNotes409ApplicationJSONObject *ListBillCreditNotes409ApplicationJSON
	// Your `query` parameter was not correctly formed
	Schema *shared.Schema
}

func (o *ListBillCreditNotesResponse) GetBillCreditNotes() *shared.BillCreditNotes {
	if o == nil {
		return nil
	}
	return o.BillCreditNotes
}

func (o *ListBillCreditNotesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListBillCreditNotesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListBillCreditNotesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListBillCreditNotesResponse) GetListBillCreditNotes409ApplicationJSONObject() *ListBillCreditNotes409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListBillCreditNotes409ApplicationJSONObject
}

func (o *ListBillCreditNotesResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
