package operations

import (
	"net/http"
)

type DownloadDirectCostAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type DownloadDirectCostAttachmentRequest struct {
	PathParams DownloadDirectCostAttachmentPathParams
}

type DownloadDirectCostAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
