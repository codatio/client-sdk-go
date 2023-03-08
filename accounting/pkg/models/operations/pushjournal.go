package operations

import (
	"net/http"
	"time"
)

type PushJournalPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PushJournalQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PushJournalSourceModifiedDateStatusEnum string

const (
	PushJournalSourceModifiedDateStatusEnumUnknown  PushJournalSourceModifiedDateStatusEnum = "Unknown"
	PushJournalSourceModifiedDateStatusEnumActive   PushJournalSourceModifiedDateStatusEnum = "Active"
	PushJournalSourceModifiedDateStatusEnumArchived PushJournalSourceModifiedDateStatusEnum = "Archived"
)

// PushJournalSourceModifiedDateInput
// > **Language tip:** For line items, or individual transactions, of a company's financial documents, refer to the [Journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) data type
//
// > View the coverage for journals in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// In accounting software, journals are used to record all the financial transactions of a company. Each transaction in a journal is represented by a separate [journal entry](https://docs.codat.io/accounting-api#/schemas/JournalEntry). These entries are used to create the general ledger, which is then used to create the financial statements of a business.
//
// When a company records all their transactions in a single journal, it can become large and difficult to maintain and track. This is why large companies often use multiple journals (also known as subjournals) to categorize and manage journal entries.
//
// Such journals can be divided into two categories:
//
// - Special journals: journals used to record specific types of transactions; for example, a purchases journal, a sales journal, or a cash management journal.
// - General journals: journals used to record transactions that fall outside the scope of the special journals.
//
// Multiple journals or subjournals are used in the following Codat integrations:
//
// - [Sage Intacct](https://docs.codat.io/integrations/accounting/sage-intacct/accounting-sage-intacct)  (mandatory)
// - [Exact Online](https://docs.codat.io/integrations/accounting/exact-online/accounting-exact-online)  (mandatory)
// - [Oracle NetSuite](https://docs.codat.io/integrations/accounting/netsuite/accounting-netsuite) (optional)
//
// > When pushing journal entries to an accounting platform that doesn’t support multiple journals (multi-book accounting), the entries will be linked to the platform-generic journal. The Journals data type will only include one object.
type PushJournalSourceModifiedDateInput struct {
	CreatedOn          *time.Time                               `json:"createdOn,omitempty"`
	HasChildren        *bool                                    `json:"hasChildren,omitempty"`
	ID                 *string                                  `json:"id,omitempty"`
	JournalCode        *string                                  `json:"journalCode,omitempty"`
	ModifiedDate       *time.Time                               `json:"modifiedDate,omitempty"`
	Name               *string                                  `json:"name,omitempty"`
	ParentID           *string                                  `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                               `json:"sourceModifiedDate,omitempty"`
	Status             *PushJournalSourceModifiedDateStatusEnum `json:"status,omitempty"`
	Type               *string                                  `json:"type,omitempty"`
}

type PushJournalRequest struct {
	PathParams  PushJournalPathParams
	QueryParams PushJournalQueryParams
	Request     *PushJournalSourceModifiedDateInput `request:"mediaType=application/json"`
}

type PushJournal200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PushJournal200ApplicationJSONChangesTypeEnum string

const (
	PushJournal200ApplicationJSONChangesTypeEnumUnknown            PushJournal200ApplicationJSONChangesTypeEnum = "Unknown"
	PushJournal200ApplicationJSONChangesTypeEnumCreated            PushJournal200ApplicationJSONChangesTypeEnum = "Created"
	PushJournal200ApplicationJSONChangesTypeEnumModified           PushJournal200ApplicationJSONChangesTypeEnum = "Modified"
	PushJournal200ApplicationJSONChangesTypeEnumDeleted            PushJournal200ApplicationJSONChangesTypeEnum = "Deleted"
	PushJournal200ApplicationJSONChangesTypeEnumAttachmentUploaded PushJournal200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PushJournal200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PushJournal200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PushJournal200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PushJournal200ApplicationJSONSourceModifiedDateMetadataMetadata
// Additional information about the entity
type PushJournal200ApplicationJSONSourceModifiedDateMetadataMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PushJournal200ApplicationJSONSourceModifiedDateMetadata struct {
	Metadata *PushJournal200ApplicationJSONSourceModifiedDateMetadataMetadata `json:"metadata,omitempty"`
}

type PushJournal200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PushJournal200ApplicationJSONSourceModifiedDateStatusEnumUnknown  PushJournal200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PushJournal200ApplicationJSONSourceModifiedDateStatusEnumActive   PushJournal200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	PushJournal200ApplicationJSONSourceModifiedDateStatusEnumArchived PushJournal200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

// PushJournal200ApplicationJSONSourceModifiedDate
// > **Language tip:** For line items, or individual transactions, of a company's financial documents, refer to the [Journal entries](https://docs.codat.io/accounting-api#/schemas/JournalEntry) data type
//
// > View the coverage for journals in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=journals" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// In accounting software, journals are used to record all the financial transactions of a company. Each transaction in a journal is represented by a separate [journal entry](https://docs.codat.io/accounting-api#/schemas/JournalEntry). These entries are used to create the general ledger, which is then used to create the financial statements of a business.
//
// When a company records all their transactions in a single journal, it can become large and difficult to maintain and track. This is why large companies often use multiple journals (also known as subjournals) to categorize and manage journal entries.
//
// Such journals can be divided into two categories:
//
// - Special journals: journals used to record specific types of transactions; for example, a purchases journal, a sales journal, or a cash management journal.
// - General journals: journals used to record transactions that fall outside the scope of the special journals.
//
// Multiple journals or subjournals are used in the following Codat integrations:
//
// - [Sage Intacct](https://docs.codat.io/integrations/accounting/sage-intacct/accounting-sage-intacct)  (mandatory)
// - [Exact Online](https://docs.codat.io/integrations/accounting/exact-online/accounting-exact-online)  (mandatory)
// - [Oracle NetSuite](https://docs.codat.io/integrations/accounting/netsuite/accounting-netsuite) (optional)
//
// > When pushing journal entries to an accounting platform that doesn’t support multiple journals (multi-book accounting), the entries will be linked to the platform-generic journal. The Journals data type will only include one object.
type PushJournal200ApplicationJSONSourceModifiedDate struct {
	CreatedOn          *time.Time                                                 `json:"createdOn,omitempty"`
	HasChildren        *bool                                                      `json:"hasChildren,omitempty"`
	ID                 *string                                                    `json:"id,omitempty"`
	JournalCode        *string                                                    `json:"journalCode,omitempty"`
	Metadata           *PushJournal200ApplicationJSONSourceModifiedDateMetadata   `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                 `json:"modifiedDate,omitempty"`
	Name               *string                                                    `json:"name,omitempty"`
	ParentID           *string                                                    `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	Status             *PushJournal200ApplicationJSONSourceModifiedDateStatusEnum `json:"status,omitempty"`
	Type               *string                                                    `json:"type,omitempty"`
}

type PushJournal200ApplicationJSONStatusEnum string

const (
	PushJournal200ApplicationJSONStatusEnumPending  PushJournal200ApplicationJSONStatusEnum = "Pending"
	PushJournal200ApplicationJSONStatusEnumFailed   PushJournal200ApplicationJSONStatusEnum = "Failed"
	PushJournal200ApplicationJSONStatusEnumSuccess  PushJournal200ApplicationJSONStatusEnum = "Success"
	PushJournal200ApplicationJSONStatusEnumTimedOut PushJournal200ApplicationJSONStatusEnum = "TimedOut"
)

type PushJournal200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PushJournal200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PushJournal200ApplicationJSONValidation struct {
	Errors   []PushJournal200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PushJournal200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PushJournal200ApplicationJSON struct {
	Changes           []PushJournal200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                           `json:"companyId"`
	CompletedOnUtc    *time.Time                                       `json:"completedOnUtc,omitempty"`
	Data              *PushJournal200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                           `json:"dataConnectionKey"`
	DataType          *string                                          `json:"dataType,omitempty"`
	ErrorMessage      *string                                          `json:"errorMessage,omitempty"`
	PushOperationKey  string                                           `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                        `json:"requestedOnUtc"`
	Status            PushJournal200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                              `json:"statusCode"`
	TimeoutInMinutes  *int                                             `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                             `json:"timeoutInSeconds,omitempty"`
	Validation        *PushJournal200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PushJournalResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	PushJournal200ApplicationJSONObject *PushJournal200ApplicationJSON
}
