package operations

import (
	"net/http"
	"time"
)

type GetCustomersPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCustomersQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type GetCustomersRequest struct {
	PathParams  GetCustomersPathParams
	QueryParams GetCustomersQueryParams
}

type GetCustomersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetCustomersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetCustomersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetCustomersLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetCustomersLinksLinks struct {
	Current  GetCustomersLinksLinksCurrent   `json:"current"`
	Next     *GetCustomersLinksLinksNext     `json:"next,omitempty"`
	Previous *GetCustomersLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetCustomersLinksLinksSelf      `json:"self"`
}

type GetCustomersLinksSourceModifiedDateAddressesTypeEnum string

const (
	GetCustomersLinksSourceModifiedDateAddressesTypeEnumUnknown  GetCustomersLinksSourceModifiedDateAddressesTypeEnum = "Unknown"
	GetCustomersLinksSourceModifiedDateAddressesTypeEnumBilling  GetCustomersLinksSourceModifiedDateAddressesTypeEnum = "Billing"
	GetCustomersLinksSourceModifiedDateAddressesTypeEnumDelivery GetCustomersLinksSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type GetCustomersLinksSourceModifiedDateAddresses struct {
	City       *string                                              `json:"city,omitempty"`
	Country    *string                                              `json:"country,omitempty"`
	Line1      *string                                              `json:"line1,omitempty"`
	Line2      *string                                              `json:"line2,omitempty"`
	PostalCode *string                                              `json:"postalCode,omitempty"`
	Region     *string                                              `json:"region,omitempty"`
	Type       GetCustomersLinksSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type GetCustomersLinksSourceModifiedDateContactsAddressTypeEnum string

const (
	GetCustomersLinksSourceModifiedDateContactsAddressTypeEnumUnknown  GetCustomersLinksSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	GetCustomersLinksSourceModifiedDateContactsAddressTypeEnumBilling  GetCustomersLinksSourceModifiedDateContactsAddressTypeEnum = "Billing"
	GetCustomersLinksSourceModifiedDateContactsAddressTypeEnumDelivery GetCustomersLinksSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

// GetCustomersLinksSourceModifiedDateContactsAddress
// An object of Address information.
type GetCustomersLinksSourceModifiedDateContactsAddress struct {
	City       *string                                                    `json:"city,omitempty"`
	Country    *string                                                    `json:"country,omitempty"`
	Line1      *string                                                    `json:"line1,omitempty"`
	Line2      *string                                                    `json:"line2,omitempty"`
	PostalCode *string                                                    `json:"postalCode,omitempty"`
	Region     *string                                                    `json:"region,omitempty"`
	Type       GetCustomersLinksSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

type GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnum string

const (
	GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnumUnknown  GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnumPrimary  GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnumLandline GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnumMobile   GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnumFax      GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

type GetCustomersLinksSourceModifiedDateContactsPhone struct {
	Number *string                                                  `json:"number,omitempty"`
	Type   GetCustomersLinksSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

type GetCustomersLinksSourceModifiedDateContactsStatusEnum string

const (
	GetCustomersLinksSourceModifiedDateContactsStatusEnumUnknown  GetCustomersLinksSourceModifiedDateContactsStatusEnum = "Unknown"
	GetCustomersLinksSourceModifiedDateContactsStatusEnumActive   GetCustomersLinksSourceModifiedDateContactsStatusEnum = "Active"
	GetCustomersLinksSourceModifiedDateContactsStatusEnumArchived GetCustomersLinksSourceModifiedDateContactsStatusEnum = "Archived"
)

type GetCustomersLinksSourceModifiedDateContacts struct {
	Address      *GetCustomersLinksSourceModifiedDateContactsAddress   `json:"address,omitempty"`
	Email        *string                                               `json:"email,omitempty"`
	ModifiedDate *time.Time                                            `json:"modifiedDate,omitempty"`
	Name         *string                                               `json:"name,omitempty"`
	Phone        []GetCustomersLinksSourceModifiedDateContactsPhone    `json:"phone,omitempty"`
	Status       GetCustomersLinksSourceModifiedDateContactsStatusEnum `json:"status"`
}

type GetCustomersLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetCustomersLinksSourceModifiedDateStatusEnum string

const (
	GetCustomersLinksSourceModifiedDateStatusEnumUnknown  GetCustomersLinksSourceModifiedDateStatusEnum = "Unknown"
	GetCustomersLinksSourceModifiedDateStatusEnumActive   GetCustomersLinksSourceModifiedDateStatusEnum = "Active"
	GetCustomersLinksSourceModifiedDateStatusEnumArchived GetCustomersLinksSourceModifiedDateStatusEnum = "Archived"
)

// GetCustomersLinksSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetCustomersLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetCustomersLinksSourceModifiedDate
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type GetCustomersLinksSourceModifiedDate struct {
	Addresses          []GetCustomersLinksSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                              `json:"contactName,omitempty"`
	Contacts           []GetCustomersLinksSourceModifiedDateContacts        `json:"contacts,omitempty"`
	CustomerName       *string                                              `json:"customerName,omitempty"`
	DefaultCurrency    *string                                              `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                              `json:"emailAddress,omitempty"`
	ID                 *string                                              `json:"id,omitempty"`
	Metadata           *GetCustomersLinksSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                           `json:"modifiedDate,omitempty"`
	Phone              *string                                              `json:"phone,omitempty"`
	RegistrationNumber *string                                              `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                           `json:"sourceModifiedDate,omitempty"`
	Status             GetCustomersLinksSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *GetCustomersLinksSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TaxNumber          *string                                              `json:"taxNumber,omitempty"`
}

// GetCustomersLinks
// Codat's Paging Model
type GetCustomersLinks struct {
	Links        GetCustomersLinksLinks                `json:"_links"`
	PageNumber   int64                                 `json:"pageNumber"`
	PageSize     int64                                 `json:"pageSize"`
	Results      []GetCustomersLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                 `json:"totalResults"`
}

type GetCustomersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *GetCustomersLinks
}
