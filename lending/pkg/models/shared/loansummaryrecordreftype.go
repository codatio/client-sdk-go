// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// LoanSummaryRecordRefType - The datatype being referred to.
type LoanSummaryRecordRefType string

const (
	LoanSummaryRecordRefTypeAccounts             LoanSummaryRecordRefType = "accounts"
	LoanSummaryRecordRefTypeBankingAccounts      LoanSummaryRecordRefType = "banking-accounts"
	LoanSummaryRecordRefTypeCommerceTransactions LoanSummaryRecordRefType = "commerce-transactions"
)

func (e LoanSummaryRecordRefType) ToPointer() *LoanSummaryRecordRefType {
	return &e
}
func (e *LoanSummaryRecordRefType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "accounts":
		fallthrough
	case "banking-accounts":
		fallthrough
	case "commerce-transactions":
		*e = LoanSummaryRecordRefType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoanSummaryRecordRefType: %v", v)
	}
}
