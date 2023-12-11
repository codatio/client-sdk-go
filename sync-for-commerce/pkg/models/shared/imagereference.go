// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ImageReference - Image reference.
type ImageReference struct {
	// Alternative text when image is not available.
	Alt *string `json:"alt,omitempty"`
	// Source URL for image.
	Src *string `json:"src,omitempty"`
}

func (o *ImageReference) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *ImageReference) GetSrc() *string {
	if o == nil {
		return nil
	}
	return o.Src
}
