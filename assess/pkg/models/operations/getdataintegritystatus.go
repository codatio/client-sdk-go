package operations

import (
	"net/http"
	"time"
)

type GetDataIntegrityStatusDataTypeEnum string

const (
	GetDataIntegrityStatusDataTypeEnumBankingAccounts     GetDataIntegrityStatusDataTypeEnum = "banking-accounts"
	GetDataIntegrityStatusDataTypeEnumBankingTransactions GetDataIntegrityStatusDataTypeEnum = "banking-transactions"
	GetDataIntegrityStatusDataTypeEnumBankAccounts        GetDataIntegrityStatusDataTypeEnum = "bankAccounts"
	GetDataIntegrityStatusDataTypeEnumAccountTransactions GetDataIntegrityStatusDataTypeEnum = "accountTransactions"
)

type GetDataIntegrityStatusPathParams struct {
	CompanyID string                             `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  GetDataIntegrityStatusDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetDataIntegrityStatusRequest struct {
	PathParams GetDataIntegrityStatusPathParams
}

// GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeAmounts
// Only returned for transactions. For accounts, there is nothing returned.
type GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeAmounts struct {
	Currency *string  `json:"currency,omitempty"`
	Max      *float64 `json:"max,omitempty"`
	Min      *float64 `json:"min,omitempty"`
}

type GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeConnectionIds struct {
	Source []string `json:"source,omitempty"`
	Target []string `json:"target,omitempty"`
}

// GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeDates
// Only returned for transactions. For accounts, there is nothing returned.
type GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeDates struct {
	MaxDate            *time.Time `json:"maxDate,omitempty"`
	MaxOverlappingDate *time.Time `json:"maxOverlappingDate,omitempty"`
	MinDate            *time.Time `json:"minDate,omitempty"`
	MinOverlappingDate *time.Time `json:"minOverlappingDate,omitempty"`
}

type GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum string

const (
	GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnumUnknown      GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum = "Unknown"
	GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnumDoesNotExist GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum = "DoesNotExist"
	GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnumError        GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum = "Error"
	GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnumComplete     GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum = "Complete"
)

type GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfo struct {
	CurrentStatus *GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfoCurrentStatusEnum `json:"currentStatus,omitempty"`
	LastMatched   *time.Time                                                                            `json:"lastMatched,omitempty"`
	StatusMessage *string                                                                               `json:"statusMessage,omitempty"`
}

type GetDataIntegrityStatus200ApplicationJSONDataIntegrityType struct {
	Amounts       *GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeAmounts       `json:"amounts,omitempty"`
	ConnectionIds *GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeConnectionIds `json:"connectionIds,omitempty"`
	Dates         *GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeDates         `json:"dates,omitempty"`
	StatusInfo    *GetDataIntegrityStatus200ApplicationJSONDataIntegrityTypeStatusInfo    `json:"statusInfo,omitempty"`
	Type          *string                                                                 `json:"type,omitempty"`
}

type GetDataIntegrityStatus200ApplicationJSON struct {
	Metadata []GetDataIntegrityStatus200ApplicationJSONDataIntegrityType `json:"metadata,omitempty"`
}

type GetDataIntegrityStatusResponse struct {
	ContentType                                    string
	StatusCode                                     int
	RawResponse                                    *http.Response
	GetDataIntegrityStatus200ApplicationJSONObject *GetDataIntegrityStatus200ApplicationJSON
}
