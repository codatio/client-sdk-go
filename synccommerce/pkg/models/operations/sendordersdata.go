package operations

import (
	"net/http"
	"time"
)

type SendOrdersDataRequestBodyOrdersCustomerRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SendOrdersDataRequestBodyOrdersLocationRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SendOrdersDataRequestBodyOrdersOrderLineItemsDiscountAllocations struct {
	Name        *string  `json:"name,omitempty"`
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

type SendOrdersDataRequestBodyOrdersOrderLineItemsProductRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SendOrdersDataRequestBodyOrdersOrderLineItemsProductVariantRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SendOrdersDataRequestBodyOrdersOrderLineItems struct {
	DiscountAllocations []SendOrdersDataRequestBodyOrdersOrderLineItemsDiscountAllocations `json:"discountAllocations,omitempty"`
	ID                  *string                                                            `json:"id,omitempty"`
	ProductRef          *SendOrdersDataRequestBodyOrdersOrderLineItemsProductRef           `json:"productRef,omitempty"`
	ProductVariantRef   *SendOrdersDataRequestBodyOrdersOrderLineItemsProductVariantRef    `json:"productVariantRef,omitempty"`
	Quantity            *float64                                                           `json:"quantity,omitempty"`
	TaxPercentage       *float64                                                           `json:"taxPercentage,omitempty"`
	TotalAmount         *float64                                                           `json:"totalAmount,omitempty"`
	TotalTaxAmount      *float64                                                           `json:"totalTaxAmount,omitempty"`
	UnitPrice           *float64                                                           `json:"unitPrice,omitempty"`
}

type SendOrdersDataRequestBodyOrdersPaymentsStatusEnum string

const (
	SendOrdersDataRequestBodyOrdersPaymentsStatusEnumUnknown    SendOrdersDataRequestBodyOrdersPaymentsStatusEnum = "Unknown"
	SendOrdersDataRequestBodyOrdersPaymentsStatusEnumPending    SendOrdersDataRequestBodyOrdersPaymentsStatusEnum = "Pending"
	SendOrdersDataRequestBodyOrdersPaymentsStatusEnumAuthorized SendOrdersDataRequestBodyOrdersPaymentsStatusEnum = "Authorized"
	SendOrdersDataRequestBodyOrdersPaymentsStatusEnumPaid       SendOrdersDataRequestBodyOrdersPaymentsStatusEnum = "Paid"
	SendOrdersDataRequestBodyOrdersPaymentsStatusEnumFailed     SendOrdersDataRequestBodyOrdersPaymentsStatusEnum = "Failed"
	SendOrdersDataRequestBodyOrdersPaymentsStatusEnumCancelled  SendOrdersDataRequestBodyOrdersPaymentsStatusEnum = "Cancelled"
)

type SendOrdersDataRequestBodyOrdersPaymentsTypeEnum string

const (
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumUnknown     SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Unknown"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumCash        SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Cash"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumCard        SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Card"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumInvoice     SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Invoice"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumOnlineCard  SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "OnlineCard"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumSwish       SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Swish"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumVipps       SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Vipps"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumMobile      SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Mobile"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumStoreCredit SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "StoreCredit"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumPaypal      SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Paypal"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumCustom      SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Custom"
	SendOrdersDataRequestBodyOrdersPaymentsTypeEnumPrepaid     SendOrdersDataRequestBodyOrdersPaymentsTypeEnum = "Prepaid"
)

type SendOrdersDataRequestBodyOrdersPayments struct {
	Amount             *float64                                           `json:"amount,omitempty"`
	CreatedDate        *time.Time                                         `json:"createdDate,omitempty"`
	Currency           *string                                            `json:"currency,omitempty"`
	DueDate            *time.Time                                         `json:"dueDate,omitempty"`
	ID                 *string                                            `json:"id,omitempty"`
	ModifiedDate       *time.Time                                         `json:"modifiedDate,omitempty"`
	PaymentProvider    *string                                            `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time                                         `json:"sourceModifiedDate,omitempty"`
	Status             *SendOrdersDataRequestBodyOrdersPaymentsStatusEnum `json:"status,omitempty"`
	Type               *SendOrdersDataRequestBodyOrdersPaymentsTypeEnum   `json:"type,omitempty"`
}

type SendOrdersDataRequestBodyOrdersServiceChargesTypeEnum string

const (
	SendOrdersDataRequestBodyOrdersServiceChargesTypeEnumUnknown     SendOrdersDataRequestBodyOrdersServiceChargesTypeEnum = "Unknown"
	SendOrdersDataRequestBodyOrdersServiceChargesTypeEnumGeneric     SendOrdersDataRequestBodyOrdersServiceChargesTypeEnum = "Generic"
	SendOrdersDataRequestBodyOrdersServiceChargesTypeEnumShipping    SendOrdersDataRequestBodyOrdersServiceChargesTypeEnum = "Shipping"
	SendOrdersDataRequestBodyOrdersServiceChargesTypeEnumOverpayment SendOrdersDataRequestBodyOrdersServiceChargesTypeEnum = "Overpayment"
)

type SendOrdersDataRequestBodyOrdersServiceCharges struct {
	Description   *string                                                `json:"description,omitempty"`
	Quantity      *int                                                   `json:"quantity,omitempty"`
	TaxAmount     *float64                                               `json:"taxAmount,omitempty"`
	TaxPercentage *float64                                               `json:"taxPercentage,omitempty"`
	TotalAmount   *float64                                               `json:"totalAmount,omitempty"`
	Type          *SendOrdersDataRequestBodyOrdersServiceChargesTypeEnum `json:"type,omitempty"`
}

type SendOrdersDataRequestBodyOrders struct {
	ClosedDate         *time.Time                                      `json:"closedDate,omitempty"`
	Country            *string                                         `json:"country,omitempty"`
	CreatedDate        *time.Time                                      `json:"createdDate,omitempty"`
	Currency           *string                                         `json:"currency,omitempty"`
	CustomerRef        *SendOrdersDataRequestBodyOrdersCustomerRef     `json:"customerRef,omitempty"`
	ID                 *string                                         `json:"id,omitempty"`
	LocationRef        *SendOrdersDataRequestBodyOrdersLocationRef     `json:"locationRef,omitempty"`
	ModifiedDate       *time.Time                                      `json:"modifiedDate,omitempty"`
	OrderLineItems     []SendOrdersDataRequestBodyOrdersOrderLineItems `json:"orderLineItems,omitempty"`
	OrderNumber        *string                                         `json:"orderNumber,omitempty"`
	Payments           []SendOrdersDataRequestBodyOrdersPayments       `json:"payments,omitempty"`
	ServiceCharges     []SendOrdersDataRequestBodyOrdersServiceCharges `json:"serviceCharges,omitempty"`
	SourceModifiedDate *time.Time                                      `json:"sourceModifiedDate,omitempty"`
	TotalAmount        *float64                                        `json:"totalAmount,omitempty"`
	TotalDiscount      *float64                                        `json:"totalDiscount,omitempty"`
	TotalGratuity      *float64                                        `json:"totalGratuity,omitempty"`
	TotalRefund        *float64                                        `json:"totalRefund,omitempty"`
	TotalTaxAmount     *float64                                        `json:"totalTaxAmount,omitempty"`
}

type SendOrdersDataRequestBody struct {
	ContractVersion *string                           `json:"contractVersion,omitempty"`
	Orders          []SendOrdersDataRequestBodyOrders `json:"orders,omitempty"`
}

type SendOrdersDataRequest struct {
	RequestBody *SendOrdersDataRequestBody `request:"mediaType=application/json"`
	CompanyID   string                     `pathParam:"style=simple,explode=false,name=companyId"`
}

type SendOrdersData200ApplicationJSONDataOrdersCustomerRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SendOrdersData200ApplicationJSONDataOrdersLocationRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SendOrdersData200ApplicationJSONDataOrdersOrderLineItemsDiscountAllocations struct {
	Name        *string  `json:"name,omitempty"`
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

type SendOrdersData200ApplicationJSONDataOrdersOrderLineItemsProductRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SendOrdersData200ApplicationJSONDataOrdersOrderLineItemsProductVariantRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SendOrdersData200ApplicationJSONDataOrdersOrderLineItems struct {
	DiscountAllocations []SendOrdersData200ApplicationJSONDataOrdersOrderLineItemsDiscountAllocations `json:"discountAllocations,omitempty"`
	ID                  *string                                                                       `json:"id,omitempty"`
	ProductRef          *SendOrdersData200ApplicationJSONDataOrdersOrderLineItemsProductRef           `json:"productRef,omitempty"`
	ProductVariantRef   *SendOrdersData200ApplicationJSONDataOrdersOrderLineItemsProductVariantRef    `json:"productVariantRef,omitempty"`
	Quantity            *float64                                                                      `json:"quantity,omitempty"`
	TaxPercentage       *float64                                                                      `json:"taxPercentage,omitempty"`
	TotalAmount         *float64                                                                      `json:"totalAmount,omitempty"`
	TotalTaxAmount      *float64                                                                      `json:"totalTaxAmount,omitempty"`
	UnitPrice           *float64                                                                      `json:"unitPrice,omitempty"`
}

type SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnum string

const (
	SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnumUnknown    SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnum = "Unknown"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnumPending    SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnum = "Pending"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnumAuthorized SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnum = "Authorized"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnumPaid       SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnum = "Paid"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnumFailed     SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnum = "Failed"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnumCancelled  SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnum = "Cancelled"
)

type SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum string

const (
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumUnknown     SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Unknown"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumCash        SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Cash"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumCard        SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Card"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumInvoice     SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Invoice"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumOnlineCard  SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "OnlineCard"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumSwish       SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Swish"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumVipps       SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Vipps"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumMobile      SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Mobile"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumStoreCredit SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "StoreCredit"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumPaypal      SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Paypal"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumCustom      SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Custom"
	SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnumPrepaid     SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum = "Prepaid"
)

type SendOrdersData200ApplicationJSONDataOrdersPayments struct {
	Amount             *float64                                                      `json:"amount,omitempty"`
	CreatedDate        *time.Time                                                    `json:"createdDate,omitempty"`
	Currency           *string                                                       `json:"currency,omitempty"`
	DueDate            *time.Time                                                    `json:"dueDate,omitempty"`
	ID                 *string                                                       `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                    `json:"modifiedDate,omitempty"`
	PaymentProvider    *string                                                       `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time                                                    `json:"sourceModifiedDate,omitempty"`
	Status             *SendOrdersData200ApplicationJSONDataOrdersPaymentsStatusEnum `json:"status,omitempty"`
	Type               *SendOrdersData200ApplicationJSONDataOrdersPaymentsTypeEnum   `json:"type,omitempty"`
}

type SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnum string

const (
	SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnumUnknown     SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnum = "Unknown"
	SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnumGeneric     SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnum = "Generic"
	SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnumShipping    SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnum = "Shipping"
	SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnumOverpayment SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnum = "Overpayment"
)

type SendOrdersData200ApplicationJSONDataOrdersServiceCharges struct {
	Description   *string                                                           `json:"description,omitempty"`
	Quantity      *int                                                              `json:"quantity,omitempty"`
	TaxAmount     *float64                                                          `json:"taxAmount,omitempty"`
	TaxPercentage *float64                                                          `json:"taxPercentage,omitempty"`
	TotalAmount   *float64                                                          `json:"totalAmount,omitempty"`
	Type          *SendOrdersData200ApplicationJSONDataOrdersServiceChargesTypeEnum `json:"type,omitempty"`
}

type SendOrdersData200ApplicationJSONDataOrders struct {
	ClosedDate         *time.Time                                                 `json:"closedDate,omitempty"`
	Country            *string                                                    `json:"country,omitempty"`
	CreatedDate        *time.Time                                                 `json:"createdDate,omitempty"`
	Currency           *string                                                    `json:"currency,omitempty"`
	CustomerRef        *SendOrdersData200ApplicationJSONDataOrdersCustomerRef     `json:"customerRef,omitempty"`
	ID                 *string                                                    `json:"id,omitempty"`
	LocationRef        *SendOrdersData200ApplicationJSONDataOrdersLocationRef     `json:"locationRef,omitempty"`
	ModifiedDate       *time.Time                                                 `json:"modifiedDate,omitempty"`
	OrderLineItems     []SendOrdersData200ApplicationJSONDataOrdersOrderLineItems `json:"orderLineItems,omitempty"`
	OrderNumber        *string                                                    `json:"orderNumber,omitempty"`
	Payments           []SendOrdersData200ApplicationJSONDataOrdersPayments       `json:"payments,omitempty"`
	ServiceCharges     []SendOrdersData200ApplicationJSONDataOrdersServiceCharges `json:"serviceCharges,omitempty"`
	SourceModifiedDate *time.Time                                                 `json:"sourceModifiedDate,omitempty"`
	TotalAmount        *float64                                                   `json:"totalAmount,omitempty"`
	TotalDiscount      *float64                                                   `json:"totalDiscount,omitempty"`
	TotalGratuity      *float64                                                   `json:"totalGratuity,omitempty"`
	TotalRefund        *float64                                                   `json:"totalRefund,omitempty"`
	TotalTaxAmount     *float64                                                   `json:"totalTaxAmount,omitempty"`
}

type SendOrdersData200ApplicationJSONData struct {
	ContractVersion *string                                      `json:"contractVersion,omitempty"`
	Orders          []SendOrdersData200ApplicationJSONDataOrders `json:"orders,omitempty"`
}

type SendOrdersData200ApplicationJSON struct {
	Data      *SendOrdersData200ApplicationJSONData `json:"data,omitempty"`
	DatasetID *string                               `json:"datasetId,omitempty"`
}

type SendOrdersDataResponse struct {
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
	SendOrdersData200ApplicationJSONObject *SendOrdersData200ApplicationJSON
}
