package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type CreatePullOperationPathParams struct {
	CompanyID string              `pathParam:"style=simple,explode=false,name=companyId"`
	DataType  shared.DataTypeEnum `pathParam:"style=simple,explode=false,name=dataType"`
}

type CreatePullOperationQueryParams struct {
	ConnectionID *string `queryParam:"style=form,explode=true,name=connectionId"`
}

type CreatePullOperationRequest struct {
	PathParams  CreatePullOperationPathParams
	QueryParams CreatePullOperationQueryParams
}

type CreatePullOperation404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type CreatePullOperation401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type CreatePullOperationResponse struct {
	ContentType                                 string
	PullOperation                               *shared.PullOperation
	StatusCode                                  int
	CreatePullOperation401ApplicationJSONObject *CreatePullOperation401ApplicationJSON
	CreatePullOperation404ApplicationJSONObject *CreatePullOperation404ApplicationJSON
}
