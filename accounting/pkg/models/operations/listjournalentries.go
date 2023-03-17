package operations

import (
	"net/http"
	"time"
)

type ListJournalEntriesRequest struct {
	CompanyID string  `pathParam:"style=simple,explode=false,name=companyId"`
	OrderBy   *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page      int     `queryParam:"style=form,explode=true,name=page"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query     *string `queryParam:"style=form,explode=true,name=query"`
}

type ListJournalEntriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListJournalEntriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListJournalEntriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListJournalEntriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListJournalEntriesLinksLinks struct {
	Current  ListJournalEntriesLinksLinksCurrent   `json:"current"`
	Next     *ListJournalEntriesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListJournalEntriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListJournalEntriesLinksLinksSelf      `json:"self"`
}

// ListJournalEntriesLinksSourceModifiedDateJournalLinesAccountRef
// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
type ListJournalEntriesLinksSourceModifiedDateJournalLinesAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListJournalEntriesLinksSourceModifiedDateJournalLinesTracking
// List of record refs associated with the tracking information for the line (eg to a Tracking Category, or customer etc.)
type ListJournalEntriesLinksSourceModifiedDateJournalLinesTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type ListJournalEntriesLinksSourceModifiedDateJournalLines struct {
	AccountRef  *ListJournalEntriesLinksSourceModifiedDateJournalLinesAccountRef `json:"accountRef,omitempty"`
	Currency    *string                                                          `json:"currency,omitempty"`
	Description *string                                                          `json:"description,omitempty"`
	NetAmount   float64                                                          `json:"netAmount"`
	Tracking    *ListJournalEntriesLinksSourceModifiedDateJournalLinesTracking   `json:"tracking,omitempty"`
}

// ListJournalEntriesLinksSourceModifiedDateJournalRef
// Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).
type ListJournalEntriesLinksSourceModifiedDateJournalRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type ListJournalEntriesLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// ListJournalEntriesLinksSourceModifiedDateRecordRef
// Links to the underlying record or data type.
//
// Found on:
//
// - Journal entries
// - Account transactions
// - Invoices
// - Transfers
type ListJournalEntriesLinksSourceModifiedDateRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

// ListJournalEntriesLinksSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type ListJournalEntriesLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// ListJournalEntriesLinksSourceModifiedDate
// > **Language tip:** For the top-level record of a company's financial transactions, refer to the [Journals](https://docs.codat.io/accounting-api#/schemas/Journal) data type
//
// > View the coverage for journal entries in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journalEntries" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A journal entry report shows the entries made in a company's general ledger, or [accounts](https://api.codat.io/swagger/index.html#/Accounts/get_companies__companyId__data_accounts), when transactions are approved. The journal line items for each journal entry should balance.
//
// A journal entry line item is a single transaction line on the journal entry. For example:
//
// - When a journal entry is recording a receipt of cash, the credit to accounts receivable and the debit to cash are separate line items.
// - When a company needs to recognise revenue from an annual contract on a monthly basis, on receipt of cash for month one, they make a debit to deferred income and a credit to revenue.
//
// In Codat a journal entry contains details of:
//
// - The date on which the entry was created and posted.
// - Itemised lines, including amounts and currency.
// - A reference to the associated accounts.
// - A reference to the underlying record. For example, the invoice, bill, or other data type that triggered the posting of the journal entry to the general ledger.
//
// > **Pushing journal entries **
// > Codat only supports journal entries in the base currency of the company that are pushed into accounts denominated in the same base currency.
type ListJournalEntriesLinksSourceModifiedDate struct {
	CreatedOn          *time.Time                                                 `json:"createdOn,omitempty"`
	Description        *string                                                    `json:"description,omitempty"`
	ID                 *string                                                    `json:"id,omitempty"`
	JournalLines       []ListJournalEntriesLinksSourceModifiedDateJournalLines    `json:"journalLines,omitempty"`
	JournalRef         *ListJournalEntriesLinksSourceModifiedDateJournalRef       `json:"journalRef,omitempty"`
	Metadata           *ListJournalEntriesLinksSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                 `json:"modifiedDate,omitempty"`
	PostedOn           *time.Time                                                 `json:"postedOn,omitempty"`
	RecordRef          *ListJournalEntriesLinksSourceModifiedDateRecordRef        `json:"recordRef,omitempty"`
	SourceModifiedDate *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *ListJournalEntriesLinksSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	UpdatedOn          *time.Time                                                 `json:"updatedOn,omitempty"`
}

// ListJournalEntriesLinks
// Codat's Paging Model
type ListJournalEntriesLinks struct {
	Links        ListJournalEntriesLinksLinks                `json:"_links"`
	PageNumber   int64                                       `json:"pageNumber"`
	PageSize     int64                                       `json:"pageSize"`
	Results      []ListJournalEntriesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                       `json:"totalResults"`
}

type ListJournalEntriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListJournalEntriesLinks
}
