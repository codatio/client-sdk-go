package operations

import (
	"net/http"
)

type GetIntegrationBrandingRequest struct {
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

type GetIntegrationBrandingResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
