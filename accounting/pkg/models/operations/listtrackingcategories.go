package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListTrackingCategoriesPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListTrackingCategoriesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListTrackingCategoriesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListTrackingCategoriesRequest struct {
	PathParams  ListTrackingCategoriesPathParams
	QueryParams ListTrackingCategoriesQueryParams
	Security    ListTrackingCategoriesSecurity
}

type ListTrackingCategoriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListTrackingCategoriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListTrackingCategoriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListTrackingCategoriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListTrackingCategoriesLinksLinks struct {
	Current  ListTrackingCategoriesLinksLinksCurrent   `json:"current"`
	Next     *ListTrackingCategoriesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListTrackingCategoriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListTrackingCategoriesLinksLinksSelf      `json:"self"`
}

// ListTrackingCategoriesLinks
// Codat's Paging Model
type ListTrackingCategoriesLinks struct {
	Links        ListTrackingCategoriesLinksLinks `json:"_links"`
	PageNumber   int64                            `json:"pageNumber"`
	PageSize     int64                            `json:"pageSize"`
	Results      []shared.TrackingCategory        `json:"results,omitempty"`
	TotalResults int64                            `json:"totalResults"`
}

type ListTrackingCategoriesResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListTrackingCategoriesLinks
}
