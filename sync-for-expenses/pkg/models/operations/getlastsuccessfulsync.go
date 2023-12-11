// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/models/shared"
	"net/http"
)

type GetLastSuccessfulSyncRequest struct {
	// Unique identifier for a company.
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
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
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
