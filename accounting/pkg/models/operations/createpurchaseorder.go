package operations

import (
	"net/http"
	"time"
)

type CreatePurchaseOrderPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type CreatePurchaseOrderQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// CreatePurchaseOrderSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreatePurchaseOrderSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreatePurchaseOrderSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type CreatePurchaseOrderSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreatePurchaseOrderSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreatePurchaseOrderSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type CreatePurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreatePurchaseOrderSourceModifiedDateLineItems struct {
	AccountRef           *CreatePurchaseOrderSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                              `json:"description,omitempty"`
	DiscountAmount       *float64                                                             `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                             `json:"discountPercentage,omitempty"`
	ItemRef              *CreatePurchaseOrderSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             *float64                                                             `json:"quantity,omitempty"`
	SubTotal             *float64                                                             `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                             `json:"taxAmount,omitempty"`
	TaxRateRef           *CreatePurchaseOrderSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                             `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []CreatePurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64                                                             `json:"unitAmount,omitempty"`
}

type CreatePurchaseOrderSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum string

const (
	CreatePurchaseOrderSourceModifiedDateShipToAddressTypeEnumUnknown  CreatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	CreatePurchaseOrderSourceModifiedDateShipToAddressTypeEnumBilling  CreatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Billing"
	CreatePurchaseOrderSourceModifiedDateShipToAddressTypeEnumDelivery CreatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// CreatePurchaseOrderSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type CreatePurchaseOrderSourceModifiedDateShipToAddress struct {
	City       *string                                                    `json:"city,omitempty"`
	Country    *string                                                    `json:"country,omitempty"`
	Line1      *string                                                    `json:"line1,omitempty"`
	Line2      *string                                                    `json:"line2,omitempty"`
	PostalCode *string                                                    `json:"postalCode,omitempty"`
	Region     *string                                                    `json:"region,omitempty"`
	Type       CreatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// CreatePurchaseOrderSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type CreatePurchaseOrderSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// CreatePurchaseOrderSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type CreatePurchaseOrderSourceModifiedDateShipTo struct {
	Address *CreatePurchaseOrderSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *CreatePurchaseOrderSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type CreatePurchaseOrderSourceModifiedDateStatusEnum string

const (
	CreatePurchaseOrderSourceModifiedDateStatusEnumUnknown CreatePurchaseOrderSourceModifiedDateStatusEnum = "Unknown"
	CreatePurchaseOrderSourceModifiedDateStatusEnumDraft   CreatePurchaseOrderSourceModifiedDateStatusEnum = "Draft"
	CreatePurchaseOrderSourceModifiedDateStatusEnumOpen    CreatePurchaseOrderSourceModifiedDateStatusEnum = "Open"
	CreatePurchaseOrderSourceModifiedDateStatusEnumClosed  CreatePurchaseOrderSourceModifiedDateStatusEnum = "Closed"
	CreatePurchaseOrderSourceModifiedDateStatusEnumVoid    CreatePurchaseOrderSourceModifiedDateStatusEnum = "Void"
)

// CreatePurchaseOrderSourceModifiedDateSupplierRef
// Supplier that the purchase order is recorded against in the accounting system.
type CreatePurchaseOrderSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// CreatePurchaseOrderSourceModifiedDate
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type CreatePurchaseOrderSourceModifiedDate struct {
	Currency             *string                                           `json:"currency,omitempty"`
	CurrencyRate         *float64                                          `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time                                        `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time                                        `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                                           `json:"id,omitempty"`
	IssueDate            *time.Time                                        `json:"issueDate,omitempty"`
	LineItems            []CreatePurchaseOrderSourceModifiedDateLineItems  `json:"lineItems,omitempty"`
	Metadata             *CreatePurchaseOrderSourceModifiedDateMetadata    `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                        `json:"modifiedDate,omitempty"`
	Note                 *string                                           `json:"note,omitempty"`
	PaymentDueDate       *time.Time                                        `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                                           `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *CreatePurchaseOrderSourceModifiedDateShipTo      `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time                                        `json:"sourceModifiedDate,omitempty"`
	Status               *CreatePurchaseOrderSourceModifiedDateStatusEnum  `json:"status,omitempty"`
	SubTotal             *float64                                          `json:"subTotal,omitempty"`
	SupplierRef          *CreatePurchaseOrderSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	TotalAmount          *float64                                          `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                                          `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                                          `json:"totalTaxAmount,omitempty"`
}

type CreatePurchaseOrderRequest struct {
	PathParams  CreatePurchaseOrderPathParams
	QueryParams CreatePurchaseOrderQueryParams
	Request     *CreatePurchaseOrderSourceModifiedDate `request:"mediaType=application/json"`
}

type CreatePurchaseOrder200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type CreatePurchaseOrder200ApplicationJSONChangesTypeEnum string

