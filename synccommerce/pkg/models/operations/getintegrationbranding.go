package operations

import (
	"net/http"
)

type GetIntegrationBrandingPathParams struct {
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

type GetIntegrationBrandingRequest struct {
	PathParams GetIntegrationBrandingPathParams
}

type GetIntegrationBrandingResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
