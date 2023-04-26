// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type FeatureStateEnum string

const (
	FeatureStateEnumRelease        FeatureStateEnum = "Release"
	FeatureStateEnumBeta           FeatureStateEnum = "Beta"
	FeatureStateEnumDeprecated     FeatureStateEnum = "Deprecated"
	FeatureStateEnumNotSupported   FeatureStateEnum = "NotSupported"
	FeatureStateEnumNotImplemented FeatureStateEnum = "NotImplemented"
)

func (e FeatureStateEnum) ToPointer() *FeatureStateEnum {
	return &e
}

func (e *FeatureStateEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Release":
		fallthrough
	case "Beta":
		fallthrough
	case "Deprecated":
		fallthrough
	case "NotSupported":
		fallthrough
	case "NotImplemented":
		*e = FeatureStateEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for FeatureStateEnum: %s", s)
	}
}
