package operations

import (
	"net/http"
	"time"
)

type UpdateCompanyRequestBody struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type UpdateCompanyRequest struct {
	RequestBody *UpdateCompanyRequestBody `request:"mediaType=application/json"`
	CompanyID   string                    `pathParam:"style=simple,explode=false,name=companyId"`
}

type UpdateCompany401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type UpdateCompanyCompanyConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type UpdateCompanyCompanyConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type UpdateCompanyCompanyConnectionSourceTypeEnum string

const (
	UpdateCompanyCompanyConnectionSourceTypeEnumAccounting UpdateCompanyCompanyConnectionSourceTypeEnum = "Accounting"
	UpdateCompanyCompanyConnectionSourceTypeEnumBanking    UpdateCompanyCompanyConnectionSourceTypeEnum = "Banking"
	UpdateCompanyCompanyConnectionSourceTypeEnumCommerce   UpdateCompanyCompanyConnectionSourceTypeEnum = "Commerce"
	UpdateCompanyCompanyConnectionSourceTypeEnumOther      UpdateCompanyCompanyConnectionSourceTypeEnum = "Other"
	UpdateCompanyCompanyConnectionSourceTypeEnumUnknown    UpdateCompanyCompanyConnectionSourceTypeEnum = "Unknown"
)

type UpdateCompanyCompanyConnectionDataConnectionStatusEnum string

const (
	UpdateCompanyCompanyConnectionDataConnectionStatusEnumPendingAuth  UpdateCompanyCompanyConnectionDataConnectionStatusEnum = "PendingAuth"
	UpdateCompanyCompanyConnectionDataConnectionStatusEnumLinked       UpdateCompanyCompanyConnectionDataConnectionStatusEnum = "Linked"
	UpdateCompanyCompanyConnectionDataConnectionStatusEnumUnlinked     UpdateCompanyCompanyConnectionDataConnectionStatusEnum = "Unlinked"
	UpdateCompanyCompanyConnectionDataConnectionStatusEnumDeauthorized UpdateCompanyCompanyConnectionDataConnectionStatusEnum = "Deauthorized"
)

// UpdateCompanyCompanyConnection
// A connection represents the link between a `company` and a source of data.
type UpdateCompanyCompanyConnection struct {
	ConnectionInfo       *UpdateCompanyCompanyConnectionConnectionInfo          `json:"connectionInfo,omitempty"`
	Created              time.Time                                              `json:"created"`
	DataConnectionErrors []UpdateCompanyCompanyConnectionDataConnectionErrors   `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                 `json:"id"`
	IntegrationID        string                                                 `json:"integrationId"`
	IntegrationKey       string                                                 `json:"integrationKey"`
	LastSync             *time.Time                                             `json:"lastSync,omitempty"`
	LinkURL              string                                                 `json:"linkUrl"`
	PlatformName         string                                                 `json:"platformName"`
	SourceID             string                                                 `json:"sourceId"`
	SourceType           UpdateCompanyCompanyConnectionSourceTypeEnum           `json:"sourceType"`
	Status               UpdateCompanyCompanyConnectionDataConnectionStatusEnum `json:"status"`
}

// UpdateCompanyCompany
// A company in Codat represent a small or medium sized business, whose data you wish to share
type UpdateCompanyCompany struct {
	Created           *time.Time                       `json:"created,omitempty"`
	CreatedByUserName *string                          `json:"createdByUserName,omitempty"`
	DataConnections   []UpdateCompanyCompanyConnection `json:"dataConnections,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	ID                string                           `json:"id"`
	LastSync          *time.Time                       `json:"lastSync,omitempty"`
	Name              string                           `json:"name"`
	Platform          *string                          `json:"platform,omitempty"`
	Redirect          string                           `json:"redirect"`
}

type UpdateCompanyResponse struct {
	Company                               *UpdateCompanyCompany
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	UpdateCompany401ApplicationJSONObject *UpdateCompany401ApplicationJSON
}
