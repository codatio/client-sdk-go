package operations

import (
	"net/http"
)

type DownloadSupplierAttachmentRequest struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	SupplierID   string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type DownloadSupplierAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
