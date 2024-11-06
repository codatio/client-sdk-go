// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	"net/http"
)

type GetAccountingInvoiceRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for an invoice.
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

func (o *GetAccountingInvoiceRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingInvoiceRequest) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type GetAccountingInvoiceResponse struct {
	// Success
	AccountingInvoice *shared.AccountingInvoice
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingInvoiceResponse) GetAccountingInvoice() *shared.AccountingInvoice {
	if o == nil {
		return nil
	}
	return o.AccountingInvoice
}

func (o *GetAccountingInvoiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingInvoiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingInvoiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
