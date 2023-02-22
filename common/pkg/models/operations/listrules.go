package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type ListRulesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListRulesRequest struct {
	QueryParams ListRulesQueryParams
}

type ListRules401ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListRules400ApplicationJSON struct {
	CanBeRetried      *string `json:"canBeRetried,omitempty"`
	CorrelationID     *string `json:"correlationId,omitempty"`
	DetailedErrorCode *int64  `json:"detailedErrorCode,omitempty"`
	Error             *string `json:"error,omitempty"`
	Service           *string `json:"service,omitempty"`
	StatusCode        *int64  `json:"statusCode,omitempty"`
}

type ListRulesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListRulesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListRulesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListRulesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListRulesLinksLinks struct {
	Current  ListRulesLinksLinksCurrent   `json:"current"`
	Next     *ListRulesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListRulesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListRulesLinksLinksSelf      `json:"self"`
}

// ListRulesLinks
// Codat's Paging Model
type ListRulesLinks struct {
	Links        ListRulesLinksLinks `json:"_links"`
	PageNumber   int64               `json:"pageNumber"`
	PageSize     int64               `json:"pageSize"`
	Results      []shared.Rule       `json:"results,omitempty"`
	TotalResults int64               `json:"totalResults"`
}

type ListRulesResponse struct {
	ContentType                       string
	StatusCode                        int
	Links                             *ListRulesLinks
	ListRules400ApplicationJSONObject *ListRules400ApplicationJSON
	ListRules401ApplicationJSONObject *ListRules401ApplicationJSON
}
