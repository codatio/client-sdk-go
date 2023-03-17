package operations

import (
	"net/http"
)

type DownloadDirectIncomeAttachmentRequest struct {
	AttachmentID   string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID   string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectIncomeID string `pathParam:"style=simple,explode=false,name=directIncomeId"`
}

type DownloadDirectIncomeAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
