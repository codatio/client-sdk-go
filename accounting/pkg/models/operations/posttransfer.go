package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostTransferPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

// PostTransferSourceModifiedDateContactRef
// The customer or supplier for the transfer, if available.
type PostTransferSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// PostTransferSourceModifiedDateTransferAccountAccountRef
// The account that the transfer is moving from or to.
type PostTransferSourceModifiedDateTransferAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostTransferSourceModifiedDateTransferAccount
// The details of the accounts the transfer is moving from.
type PostTransferSourceModifiedDateTransferAccount struct {
	AccountRef *PostTransferSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	Amount     *float64                                                 `json:"amount,omitempty"`
	Currency   *string                                                  `json:"currency,omitempty"`
}

type PostTransferSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostTransferSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type PostTransferSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostTransferSourceModifiedDate
// > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type PostTransferSourceModifiedDate struct {
	ContactRef           *PostTransferSourceModifiedDateContactRef            `json:"contactRef,omitempty"`
	Date                 *time.Time                                           `json:"date,omitempty"`
	DepositedRecordRefs  []string                                             `json:"depositedRecordRefs,omitempty"`
	Description          *string                                              `json:"description,omitempty"`
	From                 *PostTransferSourceModifiedDateTransferAccount       `json:"from,omitempty"`
	ID                   *string                                              `json:"id,omitempty"`
	Metadata             *PostTransferSourceModifiedDateMetadata              `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                           `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                           `json:"sourceModifiedDate,omitempty"`
	SupplementalData     *PostTransferSourceModifiedDateSupplementalData      `json:"supplementalData,omitempty"`
	To                   *PostTransferSourceModifiedDateTransferAccount       `json:"to,omitempty"`
	TrackingCategoryRefs []PostTransferSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

type PostTransferSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostTransferRequest struct {
	PathParams PostTransferPathParams
	Request    *PostTransferSourceModifiedDate `request:"mediaType=application/json"`
	Security   PostTransferSecurity
}

type PostTransfer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostTransfer200ApplicationJSONChangesTypeEnum string

const (
	PostTransfer200ApplicationJSONChangesTypeEnumUnknown            PostTransfer200ApplicationJSONChangesTypeEnum = "Unknown"
	PostTransfer200ApplicationJSONChangesTypeEnumCreated            PostTransfer200ApplicationJSONChangesTypeEnum = "Created"
	PostTransfer200ApplicationJSONChangesTypeEnumModified           PostTransfer200ApplicationJSONChangesTypeEnum = "Modified"
	PostTransfer200ApplicationJSONChangesTypeEnumDeleted            PostTransfer200ApplicationJSONChangesTypeEnum = "Deleted"
	PostTransfer200ApplicationJSONChangesTypeEnumAttachmentUploaded PostTransfer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostTransfer200ApplicationJSONChanges struct {
	AttachmentID *string                                                      `json:"attachmentId,omitempty"`
	RecordRef    *PostTransfer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostTransfer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostTransfer200ApplicationJSONSourceModifiedDateContactRef
// The customer or supplier for the transfer, if available.
type PostTransfer200ApplicationJSONSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// PostTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef
// The account that the transfer is moving from or to.
type PostTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostTransfer200ApplicationJSONSourceModifiedDateTransferAccount
// The details of the accounts the transfer is moving from.
type PostTransfer200ApplicationJSONSourceModifiedDateTransferAccount struct {
	AccountRef *PostTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	Amount     *float64                                                                   `json:"amount,omitempty"`
	Currency   *string                                                                    `json:"currency,omitempty"`
}

type PostTransfer200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostTransfer200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

type PostTransfer200ApplicationJSONSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostTransfer200ApplicationJSONSourceModifiedDate
// > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type PostTransfer200ApplicationJSONSourceModifiedDate struct {
	ContactRef           *PostTransfer200ApplicationJSONSourceModifiedDateContactRef            `json:"contactRef,omitempty"`
	Date                 *time.Time                                                             `json:"date,omitempty"`
	DepositedRecordRefs  []string                                                               `json:"depositedRecordRefs,omitempty"`
	Description          *string                                                                `json:"description,omitempty"`
	From                 *PostTransfer200ApplicationJSONSourceModifiedDateTransferAccount       `json:"from,omitempty"`
	ID                   *string                                                                `json:"id,omitempty"`
	Metadata             *PostTransfer200ApplicationJSONSourceModifiedDateMetadata              `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                             `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                                             `json:"sourceModifiedDate,omitempty"`
	SupplementalData     *PostTransfer200ApplicationJSONSourceModifiedDateSupplementalData      `json:"supplementalData,omitempty"`
	To                   *PostTransfer200ApplicationJSONSourceModifiedDateTransferAccount       `json:"to,omitempty"`
	TrackingCategoryRefs []PostTransfer200ApplicationJSONSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

type PostTransfer200ApplicationJSONStatusEnum string

const (
	PostTransfer200ApplicationJSONStatusEnumPending  PostTransfer200ApplicationJSONStatusEnum = "Pending"
	PostTransfer200ApplicationJSONStatusEnumFailed   PostTransfer200ApplicationJSONStatusEnum = "Failed"
	PostTransfer200ApplicationJSONStatusEnumSuccess  PostTransfer200ApplicationJSONStatusEnum = "Success"
	PostTransfer200ApplicationJSONStatusEnumTimedOut PostTransfer200ApplicationJSONStatusEnum = "TimedOut"
)

type PostTransfer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostTransfer200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostTransfer200ApplicationJSONValidation struct {
	Errors   []PostTransfer200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostTransfer200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostTransfer200ApplicationJSON struct {
	Changes           []PostTransfer200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                            `json:"companyId"`
	CompletedOnUtc    *time.Time                                        `json:"completedOnUtc,omitempty"`
	Data              *PostTransfer200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                            `json:"dataConnectionKey"`
	DataType          *string                                           `json:"dataType,omitempty"`
	ErrorMessage      *string                                           `json:"errorMessage,omitempty"`
	PushOperationKey  string                                            `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                         `json:"requestedOnUtc"`
	Status            PostTransfer200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                               `json:"statusCode"`
	TimeoutInMinutes  *int                                              `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                              `json:"timeoutInSeconds,omitempty"`
	Validation        *PostTransfer200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostTransferResponse struct {
	ContentType                          string
	StatusCode                           int
	PostTransfer200ApplicationJSONObject *PostTransfer200ApplicationJSON
}
