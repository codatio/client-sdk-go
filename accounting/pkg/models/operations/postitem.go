package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type PostItemPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostItemQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostItemSourceModifiedDateBillItemAccountRef
// Reference of the account to which the item is linked.
type PostItemSourceModifiedDateBillItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostItemSourceModifiedDateBillItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type PostItemSourceModifiedDateBillItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// PostItemSourceModifiedDateBillItem
// Item details that are only for bills.
type PostItemSourceModifiedDateBillItem struct {
	AccountRef  *PostItemSourceModifiedDateBillItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                       `json:"description,omitempty"`
	TaxRateRef  *PostItemSourceModifiedDateBillItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                      `json:"unitPrice,omitempty"`
}

// PostItemSourceModifiedDateInvoiceItemAccountRef
// Reference of the account to which the item is linked.
type PostItemSourceModifiedDateInvoiceItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostItemSourceModifiedDateInvoiceItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type PostItemSourceModifiedDateInvoiceItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// PostItemSourceModifiedDateInvoiceItem
// Item details that are only for bills.
type PostItemSourceModifiedDateInvoiceItem struct {
	AccountRef  *PostItemSourceModifiedDateInvoiceItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                          `json:"description,omitempty"`
	TaxRateRef  *PostItemSourceModifiedDateInvoiceItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                         `json:"unitPrice,omitempty"`
}

type PostItemSourceModifiedDateItemStatusEnum string

const (
	PostItemSourceModifiedDateItemStatusEnumUnknown  PostItemSourceModifiedDateItemStatusEnum = "Unknown"
	PostItemSourceModifiedDateItemStatusEnumActive   PostItemSourceModifiedDateItemStatusEnum = "Active"
	PostItemSourceModifiedDateItemStatusEnumArchived PostItemSourceModifiedDateItemStatusEnum = "Archived"
)

type PostItemSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostItemSourceModifiedDateTypeEnum string

const (
	PostItemSourceModifiedDateTypeEnumUnknown      PostItemSourceModifiedDateTypeEnum = "Unknown"
	PostItemSourceModifiedDateTypeEnumInventory    PostItemSourceModifiedDateTypeEnum = "Inventory"
	PostItemSourceModifiedDateTypeEnumNonInventory PostItemSourceModifiedDateTypeEnum = "NonInventory"
	PostItemSourceModifiedDateTypeEnumService      PostItemSourceModifiedDateTypeEnum = "Service"
)

