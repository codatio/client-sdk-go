package shared

import (
	"time"
)

type ProductCategoryRecordRef struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// ProductCategory
// Product categories are used to classify a group of products together, either by type (eg "Furniture"), or sometimes by tax profile.
type ProductCategory struct {
	AncestorRefs       []ProductCategoryRecordRef `json:"ancestorRefs,omitempty"`
	HasChildren        *bool                      `json:"hasChildren,omitempty"`
	ID                 *string                    `json:"id,omitempty"`
	ModifiedDate       *time.Time                 `json:"modifiedDate,omitempty"`
	Name               *string                    `json:"name,omitempty"`
	SourceModifiedDate *time.Time                 `json:"sourceModifiedDate,omitempty"`
}
