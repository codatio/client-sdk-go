package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetInvoicePathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type GetInvoiceSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetInvoiceRequest struct {
	PathParams GetInvoicePathParams
	Security   GetInvoiceSecurity
}

type GetInvoiceResponse struct {
	ContentType string
	Invoice     *shared.Invoice
	StatusCode  int
}
