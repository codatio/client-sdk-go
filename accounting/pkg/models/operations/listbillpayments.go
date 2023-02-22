package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListBillPaymentsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListBillPaymentsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListBillPaymentsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListBillPaymentsRequest struct {
	PathParams  ListBillPaymentsPathParams
	QueryParams ListBillPaymentsQueryParams
	Security    ListBillPaymentsSecurity
}

type ListBillPaymentsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListBillPaymentsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListBillPaymentsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListBillPaymentsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListBillPaymentsLinksLinks struct {
	Current  ListBillPaymentsLinksLinksCurrent   `json:"current"`
	Next     *ListBillPaymentsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListBillPaymentsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListBillPaymentsLinksLinksSelf      `json:"self"`
}

// ListBillPaymentsLinks
// Codat's Paging Model
type ListBillPaymentsLinks struct {
	Links        ListBillPaymentsLinksLinks `json:"_links"`
	PageNumber   int64                      `json:"pageNumber"`
	PageSize     int64                      `json:"pageSize"`
	Results      []shared.BillPayment       `json:"results,omitempty"`
	TotalResults int64                      `json:"totalResults"`
}

type ListBillPaymentsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListBillPaymentsLinks
}
