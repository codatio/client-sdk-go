package operations

import (
	"net/http"
	"time"
)

type GetJournalEntryRequest struct {
	CompanyID      string `pathParam:"style=simple,explode=false,name=companyId"`
	JournalEntryID string `pathParam:"style=simple,explode=false,name=journalEntryId"`
}

// GetJournalEntrySourceModifiedDateJournalLinesAccountRef
// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
type GetJournalEntrySourceModifiedDateJournalLinesAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetJournalEntrySourceModifiedDateJournalLinesTracking
// List of record refs associated with the tracking information for the line (eg to a Tracking Category, or customer etc.)
type GetJournalEntrySourceModifiedDateJournalLinesTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type GetJournalEntrySourceModifiedDateJournalLines struct {
	AccountRef  *GetJournalEntrySourceModifiedDateJournalLinesAccountRef `json:"accountRef,omitempty"`
	Currency    *string                                                  `json:"currency,omitempty"`
	Description *string                                                  `json:"description,omitempty"`
	NetAmount   float64                                                  `json:"netAmount"`
	Tracking    *GetJournalEntrySourceModifiedDateJournalLinesTracking   `json:"tracking,omitempty"`
}

// GetJournalEntrySourceModifiedDateJournalRef
// Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).
type GetJournalEntrySourceModifiedDateJournalRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type GetJournalEntrySourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// GetJournalEntrySourceModifiedDateRecordRef
// Links to the underlying record or data type.
//
// Found on:
//
// - Journal entries
// - Account transactions
// - Invoices
// - Transfers
type GetJournalEntrySourceModifiedDateRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

// GetJournalEntrySourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetJournalEntrySourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetJournalEntrySourceModifiedDate
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
type GetJournalEntrySourceModifiedDate struct {
	CreatedOn          *time.Time                                         `json:"createdOn,omitempty"`
	Description        *string                                            `json:"description,omitempty"`
	ID                 *string                                            `json:"id,omitempty"`
	JournalLines       []GetJournalEntrySourceModifiedDateJournalLines    `json:"journalLines,omitempty"`
	JournalRef         *GetJournalEntrySourceModifiedDateJournalRef       `json:"journalRef,omitempty"`
	Metadata           *GetJournalEntrySourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                         `json:"modifiedDate,omitempty"`
	PostedOn           *time.Time                                         `json:"postedOn,omitempty"`
	RecordRef          *GetJournalEntrySourceModifiedDateRecordRef        `json:"recordRef,omitempty"`
	SourceModifiedDate *time.Time                                         `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *GetJournalEntrySourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	UpdatedOn          *time.Time                                         `json:"updatedOn,omitempty"`
}

type GetJournalEntryResponse struct {
	ContentType        string
	SourceModifiedDate *GetJournalEntrySourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
