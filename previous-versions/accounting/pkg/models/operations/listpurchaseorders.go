// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"net/http"
)

type ListPurchaseOrdersRequest struct {
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

func (o *ListPurchaseOrdersRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListPurchaseOrdersRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListPurchaseOrdersRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListPurchaseOrdersRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListPurchaseOrdersRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type ListPurchaseOrdersResponse struct {
	ContentType string
	// Your `query` parameter was not correctly formed
	ErrorMessage *shared.ErrorMessage
	// Success
	PurchaseOrders *shared.PurchaseOrders
	StatusCode     int
	RawResponse    *http.Response

	Next func() (*ListPurchaseOrdersResponse, error)
}

func (o *ListPurchaseOrdersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPurchaseOrdersResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ListPurchaseOrdersResponse) GetPurchaseOrders() *shared.PurchaseOrders {
	if o == nil {
		return nil
	}
	return o.PurchaseOrders
}

func (o *ListPurchaseOrdersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPurchaseOrdersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
