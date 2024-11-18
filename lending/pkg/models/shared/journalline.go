// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/lending/v7/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// JournalLineDataType - Allowed name of the 'dataType'.
type JournalLineDataType string

const (
	JournalLineDataTypeCustomers JournalLineDataType = "customers"
	JournalLineDataTypeSuppliers JournalLineDataType = "suppliers"
)

func (e JournalLineDataType) ToPointer() *JournalLineDataType {
	return &e
}
func (e *JournalLineDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "customers":
		fallthrough
	case "suppliers":
		*e = JournalLineDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JournalLineDataType: %v", v)
	}
}

type ContactReference struct {
	// Allowed name of the 'dataType'.
	DataType *JournalLineDataType `json:"dataType,omitempty"`
	// Unique identifier for a customer or supplier.
	ID string `json:"id"`
}

func (o *ContactReference) GetDataType() *JournalLineDataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *ContactReference) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// JournalLineTracking - List of record refs associated with the tracking information for the line (eg to a Tracking Category, or customer etc.)
type JournalLineTracking struct {
	RecordRefs []TrackingRecordRef `json:"recordRefs,omitempty"`
}

func (o *JournalLineTracking) GetRecordRefs() []TrackingRecordRef {
	if o == nil {
		return nil
	}
	return o.RecordRefs
}

type JournalLine struct {
	// Data types that reference an account, for example bill and invoice line items, use an accountRef that includes the ID and name of the linked account.
	AccountRef *AccountRef       `json:"accountRef,omitempty"`
	ContactRef *ContactReference `json:"contactRef,omitempty"`
	// Currency for the journal line item.
	Currency *string `json:"currency,omitempty"`
	// Description of the journal line item.
	Description *string `json:"description,omitempty"`
	// Amount for the journal line. Debit entries are considered positive, and credit entries are considered negative.
	NetAmount *decimal.Big `decimal:"number" json:"netAmount"`
	// List of record refs associated with the tracking information for the line (eg to a Tracking Category, or customer etc.)
	Tracking *JournalLineTracking `json:"tracking,omitempty"`
}

func (j JournalLine) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JournalLine) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JournalLine) GetAccountRef() *AccountRef {
	if o == nil {
		return nil
	}
	return o.AccountRef
}

func (o *JournalLine) GetContactRef() *ContactReference {
	if o == nil {
		return nil
	}
	return o.ContactRef
}

func (o *JournalLine) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *JournalLine) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *JournalLine) GetNetAmount() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.NetAmount
}

func (o *JournalLine) GetTracking() *JournalLineTracking {
	if o == nil {
		return nil
	}
	return o.Tracking
}
