// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/pkg/types"
)

// AccountingBankTransaction - > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/banking-api#/operations/list-all-banking-transactions)
//
// > View the coverage for bank transactions in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=bankTransactions" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Transactional banking data for a specific company and account.
//
// Bank transactions include the:
// * Amount of the transaction.
// * Current account balance.
// * Transaction type, for example, credit, debit, or transfer.
type AccountingBankTransaction struct {
	AccountID *string        `json:"accountId,omitempty"`
	Amount    *types.Decimal `json:"amount,omitempty"`
	Balance   *types.Decimal `json:"balance,omitempty"`
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
	ClearedOnDate      *string `json:"clearedOnDate,omitempty"`
	Description        *string `json:"description,omitempty"`
	ID                 *string `json:"id,omitempty"`
	ModifiedDate       *string `json:"modifiedDate,omitempty"`
	Reconciled         *bool   `json:"reconciled,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Type of transaction for the bank statement line
	TransactionType *BankTransactionType `json:"transactionType,omitempty"`
}

func (o *AccountingBankTransaction) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *AccountingBankTransaction) GetAmount() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AccountingBankTransaction) GetBalance() *types.Decimal {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *AccountingBankTransaction) GetClearedOnDate() *string {
	if o == nil {
		return nil
	}
	return o.ClearedOnDate
}

func (o *AccountingBankTransaction) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingBankTransaction) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingBankTransaction) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *AccountingBankTransaction) GetReconciled() *bool {
	if o == nil {
		return nil
	}
	return o.Reconciled
}

func (o *AccountingBankTransaction) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *AccountingBankTransaction) GetTransactionType() *BankTransactionType {
	if o == nil {
		return nil
	}
	return o.TransactionType
}
