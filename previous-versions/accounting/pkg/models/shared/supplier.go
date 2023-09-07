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
	Status SupplierStatus `json:"status"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Name of the supplier as recorded in the accounting system, typically the company name.
	SupplierName *string `json:"supplierName,omitempty"`
	// Supplier's company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
}

func (o *Supplier) GetAddresses() []Addressesitems {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *Supplier) GetContactName() *string {
	if o == nil {
		return nil
	}
	return o.ContactName
}

func (o *Supplier) GetDefaultCurrency() *string {
	if o == nil {
		return nil
	}
	return o.DefaultCurrency
}

func (o *Supplier) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *Supplier) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Supplier) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Supplier) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Supplier) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *Supplier) GetRegistrationNumber() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationNumber
}

func (o *Supplier) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *Supplier) GetStatus() SupplierStatus {
	if o == nil {
		return SupplierStatus("")
	}
	return o.Status
}

func (o *Supplier) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *Supplier) GetSupplierName() *string {
	if o == nil {
		return nil
	}
	return o.SupplierName
}

func (o *Supplier) GetTaxNumber() *string {
	if o == nil {
		return nil
	}
	return o.TaxNumber
}
