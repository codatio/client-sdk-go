package operations

import (
	"net/http"
	"time"
)

type ListFilesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListFilesRequest struct {
	PathParams ListFilesPathParams
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
	RawResponse *http.Response
}
