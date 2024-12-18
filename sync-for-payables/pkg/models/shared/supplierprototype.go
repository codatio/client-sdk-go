// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

type SupplierPrototype struct {
	// Name of the supplier as recorded in the accounting system, typically the company name.
	SupplierName string `json:"supplierName"`
	// Name of the main contact for the supplier.
	ContactName *string `json:"contactName,omitempty"`
	// Email address that the supplier may be contacted on.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Phone number that the supplier may be contacted on.
	Phone *string `json:"phone,omitempty"`
	// An array of Addresses.
	Addresses []Address `json:"addresses,omitempty"`
	// Status of the supplier.
	Status SupplierStatus `json:"status"`
	// Amount outstanding against the supplier.
	Balance *decimal.Big `decimal:"number" json:"balance,omitempty"`
	// Default currency the supplier's transactional data is recorded in.
	DefaultCurrency *string `json:"defaultCurrency,omitempty"`
}

func (s SupplierPrototype) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SupplierPrototype) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SupplierPrototype) GetSupplierName() string {
	if o == nil {
		return ""
	}
	return o.SupplierName
}

func (o *SupplierPrototype) GetContactName() *string {
	if o == nil {
		return nil
	}
	return o.ContactName
}

func (o *SupplierPrototype) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *SupplierPrototype) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *SupplierPrototype) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *SupplierPrototype) GetStatus() SupplierStatus {
	if o == nil {
		return SupplierStatus("")
	}
	return o.Status
}

func (o *SupplierPrototype) GetBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *SupplierPrototype) GetDefaultCurrency() *string {
	if o == nil {
		return nil
	}
	return o.DefaultCurrency
}
