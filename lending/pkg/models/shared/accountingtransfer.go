// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountingTransferStatus - The status of the transfer in the account
type AccountingTransferStatus string

const (
	AccountingTransferStatusUnknown      AccountingTransferStatus = "Unknown"
	AccountingTransferStatusUnreconciled AccountingTransferStatus = "Unreconciled"
	AccountingTransferStatusReconciled   AccountingTransferStatus = "Reconciled"
	AccountingTransferStatusVoid         AccountingTransferStatus = "Void"
)

func (e AccountingTransferStatus) ToPointer() *AccountingTransferStatus {
	return &e
}
func (e *AccountingTransferStatus) UnmarshalJSON(data []byte) error {
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
		*e = AccountingTransferStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingTransferStatus: %v", v)
	}
}

// AccountingTransfer - A transfer records the movement of money between two bank accounts, or between a bank account and a nominal account. It is a child data type of [account transactions](https://docs.codat.io/lending-api#/schemas/AccountTransaction).
type AccountingTransfer struct {
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
	DepositedRecordRefs []RecordRef `json:"depositedRecordRefs,omitempty"`
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
	Status *AccountingTransferStatus `json:"status,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Account details of the account sending or receiving the transfer.
	To *TransferAccount `json:"to,omitempty"`
	// Reference to the tracking categories this transfer is being tracked against.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
}

func (o *AccountingTransfer) GetContactRef() *ContactRef {
	if o == nil {
		return nil
	}
	return o.ContactRef
}

func (o *AccountingTransfer) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *AccountingTransfer) GetDepositedRecordRefs() []RecordRef {
	if o == nil {
		return nil
	}
	return o.DepositedRecordRefs
}

func (o *AccountingTransfer) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingTransfer) GetFrom() *TransferAccount {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *AccountingTransfer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingTransfer) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingTransfer) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingTransfer) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingTransfer) GetStatus() *AccountingTransferStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AccountingTransfer) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingTransfer) GetTo() *TransferAccount {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *AccountingTransfer) GetTrackingCategoryRefs() []TrackingCategoryRef {
	if o == nil {
		return nil
	}
	return o.TrackingCategoryRefs
}
