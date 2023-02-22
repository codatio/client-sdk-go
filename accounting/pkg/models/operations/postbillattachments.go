package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type PostBillAttachmentsPathParams struct {
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostBillAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostBillAttachmentsRequest struct {
	PathParams PostBillAttachmentsPathParams
	Security   PostBillAttachmentsSecurity
}

type PostBillAttachmentsResponse struct {
	ContentType string
	StatusCode  int
}
