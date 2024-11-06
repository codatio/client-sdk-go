// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ReportOperationStatus - The status of the report generation.
type ReportOperationStatus string

const (
	ReportOperationStatusInProgress ReportOperationStatus = "InProgress"
	ReportOperationStatusComplete   ReportOperationStatus = "Complete"
	ReportOperationStatusError      ReportOperationStatus = "Error"
)

func (e ReportOperationStatus) ToPointer() *ReportOperationStatus {
	return &e
}
func (e *ReportOperationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "InProgress":
		fallthrough
	case "Complete":
		fallthrough
	case "Error":
		*e = ReportOperationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReportOperationStatus: %v", v)
	}
}

// ReportOperationType - The name of the report generated.
type ReportOperationType string

const (
	ReportOperationTypeCategorizedBankStatement ReportOperationType = "categorizedBankStatement"
)

func (e ReportOperationType) ToPointer() *ReportOperationType {
	return &e
}
func (e *ReportOperationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "categorizedBankStatement":
		*e = ReportOperationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReportOperationType: %v", v)
	}
}

// ReportOperation - Information about a report generation.
type ReportOperation struct {
	// A short message describing any errors that occurred while generating the report.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique identifier of the report
	ID *string `json:"id,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	RequestedDate *string `json:"requestedDate,omitempty"`
	// The status of the report generation.
	Status *ReportOperationStatus `json:"status,omitempty"`
	// The name of the report generated.
	Type *ReportOperationType `json:"type,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	UpdatedDate *string `json:"updatedDate,omitempty"`
}

func (o *ReportOperation) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ReportOperation) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ReportOperation) GetRequestedDate() *string {
	if o == nil {
		return nil
	}
	return o.RequestedDate
}

func (o *ReportOperation) GetStatus() *ReportOperationStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ReportOperation) GetType() *ReportOperationType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ReportOperation) GetUpdatedDate() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedDate
}
