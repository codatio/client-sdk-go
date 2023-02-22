package operations

import (
	"github.com/codatio/client-sdk-go/assess/pkg/models/shared"
)

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesPathParams struct {
	CompanyID string                           `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  shared.DataIntegritydataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesQueryParams struct {
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesRequest struct {
	PathParams  GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesPathParams
	QueryParams GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesQueryParams
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSON struct {
	Summaries []shared.DataIntegritySummary `json:"summaries,omitempty"`
}

type GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummariesResponse struct {
	ContentType                                                                                    string
	StatusCode                                                                                     int
	GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSONObject *GetDataCompaniesCompanyIDAssessDataTypesDataTypeDataIntegritySummaries200ApplicationJSON
}
