package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type DownloadDirectIncomeAttachmentPathParams struct {
	AttachmentID   string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type DownloadDirectIncomeAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DownloadDirectIncomeAttachmentRequest struct {
	PathParams DownloadDirectIncomeAttachmentPathParams
	Security   DownloadDirectIncomeAttachmentSecurity
}

type DownloadDirectIncomeAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
