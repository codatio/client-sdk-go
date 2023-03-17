package operations

import (
	"net/http"
	"time"
)

// CreateItemSourceModifiedDateBillItemAccountRef
// Reference of the account to which the item is linked.
type CreateItemSourceModifiedDateBillItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateItemSourceModifiedDateBillItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type CreateItemSourceModifiedDateBillItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateItemSourceModifiedDateBillItem
// Item details that are only for bills.
type CreateItemSourceModifiedDateBillItem struct {
	AccountRef  *CreateItemSourceModifiedDateBillItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                         `json:"description,omitempty"`
	TaxRateRef  *CreateItemSourceModifiedDateBillItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                        `json:"unitPrice,omitempty"`
}

// CreateItemSourceModifiedDateInvoiceItemAccountRef
// Reference of the account to which the item is linked.
type CreateItemSourceModifiedDateInvoiceItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateItemSourceModifiedDateInvoiceItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type CreateItemSourceModifiedDateInvoiceItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateItemSourceModifiedDateInvoiceItem
// Item details that are only for bills.
type CreateItemSourceModifiedDateInvoiceItem struct {
	AccountRef  *CreateItemSourceModifiedDateInvoiceItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                            `json:"description,omitempty"`
	TaxRateRef  *CreateItemSourceModifiedDateInvoiceItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                           `json:"unitPrice,omitempty"`
}

type CreateItemSourceModifiedDateItemStatusEnum string

const (
	CreateItemSourceModifiedDateItemStatusEnumUnknown  CreateItemSourceModifiedDateItemStatusEnum = "Unknown"
	CreateItemSourceModifiedDateItemStatusEnumActive   CreateItemSourceModifiedDateItemStatusEnum = "Active"
	CreateItemSourceModifiedDateItemStatusEnumArchived CreateItemSourceModifiedDateItemStatusEnum = "Archived"
)

type CreateItemSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateItemSourceModifiedDateTypeEnum string

const (
	CreateItemSourceModifiedDateTypeEnumUnknown      CreateItemSourceModifiedDateTypeEnum = "Unknown"
	CreateItemSourceModifiedDateTypeEnumInventory    CreateItemSourceModifiedDateTypeEnum = "Inventory"
	CreateItemSourceModifiedDateTypeEnumNonInventory CreateItemSourceModifiedDateTypeEnum = "NonInventory"
	CreateItemSourceModifiedDateTypeEnumService      CreateItemSourceModifiedDateTypeEnum = "Service"
)

// CreateItemSourceModifiedDate
// > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
type CreateItemSourceModifiedDate struct {
	BillItem           *CreateItemSourceModifiedDateBillItem      `json:"billItem,omitempty"`
	Code               *string                                    `json:"code,omitempty"`
	ID                 *string                                    `json:"id,omitempty"`
	InvoiceItem        *CreateItemSourceModifiedDateInvoiceItem   `json:"invoiceItem,omitempty"`
	IsBillItem         bool                                       `json:"isBillItem"`
	IsInvoiceItem      bool                                       `json:"isInvoiceItem"`
	ItemStatus         CreateItemSourceModifiedDateItemStatusEnum `json:"itemStatus"`
	Metadata           *CreateItemSourceModifiedDateMetadata      `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                 `json:"modifiedDate,omitempty"`
	Name               *string                                    `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                 `json:"sourceModifiedDate,omitempty"`
	Type               CreateItemSourceModifiedDateTypeEnum       `json:"type"`
}

