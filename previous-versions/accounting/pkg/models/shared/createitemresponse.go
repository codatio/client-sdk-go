// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// AccountingItem - > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type AccountingItem struct {
	// Item details that are only for bills.
	BillItem *BillItem `json:"billItem,omitempty"`
	// Friendly reference for the item.
	Code *string `json:"code,omitempty"`
	// Identifier for the item that is unique to a company in the accounting software.
	ID *string `json:"id,omitempty"`
	// Item details that are only for bills.
	InvoiceItem *InvoiceItem `json:"invoiceItem,omitempty"`
	// Whether you can use this item for bills.
	IsBillItem bool `json:"isBillItem"`
	// Whether you can use this item for invoices.
	IsInvoiceItem bool `json:"isInvoiceItem"`
	// Current state of the item, either:
	//
	// - `Active`: Available for use
	// - `Archived`: Unavailable
	// - `Unknown`
	//
	// Due to a [limitation in Xero's API](https://docs.codat.io/integrations/accounting/xero/xero-faq#why-do-all-of-my-items-from-xero-have-their-status-as-unknown), all items from Xero are mapped as `Unknown`.
	ItemStatus   ItemStatus `json:"itemStatus"`
	Metadata     *Metadata  `json:"metadata,omitempty"`
	ModifiedDate *string    `json:"modifiedDate,omitempty"`
	// Name of the item in the accounting software.
	Name               *string `json:"name,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Type of the item.
	Type ItemType `json:"type"`
}

func (o *AccountingItem) GetBillItem() *BillItem {
	if o == nil {
		return nil
	}
	return o.BillItem
}

func (o *AccountingItem) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *AccountingItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingItem) GetInvoiceItem() *InvoiceItem {
	if o == nil {
		return nil
	}
	return o.InvoiceItem
}

func (o *AccountingItem) GetIsBillItem() bool {
	if o == nil {
		return false
	}
	return o.IsBillItem
}

func (o *AccountingItem) GetIsInvoiceItem() bool {
	if o == nil {
		return false
	}
	return o.IsInvoiceItem
}

func (o *AccountingItem) GetItemStatus() ItemStatus {
	if o == nil {
		return ItemStatus("")
	}
	return o.ItemStatus
}

func (o *AccountingItem) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingItem) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountingItem) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingItem) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingItem) GetType() ItemType {
	if o == nil {
		return ItemType("")
	}
	return o.Type
}

type CreateItemResponse struct {
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
	CompletedOnUtc *string         `json:"completedOnUtc,omitempty"`
	Data           *AccountingItem `json:"data,omitempty"`
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

func (o *CreateItemResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *CreateItemResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateItemResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *CreateItemResponse) GetData() *AccountingItem {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *CreateItemResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *CreateItemResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *CreateItemResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateItemResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *CreateItemResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *CreateItemResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *CreateItemResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateItemResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *CreateItemResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *CreateItemResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}
