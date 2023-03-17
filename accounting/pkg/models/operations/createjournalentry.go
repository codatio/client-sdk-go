package operations

import (
	"net/http"
	"time"
)

// CreateJournalEntrySourceModifiedDateJournalLinesAccountRef
// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
type CreateJournalEntrySourceModifiedDateJournalLinesAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateJournalEntrySourceModifiedDateJournalLinesTracking
// List of record refs associated with the tracking information for the line (eg to a Tracking Category, or customer etc.)
type CreateJournalEntrySourceModifiedDateJournalLinesTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type CreateJournalEntrySourceModifiedDateJournalLines struct {
	AccountRef  *CreateJournalEntrySourceModifiedDateJournalLinesAccountRef `json:"accountRef,omitempty"`
	Currency    *string                                                     `json:"currency,omitempty"`
	Description *string                                                     `json:"description,omitempty"`
	NetAmount   float64                                                     `json:"netAmount"`
	Tracking    *CreateJournalEntrySourceModifiedDateJournalLinesTracking   `json:"tracking,omitempty"`
}

// CreateJournalEntrySourceModifiedDateJournalRef
// Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).
type CreateJournalEntrySourceModifiedDateJournalRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateJournalEntrySourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateJournalEntrySourceModifiedDateRecordRef
// Links to the underlying record or data type.
//
// Found on:
//
// - Journal entries
// - Account transactions
// - Invoices
// - Transfers
type CreateJournalEntrySourceModifiedDateRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

// CreateJournalEntrySourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateJournalEntrySourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateJournalEntrySourceModifiedDate
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
type CreateJournalEntrySourceModifiedDate struct {
	CreatedOn          *time.Time                                            `json:"createdOn,omitempty"`
	Description        *string                                               `json:"description,omitempty"`
	ID                 *string                                               `json:"id,omitempty"`
	JournalLines       []CreateJournalEntrySourceModifiedDateJournalLines    `json:"journalLines,omitempty"`
	JournalRef         *CreateJournalEntrySourceModifiedDateJournalRef       `json:"journalRef,omitempty"`
	Metadata           *CreateJournalEntrySourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                            `json:"modifiedDate,omitempty"`
	PostedOn           *time.Time                                            `json:"postedOn,omitempty"`
	RecordRef          *CreateJournalEntrySourceModifiedDateRecordRef        `json:"recordRef,omitempty"`
	SourceModifiedDate *time.Time                                            `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *CreateJournalEntrySourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	UpdatedOn          *time.Time                                            `json:"updatedOn,omitempty"`
}

type CreateJournalEntryRequest struct {
	RequestBody      *CreateJournalEntrySourceModifiedDate `request:"mediaType=application/json"`
	CompanyID        string                                `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                                `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                                  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateJournalEntry200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateJournalEntry200ApplicationJSONChangesTypeEnum string

const (
	CreateJournalEntry200ApplicationJSONChangesTypeEnumUnknown            CreateJournalEntry200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateJournalEntry200ApplicationJSONChangesTypeEnumCreated            CreateJournalEntry200ApplicationJSONChangesTypeEnum = "Created"
	CreateJournalEntry200ApplicationJSONChangesTypeEnumModified           CreateJournalEntry200ApplicationJSONChangesTypeEnum = "Modified"
	CreateJournalEntry200ApplicationJSONChangesTypeEnumDeleted            CreateJournalEntry200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateJournalEntry200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateJournalEntry200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateJournalEntry200ApplicationJSONChanges struct {
	AttachmentID *string                                                            `json:"attachmentId,omitempty"`
	RecordRef    *CreateJournalEntry200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateJournalEntry200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesAccountRef
// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
type CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesTracking
// List of record refs associated with the tracking information for the line (eg to a Tracking Category, or customer etc.)
type CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesTracking struct {
	RecordRefs []string `json:"recordRefs,omitempty"`
}

type CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalLines struct {
	AccountRef  *CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesAccountRef `json:"accountRef,omitempty"`
	Currency    *string                                                                       `json:"currency,omitempty"`
	Description *string                                                                       `json:"description,omitempty"`
	NetAmount   float64                                                                       `json:"netAmount"`
	Tracking    *CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalLinesTracking   `json:"tracking,omitempty"`
}

// CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalRef
// Links journal entries to the relevant journal in accounting integrations that use multi-book accounting (multiple journals).
type CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreateJournalEntry200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateJournalEntry200ApplicationJSONSourceModifiedDateRecordRef
// Links to the underlying record or data type.
//
// Found on:
//
// - Journal entries
// - Account transactions
// - Invoices
// - Transfers
type CreateJournalEntry200ApplicationJSONSourceModifiedDateRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

// CreateJournalEntry200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateJournalEntry200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateJournalEntry200ApplicationJSONSourceModifiedDate
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
type CreateJournalEntry200ApplicationJSONSourceModifiedDate struct {
	CreatedOn          *time.Time                                                              `json:"createdOn,omitempty"`
	Description        *string                                                                 `json:"description,omitempty"`
	ID                 *string                                                                 `json:"id,omitempty"`
	JournalLines       []CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalLines    `json:"journalLines,omitempty"`
	JournalRef         *CreateJournalEntry200ApplicationJSONSourceModifiedDateJournalRef       `json:"journalRef,omitempty"`
	Metadata           *CreateJournalEntry200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                              `json:"modifiedDate,omitempty"`
	PostedOn           *time.Time                                                              `json:"postedOn,omitempty"`
	RecordRef          *CreateJournalEntry200ApplicationJSONSourceModifiedDateRecordRef        `json:"recordRef,omitempty"`
	SourceModifiedDate *time.Time                                                              `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *CreateJournalEntry200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	UpdatedOn          *time.Time                                                              `json:"updatedOn,omitempty"`
}

type CreateJournalEntry200ApplicationJSONStatusEnum string

const (
	CreateJournalEntry200ApplicationJSONStatusEnumPending  CreateJournalEntry200ApplicationJSONStatusEnum = "Pending"
	CreateJournalEntry200ApplicationJSONStatusEnumFailed   CreateJournalEntry200ApplicationJSONStatusEnum = "Failed"
	CreateJournalEntry200ApplicationJSONStatusEnumSuccess  CreateJournalEntry200ApplicationJSONStatusEnum = "Success"
	CreateJournalEntry200ApplicationJSONStatusEnumTimedOut CreateJournalEntry200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateJournalEntry200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateJournalEntry200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateJournalEntry200ApplicationJSONValidation struct {
	Errors   []CreateJournalEntry200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateJournalEntry200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateJournalEntry200ApplicationJSON struct {
	Changes           []CreateJournalEntry200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                  `json:"companyId"`
	CompletedOnUtc    *time.Time                                              `json:"completedOnUtc,omitempty"`
	Data              *CreateJournalEntry200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                  `json:"dataConnectionKey"`
	DataType          *string                                                 `json:"dataType,omitempty"`
	ErrorMessage      *string                                                 `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                  `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                               `json:"requestedOnUtc"`
	Status            CreateJournalEntry200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                     `json:"statusCode"`
	TimeoutInMinutes  *int                                                    `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                    `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateJournalEntry200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateJournalEntryResponse struct {
	ContentType                                string
	StatusCode                                 int
	RawResponse                                *http.Response
	CreateJournalEntry200ApplicationJSONObject *CreateJournalEntry200ApplicationJSON
}
