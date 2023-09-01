// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProductCategories - OK
type ProductCategories struct {
	Links        Links             `json:"_links"`
	PageNumber   int64             `json:"pageNumber"`
	PageSize     int64             `json:"pageSize"`
	Results      []ProductCategory `json:"results,omitempty"`
	TotalResults int64             `json:"totalResults"`
}

func (o *ProductCategories) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *ProductCategories) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *ProductCategories) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *ProductCategories) GetResults() []ProductCategory {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *ProductCategories) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
