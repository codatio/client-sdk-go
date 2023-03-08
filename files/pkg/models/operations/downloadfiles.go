package operations

import (
	"net/http"
	"time"
)

type DownloadFilesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type DownloadFilesQueryParams struct {
	Date *time.Time `queryParam:"style=form,explode=true,name=date"`
}

type DownloadFilesRequest struct {
	PathParams  DownloadFilesPathParams
	QueryParams DownloadFilesQueryParams
}

type DownloadFilesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
