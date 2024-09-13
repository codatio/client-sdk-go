// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// WriteType - Type of write request.
type WriteType string

const (
	WriteTypeCreate           WriteType = "Create"
	WriteTypeUpdate           WriteType = "Update"
	WriteTypeDelete           WriteType = "Delete"
	WriteTypeUploadAttachment WriteType = "UploadAttachment"
)

func (e WriteType) ToPointer() *WriteType {
	return &e
}
func (e *WriteType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Create":
		fallthrough
	case "Update":
		fallthrough
	case "Delete":
		fallthrough
	case "UploadAttachment":
		*e = WriteType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WriteType: %v", v)
	}
}
