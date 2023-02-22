package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetCompanyDataHistoryPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompanyDataHistoryQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetCompanyDataHistoryRequest struct {
	PathParams  GetCompanyDataHistoryPathParams
	QueryParams GetCompanyDataHistoryQueryParams
}

type GetCompanyDataHistory404ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyDataHistory401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyDataHistory400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type GetCompanyDataHistoryLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetCompanyDataHistoryLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetCompanyDataHistoryLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetCompanyDataHistoryLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetCompanyDataHistoryLinksLinks struct {
	Current  GetCompanyDataHistoryLinksLinksCurrent   `json:"current"`
	Next     *GetCompanyDataHistoryLinksLinksNext     `json:"next,omitempty"`
	Previous *GetCompanyDataHistoryLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetCompanyDataHistoryLinksLinksSelf      `json:"self"`
}

// GetCompanyDataHistoryLinks
// Codat's Paging Model
type GetCompanyDataHistoryLinks struct {
	Links        GetCompanyDataHistoryLinksLinks `json:"_links"`
	PageNumber   int64                           `json:"pageNumber"`
	PageSize     int64                           `json:"pageSize"`
	Results      []shared.PullOperation          `json:"results,omitempty"`
	TotalResults int64                           `json:"totalResults"`
}

type GetCompanyDataHistoryResponse struct {
	ContentType                                   string
	StatusCode                                    int
	Links                                         *GetCompanyDataHistoryLinks
	GetCompanyDataHistory400ApplicationJSONObject *GetCompanyDataHistory400ApplicationJSON
	GetCompanyDataHistory401ApplicationJSONObject *GetCompanyDataHistory401ApplicationJSON
	GetCompanyDataHistory404ApplicationJSONObject *GetCompanyDataHistory404ApplicationJSON
}
