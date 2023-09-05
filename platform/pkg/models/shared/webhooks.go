// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Webhooks - OK
type Webhooks struct {
	Links        Links     `json:"_links"`
	PageNumber   int64     `json:"pageNumber"`
	PageSize     int64     `json:"pageSize"`
	Results      []Webhook `json:"results,omitempty"`
	TotalResults int64     `json:"totalResults"`
}

func (o *Webhooks) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Webhooks) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Webhooks) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Webhooks) GetResults() []Webhook {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Webhooks) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
