package operations

import (
	"net/http"
	"time"
)

type DownloadFilesRequest struct {
	CompanyID string     `pathParam:"style=simple,explode=false,name=companyId"`
	Date      *time.Time `queryParam:"style=form,explode=true,name=date"`
}

type DownloadFilesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
