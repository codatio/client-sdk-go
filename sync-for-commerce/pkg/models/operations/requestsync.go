// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"net/http"
)

type RequestSyncRequest struct {
	SyncToLatestArgs *shared.SyncToLatestArgs `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *RequestSyncRequest) GetSyncToLatestArgs() *shared.SyncToLatestArgs {
	if o == nil {
		return nil
	}
	return o.SyncToLatestArgs
}

func (o *RequestSyncRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type RequestSyncResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	SyncSummary *shared.SyncSummary
}

func (o *RequestSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RequestSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RequestSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RequestSyncResponse) GetSyncSummary() *shared.SyncSummary {
	if o == nil {
		return nil
	}
	return o.SyncSummary
}
