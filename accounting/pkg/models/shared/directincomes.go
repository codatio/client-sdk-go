// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DirectIncomes - Success
type DirectIncomes struct {
	Links        Links          `json:"_links"`
	PageNumber   int64          `json:"pageNumber"`
	PageSize     int64          `json:"pageSize"`
	Results      []DirectIncome `json:"results,omitempty"`
	TotalResults int64          `json:"totalResults"`
}

func (o *DirectIncomes) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *DirectIncomes) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *DirectIncomes) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *DirectIncomes) GetResults() []DirectIncome {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *DirectIncomes) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
