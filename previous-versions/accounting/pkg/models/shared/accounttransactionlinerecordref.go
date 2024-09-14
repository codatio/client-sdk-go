// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AccountTransactionLineRecordRefDataType - Name of underlying data type.
type AccountTransactionLineRecordRefDataType string

const (
	AccountTransactionLineRecordRefDataTypeBankTransactions AccountTransactionLineRecordRefDataType = "bankTransactions"
	AccountTransactionLineRecordRefDataTypeBillCreditNotes  AccountTransactionLineRecordRefDataType = "billCreditNotes"
	AccountTransactionLineRecordRefDataTypeBillPayments     AccountTransactionLineRecordRefDataType = "billPayments"
	AccountTransactionLineRecordRefDataTypeBills            AccountTransactionLineRecordRefDataType = "bills"
	AccountTransactionLineRecordRefDataTypeCreditNotes      AccountTransactionLineRecordRefDataType = "creditNotes"
	AccountTransactionLineRecordRefDataTypeDirectCosts      AccountTransactionLineRecordRefDataType = "directCosts"
	AccountTransactionLineRecordRefDataTypeDirectIncomes    AccountTransactionLineRecordRefDataType = "directIncomes"
	AccountTransactionLineRecordRefDataTypeInvoices         AccountTransactionLineRecordRefDataType = "invoices"
	AccountTransactionLineRecordRefDataTypeJournalEntries   AccountTransactionLineRecordRefDataType = "journalEntries"
	AccountTransactionLineRecordRefDataTypePayments         AccountTransactionLineRecordRefDataType = "payments"
	AccountTransactionLineRecordRefDataTypeTransfers        AccountTransactionLineRecordRefDataType = "transfers"
)

func (e AccountTransactionLineRecordRefDataType) ToPointer() *AccountTransactionLineRecordRefDataType {
	return &e
}
func (e *AccountTransactionLineRecordRefDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bankTransactions":
		fallthrough
	case "billCreditNotes":
		fallthrough
	case "billPayments":
		fallthrough
	case "bills":
		fallthrough
	case "creditNotes":
		fallthrough
	case "directCosts":
		fallthrough
	case "directIncomes":
		fallthrough
	case "invoices":
		fallthrough
	case "journalEntries":
		fallthrough
	case "payments":
		fallthrough
	case "transfers":
		*e = AccountTransactionLineRecordRefDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountTransactionLineRecordRefDataType: %v", v)
	}
}

// AccountTransactionLineRecordRef - Links an account transaction line to the underlying record that created it.
type AccountTransactionLineRecordRef struct {
	// Name of underlying data type.
	DataType *AccountTransactionLineRecordRefDataType `json:"dataType,omitempty"`
	// 'id' of the underlying record or data type.
	ID *string `json:"id,omitempty"`
}

func (o *AccountTransactionLineRecordRef) GetDataType() *AccountTransactionLineRecordRefDataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *AccountTransactionLineRecordRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}