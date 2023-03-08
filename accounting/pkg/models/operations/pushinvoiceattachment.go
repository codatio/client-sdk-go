package operations

import (
	"net/http"
)

type PushInvoiceAttachmentPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type PushInvoiceAttachmentRequest struct {
	PathParams PushInvoiceAttachmentPathParams
}

type PushInvoiceAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
