// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/codatio/client-sdk-go/bank-feeds/v6/pkg/utils"
)

// Type - The type of routing number.
type Type string

const (
	TypeRtn      Type = "rtn"
	TypeAba      Type = "aba"
	TypeSwift    Type = "swift"
	TypeBsb      Type = "bsb"
	TypeIban     Type = "iban"
	TypeNz2      Type = "nz2"
	TypeTrno     Type = "trno"
	TypeSortcode Type = "sortcode"
	TypeBlz      Type = "blz"
	TypeIfsc     Type = "ifsc"
	TypeBankcode Type = "bankcode"
	TypeApca     Type = "apca"
	TypeClabe    Type = "clabe"
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
	case "rtn":
		fallthrough
	case "aba":
		fallthrough
	case "swift":
		fallthrough
	case "bsb":
		fallthrough
	case "iban":
		fallthrough
	case "nz2":
		fallthrough
	case "trno":
		fallthrough
	case "sortcode":
		fallthrough
	case "blz":
		fallthrough
	case "ifsc":
		fallthrough
	case "bankcode":
		fallthrough
	case "apca":
		fallthrough
	case "clabe":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// RoutingInfo - Routing information for the bank. This does not include account number.
type RoutingInfo struct {
	// The numeric identifier of the routing number
	BankCode *string `json:"bankCode,omitempty"`
	// The type of routing number.
	Type *Type `default:"bankcode" json:"type"`
}

func (r RoutingInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RoutingInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RoutingInfo) GetBankCode() *string {
	if o == nil {
		return nil
	}
	return o.BankCode
}

func (o *RoutingInfo) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
