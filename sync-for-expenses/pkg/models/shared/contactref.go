// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v4/pkg/utils"
)

// Type - The type of contact.
type Type string

const (
	TypeSupplier Type = "Supplier"
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
	case "Supplier":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type ContactRef struct {
	// Identifier of supplier or customer.
	ID *string `json:"id,omitempty"`
	// The type of contact.
	Type *Type `default:"Supplier" json:"type"`
}

func (c ContactRef) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ContactRef) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ContactRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ContactRef) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
