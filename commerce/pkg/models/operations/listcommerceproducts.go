package operations

import (
	"net/http"
	"time"
)

type ListCommerceProductsRequest struct {
	CompanyID    string  `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string  `pathParam:"style=simple,explode=false,name=connectionId"`
	OrderBy      *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page         int     `queryParam:"style=form,explode=true,name=page"`
	PageSize     *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query        *string `queryParam:"style=form,explode=true,name=query"`
}

type ListCommerceProductsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListCommerceProductsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceProductsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListCommerceProductsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListCommerceProductsLinksLinks struct {
	Current  ListCommerceProductsLinksLinksCurrent   `json:"current"`
	Next     *ListCommerceProductsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListCommerceProductsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListCommerceProductsLinksLinksSelf      `json:"self"`
}

type ListCommerceProductsLinksProductSourceModifiedDateStatusEnum string

const (
	ListCommerceProductsLinksProductSourceModifiedDateStatusEnumUnknown     ListCommerceProductsLinksProductSourceModifiedDateStatusEnum = "Unknown"
	ListCommerceProductsLinksProductSourceModifiedDateStatusEnumPublished   ListCommerceProductsLinksProductSourceModifiedDateStatusEnum = "Published"
	ListCommerceProductsLinksProductSourceModifiedDateStatusEnumUnpublished ListCommerceProductsLinksProductSourceModifiedDateStatusEnum = "Unpublished"
)

// ListCommerceProductsLinksProductSourceModifiedDate
// Represents a variation of a product available for sale, for example an item of clothing
// may be available for sale in multiple sizes and colors
type ListCommerceProductsLinksProductSourceModifiedDate struct {
	Barcode            *string                                                       `json:"barcode,omitempty"`
	CreatedDate        *time.Time                                                    `json:"createdDate,omitempty"`
	ID                 string                                                        `json:"id"`
	Inventory          []interface{}                                                 `json:"inventory,omitempty"`
	IsTaxEnabled       *bool                                                         `json:"isTaxEnabled,omitempty"`
	ModifiedDate       *time.Time                                                    `json:"modifiedDate,omitempty"`
	Name               *string                                                       `json:"name,omitempty"`
	Prices             []interface{}                                                 `json:"prices,omitempty"`
	Quantity           *float64                                                      `json:"quantity,omitempty"`
	ShippingRequired   *bool                                                         `json:"shippingRequired,omitempty"`
	Sku                *string                                                       `json:"sku,omitempty"`
	SourceModifiedDate *time.Time                                                    `json:"sourceModifiedDate,omitempty"`
	Status             *ListCommerceProductsLinksProductSourceModifiedDateStatusEnum `json:"status,omitempty"`
	UnitOfMeasure      *string                                                       `json:"unitOfMeasure,omitempty"`
	VatPercentage      *float64                                                      `json:"vatPercentage,omitempty"`
}

// ListCommerceProductsLinksProduct
// A Product is an item in the company's inventory, and includes information
// about the price and quantity of all products, and variants thereof, available for sale
type ListCommerceProductsLinksProduct struct {
	Categorization *string                                              `json:"categorization,omitempty"`
	Description    *string                                              `json:"description,omitempty"`
	ID             string                                               `json:"id"`
	IsGiftCard     *bool                                                `json:"isGiftCard,omitempty"`
	Name           *string                                              `json:"name,omitempty"`
	Variants       []ListCommerceProductsLinksProductSourceModifiedDate `json:"variants,omitempty"`
}

// ListCommerceProductsLinks
// Codat's Paging Model
type ListCommerceProductsLinks struct {
	Links        ListCommerceProductsLinksLinks     `json:"_links"`
	PageNumber   int64                              `json:"pageNumber"`
	PageSize     int64                              `json:"pageSize"`
	Results      []ListCommerceProductsLinksProduct `json:"results,omitempty"`
	TotalResults int64                              `json:"totalResults"`
}

type ListCommerceProductsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListCommerceProductsLinks
}
