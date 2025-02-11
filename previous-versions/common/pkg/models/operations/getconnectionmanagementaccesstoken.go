// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/common/pkg/models/shared"
	"net/http"
)

type GetConnectionManagementAccessTokenRequest struct {
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

func (o *GetConnectionManagementAccessTokenRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

type GetConnectionManagementAccessTokenResponse struct {
	// Success
	ConnectionManagementAccessToken *shared.ConnectionManagementAccessToken
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetConnectionManagementAccessTokenResponse) GetConnectionManagementAccessToken() *shared.ConnectionManagementAccessToken {
	if o == nil {
		return nil
	}
	return o.ConnectionManagementAccessToken
}

func (o *GetConnectionManagementAccessTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectionManagementAccessTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectionManagementAccessTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
