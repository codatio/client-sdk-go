package operations

import (
	"net/http"
	"time"
)

type GetCompanyConnectionPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCompanyConnectionRequest struct {
	PathParams GetCompanyConnectionPathParams
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

type GetCompanyConnectionConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type GetCompanyConnectionConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type GetCompanyConnectionConnectionSourceTypeEnum string

const (
	GetCompanyConnectionConnectionSourceTypeEnumAccounting GetCompanyConnectionConnectionSourceTypeEnum = "Accounting"
	GetCompanyConnectionConnectionSourceTypeEnumBanking    GetCompanyConnectionConnectionSourceTypeEnum = "Banking"
	GetCompanyConnectionConnectionSourceTypeEnumCommerce   GetCompanyConnectionConnectionSourceTypeEnum = "Commerce"
	GetCompanyConnectionConnectionSourceTypeEnumOther      GetCompanyConnectionConnectionSourceTypeEnum = "Other"
	GetCompanyConnectionConnectionSourceTypeEnumUnknown    GetCompanyConnectionConnectionSourceTypeEnum = "Unknown"
)

type GetCompanyConnectionConnectionDataConnectionStatusEnum string

const (
	GetCompanyConnectionConnectionDataConnectionStatusEnumPendingAuth  GetCompanyConnectionConnectionDataConnectionStatusEnum = "PendingAuth"
	GetCompanyConnectionConnectionDataConnectionStatusEnumLinked       GetCompanyConnectionConnectionDataConnectionStatusEnum = "Linked"
	GetCompanyConnectionConnectionDataConnectionStatusEnumUnlinked     GetCompanyConnectionConnectionDataConnectionStatusEnum = "Unlinked"
	GetCompanyConnectionConnectionDataConnectionStatusEnumDeauthorized GetCompanyConnectionConnectionDataConnectionStatusEnum = "Deauthorized"
)

// GetCompanyConnectionConnection
// A connection represents the link between a `company` and a source of data.
type GetCompanyConnectionConnection struct {
	ConnectionInfo       *GetCompanyConnectionConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                              `json:"created"`
	DataConnectionErrors []GetCompanyConnectionConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                 `json:"id"`
	IntegrationID        string                                                 `json:"integrationId"`
	IntegrationKey       string                                                 `json:"integrationKey"`
	LastSync             *time.Time                                             `json:"lastSync,omitempty"`
	LinkURL              string                                                 `json:"linkUrl"`
	PlatformName         string                                                 `json:"platformName"`
	SourceID             string                                                 `json:"sourceId"`
	SourceType           GetCompanyConnectionConnectionSourceTypeEnum           `json:"sourceType"`
	Status               GetCompanyConnectionConnectionDataConnectionStatusEnum `json:"status"`
}

type GetCompanyConnectionResponse struct {
	Connection                                   *GetCompanyConnectionConnection
	ContentType                                  string
	StatusCode                                   int
	RawResponse                                  *http.Response
	GetCompanyConnection401ApplicationJSONObject *GetCompanyConnection401ApplicationJSON
	GetCompanyConnection404ApplicationJSONObject *GetCompanyConnection404ApplicationJSON
}
