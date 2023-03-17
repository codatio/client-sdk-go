package operations

import (
	"net/http"
	"time"
)

type PutSupplierSourceModifiedDateAddressesTypeEnum string

const (
	PutSupplierSourceModifiedDateAddressesTypeEnumUnknown  PutSupplierSourceModifiedDateAddressesTypeEnum = "Unknown"
	PutSupplierSourceModifiedDateAddressesTypeEnumBilling  PutSupplierSourceModifiedDateAddressesTypeEnum = "Billing"
	PutSupplierSourceModifiedDateAddressesTypeEnumDelivery PutSupplierSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type PutSupplierSourceModifiedDateAddresses struct {
	City       *string                                        `json:"city,omitempty"`
	Country    *string                                        `json:"country,omitempty"`
	Line1      *string                                        `json:"line1,omitempty"`
	Line2      *string                                        `json:"line2,omitempty"`
	PostalCode *string                                        `json:"postalCode,omitempty"`
	Region     *string                                        `json:"region,omitempty"`
	Type       PutSupplierSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type PutSupplierSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PutSupplierSourceModifiedDateStatusEnum string

const (
	PutSupplierSourceModifiedDateStatusEnumUnknown  PutSupplierSourceModifiedDateStatusEnum = "Unknown"
	PutSupplierSourceModifiedDateStatusEnumActive   PutSupplierSourceModifiedDateStatusEnum = "Active"
	PutSupplierSourceModifiedDateStatusEnumArchived PutSupplierSourceModifiedDateStatusEnum = "Archived"
)

// PutSupplierSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type PutSupplierSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PutSupplierSourceModifiedDate
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type PutSupplierSourceModifiedDate struct {
	Addresses          []PutSupplierSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                        `json:"contactName,omitempty"`
	DefaultCurrency    *string                                        `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                        `json:"emailAddress,omitempty"`
	ID                 *string                                        `json:"id,omitempty"`
	Metadata           *PutSupplierSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                     `json:"modifiedDate,omitempty"`
	Phone              *string                                        `json:"phone,omitempty"`
	RegistrationNumber *string                                        `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                     `json:"sourceModifiedDate,omitempty"`
	Status             PutSupplierSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *PutSupplierSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	SupplierName       *string                                        `json:"supplierName,omitempty"`
	TaxNumber          *string                                        `json:"taxNumber,omitempty"`
}

type PutSupplierRequest struct {
	RequestBody      *PutSupplierSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID        string                         `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                         `pathParam:"style=simple,explode=false,name=connectionId"`
	ForceUpdate      *bool                          `queryParam:"style=form,explode=true,name=forceUpdate"`
	SupplierID       string                         `pathParam:"style=simple,explode=false,name=supplierId"`
	TimeoutInMinutes *int                           `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type PutSupplier200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PutSupplier200ApplicationJSONChangesTypeEnum string

const (
	PutSupplier200ApplicationJSONChangesTypeEnumUnknown            PutSupplier200ApplicationJSONChangesTypeEnum = "Unknown"
	PutSupplier200ApplicationJSONChangesTypeEnumCreated            PutSupplier200ApplicationJSONChangesTypeEnum = "Created"
	PutSupplier200ApplicationJSONChangesTypeEnumModified           PutSupplier200ApplicationJSONChangesTypeEnum = "Modified"
	PutSupplier200ApplicationJSONChangesTypeEnumDeleted            PutSupplier200ApplicationJSONChangesTypeEnum = "Deleted"
	PutSupplier200ApplicationJSONChangesTypeEnumAttachmentUploaded PutSupplier200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PutSupplier200ApplicationJSONChanges struct {
	AttachmentID *string                                                     `json:"attachmentId,omitempty"`
	RecordRef    *PutSupplier200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PutSupplier200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

type PutSupplier200ApplicationJSONSourceModifiedDateAddressesTypeEnum string

const (
	PutSupplier200ApplicationJSONSourceModifiedDateAddressesTypeEnumUnknown  PutSupplier200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Unknown"
	PutSupplier200ApplicationJSONSourceModifiedDateAddressesTypeEnumBilling  PutSupplier200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Billing"
	PutSupplier200ApplicationJSONSourceModifiedDateAddressesTypeEnumDelivery PutSupplier200ApplicationJSONSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type PutSupplier200ApplicationJSONSourceModifiedDateAddresses struct {
	City       *string                                                          `json:"city,omitempty"`
	Country    *string                                                          `json:"country,omitempty"`
	Line1      *string                                                          `json:"line1,omitempty"`
	Line2      *string                                                          `json:"line2,omitempty"`
	PostalCode *string                                                          `json:"postalCode,omitempty"`
	Region     *string                                                          `json:"region,omitempty"`
	Type       PutSupplier200ApplicationJSONSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type PutSupplier200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PutSupplier200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PutSupplier200ApplicationJSONSourceModifiedDateStatusEnumUnknown  PutSupplier200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PutSupplier200ApplicationJSONSourceModifiedDateStatusEnumActive   PutSupplier200ApplicationJSONSourceModifiedDateStatusEnum = "Active"
	PutSupplier200ApplicationJSONSourceModifiedDateStatusEnumArchived PutSupplier200ApplicationJSONSourceModifiedDateStatusEnum = "Archived"
)

// PutSupplier200ApplicationJSONSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type PutSupplier200ApplicationJSONSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// PutSupplier200ApplicationJSONSourceModifiedDate
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type PutSupplier200ApplicationJSONSourceModifiedDate struct {
	Addresses          []PutSupplier200ApplicationJSONSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                                          `json:"contactName,omitempty"`
	DefaultCurrency    *string                                                          `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                                          `json:"emailAddress,omitempty"`
	ID                 *string                                                          `json:"id,omitempty"`
	Metadata           *PutSupplier200ApplicationJSONSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                       `json:"modifiedDate,omitempty"`
	Phone              *string                                                          `json:"phone,omitempty"`
	RegistrationNumber *string                                                          `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                                       `json:"sourceModifiedDate,omitempty"`
	Status             PutSupplier200ApplicationJSONSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *PutSupplier200ApplicationJSONSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	SupplierName       *string                                                          `json:"supplierName,omitempty"`
	TaxNumber          *string                                                          `json:"taxNumber,omitempty"`
}

type PutSupplier200ApplicationJSONStatusEnum string

const (
	PutSupplier200ApplicationJSONStatusEnumPending  PutSupplier200ApplicationJSONStatusEnum = "Pending"
	PutSupplier200ApplicationJSONStatusEnumFailed   PutSupplier200ApplicationJSONStatusEnum = "Failed"
	PutSupplier200ApplicationJSONStatusEnumSuccess  PutSupplier200ApplicationJSONStatusEnum = "Success"
	PutSupplier200ApplicationJSONStatusEnumTimedOut PutSupplier200ApplicationJSONStatusEnum = "TimedOut"
)

type PutSupplier200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PutSupplier200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PutSupplier200ApplicationJSONValidation struct {
	Errors   []PutSupplier200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PutSupplier200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PutSupplier200ApplicationJSON struct {
	Changes           []PutSupplier200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                           `json:"companyId"`
	CompletedOnUtc    *time.Time                                       `json:"completedOnUtc,omitempty"`
	Data              *PutSupplier200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                           `json:"dataConnectionKey"`
	DataType          *string                                          `json:"dataType,omitempty"`
	ErrorMessage      *string                                          `json:"errorMessage,omitempty"`
	PushOperationKey  string                                           `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                        `json:"requestedOnUtc"`
	Status            PutSupplier200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                              `json:"statusCode"`
	TimeoutInMinutes  *int                                             `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                             `json:"timeoutInSeconds,omitempty"`
	Validation        *PutSupplier200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PutSupplierResponse struct {
	ContentType                         string
	StatusCode                          int
	RawResponse                         *http.Response
	PutSupplier200ApplicationJSONObject *PutSupplier200ApplicationJSON
}
