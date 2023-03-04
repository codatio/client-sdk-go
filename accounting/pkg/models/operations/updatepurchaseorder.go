package operations

import (
	"net/http"
	"time"
)

type UpdatePurchaseOrderPathParams struct {
	CompanyID       string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID    string `pathParam:"style=simple,explode=false,name=connectionId"`
	PurchaseOrderID string `pathParam:"style=simple,explode=false,name=purchaseOrderId"`
}

type UpdatePurchaseOrderQueryParams struct {
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// UpdatePurchaseOrderSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type UpdatePurchaseOrderSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdatePurchaseOrderSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type UpdatePurchaseOrderSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdatePurchaseOrderSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type UpdatePurchaseOrderSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type UpdatePurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdatePurchaseOrderSourceModifiedDateLineItems struct {
	AccountRef           *UpdatePurchaseOrderSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                              `json:"description,omitempty"`
	DiscountAmount       *float64                                                             `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                             `json:"discountPercentage,omitempty"`
	ItemRef              *UpdatePurchaseOrderSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             *float64                                                             `json:"quantity,omitempty"`
	SubTotal             *float64                                                             `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                             `json:"taxAmount,omitempty"`
	TaxRateRef           *UpdatePurchaseOrderSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                             `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []UpdatePurchaseOrderSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64                                                             `json:"unitAmount,omitempty"`
}

type UpdatePurchaseOrderSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum string

const (
	UpdatePurchaseOrderSourceModifiedDateShipToAddressTypeEnumUnknown  UpdatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	UpdatePurchaseOrderSourceModifiedDateShipToAddressTypeEnumBilling  UpdatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Billing"
	UpdatePurchaseOrderSourceModifiedDateShipToAddressTypeEnumDelivery UpdatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// UpdatePurchaseOrderSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type UpdatePurchaseOrderSourceModifiedDateShipToAddress struct {
	City       *string                                                    `json:"city,omitempty"`
	Country    *string                                                    `json:"country,omitempty"`
	Line1      *string                                                    `json:"line1,omitempty"`
	Line2      *string                                                    `json:"line2,omitempty"`
	PostalCode *string                                                    `json:"postalCode,omitempty"`
	Region     *string                                                    `json:"region,omitempty"`
	Type       UpdatePurchaseOrderSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// UpdatePurchaseOrderSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type UpdatePurchaseOrderSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// UpdatePurchaseOrderSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type UpdatePurchaseOrderSourceModifiedDateShipTo struct {
	Address *UpdatePurchaseOrderSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *UpdatePurchaseOrderSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type UpdatePurchaseOrderSourceModifiedDateStatusEnum string

const (
	UpdatePurchaseOrderSourceModifiedDateStatusEnumUnknown UpdatePurchaseOrderSourceModifiedDateStatusEnum = "Unknown"
	UpdatePurchaseOrderSourceModifiedDateStatusEnumDraft   UpdatePurchaseOrderSourceModifiedDateStatusEnum = "Draft"
	UpdatePurchaseOrderSourceModifiedDateStatusEnumOpen    UpdatePurchaseOrderSourceModifiedDateStatusEnum = "Open"
	UpdatePurchaseOrderSourceModifiedDateStatusEnumClosed  UpdatePurchaseOrderSourceModifiedDateStatusEnum = "Closed"
	UpdatePurchaseOrderSourceModifiedDateStatusEnumVoid    UpdatePurchaseOrderSourceModifiedDateStatusEnum = "Void"
)

// UpdatePurchaseOrderSourceModifiedDateSupplierRef
// Supplier that the purchase order is recorded against in the accounting system.
type UpdatePurchaseOrderSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// UpdatePurchaseOrderSourceModifiedDate
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type UpdatePurchaseOrderSourceModifiedDate struct {
	Currency             *string                                           `json:"currency,omitempty"`
	CurrencyRate         *float64                                          `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time                                        `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time                                        `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                                           `json:"id,omitempty"`
	IssueDate            *time.Time                                        `json:"issueDate,omitempty"`
	LineItems            []UpdatePurchaseOrderSourceModifiedDateLineItems  `json:"lineItems,omitempty"`
	Metadata             *UpdatePurchaseOrderSourceModifiedDateMetadata    `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                        `json:"modifiedDate,omitempty"`
	Note                 *string                                           `json:"note,omitempty"`
	PaymentDueDate       *time.Time                                        `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                                           `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *UpdatePurchaseOrderSourceModifiedDateShipTo      `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time                                        `json:"sourceModifiedDate,omitempty"`
	Status               *UpdatePurchaseOrderSourceModifiedDateStatusEnum  `json:"status,omitempty"`
	SubTotal             *float64                                          `json:"subTotal,omitempty"`
	SupplierRef          *UpdatePurchaseOrderSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	TotalAmount          *float64                                          `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                                          `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                                          `json:"totalTaxAmount,omitempty"`
}

