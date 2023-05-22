// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DownloadInvoicesAttachmentRequest struct {
	// Unique identifier for an attachment
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Unique identifier for an invoice
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type DownloadInvoicesAttachmentResponse struct {
	ContentType string
	// Success
	Data        []byte
	StatusCode  int
	RawResponse *http.Response
}