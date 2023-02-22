package operations

import (
	"github.com/codatio/client-sdk-go/files/pkg/models/shared"
)

type UploadFilesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UploadFilesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UploadFilesRequest struct {
	PathParams UploadFilesPathParams
	Security   UploadFilesSecurity
}

type UploadFilesResponse struct {
	ContentType string
	StatusCode  int
}
