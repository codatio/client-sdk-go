package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetBillAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBillAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBillAttachmentRequest struct {
	PathParams GetBillAttachmentPathParams
	Security   GetBillAttachmentSecurity
}

type GetBillAttachmentResponse struct {
	Attachment  *shared.Attachment
	ContentType string
	StatusCode  int
}
