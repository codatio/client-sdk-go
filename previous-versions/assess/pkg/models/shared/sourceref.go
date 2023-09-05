// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SourceRef - A source reference containing the `sourceType` object "Banking".
type SourceRef struct {
	// The data source type.
	SourceType *string `json:"sourceType,omitempty"`
}

func (o *SourceRef) GetSourceType() *string {
	if o == nil {
		return nil
	}
	return o.SourceType
}