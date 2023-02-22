package shared

import (
	"time"
)

// SyncSetting
// Describes how often, and how much history, should be fetched for the given data type when a pull operation is queued.
type SyncSetting struct {
	DataType         string     `json:"dataType"`
	FetchOnFirstLink bool       `json:"fetchOnFirstLink"`
	IsLocked         *bool      `json:"isLocked,omitempty"`
	MonthsToSync     *int64     `json:"monthsToSync,omitempty"`
	SyncFromUtc      *time.Time `json:"syncFromUtc,omitempty"`
	SyncFromWindow   *int64     `json:"syncFromWindow,omitempty"`
	SyncOrder        int64      `json:"syncOrder"`
	SyncSchedule     int64      `json:"syncSchedule"`
}
