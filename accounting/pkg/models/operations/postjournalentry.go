package operations

import (
	"net/http"
	"time"
)

type PostJournalEntryPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostJournalEntryQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostJournalEntrySourceModifiedDateJournalLinesAccountRef
// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
type PostJournalEntrySourceModifiedDateJournalLinesAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostJournalEntrySourceModifiedDateJournalLinesTracking
// List of record refs associated with the tracking information for the line (eg to a Tracking Category, or customer etc.)
type PostJournalEntrySourceModifiedDateJournalLinesTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type PostJournalEntrySourceModifiedDateJournalLines struct {
	AccountRef  *PostJournalEntrySourceModifiedDateJournalLinesAccountRef `json:"accountRef,omitempty"`
	Currency    *string                                                   `json:"currency,omitempty"`
	Description *string                                                   `json:"description,omitempty"`
	NetAmount   float64                                                   `json:"netAmount"`
	Tracking    *PostJournalEntrySourceModifiedDateJournalLinesTracking   `json:"tracking,omitempty"`
}

// PostJournalEntrySourceModifiedDateJournalRef
// Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).
type PostJournalEntrySourceModifiedDateJournalRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostJournalEntrySourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// PostJournalEntrySourceModifiedDateRecordRef
// Links to the underlying record or data type.
//
// Found on:
//
// - Journal entries
// - Account transactions
// - Invoices
// - Transfers
type PostJournalEntrySourceModifiedDateRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostJournalEntrySourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostJournalEntrySourceModifiedDate
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
type PostJournalEntrySourceModifiedDate struct {
	CreatedOn          *time.Time                                          `json:"createdOn,omitempty"`
	Description        *string                                             `json:"description,omitempty"`
	ID                 *string                                             `json:"id,omitempty"`
	JournalLines       []PostJournalEntrySourceModifiedDateJournalLines    `json:"journalLines,omitempty"`
	JournalRef         *PostJournalEntrySourceModifiedDateJournalRef       `json:"journalRef,omitempty"`
	Metadata           *PostJournalEntrySourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                          `json:"modifiedDate,omitempty"`
	PostedOn           *time.Time                                          `json:"postedOn,omitempty"`
	RecordRef          *PostJournalEntrySourceModifiedDateRecordRef        `json:"recordRef,omitempty"`
	SourceModifiedDate *time.Time                                          `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *PostJournalEntrySourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	UpdatedOn          *time.Time                                          `json:"updatedOn,omitempty"`
}

type PostJournalEntryRequest struct {
	PathParams  PostJournalEntryPathParams
	QueryParams PostJournalEntryQueryParams
	Request     *PostJournalEntrySourceModifiedDate `request:"mediaType=application/json"`
}

type PostJournalEntry200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostJournalEntry200ApplicationJSONChangesTypeEnum string

const (
	PostJournalEntry200ApplicationJSONChangesTypeEnumUnknown            PostJournalEntry200ApplicationJSONChangesTypeEnum = "Unknown"
	PostJournalEntry200ApplicationJSONChangesTypeEnumCreated            PostJournalEntry200ApplicationJSONChangesTypeEnum = "Created"
	PostJournalEntry200ApplicationJSONChangesTypeEnumModified           PostJournalEntry200ApplicationJSONChangesTypeEnum = "Modified"
	PostJournalEntry200ApplicationJSONChangesTypeEnumDeleted            PostJournalEntry200ApplicationJSONChangesTypeEnum = "Deleted"
	PostJournalEntry200ApplicationJSONChangesTypeEnumAttachmentUploaded PostJournalEntry200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostJournalEntry200ApplicationJSONChanges struct {
	AttachmentID *string                                                          `json:"attachmentId,omitempty"`
	RecordRef    *PostJournalEntry200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostJournalEntry200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesAccountRef
// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
type PostJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesTracking
// List of record refs associated with the tracking information for the line (eg to a Tracking Category, or customer etc.)
type PostJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type PostJournalEntry200ApplicationJSONSourceModifiedDateJournalLines struct {
	AccountRef  *PostJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesAccountRef `json:"accountRef,omitempty"`
	Currency    *string                                                                     `json:"currency,omitempty"`
	Description *string                                                                     `json:"description,omitempty"`
	NetAmount   float64                                                                     `json:"netAmount"`
	Tracking    *PostJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesTracking   `json:"tracking,omitempty"`
}

// PostJournalEntry200ApplicationJSONSourceModifiedDateJournalRef
// Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).
type PostJournalEntry200ApplicationJSONSourceModifiedDateJournalRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostJournalEntry200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// PostJournalEntry200ApplicationJSONSourceModifiedDateRecordRef
// Links to the underlying record or data type.
//
// Found on:
//
// - Journal entries
// - Account transactions
// - Invoices
// - Transfers
type PostJournalEntry200ApplicationJSONSourceModifiedDateRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostJournalEntry200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostJournalEntry200ApplicationJSONSourceModifiedDate
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
type PostJournalEntry200ApplicationJSONSourceModifiedDate struct {
	CreatedOn          *time.Time                                                            `json:"createdOn,omitempty"`
	Description        *string                                                               `json:"description,omitempty"`
	ID                 *string                                                               `json:"id,omitempty"`
	JournalLines       []PostJournalEntry200ApplicationJSONSourceModifiedDateJournalLines    `json:"journalLines,omitempty"`
	JournalRef         *PostJournalEntry200ApplicationJSONSourceModifiedDateJournalRef       `json:"journalRef,omitempty"`
	Metadata           *PostJournalEntry200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                            `json:"modifiedDate,omitempty"`
	PostedOn           *time.Time                                                            `json:"postedOn,omitempty"`
	RecordRef          *PostJournalEntry200ApplicationJSONSourceModifiedDateRecordRef        `json:"recordRef,omitempty"`
	SourceModifiedDate *time.Time                                                            `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *PostJournalEntry200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	UpdatedOn          *time.Time                                                            `json:"updatedOn,omitempty"`
}

type PostJournalEntry200ApplicationJSONStatusEnum string

const (
	PostJournalEntry200ApplicationJSONStatusEnumPending  PostJournalEntry200ApplicationJSONStatusEnum = "Pending"
	PostJournalEntry200ApplicationJSONStatusEnumFailed   PostJournalEntry200ApplicationJSONStatusEnum = "Failed"
	PostJournalEntry200ApplicationJSONStatusEnumSuccess  PostJournalEntry200ApplicationJSONStatusEnum = "Success"
	PostJournalEntry200ApplicationJSONStatusEnumTimedOut PostJournalEntry200ApplicationJSONStatusEnum = "TimedOut"
)

type PostJournalEntry200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostJournalEntry200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostJournalEntry200ApplicationJSONValidation struct {
	Errors   []PostJournalEntry200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostJournalEntry200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostJournalEntry200ApplicationJSON struct {
	Changes           []PostJournalEntry200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                `json:"companyId"`
	CompletedOnUtc    *time.Time                                            `json:"completedOnUtc,omitempty"`
	Data              *PostJournalEntry200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                `json:"dataConnectionKey"`
	DataType          *string                                               `json:"dataType,omitempty"`
	ErrorMessage      *string                                               `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                             `json:"requestedOnUtc"`
	Status            PostJournalEntry200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                   `json:"statusCode"`
	TimeoutInMinutes  *int                                                  `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                  `json:"timeoutInSeconds,omitempty"`
	Validation        *PostJournalEntry200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostJournalEntryResponse struct {
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
	PostJournalEntry200ApplicationJSONObject *PostJournalEntry200ApplicationJSON
}
