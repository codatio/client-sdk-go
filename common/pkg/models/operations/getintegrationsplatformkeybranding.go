package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetIntegrationsPlatformKeyBrandingPathParams struct {
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

type GetIntegrationsPlatformKeyBrandingRequest struct {
	PathParams GetIntegrationsPlatformKeyBrandingPathParams
}

type GetIntegrationsPlatformKeyBrandingResponse struct {
	Branding    *shared.Branding
	ContentType string
	StatusCode  int
}
