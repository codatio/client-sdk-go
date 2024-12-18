// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/models/shared"
	"net/http"
)

type ListAccountingInvoiceAttachmentsRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for an invoice.
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

func (o *ListAccountingInvoiceAttachmentsRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *ListAccountingInvoiceAttachmentsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListAccountingInvoiceAttachmentsRequest) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type ListAccountingInvoiceAttachmentsResponse struct {
	// Success
	Attachments *shared.Attachments
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountingInvoiceAttachmentsResponse) GetAttachments() *shared.Attachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *ListAccountingInvoiceAttachmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountingInvoiceAttachmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountingInvoiceAttachmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
