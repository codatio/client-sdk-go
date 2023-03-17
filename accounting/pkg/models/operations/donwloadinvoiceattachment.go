package operations

import (
	"net/http"
)

type DonwloadInvoiceAttachmentRequest struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type DonwloadInvoiceAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
