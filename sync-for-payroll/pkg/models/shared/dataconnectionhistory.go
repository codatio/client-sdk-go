// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DataConnectionHistory - OK
type DataConnectionHistory struct {
	Links        Links           `json:"_links"`
	PageNumber   int64           `json:"pageNumber"`
	PageSize     int64           `json:"pageSize"`
	Results      []PullOperation `json:"results,omitempty"`
	TotalResults int64           `json:"totalResults"`
}

func (o *DataConnectionHistory) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *DataConnectionHistory) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *DataConnectionHistory) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *DataConnectionHistory) GetResults() []PullOperation {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *DataConnectionHistory) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}