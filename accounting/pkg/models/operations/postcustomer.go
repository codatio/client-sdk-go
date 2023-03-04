package operations

import (
	"net/http"
	"time"
)

type PostCustomerPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostCustomerQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostCustomerSourceModifiedDateAddressesTypeEnum string

const (
	PostCustomerSourceModifiedDateAddressesTypeEnumUnknown  PostCustomerSourceModifiedDateAddressesTypeEnum = "Unknown"
	PostCustomerSourceModifiedDateAddressesTypeEnumBilling  PostCustomerSourceModifiedDateAddressesTypeEnum = "Billing"
	PostCustomerSourceModifiedDateAddressesTypeEnumDelivery PostCustomerSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type PostCustomerSourceModifiedDateAddresses struct {
	City       *string                                         `json:"city,omitempty"`
	Country    *string                                         `json:"country,omitempty"`
	Line1      *string                                         `json:"line1,omitempty"`
	Line2      *string                                         `json:"line2,omitempty"`
	PostalCode *string                                         `json:"postalCode,omitempty"`
	Region     *string                                         `json:"region,omitempty"`
	Type       PostCustomerSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type PostCustomerSourceModifiedDateContactsAddressTypeEnum string

const (
	PostCustomerSourceModifiedDateContactsAddressTypeEnumUnknown  PostCustomerSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	PostCustomerSourceModifiedDateContactsAddressTypeEnumBilling  PostCustomerSourceModifiedDateContactsAddressTypeEnum = "Billing"
	PostCustomerSourceModifiedDateContactsAddressTypeEnumDelivery PostCustomerSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

// PostCustomerSourceModifiedDateContactsAddress
// An object of Address information.
type PostCustomerSourceModifiedDateContactsAddress struct {
	City       *string                                               `json:"city,omitempty"`
	Country    *string                                               `json:"country,omitempty"`
	Line1      *string                                               `json:"line1,omitempty"`
	Line2      *string                                               `json:"line2,omitempty"`
	PostalCode *string                                               `json:"postalCode,omitempty"`
	Region     *string                                               `json:"region,omitempty"`
	Type       PostCustomerSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

type PostCustomerSourceModifiedDateContactsPhoneTypeEnum string

const (
	PostCustomerSourceModifiedDateContactsPhoneTypeEnumUnknown  PostCustomerSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	PostCustomerSourceModifiedDateContactsPhoneTypeEnumPrimary  PostCustomerSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	PostCustomerSourceModifiedDateContactsPhoneTypeEnumLandline PostCustomerSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	PostCustomerSourceModifiedDateContactsPhoneTypeEnumMobile   PostCustomerSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	PostCustomerSourceModifiedDateContactsPhoneTypeEnumFax      PostCustomerSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

type PostCustomerSourceModifiedDateContactsPhone struct {
	Number *string                                             `json:"number,omitempty"`
	Type   PostCustomerSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

type PostCustomerSourceModifiedDateContactsStatusEnum string

const (
	PostCustomerSourceModifiedDateContactsStatusEnumUnknown  PostCustomerSourceModifiedDateContactsStatusEnum = "Unknown"
	PostCustomerSourceModifiedDateContactsStatusEnumActive   PostCustomerSourceModifiedDateContactsStatusEnum = "Active"
	PostCustomerSourceModifiedDateContactsStatusEnumArchived PostCustomerSourceModifiedDateContactsStatusEnum = "Archived"
)

type PostCustomerSourceModifiedDateContacts struct {
	Address      *PostCustomerSourceModifiedDateContactsAddress   `json:"address,omitempty"`
	Email        *string                                          `json:"email,omitempty"`
	ModifiedDate *time.Time                                       `json:"modifiedDate,omitempty"`
	Name         *string                                          `json:"name,omitempty"`
	Phone        []PostCustomerSourceModifiedDateContactsPhone    `json:"phone,omitempty"`
	Status       PostCustomerSourceModifiedDateContactsStatusEnum `json:"status"`
}

type PostCustomerSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostCustomerSourceModifiedDateStatusEnum string

const (
	PostCustomerSourceModifiedDateStatusEnumUnknown  PostCustomerSourceModifiedDateStatusEnum = "Unknown"
	PostCustomerSourceModifiedDateStatusEnumActive   PostCustomerSourceModifiedDateStatusEnum = "Active"
	PostCustomerSourceModifiedDateStatusEnumArchived PostCustomerSourceModifiedDateStatusEnum = "Archived"
)

type PostCustomerSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostCustomerSourceModifiedDate
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type PostCustomerSourceModifiedDate struct {
	Addresses          []PostCustomerSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                         `json:"contactName,omitempty"`
	Contacts           []PostCustomerSourceModifiedDateContacts        `json:"contacts,omitempty"`
	CustomerName       *string                                         `json:"customerName,omitempty"`
	DefaultCurrency    *string                                         `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                         `json:"emailAddress,omitempty"`
	ID                 *string                                         `json:"id,omitempty"`
	Metadata           *PostCustomerSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                      `json:"modifiedDate,omitempty"`
	Phone              *string                                         `json:"phone,omitempty"`
	RegistrationNumber *string                                         `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                      `json:"sourceModifiedDate,omitempty"`
	Status             PostCustomerSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *PostCustomerSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TaxNumber          *string                                         `json:"taxNumber,omitempty"`
}

type PostCustomerRequest struct {
	PathParams  PostCustomerPathParams
	QueryParams PostCustomerQueryParams
	Request     *PostCustomerSourceModifiedDate `request:"mediaType=application/json"`
}

type PostCustomer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostCustomer200ApplicationJSONChangesTypeEnum string

const (
	PostCustomer200ApplicationJSONChangesTypeEnumUnknown            PostCustomer200ApplicationJSONChangesTypeEnum = "Unknown"
	PostCustomer200ApplicationJSONChangesTypeEnumCreated            PostCustomer200ApplicationJSONChangesTypeEnum = "Created"
	PostCustomer200ApplicationJSONChangesTypeEnumModified           PostCustomer200ApplicationJSONChangesTypeEnum = "Modified"
	PostCustomer200ApplicationJSONChangesTypeEnumDeleted            PostCustomer200ApplicationJSONChangesTypeEnum = "Deleted"
	PostCustomer200ApplicationJSONChangesTypeEnumAttachmentUploaded PostCustomer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostCustomer200ApplicationJSONChanges struct {
	AttachmentID *string                                                      `json:"attachmentId,omitempty"`
	RecordRef    *PostCustomer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostCustomer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum string

const (
	PostCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumUnknown  PostCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Unknown"
	PostCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumBilling  PostCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Billing"
	PostCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumDelivery PostCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type PostCustomer200ApplicationJSONSourceModifiedDateAddresses struct {
	City       *string                                                           `json:"city,omitempty"`
	Country    *string                                                           `json:"country,omitempty"`
	Line1      *string                                                           `json:"line1,omitempty"`
	Line2      *string                                                           `json:"line2,omitempty"`
	PostalCode *string                                                           `json:"postalCode,omitempty"`
	Region     *string                                                           `json:"region,omitempty"`
	Type       PostCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type PostCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum string

const (
	PostCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumUnknown  PostCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	PostCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumBilling  PostCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Billing"
	PostCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumDelivery PostCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

// PostCustomer200ApplicationJSONSourceModifiedDateContactsAddress
// An object of Address information.
type PostCustomer200ApplicationJSONSourceModifiedDateContactsAddress struct {
	City       *string                                                                 `json:"city,omitempty"`
	Country    *string                                                                 `json:"country,omitempty"`
	Line1      *string                                                                 `json:"line1,omitempty"`
	Line2      *string                                                                 `json:"line2,omitempty"`
	PostalCode *string                                                                 `json:"postalCode,omitempty"`
	Region     *string                                                                 `json:"region,omitempty"`
	Type       PostCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

type PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum string

const (
	PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumUnknown  PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumPrimary  PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumLandline PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumMobile   PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumFax      PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

type PostCustomer200ApplicationJSONSourceModifiedDateContactsPhone struct {
	Number *string                                                               `json:"number,omitempty"`
	Type   PostCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

type PostCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum string

const (
	PostCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumUnknown  PostCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Unknown"
	PostCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumActive   PostCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Active"
	PostCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumArchived PostCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Archived"
)

type PostCustomer200ApplicationJSONSourceModifiedDateContacts struct {
	Address      *PostCustomer200ApplicationJSONSourceModifiedDateContactsAddress   `json:"address,omitempty"`
	Email        *string                                                            `json:"email,omitempty"`
	ModifiedDate *time.Time                                                         `json:"modifiedDate,omitempty"`
	Name         *string                                                            `json:"name,omitempty"`
	Phone        []PostCustomer200ApplicationJSONSourceModifiedDateContactsPhone    `json:"phone,omitempty"`
	Status       PostCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum `json:"status"`
}

type PostCustomer200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostCustomer200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PostCustomer200ApplicationJSONSourceModifiedDateStatusEnumUnknown  PostCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PostCustomer200ApplicationJSONSourceModifiedDateStatusEnumActive   PostCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	PostCustomer200ApplicationJSONSourceModifiedDateStatusEnumArchived PostCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

type PostCustomer200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostCustomer200ApplicationJSONSourceModifiedDate
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type PostCustomer200ApplicationJSONSourceModifiedDate struct {
	Addresses          []PostCustomer200ApplicationJSONSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                                           `json:"contactName,omitempty"`
	Contacts           []PostCustomer200ApplicationJSONSourceModifiedDateContacts        `json:"contacts,omitempty"`
	CustomerName       *string                                                           `json:"customerName,omitempty"`
	DefaultCurrency    *string                                                           `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                                           `json:"emailAddress,omitempty"`
	ID                 *string                                                           `json:"id,omitempty"`
	Metadata           *PostCustomer200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                        `json:"modifiedDate,omitempty"`
	Phone              *string                                                           `json:"phone,omitempty"`
	RegistrationNumber *string                                                           `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                                        `json:"sourceModifiedDate,omitempty"`
	Status             PostCustomer200ApplicationJSONSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *PostCustomer200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TaxNumber          *string                                                           `json:"taxNumber,omitempty"`
}

type PostCustomer200ApplicationJSONStatusEnum string

const (
	PostCustomer200ApplicationJSONStatusEnumPending  PostCustomer200ApplicationJSONStatusEnum = "Pending"
	PostCustomer200ApplicationJSONStatusEnumFailed   PostCustomer200ApplicationJSONStatusEnum = "Failed"
	PostCustomer200ApplicationJSONStatusEnumSuccess  PostCustomer200ApplicationJSONStatusEnum = "Success"
	PostCustomer200ApplicationJSONStatusEnumTimedOut PostCustomer200ApplicationJSONStatusEnum = "TimedOut"
)

type PostCustomer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostCustomer200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostCustomer200ApplicationJSONValidation struct {
	Errors   []PostCustomer200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostCustomer200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostCustomer200ApplicationJSON struct {
	Changes           []PostCustomer200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                            `json:"companyId"`
	CompletedOnUtc    *time.Time                                        `json:"completedOnUtc,omitempty"`
	Data              *PostCustomer200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                            `json:"dataConnectionKey"`
	DataType          *string                                           `json:"dataType,omitempty"`
	ErrorMessage      *string                                           `json:"errorMessage,omitempty"`
	PushOperationKey  string                                            `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                         `json:"requestedOnUtc"`
	Status            PostCustomer200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                               `json:"statusCode"`
	TimeoutInMinutes  *int                                              `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                              `json:"timeoutInSeconds,omitempty"`
	Validation        *PostCustomer200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostCustomerResponse struct {
	ContentType                          string
	StatusCode                           int
	RawResponse                          *http.Response
	PostCustomer200ApplicationJSONObject *PostCustomer200ApplicationJSON
}
