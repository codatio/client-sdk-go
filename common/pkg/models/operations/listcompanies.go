package operations

import (
	"time"
)

type ListCompaniesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCompaniesRequest struct {
	QueryParams ListCompaniesQueryParams
}

type ListCompanies401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListCompanies400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListCompaniesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCompaniesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCompaniesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCompaniesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCompaniesLinksLinks struct {
	Current  ListCompaniesLinksLinksCurrent   `json:"current"`
	Next     *ListCompaniesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCompaniesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCompaniesLinksLinksSelf      `json:"self"`
}

type ListCompaniesLinksCompanyConnectionConnectionInfo struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty"`
}

type ListCompaniesLinksCompanyConnectionDataConnectionErrors struct {
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	ErroredOnUtc *time.Time `json:"erroredOnUtc,omitempty"`
	StatusCode   *string    `json:"statusCode,omitempty"`
	StatusText   *string    `json:"statusText,omitempty"`
}

type ListCompaniesLinksCompanyConnectionSourceTypeEnum string

const (
	ListCompaniesLinksCompanyConnectionSourceTypeEnumAccounting ListCompaniesLinksCompanyConnectionSourceTypeEnum = "Accounting"
	ListCompaniesLinksCompanyConnectionSourceTypeEnumBanking    ListCompaniesLinksCompanyConnectionSourceTypeEnum = "Banking"
	ListCompaniesLinksCompanyConnectionSourceTypeEnumCommerce   ListCompaniesLinksCompanyConnectionSourceTypeEnum = "Commerce"
	ListCompaniesLinksCompanyConnectionSourceTypeEnumOther      ListCompaniesLinksCompanyConnectionSourceTypeEnum = "Other"
	ListCompaniesLinksCompanyConnectionSourceTypeEnumUnknown    ListCompaniesLinksCompanyConnectionSourceTypeEnum = "Unknown"
)

type ListCompaniesLinksCompanyConnectionStatusEnum string

const (
	ListCompaniesLinksCompanyConnectionStatusEnumPendingAuth  ListCompaniesLinksCompanyConnectionStatusEnum = "PendingAuth"
	ListCompaniesLinksCompanyConnectionStatusEnumLinked       ListCompaniesLinksCompanyConnectionStatusEnum = "Linked"
	ListCompaniesLinksCompanyConnectionStatusEnumUnlinked     ListCompaniesLinksCompanyConnectionStatusEnum = "Unlinked"
	ListCompaniesLinksCompanyConnectionStatusEnumDeauthorized ListCompaniesLinksCompanyConnectionStatusEnum = "Deauthorized"
)

// ListCompaniesLinksCompanyConnection
// A connection represents the link between a `company` and a source of data.
type ListCompaniesLinksCompanyConnection struct {
	ConnectionInfo       *ListCompaniesLinksCompanyConnectionConnectionInfo        `json:"connectionInfo,omitempty"`
	Created              time.Time                                                 `json:"created"`
	DataConnectionErrors []ListCompaniesLinksCompanyConnectionDataConnectionErrors `json:"dataConnectionErrors,omitempty"`
	ID                   string                                                    `json:"id"`
	IntegrationID        string                                                    `json:"integrationId"`
	IntegrationKey       string                                                    `json:"integrationKey"`
	LastSync             *time.Time                                                `json:"lastSync,omitempty"`
	LinkURL              string                                                    `json:"linkUrl"`
	PlatformName         string                                                    `json:"platformName"`
	SourceID             string                                                    `json:"sourceId"`
	SourceType           ListCompaniesLinksCompanyConnectionSourceTypeEnum         `json:"sourceType"`
	Status               ListCompaniesLinksCompanyConnectionStatusEnum             `json:"status"`
}

// ListCompaniesLinksCompany
// A company in Codat represent a small or medium sized business, whose data you wish to share
type ListCompaniesLinksCompany struct {
	Created           *time.Time                            `json:"created,omitempty"`
	CreatedByUserName *string                               `json:"createdByUserName,omitempty"`
	DataConnections   []ListCompaniesLinksCompanyConnection `json:"dataConnections,omitempty"`
	Description       *string                               `json:"description,omitempty"`
	ID                string                                `json:"id"`
	LastSync          *time.Time                            `json:"lastSync,omitempty"`
	Name              string                                `json:"name"`
	Platform          *string                               `json:"platform,omitempty"`
	Redirect          string                                `json:"redirect"`
}

// ListCompaniesLinks
// Codat's Paging Model
type ListCompaniesLinks struct {
	Links        ListCompaniesLinksLinks     `json:"_links"`
	PageNumber   int64                       `json:"pageNumber"`
	PageSize     int64                       `json:"pageSize"`
	Results      []ListCompaniesLinksCompany `json:"results,omitempty"`
	TotalResults int64                       `json:"totalResults"`
}

type ListCompaniesResponse struct {
	ContentType                           string
	StatusCode                            int
	Links                                 *ListCompaniesLinks
	ListCompanies400ApplicationJSONObject *ListCompanies400ApplicationJSON
	ListCompanies401ApplicationJSONObject *ListCompanies401ApplicationJSON
}
