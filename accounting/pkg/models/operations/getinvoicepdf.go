package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetInvoicePdfPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	InvoiceID string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type GetInvoicePdfSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetInvoicePdfRequest struct {
	PathParams GetInvoicePdfPathParams
	Security   GetInvoicePdfSecurity
}

type GetInvoicePdfResponse struct {
	ContentType string
	StatusCode  int
}
