// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/assess/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type EnhancedCashFlowTransaction struct {
	// An account reference containing the account id and name.
	AccountRef *AccountRef `json:"accountRef,omitempty"`
	// The bank transaction amount.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
	// An array of counterparty names involved in the transaction.
	CounterpartyNames []string `json:"counterpartyNames,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
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
	// The description of the bank transaction.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the bank transaction.
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
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Returns the payment processor responsible for the transaction.
	PlatformName *string `json:"platformName,omitempty"`
	// A source reference containing the `sourceType` object "Banking".
	SourceRef           *SourceRef           `json:"sourceRef,omitempty"`
	TransactionCategory *TransactionCategory `json:"transactionCategory,omitempty"`
}

func (e EnhancedCashFlowTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EnhancedCashFlowTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EnhancedCashFlowTransaction) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *EnhancedCashFlowTransaction) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *EnhancedCashFlowTransaction) GetCounterpartyNames() []string {
	if o == nil {
		return nil
	}
	return o.CounterpartyNames
}

func (o *EnhancedCashFlowTransaction) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *EnhancedCashFlowTransaction) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *EnhancedCashFlowTransaction) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *EnhancedCashFlowTransaction) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EnhancedCashFlowTransaction) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *EnhancedCashFlowTransaction) GetPlatformName() *string {
	if o == nil {
		return nil
	}
	return o.PlatformName
}

func (o *EnhancedCashFlowTransaction) GetSourceRef() *SourceRef {
	if o == nil {
		return nil
	}
	return o.SourceRef
}

func (o *EnhancedCashFlowTransaction) GetTransactionCategory() *TransactionCategory {
	if o == nil {
		return nil
	}
	return o.TransactionCategory
}
