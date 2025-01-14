// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Journals struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64     `json:"pageSize"`
	Results  []Journal `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *Journals) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Journals) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Journals) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Journals) GetResults() []Journal {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Journals) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
