// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Summaries struct {
	Summaries []DataIntegritySummary `json:"summaries,omitempty"`
}

func (o *Summaries) GetSummaries() []DataIntegritySummary {
	if o == nil {
		return nil
	}
	return o.Summaries
}
