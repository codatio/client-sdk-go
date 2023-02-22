package shared

import (
	"time"
)

// DataIntegrityStatusAmounts
// Only returned for transactions. For accounts, there is nothing returned.
type DataIntegrityStatusAmounts struct {
	Currency *string  `json:"currency,omitempty"`
	Max      *float64 `json:"max,omitempty"`
	Min      *float64 `json:"min,omitempty"`
}

type DataIntegrityStatusConnectionIds struct {
	Source []string `json:"source,omitempty"`
	Target []string `json:"target,omitempty"`
}

// DataIntegrityStatusDates
// Only returned for transactions. For accounts, there is nothing returned.
type DataIntegrityStatusDates struct {
	MaxDate            *time.Time `json:"maxDate,omitempty"`
	MaxOverlappingDate *time.Time `json:"maxOverlappingDate,omitempty"`
	MinDate            *time.Time `json:"minDate,omitempty"`
	MinOverlappingDate *time.Time `json:"minOverlappingDate,omitempty"`
}

type DataIntegrityStatusStatusInfoCurrentStatusEnum string

const (
	DataIntegrityStatusStatusInfoCurrentStatusEnumUnknown      DataIntegrityStatusStatusInfoCurrentStatusEnum = "Unknown"
	DataIntegrityStatusStatusInfoCurrentStatusEnumDoesNotExist DataIntegrityStatusStatusInfoCurrentStatusEnum = "DoesNotExist"
	DataIntegrityStatusStatusInfoCurrentStatusEnumError        DataIntegrityStatusStatusInfoCurrentStatusEnum = "Error"
	DataIntegrityStatusStatusInfoCurrentStatusEnumComplete     DataIntegrityStatusStatusInfoCurrentStatusEnum = "Complete"
)

type DataIntegrityStatusStatusInfo struct {
	CurrentStatus *DataIntegrityStatusStatusInfoCurrentStatusEnum `json:"currentStatus,omitempty"`
	LastMatched   *time.Time                                      `json:"lastMatched,omitempty"`
	StatusMessage *string                                         `json:"statusMessage,omitempty"`
}

type DataIntegrityStatus struct {
	Amounts       *DataIntegrityStatusAmounts       `json:"amounts,omitempty"`
	ConnectionIds *DataIntegrityStatusConnectionIds `json:"connectionIds,omitempty"`
	Dates         *DataIntegrityStatusDates         `json:"dates,omitempty"`
	StatusInfo    *DataIntegrityStatusStatusInfo    `json:"statusInfo,omitempty"`
	Type          *string                           `json:"type,omitempty"`
}
