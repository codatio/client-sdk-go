package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
)

type ListPurchaseOrdersPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListPurchaseOrdersQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListPurchaseOrdersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListPurchaseOrdersRequest struct {
	PathParams  ListPurchaseOrdersPathParams
	QueryParams ListPurchaseOrdersQueryParams
	Security    ListPurchaseOrdersSecurity
}

type ListPurchaseOrdersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListPurchaseOrdersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListPurchaseOrdersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListPurchaseOrdersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListPurchaseOrdersLinksLinks struct {
	Current  ListPurchaseOrdersLinksLinksCurrent   `json:"current"`
	Next     *ListPurchaseOrdersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListPurchaseOrdersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListPurchaseOrdersLinksLinksSelf      `json:"self"`
}

// ListPurchaseOrdersLinks
// Codat's Paging Model
type ListPurchaseOrdersLinks struct {
	Links        ListPurchaseOrdersLinksLinks `json:"_links"`
	PageNumber   int64                        `json:"pageNumber"`
	PageSize     int64                        `json:"pageSize"`
	Results      []shared.PurchaseOrder       `json:"results,omitempty"`
	TotalResults int64                        `json:"totalResults"`
}

type ListPurchaseOrdersResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListPurchaseOrdersLinks
}
