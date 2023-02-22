package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetCompaniesCompanyIDPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDRequest struct {
	PathParams GetCompaniesCompanyIDPathParams
}

type GetCompaniesCompanyID401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompaniesCompanyIDResponse struct {
	Company                                       *shared.Company
	ContentType                                   string
	StatusCode                                    int
	GetCompaniesCompanyID401ApplicationJSONObject *GetCompaniesCompanyID401ApplicationJSON
}
