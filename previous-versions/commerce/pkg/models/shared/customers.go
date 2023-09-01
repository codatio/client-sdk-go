// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Customers - OK
type Customers struct {
	Links        Links      `json:"_links"`
	PageNumber   int64      `json:"pageNumber"`
	PageSize     int64      `json:"pageSize"`
	Results      []Customer `json:"results,omitempty"`
	TotalResults int64      `json:"totalResults"`
}

func (o *Customers) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Customers) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Customers) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Customers) GetResults() []Customer {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Customers) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
