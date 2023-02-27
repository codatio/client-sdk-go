package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetPaymentMethodPathParams struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	PaymentMethodID string `pathParam:"style=simple,explode=false,name=paymentMethodId"`
}

type GetPaymentMethodSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetPaymentMethodRequest struct {
	PathParams GetPaymentMethodPathParams
	Security   GetPaymentMethodSecurity
}

type GetPaymentMethodSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetPaymentMethodSourceModifiedDateStatusEnum string

const (
	GetPaymentMethodSourceModifiedDateStatusEnumUnknown  GetPaymentMethodSourceModifiedDateStatusEnum = "Unknown"
	GetPaymentMethodSourceModifiedDateStatusEnumActive   GetPaymentMethodSourceModifiedDateStatusEnum = "Active"
	GetPaymentMethodSourceModifiedDateStatusEnumArchived GetPaymentMethodSourceModifiedDateStatusEnum = "Archived"
)

type GetPaymentMethodSourceModifiedDateTypeEnum string

const (
	GetPaymentMethodSourceModifiedDateTypeEnumUnknown      GetPaymentMethodSourceModifiedDateTypeEnum = "Unknown"
	GetPaymentMethodSourceModifiedDateTypeEnumCash         GetPaymentMethodSourceModifiedDateTypeEnum = "Cash"
	GetPaymentMethodSourceModifiedDateTypeEnumCheck        GetPaymentMethodSourceModifiedDateTypeEnum = "Check"
	GetPaymentMethodSourceModifiedDateTypeEnumCreditCard   GetPaymentMethodSourceModifiedDateTypeEnum = "CreditCard"
	GetPaymentMethodSourceModifiedDateTypeEnumDebitCard    GetPaymentMethodSourceModifiedDateTypeEnum = "DebitCard"
	GetPaymentMethodSourceModifiedDateTypeEnumBankTransfer GetPaymentMethodSourceModifiedDateTypeEnum = "BankTransfer"
	GetPaymentMethodSourceModifiedDateTypeEnumOther        GetPaymentMethodSourceModifiedDateTypeEnum = "Other"
)

// GetPaymentMethodSourceModifiedDate
// > View the coverage for payment methods in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=paymentMethods" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A Payment Method represents the payment method(s) used to pay a Bill. Payment Methods are referenced on [Bill Payments](https://docs.codat.io/accounting-api#/schemas/BillPayment) and [Payments](https://docs.codat.io/accounting-api#/schemas/Payment).
//
// From the Payment Methods endpoints you can retrieve:
//
//   - A list of all the Payment Methods used by a company: `GET/companies/{companyId}/data/paymentMethods`.
//   - The details of an individual Payment Method:
//     `GET /companies/{companyId}/data/paymentMethods/{paymentMethodId}`.
type GetPaymentMethodSourceModifiedDate struct {
	ID                 *string                                       `json:"id,omitempty"`
	Metadata           *GetPaymentMethodSourceModifiedDateMetadata   `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                    `json:"modifiedDate,omitempty"`
	Name               *string                                       `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                    `json:"sourceModifiedDate,omitempty"`
	Status             *GetPaymentMethodSourceModifiedDateStatusEnum `json:"status,omitempty"`
	Type               *GetPaymentMethodSourceModifiedDateTypeEnum   `json:"type,omitempty"`
}

type GetPaymentMethodResponse struct {
	ContentType        string
	SourceModifiedDate *GetPaymentMethodSourceModifiedDate
	StatusCode         int
}
