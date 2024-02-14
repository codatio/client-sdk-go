// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// JournalEntryRecordRefDataType - Name of underlying data type.
type JournalEntryRecordRefDataType string

const (
	JournalEntryRecordRefDataTypeBankTransactions JournalEntryRecordRefDataType = "bankTransactions"
	JournalEntryRecordRefDataTypeBillCreditNotes  JournalEntryRecordRefDataType = "billCreditNotes"
	JournalEntryRecordRefDataTypeBillPayments     JournalEntryRecordRefDataType = "billPayments"
	JournalEntryRecordRefDataTypeBills            JournalEntryRecordRefDataType = "bills"
	JournalEntryRecordRefDataTypeCreditNotes      JournalEntryRecordRefDataType = "creditNotes"
	JournalEntryRecordRefDataTypeDirectCosts      JournalEntryRecordRefDataType = "directCosts"
	JournalEntryRecordRefDataTypeDirectIncomes    JournalEntryRecordRefDataType = "directIncomes"
	JournalEntryRecordRefDataTypeInvoices         JournalEntryRecordRefDataType = "invoices"
	JournalEntryRecordRefDataTypeJournalEntries   JournalEntryRecordRefDataType = "journalEntries"
	JournalEntryRecordRefDataTypePayments         JournalEntryRecordRefDataType = "payments"
	JournalEntryRecordRefDataTypeTransfers        JournalEntryRecordRefDataType = "transfers"
)

func (e JournalEntryRecordRefDataType) ToPointer() *JournalEntryRecordRefDataType {
	return &e
}

func (e *JournalEntryRecordRefDataType) UnmarshalJSON(data []byte) error {
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
		*e = JournalEntryRecordRefDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JournalEntryRecordRefDataType: %v", v)
	}
}

// JournalEntryRecordRef - Links a journal entry to the underlying record that created it.
type JournalEntryRecordRef struct {
	// Name of underlying data type.
	DataType *JournalEntryRecordRefDataType `json:"dataType,omitempty"`
	// 'id' of the underlying record or data type.
	ID *string `json:"id,omitempty"`
}

func (o *JournalEntryRecordRef) GetDataType() *JournalEntryRecordRefDataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *JournalEntryRecordRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
