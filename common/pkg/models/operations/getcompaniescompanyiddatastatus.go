package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetCompaniesCompanyIDDataStatusPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDDataStatusRequest struct {
	PathParams GetCompaniesCompanyIDDataStatusPathParams
}

type GetCompaniesCompanyIDDataStatus404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompaniesCompanyIDDataStatus401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompaniesCompanyIDDataStatus200ApplicationJSON struct {
	DataType1 *shared.DataStatus `json:"dataType1,omitempty"`
	DataType2 *shared.DataStatus `json:"dataType2,omitempty"`
}

type GetCompaniesCompanyIDDataStatusResponse struct {
	ContentType                                             string
	StatusCode                                              int
	GetCompaniesCompanyIDDataStatus200ApplicationJSONObject *GetCompaniesCompanyIDDataStatus200ApplicationJSON
	GetCompaniesCompanyIDDataStatus401ApplicationJSONObject *GetCompaniesCompanyIDDataStatus401ApplicationJSON
	GetCompaniesCompanyIDDataStatus404ApplicationJSONObject *GetCompaniesCompanyIDDataStatus404ApplicationJSON
}
