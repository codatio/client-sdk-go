package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetCompanyConnectionPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCompanyConnectionRequest struct {
	PathParams GetCompanyConnectionPathParams
	Request    *bool `request:"mediaType=application/json"`
}

type GetCompanyConnection404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyConnection401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyConnectionResponse struct {
	Connection                                   *shared.Connection
	ContentType                                  string
	StatusCode                                   int
	GetCompanyConnection401ApplicationJSONObject *GetCompanyConnection401ApplicationJSON
	GetCompanyConnection404ApplicationJSONObject *GetCompanyConnection404ApplicationJSON
}
