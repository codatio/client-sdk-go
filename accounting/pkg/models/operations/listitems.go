package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type ListItemsPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
}

type ListItemsQueryParams struct {
	OrderBy  *string  `queryParam:"style=form,explode=true,name=orderBy"`
	Page     float64  `queryParam:"style=form,explode=true,name=page"`
	PageSize *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Query    *string  `queryParam:"style=form,explode=true,name=query"`
}

type ListItemsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListItemsRequest struct {
	PathParams  ListItemsPathParams
	QueryParams ListItemsQueryParams
	Security    ListItemsSecurity
}

type ListItemsLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListItemsLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListItemsLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListItemsLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListItemsLinksLinks struct {
	Current  ListItemsLinksLinksCurrent   `json:"current"`
	Next     *ListItemsLinksLinksNext     `json:"next,omitempty"`
	Previous *ListItemsLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListItemsLinksLinksSelf      `json:"self"`
}

// ListItemsLinksSourceModifiedDateBillItemAccountRef
// Reference of the account to which the item is linked.
type ListItemsLinksSourceModifiedDateBillItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListItemsLinksSourceModifiedDateBillItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type ListItemsLinksSourceModifiedDateBillItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// ListItemsLinksSourceModifiedDateBillItem
// Item details that are only for bills.
type ListItemsLinksSourceModifiedDateBillItem struct {
	AccountRef  *ListItemsLinksSourceModifiedDateBillItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                             `json:"description,omitempty"`
	TaxRateRef  *ListItemsLinksSourceModifiedDateBillItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                            `json:"unitPrice,omitempty"`
}

// ListItemsLinksSourceModifiedDateInvoiceItemAccountRef
// Reference of the account to which the item is linked.
type ListItemsLinksSourceModifiedDateInvoiceItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ListItemsLinksSourceModifiedDateInvoiceItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type ListItemsLinksSourceModifiedDateInvoiceItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// ListItemsLinksSourceModifiedDateInvoiceItem
// Item details that are only for bills.
type ListItemsLinksSourceModifiedDateInvoiceItem struct {
	AccountRef  *ListItemsLinksSourceModifiedDateInvoiceItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                                `json:"description,omitempty"`
	TaxRateRef  *ListItemsLinksSourceModifiedDateInvoiceItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                               `json:"unitPrice,omitempty"`
}

type ListItemsLinksSourceModifiedDateItemStatusEnum string

const (
	ListItemsLinksSourceModifiedDateItemStatusEnumUnknown  ListItemsLinksSourceModifiedDateItemStatusEnum = "Unknown"
	ListItemsLinksSourceModifiedDateItemStatusEnumActive   ListItemsLinksSourceModifiedDateItemStatusEnum = "Active"
	ListItemsLinksSourceModifiedDateItemStatusEnumArchived ListItemsLinksSourceModifiedDateItemStatusEnum = "Archived"
)

type ListItemsLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListItemsLinksSourceModifiedDateTypeEnum string

const (
	ListItemsLinksSourceModifiedDateTypeEnumUnknown      ListItemsLinksSourceModifiedDateTypeEnum = "Unknown"
	ListItemsLinksSourceModifiedDateTypeEnumInventory    ListItemsLinksSourceModifiedDateTypeEnum = "Inventory"
	ListItemsLinksSourceModifiedDateTypeEnumNonInventory ListItemsLinksSourceModifiedDateTypeEnum = "NonInventory"
	ListItemsLinksSourceModifiedDateTypeEnumService      ListItemsLinksSourceModifiedDateTypeEnum = "Service"
)

// ListItemsLinksSourceModifiedDate
// > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
type ListItemsLinksSourceModifiedDate struct {
	BillItem           *ListItemsLinksSourceModifiedDateBillItem      `json:"billItem,omitempty"`
	Code               *string                                        `json:"code,omitempty"`
	ID                 *string                                        `json:"id,omitempty"`
	InvoiceItem        *ListItemsLinksSourceModifiedDateInvoiceItem   `json:"invoiceItem,omitempty"`
	IsBillItem         bool                                           `json:"isBillItem"`
	IsInvoiceItem      bool                                           `json:"isInvoiceItem"`
	ItemStatus         ListItemsLinksSourceModifiedDateItemStatusEnum `json:"itemStatus"`
	Metadata           *ListItemsLinksSourceModifiedDateMetadata      `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                     `json:"modifiedDate,omitempty"`
	Name               *string                                        `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                     `json:"sourceModifiedDate,omitempty"`
	Type               ListItemsLinksSourceModifiedDateTypeEnum       `json:"type"`
}

// ListItemsLinks
// Codat's Paging Model
type ListItemsLinks struct {
	Links        ListItemsLinksLinks                `json:"_links"`
	PageNumber   int64                              `json:"pageNumber"`
	PageSize     int64                              `json:"pageSize"`
	Results      []ListItemsLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                              `json:"totalResults"`
}

type ListItemsResponse struct {
	ContentType string
	StatusCode  int
	Links       *ListItemsLinks
}
