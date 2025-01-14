// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/codatio/client-sdk-go/sync-for-payables/v5/pkg/utils"
	"github.com/ericlagergren/decimal"
)

// Supplier - Suppliers are people or organizations that provide something, such as a product or service. Use the [List suppliers](https://docs.codat.io/sync-for-payables-v2-api#/operations/list-suppliers) endpoint to retrieve a list of all suppliers for a company.
//
// Suppliers' data links to accounts payable [bills](https://docs.codat.io/sync-for-payables-v2-api#/schemas/Bill).
type Supplier struct {
	// Identifier for the supplier, unique to the company in the accounting software.
	ID *string `json:"id,omitempty"`
	// Name of the supplier as recorded in the accounting system, typically the company name.
	SupplierName *string `json:"supplierName,omitempty"`
	// Name of the main contact for the supplier.
	ContactName *string `json:"contactName,omitempty"`
	// Email address that the supplier may be contacted on.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Phone number that the supplier may be contacted on.
	Phone *string `json:"phone,omitempty"`
	// An array of Addresses.
	Addresses []Address `json:"addresses,omitempty"`
	// Status of the supplier.
	Status *SupplierStatus `json:"status,omitempty"`
	// Amount outstanding against the supplier.
	Balance *decimal.Big `decimal:"number" json:"balance,omitempty"`
	// Default currency the supplier's transactional data is recorded in.
	DefaultCurrency    *string `json:"defaultCurrency,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
}

func (s Supplier) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Supplier) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Supplier) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Supplier) GetSupplierName() *string {
	if o == nil {
		return nil
	}
	return o.SupplierName
}

func (o *Supplier) GetContactName() *string {
	if o == nil {
		return nil
	}
	return o.ContactName
}

func (o *Supplier) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *Supplier) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *Supplier) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *Supplier) GetStatus() *SupplierStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Supplier) GetBalance() *decimal.Big {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *Supplier) GetDefaultCurrency() *string {
	if o == nil {
		return nil
	}
	return o.DefaultCurrency
}

func (o *Supplier) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}
