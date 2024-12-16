// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/bank-feeds/v8/pkg/models/shared"
	"net/http"
)

type GetCompanyAccessTokenRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetCompanyAccessTokenRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetCompanyAccessTokenResponse struct {
	// OK
	CompanyAccessToken *shared.CompanyAccessToken
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetCompanyAccessTokenResponse) GetCompanyAccessToken() *shared.CompanyAccessToken {
	if o == nil {
		return nil
	}
	return o.CompanyAccessToken
}

func (o *GetCompanyAccessTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCompanyAccessTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCompanyAccessTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
