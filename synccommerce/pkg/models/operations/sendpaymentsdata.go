package operations

import (
	"net/http"
	"time"
)

type SendPaymentsDataRequestBodyPaymentsStatusEnum string

const (
	SendPaymentsDataRequestBodyPaymentsStatusEnumUnknown    SendPaymentsDataRequestBodyPaymentsStatusEnum = "Unknown"
	SendPaymentsDataRequestBodyPaymentsStatusEnumPending    SendPaymentsDataRequestBodyPaymentsStatusEnum = "Pending"
	SendPaymentsDataRequestBodyPaymentsStatusEnumAuthorized SendPaymentsDataRequestBodyPaymentsStatusEnum = "Authorized"
	SendPaymentsDataRequestBodyPaymentsStatusEnumPaid       SendPaymentsDataRequestBodyPaymentsStatusEnum = "Paid"
	SendPaymentsDataRequestBodyPaymentsStatusEnumFailed     SendPaymentsDataRequestBodyPaymentsStatusEnum = "Failed"
	SendPaymentsDataRequestBodyPaymentsStatusEnumCancelled  SendPaymentsDataRequestBodyPaymentsStatusEnum = "Cancelled"
)

type SendPaymentsDataRequestBodyPaymentsTypeEnum string

const (
	SendPaymentsDataRequestBodyPaymentsTypeEnumUnknown     SendPaymentsDataRequestBodyPaymentsTypeEnum = "Unknown"
	SendPaymentsDataRequestBodyPaymentsTypeEnumCash        SendPaymentsDataRequestBodyPaymentsTypeEnum = "Cash"
	SendPaymentsDataRequestBodyPaymentsTypeEnumCard        SendPaymentsDataRequestBodyPaymentsTypeEnum = "Card"
	SendPaymentsDataRequestBodyPaymentsTypeEnumInvoice     SendPaymentsDataRequestBodyPaymentsTypeEnum = "Invoice"
	SendPaymentsDataRequestBodyPaymentsTypeEnumOnlineCard  SendPaymentsDataRequestBodyPaymentsTypeEnum = "OnlineCard"
	SendPaymentsDataRequestBodyPaymentsTypeEnumSwish       SendPaymentsDataRequestBodyPaymentsTypeEnum = "Swish"
	SendPaymentsDataRequestBodyPaymentsTypeEnumVipps       SendPaymentsDataRequestBodyPaymentsTypeEnum = "Vipps"
	SendPaymentsDataRequestBodyPaymentsTypeEnumMobile      SendPaymentsDataRequestBodyPaymentsTypeEnum = "Mobile"
	SendPaymentsDataRequestBodyPaymentsTypeEnumStoreCredit SendPaymentsDataRequestBodyPaymentsTypeEnum = "StoreCredit"
	SendPaymentsDataRequestBodyPaymentsTypeEnumPaypal      SendPaymentsDataRequestBodyPaymentsTypeEnum = "Paypal"
	SendPaymentsDataRequestBodyPaymentsTypeEnumCustom      SendPaymentsDataRequestBodyPaymentsTypeEnum = "Custom"
	SendPaymentsDataRequestBodyPaymentsTypeEnumPrepaid     SendPaymentsDataRequestBodyPaymentsTypeEnum = "Prepaid"
)

type SendPaymentsDataRequestBodyPayments struct {
	Amount             *float64                                       `json:"amount,omitempty"`
	CreatedDate        *time.Time                                     `json:"createdDate,omitempty"`
	Currency           *string                                        `json:"currency,omitempty"`
	DueDate            *time.Time                                     `json:"dueDate,omitempty"`
	ID                 *string                                        `json:"id,omitempty"`
	ModifiedDate       *time.Time                                     `json:"modifiedDate,omitempty"`
	PaymentProvider    *string                                        `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time                                     `json:"sourceModifiedDate,omitempty"`
	Status             *SendPaymentsDataRequestBodyPaymentsStatusEnum `json:"status,omitempty"`
	Type               *SendPaymentsDataRequestBodyPaymentsTypeEnum   `json:"type,omitempty"`
}

type SendPaymentsDataRequestBody struct {
	ContractVersion *string                               `json:"contractVersion,omitempty"`
	Payments        []SendPaymentsDataRequestBodyPayments `json:"payments,omitempty"`
}

type SendPaymentsDataRequest struct {
	RequestBody *SendPaymentsDataRequestBody `request:"mediaType=application/json"`
	CompanyID   string                       `pathParam:"style=simple,explode=false,name=companyId"`
}

type SendPaymentsData200ApplicationJSONDataPaymentsStatusEnum string

const (
	SendPaymentsData200ApplicationJSONDataPaymentsStatusEnumUnknown    SendPaymentsData200ApplicationJSONDataPaymentsStatusEnum = "Unknown"
	SendPaymentsData200ApplicationJSONDataPaymentsStatusEnumPending    SendPaymentsData200ApplicationJSONDataPaymentsStatusEnum = "Pending"
	SendPaymentsData200ApplicationJSONDataPaymentsStatusEnumAuthorized SendPaymentsData200ApplicationJSONDataPaymentsStatusEnum = "Authorized"
	SendPaymentsData200ApplicationJSONDataPaymentsStatusEnumPaid       SendPaymentsData200ApplicationJSONDataPaymentsStatusEnum = "Paid"
	SendPaymentsData200ApplicationJSONDataPaymentsStatusEnumFailed     SendPaymentsData200ApplicationJSONDataPaymentsStatusEnum = "Failed"
	SendPaymentsData200ApplicationJSONDataPaymentsStatusEnumCancelled  SendPaymentsData200ApplicationJSONDataPaymentsStatusEnum = "Cancelled"
)

type SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum string

const (
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumUnknown     SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Unknown"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumCash        SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Cash"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumCard        SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Card"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumInvoice     SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Invoice"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumOnlineCard  SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "OnlineCard"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumSwish       SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Swish"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumVipps       SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Vipps"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumMobile      SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Mobile"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumStoreCredit SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "StoreCredit"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumPaypal      SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Paypal"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumCustom      SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Custom"
	SendPaymentsData200ApplicationJSONDataPaymentsTypeEnumPrepaid     SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum = "Prepaid"
)

type SendPaymentsData200ApplicationJSONDataPayments struct {
	Amount             *float64                                                  `json:"amount,omitempty"`
	CreatedDate        *time.Time                                                `json:"createdDate,omitempty"`
	Currency           *string                                                   `json:"currency,omitempty"`
	DueDate            *time.Time                                                `json:"dueDate,omitempty"`
	ID                 *string                                                   `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                `json:"modifiedDate,omitempty"`
	PaymentProvider    *string                                                   `json:"paymentProvider,omitempty"`
	SourceModifiedDate *time.Time                                                `json:"sourceModifiedDate,omitempty"`
	Status             *SendPaymentsData200ApplicationJSONDataPaymentsStatusEnum `json:"status,omitempty"`
	Type               *SendPaymentsData200ApplicationJSONDataPaymentsTypeEnum   `json:"type,omitempty"`
}

type SendPaymentsData200ApplicationJSONData struct {
	ContractVersion *string                                          `json:"contractVersion,omitempty"`
	Payments        []SendPaymentsData200ApplicationJSONDataPayments `json:"payments,omitempty"`
}

type SendPaymentsData200ApplicationJSON struct {
	Data      *SendPaymentsData200ApplicationJSONData `json:"data,omitempty"`
	DatasetID *string                                 `json:"datasetId,omitempty"`
}

type SendPaymentsDataResponse struct {
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
	SendPaymentsData200ApplicationJSONObject *SendPaymentsData200ApplicationJSON
}
