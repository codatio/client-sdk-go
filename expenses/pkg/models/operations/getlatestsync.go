// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/expenses/pkg/models/shared"
	"net/http"
)

type GetLatestSyncRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetLatestSyncRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetLatestSyncResponse struct {
	// Success
	CompanySyncStatus *shared.CompanySyncStatus
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
	// Your API request was not properly authorized.
	Schema *shared.Schema
}

func (o *GetLatestSyncResponse) GetCompanySyncStatus() *shared.CompanySyncStatus {
	if o == nil {
		return nil
	}
	return o.CompanySyncStatus
}

func (o *GetLatestSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLatestSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLatestSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLatestSyncResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
