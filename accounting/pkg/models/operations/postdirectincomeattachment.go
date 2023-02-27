package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type PostDirectIncomeAttachmentPathParams struct {
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type PostDirectIncomeAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostDirectIncomeAttachmentRequest struct {
	PathParams PostDirectIncomeAttachmentPathParams
	Security   PostDirectIncomeAttachmentSecurity
}

type PostDirectIncomeAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
