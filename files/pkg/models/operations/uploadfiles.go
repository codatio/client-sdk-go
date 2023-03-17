package operations

import (
	"net/http"
)

type UploadFilesRequest struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UploadFilesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
