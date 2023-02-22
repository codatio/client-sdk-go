package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetPullOperationPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	DatasetID string `pathParam:"style=simple,explode=false,name=datasetId"`
}

type GetPullOperationRequest struct {
	PathParams GetPullOperationPathParams
}

type GetPullOperation404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetPullOperation401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetPullOperationResponse struct {
	ContentType                              string
	PullOperation                            *shared.PullOperation
	StatusCode                               int
	GetPullOperation401ApplicationJSONObject *GetPullOperation401ApplicationJSON
	GetPullOperation404ApplicationJSONObject *GetPullOperation404ApplicationJSON
}
