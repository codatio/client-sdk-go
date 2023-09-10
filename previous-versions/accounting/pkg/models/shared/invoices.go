// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Invoices struct {
	Links        Links     `json:"_links"`
	PageNumber   int64     `json:"pageNumber"`
	PageSize     int64     `json:"pageSize"`
	Results      []Invoice `json:"results,omitempty"`
	TotalResults int64     `json:"totalResults"`
}

func (o *Invoices) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Invoices) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Invoices) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Invoices) GetResults() []Invoice {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Invoices) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
