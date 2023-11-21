// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DataConnectionStatus - The current authorization status of the data connection.
type DataConnectionStatus string

const (
	DataConnectionStatusPendingAuth  DataConnectionStatus = "PendingAuth"
	DataConnectionStatusLinked       DataConnectionStatus = "Linked"
	DataConnectionStatusUnlinked     DataConnectionStatus = "Unlinked"
	DataConnectionStatusDeauthorized DataConnectionStatus = "Deauthorized"
)

func (e DataConnectionStatus) ToPointer() *DataConnectionStatus {
	return &e
}

func (e *DataConnectionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PendingAuth":
		fallthrough
	case "Linked":
		fallthrough
	case "Unlinked":
		fallthrough
	case "Deauthorized":
		*e = DataConnectionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DataConnectionStatus: %v", v)
	}
}
