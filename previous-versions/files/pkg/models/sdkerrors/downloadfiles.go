// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// DownloadFilesErrorMessage - You are using an outdated API key or a key not associated with that resource.
type DownloadFilesErrorMessage struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	// `True` if the error occurred transiently and can be retried.
	CanBeRetried *string `json:"canBeRetried,omitempty"`
	// Unique identifier used to propagate to all downstream services and determine the source of the error.
	CorrelationID *string `json:"correlationId,omitempty"`
	// Machine readable error code used to automate processes based on the code returned.
	DetailedErrorCode *int64 `json:"detailedErrorCode,omitempty"`
	// A brief description of the error.
	Error_ *string `json:"error,omitempty"`
	// Codat's service the returned the error.
	Service *string `json:"service,omitempty"`
	// The HTTP status code returned by the error.
	StatusCode *int64 `json:"statusCode,omitempty"`
}

var _ error = &DownloadFilesErrorMessage{}

func (e *DownloadFilesErrorMessage) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}