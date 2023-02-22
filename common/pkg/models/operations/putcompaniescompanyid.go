package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type PutCompaniesCompanyIDPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PutCompaniesCompanyIDRequestBody struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type PutCompaniesCompanyIDRequest struct {
	PathParams PutCompaniesCompanyIDPathParams
	Request    *PutCompaniesCompanyIDRequestBody `request:"mediaType=application/json"`
}

type PutCompaniesCompanyID401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type PutCompaniesCompanyIDResponse struct {
	Company                                       *shared.Company
	ContentType                                   string
	StatusCode                                    int
	PutCompaniesCompanyID401ApplicationJSONObject *PutCompaniesCompanyID401ApplicationJSON
}
