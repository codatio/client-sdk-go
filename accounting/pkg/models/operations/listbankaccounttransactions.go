// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type ListBankAccountTransactionsRequest struct {
	// Unique identifier for an account
	AccountID    string `pathParam:"style=simple,explode=false,name=accountId"`
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

func (o *ListBankAccountTransactionsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListBankAccountTransactionsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListBankAccountTransactionsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListBankAccountTransactionsRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListBankAccountTransactionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListBankAccountTransactionsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListBankAccountTransactionsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

// ListBankAccountTransactions409ApplicationJSON - The data type's dataset has not been requested or is still syncing.
type ListBankAccountTransactions409ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

func (o *ListBankAccountTransactions409ApplicationJSON) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *ListBankAccountTransactions409ApplicationJSON) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *ListBankAccountTransactions409ApplicationJSON) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *ListBankAccountTransactions409ApplicationJSON) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListBankAccountTransactions409ApplicationJSON) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ListBankAccountTransactions409ApplicationJSON) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type ListBankAccountTransactionsResponse struct {
	// Success
	BankTransactionsResponse *shared.BankTransactionsResponse
	ContentType              string
	StatusCode               int
	RawResponse              *http.Response
	// The data type's dataset has not been requested or is still syncing.
	ListBankAccountTransactions409ApplicationJSONObject *ListBankAccountTransactions409ApplicationJSON
	// Your `query` parameter was not correctly formed
	Schema *shared.Schema
}

func (o *ListBankAccountTransactionsResponse) GetBankTransactionsResponse() *shared.BankTransactionsResponse {
	if o == nil {
		return nil
	}
	return o.BankTransactionsResponse
}

func (o *ListBankAccountTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListBankAccountTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListBankAccountTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListBankAccountTransactionsResponse) GetListBankAccountTransactions409ApplicationJSONObject() *ListBankAccountTransactions409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListBankAccountTransactions409ApplicationJSONObject
}

func (o *ListBankAccountTransactionsResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
