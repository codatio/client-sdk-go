// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreditNoteStatusEnum string

const (
	CreditNoteStatusEnumUnknown       CreditNoteStatusEnum = "Unknown"
	CreditNoteStatusEnumDraft         CreditNoteStatusEnum = "Draft"
	CreditNoteStatusEnumSubmitted     CreditNoteStatusEnum = "Submitted"
	CreditNoteStatusEnumPaid          CreditNoteStatusEnum = "Paid"
	CreditNoteStatusEnumVoid          CreditNoteStatusEnum = "Void"
	CreditNoteStatusEnumPartiallyPaid CreditNoteStatusEnum = "PartiallyPaid"
)

func (e CreditNoteStatusEnum) ToPointer() *CreditNoteStatusEnum {
	return &e
}

func (e *CreditNoteStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Draft":
		fallthrough
	case "Submitted":
		fallthrough
	case "Paid":
		fallthrough
	case "Void":
		fallthrough
	case "PartiallyPaid":
		*e = CreditNoteStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreditNoteStatusEnum: %v", v)
	}
}
