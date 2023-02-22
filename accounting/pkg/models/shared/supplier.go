package shared

import (
	"time"
)

type SupplierStatusEnum string

const (
	SupplierStatusEnumUnknown  SupplierStatusEnum = "Unknown"
	SupplierStatusEnumActive   SupplierStatusEnum = "Active"
	SupplierStatusEnumArchived SupplierStatusEnum = "Archived"
)

// Supplier
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/docs/datamodel-accounting-bills).
type Supplier struct {
	Addresses          []Addressesitems   `json:"addresses,omitempty"`
	ContactName        *string            `json:"contactName,omitempty"`
	DefaultCurrency    *string            `json:"defaultCurrency,omitempty"`
	EmailAddress       *string            `json:"emailAddress,omitempty"`
	ID                 *string            `json:"id,omitempty"`
	Metadata           *Metadata          `json:"metadata,omitempty"`
	ModifiedDate       *time.Time         `json:"modifiedDate,omitempty"`
	Phone              *string            `json:"phone,omitempty"`
	RegistrationNumber *string            `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time         `json:"sourceModifiedDate,omitempty"`
	Status             SupplierStatusEnum `json:"status"`
	SupplementalData   *SupplementalData  `json:"supplementalData,omitempty"`
	SupplierName       *string            `json:"supplierName,omitempty"`
	TaxNumber          *string            `json:"taxNumber,omitempty"`
}
