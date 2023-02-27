package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type PushInvoiceAttachmentPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type PushInvoiceAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PushInvoiceAttachmentRequest struct {
	PathParams PushInvoiceAttachmentPathParams
	Security   PushInvoiceAttachmentSecurity
}

type PushInvoiceAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
