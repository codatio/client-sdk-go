package operations

import (
	"time"
)

type GetDataIntegrityStatusForDataTypeDataTypeEnum string

const (
	GetDataIntegrityStatusForDataTypeDataTypeEnumBankingAccounts     GetDataIntegrityStatusForDataTypeDataTypeEnum = "banking-accounts"
	GetDataIntegrityStatusForDataTypeDataTypeEnumBankingTransactions GetDataIntegrityStatusForDataTypeDataTypeEnum = "banking-transactions"
	GetDataIntegrityStatusForDataTypeDataTypeEnumBankAccounts        GetDataIntegrityStatusForDataTypeDataTypeEnum = "bankAccounts"
	GetDataIntegrityStatusForDataTypeDataTypeEnumAccountTransactions GetDataIntegrityStatusForDataTypeDataTypeEnum = "accountTransactions"
)

type GetDataIntegrityStatusForDataTypePathParams struct {
	CompanyID string                                        `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  GetDataIntegrityStatusForDataTypeDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetDataIntegrityStatusForDataTypeRequest struct {
	PathParams GetDataIntegrityStatusForDataTypePathParams
}

// GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeAmounts
// Only returned for transactions. For accounts, there is nothing returned.
type GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeAmounts struct {
	Currency *string  `json:"currency,omitempty"`
	Max      *float64 `json:"max,omitempty"`
	Min      *float64 `json:"min,omitempty"`
}

type GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeConnectionIds struct {
	Source []string `json:"source,omitempty"`
	Target []string `json:"target,omitempty"`
}

// GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeDates
// Only returned for transactions. For accounts, there is nothing returned.
type GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeDates struct {
	MaxDate            *time.Time `json:"maxDate,omitempty"`
	MaxOverlappingDate *time.Time `json:"maxOverlappingDate,omitempty"`
	MinDate            *time.Time `json:"minDate,omitempty"`
	MinOverlappingDate *time.Time `json:"minOverlappingDate,omitempty"`
}

type GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum string

const (
	GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnumUnknown      GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum = "Unknown"
	GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnumDoesNotExist GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum = "DoesNotExist"
	GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnumError        GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum = "Error"
	GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnumComplete     GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum = "Complete"
)

type GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfo struct {
	CurrentStatus *GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum `json:"currentStatus,omitempty"`
	LastMatched   *time.Time                                                                                       `json:"lastMatched,omitempty"`
	StatusMessage *string                                                                                          `json:"statusMessage,omitempty"`
}

type GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityType struct {
	Amounts       *GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeAmounts       `json:"amounts,omitempty"`
	ConnectionIds *GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeConnectionIds `json:"connectionIds,omitempty"`
	Dates         *GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeDates         `json:"dates,omitempty"`
	StatusInfo    *GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityTypeStatusInfo    `json:"statusInfo,omitempty"`
	Type          *string                                                                            `json:"type,omitempty"`
}

type GetDataIntegrityStatusForDataType200ApplicationJSON struct {
	Metadata []GetDataIntegrityStatusForDataType200ApplicationJSONDataIntegrityType `json:"metadata,omitempty"`
}

type GetDataIntegrityStatusForDataTypeResponse struct {
	ContentType                                               string
	StatusCode                                                int
	GetDataIntegrityStatusForDataType200ApplicationJSONObject *GetDataIntegrityStatusForDataType200ApplicationJSON
}
