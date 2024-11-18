// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountingCreateTransferResponseStatus - The status of the transfer in the account
type AccountingCreateTransferResponseStatus string

const (
	AccountingCreateTransferResponseStatusUnknown      AccountingCreateTransferResponseStatus = "Unknown"
	AccountingCreateTransferResponseStatusUnreconciled AccountingCreateTransferResponseStatus = "Unreconciled"
	AccountingCreateTransferResponseStatusReconciled   AccountingCreateTransferResponseStatus = "Reconciled"
	AccountingCreateTransferResponseStatusVoid         AccountingCreateTransferResponseStatus = "Void"
)

func (e AccountingCreateTransferResponseStatus) ToPointer() *AccountingCreateTransferResponseStatus {
	return &e
}
func (e *AccountingCreateTransferResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Unreconciled":
		fallthrough
	case "Reconciled":
		fallthrough
	case "Void":
		*e = AccountingCreateTransferResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingCreateTransferResponseStatus: %v", v)
	}
}

// AccountingCreateTransferResponseAccountingTransfer - A transfer records the movement of money between two bank accounts, or between a bank account and a nominal account. It is a child data type of [account transactions](https://docs.codat.io/lending-api#/schemas/AccountTransaction).
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type AccountingCreateTransferResponseAccountingTransfer struct {
	ContactRef *ContactRef `json:"contactRef,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	Date *string `json:"date,omitempty"`
	// List of selected transactions to associate with the transfer. Use this field to include transactions which are posted to the _undeposited funds_ (or other holding) account within the transfer.
	DepositedRecordRefs []AccountingRecordRef `json:"depositedRecordRefs,omitempty"`
	// Description of the transfer.
	Description *string `json:"description,omitempty"`
	// Account details of the account sending or receiving the transfer.
	From *TransferAccount `json:"from,omitempty"`
	// Unique identifier for the transfer.
	ID                 *string   `json:"id,omitempty"`
	Metadata           *Metadata `json:"metadata,omitempty"`
	ModifiedDate       *string   `json:"modifiedDate,omitempty"`
	SourceModifiedDate *string   `json:"sourceModifiedDate,omitempty"`
	// The status of the transfer in the account
	Status *AccountingCreateTransferResponseStatus `json:"status,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Account details of the account sending or receiving the transfer.
	To *TransferAccount `json:"to,omitempty"`
	// Reference to the tracking categories this transfer is being tracked against.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetContactRef() *ContactRef {
	if o == nil {
		return nil
	}
	return o.ContactRef
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetDepositedRecordRefs() []AccountingRecordRef {
	if o == nil {
		return nil
	}
	return o.DepositedRecordRefs
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetFrom() *TransferAccount {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetStatus() *AccountingCreateTransferResponseStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetTo() *TransferAccount {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *AccountingCreateTransferResponseAccountingTransfer) GetTrackingCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}

type AccountingCreateTransferResponse struct {
	// Contains a single entry that communicates which record has changed and the manner in which it changed.
	Changes []PushOperationChange `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	CompletedOnUtc *string                                             `json:"completedOnUtc,omitempty"`
	Data           *AccountingCreateTransferResponseAccountingTransfer `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// Available data types
	DataType *DataType `json:"dataType,omitempty"`
	// A message about the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey string `json:"pushOperationKey"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	RequestedOnUtc string `json:"requestedOnUtc"`
	// The current status of the push operation.
	Status PushOperationStatus `json:"status"`
	// Push status code.
	StatusCode int64 `json:"statusCode"`
	// Number of minutes the push operation must complete within before it times out.
	TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
	// Number of seconds the push operation must complete within before it times out.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *Validation `json:"validation,omitempty"`
}

func (o *AccountingCreateTransferResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *AccountingCreateTransferResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *AccountingCreateTransferResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *AccountingCreateTransferResponse) GetData() *AccountingCreateTransferResponseAccountingTransfer {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AccountingCreateTransferResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *AccountingCreateTransferResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountingCreateTransferResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *AccountingCreateTransferResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *AccountingCreateTransferResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *AccountingCreateTransferResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *AccountingCreateTransferResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountingCreateTransferResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *AccountingCreateTransferResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *AccountingCreateTransferResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}
