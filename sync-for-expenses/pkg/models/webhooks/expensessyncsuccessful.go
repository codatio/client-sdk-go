// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package webhooks

import (
	"net/http"
)

type ExpensesSyncSuccessfulResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ExpensesSyncSuccessfulResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ExpensesSyncSuccessfulResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ExpensesSyncSuccessfulResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
