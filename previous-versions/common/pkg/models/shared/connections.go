// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Connections struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64        `json:"pageSize"`
	Results  []Connection `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *Connections) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *Connections) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *Connections) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *Connections) GetResults() []Connection {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *Connections) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
