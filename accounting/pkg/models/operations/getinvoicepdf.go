package operations

import (
	"net/http"
)

type GetInvoicePdfRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type GetInvoicePdfResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
