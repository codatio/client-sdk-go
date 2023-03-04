package operations

import (
	"net/http"
	"time"
)

type UpdateCustomerPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string `pathParam:"style=simple,explode=false,name=customerId"`
}

type UpdateCustomerQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type UpdateCustomerSourceModifiedDateAddressesTypeEnum string

const (
	UpdateCustomerSourceModifiedDateAddressesTypeEnumUnknown  UpdateCustomerSourceModifiedDateAddressesTypeEnum = "Unknown"
	UpdateCustomerSourceModifiedDateAddressesTypeEnumBilling  UpdateCustomerSourceModifiedDateAddressesTypeEnum = "Billing"
	UpdateCustomerSourceModifiedDateAddressesTypeEnumDelivery UpdateCustomerSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type UpdateCustomerSourceModifiedDateAddresses struct {
	City       *string                                           `json:"city,omitempty"`
	Country    *string                                           `json:"country,omitempty"`
	Line1      *string                                           `json:"line1,omitempty"`
	Line2      *string                                           `json:"line2,omitempty"`
	PostalCode *string                                           `json:"postalCode,omitempty"`
	Region     *string                                           `json:"region,omitempty"`
	Type       UpdateCustomerSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type UpdateCustomerSourceModifiedDateContactsAddressTypeEnum string

const (
	UpdateCustomerSourceModifiedDateContactsAddressTypeEnumUnknown  UpdateCustomerSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	UpdateCustomerSourceModifiedDateContactsAddressTypeEnumBilling  UpdateCustomerSourceModifiedDateContactsAddressTypeEnum = "Billing"
	UpdateCustomerSourceModifiedDateContactsAddressTypeEnumDelivery UpdateCustomerSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

// UpdateCustomerSourceModifiedDateContactsAddress
// An object of Address information.
type UpdateCustomerSourceModifiedDateContactsAddress struct {
	City       *string                                                 `json:"city,omitempty"`
	Country    *string                                                 `json:"country,omitempty"`
	Line1      *string                                                 `json:"line1,omitempty"`
	Line2      *string                                                 `json:"line2,omitempty"`
	PostalCode *string                                                 `json:"postalCode,omitempty"`
	Region     *string                                                 `json:"region,omitempty"`
	Type       UpdateCustomerSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

type UpdateCustomerSourceModifiedDateContactsPhoneTypeEnum string

const (
	UpdateCustomerSourceModifiedDateContactsPhoneTypeEnumUnknown  UpdateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	UpdateCustomerSourceModifiedDateContactsPhoneTypeEnumPrimary  UpdateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	UpdateCustomerSourceModifiedDateContactsPhoneTypeEnumLandline UpdateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	UpdateCustomerSourceModifiedDateContactsPhoneTypeEnumMobile   UpdateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	UpdateCustomerSourceModifiedDateContactsPhoneTypeEnumFax      UpdateCustomerSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

type UpdateCustomerSourceModifiedDateContactsPhone struct {
	Number *string                                               `json:"number,omitempty"`
	Type   UpdateCustomerSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

type UpdateCustomerSourceModifiedDateContactsStatusEnum string

const (
	UpdateCustomerSourceModifiedDateContactsStatusEnumUnknown  UpdateCustomerSourceModifiedDateContactsStatusEnum = "Unknown"
	UpdateCustomerSourceModifiedDateContactsStatusEnumActive   UpdateCustomerSourceModifiedDateContactsStatusEnum = "Active"
	UpdateCustomerSourceModifiedDateContactsStatusEnumArchived UpdateCustomerSourceModifiedDateContactsStatusEnum = "Archived"
)

type UpdateCustomerSourceModifiedDateContacts struct {
	Address      *UpdateCustomerSourceModifiedDateContactsAddress   `json:"address,omitempty"`
	Email        *string                                            `json:"email,omitempty"`
	ModifiedDate *time.Time                                         `json:"modifiedDate,omitempty"`
	Name         *string                                            `json:"name,omitempty"`
	Phone        []UpdateCustomerSourceModifiedDateContactsPhone    `json:"phone,omitempty"`
	Status       UpdateCustomerSourceModifiedDateContactsStatusEnum `json:"status"`
}

type UpdateCustomerSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdateCustomerSourceModifiedDateStatusEnum string

const (
	UpdateCustomerSourceModifiedDateStatusEnumUnknown  UpdateCustomerSourceModifiedDateStatusEnum = "Unknown"
	UpdateCustomerSourceModifiedDateStatusEnumActive   UpdateCustomerSourceModifiedDateStatusEnum = "Active"
	UpdateCustomerSourceModifiedDateStatusEnumArchived UpdateCustomerSourceModifiedDateStatusEnum = "Archived"
)

type UpdateCustomerSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// UpdateCustomerSourceModifiedDate
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type UpdateCustomerSourceModifiedDate struct {
	Addresses          []UpdateCustomerSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                           `json:"contactName,omitempty"`
	Contacts           []UpdateCustomerSourceModifiedDateContacts        `json:"contacts,omitempty"`
	CustomerName       *string                                           `json:"customerName,omitempty"`
	DefaultCurrency    *string                                           `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                           `json:"emailAddress,omitempty"`
	ID                 *string                                           `json:"id,omitempty"`
	Metadata           *UpdateCustomerSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                        `json:"modifiedDate,omitempty"`
	Phone              *string                                           `json:"phone,omitempty"`
	RegistrationNumber *string                                           `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                        `json:"sourceModifiedDate,omitempty"`
	Status             UpdateCustomerSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *UpdateCustomerSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TaxNumber          *string                                           `json:"taxNumber,omitempty"`
}

type UpdateCustomerRequest struct {
	PathParams  UpdateCustomerPathParams
	QueryParams UpdateCustomerQueryParams
	Request     *UpdateCustomerSourceModifiedDate `request:"mediaType=application/json"`
}

type UpdateCustomer200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdateCustomer200ApplicationJSONChangesTypeEnum string

const (
	UpdateCustomer200ApplicationJSONChangesTypeEnumUnknown            UpdateCustomer200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdateCustomer200ApplicationJSONChangesTypeEnumCreated            UpdateCustomer200ApplicationJSONChangesTypeEnum = "Created"
	UpdateCustomer200ApplicationJSONChangesTypeEnumModified           UpdateCustomer200ApplicationJSONChangesTypeEnum = "Modified"
	UpdateCustomer200ApplicationJSONChangesTypeEnumDeleted            UpdateCustomer200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdateCustomer200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdateCustomer200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdateCustomer200ApplicationJSONChanges struct {
	AttachmentID *string                                                        `json:"attachmentId,omitempty"`
	RecordRef    *UpdateCustomer200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdateCustomer200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type UpdateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum string

const (
	UpdateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumUnknown  UpdateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Unknown"
	UpdateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumBilling  UpdateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Billing"
	UpdateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnumDelivery UpdateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type UpdateCustomer200ApplicationJSONSourceModifiedDateAddresses struct {
	City       *string                                                             `json:"city,omitempty"`
	Country    *string                                                             `json:"country,omitempty"`
	Line1      *string                                                             `json:"line1,omitempty"`
	Line2      *string                                                             `json:"line2,omitempty"`
	PostalCode *string                                                             `json:"postalCode,omitempty"`
	Region     *string                                                             `json:"region,omitempty"`
	Type       UpdateCustomer200ApplicationJSONSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum string

const (
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumUnknown  UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Unknown"
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumBilling  UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Billing"
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnumDelivery UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum = "Delivery"
)

// UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddress
// An object of Address information.
type UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddress struct {
	City       *string                                                                   `json:"city,omitempty"`
	Country    *string                                                                   `json:"country,omitempty"`
	Line1      *string                                                                   `json:"line1,omitempty"`
	Line2      *string                                                                   `json:"line2,omitempty"`
	PostalCode *string                                                                   `json:"postalCode,omitempty"`
	Region     *string                                                                   `json:"region,omitempty"`
	Type       UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddressTypeEnum `json:"type"`
}

type UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum string

const (
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumUnknown  UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Unknown"
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumPrimary  UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Primary"
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumLandline UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Landline"
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumMobile   UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Mobile"
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnumFax      UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum = "Fax"
)

type UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhone struct {
	Number *string                                                                 `json:"number,omitempty"`
	Type   UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhoneTypeEnum `json:"type"`
}

type UpdateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum string

const (
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumUnknown  UpdateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Unknown"
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumActive   UpdateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Active"
	UpdateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnumArchived UpdateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum = "Archived"
)

type UpdateCustomer200ApplicationJSONSourceModifiedDateContacts struct {
	Address      *UpdateCustomer200ApplicationJSONSourceModifiedDateContactsAddress   `json:"address,omitempty"`
	Email        *string                                                              `json:"email,omitempty"`
	ModifiedDate *time.Time                                                           `json:"modifiedDate,omitempty"`
	Name         *string                                                              `json:"name,omitempty"`
	Phone        []UpdateCustomer200ApplicationJSONSourceModifiedDateContactsPhone    `json:"phone,omitempty"`
	Status       UpdateCustomer200ApplicationJSONSourceModifiedDateContactsStatusEnum `json:"status"`
}

type UpdateCustomer200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdateCustomer200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	UpdateCustomer200ApplicationJSONSourceModifiedDateStatusEnumUnknown  UpdateCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	UpdateCustomer200ApplicationJSONSourceModifiedDateStatusEnumActive   UpdateCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	UpdateCustomer200ApplicationJSONSourceModifiedDateStatusEnumArchived UpdateCustomer200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

type UpdateCustomer200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// UpdateCustomer200ApplicationJSONSourceModifiedDate
// > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
type UpdateCustomer200ApplicationJSONSourceModifiedDate struct {
	Addresses          []UpdateCustomer200ApplicationJSONSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                                             `json:"contactName,omitempty"`
	Contacts           []UpdateCustomer200ApplicationJSONSourceModifiedDateContacts        `json:"contacts,omitempty"`
	CustomerName       *string                                                             `json:"customerName,omitempty"`
	DefaultCurrency    *string                                                             `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                                             `json:"emailAddress,omitempty"`
	ID                 *string                                                             `json:"id,omitempty"`
	Metadata           *UpdateCustomer200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                          `json:"modifiedDate,omitempty"`
	Phone              *string                                                             `json:"phone,omitempty"`
	RegistrationNumber *string                                                             `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                                          `json:"sourceModifiedDate,omitempty"`
	Status             UpdateCustomer200ApplicationJSONSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *UpdateCustomer200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	TaxNumber          *string                                                             `json:"taxNumber,omitempty"`
}

type UpdateCustomer200ApplicationJSONStatusEnum string

const (
	UpdateCustomer200ApplicationJSONStatusEnumPending  UpdateCustomer200ApplicationJSONStatusEnum = "Pending"
	UpdateCustomer200ApplicationJSONStatusEnumFailed   UpdateCustomer200ApplicationJSONStatusEnum = "Failed"
	UpdateCustomer200ApplicationJSONStatusEnumSuccess  UpdateCustomer200ApplicationJSONStatusEnum = "Success"
	UpdateCustomer200ApplicationJSONStatusEnumTimedOut UpdateCustomer200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdateCustomer200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdateCustomer200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdateCustomer200ApplicationJSONValidation struct {
	Errors   []UpdateCustomer200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []UpdateCustomer200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type UpdateCustomer200ApplicationJSON struct {
	Changes           []UpdateCustomer200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                              `json:"companyId"`
	CompletedOnUtc    *time.Time                                          `json:"completedOnUtc,omitempty"`
	Data              *UpdateCustomer200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                              `json:"dataConnectionKey"`
	DataType          *string                                             `json:"dataType,omitempty"`
	ErrorMessage      *string                                             `json:"errorMessage,omitempty"`
	PushOperationKey  string                                              `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                           `json:"requestedOnUtc"`
	Status            UpdateCustomer200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                 `json:"statusCode"`
	TimeoutInMinutes  *int                                                `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdateCustomer200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type UpdateCustomerResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	UpdateCustomer200ApplicationJSONObject *UpdateCustomer200ApplicationJSON
}
