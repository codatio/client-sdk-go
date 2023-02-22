package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type GetSupplierAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	SupplierID   string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type GetSupplierAttachmentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetSupplierAttachmentRequest struct {
	PathParams GetSupplierAttachmentPathParams
	Security   GetSupplierAttachmentSecurity
}

type GetSupplierAttachmentResponse struct {
	Attachment  *shared.Attachment
	ContentType string
	StatusCode  int
}
