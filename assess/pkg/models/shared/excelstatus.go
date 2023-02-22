package shared

import (
	"time"
)

type ExcelStatus struct {
	ErrorMessage     *string    `json:"errorMessage,omitempty"`
	FileSize         *int64     `json:"fileSize,omitempty"`
	InProgress       *bool      `json:"inProgress,omitempty"`
	LastGenerated    *time.Time `json:"lastGenerated,omitempty"`
	LastInvocationID *string    `json:"lastInvocationId,omitempty"`
	Queued           *string    `json:"queued,omitempty"`
	ReportType       *string    `json:"reportType,omitempty"`
	Success          *bool      `json:"success,omitempty"`
}
