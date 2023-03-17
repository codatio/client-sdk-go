package operations

import (
	"net/http"
	"time"
)

type SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnum string

const (
	SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnumUnknown       SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnum = "Unknown"
	SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnumFee           SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnum = "Fee"
	SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnumOrder         SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnum = "Order"
	SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnumPayment       SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnum = "Payment"
	SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnumServiceCharge SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnum = "ServiceCharge"
)

type SendTransactionsDataRequestBodyTransactionsTransactionSourceRef struct {
	ID   *string                                                                  `json:"id,omitempty"`
	Type *SendTransactionsDataRequestBodyTransactionsTransactionSourceRefTypeEnum `json:"type,omitempty"`
}

type SendTransactionsDataRequestBodyTransactionsTypeEnum string

const (
	SendTransactionsDataRequestBodyTransactionsTypeEnumUnknown          SendTransactionsDataRequestBodyTransactionsTypeEnum = "Unknown"
	SendTransactionsDataRequestBodyTransactionsTypeEnumPayment          SendTransactionsDataRequestBodyTransactionsTypeEnum = "Payment"
	SendTransactionsDataRequestBodyTransactionsTypeEnumRefund           SendTransactionsDataRequestBodyTransactionsTypeEnum = "Refund"
	SendTransactionsDataRequestBodyTransactionsTypeEnumPayout           SendTransactionsDataRequestBodyTransactionsTypeEnum = "Payout"
	SendTransactionsDataRequestBodyTransactionsTypeEnumFailedPayout     SendTransactionsDataRequestBodyTransactionsTypeEnum = "FailedPayout"
	SendTransactionsDataRequestBodyTransactionsTypeEnumTransfer         SendTransactionsDataRequestBodyTransactionsTypeEnum = "Transfer"
	SendTransactionsDataRequestBodyTransactionsTypeEnumPaymentFee       SendTransactionsDataRequestBodyTransactionsTypeEnum = "PaymentFee"
	SendTransactionsDataRequestBodyTransactionsTypeEnumPaymentFeeRefund SendTransactionsDataRequestBodyTransactionsTypeEnum = "PaymentFeeRefund"
)

type SendTransactionsDataRequestBodyTransactions struct {
	CreatedDate          *time.Time                                                       `json:"createdDate,omitempty"`
	Currency             *string                                                          `json:"currency,omitempty"`
	ID                   *string                                                          `json:"id,omitempty"`
	ModifiedDate         *time.Time                                                       `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                                       `json:"sourceModifiedDate,omitempty"`
	SubType              *string                                                          `json:"subType,omitempty"`
	TotalAmount          *float64                                                         `json:"totalAmount,omitempty"`
	TransactionSourceRef *SendTransactionsDataRequestBodyTransactionsTransactionSourceRef `json:"transactionSourceRef,omitempty"`
	Type                 *SendTransactionsDataRequestBodyTransactionsTypeEnum             `json:"type,omitempty"`
}

type SendTransactionsDataRequestBody struct {
	ContractVersion *string                                       `json:"contractVersion,omitempty"`
	Transactions    []SendTransactionsDataRequestBodyTransactions `json:"transactions,omitempty"`
}

type SendTransactionsDataRequest struct {
	RequestBody *SendTransactionsDataRequestBody `request:"mediaType=application/json"`
	CompanyID   string                           `pathParam:"style=simple,explode=false,name=companyId"`
}

type SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnum string

const (
	SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnumUnknown       SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnum = "Unknown"
	SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnumFee           SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnum = "Fee"
	SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnumOrder         SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnum = "Order"
	SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnumPayment       SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnum = "Payment"
	SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnumServiceCharge SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnum = "ServiceCharge"
)

type SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRef struct {
	ID   *string                                                                             `json:"id,omitempty"`
	Type *SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRefTypeEnum `json:"type,omitempty"`
}

type SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum string

const (
	SendTransactionsData200ApplicationJSONDataTransactionsTypeEnumUnknown          SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum = "Unknown"
	SendTransactionsData200ApplicationJSONDataTransactionsTypeEnumPayment          SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum = "Payment"
	SendTransactionsData200ApplicationJSONDataTransactionsTypeEnumRefund           SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum = "Refund"
	SendTransactionsData200ApplicationJSONDataTransactionsTypeEnumPayout           SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum = "Payout"
	SendTransactionsData200ApplicationJSONDataTransactionsTypeEnumFailedPayout     SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum = "FailedPayout"
	SendTransactionsData200ApplicationJSONDataTransactionsTypeEnumTransfer         SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum = "Transfer"
	SendTransactionsData200ApplicationJSONDataTransactionsTypeEnumPaymentFee       SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum = "PaymentFee"
	SendTransactionsData200ApplicationJSONDataTransactionsTypeEnumPaymentFeeRefund SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum = "PaymentFeeRefund"
)

type SendTransactionsData200ApplicationJSONDataTransactions struct {
	CreatedDate          *time.Time                                                                  `json:"createdDate,omitempty"`
	Currency             *string                                                                     `json:"currency,omitempty"`
	ID                   *string                                                                     `json:"id,omitempty"`
	ModifiedDate         *time.Time                                                                  `json:"modifiedDate,omitempty"`
	SourceModifiedDate   *time.Time                                                                  `json:"sourceModifiedDate,omitempty"`
	SubType              *string                                                                     `json:"subType,omitempty"`
	TotalAmount          *float64                                                                    `json:"totalAmount,omitempty"`
	TransactionSourceRef *SendTransactionsData200ApplicationJSONDataTransactionsTransactionSourceRef `json:"transactionSourceRef,omitempty"`
	Type                 *SendTransactionsData200ApplicationJSONDataTransactionsTypeEnum             `json:"type,omitempty"`
}

type SendTransactionsData200ApplicationJSONData struct {
	ContractVersion *string                                                  `json:"contractVersion,omitempty"`
	Transactions    []SendTransactionsData200ApplicationJSONDataTransactions `json:"transactions,omitempty"`
}

type SendTransactionsData200ApplicationJSON struct {
	Data      *SendTransactionsData200ApplicationJSONData `json:"data,omitempty"`
	DatasetID *string                                     `json:"datasetId,omitempty"`
}

type SendTransactionsDataResponse struct {
	ContentType                                  string
	StatusCode                                   int
	RawResponse                                  *http.Response
	SendTransactionsData200ApplicationJSONObject *SendTransactionsData200ApplicationJSON
}
