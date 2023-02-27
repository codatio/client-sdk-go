package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type DownloadSupplierAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	SupplierID   string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type DownloadSupplierAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DownloadSupplierAttachmentRequest struct {
	PathParams DownloadSupplierAttachmentPathParams
	Security   DownloadSupplierAttachmentSecurity
}

type DownloadSupplierAttachmentResponse struct {
	ContentType string
	StatusCode  int
}
