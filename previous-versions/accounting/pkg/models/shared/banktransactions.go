// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// SourceModifiedDate - > **Accessing Bank Accounts through Banking API**
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
type SourceModifiedDate struct {
	// Unique identifier to the `accountId` the bank transactions originates from.
	AccountID *string `json:"accountId,omitempty"`
	// The amount transacted in the bank transaction.
	Amount *decimal.Big `decimal:"number" json:"amount,omitempty"`
	// The remaining balance in the account with ID `accountId`.
	Balance *decimal.Big `decimal:"number" json:"balance,omitempty"`
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
	ClearedOnDate *string `json:"clearedOnDate,omitempty"`
	// Description of the bank transaction.
	Description *string `json:"description,omitempty"`
	// Identifier for the bank transaction, unique to the company in the accounting platform.
	ID           *string `json:"id,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// `True` if the bank transaction has been [reconciled](https://www.xero.com/uk/guides/what-is-bank-reconciliation/) in the accounting platform.
	Reconciled         *bool   `json:"reconciled,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Type of transaction for the bank statement line.
	TransactionType *BankTransactionType `json:"transactionType,omitempty"`
}

func (s SourceModifiedDate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceModifiedDate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceModifiedDate) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *SourceModifiedDate) GetAmount() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *SourceModifiedDate) GetBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *SourceModifiedDate) GetClearedOnDate() *string {
	if o == nil {
		return nil
	}
	return o.ClearedOnDate
}

func (o *SourceModifiedDate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SourceModifiedDate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SourceModifiedDate) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *SourceModifiedDate) GetReconciled() *bool {
	if o == nil {
		return nil
	}
	return o.Reconciled
}

func (o *SourceModifiedDate) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *SourceModifiedDate) GetTransactionType() *BankTransactionType {
	if o == nil {
		return nil
	}
	return o.TransactionType
}

type BankTransactions struct {
	Links Links `json:"_links"`
	// Current page number.
	PageNumber int64 `json:"pageNumber"`
	// Number of items to return in results array.
	PageSize int64                `json:"pageSize"`
	Results  []SourceModifiedDate `json:"results,omitempty"`
	// Total number of items.
	TotalResults int64 `json:"totalResults"`
}

func (o *BankTransactions) GetLinks() Links {
	if o == nil {
		return Links{}
	}
	return o.Links
}

func (o *BankTransactions) GetPageNumber() int64 {
	if o == nil {
		return 0
	}
	return o.PageNumber
}

func (o *BankTransactions) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *BankTransactions) GetResults() []SourceModifiedDate {
	if o == nil {
		return nil
	}
	return o.Results
}

func (o *BankTransactions) GetTotalResults() int64 {
	if o == nil {
		return 0
	}
	return o.TotalResults
}
