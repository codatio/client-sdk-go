package shared

import (
	"time"
)

type PaymentMethodStatusEnum string

const (
	PaymentMethodStatusEnumActive   PaymentMethodStatusEnum = "Active"
	PaymentMethodStatusEnumArchived PaymentMethodStatusEnum = "Archived"
	PaymentMethodStatusEnumUnknown  PaymentMethodStatusEnum = "Unknown"
)

type PaymentMethod struct {
	ID                 string                   `json:"id"`
	ModifiedDate       *time.Time               `json:"modifiedDate,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	SourceModifiedDate *time.Time               `json:"sourceModifiedDate,omitempty"`
	Status             *PaymentMethodStatusEnum `json:"status,omitempty"`
}
