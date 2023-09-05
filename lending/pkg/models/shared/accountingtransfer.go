// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AccountingTransferContactRef - The customer or supplier for the transfer, if available.
type AccountingTransferContactRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       string  `json:"id"`
}

func (o *AccountingTransferContactRef) GetDataType() *string {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountingTransferContactRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// AccountingTransfer - > View the coverage for transfers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=transfers" target="_blank">Data coverage explorer</a>.
//
// A transfer records the movement of money between two bank accounts, or between a bank account and a nominal account. It is a child data type of [account transactions](https://docs.codat.io/accounting-api#/schemas/AccountTransaction).
type AccountingTransfer struct {
	// The customer or supplier for the transfer, if available.
	ContactRef *AccountingTransferContactRef `json:"contactRef,omitempty"`
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
	Description *string          `json:"description,omitempty"`
	From        *TransferAccount `json:"from,omitempty"`
	// Unique identifier for the transfer.
	ID                 *string   `json:"id,omitempty"`
	Metadata           *Metadata `json:"metadata,omitempty"`
	ModifiedDate       *string   `json:"modifiedDate,omitempty"`
	SourceModifiedDate *string   `json:"sourceModifiedDate,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/additional-data) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	To               *TransferAccount  `json:"to,omitempty"`
	// Reference to the tracking categories this transfer is being tracked against.
	TrackingCategoryRefs []TrackingCategoryRef `json:"trackingCategoryRefs,omitempty"`
}

func (o *AccountingTransfer) GetContactRef() *AccountingTransferContactRef {
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