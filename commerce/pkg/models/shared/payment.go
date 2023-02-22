package shared

import (
	"time"
)

type PaymentStatusEnum string

const (
	PaymentStatusEnumPending    PaymentStatusEnum = "Pending"
	PaymentStatusEnumAuthorized PaymentStatusEnum = "Authorized"
	PaymentStatusEnumPaid       PaymentStatusEnum = "Paid"
	PaymentStatusEnumFailed     PaymentStatusEnum = "Failed"
	PaymentStatusEnumCancelled  PaymentStatusEnum = "Cancelled"
	PaymentStatusEnumUnknown    PaymentStatusEnum = "Unknown"
)

// Payment
// A payment made in a commerce platform
type Payment struct {
	Amount             *float64           `json:"amount,omitempty"`
	CreatedDate        *time.Time         `json:"createdDate,omitempty"`
	Currency           *string            `json:"currency,omitempty"`
	DueDate            *time.Time         `json:"dueDate,omitempty"`
	ID                 string             `json:"id"`
	ModifiedDate       *time.Time         `json:"modifiedDate,omitempty"`
	PaymentMethodRef   *Zero              `json:"paymentMethodRef,omitempty"`
	PaymentProvider    *string            `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time         `json:"sourceModifiedDate,omitempty"`
	Status             *PaymentStatusEnum `json:"status,omitempty"`
}
