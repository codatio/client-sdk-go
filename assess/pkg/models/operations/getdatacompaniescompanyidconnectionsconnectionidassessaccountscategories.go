package operations

import (
	"time"
)

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesPathParams struct {
	CompanyID    string `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesRequest struct {
	PathParams  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesPathParams
	QueryParams GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesQueryParams
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinks struct {
	Current  GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinksCurrent   `json:"current"`
	Next     *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinksNext     `json:"next,omitempty"`
	Previous *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinksSelf      `json:"self"`
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksCategorisedAccountAccountRef
// An object containing account reference data.
type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksCategorisedAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksCategorisedAccountModifiedDate struct {
	DetailType   *string    `json:"detailType,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Subtype      *string    `json:"subtype,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksCategorisedAccount struct {
	AccountRef *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksCategorisedAccountAccountRef   `json:"accountRef,omitempty"`
	Confirmed  *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksCategorisedAccountModifiedDate `json:"confirmed,omitempty"`
	Suggested  *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksCategorisedAccountModifiedDate `json:"suggested,omitempty"`
}

// GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinks
// Codat's Paging Model
type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinks struct {
	Links        GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksLinks                `json:"_links"`
	PageNumber   int64                                                                                             `json:"pageNumber"`
	PageSize     int64                                                                                             `json:"pageSize"`
	Results      []GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinksCategorisedAccount `json:"results,omitempty"`
	TotalResults int64                                                                                             `json:"totalResults"`
}

type GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesResponse struct {
	ContentType string
	StatusCode  int
	Links       *GetDataCompaniesCompanyIDConnectionsConnectionIDAssessAccountsCategoriesLinks
}
