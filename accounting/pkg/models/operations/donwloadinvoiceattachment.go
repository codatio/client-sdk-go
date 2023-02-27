package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type DonwloadInvoiceAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type DonwloadInvoiceAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DonwloadInvoiceAttachmentRequest struct {
	PathParams DonwloadInvoiceAttachmentPathParams
	Security   DonwloadInvoiceAttachmentSecurity
}

type DonwloadInvoiceAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
