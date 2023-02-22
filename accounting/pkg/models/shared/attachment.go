package shared

import (
	"time"
)

type Attachment struct {
	ContentType        *string    `json:"contentType,omitempty"`
	DateCreated        *time.Time `json:"dateCreated,omitempty"`
	FileSize           *int       `json:"fileSize,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	IncludeWhenSent    *bool      `json:"includeWhenSent,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}
