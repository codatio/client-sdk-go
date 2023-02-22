package shared

import (
	"time"
)

// CategorisedAccountAccountRef
// An object containing account reference data.
type CategorisedAccountAccountRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CategorisedAccountModifiedDate struct {
	DetailType   *string    `json:"detailType,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	Subtype      *string    `json:"subtype,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

type CategorisedAccount struct {
	AccountRef *CategorisedAccountAccountRef   `json:"accountRef,omitempty"`
	Confirmed  *CategorisedAccountModifiedDate `json:"confirmed,omitempty"`
	Suggested  *CategorisedAccountModifiedDate `json:"suggested,omitempty"`
}
