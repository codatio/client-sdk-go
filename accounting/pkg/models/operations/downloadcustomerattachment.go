package operations

import (
	"net/http"
)

type DownloadCustomerAttachmentPathParams struct {
	AttachmentID string `pathParam:"style=simple,explode=false,name=attachmentId"`
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

type DownloadCustomerAttachmentRequest struct {
	PathParams DownloadCustomerAttachmentPathParams
}

type DownloadCustomerAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
