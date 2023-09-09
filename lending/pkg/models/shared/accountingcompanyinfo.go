// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountingCompanyInfoAccountingAddress struct {
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
	Type AccountingAddressType `json:"type"`
}

func (o *AccountingCompanyInfoAccountingAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *AccountingCompanyInfoAccountingAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *AccountingCompanyInfoAccountingAddress) GetLine1() *string {
	if o == nil {
		return nil
	}
	return o.Line1
}

func (o *AccountingCompanyInfoAccountingAddress) GetLine2() *string {
	if o == nil {
		return nil
	}
	return o.Line2
}

func (o *AccountingCompanyInfoAccountingAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *AccountingCompanyInfoAccountingAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *AccountingCompanyInfoAccountingAddress) GetType() AccountingAddressType {
	if o == nil {
		return AccountingAddressType("")
	}
	return o.Type
}

type AccountingCompanyInfoPhone struct {
	// A phone number.
	Number *string `json:"number"`
	// The type of phone number
	Type PhoneNumberType `json:"type"`
}

func (o *AccountingCompanyInfoPhone) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *AccountingCompanyInfoPhone) GetType() PhoneNumberType {
	if o == nil {
		return PhoneNumberType("")
	}
	return o.Type
}

// AccountingCompanyInfoWeblinkType - The type of the weblink.
type AccountingCompanyInfoWeblinkType string

const (
	AccountingCompanyInfoWeblinkTypeWebsite AccountingCompanyInfoWeblinkType = "Website"
	AccountingCompanyInfoWeblinkTypeSocial  AccountingCompanyInfoWeblinkType = "Social"
	AccountingCompanyInfoWeblinkTypeUnknown AccountingCompanyInfoWeblinkType = "Unknown"
)

func (e AccountingCompanyInfoWeblinkType) ToPointer() *AccountingCompanyInfoWeblinkType {
	return &e
}

func (e *AccountingCompanyInfoWeblinkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Website":
		fallthrough
	case "Social":
		fallthrough
	case "Unknown":
		*e = AccountingCompanyInfoWeblinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingCompanyInfoWeblinkType: %v", v)
	}
}

// AccountingCompanyInfoWeblink - Weblink associated with the company.
type AccountingCompanyInfoWeblink struct {
	// The type of the weblink.
	Type *AccountingCompanyInfoWeblinkType `json:"type,omitempty"`
	// The full URL for the weblink.
	URL *string `json:"url,omitempty"`
}

func (o *AccountingCompanyInfoWeblink) GetType() *AccountingCompanyInfoWeblinkType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AccountingCompanyInfoWeblink) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// AccountingCompanyInfo - > View the coverage for company info in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=cashFlowStatement" target="_blank">Data coverage explorer</a>.
//
// Company info provides standard details about a linked company such as their address, phone number, and company registration.
//
// > **Company information or companies?**
// >
// > Company information is standard information that is held in the accounting platform about a company. `Companies` is an endpoint that lists businesses in the Codat system that have linked and shared their data sources.
type AccountingCompanyInfo struct {
	// Identifier or reference for the company in the accounting platform.
	AccountingPlatformRef *string `json:"accountingPlatformRef,omitempty"`
	// An array of Addresses.
	Addresses []AccountingCompanyInfoAccountingAddress `json:"addresses,omitempty"`
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
	PhoneNumbers []AccountingCompanyInfoPhone `json:"phoneNumbers,omitempty"`
	// Registration number given to the linked company by the companies authority in the country of origin. In the UK this is Companies House.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// URL addresses for the accounting source.
	//
	// For example, for Xero integrations two URLs are returned. These have many potential use cases, such as [deep linking](https://developer.xero.com/documentation/api-guides/deep-link-xero).
	SourceUrls map[string]string `json:"sourceUrls,omitempty"`
	// Company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
	// An array of weblinks.
	WebLinks []AccountingCompanyInfoWeblink `json:"webLinks,omitempty"`
}

func (o *AccountingCompanyInfo) GetAccountingPlatformRef() *string {
	if o == nil {
		return nil
	}
	return o.AccountingPlatformRef
}

func (o *AccountingCompanyInfo) GetAddresses() []AccountingCompanyInfoAccountingAddress {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *AccountingCompanyInfo) GetBaseCurrency() *string {
	if o == nil {
		return nil
	}
	return o.BaseCurrency
}

func (o *AccountingCompanyInfo) GetCompanyLegalName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyLegalName
}

func (o *AccountingCompanyInfo) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *AccountingCompanyInfo) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *AccountingCompanyInfo) GetFinancialYearStartDate() *string {
	if o == nil {
		return nil
	}
	return o.FinancialYearStartDate
}

func (o *AccountingCompanyInfo) GetLedgerLockDate() *string {
	if o == nil {
		return nil
	}
	return o.LedgerLockDate
}

func (o *AccountingCompanyInfo) GetPhoneNumbers() []AccountingCompanyInfoPhone {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *AccountingCompanyInfo) GetRegistrationNumber() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationNumber
}

func (o *AccountingCompanyInfo) GetSourceUrls() map[string]string {
	if o == nil {
		return nil
	}
	return o.SourceUrls
}

func (o *AccountingCompanyInfo) GetTaxNumber() *string {
	if o == nil {
		return nil
	}
	return o.TaxNumber
}

func (o *AccountingCompanyInfo) GetWebLinks() []AccountingCompanyInfoWeblink {
	if o == nil {
		return nil
	}
	return o.WebLinks
}