type UpdatePurchaseOrderRequest struct {
	PathParams  UpdatePurchaseOrderPathParams
	QueryParams UpdatePurchaseOrderQueryParams
	Request     *UpdatePurchaseOrderSourceModifiedDate `request:"mediaType=application/json"`
}

type UpdatePurchaseOrder200ApplicationJSONChangesPushOperationRecordRef struct {
	DataType *string `json:"dataType,omitempty"`
	ID       *string `json:"id,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum string

const (
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumUnknown            UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Unknown"
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumCreated            UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Created"
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumModified           UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Modified"
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumDeleted            UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "Deleted"
	UpdatePurchaseOrder200ApplicationJSONChangesTypeEnumAttachmentUploaded UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum = "AttachmentUploaded"
)

type UpdatePurchaseOrder200ApplicationJSONChanges struct {
	AttachmentID *string                                                             `json:"attachmentId,omitempty"`
	RecordRef    *UpdatePurchaseOrder200ApplicationJSONChangesPushOperationRecordRef `json:"recordRef,omitempty"`
	Type         *UpdatePurchaseOrder200ApplicationJSONChangesTypeEnum               `json:"type,omitempty"`
}

// UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef
// Reference to the account to which the line item is linked.
type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef
// Reference to the product or inventory item to which the line item is linked.
type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

// UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef
// Reference to the tax rate to which the line item is linked.
type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItems struct {
	AccountRef           *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsAccountRef            `json:"accountRef,omitempty"`
	Description          *string                                                                                `json:"description,omitempty"`
	DiscountAmount       *float64                                                                               `json:"discountAmount,omitempty"`
	DiscountPercentage   *float64                                                                               `json:"discountPercentage,omitempty"`
	ItemRef              *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsItemRef               `json:"itemRef,omitempty"`
	Quantity             *float64                                                                               `json:"quantity,omitempty"`
	SubTotal             *float64                                                                               `json:"subTotal,omitempty"`
	TaxAmount            *float64                                                                               `json:"taxAmount,omitempty"`
	TaxRateRef           *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTaxRateRef            `json:"taxRateRef,omitempty"`
	TotalAmount          *float64                                                                               `json:"totalAmount,omitempty"`
	TrackingCategoryRefs []UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItemsTrackingCategoryRefs `json:"trackingCategoryRefs,omitempty"`
	UnitAmount           *float64                                                                               `json:"unitAmount,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum string

const (
	UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumUnknown  UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Unknown"
	UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumBilling  UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Billing"
	UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnumDelivery UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum = "Delivery"
)

// UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress
// Delivery address for any goods that have been ordered.
type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress struct {
	City       *string                                                                      `json:"city,omitempty"`
	Country    *string                                                                      `json:"country,omitempty"`
	Line1      *string                                                                      `json:"line1,omitempty"`
	Line2      *string                                                                      `json:"line2,omitempty"`
	PostalCode *string                                                                      `json:"postalCode,omitempty"`
	Region     *string                                                                      `json:"region,omitempty"`
	Type       UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddressTypeEnum `json:"type"`
}

// UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact
// Details of the named contact at the delivery address.
type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipTo
// Delivery details for any goods that have been ordered.
type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipTo struct {
	Address *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToAddress `json:"address,omitempty"`
	Contact *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipToContact `json:"contact,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum string

const (
	UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumUnknown UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Unknown"
	UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumDraft   UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Draft"
	UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumOpen    UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Open"
	UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumClosed  UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Closed"
	UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnumVoid    UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum = "Void"
)

// UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef
// Supplier that the purchase order is recorded against in the accounting system.
type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef struct {
	ID           string  `json:"id"`
	SupplierName *string `json:"supplierName,omitempty"`
}

// UpdatePurchaseOrder200ApplicationJSONSourceModifiedDate
// > View the coverage for purchase orders in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=purchaseOrders" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Purchase orders represent a business's intent to purchase goods or services from a supplier and normally include information such as expected delivery dates and shipping details.
//
// This information can be used to provide visibility on a business's expected payables and to track a purchase through the full procurement process.
type UpdatePurchaseOrder200ApplicationJSONSourceModifiedDate struct {
	Currency             *string                                                             `json:"currency,omitempty"`
	CurrencyRate         *float64                                                            `json:"currencyRate,omitempty"`
	DeliveryDate         *time.Time                                                          `json:"deliveryDate,omitempty"`
	ExpectedDeliveryDate *time.Time                                                          `json:"expectedDeliveryDate,omitempty"`
	ID                   *string                                                             `json:"id,omitempty"`
	IssueDate            *time.Time                                                          `json:"issueDate,omitempty"`
	LineItems            []UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateLineItems  `json:"lineItems,omitempty"`
	Metadata             *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateMetadata    `json:"metadata,omitempty"`
	ModifiedDate         *time.Time                                                          `json:"modifiedDate,omitempty"`
	Note                 *string                                                             `json:"note,omitempty"`
	PaymentDueDate       *time.Time                                                          `json:"paymentDueDate,omitempty"`
	PurchaseOrderNumber  *string                                                             `json:"purchaseOrderNumber,omitempty"`
	ShipTo               *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateShipTo      `json:"shipTo,omitempty"`
	SourceModifiedDate   *time.Time                                                          `json:"sourceModifiedDate,omitempty"`
	Status               *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateStatusEnum  `json:"status,omitempty"`
	SubTotal             *float64                                                            `json:"subTotal,omitempty"`
	SupplierRef          *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDateSupplierRef `json:"supplierRef,omitempty"`
	TotalAmount          *float64                                                            `json:"totalAmount,omitempty"`
	TotalDiscount        *float64                                                            `json:"totalDiscount,omitempty"`
	TotalTaxAmount       *float64                                                            `json:"totalTaxAmount,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSONStatusEnum string

const (
	UpdatePurchaseOrder200ApplicationJSONStatusEnumPending  UpdatePurchaseOrder200ApplicationJSONStatusEnum = "Pending"
	UpdatePurchaseOrder200ApplicationJSONStatusEnumFailed   UpdatePurchaseOrder200ApplicationJSONStatusEnum = "Failed"
	UpdatePurchaseOrder200ApplicationJSONStatusEnumSuccess  UpdatePurchaseOrder200ApplicationJSONStatusEnum = "Success"
	UpdatePurchaseOrder200ApplicationJSONStatusEnumTimedOut UpdatePurchaseOrder200ApplicationJSONStatusEnum = "TimedOut"
)

type UpdatePurchaseOrder200ApplicationJSONValidationValidationItem struct {
	ItemID        *string `json:"itemId,omitempty"`
	Message       *string `json:"message,omitempty"`
	ValidatorName *string `json:"validatorName,omitempty"`
}

// UpdatePurchaseOrder200ApplicationJSONValidation
// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
type UpdatePurchaseOrder200ApplicationJSONValidation struct {
	Errors   []UpdatePurchaseOrder200ApplicationJSONValidationValidationItem `json:"errors,omitempty"`
	Warnings []UpdatePurchaseOrder200ApplicationJSONValidationValidationItem `json:"warnings,omitempty"`
}

type UpdatePurchaseOrder200ApplicationJSON struct {
	Changes           []UpdatePurchaseOrder200ApplicationJSONChanges           `json:"changes,omitempty"`
	CompanyID         string                                                   `json:"companyId"`
	CompletedOnUtc    *time.Time                                               `json:"completedOnUtc,omitempty"`
	Data              *UpdatePurchaseOrder200ApplicationJSONSourceModifiedDate `json:"data,omitempty"`
	DataConnectionKey string                                                   `json:"dataConnectionKey"`
	DataType          *string                                                  `json:"dataType,omitempty"`
	ErrorMessage      *string                                                  `json:"errorMessage,omitempty"`
	PushOperationKey  string                                                   `json:"pushOperationKey"`
	RequestedOnUtc    time.Time                                                `json:"requestedOnUtc"`
	Status            UpdatePurchaseOrder200ApplicationJSONStatusEnum          `json:"status"`
	StatusCode        int                                                      `json:"statusCode"`
	TimeoutInMinutes  *int                                                     `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds  *int                                                     `json:"timeoutInSeconds,omitempty"`
	Validation        *UpdatePurchaseOrder200ApplicationJSONValidation         `json:"validation,omitempty"`
}

type UpdatePurchaseOrderResponse struct {
	ContentType                                 string
	StatusCode                                  int
	RawResponse                                 *http.Response
	UpdatePurchaseOrder200ApplicationJSONObject *UpdatePurchaseOrder200ApplicationJSON
}
