// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v4/pkg/models/shared"
	"net/http"
)

type GetAccountingInvoiceAttachmentRequest struct {
	// Unique identifier for an attachment
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for an invoice
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

func (o *GetAccountingInvoiceAttachmentRequest) GetAttachmentID() string {
	if o == nil {
		return ""
	}
	return o.AttachmentID
}

func (o *GetAccountingInvoiceAttachmentRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *GetAccountingInvoiceAttachmentRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetAccountingInvoiceAttachmentRequest) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type GetAccountingInvoiceAttachmentResponse struct {
	// Success
	AccountingAttachment *shared.AccountingAttachment
	ContentType          string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetAccountingInvoiceAttachmentResponse) GetAccountingAttachment() *shared.AccountingAttachment {
	if o == nil {
		return nil
	}
	return o.AccountingAttachment
}

func (o *GetAccountingInvoiceAttachmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingInvoiceAttachmentResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingInvoiceAttachmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingInvoiceAttachmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
