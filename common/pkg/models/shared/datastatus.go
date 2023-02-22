package shared

import (
	"time"
)

// DataStatus
// Describes the state of data in the Codat cache for a company and data type
type DataStatus struct {
	CurrentStatus          string    `json:"currentStatus"`
	DataType               string    `json:"dataType"`
	LastSuccessfulSync     time.Time `json:"lastSuccessfulSync"`
	LatestSuccessfulSyncID *string   `json:"latestSuccessfulSyncId,omitempty"`
	LatestSyncID           *string   `json:"latestSyncId,omitempty"`
}
