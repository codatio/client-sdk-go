// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CommerceCompanyInfo - In the Codat system, company information includes standard commercial details about
// a linked company, such as their address, phone number, and company registration.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=companyInfo) for this data type.
type CommerceCompanyInfo struct {
	// The available and current cash balances for the company's accounts
	AccountBalances []AccountBalance `json:"accountBalances,omitempty"`
	// Addresses associated with the company
	Addresses []CommerceAddress `json:"addresses,omitempty"`
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
	PhoneNumbers []Items `json:"phoneNumbers,omitempty"`
	// The registration number of the company
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// URL addresses for the originating system. For example, potential use cases include 'deeplinking' to the originating system
	SourceUrls map[string]string `json:"sourceUrls,omitempty"`
	// Weblinks associated with the company
	WebLinks []WebLinksitems `json:"webLinks,omitempty"`
}

func (o *CommerceCompanyInfo) GetAccountBalances() []AccountBalance {
	if o == nil {
		return nil
	}
	return o.AccountBalances
}

func (o *CommerceCompanyInfo) GetAddresses() []CommerceAddress {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *CommerceCompanyInfo) GetBaseCurrency() *string {
	if o == nil {
		return nil
	}
	return o.BaseCurrency
}

func (o *CommerceCompanyInfo) GetCommercePlatformRef() *string {
	if o == nil {
		return nil
	}
	return o.CommercePlatformRef
}

func (o *CommerceCompanyInfo) GetCompanyLegalName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyLegalName
}

func (o *CommerceCompanyInfo) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *CommerceCompanyInfo) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *CommerceCompanyInfo) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *CommerceCompanyInfo) GetPhoneNumbers() []Items {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *CommerceCompanyInfo) GetRegistrationNumber() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationNumber
}

func (o *CommerceCompanyInfo) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *CommerceCompanyInfo) GetSourceUrls() map[string]string {
	if o == nil {
		return nil
	}
	return o.SourceUrls
}

func (o *CommerceCompanyInfo) GetWebLinks() []WebLinksitems {
	if o == nil {
		return nil
	}
	return o.WebLinks
}