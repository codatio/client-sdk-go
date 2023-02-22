package shared

import (
	"time"
)

type CompanyInfoAccountBalances struct {
	Available *float64 `json:"available,omitempty"`
	Currency  *string  `json:"currency,omitempty"`
	Pending   *float64 `json:"pending,omitempty"`
	Reserved  *float64 `json:"reserved,omitempty"`
}

type CompanyInfoAddress struct {
	City       *string      `json:"city,omitempty"`
	Country    *string      `json:"country,omitempty"`
	Line1      *string      `json:"line1,omitempty"`
	Line2      *string      `json:"line2,omitempty"`
	PostalCode *string      `json:"postalCode,omitempty"`
	Region     *string      `json:"region,omitempty"`
	Type       *interface{} `json:"type,omitempty"`
}

type CompanyInfoPhoneNumbersTypeEnum string

const (
	CompanyInfoPhoneNumbersTypeEnumPrimary  CompanyInfoPhoneNumbersTypeEnum = "Primary"
	CompanyInfoPhoneNumbersTypeEnumLandline CompanyInfoPhoneNumbersTypeEnum = "Landline"
	CompanyInfoPhoneNumbersTypeEnumMobile   CompanyInfoPhoneNumbersTypeEnum = "Mobile"
	CompanyInfoPhoneNumbersTypeEnumFax      CompanyInfoPhoneNumbersTypeEnum = "Fax"
	CompanyInfoPhoneNumbersTypeEnumUnknown  CompanyInfoPhoneNumbersTypeEnum = "Unknown"
)

type CompanyInfoPhoneNumbers struct {
	Number string                          `json:"number"`
	Type   CompanyInfoPhoneNumbersTypeEnum `json:"type"`
}

// CompanyInfo
// In the Codat system, company information includes standard commercial details about
// a linked company, such as their address, phone number, and company registration.
type CompanyInfo struct {
	AccountBalances     []CompanyInfoAccountBalances `json:"accountBalances,omitempty"`
	Addresses           []CompanyInfoAddress         `json:"addresses,omitempty"`
	BaseCurrency        *string                      `json:"baseCurrency,omitempty"`
	CommercePlatformRef *string                      `json:"commercePlatformRef,omitempty"`
	CompanyLegalName    *string                      `json:"companyLegalName,omitempty"`
	CompanyName         *string                      `json:"companyName,omitempty"`
	CreatedDate         *time.Time                   `json:"createdDate,omitempty"`
	ModifiedDate        *time.Time                   `json:"modifiedDate,omitempty"`
	PhoneNumbers        []CompanyInfoPhoneNumbers    `json:"phoneNumbers,omitempty"`
	RegistrationNumber  *string                      `json:"registrationNumber,omitempty"`
	SourceModifiedDate  *time.Time                   `json:"sourceModifiedDate,omitempty"`
	SourceUrls          map[string]string            `json:"sourceUrls,omitempty"`
	WebLinks            []interface{}                `json:"webLinks,omitempty"`
}
