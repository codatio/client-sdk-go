// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CompanyInfoWeblinkType - The type of the weblink.
type CompanyInfoWeblinkType string

const (
	CompanyInfoWeblinkTypeWebsite CompanyInfoWeblinkType = "Website"
	CompanyInfoWeblinkTypeSocial  CompanyInfoWeblinkType = "Social"
	CompanyInfoWeblinkTypeUnknown CompanyInfoWeblinkType = "Unknown"
)

func (e CompanyInfoWeblinkType) ToPointer() *CompanyInfoWeblinkType {
	return &e
}

func (e *CompanyInfoWeblinkType) UnmarshalJSON(data []byte) error {
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
		*e = CompanyInfoWeblinkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CompanyInfoWeblinkType: %v", v)
	}
}

// CompanyInfoWeblink - Weblink associated with the company.
type CompanyInfoWeblink struct {
	// The type of the weblink.
	Type *CompanyInfoWeblinkType `json:"type,omitempty"`
	// The full URL for the weblink.
	URL *string `json:"url,omitempty"`
}

func (o *CompanyInfoWeblink) GetType() *CompanyInfoWeblinkType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CompanyInfoWeblink) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// CompanyInfo - In the Codat system, company information includes standard commercial details about
// a linked company, such as their address, phone number, and company registration.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=companyInfo) for this data type.
type CompanyInfo struct {
	// The available and current cash balances for the company's accounts
	AccountBalances []AccountBalance `json:"accountBalances,omitempty"`
	// Addresses associated with the company
	Addresses []Address `json:"addresses,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	BaseCurrency *string `json:"baseCurrency,omitempty"`
	// Identifier or reference for the company in the commerce platform
	CommercePlatformRef *string `json:"commercePlatformRef,omitempty"`
	// The full legal name of the company
	CompanyLegalName *string `json:"companyLegalName,omitempty"`
	// The name of the company
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
	CreatedDate  *string `json:"createdDate,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Phone numbers associated with the company
	PhoneNumbers []PhoneNumber `json:"phoneNumbers,omitempty"`
	// The registration number of the company
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// URL addresses for the originating system. For example, potential use cases include 'deeplinking' to the originating system
	SourceUrls map[string]string `json:"sourceUrls,omitempty"`
	// Weblinks associated with the company
	WebLinks []CompanyInfoWeblink `json:"webLinks,omitempty"`
}

func (o *CompanyInfo) GetAccountBalances() []AccountBalance {
	if o == nil {
		return nil
	}
	return o.AccountBalances
}

func (o *CompanyInfo) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *CompanyInfo) GetBaseCurrency() *string {
	if o == nil {
		return nil
	}
	return o.BaseCurrency
}

func (o *CompanyInfo) GetCommercePlatformRef() *string {
	if o == nil {
		return nil
	}
	return o.CommercePlatformRef
}

func (o *CompanyInfo) GetCompanyLegalName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyLegalName
}

func (o *CompanyInfo) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *CompanyInfo) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *CompanyInfo) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *CompanyInfo) GetPhoneNumbers() []PhoneNumber {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *CompanyInfo) GetRegistrationNumber() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationNumber
}

func (o *CompanyInfo) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *CompanyInfo) GetSourceUrls() map[string]string {
	if o == nil {
		return nil
	}
	return o.SourceUrls
}

func (o *CompanyInfo) GetWebLinks() []CompanyInfoWeblink {
	if o == nil {
		return nil
	}
	return o.WebLinks
}