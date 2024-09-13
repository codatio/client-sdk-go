// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	"net/http"
)

type GetConnectionManagementCorsSettingsResponse struct {
	// Success
	ConnectionManagementAllowedOrigins *shared.ConnectionManagementAllowedOrigins
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetConnectionManagementCorsSettingsResponse) GetConnectionManagementAllowedOrigins() *shared.ConnectionManagementAllowedOrigins {
	if o == nil {
		return nil
	}
	return o.ConnectionManagementAllowedOrigins
}

func (o *GetConnectionManagementCorsSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetConnectionManagementCorsSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectionManagementCorsSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
