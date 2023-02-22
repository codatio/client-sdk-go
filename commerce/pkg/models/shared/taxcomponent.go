package shared

import (
	"time"
)

type TaxComponent struct {
	ID                 string     `json:"id"`
	IsCompound         *bool      `json:"isCompound,omitempty"`
	ModifiedDate       *time.Time `json:"modifiedDate,omitempty"`
	Name               string     `json:"name"`
	Rate               *float32   `json:"rate,omitempty"`
	SourceModifiedDate *time.Time `json:"sourceModifiedDate,omitempty"`
}
