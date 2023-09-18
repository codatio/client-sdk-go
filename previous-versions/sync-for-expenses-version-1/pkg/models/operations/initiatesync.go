// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-expenses-version-1/pkg/models/shared"
	"net/http"
)

type InitiateSyncRequest struct {
	PostSync  *shared.PostSync `request:"mediaType=application/json"`
	CompanyID string           `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *InitiateSyncRequest) GetPostSync() *shared.PostSync {
	if o == nil {
		return nil
	}
	return o.PostSync
}

func (o *InitiateSyncRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type InitiateSyncResponse struct {
	ContentType string
	// If model is incorrect
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
	// Returns the newly created SyncId
	SyncInitiated *shared.SyncInitiated
}

func (o *InitiateSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InitiateSyncResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
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