// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
	"net/http"
)

type GetEnhancedInvoicesReportRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Page number. [Read more](https://docs.codat.io/using-the-api/paging).
	Page *int `queryParam:"style=form,explode=true,name=page"`
	// Number of records to return in a page. [Read more](https://docs.codat.io/using-the-api/paging).
	PageSize *int `queryParam:"style=form,explode=true,name=pageSize"`
	// Codat query string. [Read more](https://docs.codat.io/using-the-api/querying).
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *GetEnhancedInvoicesReportRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetEnhancedInvoicesReportRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetEnhancedInvoicesReportRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *GetEnhancedInvoicesReportRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type GetEnhancedInvoicesReportResponse struct {
	ContentType string
	// OK
	EnhancedInvoicesReport *shared.EnhancedInvoicesReport
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetEnhancedInvoicesReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEnhancedInvoicesReportResponse) GetEnhancedInvoicesReport() *shared.EnhancedInvoicesReport {
	if o == nil {
		return nil
	}
	return o.EnhancedInvoicesReport
}

func (o *GetEnhancedInvoicesReportResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetEnhancedInvoicesReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEnhancedInvoicesReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
