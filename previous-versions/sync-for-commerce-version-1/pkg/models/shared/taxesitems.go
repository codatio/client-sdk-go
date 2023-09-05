// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ItemsTaxComponentRef - Taxes rates reference object depending on the rates being available on source commerce package.
type ItemsTaxComponentRef struct {
	// The unique identitifer of the tax component being referenced.
	ID string `json:"id"`
	// Name of the tax component being referenced.
	Name string `json:"name"`
}

func (o *ItemsTaxComponentRef) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ItemsTaxComponentRef) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type Taxesitems struct {
	// Tax amount on order line sale as available from source commerce platform.
	Rate *float64 `json:"rate,omitempty"`
	// Taxes rates reference object depending on the rates being available on source commerce package.
	TaxComponentRef *ItemsTaxComponentRef `json:"taxComponentRef,omitempty"`
}

func (o *Taxesitems) GetRate() *float64 {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *Taxesitems) GetTaxComponentRef() *ItemsTaxComponentRef {
	if o == nil {
		return nil
	}
	return o.TaxComponentRef
}