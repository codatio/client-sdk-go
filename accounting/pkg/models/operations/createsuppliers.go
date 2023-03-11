package operations

import (
	"net/http"
	"time"
)

type CreateSuppliersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreateSuppliersQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateSuppliersSourceModifiedDateAddressesTypeEnum string

const (
	CreateSuppliersSourceModifiedDateAddressesTypeEnumUnknown  CreateSuppliersSourceModifiedDateAddressesTypeEnum = "Unknown"
	CreateSuppliersSourceModifiedDateAddressesTypeEnumBilling  CreateSuppliersSourceModifiedDateAddressesTypeEnum = "Billing"
	CreateSuppliersSourceModifiedDateAddressesTypeEnumDelivery CreateSuppliersSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type CreateSuppliersSourceModifiedDateAddresses struct {
	City       *string                                            `json:"city,omitempty"`
	Country    *string                                            `json:"country,omitempty"`
	Line1      *string                                            `json:"line1,omitempty"`
	Line2      *string                                            `json:"line2,omitempty"`
	PostalCode *string                                            `json:"postalCode,omitempty"`
	Region     *string                                            `json:"region,omitempty"`
	Type       CreateSuppliersSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type CreateSuppliersSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateSuppliersSourceModifiedDateStatusEnum string

const (
	CreateSuppliersSourceModifiedDateStatusEnumUnknown  CreateSuppliersSourceModifiedDateStatusEnum = "Unknown"
	CreateSuppliersSourceModifiedDateStatusEnumActive   CreateSuppliersSourceModifiedDateStatusEnum = "Active"
	CreateSuppliersSourceModifiedDateStatusEnumArchived CreateSuppliersSourceModifiedDateStatusEnum = "Archived"
)

// CreateSuppliersSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateSuppliersSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateSuppliersSourceModifiedDate
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type CreateSuppliersSourceModifiedDate struct {
	Addresses          []CreateSuppliersSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                            `json:"contactName,omitempty"`
	DefaultCurrency    *string                                            `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                            `json:"emailAddress,omitempty"`
	ID                 *string                                            `json:"id,omitempty"`
	Metadata           *CreateSuppliersSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                         `json:"modifiedDate,omitempty"`
	Phone              *string                                            `json:"phone,omitempty"`
	RegistrationNumber *string                                            `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                         `json:"sourceModifiedDate,omitempty"`
	Status             CreateSuppliersSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *CreateSuppliersSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	SupplierName       *string                                            `json:"supplierName,omitempty"`
	TaxNumber          *string                                            `json:"taxNumber,omitempty"`
}

type CreateSuppliersRequest struct {
	PathParams  CreateSuppliersPathParams
	QueryParams CreateSuppliersQueryParams
	Request     *CreateSuppliersSourceModifiedDate `request:"mediaType=application/json"`
}

type CreateSuppliers200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateSuppliers200ApplicationJSONChangesTypeEnum string

const (
	CreateSuppliers200ApplicationJSONChangesTypeEnumUnknown            CreateSuppliers200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateSuppliers200ApplicationJSONChangesTypeEnumCreated            CreateSuppliers200ApplicationJSONChangesTypeEnum = "Created"
	CreateSuppliers200ApplicationJSONChangesTypeEnumModified           CreateSuppliers200ApplicationJSONChangesTypeEnum = "Modified"
	CreateSuppliers200ApplicationJSONChangesTypeEnumDeleted            CreateSuppliers200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateSuppliers200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateSuppliers200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateSuppliers200ApplicationJSONChanges struct {
	AttachmentID *string                                                         `json:"attachmentId,omitempty"`
	RecordRef    *CreateSuppliers200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateSuppliers200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum string

const (
	CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumUnknown  CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Unknown"
	CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumBilling  CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Billing"
	CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumDelivery CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type CreateSuppliers200ApplicationJSONSourceModifiedDateAddresses struct {
	City       *string                                                              `json:"city,omitempty"`
	Country    *string                                                              `json:"country,omitempty"`
	Line1      *string                                                              `json:"line1,omitempty"`
	Line2      *string                                                              `json:"line2,omitempty"`
	PostalCode *string                                                              `json:"postalCode,omitempty"`
	Region     *string                                                              `json:"region,omitempty"`
	Type       CreateSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type CreateSuppliers200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnumUnknown  CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnumActive   CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnumArchived CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

// CreateSuppliers200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type CreateSuppliers200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// CreateSuppliers200ApplicationJSONSourceModifiedDate
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type CreateSuppliers200ApplicationJSONSourceModifiedDate struct {
	Addresses          []CreateSuppliers200ApplicationJSONSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                                              `json:"contactName,omitempty"`
	DefaultCurrency    *string                                                              `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                                              `json:"emailAddress,omitempty"`
	ID                 *string                                                              `json:"id,omitempty"`
	Metadata           *CreateSuppliers200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                           `json:"modifiedDate,omitempty"`
	Phone              *string                                                              `json:"phone,omitempty"`
	RegistrationNumber *string                                                              `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                                           `json:"sourceModifiedDate,omitempty"`
	Status             CreateSuppliers200ApplicationJSONSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *CreateSuppliers200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	SupplierName       *string                                                              `json:"supplierName,omitempty"`
	TaxNumber          *string                                                              `json:"taxNumber,omitempty"`
}

type CreateSuppliers200ApplicationJSONStatusEnum string

const (
	CreateSuppliers200ApplicationJSONStatusEnumPending  CreateSuppliers200ApplicationJSONStatusEnum = "Pending"
	CreateSuppliers200ApplicationJSONStatusEnumFailed   CreateSuppliers200ApplicationJSONStatusEnum = "Failed"
	CreateSuppliers200ApplicationJSONStatusEnumSuccess  CreateSuppliers200ApplicationJSONStatusEnum = "Success"
	CreateSuppliers200ApplicationJSONStatusEnumTimedOut CreateSuppliers200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateSuppliers200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateSuppliers200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateSuppliers200ApplicationJSONValidation struct {
	Errors   []CreateSuppliers200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateSuppliers200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateSuppliers200ApplicationJSON struct {
	Changes           []CreateSuppliers200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                               `json:"companyId"`
	CompletedOnUtc    *time.Time                                           `json:"completedOnUtc,omitempty"`
	Data              *CreateSuppliers200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                               `json:"dataConnectionKey"`
	DataType          *string                                              `json:"dataType,omitempty"`
	ErrorMessage      *string                                              `json:"errorMessage,omitempty"`
	PushOperationKey  string                                               `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                            `json:"requestedOnUtc"`
	Status            CreateSuppliers200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                  `json:"statusCode"`
	TimeoutInMinutes  *int                                                 `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                 `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateSuppliers200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateSuppliersResponse struct {
	ContentType                             string
	StatusCode                              int
	RawResponse                             *http.Response
	CreateSuppliers200ApplicationJSONObject *CreateSuppliers200ApplicationJSON
}
