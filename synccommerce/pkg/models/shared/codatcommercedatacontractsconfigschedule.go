package shared

import (
	"time"
)

type CodatCommerceDataContractsConfigSchedule struct {
	FrequencyOptions  []string   `json:"frequencyOptions,omitempty"`
	SelectedFrequency *string    `json:"selectedFrequency,omitempty"`
	StartDate         *time.Time `json:"startDate,omitempty"`
	SyncHourUtc       *int       `json:"syncHourUtc,omitempty"`
	TimeZone          *string    `json:"timeZone,omitempty"`
}
