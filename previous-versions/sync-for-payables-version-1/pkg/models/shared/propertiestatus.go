// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PropertieStatus - The current status of the dataset.
type PropertieStatus string

const (
	PropertieStatusInitial            PropertieStatus = "Initial"
	PropertieStatusQueued             PropertieStatus = "Queued"
	PropertieStatusFetching           PropertieStatus = "Fetching"
	PropertieStatusMapQueued          PropertieStatus = "MapQueued"
	PropertieStatusMapping            PropertieStatus = "Mapping"
	PropertieStatusComplete           PropertieStatus = "Complete"
	PropertieStatusFetchError         PropertieStatus = "FetchError"
	PropertieStatusMapError           PropertieStatus = "MapError"
	PropertieStatusInternalError      PropertieStatus = "InternalError"
	PropertieStatusProcessingQueued   PropertieStatus = "ProcessingQueued"
	PropertieStatusProcessing         PropertieStatus = "Processing"
	PropertieStatusProcessingError    PropertieStatus = "ProcessingError"
	PropertieStatusValidationQueued   PropertieStatus = "ValidationQueued"
	PropertieStatusValidating         PropertieStatus = "Validating"
	PropertieStatusValidationError    PropertieStatus = "ValidationError"
	PropertieStatusAuthError          PropertieStatus = "AuthError"
	PropertieStatusCancelled          PropertieStatus = "Cancelled"
	PropertieStatusNotSupported       PropertieStatus = "NotSupported"
	PropertieStatusRateLimitError     PropertieStatus = "RateLimitError"
	PropertieStatusPermissionsError   PropertieStatus = "PermissionsError"
	PropertieStatusPrerequisiteNotMet PropertieStatus = "PrerequisiteNotMet"
)

func (e PropertieStatus) ToPointer() *PropertieStatus {
	return &e
}
func (e *PropertieStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Initial":
		fallthrough
	case "Queued":
		fallthrough
	case "Fetching":
		fallthrough
	case "MapQueued":
		fallthrough
	case "Mapping":
		fallthrough
	case "Complete":
		fallthrough
	case "FetchError":
		fallthrough
	case "MapError":
		fallthrough
	case "InternalError":
		fallthrough
	case "ProcessingQueued":
		fallthrough
	case "Processing":
		fallthrough
	case "ProcessingError":
		fallthrough
	case "ValidationQueued":
		fallthrough
	case "Validating":
		fallthrough
	case "ValidationError":
		fallthrough
	case "AuthError":
		fallthrough
	case "Cancelled":
		fallthrough
	case "NotSupported":
		fallthrough
	case "RateLimitError":
		fallthrough
	case "PermissionsError":
		fallthrough
	case "PrerequisiteNotMet":
		*e = PropertieStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertieStatus: %v", v)
	}
}