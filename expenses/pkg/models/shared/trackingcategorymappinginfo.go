package shared

import (
	"time"
)

type TrackingCategoryMappingInfo struct {
	HasChildren  *bool      `json:"hasChildren,omitempty"`
	ID           *string    `json:"id,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Name         *string    `json:"name,omitempty"`
	ParentID     *string    `json:"parentId,omitempty"`
}
