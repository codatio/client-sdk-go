// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type ListJournalsRequest struct {
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

func (o *ListJournalsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListJournalsRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListJournalsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListJournalsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListJournalsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

// ListJournals409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type ListJournals409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

func (o *ListJournals409ApplicationJSON) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *ListJournals409ApplicationJSON) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *ListJournals409ApplicationJSON) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *ListJournals409ApplicationJSON) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListJournals409ApplicationJSON) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ListJournals409ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type ListJournalsResponse struct {
	ContentType string
	// Success
	Journals    *shared.Journals
	StatusCode  int
	RawResponse *http.Response
	// The data type's dataset has not been requested or is still syncing.
	ListJournals409ApplicationJSONObject *ListJournals409ApplicationJSON
	// Your `query` parameter was not correctly formed
	Schema *shared.Schema
}

func (o *ListJournalsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListJournalsResponse) GetJournals() *shared.Journals {
	if o == nil {
		return nil
	}
	return o.Journals
}

func (o *ListJournalsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListJournalsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListJournalsResponse) GetListJournals409ApplicationJSONObject() *ListJournals409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListJournals409ApplicationJSONObject
}

func (o *ListJournalsResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