type CreateItemRequest struct {
	RequestBody      *CreateItemSourceModifiedDate `request:"mediaType=application/json"`
	CompanyID        string                        `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID     string                        `pathParam:"style=simple,explode=false,name=connectionId"`
	TimeoutInMinutes *int                          `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

type CreateItem200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreateItem200ApplicationJSONChangesTypeEnum string

const (
	CreateItem200ApplicationJSONChangesTypeEnumUnknown            CreateItem200ApplicationJSONChangesTypeEnum = "Unknown"
	CreateItem200ApplicationJSONChangesTypeEnumCreated            CreateItem200ApplicationJSONChangesTypeEnum = "Created"
	CreateItem200ApplicationJSONChangesTypeEnumModified           CreateItem200ApplicationJSONChangesTypeEnum = "Modified"
	CreateItem200ApplicationJSONChangesTypeEnumDeleted            CreateItem200ApplicationJSONChangesTypeEnum = "Deleted"
	CreateItem200ApplicationJSONChangesTypeEnumAttachmentUploaded CreateItem200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreateItem200ApplicationJSONChanges struct {
	AttachmentID *string                                                    `json:"attachmentId,omitempty"`
	RecordRef    *CreateItem200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreateItem200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreateItem200ApplicationJSONSourceModifiedDateBillItemAccountRef
// Reference of the account to which the item is linked.
type CreateItem200ApplicationJSONSourceModifiedDateBillItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateItem200ApplicationJSONSourceModifiedDateBillItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type CreateItem200ApplicationJSONSourceModifiedDateBillItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateItem200ApplicationJSONSourceModifiedDateBillItem
// Item details that are only for bills.
type CreateItem200ApplicationJSONSourceModifiedDateBillItem struct {
	AccountRef  *CreateItem200ApplicationJSONSourceModifiedDateBillItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                                           `json:"description,omitempty"`
	TaxRateRef  *CreateItem200ApplicationJSONSourceModifiedDateBillItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                                          `json:"unitPrice,omitempty"`
}

// CreateItem200ApplicationJSONSourceModifiedDateInvoiceItemAccountRef
// Reference of the account to which the item is linked.
type CreateItem200ApplicationJSONSourceModifiedDateInvoiceItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateItem200ApplicationJSONSourceModifiedDateInvoiceItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type CreateItem200ApplicationJSONSourceModifiedDateInvoiceItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// CreateItem200ApplicationJSONSourceModifiedDateInvoiceItem
// Item details that are only for bills.
type CreateItem200ApplicationJSONSourceModifiedDateInvoiceItem struct {
	AccountRef  *CreateItem200ApplicationJSONSourceModifiedDateInvoiceItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                                              `json:"description,omitempty"`
	TaxRateRef  *CreateItem200ApplicationJSONSourceModifiedDateInvoiceItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                                             `json:"unitPrice,omitempty"`
}

type CreateItem200ApplicationJSONSourceModifiedDateItemStatusEnum string

const (
	CreateItem200ApplicationJSONSourceModifiedDateItemStatusEnumUnknown  CreateItem200ApplicationJSONSourceModifiedDateItemStatusEnum = "Unknown"
	CreateItem200ApplicationJSONSourceModifiedDateItemStatusEnumActive   CreateItem200ApplicationJSONSourceModifiedDateItemStatusEnum = "Active"
	CreateItem200ApplicationJSONSourceModifiedDateItemStatusEnumArchived CreateItem200ApplicationJSONSourceModifiedDateItemStatusEnum = "Archived"
)

type CreateItem200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreateItem200ApplicationJSONSourceModifiedDateTypeEnum string

const (
	CreateItem200ApplicationJSONSourceModifiedDateTypeEnumUnknown      CreateItem200ApplicationJSONSourceModifiedDateTypeEnum = "Unknown"
	CreateItem200ApplicationJSONSourceModifiedDateTypeEnumInventory    CreateItem200ApplicationJSONSourceModifiedDateTypeEnum = "Inventory"
	CreateItem200ApplicationJSONSourceModifiedDateTypeEnumNonInventory CreateItem200ApplicationJSONSourceModifiedDateTypeEnum = "NonInventory"
	CreateItem200ApplicationJSONSourceModifiedDateTypeEnumService      CreateItem200ApplicationJSONSourceModifiedDateTypeEnum = "Service"
)

// CreateItem200ApplicationJSONSourceModifiedDate
// > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
type CreateItem200ApplicationJSONSourceModifiedDate struct {
	BillItem           *CreateItem200ApplicationJSONSourceModifiedDateBillItem      `json:"billItem,omitempty"`
	Code               *string                                                      `json:"code,omitempty"`
	ID                 *string                                                      `json:"id,omitempty"`
	InvoiceItem        *CreateItem200ApplicationJSONSourceModifiedDateInvoiceItem   `json:"invoiceItem,omitempty"`
	IsBillItem         bool                                                         `json:"isBillItem"`
	IsInvoiceItem      bool                                                         `json:"isInvoiceItem"`
	ItemStatus         CreateItem200ApplicationJSONSourceModifiedDateItemStatusEnum `json:"itemStatus"`
	Metadata           *CreateItem200ApplicationJSONSourceModifiedDateMetadata      `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                   `json:"modifiedDate,omitempty"`
	Name               *string                                                      `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                                   `json:"sourceModifiedDate,omitempty"`
	Type               CreateItem200ApplicationJSONSourceModifiedDateTypeEnum       `json:"type"`
}

type CreateItem200ApplicationJSONStatusEnum string

const (
	CreateItem200ApplicationJSONStatusEnumPending  CreateItem200ApplicationJSONStatusEnum = "Pending"
	CreateItem200ApplicationJSONStatusEnumFailed   CreateItem200ApplicationJSONStatusEnum = "Failed"
	CreateItem200ApplicationJSONStatusEnumSuccess  CreateItem200ApplicationJSONStatusEnum = "Success"
	CreateItem200ApplicationJSONStatusEnumTimedOut CreateItem200ApplicationJSONStatusEnum = "TimedOut"
)

type CreateItem200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreateItem200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreateItem200ApplicationJSONValidation struct {
	Errors   []CreateItem200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreateItem200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreateItem200ApplicationJSON struct {
	Changes           []CreateItem200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                          `json:"companyId"`
	CompletedOnUtc    *time.Time                                      `json:"completedOnUtc,omitempty"`
	Data              *CreateItem200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                          `json:"dataConnectionKey"`
	DataType          *string                                         `json:"dataType,omitempty"`
	ErrorMessage      *string                                         `json:"errorMessage,omitempty"`
	PushOperationKey  string                                          `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                       `json:"requestedOnUtc"`
	Status            CreateItem200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                             `json:"statusCode"`
	TimeoutInMinutes  *int                                            `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                            `json:"timeoutInSeconds,omitempty"`
	Validation        *CreateItem200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreateItemResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	CreateItem200ApplicationJSONObject *CreateItem200ApplicationJSON
}
