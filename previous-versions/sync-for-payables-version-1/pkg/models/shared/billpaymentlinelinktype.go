// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BillPaymentLineLinkType - Types of links to bill payment lines.
type BillPaymentLineLinkType string

const (
	BillPaymentLineLinkTypeUnknown          BillPaymentLineLinkType = "Unknown"
	BillPaymentLineLinkTypeUnlinked         BillPaymentLineLinkType = "Unlinked"
	BillPaymentLineLinkTypeBill             BillPaymentLineLinkType = "Bill"
	BillPaymentLineLinkTypeOther            BillPaymentLineLinkType = "Other"
	BillPaymentLineLinkTypeCreditNote       BillPaymentLineLinkType = "CreditNote"
	BillPaymentLineLinkTypeBillPayment      BillPaymentLineLinkType = "BillPayment"
	BillPaymentLineLinkTypePaymentOnAccount BillPaymentLineLinkType = "PaymentOnAccount"
	BillPaymentLineLinkTypeRefund           BillPaymentLineLinkType = "Refund"
	BillPaymentLineLinkTypeManualJournal    BillPaymentLineLinkType = "ManualJournal"
	BillPaymentLineLinkTypeDiscount         BillPaymentLineLinkType = "Discount"
)

func (e BillPaymentLineLinkType) ToPointer() *BillPaymentLineLinkType {
	return &e
}
func (e *BillPaymentLineLinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Unlinked":
		fallthrough
	case "Bill":
		fallthrough
	case "Other":
		fallthrough
	case "CreditNote":
		fallthrough
	case "BillPayment":
		fallthrough
	case "PaymentOnAccount":
		fallthrough
	case "Refund":
		fallthrough
	case "ManualJournal":
		fallthrough
	case "Discount":
		*e = BillPaymentLineLinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillPaymentLineLinkType: %v", v)
	}
}
