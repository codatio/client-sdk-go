package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type DownloadCustomerAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

type DownloadCustomerAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DownloadCustomerAttachmentRequest struct {
	PathParams DownloadCustomerAttachmentPathParams
	Security   DownloadCustomerAttachmentSecurity
}

type DownloadCustomerAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
