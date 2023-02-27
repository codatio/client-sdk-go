package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type DownloadDirectCostAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type DownloadDirectCostAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DownloadDirectCostAttachmentRequest struct {
	PathParams DownloadDirectCostAttachmentPathParams
	Security   DownloadDirectCostAttachmentSecurity
}

type DownloadDirectCostAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
