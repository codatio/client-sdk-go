package operations

import (
	"net/http"
)

type PostDirectCostAttachmentRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	DirectCostID string `pathParam:"style=simple,explode=false,name=directCostId"`
}

type PostDirectCostAttachmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
