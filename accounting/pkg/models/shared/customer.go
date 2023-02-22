package shared

import (
	"time"
)

type CustomerAddressesTypeEnum string

const (
	CustomerAddressesTypeEnumUnknown  CustomerAddressesTypeEnum = "Unknown"
	CustomerAddressesTypeEnumBilling  CustomerAddressesTypeEnum = "Billing"
	CustomerAddressesTypeEnumDelivery CustomerAddressesTypeEnum = "Delivery"
)

type CustomerAddresses struct {
	City       *string                   `json:"city,omitempty"`
	Country    *string                   `json:"country,omitempty"`
	Line1      *string                   `json:"line1,omitempty"`
	Line2      *string                   `json:"line2,omitempty"`
	PostalCode *string                   `json:"postalCode,omitempty"`
	Region     *string                   `json:"region,omitempty"`
	Type       CustomerAddressesTypeEnum `json:"type"`
}

type CustomerContacts struct {
	Address      *Addressesitems     `json:"address,omitempty"`
	Email        *string             `json:"email,omitempty"`
	ModifiedDate *time.Time          `json:"modifiedDate,omitempty"`
	Name         *string             `json:"name,omitempty"`
	Phone        []PhoneNumbersitems `json:"phone,omitempty"`
	Status       StatusEnum          `json:"status"`
}

type CustomerStatusEnum string

const (
	CustomerStatusEnumUnknown  CustomerStatusEnum = "Unknown"
	CustomerStatusEnumActive   CustomerStatusEnum = "Active"
	CustomerStatusEnumArchived CustomerStatusEnum = "Archived"
)

// Customer
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/docs/datamodel-accounting-invoices).
type Customer struct {
	Addresses          []CustomerAddresses `json:"addresses,omitempty"`
	ContactName        *string             `json:"contactName,omitempty"`
	Contacts           []CustomerContacts  `json:"contacts,omitempty"`
	CustomerName       *string             `json:"customerName,omitempty"`
	DefaultCurrency    *string             `json:"defaultCurrency,omitempty"`
	EmailAddress       *string             `json:"emailAddress,omitempty"`
	ID                 *string             `json:"id,omitempty"`
	Metadata           *Metadata           `json:"metadata,omitempty"`
	ModifiedDate       *time.Time          `json:"modifiedDate,omitempty"`
	Phone              *string             `json:"phone,omitempty"`
	RegistrationNumber *string             `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time          `json:"sourceModifiedDate,omitempty"`
	Status             CustomerStatusEnum  `json:"status"`
	SupplementalData   *SupplementalData   `json:"supplementalData,omitempty"`
	TaxNumber          *string             `json:"taxNumber,omitempty"`
}
