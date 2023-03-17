package operations

import (
	"net/http"
	"time"
)

type ListTrackingCategoriesRequest struct {
	CompanyID string  `pathParam:"style=simple,explode=false,name=companyId"`
	OrderBy   *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page      int     `queryParam:"style=form,explode=true,name=page"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query     *string `queryParam:"style=form,explode=true,name=query"`
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

type ListTrackingCategoriesLinksSourceModifiedDateTrackingCategoryStatusEnum string

const (
	ListTrackingCategoriesLinksSourceModifiedDateTrackingCategoryStatusEnumUnknown  ListTrackingCategoriesLinksSourceModifiedDateTrackingCategoryStatusEnum = "Unknown"
	ListTrackingCategoriesLinksSourceModifiedDateTrackingCategoryStatusEnumActive   ListTrackingCategoriesLinksSourceModifiedDateTrackingCategoryStatusEnum = "Active"
	ListTrackingCategoriesLinksSourceModifiedDateTrackingCategoryStatusEnumArchived ListTrackingCategoriesLinksSourceModifiedDateTrackingCategoryStatusEnum = "Archived"
)

// ListTrackingCategoriesLinksSourceModifiedDate
// Details of a category used for tracking transactions.
//
// > Language tip
// >
// > Parameters used to track types of spend in various parts of an organization can be called  **dimensions**, **projects**, **classes**, or **locations** in different accounting platforms. In Codat, we refer to these as tracking categories.
//
// View the coverage for tracking categories in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=trackingCategories" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Tracking categories are used to monitor cost centres and control budgets that sit outside the standard chart of accounts. Customers may use tracking categories to group together and track the income and costs of specific departments, projects, locations or customers.
//
// From their accounting system, customers can:
//
// - Create and maintain tracking categories and tracking category types.
// - View all tracking categories that are available for use.
// - View the relationships between the categories.
// - Assign invoices, bills, credit notes, or bill credit notes to one or more categories.
// - View the categories that a transaction belongs to.
// - View all transactions in a tracking category.
type ListTrackingCategoriesLinksSourceModifiedDate struct {
	HasChildren        *bool                                                                    `json:"hasChildren,omitempty"`
	ID                 *string                                                                  `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                               `json:"modifiedDate,omitempty"`
	Name               *string                                                                  `json:"name,omitempty"`
	ParentID           *string                                                                  `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                                               `json:"sourceModifiedDate,omitempty"`
	Status             *ListTrackingCategoriesLinksSourceModifiedDateTrackingCategoryStatusEnum `json:"status,omitempty"`
}

// ListTrackingCategoriesLinks
// Codat's Paging Model
type ListTrackingCategoriesLinks struct {
	Links        ListTrackingCategoriesLinksLinks                `json:"_links"`
	PageNumber   int64                                           `json:"pageNumber"`
	PageSize     int64                                           `json:"pageSize"`
	Results      []ListTrackingCategoriesLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                           `json:"totalResults"`
}

type ListTrackingCategoriesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListTrackingCategoriesLinks
}
