// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Customers struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64      `json:"pageSize"`
	Results  []Customer `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
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
