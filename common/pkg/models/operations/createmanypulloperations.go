package operations

import (
	"net/http"
)

type CreateManyPullOperationsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type CreateManyPullOperationsRequest struct {
	PathParams CreateManyPullOperationsPathParams
}

type CreateManyPullOperations404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type CreateManyPullOperations401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type CreateManyPullOperationsResponse struct {
	ContentType                                      string
	StatusCode                                       int
	RawResponse                                      *http.Response
	CreateManyPullOperations401ApplicationJSONObject *CreateManyPullOperations401ApplicationJSON
	CreateManyPullOperations404ApplicationJSONObject *CreateManyPullOperations404ApplicationJSON
}
