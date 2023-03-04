package operations

import (
	"net/http"
)

type UploadFilesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UploadFilesRequest struct {
	PathParams UploadFilesPathParams
}

type UploadFilesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
