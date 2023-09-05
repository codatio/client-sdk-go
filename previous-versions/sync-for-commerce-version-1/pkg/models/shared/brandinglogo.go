// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BrandingLogo - Logo branding references.
type BrandingLogo struct {
	Full   *BrandingImage `json:"full,omitempty"`
	Square *BrandingImage `json:"square,omitempty"`
}

func (o *BrandingLogo) GetFull() *BrandingImage {
	if o == nil {
		return nil
	}
	return o.Full
}

func (o *BrandingLogo) GetSquare() *BrandingImage {
	if o == nil {
		return nil
	}
	return o.Square
}