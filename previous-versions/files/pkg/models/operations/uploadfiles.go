// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/files/pkg/models/shared"
	"net/http"
)

type UploadFilesRequestBody struct {
	Content     []byte `multipartForm:"content"`
	RequestBody string `multipartForm:"name=requestBody"`
}

func (o *UploadFilesRequestBody) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *UploadFilesRequestBody) GetRequestBody() string {
	if o == nil {
		return ""
	}
	return o.RequestBody
}

type UploadFilesRequest struct {
	RequestBody *UploadFilesRequestBody `multipartForm:"file" request:"mediaType=multipart/form-data"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

func (o *UploadFilesRequest) GetRequestBody() *UploadFilesRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UploadFilesRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UploadFilesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

// UploadFilesErrorMessage - One or more of the resources you referenced could not be found.
// This might be because your company or data connection id is wrong, or was already deleted.
type UploadFilesErrorMessage struct {
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

func (o *UploadFilesErrorMessage) GetCanBeRetried() *string {
	if o == nil {
		return nil
	}
	return o.CanBeRetried
}

func (o *UploadFilesErrorMessage) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *UploadFilesErrorMessage) GetDetailedErrorCode() *int64 {
	if o == nil {
		return nil
	}
	return o.DetailedErrorCode
}

func (o *UploadFilesErrorMessage) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *UploadFilesErrorMessage) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *UploadFilesErrorMessage) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

type UploadFilesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// One or more of the resources you referenced could not be found.
	// This might be because your company or data connection id is wrong, or was already deleted.
	ErrorMessage *UploadFilesErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request made is not valid.
	Schema *shared.Schema
}

func (o *UploadFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadFilesResponse) GetErrorMessage() *UploadFilesErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UploadFilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadFilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UploadFilesResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
