package operations

import (
	"time"
)

type CreateExpenseDatasetPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type CreateExpenseDatasetRequestBodyItemsLinesRecordRef struct {
	ID *string `json:"id,omitempty"`
}

type CreateExpenseDatasetRequestBodyItemsLines struct {
	AccountRef   CreateExpenseDatasetRequestBodyItemsLinesRecordRef   `json:"accountRef"`
	NetAmount    float64                                              `json:"netAmount"`
	TaxAmount    float64                                              `json:"taxAmount"`
	TaxRateRef   *CreateExpenseDatasetRequestBodyItemsLinesRecordRef  `json:"taxRateRef,omitempty"`
	TrackingRefs []CreateExpenseDatasetRequestBodyItemsLinesRecordRef `json:"trackingRefs,omitempty"`
}

type CreateExpenseDatasetRequestBodyItemsTypeEnum string

const (
	CreateExpenseDatasetRequestBodyItemsTypeEnumPayment       CreateExpenseDatasetRequestBodyItemsTypeEnum = "Payment"
	CreateExpenseDatasetRequestBodyItemsTypeEnumRefund        CreateExpenseDatasetRequestBodyItemsTypeEnum = "Refund"
	CreateExpenseDatasetRequestBodyItemsTypeEnumReward        CreateExpenseDatasetRequestBodyItemsTypeEnum = "Reward"
	CreateExpenseDatasetRequestBodyItemsTypeEnumChargeback    CreateExpenseDatasetRequestBodyItemsTypeEnum = "Chargeback"
	CreateExpenseDatasetRequestBodyItemsTypeEnumTransferIn    CreateExpenseDatasetRequestBodyItemsTypeEnum = "TransferIn"
	CreateExpenseDatasetRequestBodyItemsTypeEnumTransferOut   CreateExpenseDatasetRequestBodyItemsTypeEnum = "TransferOut"
	CreateExpenseDatasetRequestBodyItemsTypeEnumAdjustmentIn  CreateExpenseDatasetRequestBodyItemsTypeEnum = "AdjustmentIn"
	CreateExpenseDatasetRequestBodyItemsTypeEnumAdjustmentOut CreateExpenseDatasetRequestBodyItemsTypeEnum = "AdjustmentOut"
)

type CreateExpenseDatasetRequestBodyItems struct {
	Currency     string                                       `json:"currency"`
	CurrencyRate *float64                                     `json:"currencyRate,omitempty"`
	ID           string                                       `json:"id"`
	IssueDate    time.Time                                    `json:"issueDate"`
	Lines        []CreateExpenseDatasetRequestBodyItemsLines  `json:"lines,omitempty"`
	MerchantName *string                                      `json:"merchantName,omitempty"`
	Notes        *string                                      `json:"notes,omitempty"`
	Type         CreateExpenseDatasetRequestBodyItemsTypeEnum `json:"type"`
}

type CreateExpenseDatasetRequestBody struct {
	Items []CreateExpenseDatasetRequestBodyItems `json:"items,omitempty"`
}

type CreateExpenseDatasetRequest struct {
	PathParams CreateExpenseDatasetPathParams
	Request    *CreateExpenseDatasetRequestBody `request:"mediaType=application/json"`
}

type CreateExpenseDataset200ApplicationJSON struct {
	DatasetID *string `json:"datasetId,omitempty"`
}

type CreateExpenseDatasetResponse struct {
	ContentType                                  string
	StatusCode                                   int
	CreateExpenseDataset200ApplicationJSONObject *CreateExpenseDataset200ApplicationJSON
}
