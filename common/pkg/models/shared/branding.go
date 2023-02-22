package shared

// BrandingLogoBrandingImageImage
// Image reference.
type BrandingLogoBrandingImageImage struct {
	Alt *string `json:"alt,omitempty"`
	Src *string `json:"src,omitempty"`
}

type BrandingLogoBrandingImage struct {
	Image *BrandingLogoBrandingImageImage `json:"image,omitempty"`
}

// BrandingLogo
// Logo branding references.
type BrandingLogo struct {
	Full   *BrandingLogoBrandingImage `json:"full,omitempty"`
	Square *Full                      `json:"square,omitempty"`
}

type Branding struct {
	Button   *interface{}  `json:"button,omitempty"`
	Logo     *BrandingLogo `json:"logo,omitempty"`
	SourceID *string       `json:"sourceId,omitempty"`
}
