// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// BrandingButton - Button branding references.
type BrandingButton struct {
	Default *BrandingImage `json:"default,omitempty"`
	Hover   *BrandingImage `json:"hover,omitempty"`
}

func (o *BrandingButton) GetDefault() *BrandingImage {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *BrandingButton) GetHover() *BrandingImage {
	if o == nil {
		return nil
	}
	return o.Hover
}
