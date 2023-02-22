package shared

import (
	"time"
)

type CodatDataContractsDatasetsCommercePaymentRef struct {
	Amount             *float64                                             `json:"amount,omitempty"`
	CreatedDate        *time.Time                                           `json:"createdDate,omitempty"`
	Currency           *string                                              `json:"currency,omitempty"`
	DueDate            *time.Time                                           `json:"dueDate,omitempty"`
	ID                 *string                                              `json:"id,omitempty"`
	ModifiedDate       *time.Time                                           `json:"modifiedDate,omitempty"`
	PaymentProvider    *string                                              `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time                                           `json:"sourceModifiedDate,omitempty"`
	Status             *CodatDataContractsDatasetsCommercePaymentStatusEnum `json:"status,omitempty"`
	Type               *CodatDataContractsDatasetsCommercePaymentTypeEnum   `json:"type,omitempty"`
}
