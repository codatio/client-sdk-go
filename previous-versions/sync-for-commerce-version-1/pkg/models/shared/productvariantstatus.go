// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ProductVariantStatus - The status of the product variant.
type ProductVariantStatus string

const (
	ProductVariantStatusUnknown     ProductVariantStatus = "Unknown"
	ProductVariantStatusPublished   ProductVariantStatus = "Published"
	ProductVariantStatusUnpublished ProductVariantStatus = "Unpublished"
)

func (e ProductVariantStatus) ToPointer() *ProductVariantStatus {
	return &e
}

func (e *ProductVariantStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unknown":
		fallthrough
	case "Published":
		fallthrough
	case "Unpublished":
		*e = ProductVariantStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductVariantStatus: %v", v)
	}
}
