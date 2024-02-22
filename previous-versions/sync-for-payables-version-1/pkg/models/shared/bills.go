// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Bills struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64  `json:"pageSize"`
	Results  []Bill `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *Bills) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Bills) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Bills) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Bills) GetResults() []Bill {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Bills) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}