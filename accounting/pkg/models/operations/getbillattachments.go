package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetBillAttachmentsPathParams struct {
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetBillAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetBillAttachmentsRequest struct {
	PathParams GetBillAttachmentsPathParams
	Security   GetBillAttachmentsSecurity
}

type GetBillAttachmentsResponse struct {
	AttachmentsDataset *shared.AttachmentsDataset
	ContentType        string
	StatusCode         int
}
