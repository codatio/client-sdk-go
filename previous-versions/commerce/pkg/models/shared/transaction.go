// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/commerce/v3/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// Transaction - Details of all financial transactions recorded in the commerce or point of sale system are added to the Transactions data type. For example, payments, service charges, and fees.
//
// You can use data from the Transactions endpoints to calculate key metrics, such as:
// - Transaction volumes
// - Average transaction volume
// - Average transaction value
// - Returns
// - Payouts
type Transaction struct {
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
	CreatedDate *string `json:"createdDate,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
	// A unique, persistent identifier for this record
	ID           string  `json:"id"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
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
	SourceCreatedDate  *string `json:"sourceCreatedDate,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Non-standardised transaction type data from the commerce software
	SubType *string `json:"subType,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// The total transaction amount
	TotalAmount *decimal.Big `decimal:"number" json:"totalAmount,omitempty"`
	// Link to the source event which triggered this transaction
	TransactionSourceRef *TransactionSourceRef `json:"transactionSourceRef,omitempty"`
	// The type of the platform transaction:
	// - `Unknown`
	// - `FailedPayout` — Failed transfer of funds from the seller's merchant account to their bank account.
	// - `Payment` — Credit and debit card payments.
	// - `PaymentFee` — Payment provider's fee on each card payment.
	// - `PaymentFeeRefund` — Payment provider's fee that has been refunded to the seller.
	// - `Payout` — Transfer of funds from the seller's merchant account to their bank account.
	// - `Refund` — Refunds to a customer's credit or debit card.
	// - `Transfer` — Secure transfer of funds to the seller's bank account.
	Type *TransactionType `json:"type,omitempty"`
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Transaction) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *Transaction) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Transaction) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Transaction) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Transaction) GetSourceCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceCreatedDate
}

func (o *Transaction) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *Transaction) GetSubType() *string {
	if o == nil {
		return nil
	}
	return o.SubType
}

func (o *Transaction) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *Transaction) GetTotalAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *Transaction) GetTransactionSourceRef() *TransactionSourceRef {
	if o == nil {
		return nil
	}
	return o.TransactionSourceRef
}

func (o *Transaction) GetType() *TransactionType {
	if o == nil {
		return nil
	}
	return o.Type
}
