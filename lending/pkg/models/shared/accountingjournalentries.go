// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AccountingJournalEntries struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64                    `json:"pageSize"`
	Results  []AccountingJournalEntry `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *AccountingJournalEntries) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *AccountingJournalEntries) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *AccountingJournalEntries) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *AccountingJournalEntries) GetResults() []AccountingJournalEntry {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *AccountingJournalEntries) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
