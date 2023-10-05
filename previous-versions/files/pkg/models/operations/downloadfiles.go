// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
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

// DownloadFilesErrorMessage - One or more of the resources you referenced could not be found.
// This might be because your company or data connection id is wrong, or was already deleted.
type DownloadFilesErrorMessage struct {
	// `True` if the error occurred transiently and can be retried.
	CanBeRetried *string `json:"canBeRetried,omitempty"`
	// Unique identifier used to propagate to all downstream services and determine the source of the error.
	CorrelationID *string `json:"correlationId,omitempty"`
	// Machine readable error code used to automate processes based on the code returned.
	DetailedErrorCode *int64 `json:"detailedErrorCode,omitempty"`
	// A brief description of the error.
	Error *string `json:"error,omitempty"`
	// Codat's service the returned the error.
	Service *string `json:"service,omitempty"`
	// The HTTP status code returned by the error.
	StatusCode *int64 `json:"statusCode,omitempty"`
}

func (o *DownloadFilesErrorMessage) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *DownloadFilesErrorMessage) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *DownloadFilesErrorMessage) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *DownloadFilesErrorMessage) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DownloadFilesErrorMessage) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *DownloadFilesErrorMessage) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type DownloadFilesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	Data []byte
	// One or more of the resources you referenced could not be found.
	// This might be because your company or data connection id is wrong, or was already deleted.
	ErrorMessage *DownloadFilesErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request made is not valid.
	Schema *shared.Schema
}

func (o *DownloadFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadFilesResponse) GetData() []byte {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *DownloadFilesResponse) GetErrorMessage() *DownloadFilesErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
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

func (o *DownloadFilesResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
