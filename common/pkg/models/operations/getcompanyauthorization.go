package operations

import (
	"net/http"
	"time"
)

type GetCompanyAuthorizationPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCompanyAuthorizationRequestBody struct {
	Property1 *string `json:"property1,omitempty"`
	Property2 *string `json:"property2,omitempty"`
	Property3 *string `json:"property3,omitempty"`
}

type GetCompanyAuthorizationRequest struct {
	PathParams GetCompanyAuthorizationPathParams
	Request    *GetCompanyAuthorizationRequestBody `request:"mediaType=application/json"`
}

type GetCompanyAuthorizationConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type GetCompanyAuthorizationConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type GetCompanyAuthorizationConnectionSourceTypeEnum string

const (
	GetCompanyAuthorizationConnectionSourceTypeEnumAccounting GetCompanyAuthorizationConnectionSourceTypeEnum = "Accounting"
	GetCompanyAuthorizationConnectionSourceTypeEnumBanking    GetCompanyAuthorizationConnectionSourceTypeEnum = "Banking"
	GetCompanyAuthorizationConnectionSourceTypeEnumCommerce   GetCompanyAuthorizationConnectionSourceTypeEnum = "Commerce"
	GetCompanyAuthorizationConnectionSourceTypeEnumOther      GetCompanyAuthorizationConnectionSourceTypeEnum = "Other"
	GetCompanyAuthorizationConnectionSourceTypeEnumUnknown    GetCompanyAuthorizationConnectionSourceTypeEnum = "Unknown"
)

type GetCompanyAuthorizationConnectionDataConnectionStatusEnum string

const (
	GetCompanyAuthorizationConnectionDataConnectionStatusEnumPendingAuth  GetCompanyAuthorizationConnectionDataConnectionStatusEnum = "PendingAuth"
	GetCompanyAuthorizationConnectionDataConnectionStatusEnumLinked       GetCompanyAuthorizationConnectionDataConnectionStatusEnum = "Linked"
	GetCompanyAuthorizationConnectionDataConnectionStatusEnumUnlinked     GetCompanyAuthorizationConnectionDataConnectionStatusEnum = "Unlinked"
	GetCompanyAuthorizationConnectionDataConnectionStatusEnumDeauthorized GetCompanyAuthorizationConnectionDataConnectionStatusEnum = "Deauthorized"
)

// GetCompanyAuthorizationConnection
// A connection represents the link between a `company` and a source of data.
type GetCompanyAuthorizationConnection struct {
	ConnectionInfo       *GetCompanyAuthorizationConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                                 `json:"created"`
	DataConnectionErrors []GetCompanyAuthorizationConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                    `json:"id"`
	IntegrationID        string                                                    `json:"integrationId"`
	IntegrationKey       string                                                    `json:"integrationKey"`
	LastSync             *time.Time                                                `json:"lastSync,omitempty"`
	LinkURL              string                                                    `json:"linkUrl"`
	PlatformName         string                                                    `json:"platformName"`
	SourceID             string                                                    `json:"sourceId"`
	SourceType           GetCompanyAuthorizationConnectionSourceTypeEnum           `json:"sourceType"`
	Status               GetCompanyAuthorizationConnectionDataConnectionStatusEnum `json:"status"`
}

type GetCompanyAuthorizationResponse struct {
	Connection  *GetCompanyAuthorizationConnection
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
