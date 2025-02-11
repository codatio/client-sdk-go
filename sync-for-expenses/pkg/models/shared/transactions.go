// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Transactions struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64               `json:"pageSize"`
	Results  []SchemaTransaction `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *Transactions) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Transactions) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Transactions) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Transactions) GetResults() []SchemaTransaction {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Transactions) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
