package operations

import (
	"net/http"
	"time"
)

type ListAccountsCategoriesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListAccountsCategoriesQueryParams struct {
	OrderBy  *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page     int     `queryParam:"style=form,explode=true,name=page"`
	PageSize *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string `queryParam:"style=form,explode=true,name=query"`
}

type ListAccountsCategoriesRequest struct {
	PathParams  ListAccountsCategoriesPathParams
	QueryParams ListAccountsCategoriesQueryParams
}

type ListAccountsCategoriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListAccountsCategoriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListAccountsCategoriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListAccountsCategoriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListAccountsCategoriesLinksLinks struct {
	Current  ListAccountsCategoriesLinksLinksCurrent   `json:"current"`
	Next     *ListAccountsCategoriesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListAccountsCategoriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListAccountsCategoriesLinksLinksSelf      `json:"self"`
}

// ListAccountsCategoriesLinksCategorisedAccountAccountRef
// An object containing account reference data.
type ListAccountsCategoriesLinksCategorisedAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ListAccountsCategoriesLinksCategorisedAccountModifiedDate struct {
	DetailType   *string    `json:"detailType,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Subtype      *string    `json:"subtype,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

type ListAccountsCategoriesLinksCategorisedAccount struct {
	AccountRef *ListAccountsCategoriesLinksCategorisedAccountAccountRef   `json:"accountRef,omitempty"`
	Confirmed  *ListAccountsCategoriesLinksCategorisedAccountModifiedDate `json:"confirmed,omitempty"`
	Suggested  *ListAccountsCategoriesLinksCategorisedAccountModifiedDate `json:"suggested,omitempty"`
}

// ListAccountsCategoriesLinks
// Codat's Paging Model
type ListAccountsCategoriesLinks struct {
	Links        ListAccountsCategoriesLinksLinks                `json:"_links"`
	PageNumber   int64                                           `json:"pageNumber"`
	PageSize     int64                                           `json:"pageSize"`
	Results      []ListAccountsCategoriesLinksCategorisedAccount `json:"results,omitempty"`
	TotalResults int64                                           `json:"totalResults"`
}

type ListAccountsCategoriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListAccountsCategoriesLinks
}
