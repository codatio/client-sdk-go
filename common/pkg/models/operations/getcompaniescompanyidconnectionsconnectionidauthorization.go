package operations

import (
	"time"
)

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationRequestBody struct {
	Property1 *string `json:"property1,omitempty"`
	Property2 *string `json:"property2,omitempty"`
	Property3 *string `json:"property3,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationRequest struct {
	PathParams GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationPathParams
	Request    *GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationRequestBody `request:"mediaType=application/json"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnumAccounting GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnum = "Accounting"
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnumBanking    GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnum = "Banking"
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnumCommerce   GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnum = "Commerce"
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnumOther      GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnum = "Other"
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnumUnknown    GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnum = "Unknown"
)

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnum string

const (
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnumPendingAuth  GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnum = "PendingAuth"
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnumLinked       GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnum = "Linked"
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnumUnlinked     GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnum = "Unlinked"
	GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnumDeauthorized GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnum = "Deauthorized"
)

// GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnection
// A connection represents the link between a `company` and a source of data.
type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnection struct {
	ConnectionInfo       *GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                                                                   `json:"created"`
	DataConnectionErrors []GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                                                      `json:"id"`
	IntegrationID        string                                                                                      `json:"integrationId"`
	IntegrationKey       string                                                                                      `json:"integrationKey"`
	LastSync             *time.Time                                                                                  `json:"lastSync,omitempty"`
	LinkURL              string                                                                                      `json:"linkUrl"`
	PlatformName         string                                                                                      `json:"platformName"`
	SourceID             string                                                                                      `json:"sourceId"`
	SourceType           GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionSourceTypeEnum           `json:"sourceType"`
	Status               GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnectionDataConnectionStatusEnum `json:"status"`
}

type GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationResponse struct {
	Connection  *GetCompaniesCompanyIDConnectionsConnectionIDAuthorizationConnection
	ContentType string
	StatusCode  int
}