// PostItemSourceModifiedDate
// > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
type PostItemSourceModifiedDate struct {
	BillItem           *PostItemSourceModifiedDateBillItem      `json:"billItem,omitempty"`
	Code               *string                                  `json:"code,omitempty"`
	ID                 *string                                  `json:"id,omitempty"`
	InvoiceItem        *PostItemSourceModifiedDateInvoiceItem   `json:"invoiceItem,omitempty"`
	IsBillItem         bool                                     `json:"isBillItem"`
	IsInvoiceItem      bool                                     `json:"isInvoiceItem"`
	ItemStatus         PostItemSourceModifiedDateItemStatusEnum `json:"itemStatus"`
	Metadata           *PostItemSourceModifiedDateMetadata      `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                               `json:"modifiedDate,omitempty"`
	Name               *string                                  `json:"name,omitempty"`
	SourceModifiedDate *time.Time                               `json:"sourceModifiedDate,omitempty"`
	Type               PostItemSourceModifiedDateTypeEnum       `json:"type"`
}

type PostItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostItemRequest struct {
	PathParams  PostItemPathParams
	QueryParams PostItemQueryParams
	Request     *PostItemSourceModifiedDate `request:"mediaType=application/json"`
	Security    PostItemSecurity
}

type PostItem200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostItem200ApplicationJSONChangesTypeEnum string

const (
	PostItem200ApplicationJSONChangesTypeEnumUnknown            PostItem200ApplicationJSONChangesTypeEnum = "Unknown"
	PostItem200ApplicationJSONChangesTypeEnumCreated            PostItem200ApplicationJSONChangesTypeEnum = "Created"
	PostItem200ApplicationJSONChangesTypeEnumModified           PostItem200ApplicationJSONChangesTypeEnum = "Modified"
	PostItem200ApplicationJSONChangesTypeEnumDeleted            PostItem200ApplicationJSONChangesTypeEnum = "Deleted"
	PostItem200ApplicationJSONChangesTypeEnumAttachmentUploaded PostItem200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostItem200ApplicationJSONChanges struct {
	AttachmentID *string                                                  `json:"attachmentId,omitempty"`
	RecordRef    *PostItem200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostItem200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostItem200ApplicationJSONSourceModifiedDateBillItemAccountRef
// Reference of the account to which the item is linked.
type PostItem200ApplicationJSONSourceModifiedDateBillItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostItem200ApplicationJSONSourceModifiedDateBillItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type PostItem200ApplicationJSONSourceModifiedDateBillItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// PostItem200ApplicationJSONSourceModifiedDateBillItem
// Item details that are only for bills.
type PostItem200ApplicationJSONSourceModifiedDateBillItem struct {
	AccountRef  *PostItem200ApplicationJSONSourceModifiedDateBillItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                                         `json:"description,omitempty"`
	TaxRateRef  *PostItem200ApplicationJSONSourceModifiedDateBillItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                                        `json:"unitPrice,omitempty"`
}

// PostItem200ApplicationJSONSourceModifiedDateInvoiceItemAccountRef
// Reference of the account to which the item is linked.
type PostItem200ApplicationJSONSourceModifiedDateInvoiceItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostItem200ApplicationJSONSourceModifiedDateInvoiceItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type PostItem200ApplicationJSONSourceModifiedDateInvoiceItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// PostItem200ApplicationJSONSourceModifiedDateInvoiceItem
// Item details that are only for bills.
type PostItem200ApplicationJSONSourceModifiedDateInvoiceItem struct {
	AccountRef  *PostItem200ApplicationJSONSourceModifiedDateInvoiceItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                                            `json:"description,omitempty"`
	TaxRateRef  *PostItem200ApplicationJSONSourceModifiedDateInvoiceItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                                           `json:"unitPrice,omitempty"`
}

type PostItem200ApplicationJSONSourceModifiedDateItemStatusEnum string

const (
	PostItem200ApplicationJSONSourceModifiedDateItemStatusEnumUnknown  PostItem200ApplicationJSONSourceModifiedDateItemStatusEnum = "Unknown"
	PostItem200ApplicationJSONSourceModifiedDateItemStatusEnumActive   PostItem200ApplicationJSONSourceModifiedDateItemStatusEnum = "Active"
	PostItem200ApplicationJSONSourceModifiedDateItemStatusEnumArchived PostItem200ApplicationJSONSourceModifiedDateItemStatusEnum = "Archived"
)

type PostItem200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostItem200ApplicationJSONSourceModifiedDateTypeEnum string

const (
	PostItem200ApplicationJSONSourceModifiedDateTypeEnumUnknown      PostItem200ApplicationJSONSourceModifiedDateTypeEnum = "Unknown"
	PostItem200ApplicationJSONSourceModifiedDateTypeEnumInventory    PostItem200ApplicationJSONSourceModifiedDateTypeEnum = "Inventory"
	PostItem200ApplicationJSONSourceModifiedDateTypeEnumNonInventory PostItem200ApplicationJSONSourceModifiedDateTypeEnum = "NonInventory"
	PostItem200ApplicationJSONSourceModifiedDateTypeEnumService      PostItem200ApplicationJSONSourceModifiedDateTypeEnum = "Service"
)

// PostItem200ApplicationJSONSourceModifiedDate
// > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
type PostItem200ApplicationJSONSourceModifiedDate struct {
	BillItem           *PostItem200ApplicationJSONSourceModifiedDateBillItem      `json:"billItem,omitempty"`
	Code               *string                                                    `json:"code,omitempty"`
	ID                 *string                                                    `json:"id,omitempty"`
	InvoiceItem        *PostItem200ApplicationJSONSourceModifiedDateInvoiceItem   `json:"invoiceItem,omitempty"`
	IsBillItem         bool                                                       `json:"isBillItem"`
	IsInvoiceItem      bool                                                       `json:"isInvoiceItem"`
	ItemStatus         PostItem200ApplicationJSONSourceModifiedDateItemStatusEnum `json:"itemStatus"`
	Metadata           *PostItem200ApplicationJSONSourceModifiedDateMetadata      `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                                 `json:"modifiedDate,omitempty"`
	Name               *string                                                    `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	Type               PostItem200ApplicationJSONSourceModifiedDateTypeEnum       `json:"type"`
}

type PostItem200ApplicationJSONStatusEnum string

const (
	PostItem200ApplicationJSONStatusEnumPending  PostItem200ApplicationJSONStatusEnum = "Pending"
	PostItem200ApplicationJSONStatusEnumFailed   PostItem200ApplicationJSONStatusEnum = "Failed"
	PostItem200ApplicationJSONStatusEnumSuccess  PostItem200ApplicationJSONStatusEnum = "Success"
	PostItem200ApplicationJSONStatusEnumTimedOut PostItem200ApplicationJSONStatusEnum = "TimedOut"
)

type PostItem200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostItem200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostItem200ApplicationJSONValidation struct {
	Errors   []PostItem200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostItem200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostItem200ApplicationJSON struct {
	Changes           []PostItem200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                        `json:"companyId"`
	CompletedOnUtc    *time.Time                                    `json:"completedOnUtc,omitempty"`
	Data              *PostItem200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                        `json:"dataConnectionKey"`
	DataType          *string                                       `json:"dataType,omitempty"`
	ErrorMessage      *string                                       `json:"errorMessage,omitempty"`
	PushOperationKey  string                                        `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                     `json:"requestedOnUtc"`
	Status            PostItem200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                           `json:"statusCode"`
	TimeoutInMinutes  *int                                          `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                          `json:"timeoutInSeconds,omitempty"`
	Validation        *PostItem200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostItemResponse struct {
	ContentType                      string
	StatusCode                       int
	PostItem200ApplicationJSONObject *PostItem200ApplicationJSON
}
