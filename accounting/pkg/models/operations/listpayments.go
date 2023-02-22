package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListPaymentsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListPaymentsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListPaymentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListPaymentsRequest struct {
	PathParams  ListPaymentsPathParams
	QueryParams ListPaymentsQueryParams
	Security    ListPaymentsSecurity
}

type ListPaymentsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListPaymentsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListPaymentsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListPaymentsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListPaymentsLinksLinks struct {
	Current  ListPaymentsLinksLinksCurrent   `json:"current"`
	Next     *ListPaymentsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListPaymentsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListPaymentsLinksLinksSelf      `json:"self"`
}

// ListPaymentsLinks
// Codat's Paging Model
type ListPaymentsLinks struct {
	Links        ListPaymentsLinksLinks `json:"_links"`
	PageNumber   int64                  `json:"pageNumber"`
	PageSize     int64                  `json:"pageSize"`
	Results      []shared.Payment       `json:"results,omitempty"`
	TotalResults int64                  `json:"totalResults"`
}

type ListPaymentsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListPaymentsLinks
}
