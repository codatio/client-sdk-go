package shared

import (
	"time"
)

type PaymentRefStatusEnum string

const (
	PaymentRefStatusEnumPending    PaymentRefStatusEnum = "Pending"
	PaymentRefStatusEnumAuthorized PaymentRefStatusEnum = "Authorized"
	PaymentRefStatusEnumPaid       PaymentRefStatusEnum = "Paid"
	PaymentRefStatusEnumFailed     PaymentRefStatusEnum = "Failed"
	PaymentRefStatusEnumCancelled  PaymentRefStatusEnum = "Cancelled"
	PaymentRefStatusEnumUnknown    PaymentRefStatusEnum = "Unknown"
)

type PaymentRefTypeEnum string

const (
	PaymentRefTypeEnumCash        PaymentRefTypeEnum = "Cash"
	PaymentRefTypeEnumCard        PaymentRefTypeEnum = "Card"
	PaymentRefTypeEnumInvoice     PaymentRefTypeEnum = "Invoice"
	PaymentRefTypeEnumOnlineCard  PaymentRefTypeEnum = "OnlineCard"
	PaymentRefTypeEnumSwish       PaymentRefTypeEnum = "Swish"
	PaymentRefTypeEnumVipps       PaymentRefTypeEnum = "Vipps"
	PaymentRefTypeEnumMobile      PaymentRefTypeEnum = "Mobile"
	PaymentRefTypeEnumStoreCredit PaymentRefTypeEnum = "StoreCredit"
	PaymentRefTypeEnumPaypal      PaymentRefTypeEnum = "Paypal"
	PaymentRefTypeEnumCustom      PaymentRefTypeEnum = "Custom"
	PaymentRefTypeEnumPrepaid     PaymentRefTypeEnum = "Prepaid"
	PaymentRefTypeEnumUnknown     PaymentRefTypeEnum = "Unknown"
)

type PaymentRef struct {
	Amount             *float64              `json:"amount,omitempty"`
	CreatedDate        *time.Time            `json:"createdDate,omitempty"`
	Currency           *string               `json:"currency,omitempty"`
	DueDate            *time.Time            `json:"dueDate,omitempty"`
	ID                 string                `json:"id"`
	ModifiedDate       *time.Time            `json:"modifiedDate,omitempty"`
	PaymentProvider    *string               `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time            `json:"sourceModifiedDate,omitempty"`
	Status             *PaymentRefStatusEnum `json:"status,omitempty"`
	Type               *PaymentRefTypeEnum   `json:"type,omitempty"`
}
