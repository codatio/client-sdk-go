// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v2/pkg/models/shared"
	"net/http"
)

type GetAccountingProfileRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetAccountingProfileRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetAccountingProfileResponse struct {
	// Success
	CompanyInfo *shared.CompanyInfo
	// HTTP response content type for this operation
	ContentType string
	// Your API request was not properly authorized.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAccountingProfileResponse) GetCompanyInfo() *shared.CompanyInfo {
	if o == nil {
		return nil
	}
	return o.CompanyInfo
}

func (o *GetAccountingProfileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountingProfileResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *GetAccountingProfileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountingProfileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
