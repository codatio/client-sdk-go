package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type PostDirectCostAttachmentPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type PostDirectCostAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostDirectCostAttachmentRequest struct {
	PathParams PostDirectCostAttachmentPathParams
	Security   PostDirectCostAttachmentSecurity
}

type PostDirectCostAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
