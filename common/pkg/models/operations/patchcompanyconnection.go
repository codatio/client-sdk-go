package operations

import (
	"time"
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

type PatchCompanyConnectionConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type PatchCompanyConnectionConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type PatchCompanyConnectionConnectionSourceTypeEnum string

const (
	PatchCompanyConnectionConnectionSourceTypeEnumAccounting PatchCompanyConnectionConnectionSourceTypeEnum = "Accounting"
	PatchCompanyConnectionConnectionSourceTypeEnumBanking    PatchCompanyConnectionConnectionSourceTypeEnum = "Banking"
	PatchCompanyConnectionConnectionSourceTypeEnumCommerce   PatchCompanyConnectionConnectionSourceTypeEnum = "Commerce"
	PatchCompanyConnectionConnectionSourceTypeEnumOther      PatchCompanyConnectionConnectionSourceTypeEnum = "Other"
	PatchCompanyConnectionConnectionSourceTypeEnumUnknown    PatchCompanyConnectionConnectionSourceTypeEnum = "Unknown"
)

type PatchCompanyConnectionConnectionDataConnectionStatusEnum string

const (
	PatchCompanyConnectionConnectionDataConnectionStatusEnumPendingAuth  PatchCompanyConnectionConnectionDataConnectionStatusEnum = "PendingAuth"
	PatchCompanyConnectionConnectionDataConnectionStatusEnumLinked       PatchCompanyConnectionConnectionDataConnectionStatusEnum = "Linked"
	PatchCompanyConnectionConnectionDataConnectionStatusEnumUnlinked     PatchCompanyConnectionConnectionDataConnectionStatusEnum = "Unlinked"
	PatchCompanyConnectionConnectionDataConnectionStatusEnumDeauthorized PatchCompanyConnectionConnectionDataConnectionStatusEnum = "Deauthorized"
)

// PatchCompanyConnectionConnection
// A connection represents the link between a `company` and a source of data.
type PatchCompanyConnectionConnection struct {
	ConnectionInfo       *PatchCompanyConnectionConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                                `json:"created"`
	DataConnectionErrors []PatchCompanyConnectionConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                   `json:"id"`
	IntegrationID        string                                                   `json:"integrationId"`
	IntegrationKey       string                                                   `json:"integrationKey"`
	LastSync             *time.Time                                               `json:"lastSync,omitempty"`
	LinkURL              string                                                   `json:"linkUrl"`
	PlatformName         string                                                   `json:"platformName"`
	SourceID             string                                                   `json:"sourceId"`
	SourceType           PatchCompanyConnectionConnectionSourceTypeEnum           `json:"sourceType"`
	Status               PatchCompanyConnectionConnectionDataConnectionStatusEnum `json:"status"`
}

type PatchCompanyConnectionResponse struct {
	Connection                                     *PatchCompanyConnectionConnection
	ContentType                                    string
	StatusCode                                     int
	PatchCompanyConnection401ApplicationJSONObject *PatchCompanyConnection401ApplicationJSON
	PatchCompanyConnection404ApplicationJSONObject *PatchCompanyConnection404ApplicationJSON
}
