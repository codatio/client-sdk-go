// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-commerce/v2/pkg/models/shared"
	"net/http"
)

type GetSyncStatusRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetSyncStatusRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetSyncStatusResponse struct {
	// Bad Request
	BadRequest  interface{}
	ContentType string
	// Not Found
	NotFound    interface{}
	StatusCode  int
	RawResponse *http.Response
	// Success
	SyncStatus *shared.SyncStatus
}

func (o *GetSyncStatusResponse) GetBadRequest() interface{} {
	if o == nil {
		return nil
	}
	return o.BadRequest
}

func (o *GetSyncStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSyncStatusResponse) GetNotFound() interface{} {
	if o == nil {
		return nil
	}
	return o.NotFound
}

func (o *GetSyncStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSyncStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSyncStatusResponse) GetSyncStatus() *shared.SyncStatus {
	if o == nil {
		return nil
	}
	return o.SyncStatus
}
