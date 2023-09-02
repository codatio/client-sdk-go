// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Branding - OK
type Branding struct {
	// Button branding references.
	Button *BrandingButton `json:"button,omitempty"`
	// Logo branding references.
	Logo *BrandingLogo `json:"logo,omitempty"`
	// A source-specific ID used to distinguish between different sources originating from the same data connection. In general, a data connection is a single data source. However, for TrueLayer, `sourceId` is associated with a specific bank and has a many-to-one relationship with the `integrationId`.
	SourceID *string `json:"sourceId,omitempty"`
}

func (o *Branding) GetButton() *BrandingButton {
	if o == nil {
		return nil
	}
	return o.Button
}

func (o *Branding) GetLogo() *BrandingLogo {
	if o == nil {
		return nil
	}
	return o.Logo
}

func (o *Branding) GetSourceID() *string {
	if o == nil {
		return nil
	}
	return o.SourceID
}
