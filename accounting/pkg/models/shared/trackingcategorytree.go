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
	Status *StatusEnum `json:"status,omitempty"`
	// A collection of subcategories that are nested beneath this category.
	SubCategories []TrackingCategoryTree `json:"subCategories,omitempty"`
}
