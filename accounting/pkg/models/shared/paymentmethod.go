package shared

import (
	"time"
)

type PaymentMethodStatusEnum string

const (
	PaymentMethodStatusEnumUnknown  PaymentMethodStatusEnum = "Unknown"
	PaymentMethodStatusEnumActive   PaymentMethodStatusEnum = "Active"
	PaymentMethodStatusEnumArchived PaymentMethodStatusEnum = "Archived"
)

type PaymentMethodTypeEnum string

const (
	PaymentMethodTypeEnumUnknown      PaymentMethodTypeEnum = "Unknown"
	PaymentMethodTypeEnumCash         PaymentMethodTypeEnum = "Cash"
	PaymentMethodTypeEnumCheck        PaymentMethodTypeEnum = "Check"
	PaymentMethodTypeEnumCreditCard   PaymentMethodTypeEnum = "CreditCard"
	PaymentMethodTypeEnumDebitCard    PaymentMethodTypeEnum = "DebitCard"
	PaymentMethodTypeEnumBankTransfer PaymentMethodTypeEnum = "BankTransfer"
	PaymentMethodTypeEnumOther        PaymentMethodTypeEnum = "Other"
)

// PaymentMethod
// > View the coverage for payment methods in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=paymentMethods" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A Payment Method represents the payment method(s) used to pay a Bill. Payment Methods are referenced on [Bill Payments](https://docs.codat.io/docs/datamodel-accounting-billpayments) and [Payments](https://docs.codat.io/docs/datamodel-accounting-payments).
//
// From the Payment Methods endpoints you can retrieve:
//
//   - A list of all the Payment Methods used by a company: `GET/companies/{companyId}/data/paymentMethods`.
//   - The details of an individual Payment Method:
//     `GET /companies/{companyId}/data/paymentMethods/{paymentMethodId}`.
type PaymentMethod struct {
	ID                 *string                  `json:"id,omitempty"`
	Metadata           *Metadata                `json:"metadata,omitempty"`
	ModifiedDate       *time.Time               `json:"modifiedDate,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	SourceModifiedDate *time.Time               `json:"sourceModifiedDate,omitempty"`
	Status             *PaymentMethodStatusEnum `json:"status,omitempty"`
	Type               *PaymentMethodTypeEnum   `json:"type,omitempty"`
}
