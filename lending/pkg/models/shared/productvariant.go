// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/lending/v8/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// ProductVariant - Represents a variation of a product available for sale, for example an item of clothing that may be available for sale in multiple sizes and colors.
type ProductVariant struct {
	// Unique product number of the variant. This might be a barcode, UPC, ISBN, etc.
	Barcode *string `json:"barcode,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	CreatedDate *string `json:"createdDate,omitempty"`
	// A unique, persistent identifier for this record
	ID string `json:"id"`
	// Information about the total inventory as well as the locations inventory is in.
	Inventory *ProductInventory `json:"inventory,omitempty"`
	// Whether sales taxes are enabled for this product variant.
	IsTaxEnabled *bool   `json:"isTaxEnabled,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of the product recorded in the commerce or point of sale platform.
	Name *string `json:"name,omitempty"`
	// Prices for the product variants in different currencies.
	Prices []ProductPrice `json:"prices,omitempty"`
	// Indicates whether or not the product requires physical delivery.
	ShippingRequired *bool `json:"shippingRequired,omitempty"`
	// SKU (stock keeping unit) of the variant, as defined by the merchant.
	Sku                *string `json:"sku,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// The status of the product variant.
	Status *ProductVariantStatus `json:"status,omitempty"`
	// Unit of measure for the variant, such as `kg` or `meters`.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// VAT rate for the product variant if sales taxes are enabled.
	VatPercentage *decimal.Big `decimal:"number" json:"vatPercentage,omitempty"`
}

func (p ProductVariant) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductVariant) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProductVariant) GetBarcode() *string {
	if o == nil {
		return nil
	}
	return o.Barcode
}

func (o *ProductVariant) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *ProductVariant) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ProductVariant) GetInventory() *ProductInventory {
	if o == nil {
		return nil
	}
	return o.Inventory
}

func (o *ProductVariant) GetIsTaxEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.IsTaxEnabled
}

func (o *ProductVariant) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *ProductVariant) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ProductVariant) GetPrices() []ProductPrice {
	if o == nil {
		return nil
	}
	return o.Prices
}

func (o *ProductVariant) GetShippingRequired() *bool {
	if o == nil {
		return nil
	}
	return o.ShippingRequired
}

func (o *ProductVariant) GetSku() *string {
	if o == nil {
		return nil
	}
	return o.Sku
}

func (o *ProductVariant) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *ProductVariant) GetStatus() *ProductVariantStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ProductVariant) GetUnitOfMeasure() *string {
	if o == nil {
		return nil
	}
	return o.UnitOfMeasure
}

func (o *ProductVariant) GetVatPercentage() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.VatPercentage
}
