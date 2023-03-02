package operations

type GetIntegrationsPlatformKeyBrandingPathParams struct {
	PlatformKey string `pathParam:"style=simple,explode=false,name=platformKey"`
}

type GetIntegrationsPlatformKeyBrandingRequest struct {
	PathParams GetIntegrationsPlatformKeyBrandingPathParams
}

// GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImageImage
// Image reference.
type GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImageImage struct {
	Alt *string `json:"alt,omitempty"`
	Src *string `json:"src,omitempty"`
}

type GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImage struct {
	Image *GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImageImage `json:"image,omitempty"`
}

// GetIntegrationsPlatformKeyBrandingBrandingLogo
// Logo branding references.
type GetIntegrationsPlatformKeyBrandingBrandingLogo struct {
	Full   *GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImage `json:"full,omitempty"`
	Square *GetIntegrationsPlatformKeyBrandingBrandingLogoBrandingImage `json:"square,omitempty"`
}

type GetIntegrationsPlatformKeyBrandingBranding struct {
	Button   interface{}                                     `json:"button,omitempty"`
	Logo     *GetIntegrationsPlatformKeyBrandingBrandingLogo `json:"logo,omitempty"`
	SourceID *string                                         `json:"sourceId,omitempty"`
}

type GetIntegrationsPlatformKeyBrandingResponse struct {
	Branding    *GetIntegrationsPlatformKeyBrandingBranding
	ContentType string
	StatusCode  int
}
