package operations

import (
	"time"
)

type PutCompaniesCompanyIDPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type PutCompaniesCompanyIDRequestBody struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type PutCompaniesCompanyIDRequest struct {
	PathParams PutCompaniesCompanyIDPathParams
	Request    *PutCompaniesCompanyIDRequestBody `request:"mediaType=application/json"`
}

type PutCompaniesCompanyID401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type PutCompaniesCompanyIDCompanyConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type PutCompaniesCompanyIDCompanyConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type PutCompaniesCompanyIDCompanyConnectionSourceTypeEnum string

const (
	PutCompaniesCompanyIDCompanyConnectionSourceTypeEnumAccounting PutCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Accounting"
	PutCompaniesCompanyIDCompanyConnectionSourceTypeEnumBanking    PutCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Banking"
	PutCompaniesCompanyIDCompanyConnectionSourceTypeEnumCommerce   PutCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Commerce"
	PutCompaniesCompanyIDCompanyConnectionSourceTypeEnumOther      PutCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Other"
	PutCompaniesCompanyIDCompanyConnectionSourceTypeEnumUnknown    PutCompaniesCompanyIDCompanyConnectionSourceTypeEnum = "Unknown"
)

type PutCompaniesCompanyIDCompanyConnectionStatusEnum string

const (
	PutCompaniesCompanyIDCompanyConnectionStatusEnumPendingAuth  PutCompaniesCompanyIDCompanyConnectionStatusEnum = "PendingAuth"
	PutCompaniesCompanyIDCompanyConnectionStatusEnumLinked       PutCompaniesCompanyIDCompanyConnectionStatusEnum = "Linked"
	PutCompaniesCompanyIDCompanyConnectionStatusEnumUnlinked     PutCompaniesCompanyIDCompanyConnectionStatusEnum = "Unlinked"
	PutCompaniesCompanyIDCompanyConnectionStatusEnumDeauthorized PutCompaniesCompanyIDCompanyConnectionStatusEnum = "Deauthorized"
)

// PutCompaniesCompanyIDCompanyConnection
// A connection represents the link between a `company` and a source of data.
type PutCompaniesCompanyIDCompanyConnection struct {
	ConnectionInfo       *PutCompaniesCompanyIDCompanyConnectionConnectionInfo        `json:"connectionInfo,omitempty"`
	Created              time.Time                                                    `json:"created"`
	DataConnectionErrors []PutCompaniesCompanyIDCompanyConnectionDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                       `json:"id"`
	IntegrationID        string                                                       `json:"integrationId"`
	IntegrationKey       string                                                       `json:"integrationKey"`
	LastSync             *time.Time                                                   `json:"lastSync,omitempty"`
	LinkURL              string                                                       `json:"linkUrl"`
	PlatformName         string                                                       `json:"platformName"`
	SourceID             string                                                       `json:"sourceId"`
	SourceType           PutCompaniesCompanyIDCompanyConnectionSourceTypeEnum         `json:"sourceType"`
	Status               PutCompaniesCompanyIDCompanyConnectionStatusEnum             `json:"status"`
}

// PutCompaniesCompanyIDCompany
// A company in Codat represent a small or medium sized business, whose data you wish to share
type PutCompaniesCompanyIDCompany struct {
	Created           *time.Time                               `json:"created,omitempty"`
	CreatedByUserName *string                                  `json:"createdByUserName,omitempty"`
	DataConnections   []PutCompaniesCompanyIDCompanyConnection `json:"dataConnections,omitempty"`
	Description       *string                                  `json:"description,omitempty"`
	ID                string                                   `json:"id"`
	LastSync          *time.Time                               `json:"lastSync,omitempty"`
	Name              string                                   `json:"name"`
	Platform          *string                                  `json:"platform,omitempty"`
	Redirect          string                                   `json:"redirect"`
}

type PutCompaniesCompanyIDResponse struct {
	Company                                       *PutCompaniesCompanyIDCompany
	ContentType                                   string
	StatusCode                                    int
	PutCompaniesCompanyID401ApplicationJSONObject *PutCompaniesCompanyID401ApplicationJSON
}
