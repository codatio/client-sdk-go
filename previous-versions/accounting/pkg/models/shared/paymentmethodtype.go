// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PaymentMethodType - Method of payment.
type PaymentMethodType string

const (
	PaymentMethodTypeUnknown      PaymentMethodType = "Unknown"
	PaymentMethodTypeCash         PaymentMethodType = "Cash"
	PaymentMethodTypeCheck        PaymentMethodType = "Check"
	PaymentMethodTypeCreditCard   PaymentMethodType = "CreditCard"
	PaymentMethodTypeDebitCard    PaymentMethodType = "DebitCard"
	PaymentMethodTypeBankTransfer PaymentMethodType = "BankTransfer"
	PaymentMethodTypeOther        PaymentMethodType = "Other"
)

func (e PaymentMethodType) ToPointer() *PaymentMethodType {
	return &e
}
func (e *PaymentMethodType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Cash":
		fallthrough
	case "Check":
		fallthrough
	case "CreditCard":
		fallthrough
	case "DebitCard":
		fallthrough
	case "BankTransfer":
		fallthrough
	case "Other":
		*e = PaymentMethodType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodType: %v", v)
	}
}
