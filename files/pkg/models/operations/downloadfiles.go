package operations

import (
	"github.com/codatio/client-sdk-go/files/pkg/models/shared"
	"time"
)

type DownloadFilesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type DownloadFilesQueryParams struct {
	Date *time.Time `queryParam:"style=form,explode=true,name=date"`
}

type DownloadFilesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DownloadFilesRequest struct {
	PathParams  DownloadFilesPathParams
	QueryParams DownloadFilesQueryParams
	Security    DownloadFilesSecurity
}

type DownloadFilesResponse struct {
	ContentType string
	StatusCode  int
}
