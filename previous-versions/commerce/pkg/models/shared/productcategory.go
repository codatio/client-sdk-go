// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ProductCategory - Product categories are used to classify a group of products together, either by type (e.g. "Furniture"), or sometimes by tax profile.
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-productCategories) for this data type.
type ProductCategory struct {
	// A collection of parent product categories implicitly ordered with the immediate parent last in the list.
	AncestorRefs []RecordRef `json:"ancestorRefs,omitempty"`
	// A boolean indicating whether there are other product categories beneath this one in the hierarchy.
	HasChildren *bool `json:"hasChildren,omitempty"`
	// The unique identifier of the product category
	ID           *string `json:"id,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The name of the product category
	Name               *string `json:"name,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}

func (o *ProductCategory) GetAncestorRefs() []RecordRef {
	if o == nil {
		return nil
	}
	return o.AncestorRefs
}

func (o *ProductCategory) GetHasChildren() *bool {
	if o == nil {
		return nil
	}
	return o.HasChildren
}

func (o *ProductCategory) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ProductCategory) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *ProductCategory) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ProductCategory) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}
