package operations

import (
	"net/http"
	"time"
)

type ListBankTransactionCategoriesRequest struct {
	CompanyID    string  `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connectionId"`
	OrderBy      *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page         int     `queryParam:"style=form,explode=true,name=page"`
	PageSize     *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query        *string `queryParam:"style=form,explode=true,name=query"`
}

type ListBankTransactionCategoriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBankTransactionCategoriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBankTransactionCategoriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBankTransactionCategoriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBankTransactionCategoriesLinksLinks struct {
	Current  ListBankTransactionCategoriesLinksLinksCurrent   `json:"current"`
	Next     *ListBankTransactionCategoriesLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBankTransactionCategoriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBankTransactionCategoriesLinksLinksSelf      `json:"self"`
}

type ListBankTransactionCategoriesLinksSourceModifiedDateStatusEnum string

const (
	ListBankTransactionCategoriesLinksSourceModifiedDateStatusEnumUnknown  ListBankTransactionCategoriesLinksSourceModifiedDateStatusEnum = "Unknown"
	ListBankTransactionCategoriesLinksSourceModifiedDateStatusEnumActive   ListBankTransactionCategoriesLinksSourceModifiedDateStatusEnum = "Active"
	ListBankTransactionCategoriesLinksSourceModifiedDateStatusEnumArchived ListBankTransactionCategoriesLinksSourceModifiedDateStatusEnum = "Archived"
)

// ListBankTransactionCategoriesLinksSourceModifiedDate
// The Banking Transaction Categories data type provides a list of hierarchical categories associated with a transaction for greater contextual meaning to transaction activity.
type ListBankTransactionCategoriesLinksSourceModifiedDate struct {
	HasChildren        *bool                                                           `json:"hasChildren,omitempty"`
	ID                 string                                                          `json:"id"`
	ModifiedDate       *time.Time                                                      `json:"modifiedDate,omitempty"`
	Name               string                                                          `json:"name"`
	ParentID           *string                                                         `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                                      `json:"sourceModifiedDate,omitempty"`
	Status             *ListBankTransactionCategoriesLinksSourceModifiedDateStatusEnum `json:"status,omitempty"`
}

// ListBankTransactionCategoriesLinks
// Codat's Paging Model
type ListBankTransactionCategoriesLinks struct {
	Links        ListBankTransactionCategoriesLinksLinks               `json:"_links"`
	PageNumber   int64                                                 `json:"pageNumber"`
	PageSize     int64                                                 `json:"pageSize"`
	Results      *ListBankTransactionCategoriesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                                 `json:"totalResults"`
}

type ListBankTransactionCategoriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListBankTransactionCategoriesLinks
}
