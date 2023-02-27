package operations

import (
	"github.com/codatio/client-sdk-go/files/pkg/models/shared"
	"time"
)

type ListFilesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListFilesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListFilesRequest struct {
	PathParams ListFilesPathParams
	Security   ListFilesSecurity
}

type ListFilesFile struct {
	DisplayName *string    `json:"displayName,omitempty"`
	FileName    *string    `json:"fileName,omitempty"`
	SourceType  *string    `json:"sourceType,omitempty"`
	Uploaded    *time.Time `json:"uploaded,omitempty"`
}

type ListFilesResponse struct {
	ContentType string
	Files       []ListFilesFile
	StatusCode  int
}
