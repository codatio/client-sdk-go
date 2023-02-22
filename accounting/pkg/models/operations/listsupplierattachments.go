package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListSupplierAttachmentsPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	SupplierID   string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type ListSupplierAttachmentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListSupplierAttachmentsRequest struct {
	PathParams ListSupplierAttachmentsPathParams
	Security   ListSupplierAttachmentsSecurity
}

type ListSupplierAttachmentsResponse struct {
	AttachmentsDataset *shared.AttachmentsDataset
	ContentType        string
	StatusCode         int
}
