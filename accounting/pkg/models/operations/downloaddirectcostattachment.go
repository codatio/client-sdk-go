package operations

import (
	"net/http"
)

type DownloadDirectCostAttachmentRequest struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type DownloadDirectCostAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
