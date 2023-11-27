// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CommerceProduct - A Product is an item in the company's inventory, and includes information about the price and quantity of all products, and variants thereof, available for sale.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-products) for this data type.
type CommerceProduct struct {
	// Retail category that the product is assigned to e.g. `Hardware`.
	Categorization *string `json:"categorization,omitempty"`
	// Description of the product recorded in the commerce or point of sale platform.
	Description *string `json:"description,omitempty"`
	// A unique, persistent identifier for this record
	ID string `json:"id"`
	// Whether the product represents a gift card or voucher that
	// can be redeemed in the commerce or POS platform.
	//
	IsGiftCard *bool `json:"isGiftCard,omitempty"`
	// Name of the product in the commerce or POS system
	Name *string `json:"name,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	Variants         []ProductVariant  `json:"variants,omitempty"`
}

func (o *CommerceProduct) GetCategorization() *string {
	if o == nil {
		return nil
	}
	return o.Categorization
}

func (o *CommerceProduct) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CommerceProduct) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CommerceProduct) GetIsGiftCard() *bool {
	if o == nil {
		return nil
	}
	return o.IsGiftCard
}

func (o *CommerceProduct) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CommerceProduct) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *CommerceProduct) GetVariants() []ProductVariant {
	if o == nil {
		return nil
	}
	return o.Variants
}
