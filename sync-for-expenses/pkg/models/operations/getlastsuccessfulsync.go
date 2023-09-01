// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/pkg/models/shared"
	"net/http"
)

type GetLastSuccessfulSyncRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetLastSuccessfulSyncRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetLastSuccessfulSyncResponse struct {
	// Success
	CompanySyncStatus *shared.CompanySyncStatus
	ContentType       string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetLastSuccessfulSyncResponse) GetCompanySyncStatus() *shared.CompanySyncStatus {
	if o == nil {
		return nil
	}
	return o.CompanySyncStatus
}

func (o *GetLastSuccessfulSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLastSuccessfulSyncResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetLastSuccessfulSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLastSuccessfulSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
