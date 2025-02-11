// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// CommerceProductCategory - Product categories are used to classify a group of products together, either by type (e.g. "Furniture"), or sometimes by tax profile.
type CommerceProductCategory struct {
	// A collection of parent product categories implicitly ordered with the immediate parent last in the list.
	AncestorRefs []CommerceRecordRef `json:"ancestorRefs,omitempty"`
	// A boolean indicating whether there are other product categories beneath this one in the hierarchy.
	HasChildren *bool `json:"hasChildren,omitempty"`
	// The unique identifier of the product category
	ID           *string `json:"id,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The name of the product category
	Name               *string `json:"name,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}

func (o *CommerceProductCategory) GetAncestorRefs() []CommerceRecordRef {
	if o == nil {
		return nil
	}
	return o.AncestorRefs
}

func (o *CommerceProductCategory) GetHasChildren() *bool {
	if o == nil {
		return nil
	}
	return o.HasChildren
}

func (o *CommerceProductCategory) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CommerceProductCategory) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *CommerceProductCategory) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CommerceProductCategory) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}
