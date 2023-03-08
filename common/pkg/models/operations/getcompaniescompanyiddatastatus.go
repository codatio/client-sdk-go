package operations

import (
	"net/http"
	"time"
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

// GetCompaniesCompanyIDDataStatus200ApplicationJSONDataStatus
// Describes the state of data in the Codat cache for a company and data type
type GetCompaniesCompanyIDDataStatus200ApplicationJSONDataStatus struct {
	CurrentStatus          string    `json:"currentStatus"`
	DataType               string    `json:"dataType"`
	LastSuccessfulSync     time.Time `json:"lastSuccessfulSync"`
	LatestSuccessfulSyncID *string   `json:"latestSuccessfulSyncId,omitempty"`
	LatestSyncID           *string   `json:"latestSyncId,omitempty"`
}

type GetCompaniesCompanyIDDataStatus200ApplicationJSON struct {
	DataType1 *GetCompaniesCompanyIDDataStatus200ApplicationJSONDataStatus `json:"dataType1,omitempty"`
	DataType2 *GetCompaniesCompanyIDDataStatus200ApplicationJSONDataStatus `json:"dataType2,omitempty"`
}

type GetCompaniesCompanyIDDataStatusResponse struct {
	ContentType                                             string
	StatusCode                                              int
	RawResponse                                             *http.Response
	GetCompaniesCompanyIDDataStatus200ApplicationJSONObject *GetCompaniesCompanyIDDataStatus200ApplicationJSON
	GetCompaniesCompanyIDDataStatus401ApplicationJSONObject *GetCompaniesCompanyIDDataStatus401ApplicationJSON
	GetCompaniesCompanyIDDataStatus404ApplicationJSONObject *GetCompaniesCompanyIDDataStatus404ApplicationJSON
}
