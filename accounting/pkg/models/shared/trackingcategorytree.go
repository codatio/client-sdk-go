package shared

import (
	"time"
)

type TrackingCategoryTreeTrackingCategoryStatusEnum string

const (
	TrackingCategoryTreeTrackingCategoryStatusEnumUnknown  TrackingCategoryTreeTrackingCategoryStatusEnum = "Unknown"
	TrackingCategoryTreeTrackingCategoryStatusEnumActive   TrackingCategoryTreeTrackingCategoryStatusEnum = "Active"
	TrackingCategoryTreeTrackingCategoryStatusEnumArchived TrackingCategoryTreeTrackingCategoryStatusEnum = "Archived"
)

type TrackingCategoryTree struct {
	HasChildren        *bool                                           `json:"hasChildren,omitempty"`
	ID                 *string                                         `json:"id,omitempty"`
	ModifiedDate       *time.Time                                      `json:"modifiedDate,omitempty"`
	Name               *string                                         `json:"name,omitempty"`
	ParentID           *string                                         `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                      `json:"sourceModifiedDate,omitempty"`
	Status             *TrackingCategoryTreeTrackingCategoryStatusEnum `json:"status,omitempty"`
	SubCategories      []TrackingCategoryTree                          `json:"subCategories,omitempty"`
}
