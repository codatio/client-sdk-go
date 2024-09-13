// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DataIntegrityConnectionID struct {
	// An array of strings. The connection IDs for the type specified in the url.
	Source []string `json:"source,omitempty"`
	// An array of strings. The connection IDs for the type being matched to.
	Target []string `json:"target,omitempty"`
}

func (o *DataIntegrityConnectionID) GetSource() []string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *DataIntegrityConnectionID) GetTarget() []string {
	if o == nil {
		return nil
	}
	return o.Target
}
