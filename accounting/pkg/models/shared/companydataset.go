package shared

import (
	"time"
)

type CompanyDatasetPhoneNumbersTypeEnum string

const (
	CompanyDatasetPhoneNumbersTypeEnumUnknown  CompanyDatasetPhoneNumbersTypeEnum = "Unknown"
	CompanyDatasetPhoneNumbersTypeEnumPrimary  CompanyDatasetPhoneNumbersTypeEnum = "Primary"
	CompanyDatasetPhoneNumbersTypeEnumLandline CompanyDatasetPhoneNumbersTypeEnum = "Landline"
	CompanyDatasetPhoneNumbersTypeEnumMobile   CompanyDatasetPhoneNumbersTypeEnum = "Mobile"
	CompanyDatasetPhoneNumbersTypeEnumFax      CompanyDatasetPhoneNumbersTypeEnum = "Fax"
)

type CompanyDatasetPhoneNumbers struct {
	Number *string                            `json:"number,omitempty"`
	Type   CompanyDatasetPhoneNumbersTypeEnum `json:"type"`
}

type CompanyDatasetWebLinksTypeEnum string

const (
	CompanyDatasetWebLinksTypeEnumUnknown CompanyDatasetWebLinksTypeEnum = "Unknown"
	CompanyDatasetWebLinksTypeEnumWebsite CompanyDatasetWebLinksTypeEnum = "Website"
	CompanyDatasetWebLinksTypeEnumSocial  CompanyDatasetWebLinksTypeEnum = "Social"
)

type CompanyDatasetWebLinks struct {
	Type CompanyDatasetWebLinksTypeEnum `json:"type"`
	URL  *string                        `json:"url,omitempty"`
}

type CompanyDataset struct {
	AccountingPlatformRef  *string                      `json:"accountingPlatformRef,omitempty"`
	Addresses              []Addressesitems             `json:"addresses,omitempty"`
	BaseCurrency           *string                      `json:"baseCurrency,omitempty"`
	CompanyLegalName       *string                      `json:"companyLegalName,omitempty"`
	CompanyName            *string                      `json:"companyName,omitempty"`
	CreatedDate            *time.Time                   `json:"createdDate,omitempty"`
	FinancialYearStartDate *time.Time                   `json:"financialYearStartDate,omitempty"`
	LedgerLockDate         *time.Time                   `json:"ledgerLockDate,omitempty"`
	PhoneNumbers           []CompanyDatasetPhoneNumbers `json:"phoneNumbers,omitempty"`
	RegistrationNumber     *string                      `json:"registrationNumber,omitempty"`
	SourceUrls             map[string]string            `json:"sourceUrls,omitempty"`
	TaxNumber              *string                      `json:"taxNumber,omitempty"`
	WebLinks               []CompanyDatasetWebLinks     `json:"webLinks,omitempty"`
}
