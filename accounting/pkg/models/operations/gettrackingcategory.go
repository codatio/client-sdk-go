package operations

import (
	"net/http"
	"time"
)

type GetTrackingCategoryRequest struct {
	CompanyID          string `pathParam:"style=simple,explode=false,name=companyId"`
	TrackingCategoryID string `pathParam:"style=simple,explode=false,name=trackingCategoryId"`
}

type GetTrackingCategorySourceModifiedDateTrackingCategoryStatusEnum string

const (
	GetTrackingCategorySourceModifiedDateTrackingCategoryStatusEnumUnknown  GetTrackingCategorySourceModifiedDateTrackingCategoryStatusEnum = "Unknown"
	GetTrackingCategorySourceModifiedDateTrackingCategoryStatusEnumActive   GetTrackingCategorySourceModifiedDateTrackingCategoryStatusEnum = "Active"
	GetTrackingCategorySourceModifiedDateTrackingCategoryStatusEnumArchived GetTrackingCategorySourceModifiedDateTrackingCategoryStatusEnum = "Archived"
)

type GetTrackingCategorySourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum string

const (
	GetTrackingCategorySourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumUnknown  GetTrackingCategorySourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Unknown"
	GetTrackingCategorySourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumActive   GetTrackingCategorySourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Active"
	GetTrackingCategorySourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumArchived GetTrackingCategorySourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Archived"
)

type GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum string

const (
	GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumUnknown  GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Unknown"
	GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumActive   GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Active"
	GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumArchived GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Archived"
)

type GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum string

const (
	GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumUnknown  GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Unknown"
	GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumActive   GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Active"
	GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnumArchived GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum = "Archived"
)

type GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDate struct {
	ID                 *string                                                                                                                `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                                                                             `json:"modifiedDate,omitempty"`
	Name               *string                                                                                                                `json:"name,omitempty"`
	ParentID           *string                                                                                                                `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                                                                                             `json:"sourceModifiedDate,omitempty"`
	Status             *GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum `json:"status,omitempty"`
}

type GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDate struct {
	HasChildren        *bool                                                                                                `json:"hasChildren,omitempty"`
	ID                 *string                                                                                              `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                                                           `json:"modifiedDate,omitempty"`
	Name               *string                                                                                              `json:"name,omitempty"`
	ParentID           *string                                                                                              `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                                                                           `json:"sourceModifiedDate,omitempty"`
	Status             *GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum `json:"status,omitempty"`
	SubCategories      []GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDateSourceModifiedDate        `json:"subCategories,omitempty"`
}

type GetTrackingCategorySourceModifiedDateSourceModifiedDate struct {
	HasChildren        *bool                                                                              `json:"hasChildren,omitempty"`
	ID                 *string                                                                            `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                                         `json:"modifiedDate,omitempty"`
	Name               *string                                                                            `json:"name,omitempty"`
	ParentID           *string                                                                            `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                                                         `json:"sourceModifiedDate,omitempty"`
	Status             *GetTrackingCategorySourceModifiedDateSourceModifiedDateTrackingCategoryStatusEnum `json:"status,omitempty"`
	SubCategories      []GetTrackingCategorySourceModifiedDateSourceModifiedDateSourceModifiedDate        `json:"subCategories,omitempty"`
}

type GetTrackingCategorySourceModifiedDate struct {
	HasChildren        *bool                                                            `json:"hasChildren,omitempty"`
	ID                 *string                                                          `json:"id,omitempty"`
	ModifiedDate       *time.Time                                                       `json:"modifiedDate,omitempty"`
	Name               *string                                                          `json:"name,omitempty"`
	ParentID           *string                                                          `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                                                       `json:"sourceModifiedDate,omitempty"`
	Status             *GetTrackingCategorySourceModifiedDateTrackingCategoryStatusEnum `json:"status,omitempty"`
	SubCategories      []GetTrackingCategorySourceModifiedDateSourceModifiedDate        `json:"subCategories,omitempty"`
}

type GetTrackingCategoryResponse struct {
	ContentType        string
	SourceModifiedDate *GetTrackingCategorySourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
