// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// BillPaymentAccountRef - Reference to the bank account / credit card which you are using to pay the bill.
type BillPaymentAccountRef struct {
	// Unique ID of the bank account / credit card
	ID string `json:"id"`
}

func (o *BillPaymentAccountRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
