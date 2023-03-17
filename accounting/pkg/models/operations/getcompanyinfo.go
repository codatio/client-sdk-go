package operations

import (
	"net/http"
	"time"
)

type GetCompanyInfoRequest struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyInfoCompanyInfoAddressesTypeEnum string

const (
	GetCompanyInfoCompanyInfoAddressesTypeEnumUnknown  GetCompanyInfoCompanyInfoAddressesTypeEnum = "Unknown"
	GetCompanyInfoCompanyInfoAddressesTypeEnumBilling  GetCompanyInfoCompanyInfoAddressesTypeEnum = "Billing"
	GetCompanyInfoCompanyInfoAddressesTypeEnumDelivery GetCompanyInfoCompanyInfoAddressesTypeEnum = "Delivery"
)

type GetCompanyInfoCompanyInfoAddresses struct {
	City       *string                                    `json:"city,omitempty"`
	Country    *string                                    `json:"country,omitempty"`
	Line1      *string                                    `json:"line1,omitempty"`
	Line2      *string                                    `json:"line2,omitempty"`
	PostalCode *string                                    `json:"postalCode,omitempty"`
	Region     *string                                    `json:"region,omitempty"`
	Type       GetCompanyInfoCompanyInfoAddressesTypeEnum `json:"type"`
}

type GetCompanyInfoCompanyInfoPhoneNumbersTypeEnum string

const (
	GetCompanyInfoCompanyInfoPhoneNumbersTypeEnumUnknown  GetCompanyInfoCompanyInfoPhoneNumbersTypeEnum = "Unknown"
	GetCompanyInfoCompanyInfoPhoneNumbersTypeEnumPrimary  GetCompanyInfoCompanyInfoPhoneNumbersTypeEnum = "Primary"
	GetCompanyInfoCompanyInfoPhoneNumbersTypeEnumLandline GetCompanyInfoCompanyInfoPhoneNumbersTypeEnum = "Landline"
	GetCompanyInfoCompanyInfoPhoneNumbersTypeEnumMobile   GetCompanyInfoCompanyInfoPhoneNumbersTypeEnum = "Mobile"
	GetCompanyInfoCompanyInfoPhoneNumbersTypeEnumFax      GetCompanyInfoCompanyInfoPhoneNumbersTypeEnum = "Fax"
)

type GetCompanyInfoCompanyInfoPhoneNumbers struct {
	Number *string                                       `json:"number,omitempty"`
	Type   GetCompanyInfoCompanyInfoPhoneNumbersTypeEnum `json:"type"`
}

type GetCompanyInfoCompanyInfoWebLinksTypeEnum string

const (
	GetCompanyInfoCompanyInfoWebLinksTypeEnumUnknown GetCompanyInfoCompanyInfoWebLinksTypeEnum = "Unknown"
	GetCompanyInfoCompanyInfoWebLinksTypeEnumWebsite GetCompanyInfoCompanyInfoWebLinksTypeEnum = "Website"
	GetCompanyInfoCompanyInfoWebLinksTypeEnumSocial  GetCompanyInfoCompanyInfoWebLinksTypeEnum = "Social"
)

type GetCompanyInfoCompanyInfoWebLinks struct {
	Type GetCompanyInfoCompanyInfoWebLinksTypeEnum `json:"type"`
	URL  *string                                   `json:"url,omitempty"`
}

type GetCompanyInfoCompanyInfo struct {
	AccountingPlatformRef  *string                                 `json:"accountingPlatformRef,omitempty"`
	Addresses              []GetCompanyInfoCompanyInfoAddresses    `json:"addresses,omitempty"`
	BaseCurrency           *string                                 `json:"baseCurrency,omitempty"`
	CompanyLegalName       *string                                 `json:"companyLegalName,omitempty"`
	CompanyName            *string                                 `json:"companyName,omitempty"`
	CreatedDate            *time.Time                              `json:"createdDate,omitempty"`
	FinancialYearStartDate *time.Time                              `json:"financialYearStartDate,omitempty"`
	LedgerLockDate         *time.Time                              `json:"ledgerLockDate,omitempty"`
	PhoneNumbers           []GetCompanyInfoCompanyInfoPhoneNumbers `json:"phoneNumbers,omitempty"`
	RegistrationNumber     *string                                 `json:"registrationNumber,omitempty"`
	SourceUrls             map[string]string                       `json:"sourceUrls,omitempty"`
	TaxNumber              *string                                 `json:"taxNumber,omitempty"`
	WebLinks               []GetCompanyInfoCompanyInfoWebLinks     `json:"webLinks,omitempty"`
}

type GetCompanyInfoResponse struct {
	CompanyInfo *GetCompanyInfoCompanyInfo
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
