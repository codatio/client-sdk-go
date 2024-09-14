// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountingAddress struct {
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

func (o *AccountingAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *AccountingAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *AccountingAddress) GetLine1() *string {
	if o == nil {
		return nil
	}
	return o.Line1
}

func (o *AccountingAddress) GetLine2() *string {
	if o == nil {
		return nil
	}
	return o.Line2
}

func (o *AccountingAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *AccountingAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *AccountingAddress) GetType() AccountingAddressType {
	if o == nil {
		return AccountingAddressType("")
	}
	return o.Type
}

type Phone struct {
	// A phone number.
	Number *string `json:"number,omitempty"`
	// The type of phone number
	Type PhoneNumberType `json:"type"`
}

func (o *Phone) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *Phone) GetType() PhoneNumberType {
	if o == nil {
		return PhoneNumberType("")
	}
	return o.Type
}

// AccountingCompanyDatasetType - The type of the weblink.
type AccountingCompanyDatasetType string

const (
	AccountingCompanyDatasetTypeWebsite AccountingCompanyDatasetType = "Website"
	AccountingCompanyDatasetTypeSocial  AccountingCompanyDatasetType = "Social"
	AccountingCompanyDatasetTypeUnknown AccountingCompanyDatasetType = "Unknown"
)

func (e AccountingCompanyDatasetType) ToPointer() *AccountingCompanyDatasetType {
	return &e
}
func (e *AccountingCompanyDatasetType) UnmarshalJSON(data []byte) error {
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
		*e = AccountingCompanyDatasetType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingCompanyDatasetType: %v", v)
	}
}

// Weblink associated with the company.
type Weblink struct {
	// The type of the weblink.
	Type *AccountingCompanyDatasetType `json:"type,omitempty"`
	// The full URL for the weblink.
	URL *string `json:"url,omitempty"`
}

func (o *Weblink) GetType() *AccountingCompanyDatasetType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Weblink) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// AccountingCompanyDataset - > View the coverage for company profile in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=cashFlowStatement" target="_blank">Data coverage explorer</a>.
//
// Company info provides standard details about a linked company such as their address, phone number, and company registration.
//
// > **Company information or companies?**
// >
// > Company profile is standard information that is held in the accounting software about a company. `Companies` is an endpoint that lists businesses in the Codat system that have linked and shared their data sources.
type AccountingCompanyDataset struct {
	// Identifier or reference for the company in the accounting software.
	AccountingPlatformRef *string `json:"accountingPlatformRef,omitempty"`
	// An array of Addresses.
	Addresses []AccountingAddress `json:"addresses,omitempty"`
	// Currency set in the accounting software of the linked company. Used by the currency rate.
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
	PhoneNumbers []Phone `json:"phoneNumbers,omitempty"`
	// Registration number given to the linked company by the companies authority in the country of origin. In the UK this is Companies House.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// URL addresses for the accounting source.
	//
	// For example, for Xero integrations two URLs are returned. These have many potential use cases, such as [deep linking](https://developer.xero.com/documentation/api-guides/deep-link-xero).
	SourceUrls map[string]string `json:"sourceUrls,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
	// An array of weblinks.
	WebLinks []Weblink `json:"webLinks,omitempty"`
}

func (o *AccountingCompanyDataset) GetAccountingPlatformRef() *string {
	if o == nil {
		return nil
	}
	return o.AccountingPlatformRef
}

func (o *AccountingCompanyDataset) GetAddresses() []AccountingAddress {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *AccountingCompanyDataset) GetBaseCurrency() *string {
	if o == nil {
		return nil
	}
	return o.BaseCurrency
}

func (o *AccountingCompanyDataset) GetCompanyLegalName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyLegalName
}

func (o *AccountingCompanyDataset) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *AccountingCompanyDataset) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *AccountingCompanyDataset) GetFinancialYearStartDate() *string {
	if o == nil {
		return nil
	}
	return o.FinancialYearStartDate
}

func (o *AccountingCompanyDataset) GetLedgerLockDate() *string {
	if o == nil {
		return nil
	}
	return o.LedgerLockDate
}

func (o *AccountingCompanyDataset) GetPhoneNumbers() []Phone {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *AccountingCompanyDataset) GetRegistrationNumber() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationNumber
}

func (o *AccountingCompanyDataset) GetSourceUrls() map[string]string {
	if o == nil {
		return nil
	}
	return o.SourceUrls
}

func (o *AccountingCompanyDataset) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *AccountingCompanyDataset) GetTaxNumber() *string {
	if o == nil {
		return nil
	}
	return o.TaxNumber
}

func (o *AccountingCompanyDataset) GetWebLinks() []Weblink {
	if o == nil {
		return nil
	}
	return o.WebLinks
}
