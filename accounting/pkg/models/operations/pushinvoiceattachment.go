package operations

import (
	"net/http"
)

type PushInvoiceAttachmentRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type PushInvoiceAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
