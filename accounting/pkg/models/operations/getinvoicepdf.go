package operations

import (
	"net/http"
)

type GetInvoicePdfPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type GetInvoicePdfRequest struct {
	PathParams GetInvoicePdfPathParams
}

type GetInvoicePdfResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
