// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Type - The type of the weblink.
type Type string

const (
	TypeWebsite Type = "Website"
	TypeSocial  Type = "Social"
	TypeUnknown Type = "Unknown"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Website":
		fallthrough
	case "Social":
		fallthrough
	case "Unknown":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// WebLinkItems - Weblink associated with the company.
type WebLinkItems struct {
	// The type of the weblink.
	Type *Type `json:"type,omitempty"`
	// The full URL for the weblink.
	URL *string `json:"url,omitempty"`
}

func (o *WebLinkItems) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *WebLinkItems) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
