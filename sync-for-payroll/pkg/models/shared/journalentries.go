// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// JournalEntries - Success
type JournalEntries struct {
	Links        Links          `json:"_links"`
	PageNumber   int64          `json:"pageNumber"`
	PageSize     int64          `json:"pageSize"`
	Results      []JournalEntry `json:"results,omitempty"`
	TotalResults int64          `json:"totalResults"`
}

func (o *JournalEntries) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *JournalEntries) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *JournalEntries) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *JournalEntries) GetResults() []JournalEntry {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *JournalEntries) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
