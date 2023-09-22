// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/utils"
	"net/http"
)

type ListCommercePaymentsRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Field to order results by. [Read more](https://docs.codat.io/using-the-api/ordering-results).
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `default:"1" queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `default:"100" queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (l ListCommercePaymentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCommercePaymentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCommercePaymentsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListCommercePaymentsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListCommercePaymentsRequest) GetOrderBy() *string {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListCommercePaymentsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListCommercePaymentsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListCommercePaymentsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type ListCommercePaymentsResponse struct {
	// OK
	CommercePayments *shared.CommercePayments
	ContentType      string
	// Your `query` parameter was not correctly formed
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *ListCommercePaymentsResponse) GetCommercePayments() *shared.CommercePayments {
	if o == nil {
		return nil
	}
	return o.CommercePayments
}

func (o *ListCommercePaymentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCommercePaymentsResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ListCommercePaymentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCommercePaymentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
