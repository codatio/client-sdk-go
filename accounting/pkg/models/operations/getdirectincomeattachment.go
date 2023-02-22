package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetDirectIncomeAttachmentPathParams struct {
	AttachmentID   string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type GetDirectIncomeAttachmentQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type GetDirectIncomeAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetDirectIncomeAttachmentRequest struct {
	PathParams  GetDirectIncomeAttachmentPathParams
	QueryParams GetDirectIncomeAttachmentQueryParams
	Security    GetDirectIncomeAttachmentSecurity
}

type GetDirectIncomeAttachmentResponse struct {
	Attachment  *shared.Attachment
	ContentType string
	StatusCode  int
}
