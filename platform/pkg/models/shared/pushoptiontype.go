// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PushOptionType - The option type.
type PushOptionType string

const (
	PushOptionTypeArray     PushOptionType = "Array"
	PushOptionTypeObject    PushOptionType = "Object"
	PushOptionTypeString    PushOptionType = "String"
	PushOptionTypeNumber    PushOptionType = "Number"
	PushOptionTypeBoolean   PushOptionType = "Boolean"
	PushOptionTypeDateTime  PushOptionType = "DateTime"
	PushOptionTypeFile      PushOptionType = "File"
	PushOptionTypeMultiPart PushOptionType = "MultiPart"
)

func (e PushOptionType) ToPointer() *PushOptionType {
	return &e
}

func (e *PushOptionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Array":
		fallthrough
	case "Object":
		fallthrough
	case "String":
		fallthrough
	case "Number":
		fallthrough
	case "Boolean":
		fallthrough
	case "DateTime":
		fallthrough
	case "File":
		fallthrough
	case "MultiPart":
		*e = PushOptionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PushOptionType: %v", v)
	}
}
