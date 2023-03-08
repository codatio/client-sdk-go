package operations

import (
	"net/http"
)

type DonwloadInvoiceAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type DonwloadInvoiceAttachmentRequest struct {
	PathParams DonwloadInvoiceAttachmentPathParams
}

type DonwloadInvoiceAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
