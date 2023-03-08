package operations

import (
	"net/http"
	"time"
)

type PostSuppliersPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostSuppliersQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PostSuppliersSourceModifiedDateAddressesTypeEnum string

const (
	PostSuppliersSourceModifiedDateAddressesTypeEnumUnknown  PostSuppliersSourceModifiedDateAddressesTypeEnum = "Unknown"
	PostSuppliersSourceModifiedDateAddressesTypeEnumBilling  PostSuppliersSourceModifiedDateAddressesTypeEnum = "Billing"
	PostSuppliersSourceModifiedDateAddressesTypeEnumDelivery PostSuppliersSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type PostSuppliersSourceModifiedDateAddresses struct {
	City       *string                                          `json:"city,omitempty"`
	Country    *string                                          `json:"country,omitempty"`
	Line1      *string                                          `json:"line1,omitempty"`
	Line2      *string                                          `json:"line2,omitempty"`
	PostalCode *string                                          `json:"postalCode,omitempty"`
	Region     *string                                          `json:"region,omitempty"`
	Type       PostSuppliersSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type PostSuppliersSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostSuppliersSourceModifiedDateStatusEnum string

const (
	PostSuppliersSourceModifiedDateStatusEnumUnknown  PostSuppliersSourceModifiedDateStatusEnum = "Unknown"
	PostSuppliersSourceModifiedDateStatusEnumActive   PostSuppliersSourceModifiedDateStatusEnum = "Active"
	PostSuppliersSourceModifiedDateStatusEnumArchived PostSuppliersSourceModifiedDateStatusEnum = "Archived"
)

// PostSuppliersSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type PostSuppliersSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostSuppliersSourceModifiedDate
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type PostSuppliersSourceModifiedDate struct {
	Addresses          []PostSuppliersSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                          `json:"contactName,omitempty"`
	DefaultCurrency    *string                                          `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                          `json:"emailAddress,omitempty"`
	ID                 *string                                          `json:"id,omitempty"`
	Metadata           *PostSuppliersSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                       `json:"modifiedDate,omitempty"`
	Phone              *string                                          `json:"phone,omitempty"`
	RegistrationNumber *string                                          `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                       `json:"sourceModifiedDate,omitempty"`
	Status             PostSuppliersSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *PostSuppliersSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	SupplierName       *string                                          `json:"supplierName,omitempty"`
	TaxNumber          *string                                          `json:"taxNumber,omitempty"`
}

type PostSuppliersRequest struct {
	PathParams  PostSuppliersPathParams
	QueryParams PostSuppliersQueryParams
	Request     *PostSuppliersSourceModifiedDate `request:"mediaType=application/json"`
}

type PostSuppliers200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostSuppliers200ApplicationJSONChangesTypeEnum string

const (
	PostSuppliers200ApplicationJSONChangesTypeEnumUnknown            PostSuppliers200ApplicationJSONChangesTypeEnum = "Unknown"
	PostSuppliers200ApplicationJSONChangesTypeEnumCreated            PostSuppliers200ApplicationJSONChangesTypeEnum = "Created"
	PostSuppliers200ApplicationJSONChangesTypeEnumModified           PostSuppliers200ApplicationJSONChangesTypeEnum = "Modified"
	PostSuppliers200ApplicationJSONChangesTypeEnumDeleted            PostSuppliers200ApplicationJSONChangesTypeEnum = "Deleted"
	PostSuppliers200ApplicationJSONChangesTypeEnumAttachmentUploaded PostSuppliers200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostSuppliers200ApplicationJSONChanges struct {
	AttachmentID *string                                                       `json:"attachmentId,omitempty"`
	RecordRef    *PostSuppliers200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostSuppliers200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PostSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum string

const (
	PostSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumUnknown  PostSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Unknown"
	PostSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumBilling  PostSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Billing"
	PostSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnumDelivery PostSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type PostSuppliers200ApplicationJSONSourceModifiedDateAddresses struct {
	City       *string                                                            `json:"city,omitempty"`
	Country    *string                                                            `json:"country,omitempty"`
	Line1      *string                                                            `json:"line1,omitempty"`
	Line2      *string                                                            `json:"line2,omitempty"`
	PostalCode *string                                                            `json:"postalCode,omitempty"`
	Region     *string                                                            `json:"region,omitempty"`
	Type       PostSuppliers200ApplicationJSONSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type PostSuppliers200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostSuppliers200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PostSuppliers200ApplicationJSONSourceModifiedDateStatusEnumUnknown  PostSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PostSuppliers200ApplicationJSONSourceModifiedDateStatusEnumActive   PostSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	PostSuppliers200ApplicationJSONSourceModifiedDateStatusEnumArchived PostSuppliers200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

// PostSuppliers200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type PostSuppliers200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PostSuppliers200ApplicationJSONSourceModifiedDate
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type PostSuppliers200ApplicationJSONSourceModifiedDate struct {
	Addresses          []PostSuppliers200ApplicationJSONSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                                            `json:"contactName,omitempty"`
	DefaultCurrency    *string                                                            `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                                            `json:"emailAddress,omitempty"`
	ID                 *string                                                            `json:"id,omitempty"`
	Metadata           *PostSuppliers200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                         `json:"modifiedDate,omitempty"`
	Phone              *string                                                            `json:"phone,omitempty"`
	RegistrationNumber *string                                                            `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                                         `json:"sourceModifiedDate,omitempty"`
	Status             PostSuppliers200ApplicationJSONSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *PostSuppliers200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	SupplierName       *string                                                            `json:"supplierName,omitempty"`
	TaxNumber          *string                                                            `json:"taxNumber,omitempty"`
}

type PostSuppliers200ApplicationJSONStatusEnum string

const (
	PostSuppliers200ApplicationJSONStatusEnumPending  PostSuppliers200ApplicationJSONStatusEnum = "Pending"
	PostSuppliers200ApplicationJSONStatusEnumFailed   PostSuppliers200ApplicationJSONStatusEnum = "Failed"
	PostSuppliers200ApplicationJSONStatusEnumSuccess  PostSuppliers200ApplicationJSONStatusEnum = "Success"
	PostSuppliers200ApplicationJSONStatusEnumTimedOut PostSuppliers200ApplicationJSONStatusEnum = "TimedOut"
)

type PostSuppliers200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostSuppliers200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostSuppliers200ApplicationJSONValidation struct {
	Errors   []PostSuppliers200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostSuppliers200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostSuppliers200ApplicationJSON struct {
	Changes           []PostSuppliers200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                             `json:"companyId"`
	CompletedOnUtc    *time.Time                                         `json:"completedOnUtc,omitempty"`
	Data              *PostSuppliers200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                             `json:"dataConnectionKey"`
	DataType          *string                                            `json:"dataType,omitempty"`
	ErrorMessage      *string                                            `json:"errorMessage,omitempty"`
	PushOperationKey  string                                             `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                          `json:"requestedOnUtc"`
	Status            PostSuppliers200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                `json:"statusCode"`
	TimeoutInMinutes  *int                                               `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                               `json:"timeoutInSeconds,omitempty"`
	Validation        *PostSuppliers200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostSuppliersResponse struct {
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	PostSuppliers200ApplicationJSONObject *PostSuppliers200ApplicationJSON
}
