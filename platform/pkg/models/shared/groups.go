// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Groups struct {
	Results []Group `json:"results,omitempty"`
}

func (o *Groups) GetResults() []Group {
	if o == nil {
		return nil
	}
	return o.Results
}
