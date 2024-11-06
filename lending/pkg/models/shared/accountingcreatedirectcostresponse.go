// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// AccountingCreateDirectCostResponseAccountingDirectCost - > **Language tip:** Direct costs may also be referred to as **Spend transactions**, **Spend money transactions**, or **Payments** in various accounting software.
//
// ## Overview
//
// Direct costs are the expenses associated with a business' operations. For example, purchases of raw materials and service fees are considered direct costs.
//
// Direct costs include:
//   - Purchasing an item and paying it off at the point of the purchase
//   - Receiving cash from a refunded item if the refund is made by the supplier
//   - Withdrawing money from a bank account
//   - Writing a cheque
//
// Direct costs is a child data type of [account transactions](https://docs.codat.io/lending-api#/schemas/AccountTransaction).
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type AccountingCreateDirectCostResponseAccountingDirectCost struct {
	ContactRef *ContactRef `json:"contactRef,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency string `json:"currency"`
	// Rate to convert the total amount of the payment into the base currency for the company at the time of the payment.
	//
	// Currency rates in Codat are implemented as the multiple of foreign currency units to each base currency unit.
	//
	// It is not possible to perform the currency conversion with two or more non-base currencies participating in the transaction. For example, if a company's base currency is USD, and it has a bill issued in EUR, then the bill payment must happen in USD or EUR.
	//
	// Where the currency rate is provided by the underlying accounting software, it will be available from Codat with the same precision (up to a maximum of 9 decimal places).
	//
	// For accounting software which do not provide an explicit currency rate, it is calculated as `baseCurrency / foreignCurrency` and will be returned to 9 decimal places.
	//
	// ## Examples with base currency of GBP
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (GBP) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **USD**          | $20            | 0.781         | £15.62                     |
	// | **EUR**          | €20            | 0.885         | £17.70                     |
	// | **RUB**          | ₽20            | 0.011         | £0.22                      |
	//
	// ## Examples with base currency of USD
	//
	// | Foreign Currency | Foreign Amount | Currency Rate | Base Currency Amount (USD) |
	// | :--------------- | :------------- | :------------ | :------------------------- |
	// | **GBP**          | £20            | 1.277         | $25.54                     |
	// | **EUR**          | €20            | 1.134         | $22.68                     |
	// | **RUB**          | ₽20            | 0.015         | $0.30                      |
	//
	//
	// ### Integration-specific details
	//
	// | Integration       | Scenario                                        | System behavior                                                                                                                                                      |
	// |-------------------|-------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	// | QuickBooks Online | Transaction currency differs from base currency | If currency rate value is left `null`, a rate of 1 will be used by QBO by default. To override this, specify a currencyRate in the request body.  |
	CurrencyRate *decimal.Big `decimal:"number" json:"currencyRate,omitempty"`
	// Identifier of the direct cost, unique for the company.
	ID *string `json:"id,omitempty"`
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
	IssueDate string `json:"issueDate"`
	// An array of line items.
	LineItems    []DirectCostLineItem `json:"lineItems"`
	Metadata     *Metadata            `json:"metadata,omitempty"`
	ModifiedDate *string              `json:"modifiedDate,omitempty"`
	// A note attached to the direct cost.
	Note *string `json:"note,omitempty"`
	// An array of payment allocations.
	PaymentAllocations []AccountingPaymentAllocation `json:"paymentAllocations"`
	// User-friendly reference for the direct cost.
	Reference          *string `json:"reference,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// The total amount of the direct costs, excluding any taxes.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// The total amount of tax on the direct costs.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount"`
	// The amount of the direct costs, inclusive of tax.
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount"`
}

func (a AccountingCreateDirectCostResponseAccountingDirectCost) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingCreateDirectCostResponseAccountingDirectCost) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetContactRef() *ContactRef {
	if o == nil {
		return nil
	}
	return o.ContactRef
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetLineItems() []DirectCostLineItem {
	if o == nil {
		return []DirectCostLineItem{}
	}
	return o.LineItems
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetPaymentAllocations() []AccountingPaymentAllocation {
	if o == nil {
		return []AccountingPaymentAllocation{}
	}
	return o.PaymentAllocations
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetSubTotal() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.SubTotal
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TaxAmount
}

func (o *AccountingCreateDirectCostResponseAccountingDirectCost) GetTotalAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalAmount
}

type AccountingCreateDirectCostResponse struct {
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
	CompletedOnUtc *string                                                 `json:"completedOnUtc,omitempty"`
	Data           *AccountingCreateDirectCostResponseAccountingDirectCost `json:"data,omitempty"`
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

func (o *AccountingCreateDirectCostResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *AccountingCreateDirectCostResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *AccountingCreateDirectCostResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *AccountingCreateDirectCostResponse) GetData() *AccountingCreateDirectCostResponseAccountingDirectCost {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *AccountingCreateDirectCostResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *AccountingCreateDirectCostResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountingCreateDirectCostResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *AccountingCreateDirectCostResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *AccountingCreateDirectCostResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *AccountingCreateDirectCostResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *AccountingCreateDirectCostResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountingCreateDirectCostResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *AccountingCreateDirectCostResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *AccountingCreateDirectCostResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}
