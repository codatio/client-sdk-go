// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"io"
	"net/http"
)

type DownloadFilesRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Only download files uploaded on this date.
	Date *string `queryParam:"style=form,explode=true,name=date"`
}

func (o *DownloadFilesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *DownloadFilesRequest) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

type DownloadFilesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Data io.ReadCloser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DownloadFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadFilesResponse) GetData() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DownloadFilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadFilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
