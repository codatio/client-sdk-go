package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type DownloadBillAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type DownloadBillAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DownloadBillAttachmentRequest struct {
	PathParams DownloadBillAttachmentPathParams
	Security   DownloadBillAttachmentSecurity
}

type DownloadBillAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
