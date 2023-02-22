package shared

import (
	"time"
)

type ProductSourceModifiedDateStatusEnum string

const (
	ProductSourceModifiedDateStatusEnumUnknown     ProductSourceModifiedDateStatusEnum = "Unknown"
	ProductSourceModifiedDateStatusEnumPublished   ProductSourceModifiedDateStatusEnum = "Published"
	ProductSourceModifiedDateStatusEnumUnpublished ProductSourceModifiedDateStatusEnum = "Unpublished"
)

// ProductSourceModifiedDate
// Represents a variation of a product available for sale, for example an item of clothing
// may be available for sale in multiple sizes and colors
type ProductSourceModifiedDate struct {
	Barcode            *string                              `json:"barcode,omitempty"`
	CreatedDate        *time.Time                           `json:"createdDate,omitempty"`
	ID                 string                               `json:"id"`
	Inventory          []interface{}                        `json:"inventory,omitempty"`
	IsTaxEnabled       *bool                                `json:"isTaxEnabled,omitempty"`
	ModifiedDate       *time.Time                           `json:"modifiedDate,omitempty"`
	Name               *string                              `json:"name,omitempty"`
	Prices             []interface{}                        `json:"prices,omitempty"`
	Quantity           *float64                             `json:"quantity,omitempty"`
	ShippingRequired   *bool                                `json:"shippingRequired,omitempty"`
	Sku                *string                              `json:"sku,omitempty"`
	SourceModifiedDate *time.Time                           `json:"sourceModifiedDate,omitempty"`
	Status             *ProductSourceModifiedDateStatusEnum `json:"status,omitempty"`
	UnitOfMeasure      *string                              `json:"unitOfMeasure,omitempty"`
	VatPercentage      *float64                             `json:"vatPercentage,omitempty"`
}

// Product
// A Product is an item in the company's inventory, and includes information
// about the price and quantity of all products, and variants thereof, available for sale
type Product struct {
	Categorization *string                     `json:"categorization,omitempty"`
	Description    *string                     `json:"description,omitempty"`
	ID             string                      `json:"id"`
	IsGiftCard     *bool                       `json:"isGiftCard,omitempty"`
	Name           *string                     `json:"name,omitempty"`
	Variants       []ProductSourceModifiedDate `json:"variants,omitempty"`
}