const (
	CreatePurchaseOrder200ApplicationJSONChangesTypeEnumUnknown            CreatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Unknown"
	CreatePurchaseOrder200ApplicationJSONChangesTypeEnumCreated            CreatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Created"
	CreatePurchaseOrder200ApplicationJSONChangesTypeEnumModified           CreatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Modified"
	CreatePurchaseOrder200ApplicationJSONChangesTypeEnumDeleted            CreatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Deleted"
	CreatePurchaseOrder200ApplicationJSONChangesTypeEnumAttachmentUploaded CreatePurchaseOrder200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type CreatePurchaseOrder200ApplicationJSONChanges struct {
	AttachmentID *string                                                             `json:"attachmentId,omitempty"`
	RecordRef    *CreatePurchaseOrder200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *CreatePurchaseOrder200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                                `json:"description,omitempty"`
	DiscountAmount       *float64                                                                               `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                               `json:"discountPercentage,omitempty"`
	ItemRef              *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             *float64                                                                               `json:"quantity,omitempty"`
	SubTotal             *float64                                                                               `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                               `json:"taxAmount,omitempty"`
	TaxRateRef           *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                               `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64                                                                               `json:"unitAmount,omitempty"`
}

type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum string

const (
	CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumUnknown  CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumBilling  CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Billing"
	CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumDelivery CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress struct {
	City       *string                                                                      `json:"city,omitempty"`
	Country    *string                                                                      `json:"country,omitempty"`
	Line1      *string                                                                      `json:"line1,omitempty"`
	Line2      *string                                                                      `json:"line2,omitempty"`
	PostalCode *string                                                                      `json:"postalCode,omitempty"`
	Region     *string                                                                      `json:"region,omitempty"`
	Type       CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipTo struct {
	Address *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumUnknown CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumDraft   CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumOpen    CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Open"
	CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumClosed  CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Closed"
	CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumVoid    CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
)

// CreatePurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef
// Supplier that the purchase order is recorded against in the accounting system.
type CreatePurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// CreatePurchaseOrder200ApplicationJSONSourceModifiedDate
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type CreatePurchaseOrder200ApplicationJSONSourceModifiedDate struct {
	Currency             *string                                                             `json:"currency,omitempty"`
	CurrencyRate         *float64                                                            `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time                                                          `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time                                                          `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                                                             `json:"id,omitempty"`
	IssueDate            *time.Time                                                          `json:"issueDate,omitempty"`
	LineItems            []CreatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItems  `json:"lineItems,omitempty"`
	Metadata             *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateMetadata    `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                          `json:"modifiedDate,omitempty"`
	Note                 *string                                                             `json:"note,omitempty"`
	PaymentDueDate       *time.Time                                                          `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                                                             `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateShipTo      `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time                                                          `json:"sourceModifiedDate,omitempty"`
	Status               *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum  `json:"status,omitempty"`
	SubTotal             *float64                                                            `json:"subTotal,omitempty"`
	SupplierRef          *CreatePurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	TotalAmount          *float64                                                            `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                                                            `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                                                            `json:"totalTaxAmount,omitempty"`
}

type CreatePurchaseOrder200ApplicationJSONStatusEnum string

const (
	CreatePurchaseOrder200ApplicationJSONStatusEnumPending  CreatePurchaseOrder200ApplicationJSONStatusEnum = "Pending"
	CreatePurchaseOrder200ApplicationJSONStatusEnumFailed   CreatePurchaseOrder200ApplicationJSONStatusEnum = "Failed"
	CreatePurchaseOrder200ApplicationJSONStatusEnumSuccess  CreatePurchaseOrder200ApplicationJSONStatusEnum = "Success"
	CreatePurchaseOrder200ApplicationJSONStatusEnumTimedOut CreatePurchaseOrder200ApplicationJSONStatusEnum = "TimedOut"
)

type CreatePurchaseOrder200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// CreatePurchaseOrder200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type CreatePurchaseOrder200ApplicationJSONValidation struct {
	Errors   []CreatePurchaseOrder200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []CreatePurchaseOrder200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type CreatePurchaseOrder200ApplicationJSON struct {
	Changes           []CreatePurchaseOrder200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                   `json:"companyId"`
	CompletedOnUtc    *time.Time                                               `json:"completedOnUtc,omitempty"`
	Data              *CreatePurchaseOrder200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                   `json:"dataConnectionKey"`
	DataType          *string                                                  `json:"dataType,omitempty"`
	ErrorMessage      *string                                                  `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                   `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                                `json:"requestedOnUtc"`
	Status            CreatePurchaseOrder200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                      `json:"statusCode"`
	TimeoutInMinutes  *int                                                     `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                     `json:"timeoutInSeconds,omitempty"`
	Validation        *CreatePurchaseOrder200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type CreatePurchaseOrderResponse struct {
	ContentType                                 string
	StatusCode                                  int
	RawResponse                                 *http.Response
	CreatePurchaseOrder200ApplicationJSONObject *CreatePurchaseOrder200ApplicationJSON
}
