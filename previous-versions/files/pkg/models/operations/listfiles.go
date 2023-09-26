// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"net/http"
)

type ListFilesRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *ListFilesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

// ListFilesErrorMessage - One or more of the resources you referenced could not be found.
// This might be because your company or data connection id is wrong, or was already deleted.
type ListFilesErrorMessage struct {
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

func (o *ListFilesErrorMessage) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *ListFilesErrorMessage) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *ListFilesErrorMessage) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *ListFilesErrorMessage) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ListFilesErrorMessage) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ListFilesErrorMessage) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type ListFilesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// One or more of the resources you referenced could not be found.
	// This might be because your company or data connection id is wrong, or was already deleted.
	ErrorMessage *ListFilesErrorMessage
	// Success
	Files []shared.File
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

func (o *ListFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFilesResponse) GetErrorMessage() *ListFilesErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ListFilesResponse) GetFiles() []shared.File {
	if o == nil {
		return nil
	}
	return o.Files
}

func (o *ListFilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListFilesResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
