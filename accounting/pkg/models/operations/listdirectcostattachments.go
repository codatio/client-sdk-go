package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListDirectCostAttachmentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type ListDirectCostAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListDirectCostAttachmentsRequest struct {
	PathParams ListDirectCostAttachmentsPathParams
	Security   ListDirectCostAttachmentsSecurity
}

type ListDirectCostAttachmentsResponse struct {
	AttachmentsDataset *shared.AttachmentsDataset
	ContentType        string
	StatusCode         int
}
