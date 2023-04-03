// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CompanyDatasetAddresses struct {
	// City of the customer address.
	City *string `json:"city,omitempty"`
	// Country of the customer address.
	Country *string `json:"country,omitempty"`
	// Line 1 of the customer address.
	Line1 *string `json:"line1,omitempty"`
	// Line 2 of the customer address.
	Line2 *string `json:"line2,omitempty"`
	// Postal code or zip code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Region of the customer address.
	Region *string `json:"region,omitempty"`
	// The type of the address
	Type AddressTypeEnum `json:"type"`
}

type CompanyDatasetPhone struct {
	// A phone number.
	Number string `json:"number"`
	// The type of phone number
	Type PhoneNumberTypeEnum `json:"type"`
}

// CompanyDatasetWeblinkTypeEnum - The type of the weblink.
type CompanyDatasetWeblinkTypeEnum string

const (
	CompanyDatasetWeblinkTypeEnumWebsite CompanyDatasetWeblinkTypeEnum = "Website"
	CompanyDatasetWeblinkTypeEnumSocial  CompanyDatasetWeblinkTypeEnum = "Social"
	CompanyDatasetWeblinkTypeEnumUnknown CompanyDatasetWeblinkTypeEnum = "Unknown"
)

func (e *CompanyDatasetWeblinkTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Website":
		fallthrough
	case "Social":
		fallthrough
	case "Unknown":
		*e = CompanyDatasetWeblinkTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CompanyDatasetWeblinkTypeEnum: %s", s)
	}
}

// CompanyDatasetWeblink - Weblink associated with the company.
type CompanyDatasetWeblink struct {
	// The type of the weblink.
	Type *CompanyDatasetWeblinkTypeEnum `json:"type,omitempty"`
	// The full URL for the weblink.
	URL *string `json:"url,omitempty"`
}

// CompanyDataset - > View the coverage for company info in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=cashFlowStatement" target="_blank">Data coverage explorer</a>.
//
// Company info provides standard details about a linked company such as their address, phone number, and company registration.
//
// > **Company information or companies?**
// >
// > Company information is standard information that is held in the accounting platform about a company. `Companies` is an endpoint that lists businesses in the Codat system that have linked and shared their data sources.
type CompanyDataset struct {
	// Identifier or reference for the company in the accounting platform.
	AccountingPlatformRef *string `json:"accountingPlatformRef,omitempty"`
	// An array of Addresses.
	Addresses []CompanyDatasetAddresses `json:"addresses,omitempty"`
	// Currency set in the accounting platform of the linked company. Used by the currency rate.
	BaseCurrency *string `json:"baseCurrency,omitempty"`
	// Registered legal name of the linked company.
	CompanyLegalName *string `json:"companyLegalName,omitempty"`
	// Name of the linked company.
	CompanyName *string `json:"companyName,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	CreatedDate *string `json:"createdDate,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	FinancialYearStartDate *string `json:"financialYearStartDate,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	LedgerLockDate *string `json:"ledgerLockDate,omitempty"`
	// An array of phone numbers.
	PhoneNumbers []CompanyDatasetPhone `json:"phoneNumbers,omitempty"`
	// Registration number given to the linked company by the companies authority in the country of origin. In the UK this is Companies House.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// URL addresses for the accounting source.
	//
	// For example, for Xero integrations two URLs are returned. These have many potential use cases, such as [deep linking](https://developer.xero.com/documentation/api-guides/deep-link-xero).
	SourceUrls map[string]string `json:"sourceUrls,omitempty"`
	// Company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
	// An array of weblinks.
	WebLinks []CompanyDatasetWeblink `json:"webLinks,omitempty"`
}
