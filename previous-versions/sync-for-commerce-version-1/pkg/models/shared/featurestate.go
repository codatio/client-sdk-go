// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// FeatureState - The current release state of the feature.
type FeatureState string

const (
	FeatureStateRelease        FeatureState = "Release"
	FeatureStateAlpha          FeatureState = "Alpha"
	FeatureStateBeta           FeatureState = "Beta"
	FeatureStateDeprecated     FeatureState = "Deprecated"
	FeatureStateNotSupported   FeatureState = "NotSupported"
	FeatureStateNotImplemented FeatureState = "NotImplemented"
)

func (e FeatureState) ToPointer() *FeatureState {
	return &e
}
func (e *FeatureState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Release":
		fallthrough
	case "Alpha":
		fallthrough
	case "Beta":
		fallthrough
	case "Deprecated":
		fallthrough
	case "NotSupported":
		fallthrough
	case "NotImplemented":
		*e = FeatureState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FeatureState: %v", v)
	}
}
