// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountingBillCreditNotes struct {
	Links        Links                      `json:"_links"`
	PageNumber   int64                      `json:"pageNumber"`
	PageSize     int64                      `json:"pageSize"`
	Results      []AccountingBillCreditNote `json:"results,omitempty"`
	TotalResults int64                      `json:"totalResults"`
}

func (o *AccountingBillCreditNotes) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *AccountingBillCreditNotes) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *AccountingBillCreditNotes) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *AccountingBillCreditNotes) GetResults() []AccountingBillCreditNote {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *AccountingBillCreditNotes) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
