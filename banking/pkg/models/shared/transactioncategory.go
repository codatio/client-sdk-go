package shared

import (
	"time"
)

type TransactionCategoryStatusEnum string

const (
	TransactionCategoryStatusEnumUnknown  TransactionCategoryStatusEnum = "Unknown"
	TransactionCategoryStatusEnumActive   TransactionCategoryStatusEnum = "Active"
	TransactionCategoryStatusEnumArchived TransactionCategoryStatusEnum = "Archived"
)

// TransactionCategory
// The Banking Transaction Categories data type provides a list of hierarchical categories associated with a transaction for greater contextual meaning to transaction activity.
type TransactionCategory struct {
	HasChildren        *bool                          `json:"hasChildren,omitempty"`
	ID                 string                         `json:"id"`
	ModifiedDate       *time.Time                     `json:"modifiedDate,omitempty"`
	Name               string                         `json:"name"`
	ParentID           *string                        `json:"parentId,omitempty"`
	SourceModifiedDate *time.Time                     `json:"sourceModifiedDate,omitempty"`
	Status             *TransactionCategoryStatusEnum `json:"status,omitempty"`
}
