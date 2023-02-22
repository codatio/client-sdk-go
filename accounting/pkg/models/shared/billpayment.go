package shared

import (
	"time"
)

type BillPaymentLinesLinksTypeEnum string

const (
	BillPaymentLinesLinksTypeEnumUnknown          BillPaymentLinesLinksTypeEnum = "Unknown"
	BillPaymentLinesLinksTypeEnumUnlinked         BillPaymentLinesLinksTypeEnum = "Unlinked"
	BillPaymentLinesLinksTypeEnumBill             BillPaymentLinesLinksTypeEnum = "Bill"
	BillPaymentLinesLinksTypeEnumOther            BillPaymentLinesLinksTypeEnum = "Other"
	BillPaymentLinesLinksTypeEnumCreditNote       BillPaymentLinesLinksTypeEnum = "CreditNote"
	BillPaymentLinesLinksTypeEnumBillPayment      BillPaymentLinesLinksTypeEnum = "BillPayment"
	BillPaymentLinesLinksTypeEnumPaymentOnAccount BillPaymentLinesLinksTypeEnum = "PaymentOnAccount"
	BillPaymentLinesLinksTypeEnumRefund           BillPaymentLinesLinksTypeEnum = "Refund"
	BillPaymentLinesLinksTypeEnumManualJournal    BillPaymentLinesLinksTypeEnum = "ManualJournal"
	BillPaymentLinesLinksTypeEnumDiscount         BillPaymentLinesLinksTypeEnum = "Discount"
)

type BillPaymentLinesLinks struct {
	Amount       *float64                      `json:"amount,omitempty"`
	CurrencyRate *float64                      `json:"currencyRate,omitempty"`
	ID           *string                       `json:"id,omitempty"`
	Type         BillPaymentLinesLinksTypeEnum `json:"type"`
}

type BillPaymentLines struct {
	AllocatedOnDate *time.Time              `json:"allocatedOnDate,omitempty"`
	Amount          float64                 `json:"amount"`
	Links           []BillPaymentLinesLinks `json:"links,omitempty"`
}

