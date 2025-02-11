// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ValidFor string

const (
	ValidForExpenseTransactionsPayment      ValidFor = "expense-transactions.Payment"
	ValidForExpenseTransactionsRefund       ValidFor = "expense-transactions.Refund"
	ValidForExpenseTransactionsReward       ValidFor = "expense-transactions.Reward"
	ValidForExpenseTransactionsChargeback   ValidFor = "expense-transactions.Chargeback"
	ValidForReimbursableExpenseTransactions ValidFor = "reimbursable-expense-transactions"
	ValidForTransferTransactions            ValidFor = "transfer-transactions"
	ValidForAdjustmentTransactions          ValidFor = "adjustment-transactions"
)

func (e ValidFor) ToPointer() *ValidFor {
	return &e
}
func (e *ValidFor) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "expense-transactions.Payment":
		fallthrough
	case "expense-transactions.Refund":
		fallthrough
	case "expense-transactions.Reward":
		fallthrough
	case "expense-transactions.Chargeback":
		fallthrough
	case "reimbursable-expense-transactions":
		fallthrough
	case "transfer-transactions":
		fallthrough
	case "adjustment-transactions":
		*e = ValidFor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ValidFor: %v", v)
	}
}
