// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package webhooks

import (
	"net/http"
)

type DataSyncCompletedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *DataSyncCompletedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DataSyncCompletedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DataSyncCompletedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
