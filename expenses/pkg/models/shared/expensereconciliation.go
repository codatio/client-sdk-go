package shared

import (
	"time"
)

type ExpenseReconciliationTypeEnum string

const (
	ExpenseReconciliationTypeEnumPayment    ExpenseReconciliationTypeEnum = "payment"
	ExpenseReconciliationTypeEnumRefund     ExpenseReconciliationTypeEnum = "refund"
	ExpenseReconciliationTypeEnumReward     ExpenseReconciliationTypeEnum = "reward"
	ExpenseReconciliationTypeEnumChargeback ExpenseReconciliationTypeEnum = "chargeback"
	ExpenseReconciliationTypeEnumAdjustment ExpenseReconciliationTypeEnum = "adjustment"
	ExpenseReconciliationTypeEnumTransfer   ExpenseReconciliationTypeEnum = "transfer"
)

type ExpenseReconciliation struct {
	Currency     string                        `json:"currency"`
	CurrencyRate *float64                      `json:"currencyRate,omitempty"`
	ID           string                        `json:"id"`
	IssueDate    time.Time                     `json:"issueDate"`
	Lines        []ExpenseReconciliationLines  `json:"lines,omitempty"`
	MerchantName *string                       `json:"merchantName,omitempty"`
	Notes        *string                       `json:"notes,omitempty"`
	Type         ExpenseReconciliationTypeEnum `json:"type"`
}
