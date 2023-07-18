// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/banking/pkg/models/shared"
	"net/http"
)

type ListAccountBalancesRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *ListAccountBalancesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListAccountBalancesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListAccountBalancesRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListAccountBalancesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListAccountBalancesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListAccountBalancesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

// ListAccountBalances409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type ListAccountBalances409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

func (o *ListAccountBalances409ApplicationJSON) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *ListAccountBalances409ApplicationJSON) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *ListAccountBalances409ApplicationJSON) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *ListAccountBalances409ApplicationJSON) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListAccountBalances409ApplicationJSON) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ListAccountBalances409ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type ListAccountBalancesResponse struct {
	// Success
	AccountBalances *shared.AccountBalances
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
	// The data type's dataset has not been requested or is still syncing.
	ListAccountBalances409ApplicationJSONObject *ListAccountBalances409ApplicationJSON
	// Your `query` parameter was not correctly formed
	Schema *shared.Schema
}

func (o *ListAccountBalancesResponse) GetAccountBalances() *shared.AccountBalances {
	if o == nil {
		return nil
	}
	return o.AccountBalances
}

func (o *ListAccountBalancesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountBalancesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountBalancesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAccountBalancesResponse) GetListAccountBalances409ApplicationJSONObject() *ListAccountBalances409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListAccountBalances409ApplicationJSONObject
}

func (o *ListAccountBalancesResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
