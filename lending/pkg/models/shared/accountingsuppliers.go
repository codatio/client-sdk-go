// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountingSuppliers struct {
	Links        Links                `json:"_links"`
	PageNumber   int64                `json:"pageNumber"`
	PageSize     int64                `json:"pageSize"`
	Results      []AccountingSupplier `json:"results,omitempty"`
	TotalResults int64                `json:"totalResults"`
}

func (o *AccountingSuppliers) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *AccountingSuppliers) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *AccountingSuppliers) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *AccountingSuppliers) GetResults() []AccountingSupplier {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *AccountingSuppliers) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
