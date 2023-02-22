package shared

import (
	"time"
)

type BankStatementLineTransactionTypeEnum string

const (
	BankStatementLineTransactionTypeEnumUnknown     BankStatementLineTransactionTypeEnum = "Unknown"
	BankStatementLineTransactionTypeEnumCredit      BankStatementLineTransactionTypeEnum = "Credit"
	BankStatementLineTransactionTypeEnumDebit       BankStatementLineTransactionTypeEnum = "Debit"
	BankStatementLineTransactionTypeEnumInt         BankStatementLineTransactionTypeEnum = "Int"
	BankStatementLineTransactionTypeEnumDiv         BankStatementLineTransactionTypeEnum = "Div"
	BankStatementLineTransactionTypeEnumFee         BankStatementLineTransactionTypeEnum = "Fee"
	BankStatementLineTransactionTypeEnumSerChg      BankStatementLineTransactionTypeEnum = "SerChg"
	BankStatementLineTransactionTypeEnumDep         BankStatementLineTransactionTypeEnum = "Dep"
	BankStatementLineTransactionTypeEnumAtm         BankStatementLineTransactionTypeEnum = "Atm"
	BankStatementLineTransactionTypeEnumPos         BankStatementLineTransactionTypeEnum = "Pos"
	BankStatementLineTransactionTypeEnumXfer        BankStatementLineTransactionTypeEnum = "Xfer"
	BankStatementLineTransactionTypeEnumCheck       BankStatementLineTransactionTypeEnum = "Check"
	BankStatementLineTransactionTypeEnumPayment     BankStatementLineTransactionTypeEnum = "Payment"
	BankStatementLineTransactionTypeEnumCash        BankStatementLineTransactionTypeEnum = "Cash"
	BankStatementLineTransactionTypeEnumDirectDep   BankStatementLineTransactionTypeEnum = "DirectDep"
	BankStatementLineTransactionTypeEnumDirectDebit BankStatementLineTransactionTypeEnum = "DirectDebit"
	BankStatementLineTransactionTypeEnumRepeatPmt   BankStatementLineTransactionTypeEnum = "RepeatPmt"
	BankStatementLineTransactionTypeEnumOther       BankStatementLineTransactionTypeEnum = "Other"
)

type BankStatementLine struct {
	Amount             float64                              `json:"amount"`
	Balance            float64                              `json:"balance"`
	Counterparty       *string                              `json:"counterparty,omitempty"`
	Date               time.Time                            `json:"date"`
	Description        *string                              `json:"description,omitempty"`
	ID                 *string                              `json:"id,omitempty"`
	ModifiedDate       *time.Time                           `json:"modifiedDate,omitempty"`
	Reconciled         bool                                 `json:"reconciled"`
	Reference          *string                              `json:"reference,omitempty"`
	SourceModifiedDate *time.Time                           `json:"sourceModifiedDate,omitempty"`
	TransactionType    BankStatementLineTransactionTypeEnum `json:"transactionType"`
}
