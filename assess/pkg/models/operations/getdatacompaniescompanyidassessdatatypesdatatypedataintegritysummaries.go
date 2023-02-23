package operations

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnum string

const (
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnumBankingAccounts     GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnum = "banking-accounts"
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnumBankingTransactions GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnum = "banking-transactions"
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnumBankAccounts        GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnum = "bankAccounts"
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnumAccountTransactions GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnum = "accountTransactions"
)

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesPathParams struct {
	CompanyID string                                                                             `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesDataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesQueryParams struct {
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesRequest struct {
	PathParams  GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesPathParams
	QueryParams GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesQueryParams
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONDataIntegrityTypeByAmount struct {
	Currency        *string  `json:"currency,omitempty"`
	MatchPercentage *float64 `json:"matchPercentage,omitempty"`
	Matched         *float64 `json:"matched,omitempty"`
	Total           *float64 `json:"total,omitempty"`
	Unmatched       *float64 `json:"unmatched,omitempty"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONDataIntegrityTypeByCount struct {
	MatchPercentage *float64 `json:"matchPercentage,omitempty"`
	Matched         *float64 `json:"matched,omitempty"`
	Total           *float64 `json:"total,omitempty"`
	Unmatched       *float64 `json:"unmatched,omitempty"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONDataIntegrityType struct {
	ByAmount *GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONDataIntegrityTypeByAmount `json:"byAmount,omitempty"`
	ByCount  *GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONDataIntegrityTypeByCount  `json:"byCount,omitempty"`
	Type     *string                                                                                                            `json:"type,omitempty"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSON struct {
	Summaries []GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONDataIntegrityType `json:"summaries,omitempty"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesResponse struct {
	ContentType                                                                                    string
	StatusCode                                                                                     int
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONObject *GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSON
}
