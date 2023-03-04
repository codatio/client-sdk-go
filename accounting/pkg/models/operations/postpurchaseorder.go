package operations

import (
	"net/http"
	"time"
)

type PostPurchaseOrderPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PostPurchaseOrderQueryParams struct {
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PostPurchaseOrderSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostPurchaseOrderSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostPurchaseOrderSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type PostPurchaseOrderSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostPurchaseOrderSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostPurchaseOrderSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostPurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostPurchaseOrderSourceModifiedDateLineItems struct {
	AccountRef           *PostPurchaseOrderSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                            `json:"description,omitempty"`
	DiscountAmount       *float64                                                           `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                           `json:"discountPercentage,omitempty"`
	ItemRef              *PostPurchaseOrderSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             *float64                                                           `json:"quantity,omitempty"`
	SubTotal             *float64                                                           `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                           `json:"taxAmount,omitempty"`
	TaxRateRef           *PostPurchaseOrderSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                           `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []PostPurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64                                                           `json:"unitAmount,omitempty"`
}

type PostPurchaseOrderSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostPurchaseOrderSourceModifiedDateShipToAddressTypeEnum string

const (
	PostPurchaseOrderSourceModifiedDateShipToAddressTypeEnumUnknown  PostPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	PostPurchaseOrderSourceModifiedDateShipToAddressTypeEnumBilling  PostPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Billing"
	PostPurchaseOrderSourceModifiedDateShipToAddressTypeEnumDelivery PostPurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// PostPurchaseOrderSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type PostPurchaseOrderSourceModifiedDateShipToAddress struct {
	City       *string                                                  `json:"city,omitempty"`
	Country    *string                                                  `json:"country,omitempty"`
	Line1      *string                                                  `json:"line1,omitempty"`
	Line2      *string                                                  `json:"line2,omitempty"`
	PostalCode *string                                                  `json:"postalCode,omitempty"`
	Region     *string                                                  `json:"region,omitempty"`
	Type       PostPurchaseOrderSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// PostPurchaseOrderSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type PostPurchaseOrderSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// PostPurchaseOrderSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type PostPurchaseOrderSourceModifiedDateShipTo struct {
	Address *PostPurchaseOrderSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *PostPurchaseOrderSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type PostPurchaseOrderSourceModifiedDateStatusEnum string

const (
	PostPurchaseOrderSourceModifiedDateStatusEnumUnknown PostPurchaseOrderSourceModifiedDateStatusEnum = "Unknown"
	PostPurchaseOrderSourceModifiedDateStatusEnumDraft   PostPurchaseOrderSourceModifiedDateStatusEnum = "Draft"
	PostPurchaseOrderSourceModifiedDateStatusEnumOpen    PostPurchaseOrderSourceModifiedDateStatusEnum = "Open"
	PostPurchaseOrderSourceModifiedDateStatusEnumClosed  PostPurchaseOrderSourceModifiedDateStatusEnum = "Closed"
	PostPurchaseOrderSourceModifiedDateStatusEnumVoid    PostPurchaseOrderSourceModifiedDateStatusEnum = "Void"
)

// PostPurchaseOrderSourceModifiedDateSupplierRef
// Supplier that the purchase order is recorded against in the accounting system.
type PostPurchaseOrderSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// PostPurchaseOrderSourceModifiedDate
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type PostPurchaseOrderSourceModifiedDate struct {
	Currency             *string                                         `json:"currency,omitempty"`
	CurrencyRate         *float64                                        `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time                                      `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time                                      `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                                         `json:"id,omitempty"`
	IssueDate            *time.Time                                      `json:"issueDate,omitempty"`
	LineItems            []PostPurchaseOrderSourceModifiedDateLineItems  `json:"lineItems,omitempty"`
	Metadata             *PostPurchaseOrderSourceModifiedDateMetadata    `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                      `json:"modifiedDate,omitempty"`
	Note                 *string                                         `json:"note,omitempty"`
	PaymentDueDate       *time.Time                                      `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                                         `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *PostPurchaseOrderSourceModifiedDateShipTo      `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time                                      `json:"sourceModifiedDate,omitempty"`
	Status               *PostPurchaseOrderSourceModifiedDateStatusEnum  `json:"status,omitempty"`
	SubTotal             *float64                                        `json:"subTotal,omitempty"`
	SupplierRef          *PostPurchaseOrderSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	TotalAmount          *float64                                        `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                                        `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                                        `json:"totalTaxAmount,omitempty"`
}

type PostPurchaseOrderRequest struct {
	PathParams  PostPurchaseOrderPathParams
	QueryParams PostPurchaseOrderQueryParams
	Request     *PostPurchaseOrderSourceModifiedDate `request:"mediaType=application/json"`
}

type PostPurchaseOrder200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONChangesTypeEnum string

const (
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumUnknown            PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "Unknown"
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumCreated            PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "Created"
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumModified           PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "Modified"
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumDeleted            PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "Deleted"
	PostPurchaseOrder200ApplicationJSONChangesTypeEnumAttachmentUploaded PostPurchaseOrder200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type PostPurchaseOrder200ApplicationJSONChanges struct {
	AttachmentID *string                                                           `json:"attachmentId,omitempty"`
	RecordRef    *PostPurchaseOrder200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *PostPurchaseOrder200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                              `json:"description,omitempty"`
	DiscountAmount       *float64                                                                             `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                             `json:"discountPercentage,omitempty"`
	ItemRef              *PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             *float64                                                                             `json:"quantity,omitempty"`
	SubTotal             *float64                                                                             `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                             `json:"taxAmount,omitempty"`
	TaxRateRef           *PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                             `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64                                                                             `json:"unitAmount,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum string

const (
	PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumUnknown  PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumBilling  PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Billing"
	PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumDelivery PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress struct {
	City       *string                                                                    `json:"city,omitempty"`
	Country    *string                                                                    `json:"country,omitempty"`
	Line1      *string                                                                    `json:"line1,omitempty"`
	Line2      *string                                                                    `json:"line2,omitempty"`
	PostalCode *string                                                                    `json:"postalCode,omitempty"`
	Region     *string                                                                    `json:"region,omitempty"`
	Type       PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipTo struct {
	Address *PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumUnknown PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumDraft   PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumOpen    PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Open"
	PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumClosed  PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Closed"
	PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumVoid    PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
)

// PostPurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef
// Supplier that the purchase order is recorded against in the accounting system.
type PostPurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// PostPurchaseOrder200ApplicationJSONSourceModifiedDate
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type PostPurchaseOrder200ApplicationJSONSourceModifiedDate struct {
	Currency             *string                                                           `json:"currency,omitempty"`
	CurrencyRate         *float64                                                          `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time                                                        `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time                                                        `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                                                           `json:"id,omitempty"`
	IssueDate            *time.Time                                                        `json:"issueDate,omitempty"`
	LineItems            []PostPurchaseOrder200ApplicationJSONSourceModifiedDateLineItems  `json:"lineItems,omitempty"`
	Metadata             *PostPurchaseOrder200ApplicationJSONSourceModifiedDateMetadata    `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                        `json:"modifiedDate,omitempty"`
	Note                 *string                                                           `json:"note,omitempty"`
	PaymentDueDate       *time.Time                                                        `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                                                           `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *PostPurchaseOrder200ApplicationJSONSourceModifiedDateShipTo      `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time                                                        `json:"sourceModifiedDate,omitempty"`
	Status               *PostPurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum  `json:"status,omitempty"`
	SubTotal             *float64                                                          `json:"subTotal,omitempty"`
	SupplierRef          *PostPurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	TotalAmount          *float64                                                          `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                                                          `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                                                          `json:"totalTaxAmount,omitempty"`
}

type PostPurchaseOrder200ApplicationJSONStatusEnum string

const (
	PostPurchaseOrder200ApplicationJSONStatusEnumPending  PostPurchaseOrder200ApplicationJSONStatusEnum = "Pending"
	PostPurchaseOrder200ApplicationJSONStatusEnumFailed   PostPurchaseOrder200ApplicationJSONStatusEnum = "Failed"
	PostPurchaseOrder200ApplicationJSONStatusEnumSuccess  PostPurchaseOrder200ApplicationJSONStatusEnum = "Success"
	PostPurchaseOrder200ApplicationJSONStatusEnumTimedOut PostPurchaseOrder200ApplicationJSONStatusEnum = "TimedOut"
)

type PostPurchaseOrder200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// PostPurchaseOrder200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type PostPurchaseOrder200ApplicationJSONValidation struct {
	Errors   []PostPurchaseOrder200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []PostPurchaseOrder200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type PostPurchaseOrder200ApplicationJSON struct {
	Changes           []PostPurchaseOrder200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                 `json:"companyId"`
	CompletedOnUtc    *time.Time                                             `json:"completedOnUtc,omitempty"`
	Data              *PostPurchaseOrder200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                 `json:"dataConnectionKey"`
	DataType          *string                                                `json:"dataType,omitempty"`
	ErrorMessage      *string                                                `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                 `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                              `json:"requestedOnUtc"`
	Status            PostPurchaseOrder200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                    `json:"statusCode"`
	TimeoutInMinutes  *int                                                   `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                   `json:"timeoutInSeconds,omitempty"`
	Validation        *PostPurchaseOrder200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type PostPurchaseOrderResponse struct {
	ContentType                               string
	StatusCode                                int
	RawResponse                               *http.Response
	PostPurchaseOrder200ApplicationJSONObject *PostPurchaseOrder200ApplicationJSON
}
