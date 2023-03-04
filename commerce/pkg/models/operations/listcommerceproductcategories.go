package operations

import (
	"net/http"
	"time"
)

type ListCommerceProductCategoriesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type ListCommerceProductCategoriesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceProductCategoriesRequest struct {
	PathParams  ListCommerceProductCategoriesPathParams
	QueryParams ListCommerceProductCategoriesQueryParams
}

type ListCommerceProductCategoriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceProductCategoriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceProductCategoriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceProductCategoriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceProductCategoriesLinksLinks struct {
	Current  ListCommerceProductCategoriesLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceProductCategoriesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceProductCategoriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceProductCategoriesLinksLinksSelf      `json:"self"`
}

type ListCommerceProductCategoriesLinksProductCategoryRecordRef struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// ListCommerceProductCategoriesLinksProductCategory
// Product categories are used to classify a group of products together, either by type (eg "Furniture"), or sometimes by tax profile.
type ListCommerceProductCategoriesLinksProductCategory struct {
	AncestorRefs       []ListCommerceProductCategoriesLinksProductCategoryRecordRef `json:"ancestorRefs,omitempty"`
	HasChildren        *bool                                                        `json:"hasChildren,omitempty"`
	ID                 *string                                                      `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                   `json:"modifiedDate,omitempty"`
	Name               *string                                                      `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                                   `json:"sourceModifiedDate,omitempty"`
}

// ListCommerceProductCategoriesLinks
// Codat's Paging Model
type ListCommerceProductCategoriesLinks struct {
	Links        ListCommerceProductCategoriesLinksLinks             `json:"_links"`
	PageNumber   int64                                               `json:"pageNumber"`
	PageSize     int64                                               `json:"pageSize"`
	Results      []ListCommerceProductCategoriesLinksProductCategory `json:"results,omitempty"`
	TotalResults int64                                               `json:"totalResults"`
}

type ListCommerceProductCategoriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListCommerceProductCategoriesLinks
}
