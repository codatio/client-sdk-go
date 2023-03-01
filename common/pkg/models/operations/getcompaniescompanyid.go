package operations

import (
	"time"
)

type GetCompaniesCompanyIDPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDRequest struct {
	PathParams GetCompaniesCompanyIDPathParams
}

type GetCompaniesCompanyID401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompaniesCompanyIDCompanyConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type GetCompaniesCompanyIDCompanyConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type GetCompaniesCompanyIDCompanyConnectionSourceTypeEnum string

const (
	GetCompaniesCompanyIDCompanyConnectionSourceTypeEnumAccounting GetCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Accounting"
	GetCompaniesCompanyIDCompanyConnectionSourceTypeEnumBanking    GetCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Banking"
	GetCompaniesCompanyIDCompanyConnectionSourceTypeEnumCommerce   GetCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Commerce"
	GetCompaniesCompanyIDCompanyConnectionSourceTypeEnumOther      GetCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Other"
	GetCompaniesCompanyIDCompanyConnectionSourceTypeEnumUnknown    GetCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Unknown"
)

type GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnum string

const (
	GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnumPendingAuth  GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnum = "PendingAuth"
	GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnumLinked       GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnum = "Linked"
	GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnumUnlinked     GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnum = "Unlinked"
	GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnumDeauthorized GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnum = "Deauthorized"
)

// GetCompaniesCompanyIDCompanyConnection
// A connection represents the link between a `company` and a source of data.
type GetCompaniesCompanyIDCompanyConnection struct {
	ConnectionInfo       *GetCompaniesCompanyIDCompanyConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                                      `json:"created"`
	DataConnectionErrors []GetCompaniesCompanyIDCompanyConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                         `json:"id"`
	IntegrationID        string                                                         `json:"integrationId"`
	IntegrationKey       string                                                         `json:"integrationKey"`
	LastSync             *time.Time                                                     `json:"lastSync,omitempty"`
	LinkURL              string                                                         `json:"linkUrl"`
	PlatformName         string                                                         `json:"platformName"`
	SourceID             string                                                         `json:"sourceId"`
	SourceType           GetCompaniesCompanyIDCompanyConnectionSourceTypeEnum           `json:"sourceType"`
	Status               GetCompaniesCompanyIDCompanyConnectionDataConnectionStatusEnum `json:"status"`
}

// GetCompaniesCompanyIDCompany
// A company in Codat represent a small or medium sized business, whose data you wish to share
type GetCompaniesCompanyIDCompany struct {
	Created           *time.Time                               `json:"created,omitempty"`
	CreatedByUserName *string                                  `json:"createdByUserName,omitempty"`
	DataConnections   []GetCompaniesCompanyIDCompanyConnection `json:"dataConnections,omitempty"`
	Description       *string                                  `json:"description,omitempty"`
	ID                string                                   `json:"id"`
	LastSync          *time.Time                               `json:"lastSync,omitempty"`
	Name              string                                   `json:"name"`
	Platform          *string                                  `json:"platform,omitempty"`
	Redirect          string                                   `json:"redirect"`
}

type GetCompaniesCompanyIDResponse struct {
	Company                                       *GetCompaniesCompanyIDCompany
	ContentType                                   string
	StatusCode                                    int
	GetCompaniesCompanyID401ApplicationJSONObject *GetCompaniesCompanyID401ApplicationJSON
}
