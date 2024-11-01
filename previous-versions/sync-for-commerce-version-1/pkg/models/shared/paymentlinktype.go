// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PaymentLinkType - Types of payment line links, either:
// `Unknown`
// `Unlinked` - Not used
// `Invoice` - ID refers to the invoice
// `CreditNote` - ID refers to the credit note
// `Refund` - ID refers to the sibling payment
// `Payment` - ID refers to the sibling payment
// `PaymentOnAccount` - ID refers to the customer
// `Other` - ID refers to the customer
// `Manual Journal`
// `Discount` - ID refers to the payment
type PaymentLinkType string

const (
	PaymentLinkTypeUnknown          PaymentLinkType = "Unknown"
	PaymentLinkTypeUnlinked         PaymentLinkType = "Unlinked"
	PaymentLinkTypeInvoice          PaymentLinkType = "Invoice"
	PaymentLinkTypeCreditNote       PaymentLinkType = "CreditNote"
	PaymentLinkTypeOther            PaymentLinkType = "Other"
	PaymentLinkTypeRefund           PaymentLinkType = "Refund"
	PaymentLinkTypePayment          PaymentLinkType = "Payment"
	PaymentLinkTypePaymentOnAccount PaymentLinkType = "PaymentOnAccount"
	PaymentLinkTypeManualJournal    PaymentLinkType = "ManualJournal"
	PaymentLinkTypeDiscount         PaymentLinkType = "Discount"
)

func (e PaymentLinkType) ToPointer() *PaymentLinkType {
	return &e
}
func (e *PaymentLinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Unlinked":
		fallthrough
	case "Invoice":
		fallthrough
	case "CreditNote":
		fallthrough
	case "Other":
		fallthrough
	case "Refund":
		fallthrough
	case "Payment":
		fallthrough
	case "PaymentOnAccount":
		fallthrough
	case "ManualJournal":
		fallthrough
	case "Discount":
		*e = PaymentLinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentLinkType: %v", v)
	}
}
