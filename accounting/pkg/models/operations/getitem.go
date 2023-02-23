package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"time"
)

type GetItemPathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	ItemID    string `pathParam:"style=simple,explode=false,name=itemId"`
}

type GetItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetItemRequest struct {
	PathParams GetItemPathParams
	Security   GetItemSecurity
}

// GetItemSourceModifiedDateBillItemAccountRef
// Reference of the account to which the item is linked.
type GetItemSourceModifiedDateBillItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetItemSourceModifiedDateBillItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type GetItemSourceModifiedDateBillItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// GetItemSourceModifiedDateBillItem
// Item details that are only for bills.
type GetItemSourceModifiedDateBillItem struct {
	AccountRef  *GetItemSourceModifiedDateBillItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                      `json:"description,omitempty"`
	TaxRateRef  *GetItemSourceModifiedDateBillItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                     `json:"unitPrice,omitempty"`
}

// GetItemSourceModifiedDateInvoiceItemAccountRef
// Reference of the account to which the item is linked.
type GetItemSourceModifiedDateInvoiceItemAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetItemSourceModifiedDateInvoiceItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type GetItemSourceModifiedDateInvoiceItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// GetItemSourceModifiedDateInvoiceItem
// Item details that are only for bills.
type GetItemSourceModifiedDateInvoiceItem struct {
	AccountRef  *GetItemSourceModifiedDateInvoiceItemAccountRef `json:"accountRef,omitempty"`
	Description *string                                         `json:"description,omitempty"`
	TaxRateRef  *GetItemSourceModifiedDateInvoiceItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                                        `json:"unitPrice,omitempty"`
}

type GetItemSourceModifiedDateItemStatusEnum string

const (
	GetItemSourceModifiedDateItemStatusEnumUnknown  GetItemSourceModifiedDateItemStatusEnum = "Unknown"
	GetItemSourceModifiedDateItemStatusEnumActive   GetItemSourceModifiedDateItemStatusEnum = "Active"
	GetItemSourceModifiedDateItemStatusEnumArchived GetItemSourceModifiedDateItemStatusEnum = "Archived"
)

type GetItemSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetItemSourceModifiedDateTypeEnum string

const (
	GetItemSourceModifiedDateTypeEnumUnknown      GetItemSourceModifiedDateTypeEnum = "Unknown"
	GetItemSourceModifiedDateTypeEnumInventory    GetItemSourceModifiedDateTypeEnum = "Inventory"
	GetItemSourceModifiedDateTypeEnumNonInventory GetItemSourceModifiedDateTypeEnum = "NonInventory"
	GetItemSourceModifiedDateTypeEnumService      GetItemSourceModifiedDateTypeEnum = "Service"
)

// GetItemSourceModifiedDate
// > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
type GetItemSourceModifiedDate struct {
	BillItem           *GetItemSourceModifiedDateBillItem      `json:"billItem,omitempty"`
	Code               *string                                 `json:"code,omitempty"`
	ID                 *string                                 `json:"id,omitempty"`
	InvoiceItem        *GetItemSourceModifiedDateInvoiceItem   `json:"invoiceItem,omitempty"`
	IsBillItem         bool                                    `json:"isBillItem"`
	IsInvoiceItem      bool                                    `json:"isInvoiceItem"`
	ItemStatus         GetItemSourceModifiedDateItemStatusEnum `json:"itemStatus"`
	Metadata           *GetItemSourceModifiedDateMetadata      `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                              `json:"modifiedDate,omitempty"`
	Name               *string                                 `json:"name,omitempty"`
	SourceModifiedDate *time.Time                              `json:"sourceModifiedDate,omitempty"`
	Type               GetItemSourceModifiedDateTypeEnum       `json:"type"`
}

type GetItemResponse struct {
	ContentType        string
	SourceModifiedDate *GetItemSourceModifiedDate
	StatusCode         int
}
