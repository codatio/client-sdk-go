package operations

import (
	"net/http"
	"time"
)

type GetCompanyPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyRequest struct {
	PathParams GetCompanyPathParams
}

type GetCompany401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyCompanyConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type GetCompanyCompanyConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type GetCompanyCompanyConnectionSourceTypeEnum string

const (
	GetCompanyCompanyConnectionSourceTypeEnumAccounting GetCompanyCompanyConnectionSourceTypeEnum = "Accounting"
	GetCompanyCompanyConnectionSourceTypeEnumBanking    GetCompanyCompanyConnectionSourceTypeEnum = "Banking"
	GetCompanyCompanyConnectionSourceTypeEnumCommerce   GetCompanyCompanyConnectionSourceTypeEnum = "Commerce"
	GetCompanyCompanyConnectionSourceTypeEnumOther      GetCompanyCompanyConnectionSourceTypeEnum = "Other"
	GetCompanyCompanyConnectionSourceTypeEnumUnknown    GetCompanyCompanyConnectionSourceTypeEnum = "Unknown"
)

type GetCompanyCompanyConnectionDataConnectionStatusEnum string

const (
	GetCompanyCompanyConnectionDataConnectionStatusEnumPendingAuth  GetCompanyCompanyConnectionDataConnectionStatusEnum = "PendingAuth"
	GetCompanyCompanyConnectionDataConnectionStatusEnumLinked       GetCompanyCompanyConnectionDataConnectionStatusEnum = "Linked"
	GetCompanyCompanyConnectionDataConnectionStatusEnumUnlinked     GetCompanyCompanyConnectionDataConnectionStatusEnum = "Unlinked"
	GetCompanyCompanyConnectionDataConnectionStatusEnumDeauthorized GetCompanyCompanyConnectionDataConnectionStatusEnum = "Deauthorized"
)

// GetCompanyCompanyConnection
// A connection represents the link between a `company` and a source of data.
type GetCompanyCompanyConnection struct {
	ConnectionInfo       *GetCompanyCompanyConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                           `json:"created"`
	DataConnectionErrors []GetCompanyCompanyConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                              `json:"id"`
	IntegrationID        string                                              `json:"integrationId"`
	IntegrationKey       string                                              `json:"integrationKey"`
	LastSync             *time.Time                                          `json:"lastSync,omitempty"`
	LinkURL              string                                              `json:"linkUrl"`
	PlatformName         string                                              `json:"platformName"`
	SourceID             string                                              `json:"sourceId"`
	SourceType           GetCompanyCompanyConnectionSourceTypeEnum           `json:"sourceType"`
	Status               GetCompanyCompanyConnectionDataConnectionStatusEnum `json:"status"`
}

// GetCompanyCompany
// A company in Codat represent a small or medium sized business, whose data you wish to share
type GetCompanyCompany struct {
	Created           *time.Time                    `json:"created,omitempty"`
	CreatedByUserName *string                       `json:"createdByUserName,omitempty"`
	DataConnections   []GetCompanyCompanyConnection `json:"dataConnections,omitempty"`
	Description       *string                       `json:"description,omitempty"`
	ID                string                        `json:"id"`
	LastSync          *time.Time                    `json:"lastSync,omitempty"`
	Name              string                        `json:"name"`
	Platform          *string                       `json:"platform,omitempty"`
	Redirect          string                        `json:"redirect"`
}

type GetCompanyResponse struct {
	Company                            *GetCompanyCompany
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	GetCompany401ApplicationJSONObject *GetCompany401ApplicationJSON
}
