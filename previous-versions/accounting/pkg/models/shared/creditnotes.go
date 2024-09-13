// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreditNotes struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64        `json:"pageSize"`
	Results  []CreditNote `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *CreditNotes) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *CreditNotes) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *CreditNotes) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *CreditNotes) GetResults() []CreditNote {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *CreditNotes) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
