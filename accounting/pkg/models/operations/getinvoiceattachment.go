package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetInvoiceAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type GetInvoiceAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetInvoiceAttachmentRequest struct {
	PathParams GetInvoiceAttachmentPathParams
	Security   GetInvoiceAttachmentSecurity
}

type GetInvoiceAttachmentResponse struct {
	Attachment  *shared.Attachment
	ContentType string
	StatusCode  int
}
