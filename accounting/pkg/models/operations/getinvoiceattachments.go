package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetInvoiceAttachmentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	InvoiceID    string `pathParam:"style=simple,explode=false,name=invoiceId"`
}

type GetInvoiceAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetInvoiceAttachmentsRequest struct {
	PathParams GetInvoiceAttachmentsPathParams
	Security   GetInvoiceAttachmentsSecurity
}

type GetInvoiceAttachmentsResponse struct {
	AttachmentsDataset *shared.AttachmentsDataset
	ContentType        string
	StatusCode         int
}
