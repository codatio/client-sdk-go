// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	"net/http"
)

type InitiateSyncRequest struct {
	InitiateSync *shared.InitiateSync `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *InitiateSyncRequest) GetInitiateSync() *shared.InitiateSync {
	if o == nil {
		return nil
	}
	return o.InitiateSync
}

func (o *InitiateSyncRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type InitiateSyncResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the newly created syncId
	SyncInitiated *shared.SyncInitiated
}

func (o *InitiateSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InitiateSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InitiateSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *InitiateSyncResponse) GetSyncInitiated() *shared.SyncInitiated {
	if o == nil {
		return nil
	}
	return o.SyncInitiated
}
