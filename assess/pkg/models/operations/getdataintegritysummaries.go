package operations

import (
	"net/http"
)

type GetDataIntegritySummariesDataTypeEnum string

const (
	GetDataIntegritySummariesDataTypeEnumBankingAccounts     GetDataIntegritySummariesDataTypeEnum = "banking-accounts"
	GetDataIntegritySummariesDataTypeEnumBankingTransactions GetDataIntegritySummariesDataTypeEnum = "banking-transactions"
	GetDataIntegritySummariesDataTypeEnumBankAccounts        GetDataIntegritySummariesDataTypeEnum = "bankAccounts"
	GetDataIntegritySummariesDataTypeEnumAccountTransactions GetDataIntegritySummariesDataTypeEnum = "accountTransactions"
)

type GetDataIntegritySummariesRequest struct {
	CompanyID string                                `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  GetDataIntegritySummariesDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
	Query     *string                               `queryParam:"style=form,explode=true,name=query"`
}

type GetDataIntegritySummaries200ApplicationJSONDataIntegrityTypeByAmount struct {
	Currency        *string  `json:"currency,omitempty"`
	MatchPercentage *float64 `json:"matchPercentage,omitempty"`
	Matched         *float64 `json:"matched,omitempty"`
	Total           *float64 `json:"total,omitempty"`
	Unmatched       *float64 `json:"unmatched,omitempty"`
}

type GetDataIntegritySummaries200ApplicationJSONDataIntegrityTypeByCount struct {
	MatchPercentage *float64 `json:"matchPercentage,omitempty"`
	Matched         *float64 `json:"matched,omitempty"`
	Total           *float64 `json:"total,omitempty"`
	Unmatched       *float64 `json:"unmatched,omitempty"`
}

type GetDataIntegritySummaries200ApplicationJSONDataIntegrityType struct {
	ByAmount *GetDataIntegritySummaries200ApplicationJSONDataIntegrityTypeByAmount `json:"byAmount,omitempty"`
	ByCount  *GetDataIntegritySummaries200ApplicationJSONDataIntegrityTypeByCount  `json:"byCount,omitempty"`
	Type     *string                                                               `json:"type,omitempty"`
}

type GetDataIntegritySummaries200ApplicationJSON struct {
	Summaries []GetDataIntegritySummaries200ApplicationJSONDataIntegrityType `json:"summaries,omitempty"`
}

type GetDataIntegritySummariesResponse struct {
	ContentType                                       string
	StatusCode                                        int
	RawResponse                                       *http.Response
	GetDataIntegritySummaries200ApplicationJSONObject *GetDataIntegritySummaries200ApplicationJSON
}
