// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/platform/v3/pkg/models/shared"
	"net/http"
)

type GetIntegrationsBrandingRequest struct {
	// A unique 4-letter key to represent a platform in each integration.
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

func (o *GetIntegrationsBrandingRequest) GetPlatformKey() string {
	if o == nil {
		return ""
	}
	return o.PlatformKey
}

type GetIntegrationsBrandingResponse struct {
	// OK
	Branding *shared.Branding
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetIntegrationsBrandingResponse) GetBranding() *shared.Branding {
	if o == nil {
		return nil
	}
	return o.Branding
}

func (o *GetIntegrationsBrandingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIntegrationsBrandingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIntegrationsBrandingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
