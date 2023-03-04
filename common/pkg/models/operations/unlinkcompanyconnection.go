package operations

import (
	"net/http"
	"time"
)

type UnlinkCompanyConnectionPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type UnlinkCompanyConnectionRequestBody struct {
	Status *string `json:"status,omitempty"`
}

type UnlinkCompanyConnectionRequest struct {
	PathParams UnlinkCompanyConnectionPathParams
	Request    *UnlinkCompanyConnectionRequestBody `request:"mediaType=application/json"`
}

type UnlinkCompanyConnection404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type UnlinkCompanyConnection401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type UnlinkCompanyConnectionConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type UnlinkCompanyConnectionConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type UnlinkCompanyConnectionConnectionSourceTypeEnum string

const (
	UnlinkCompanyConnectionConnectionSourceTypeEnumAccounting UnlinkCompanyConnectionConnectionSourceTypeEnum = "Accounting"
	UnlinkCompanyConnectionConnectionSourceTypeEnumBanking    UnlinkCompanyConnectionConnectionSourceTypeEnum = "Banking"
	UnlinkCompanyConnectionConnectionSourceTypeEnumCommerce   UnlinkCompanyConnectionConnectionSourceTypeEnum = "Commerce"
	UnlinkCompanyConnectionConnectionSourceTypeEnumOther      UnlinkCompanyConnectionConnectionSourceTypeEnum = "Other"
	UnlinkCompanyConnectionConnectionSourceTypeEnumUnknown    UnlinkCompanyConnectionConnectionSourceTypeEnum = "Unknown"
)

type UnlinkCompanyConnectionConnectionDataConnectionStatusEnum string

const (
	UnlinkCompanyConnectionConnectionDataConnectionStatusEnumPendingAuth  UnlinkCompanyConnectionConnectionDataConnectionStatusEnum = "PendingAuth"
	UnlinkCompanyConnectionConnectionDataConnectionStatusEnumLinked       UnlinkCompanyConnectionConnectionDataConnectionStatusEnum = "Linked"
	UnlinkCompanyConnectionConnectionDataConnectionStatusEnumUnlinked     UnlinkCompanyConnectionConnectionDataConnectionStatusEnum = "Unlinked"
	UnlinkCompanyConnectionConnectionDataConnectionStatusEnumDeauthorized UnlinkCompanyConnectionConnectionDataConnectionStatusEnum = "Deauthorized"
)

// UnlinkCompanyConnectionConnection
// A connection represents the link between a `company` and a source of data.
type UnlinkCompanyConnectionConnection struct {
	ConnectionInfo       *UnlinkCompanyConnectionConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                                 `json:"created"`
	DataConnectionErrors []UnlinkCompanyConnectionConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                    `json:"id"`
	IntegrationID        string                                                    `json:"integrationId"`
	IntegrationKey       string                                                    `json:"integrationKey"`
	LastSync             *time.Time                                                `json:"lastSync,omitempty"`
	LinkURL              string                                                    `json:"linkUrl"`
	PlatformName         string                                                    `json:"platformName"`
	SourceID             string                                                    `json:"sourceId"`
	SourceType           UnlinkCompanyConnectionConnectionSourceTypeEnum           `json:"sourceType"`
	Status               UnlinkCompanyConnectionConnectionDataConnectionStatusEnum `json:"status"`
}

type UnlinkCompanyConnectionResponse struct {
	Connection                                      *UnlinkCompanyConnectionConnection
	ContentType                                     string
	StatusCode                                      int
	RawResponse                                     *http.Response
	UnlinkCompanyConnection401ApplicationJSONObject *UnlinkCompanyConnection401ApplicationJSON
	UnlinkCompanyConnection404ApplicationJSONObject *UnlinkCompanyConnection404ApplicationJSON
}
