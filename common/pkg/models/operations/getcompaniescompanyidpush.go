package operations

import (
	"github.com/codatio/client-sdk-go/common/pkg/models/shared"
)

type GetCompaniesCompanyIDPushPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type GetCompaniesCompanyIDPushQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type GetCompaniesCompanyIDPushRequest struct {
	PathParams  GetCompaniesCompanyIDPushPathParams
	QueryParams GetCompaniesCompanyIDPushQueryParams
}

type GetCompaniesCompanyIDPushLinksLinksCurrent struct {
	Href string `json:"href"`
}

type GetCompaniesCompanyIDPushLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type GetCompaniesCompanyIDPushLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type GetCompaniesCompanyIDPushLinksLinksSelf struct {
	Href string `json:"href"`
}

type GetCompaniesCompanyIDPushLinksLinks struct {
	Current  GetCompaniesCompanyIDPushLinksLinksCurrent   `json:"current"`
	Next     *GetCompaniesCompanyIDPushLinksLinksNext     `json:"next,omitempty"`
	Previous *GetCompaniesCompanyIDPushLinksLinksPrevious `json:"previous,omitempty"`
	Self     GetCompaniesCompanyIDPushLinksLinksSelf      `json:"self"`
}

// GetCompaniesCompanyIDPushLinks
// Codat's Paging Model
type GetCompaniesCompanyIDPushLinks struct {
	Links        GetCompaniesCompanyIDPushLinksLinks                                                                                                         `json:"_links"`
	PageNumber   int64                                                                                                                                       `json:"pageNumber"`
	PageSize     int64                                                                                                                                       `json:"pageSize"`
	Results      []shared.Onecompanies1Percent7BcompanyIDPercent7D1push1Percent7BpushOperationKeyPercent7DGetResponses200ContentApplication1jsonSchemaAllOf1 `json:"results,omitempty"`
	TotalResults int64                                                                                                                                       `json:"totalResults"`
}

type GetCompaniesCompanyIDPushResponse struct {
	ContentType string
	StatusCode  int
	Links       *GetCompaniesCompanyIDPushLinks
}
