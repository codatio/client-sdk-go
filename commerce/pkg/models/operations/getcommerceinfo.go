package operations

import (
	"time"
)

type GetCommerceInfoPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCommerceInfoRequest struct {
	PathParams GetCommerceInfoPathParams
}

type GetCommerceInfoSourceModifiedDateAccountBalances struct {
	Available *float64 `json:"available,omitempty"`
	Currency  *string  `json:"currency,omitempty"`
	Pending   *float64 `json:"pending,omitempty"`
	Reserved  *float64 `json:"reserved,omitempty"`
}

type GetCommerceInfoSourceModifiedDateAddress struct {
	City       *string      `json:"city,omitempty"`
	Country    *string      `json:"country,omitempty"`
	Line1      *string      `json:"line1,omitempty"`
	Line2      *string      `json:"line2,omitempty"`
	PostalCode *string      `json:"postalCode,omitempty"`
	Region     *string      `json:"region,omitempty"`
	Type       *interface{} `json:"type,omitempty"`
}

type GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnum string

const (
	GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnumPrimary  GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnum = "Primary"
	GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnumLandline GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnum = "Landline"
	GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnumMobile   GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnum = "Mobile"
	GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnumFax      GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnum = "Fax"
	GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnumUnknown  GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnum = "Unknown"
)

type GetCommerceInfoSourceModifiedDatePhoneNumbers struct {
	Number string                                                `json:"number"`
	Type   GetCommerceInfoSourceModifiedDatePhoneNumbersTypeEnum `json:"type"`
}

// GetCommerceInfoSourceModifiedDate
// In the Codat system, company information includes standard commercial details about
// a linked company, such as their address, phone number, and company registration.
type GetCommerceInfoSourceModifiedDate struct {
	AccountBalances     []GetCommerceInfoSourceModifiedDateAccountBalances `json:"accountBalances,omitempty"`
	Addresses           []GetCommerceInfoSourceModifiedDateAddress         `json:"addresses,omitempty"`
	BaseCurrency        *string                                            `json:"baseCurrency,omitempty"`
	CommercePlatformRef *string                                            `json:"commercePlatformRef,omitempty"`
	CompanyLegalName    *string                                            `json:"companyLegalName,omitempty"`
	CompanyName         *string                                            `json:"companyName,omitempty"`
	CreatedDate         *time.Time                                         `json:"createdDate,omitempty"`
	ModifiedDate        *time.Time                                         `json:"modifiedDate,omitempty"`
	PhoneNumbers        []GetCommerceInfoSourceModifiedDatePhoneNumbers    `json:"phoneNumbers,omitempty"`
	RegistrationNumber  *string                                            `json:"registrationNumber,omitempty"`
	SourceModifiedDate  *time.Time                                         `json:"sourceModifiedDate,omitempty"`
	SourceUrls          map[string]string                                  `json:"sourceUrls,omitempty"`
	WebLinks            []interface{}                                      `json:"webLinks,omitempty"`
}

type GetCommerceInfoResponse struct {
	ContentType        string
	SourceModifiedDate *GetCommerceInfoSourceModifiedDate
	StatusCode         int
}
