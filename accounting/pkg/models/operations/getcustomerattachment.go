package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetCustomerAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

type GetCustomerAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCustomerAttachmentRequest struct {
	PathParams GetCustomerAttachmentPathParams
	Security   GetCustomerAttachmentSecurity
}

type GetCustomerAttachmentResponse struct {
	Attachment  *shared.Attachment
	ContentType string
	StatusCode  int
}
