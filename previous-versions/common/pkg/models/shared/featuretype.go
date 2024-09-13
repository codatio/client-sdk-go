// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// FeatureType - The type of feature.
type FeatureType string

const (
	FeatureTypeGet                FeatureType = "Get"
	FeatureTypePost               FeatureType = "Post"
	FeatureTypeCategorization     FeatureType = "Categorization"
	FeatureTypeDelete             FeatureType = "Delete"
	FeatureTypePut                FeatureType = "Put"
	FeatureTypeGetAsPdf           FeatureType = "GetAsPdf"
	FeatureTypeDownloadAttachment FeatureType = "DownloadAttachment"
	FeatureTypeGetAttachment      FeatureType = "GetAttachment"
	FeatureTypeGetAttachments     FeatureType = "GetAttachments"
	FeatureTypeUploadAttachment   FeatureType = "UploadAttachment"
)

func (e FeatureType) ToPointer() *FeatureType {
	return &e
}
func (e *FeatureType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Get":
		fallthrough
	case "Post":
		fallthrough
	case "Categorization":
		fallthrough
	case "Delete":
		fallthrough
	case "Put":
		fallthrough
	case "GetAsPdf":
		fallthrough
	case "DownloadAttachment":
		fallthrough
	case "GetAttachment":
		fallthrough
	case "GetAttachments":
		fallthrough
	case "UploadAttachment":
		*e = FeatureType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FeatureType: %v", v)
	}
}
