package operations

import (
	"net/http"
)

type DownloadBillAttachmentRequest struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	BillID       string `pathParam:"style=simple,explode=false,name=billId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type DownloadBillAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