// BillPayment
// > **Bill payments or payments?**
// >
// > In Codat, bill payments represent accounts payable only. For accounts receivable, see [payments](https://docs.codat.io/docs/datamodel-accounting-payments), which includes [invoices](https://docs.codat.io/docs/datamodel-accounting-invoices) and [credit notes](https://docs.codat.io/docs/datamodel-accounting-creditnotes).
//
// > View the coverage for bill payments in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=billPayments" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Bill payments include all accounts payable transaction data. This includes [bills](https://docs.codat.io/docs/datamodel-accounting-bills) and [credit notes against bills](https://docs.codat.io/docs/datamodel-accounting-billcreditnotes).
//
// A bill payment in Codat usually represents an allocation of money within any customer accounts payable account. This includes but is not strictly limited to:
//
// - A payment made against a bill—for example, a credit card payment, cheque payment, or cash payment.
// - An allocation of a supplier's credit note, to a bill or perhaps a refund.
// - A bill payment made directly to an accounts payable account. This could be an overpayment or a prepayment, or a refund of a payment made directly to an accounts payable account.
//
// Depending on the bill payments which are allowed by the underlying accounting package, some of these types may be combined. Please see the [Example data](#example-data) section for samples of what these cases look like.
//
// In Codat, a bill payment contains details of:
//
// - When the bill payment was recorded in the accounting system.
// - How much it is for and in the currency.
// - Who the payment has been paid to, the _supplier_.
// - The types of bill payments, the _line items_.
//
// Some accounting platforms give a separate name to purchases where the payment is made immediately, such as something bought with a credit card or online payment. One example of this would be QuickBooks Online's _expenses_. You can find these types of transactions in our [Direct costs](https://docs.codat.io/docs/datamodel-accounting-directcosts) data model.
//
// Bill payments is a child data type of [account transactions](https://docs.codat.io/docs/datamodel-accounting-account-transactions).
//
// ---
//
// ## Bill payment types
//
// ## Payment of a bill
//
// A payment paying a single bill should have the following properties:
//
// - A `totalAmount` indicating the amount of the bill that was paid. This is always positive.
// - A `lines` array containing one element with the following properties:
//   - An `amount` equal to the `totalAmount` above.
//   - A `links` array containing one element with the following properties:
//   - A `type` indicating the type of link, in this case a `Bill`.
//   - An `id` containing the ID of the bill that was paid.
//   - An amount of `-totalAmount` (negative `totalAmount`), indicating that the entirety of the paid amount is allocated to the bill.
//
// ## Payment of multiple bills
//
// It is possible for one payment to pay multiple bills. This can be represented using two possible formats, depending on how the supplier keeps their books:
//
// 1. The payment has multiple entries in its **lines** array, one for each bill that is paid. Each line will follow the above example for paying a bill, and the rules detailed in the data model.
// 2. The payment has a line with multiple links to each bill. This occurs when the proportion of the original payment allocated to each bill is not available.
//
// Each line is the same as those described above, with the **amount** indicating how much of the payment is allocated to the bill. The **amount** on the lines sum to the **totalAmount** on the payment.
//
// > 🚧 Pushing batch payments to Xero
// >
// > When pushing a single bill payment to Xero to pay multiple bills, only the first format is supported—multiple entries in the payment **lines** array.
//
// ## Payments and refunds on account
//
// A payment on account, that is a payment that doesn’t pay a specific bill, has one entry in its lines array.
//
// The line has the following properties:
//
// - A **totalAmount** indicating the amount paid by a supplier or refunded to them by a company. A payment to the supplier is always negative. A refund is always positive.
// - A **links** array containing one element with the following properties:
//   - A **type** indicating the type of link. For a payment this is `PaymentOnAccount`. For a refund this is `Refund`.
//   - The **id** containing the ID of the supplier.
//   - An amount for the link is `0` **totalAmount** or the amount of the payment or refund.
//
// It is possible to have a payment that is part on account and part allocated to a bill. Each line should follow the examples above.
//
// ## Using a credit note to pay a bill
//
// The payment of a bill using a credit note has one entry in its `lines` array. This **line** has the following properties:
//
// - An **amount** indicating the amount of money moved, which in this case is `0`, as the credit note and bill allocation must balance each other.
// - A **links** array containing two elements:
//   - The first link has:
//   - A **type** indicating the type of link, in this case a `Bill`.
//   - An **id** containing the ID of the bill that was paid.
//   - The second link has:
//   - A **type** indicating the type of link, in this case a `CreditNote`.
//   - An **id** containing the ID of the credit note used by this payment.
//
// The **amount** field on the **line** equals the **totalAmount** on the payment.
//
// ## Refunding a credit note
//
// A bill payment refunding a credit note has one entry in its **lines** array. This line has the following properties:
//
// - An **amount** indicating the amount of the credit note that was refunded. This is always negative, indicating that it is a refund.
// - A **links** array containing one element with the following properties:
//   - A **type** indicating the type of `link`, in this case a `CreditNote`.
//   - An **id** containing the ID of the credit note that was refunded.
//
// The **totalAmount** field on the payment equals the line's **amount** field. These are both negative, as this is money leaving accounts payable.
//
// ## Refunding a payment
//
// If a payment is refunded, for example, when a company overpaid a bill and the overpayment is returned, there are two payment records:
//
// - One for the incoming overpayment.
// - Another for the outgoing refund.
//
// The payment issuing the refund is identified by the fact that the **totalAmount** is negative. This payment has one entry in its lines array that have the following properties:
//
// - An **amount** indicating the amount that was refunded. This is always negative.
// - A **links** array containing one element with the following properties:
//   - A **type** indicating the type of a the link, in this case a `BillPayment`.
//   - An **id** containing the ID of the payment that was refunded.
//
// The **amount** field on the line equals the **totalAmount** on the payment and is negative as this is money leaving accounts payable.
//
// The payment that was refunded can be identified as it has a line where the `amount` on its `line` is positive and the type of the link is `Refund`. This payment may have several entries in its **lines** array if it was partly used to pay an bill. For example, a £1,050 payment paying a £1,000 bill with a refund of £50 has two lines:
//
// - One for £1,000 linked to the bill that was paid
// - Another for £50 linked to the payment that refunded the over payment. This link is of type `Refund` but the ID corresponds to a bill payment.
//
// The line linked to the bill payment has the following properties:
//
// - An **amount** indicating the amount that was refunded. This is positive as its money that was added to accounts payable, but is balanced out by the negative amount of the refund.
// - A **links** array containing one element with the following properties:
//   - A **type** indicating the type of the link, in this case a `Refund`.
//   - An **id** containing the ID of the payment that refunded this line.
//
// > 📘 Linked payments
// >
// > Not all accounting packages support linked payments in this way. In these platforms you may see a payment on account and a refund on account.
//
// ## Foreign currencies
//
// There are two types of currency rate that are detailed in the bill payments data type:
//
// Payment currency rate:
//
// - Base currency of the accounts payable account.
// - Foreign currency of the bill payment.
//
// Payment line link currency rate:
//
// - Base currency of the item that the link represents.
// - Foreign currency of the payment.
//
// These two rates allow the calculation of currency loss or gain for any of the transactions affected by the payment lines. The second rate is used when a bill payment is applied to an item in a currency that does not match either:
//
// - The base currency for the accounts payable account.
// - The currency of the item.
type BillPayment struct {
	AccountRef         *AccountRef        `json:"accountRef,omitempty"`
	Currency           *string            `json:"currency,omitempty"`
	CurrencyRate       *float64           `json:"currencyRate,omitempty"`
	Date               time.Time          `json:"date"`
	ID                 *string            `json:"id,omitempty"`
	Lines              []BillPaymentLines `json:"lines,omitempty"`
	Metadata           *Metadata          `json:"metadata,omitempty"`
	ModifiedDate       *time.Time         `json:"modifiedDate,omitempty"`
	Note               *string            `json:"note,omitempty"`
	PaymentMethodRef   *PaymentMethodRef  `json:"paymentMethodRef,omitempty"`
	Reference          *string            `json:"reference,omitempty"`
	SourceModifiedDate *time.Time         `json:"sourceModifiedDate,omitempty"`
	SupplementalData   *SupplementalData  `json:"supplementalData,omitempty"`
	SupplierRef        *SupplierRef       `json:"supplierRef,omitempty"`
	TotalAmount        *float64           `json:"totalAmount,omitempty"`
}
