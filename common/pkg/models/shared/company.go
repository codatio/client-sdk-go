package shared

import (
	"time"
)

// Company
// A company in Codat represent a small or medium sized business, whose data you wish to share
type Company struct {
	Created           *time.Time   `json:"created,omitempty"`
	CreatedByUserName *string      `json:"createdByUserName,omitempty"`
	DataConnections   []Connection `json:"dataConnections,omitempty"`
	Description       *string      `json:"description,omitempty"`
	ID                string       `json:"id"`
	LastSync          *time.Time   `json:"lastSync,omitempty"`
	Name              string       `json:"name"`
	Platform          *string      `json:"platform,omitempty"`
	Redirect          string       `json:"redirect"`
}
