package shared

import (
	"time"
)

type DataIntegrityDetailsMatches struct {
	Amount       *string `json:"amount,omitempty"`
	ConnectionID *string `json:"connectionId,omitempty"`
	Currency     *string `json:"currency,omitempty"`
	Date         *string `json:"date,omitempty"`
	Description  *string `json:"description,omitempty"`
	ID           *string `json:"id,omitempty"`
	Type         *string `json:"type,omitempty"`
}

type DataIntegrityDetails struct {
	Amount       *float64                      `json:"amount,omitempty"`
	ConnectionID *string                       `json:"connectionId,omitempty"`
	Currency     *string                       `json:"currency,omitempty"`
	Date         *time.Time                    `json:"date,omitempty"`
	Description  *string                       `json:"description,omitempty"`
	ID           *string                       `json:"id,omitempty"`
	Matches      []DataIntegrityDetailsMatches `json:"matches,omitempty"`
	Type         *string                       `json:"type,omitempty"`
}
