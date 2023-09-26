// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DownloadInvoicePdfRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for an invoice.
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

func (o *DownloadInvoicePdfRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DownloadInvoicePdfRequest) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type DownloadInvoicePdfResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	Data []byte
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DownloadInvoicePdfResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadInvoicePdfResponse) GetData() []byte {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DownloadInvoicePdfResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadInvoicePdfResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
