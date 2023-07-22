// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type ListAccountsRequest struct {
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

func (o *ListAccountsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListAccountsRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListAccountsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAccountsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListAccountsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

// ListAccounts409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type ListAccounts409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

func (o *ListAccounts409ApplicationJSON) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *ListAccounts409ApplicationJSON) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *ListAccounts409ApplicationJSON) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *ListAccounts409ApplicationJSON) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListAccounts409ApplicationJSON) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ListAccounts409ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type ListAccountsResponse struct {
	// Success
	Accounts    *shared.Accounts
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The data type's dataset has not been requested or is still syncing.
	ListAccounts409ApplicationJSONObject *ListAccounts409ApplicationJSON
	// Your `query` parameter was not correctly formed
	Schema *shared.Schema
}

func (o *ListAccountsResponse) GetAccounts() *shared.Accounts {
	if o == nil {
		return nil
	}
	return o.Accounts
}

func (o *ListAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAccountsResponse) GetListAccounts409ApplicationJSONObject() *ListAccounts409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListAccounts409ApplicationJSONObject
}

func (o *ListAccountsResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
