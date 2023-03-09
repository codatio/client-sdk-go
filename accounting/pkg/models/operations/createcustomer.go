package operations

import (
	"net/http"
	"time"
)

type CreateCustomerPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateCustomerQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateCustomerSourceModifiedDateAddressesTypeEnum string

const (
	CreateCustomerSourceModifiedDateAddressesTypeEnumUnknown  CreateCustomerSourceModifiedDateAddressesTypeEnum = "Unknown"
	CreateCustomerSourceModifiedDateAddressesTypeEnumBilling  CreateCustomerSourceModifiedDateAddressesTypeEnum = "Billing"
	CreateCustomerSourceModifiedDateAddressesTypeEnumDelivery CreateCustomerSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type CreateCustomerSourceModifiedDateAddresses struct {
	City       *string                                           `json:"city,omitempty"`
	Country    *string                                           `json:"country,omitempty"`
	Line1      *string                                           `json:"line1,omitempty"`
	Line2      *string                                           `json:"line2,omitempty"`
	PostalCode *string                                           `json:"postalCode,omitempty"`
	Region     *string                                           `json:"region,omitempty"`
	Type       CreateCustomerSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type CreateCustomerSourceModifiedDateContactsAddressTypeEnum string

const (
	CreateCustomerSourceModifiedDateContactsAddressTypeEnumUnknown  CreateCustomerSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	CreateCustomerSourceModifiedDateContactsAddressTypeEnumBilling  CreateCustomerSourceModifiedDateContactsAddressTypeEnum = "Billing"
	CreateCustomerSourceModifiedDateContactsAddressTypeEnumDelivery CreateCustomerSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

// CreateCustomerSourceModifiedDateContactsAddress
// An object of Address information.
type CreateCustomerSourceModifiedDateContactsAddress struct {
	City       *string                                                 `json:"city,omitempty"`
	Country    *string                                                 `json:"country,omitempty"`
	Line1      *string                                                 `json:"line1,omitempty"`
	Line2      *string                                                 `json:"line2,omitempty"`
	PostalCode *string                                                 `json:"postalCode,omitempty"`
	Region     *string                                                 `json:"region,omitempty"`
	Type       CreateCustomerSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

type CreateCustomerSourceModifiedDateContactsPhoneTypeEnum string

const (
	CreateCustomerSourceModifiedDateContactsPhoneTypeEnumUnknown  CreateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	CreateCustomerSourceModifiedDateContactsPhoneTypeEnumPrimary  CreateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	CreateCustomerSourceModifiedDateContactsPhoneTypeEnumLandline CreateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	CreateCustomerSourceModifiedDateContactsPhoneTypeEnumMobile   CreateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	CreateCustomerSourceModifiedDateContactsPhoneTypeEnumFax      CreateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

type CreateCustomerSourceModifiedDateContactsPhone struct {
	Number *string                                               `json:"number,omitempty"`
	Type   CreateCustomerSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

type CreateCustomerSourceModifiedDateContactsStatusEnum string

const (
	CreateCustomerSourceModifiedDateContactsStatusEnumUnknown  CreateCustomerSourceModifiedDateContactsStatusEnum = "Unknown"
	CreateCustomerSourceModifiedDateContactsStatusEnumActive   CreateCustomerSourceModifiedDateContactsStatusEnum = "Active"
	CreateCustomerSourceModifiedDateContactsStatusEnumArchived CreateCustomerSourceModifiedDateContactsStatusEnum = "Archived"
)

type CreateCustomerSourceModifiedDateContacts struct {
	Address      *CreateCustomerSourceModifiedDateContactsAddress   `json:"address,omitempty"`
	Email        *string                                            `json:"email,omitempty"`
	ModifiedDate *time.Time                                         `json:"modifiedDate,omitempty"`
	Name         *string                                            `json:"name,omitempty"`
	Phone        []CreateCustomerSourceModifiedDateContactsPhone    `json:"phone,omitempty"`
	Status       CreateCustomerSourceModifiedDateContactsStatusEnum `json:"status"`
}

type CreateCustomerSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateCustomerSourceModifiedDateStatusEnum string

const (
	CreateCustomerSourceModifiedDateStatusEnumUnknown  CreateCustomerSourceModifiedDateStatusEnum = "Unknown"
	CreateCustomerSourceModifiedDateStatusEnumActive   CreateCustomerSourceModifiedDateStatusEnum = "Active"
	CreateCustomerSourceModifiedDateStatusEnumArchived CreateCustomerSourceModifiedDateStatusEnum = "Archived"
)

// CreateCustomerSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateCustomerSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateCustomerSourceModifiedDate
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type CreateCustomerSourceModifiedDate struct {
	Addresses          []CreateCustomerSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                           `json:"contactName,omitempty"`
	Contacts           []CreateCustomerSourceModifiedDateContacts        `json:"contacts,omitempty"`
	CustomerName       *string                                           `json:"customerName,omitempty"`
	DefaultCurrency    *string                                           `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                           `json:"emailAddress,omitempty"`
	ID                 *string                                           `json:"id,omitempty"`
	Metadata           *CreateCustomerSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                        `json:"modifiedDate,omitempty"`
	Phone              *string                                           `json:"phone,omitempty"`
	RegistrationNumber *string                                           `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                        `json:"sourceModifiedDate,omitempty"`
	Status             CreateCustomerSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *CreateCustomerSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TaxNumber          *string                                           `json:"taxNumber,omitempty"`
}

type CreateCustomerRequest struct {
	PathParams  CreateCustomerPathParams
	QueryParams CreateCustomerQueryParams
	Request     *CreateCustomerSourceModifiedDate `request:"mediaType=application/json"`
}

type CreateCustomer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateCustomer200ApplicationJSONChangesTypeEnum string

const (
	CreateCustomer200ApplicationJSONChangesTypeEnumUnknown            CreateCustomer200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateCustomer200ApplicationJSONChangesTypeEnumCreated            CreateCustomer200ApplicationJSONChangesTypeEnum = "Created"
	CreateCustomer200ApplicationJSONChangesTypeEnumModified           CreateCustomer200ApplicationJSONChangesTypeEnum = "Modified"
	CreateCustomer200ApplicationJSONChangesTypeEnumDeleted            CreateCustomer200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateCustomer200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateCustomer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateCustomer200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *CreateCustomer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateCustomer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type CreateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum string

const (
	CreateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumUnknown  CreateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Unknown"
	CreateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumBilling  CreateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Billing"
	CreateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumDelivery CreateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type CreateCustomer200ApplicationJSONSourceModifiedDateAddresses struct {
	City       *string                                                             `json:"city,omitempty"`
	Country    *string                                                             `json:"country,omitempty"`
	Line1      *string                                                             `json:"line1,omitempty"`
	Line2      *string                                                             `json:"line2,omitempty"`
	PostalCode *string                                                             `json:"postalCode,omitempty"`
	Region     *string                                                             `json:"region,omitempty"`
	Type       CreateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum string

const (
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumUnknown  CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumBilling  CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Billing"
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumDelivery CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

// CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddress
// An object of Address information.
type CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddress struct {
	City       *string                                                                   `json:"city,omitempty"`
	Country    *string                                                                   `json:"country,omitempty"`
	Line1      *string                                                                   `json:"line1,omitempty"`
	Line2      *string                                                                   `json:"line2,omitempty"`
	PostalCode *string                                                                   `json:"postalCode,omitempty"`
	Region     *string                                                                   `json:"region,omitempty"`
	Type       CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

type CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum string

const (
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumUnknown  CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumPrimary  CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumLandline CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumMobile   CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumFax      CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

type CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhone struct {
	Number *string                                                                 `json:"number,omitempty"`
	Type   CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

type CreateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum string

const (
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumUnknown  CreateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Unknown"
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumActive   CreateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Active"
	CreateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumArchived CreateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Archived"
)

type CreateCustomer200ApplicationJSONSourceModifiedDateContacts struct {
	Address      *CreateCustomer200ApplicationJSONSourceModifiedDateContactsAddress   `json:"address,omitempty"`
	Email        *string                                                              `json:"email,omitempty"`
	ModifiedDate *time.Time                                                           `json:"modifiedDate,omitempty"`
	Name         *string                                                              `json:"name,omitempty"`
	Phone        []CreateCustomer200ApplicationJSONSourceModifiedDateContactsPhone    `json:"phone,omitempty"`
	Status       CreateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum `json:"status"`
}

type CreateCustomer200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateCustomer200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreateCustomer200ApplicationJSONSourceModifiedDateStatusEnumUnknown  CreateCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreateCustomer200ApplicationJSONSourceModifiedDateStatusEnumActive   CreateCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	CreateCustomer200ApplicationJSONSourceModifiedDateStatusEnumArchived CreateCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

// CreateCustomer200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateCustomer200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateCustomer200ApplicationJSONSourceModifiedDate
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type CreateCustomer200ApplicationJSONSourceModifiedDate struct {
	Addresses          []CreateCustomer200ApplicationJSONSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                                             `json:"contactName,omitempty"`
	Contacts           []CreateCustomer200ApplicationJSONSourceModifiedDateContacts        `json:"contacts,omitempty"`
	CustomerName       *string                                                             `json:"customerName,omitempty"`
	DefaultCurrency    *string                                                             `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                                             `json:"emailAddress,omitempty"`
	ID                 *string                                                             `json:"id,omitempty"`
	Metadata           *CreateCustomer200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                          `json:"modifiedDate,omitempty"`
	Phone              *string                                                             `json:"phone,omitempty"`
	RegistrationNumber *string                                                             `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                                          `json:"sourceModifiedDate,omitempty"`
	Status             CreateCustomer200ApplicationJSONSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *CreateCustomer200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TaxNumber          *string                                                             `json:"taxNumber,omitempty"`
}

type CreateCustomer200ApplicationJSONStatusEnum string

const (
	CreateCustomer200ApplicationJSONStatusEnumPending  CreateCustomer200ApplicationJSONStatusEnum = "Pending"
	CreateCustomer200ApplicationJSONStatusEnumFailed   CreateCustomer200ApplicationJSONStatusEnum = "Failed"
	CreateCustomer200ApplicationJSONStatusEnumSuccess  CreateCustomer200ApplicationJSONStatusEnum = "Success"
	CreateCustomer200ApplicationJSONStatusEnumTimedOut CreateCustomer200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateCustomer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateCustomer200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateCustomer200ApplicationJSONValidation struct {
	Errors   []CreateCustomer200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateCustomer200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateCustomer200ApplicationJSON struct {
	Changes           []CreateCustomer200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                              `json:"companyId"`
	CompletedOnUtc    *time.Time                                          `json:"completedOnUtc,omitempty"`
	Data              *CreateCustomer200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                              `json:"dataConnectionKey"`
	DataType          *string                                             `json:"dataType,omitempty"`
	ErrorMessage      *string                                             `json:"errorMessage,omitempty"`
	PushOperationKey  string                                              `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                           `json:"requestedOnUtc"`
	Status            CreateCustomer200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                 `json:"statusCode"`
	TimeoutInMinutes  *int                                                `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateCustomer200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateCustomerResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	CreateCustomer200ApplicationJSONObject *CreateCustomer200ApplicationJSON
}
