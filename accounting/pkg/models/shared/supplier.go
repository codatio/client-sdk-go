// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Supplier - > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://docs.codat.io/accounting-api#/operations/list-suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type Supplier struct {
	// An array of Addresses.
	Addresses []Addressesitems `json:"addresses,omitempty"`
	// Name of the main contact for the supplier.
	ContactName *string `json:"contactName,omitempty"`
	// Default currency the supplier's transactional data is recorded in.
	DefaultCurrency *string `json:"defaultCurrency,omitempty"`
	// Email address that the supplier may be contacted on.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Identifier for the supplier, unique to the company in the accounting platform.
	ID           *string   `json:"id,omitempty"`
	Metadata     *Metadata `json:"metadata,omitempty"`
	ModifiedDate *string   `json:"modifiedDate,omitempty"`
	// Phone number that the supplier may be contacted on.
	Phone *string `json:"phone,omitempty"`
	// Company number of the supplier. In the UK, this is typically the company registration number issued by Companies House.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the supplier.
	Status SupplierStatusEnum `json:"status"`
	// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Name of the supplier as recorded in the accounting system, typically the company name.
	SupplierName *string `json:"supplierName,omitempty"`
	// Supplier's company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
}
