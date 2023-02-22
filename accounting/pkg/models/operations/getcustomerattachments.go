package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetCustomerAttachmentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

type GetCustomerAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetCustomerAttachmentsRequest struct {
	PathParams GetCustomerAttachmentsPathParams
	Security   GetCustomerAttachmentsSecurity
}

type GetCustomerAttachmentsResponse struct {
	AttachmentsDataset *shared.AttachmentsDataset
	ContentType        string
	StatusCode         int
}
