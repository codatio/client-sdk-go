package operations

import (
	"net/http"
	"time"
)

type GetCustomerRequest struct {
	CompanyID  string `pathParam:"style=simple,explode=false,name=companyId"`
	CustomerID string `pathParam:"style=simple,explode=false,name=customerId"`
}

type GetCustomerSourceModifiedDateAddressesTypeEnum string

const (
	GetCustomerSourceModifiedDateAddressesTypeEnumUnknown  GetCustomerSourceModifiedDateAddressesTypeEnum = "Unknown"
	GetCustomerSourceModifiedDateAddressesTypeEnumBilling  GetCustomerSourceModifiedDateAddressesTypeEnum = "Billing"
	GetCustomerSourceModifiedDateAddressesTypeEnumDelivery GetCustomerSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type GetCustomerSourceModifiedDateAddresses struct {
	City       *string                                        `json:"city,omitempty"`
	Country    *string                                        `json:"country,omitempty"`
	Line1      *string                                        `json:"line1,omitempty"`
	Line2      *string                                        `json:"line2,omitempty"`
	PostalCode *string                                        `json:"postalCode,omitempty"`
	Region     *string                                        `json:"region,omitempty"`
	Type       GetCustomerSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type GetCustomerSourceModifiedDateContactsAddressTypeEnum string

const (
	GetCustomerSourceModifiedDateContactsAddressTypeEnumUnknown  GetCustomerSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	GetCustomerSourceModifiedDateContactsAddressTypeEnumBilling  GetCustomerSourceModifiedDateContactsAddressTypeEnum = "Billing"
	GetCustomerSourceModifiedDateContactsAddressTypeEnumDelivery GetCustomerSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

// GetCustomerSourceModifiedDateContactsAddress
// An object of Address information.
type GetCustomerSourceModifiedDateContactsAddress struct {
	City       *string                                              `json:"city,omitempty"`
	Country    *string                                              `json:"country,omitempty"`
	Line1      *string                                              `json:"line1,omitempty"`
	Line2      *string                                              `json:"line2,omitempty"`
	PostalCode *string                                              `json:"postalCode,omitempty"`
	Region     *string                                              `json:"region,omitempty"`
	Type       GetCustomerSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

type GetCustomerSourceModifiedDateContactsPhoneTypeEnum string

const (
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumUnknown  GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumPrimary  GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumLandline GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumMobile   GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	GetCustomerSourceModifiedDateContactsPhoneTypeEnumFax      GetCustomerSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

type GetCustomerSourceModifiedDateContactsPhone struct {
	Number *string                                            `json:"number,omitempty"`
	Type   GetCustomerSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

type GetCustomerSourceModifiedDateContactsStatusEnum string

const (
	GetCustomerSourceModifiedDateContactsStatusEnumUnknown  GetCustomerSourceModifiedDateContactsStatusEnum = "Unknown"
	GetCustomerSourceModifiedDateContactsStatusEnumActive   GetCustomerSourceModifiedDateContactsStatusEnum = "Active"
	GetCustomerSourceModifiedDateContactsStatusEnumArchived GetCustomerSourceModifiedDateContactsStatusEnum = "Archived"
)

type GetCustomerSourceModifiedDateContacts struct {
	Address      *GetCustomerSourceModifiedDateContactsAddress   `json:"address,omitempty"`
	Email        *string                                         `json:"email,omitempty"`
	ModifiedDate *time.Time                                      `json:"modifiedDate,omitempty"`
	Name         *string                                         `json:"name,omitempty"`
	Phone        []GetCustomerSourceModifiedDateContactsPhone    `json:"phone,omitempty"`
	Status       GetCustomerSourceModifiedDateContactsStatusEnum `json:"status"`
}

type GetCustomerSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetCustomerSourceModifiedDateStatusEnum string

const (
	GetCustomerSourceModifiedDateStatusEnumUnknown  GetCustomerSourceModifiedDateStatusEnum = "Unknown"
	GetCustomerSourceModifiedDateStatusEnumActive   GetCustomerSourceModifiedDateStatusEnum = "Active"
	GetCustomerSourceModifiedDateStatusEnumArchived GetCustomerSourceModifiedDateStatusEnum = "Archived"
)

// GetCustomerSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type GetCustomerSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetCustomerSourceModifiedDate
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type GetCustomerSourceModifiedDate struct {
	Addresses          []GetCustomerSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                        `json:"contactName,omitempty"`
	Contacts           []GetCustomerSourceModifiedDateContacts        `json:"contacts,omitempty"`
	CustomerName       *string                                        `json:"customerName,omitempty"`
	DefaultCurrency    *string                                        `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                        `json:"emailAddress,omitempty"`
	ID                 *string                                        `json:"id,omitempty"`
	Metadata           *GetCustomerSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                     `json:"modifiedDate,omitempty"`
	Phone              *string                                        `json:"phone,omitempty"`
	RegistrationNumber *string                                        `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                     `json:"sourceModifiedDate,omitempty"`
	Status             GetCustomerSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *GetCustomerSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TaxNumber          *string                                        `json:"taxNumber,omitempty"`
}

type GetCustomerResponse struct {
	ContentType        string
	SourceModifiedDate *GetCustomerSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
