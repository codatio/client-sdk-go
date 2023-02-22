package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type PatchCompanyConnectionPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type PatchCompanyConnectionRequestBody struct {
	Status *string `json:"status,omitempty"`
}

type PatchCompanyConnectionRequest struct {
	PathParams PatchCompanyConnectionPathParams
	Request    *PatchCompanyConnectionRequestBody `request:"mediaType=application/json"`
}

type PatchCompanyConnection404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type PatchCompanyConnection401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type PatchCompanyConnectionResponse struct {
	Connection                                     *shared.Connection
	ContentType                                    string
	StatusCode                                     int
	PatchCompanyConnection401ApplicationJSONObject *PatchCompanyConnection401ApplicationJSON
	PatchCompanyConnection404ApplicationJSONObject *PatchCompanyConnection404ApplicationJSON
}
