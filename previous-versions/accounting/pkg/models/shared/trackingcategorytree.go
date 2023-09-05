// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TrackingCategoryTree - The full structure of a specific tracking category including any child or subcategories.
type TrackingCategoryTree struct {
	// Boolean value indicating whether this category has SubCategories
	HasChildren *bool `json:"hasChildren,omitempty"`
	// The identifier for the item, unique per tracking category
	ID           *string `json:"id,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The name of the tracking category
	Name *string `json:"name,omitempty"`
	// The identifier for this item's immediate parent
	ParentID           *string `json:"parentId,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Current state of the tracking category.
	Status *Status `json:"status,omitempty"`
	// A collection of subcategories that are nested beneath this category.
	SubCategories []TrackingCategoryTree `json:"subCategories,omitempty"`
}

func (o *TrackingCategoryTree) GetHasChildren() *bool {
	if o == nil {
		return nil
	}
	return o.HasChildren
}

func (o *TrackingCategoryTree) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TrackingCategoryTree) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *TrackingCategoryTree) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TrackingCategoryTree) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *TrackingCategoryTree) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *TrackingCategoryTree) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TrackingCategoryTree) GetSubCategories() []TrackingCategoryTree {
	if o == nil {
		return nil
	}
	return o.SubCategories
}