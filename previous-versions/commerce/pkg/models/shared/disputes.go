// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Disputes struct {
	Links        Links     `json:"_links"`
	PageNumber   int64     `json:"pageNumber"`
	PageSize     int64     `json:"pageSize"`
	Results      []Dispute `json:"results,omitempty"`
	TotalResults int64     `json:"totalResults"`
}

func (o *Disputes) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Disputes) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Disputes) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Disputes) GetResults() []Dispute {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Disputes) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
