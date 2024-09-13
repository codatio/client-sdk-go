// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type UpdateBillResponseWithholdingTax struct {
	// Amount of tax withheld.
	Amount *decimal.Big `decimal:"number" json:"amount"`
	// Name assigned to withheld tax.
	Name string `json:"name"`
}

func (u UpdateBillResponseWithholdingTax) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateBillResponseWithholdingTax) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateBillResponseWithholdingTax) GetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.Amount
}

func (o *UpdateBillResponseWithholdingTax) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// UpdateBillResponseAccountingBill - > **Invoices or bills?**
// >
// > We distinguish between invoices where the company *owes money* vs. *is owed money*. If the company has received an invoice, and owes money to someone else (accounts payable) we call this a Bill.
// >
// > See [Invoices](https://docs.codat.io/accounting-api#/schemas/Invoice) for the accounts receivable equivalent of bills.
//
// View the coverage for bills in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bills" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// In Codat, a bill contains details of:
// * When the bill was recorded in the accounting system.
// * How much the bill is for and the currency of the amount.
// * Who the bill was received from — the *supplier*.
// * What the bill is for — the *line items*.
//
// Some accounting software give a separate name to purchases where the payment is made immediately, such as something bought with a credit card or online payment. One example of this would be QuickBooks Online's *expenses*.
//
// You can find these types of transactions in our [Direct costs](https://docs.codat.io/accounting-api#/schemas/DirectCost) data model.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type UpdateBillResponseAccountingBill struct {
	// Amount outstanding on the bill.
	AmountDue *decimal.Big `decimal:"number" json:"amountDue,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
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
	DueDate      *string      `json:"dueDate,omitempty"`
	// Identifier for the bill, unique for the company in the accounting software.
	ID        *string `json:"id,omitempty"`
	IssueDate string  `json:"issueDate"`
	// Array of Bill line items.
	LineItems    []BillLineItem `json:"lineItems,omitempty"`
	Metadata     *Metadata      `json:"metadata,omitempty"`
	ModifiedDate *string        `json:"modifiedDate,omitempty"`
	// Any private, company notes about the bill, such as payment information.
	Note *string `json:"note,omitempty"`
	// An array of payment allocations.
	PaymentAllocations []PaymentAllocationItems `json:"paymentAllocations,omitempty"`
	PurchaseOrderRefs  []PurchaseOrderRef       `json:"purchaseOrderRefs,omitempty"`
	// User-friendly reference for the bill.
	Reference          *string `json:"reference,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Current state of the bill.
	Status BillStatus `json:"status"`
	// Total amount of the bill, excluding any taxes.
	SubTotal *decimal.Big `decimal:"number" json:"subTotal"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Reference to the supplier the record relates to.
	SupplierRef *SupplierRef `json:"supplierRef,omitempty"`
	// Amount of tax on the bill.
	TaxAmount *decimal.Big `decimal:"number" json:"taxAmount"`
	// Amount of the bill, including tax.
	TotalAmount    *decimal.Big                       `decimal:"number" json:"totalAmount"`
	WithholdingTax []UpdateBillResponseWithholdingTax `json:"withholdingTax,omitempty"`
}

func (u UpdateBillResponseAccountingBill) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateBillResponseAccountingBill) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateBillResponseAccountingBill) GetAmountDue() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.AmountDue
}

func (o *UpdateBillResponseAccountingBill) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *UpdateBillResponseAccountingBill) GetCurrencyRate() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *UpdateBillResponseAccountingBill) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *UpdateBillResponseAccountingBill) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateBillResponseAccountingBill) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *UpdateBillResponseAccountingBill) GetLineItems() []BillLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *UpdateBillResponseAccountingBill) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *UpdateBillResponseAccountingBill) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *UpdateBillResponseAccountingBill) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *UpdateBillResponseAccountingBill) GetPaymentAllocations() []PaymentAllocationItems {
	if o == nil {
		return nil
	}
	return o.PaymentAllocations
}

func (o *UpdateBillResponseAccountingBill) GetPurchaseOrderRefs() []PurchaseOrderRef {
	if o == nil {
		return nil
	}
	return o.PurchaseOrderRefs
}

func (o *UpdateBillResponseAccountingBill) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *UpdateBillResponseAccountingBill) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *UpdateBillResponseAccountingBill) GetStatus() BillStatus {
	if o == nil {
		return BillStatus("")
	}
	return o.Status
}

func (o *UpdateBillResponseAccountingBill) GetSubTotal() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.SubTotal
}

func (o *UpdateBillResponseAccountingBill) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *UpdateBillResponseAccountingBill) GetSupplierRef() *SupplierRef {
	if o == nil {
		return nil
	}
	return o.SupplierRef
}

func (o *UpdateBillResponseAccountingBill) GetTaxAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TaxAmount
}

func (o *UpdateBillResponseAccountingBill) GetTotalAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.TotalAmount
}

func (o *UpdateBillResponseAccountingBill) GetWithholdingTax() []UpdateBillResponseWithholdingTax {
	if o == nil {
		return nil
	}
	return o.WithholdingTax
}

type UpdateBillResponse struct {
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
	CompletedOnUtc *string                           `json:"completedOnUtc,omitempty"`
	Data           *UpdateBillResponseAccountingBill `json:"data,omitempty"`
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

func (o *UpdateBillResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *UpdateBillResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateBillResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *UpdateBillResponse) GetData() *UpdateBillResponseAccountingBill {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UpdateBillResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *UpdateBillResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *UpdateBillResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateBillResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *UpdateBillResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *UpdateBillResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *UpdateBillResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateBillResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *UpdateBillResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *UpdateBillResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}
