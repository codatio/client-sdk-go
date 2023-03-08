package operations

import (
	"net/http"
)

type DownloadSupplierAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	SupplierID   string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type DownloadSupplierAttachmentRequest struct {
	PathParams DownloadSupplierAttachmentPathParams
}

type DownloadSupplierAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
