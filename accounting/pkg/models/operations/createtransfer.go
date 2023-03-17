package operations

import (
	"net/http"
	"time"
)

// CreateTransferSourceModifiedDateContactRef
// The customer or supplier for the transfer, if available.
type CreateTransferSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateTransferSourceModifiedDateTransferAccountAccountRef
// The account that the transfer is moving from or to.
type CreateTransferSourceModifiedDateTransferAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateTransferSourceModifiedDateTransferAccount
// The details of the accounts the transfer is moving from.
type CreateTransferSourceModifiedDateTransferAccount struct {
	AccountRef *CreateTransferSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	Amount     *float64                                                   `json:"amount,omitempty"`
	Currency   *string                                                    `json:"currency,omitempty"`
}

type CreateTransferSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateTransferSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateTransferSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateTransferSourceModifiedDateTrackingCategoryRefs
// References a category against which the item is tracked.
type CreateTransferSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateTransferSourceModifiedDate
// > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type CreateTransferSourceModifiedDate struct {
	ContactRef           *CreateTransferSourceModifiedDateContactRef            `json:"contactRef,omitempty"`
	Date                 *time.Time                                             `json:"date,omitempty"`
	DepositedRecordRefs  []string                                               `json:"depositedRecordRefs,omitempty"`
	Description          *string                                                `json:"description,omitempty"`
	From                 *CreateTransferSourceModifiedDateTransferAccount       `json:"from,omitempty"`
	ID                   *string                                                `json:"id,omitempty"`
	Metadata             *CreateTransferSourceModifiedDateMetadata              `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                             `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                             `json:"sourceModifiedDate,omitempty"`
	SupplementalData     *CreateTransferSourceModifiedDateSupplementalData      `json:"supplementalData,omitempty"`
	To                   *CreateTransferSourceModifiedDateTransferAccount       `json:"to,omitempty"`
	TrackingCategoryRefs []CreateTransferSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

type CreateTransferRequest struct {
	RequestBody  *CreateTransferSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID    string                            `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string                            `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateTransfer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateTransfer200ApplicationJSONChangesTypeEnum string

const (
	CreateTransfer200ApplicationJSONChangesTypeEnumUnknown            CreateTransfer200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateTransfer200ApplicationJSONChangesTypeEnumCreated            CreateTransfer200ApplicationJSONChangesTypeEnum = "Created"
	CreateTransfer200ApplicationJSONChangesTypeEnumModified           CreateTransfer200ApplicationJSONChangesTypeEnum = "Modified"
	CreateTransfer200ApplicationJSONChangesTypeEnumDeleted            CreateTransfer200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateTransfer200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateTransfer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateTransfer200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *CreateTransfer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateTransfer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateContactRef
// The customer or supplier for the transfer, if available.
type CreateTransfer200ApplicationJSONSourceModifiedDateContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef
// The account that the transfer is moving from or to.
type CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccount
// The details of the accounts the transfer is moving from.
type CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccount struct {
	AccountRef *CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccountAccountRef `json:"accountRef,omitempty"`
	Amount     *float64                                                                     `json:"amount,omitempty"`
	Currency   *string                                                                      `json:"currency,omitempty"`
}

type CreateTransfer200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateTransfer200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDateTrackingCategoryRefs
// References a category against which the item is tracked.
type CreateTransfer200ApplicationJSONSourceModifiedDateTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreateTransfer200ApplicationJSONSourceModifiedDate
// > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// From the **Transfers** endpoints, you can:
//
// - [Retrieve a list of all transfers for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers)
// - [Retrieve a single transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/get_companies__companyId__connections__connectionId__data_transfers__transferId_)
// - [Add a new transfer for a specified company](https://api.codat.io/swagger/index.html#/Transfers/post_companies__companyId__connections__connectionId__push_transfers)
//
// **Transfers** is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type CreateTransfer200ApplicationJSONSourceModifiedDate struct {
	ContactRef           *CreateTransfer200ApplicationJSONSourceModifiedDateContactRef            `json:"contactRef,omitempty"`
	Date                 *time.Time                                                               `json:"date,omitempty"`
	DepositedRecordRefs  []string                                                                 `json:"depositedRecordRefs,omitempty"`
	Description          *string                                                                  `json:"description,omitempty"`
	From                 *CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccount       `json:"from,omitempty"`
	ID                   *string                                                                  `json:"id,omitempty"`
	Metadata             *CreateTransfer200ApplicationJSONSourceModifiedDateMetadata              `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                               `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                                               `json:"sourceModifiedDate,omitempty"`
	SupplementalData     *CreateTransfer200ApplicationJSONSourceModifiedDateSupplementalData      `json:"supplementalData,omitempty"`
	To                   *CreateTransfer200ApplicationJSONSourceModifiedDateTransferAccount       `json:"to,omitempty"`
	TrackingCategoryRefs []CreateTransfer200ApplicationJSONSourceModifiedDateTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
}

type CreateTransfer200ApplicationJSONStatusEnum string

const (
	CreateTransfer200ApplicationJSONStatusEnumPending  CreateTransfer200ApplicationJSONStatusEnum = "Pending"
	CreateTransfer200ApplicationJSONStatusEnumFailed   CreateTransfer200ApplicationJSONStatusEnum = "Failed"
	CreateTransfer200ApplicationJSONStatusEnumSuccess  CreateTransfer200ApplicationJSONStatusEnum = "Success"
	CreateTransfer200ApplicationJSONStatusEnumTimedOut CreateTransfer200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateTransfer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateTransfer200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateTransfer200ApplicationJSONValidation struct {
	Errors   []CreateTransfer200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateTransfer200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateTransfer200ApplicationJSON struct {
	Changes           []CreateTransfer200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                              `json:"companyId"`
	CompletedOnUtc    *time.Time                                          `json:"completedOnUtc,omitempty"`
	Data              *CreateTransfer200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                              `json:"dataConnectionKey"`
	DataType          *string                                             `json:"dataType,omitempty"`
	ErrorMessage      *string                                             `json:"errorMessage,omitempty"`
	PushOperationKey  string                                              `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                           `json:"requestedOnUtc"`
	Status            CreateTransfer200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                 `json:"statusCode"`
	TimeoutInMinutes  *int                                                `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateTransfer200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateTransferResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	CreateTransfer200ApplicationJSONObject *CreateTransfer200ApplicationJSON
}
