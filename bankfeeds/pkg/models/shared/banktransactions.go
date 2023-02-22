package shared

import (
	"time"
)

type BankTransactionsTransactionsTransactionTypeEnum string

const (
	BankTransactionsTransactionsTransactionTypeEnumUnknown     BankTransactionsTransactionsTransactionTypeEnum = "Unknown"
	BankTransactionsTransactionsTransactionTypeEnumCredit      BankTransactionsTransactionsTransactionTypeEnum = "Credit"
	BankTransactionsTransactionsTransactionTypeEnumDebit       BankTransactionsTransactionsTransactionTypeEnum = "Debit"
	BankTransactionsTransactionsTransactionTypeEnumInt         BankTransactionsTransactionsTransactionTypeEnum = "Int"
	BankTransactionsTransactionsTransactionTypeEnumDiv         BankTransactionsTransactionsTransactionTypeEnum = "Div"
	BankTransactionsTransactionsTransactionTypeEnumFee         BankTransactionsTransactionsTransactionTypeEnum = "Fee"
	BankTransactionsTransactionsTransactionTypeEnumSerChg      BankTransactionsTransactionsTransactionTypeEnum = "SerChg"
	BankTransactionsTransactionsTransactionTypeEnumDep         BankTransactionsTransactionsTransactionTypeEnum = "Dep"
	BankTransactionsTransactionsTransactionTypeEnumAtm         BankTransactionsTransactionsTransactionTypeEnum = "Atm"
	BankTransactionsTransactionsTransactionTypeEnumPos         BankTransactionsTransactionsTransactionTypeEnum = "Pos"
	BankTransactionsTransactionsTransactionTypeEnumXfer        BankTransactionsTransactionsTransactionTypeEnum = "Xfer"
	BankTransactionsTransactionsTransactionTypeEnumCheck       BankTransactionsTransactionsTransactionTypeEnum = "Check"
	BankTransactionsTransactionsTransactionTypeEnumPayment     BankTransactionsTransactionsTransactionTypeEnum = "Payment"
	BankTransactionsTransactionsTransactionTypeEnumCash        BankTransactionsTransactionsTransactionTypeEnum = "Cash"
	BankTransactionsTransactionsTransactionTypeEnumDirectDep   BankTransactionsTransactionsTransactionTypeEnum = "DirectDep"
	BankTransactionsTransactionsTransactionTypeEnumDirectDebit BankTransactionsTransactionsTransactionTypeEnum = "DirectDebit"
	BankTransactionsTransactionsTransactionTypeEnumRepeatPmt   BankTransactionsTransactionsTransactionTypeEnum = "RepeatPmt"
	BankTransactionsTransactionsTransactionTypeEnumOther       BankTransactionsTransactionsTransactionTypeEnum = "Other"
)

type BankTransactionsTransactions struct {
	Amount             float64                                         `json:"amount"`
	Balance            float64                                         `json:"balance"`
	Counterparty       *string                                         `json:"counterparty,omitempty"`
	Date               time.Time                                       `json:"date"`
	Description        *string                                         `json:"description,omitempty"`
	ID                 *string                                         `json:"id,omitempty"`
	ModifiedDate       *time.Time                                      `json:"modifiedDate,omitempty"`
	Reconciled         bool                                            `json:"reconciled"`
	Reference          *string                                         `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                                      `json:"sourceModifiedDate,omitempty"`
	TransactionType    BankTransactionsTransactionsTransactionTypeEnum `json:"transactionType"`
}

// BankTransactions
// > **Accessing Bank Accounts through Banking API**
// >
// > This datatype was originally used for accessing bank account data both in accounting integrations and open banking aggregators.
// >
// > To view bank account data through the Banking API, please refer to the new datatype [here](https://docs.codat.io/docs/data-model-banking-banking-transactions)
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
type BankTransactions struct {
	AccountID       *string                        `json:"accountId,omitempty"`
	ContractVersion *string                        `json:"contractVersion,omitempty"`
	Transactions    []BankTransactionsTransactions `json:"transactions,omitempty"`
}
