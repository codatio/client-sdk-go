// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DataIntegritySummaries struct {
	Summaries []DataIntegritySummary `json:"summaries,omitempty"`
}

func (o *DataIntegritySummaries) GetSummaries() []DataIntegritySummary {
	if o == nil {
		return nil
	}
	return o.Summaries
}
