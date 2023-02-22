package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetDirectCostAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type GetDirectCostAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetDirectCostAttachmentRequest struct {
	PathParams GetDirectCostAttachmentPathParams
	Security   GetDirectCostAttachmentSecurity
}

type GetDirectCostAttachmentResponse struct {
	Attachment  *shared.Attachment
	ContentType string
	StatusCode  int
}
