package shared

import (
	"time"
)

// ItemBillItemTaxRateRef
// Reference of the tax rate to which the item is linked.
type ItemBillItemTaxRateRef struct {
	EffectiveTaxRate *float64 `json:"effectiveTaxRate,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Name             *string  `json:"name,omitempty"`
}

// ItemBillItem
// Item details that are only for bills.
type ItemBillItem struct {
	AccountRef  *AccountRef             `json:"accountRef,omitempty"`
	Description *string                 `json:"description,omitempty"`
	TaxRateRef  *ItemBillItemTaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64                `json:"unitPrice,omitempty"`
}

// ItemInvoiceItem
// Item details that are only for bills.
type ItemInvoiceItem struct {
	AccountRef  *AccountRef `json:"accountRef,omitempty"`
	Description *string     `json:"description,omitempty"`
	TaxRateRef  *TaxRateRef `json:"taxRateRef,omitempty"`
	UnitPrice   *float64    `json:"unitPrice,omitempty"`
}

type ItemItemStatusEnum string

const (
	ItemItemStatusEnumUnknown  ItemItemStatusEnum = "Unknown"
	ItemItemStatusEnumActive   ItemItemStatusEnum = "Active"
	ItemItemStatusEnumArchived ItemItemStatusEnum = "Archived"
)

type ItemTypeEnum string

const (
	ItemTypeEnumUnknown      ItemTypeEnum = "Unknown"
	ItemTypeEnumInventory    ItemTypeEnum = "Inventory"
	ItemTypeEnumNonInventory ItemTypeEnum = "NonInventory"
	ItemTypeEnumService      ItemTypeEnum = "Service"
)

// Item
// > View the coverage for items in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=items" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// **Items** allow your customers to save and track details of the products and services that they buy and sell.
type Item struct {
	BillItem           *ItemBillItem      `json:"billItem,omitempty"`
	Code               *string            `json:"code,omitempty"`
	ID                 *string            `json:"id,omitempty"`
	InvoiceItem        *ItemInvoiceItem   `json:"invoiceItem,omitempty"`
	IsBillItem         bool               `json:"isBillItem"`
	IsInvoiceItem      bool               `json:"isInvoiceItem"`
	ItemStatus         ItemItemStatusEnum `json:"itemStatus"`
	Metadata           *Metadata          `json:"metadata,omitempty"`
	ModifiedDate       *time.Time         `json:"modifiedDate,omitempty"`
	Name               *string            `json:"name,omitempty"`
	SourceModifiedDate *time.Time         `json:"sourceModifiedDate,omitempty"`
	Type               ItemTypeEnum       `json:"type"`
}
