package shared

import (
	"time"
)

type TaxRateComponents struct {
	IsCompound bool     `json:"isCompound"`
	Name       *string  `json:"name,omitempty"`
	Rate       *float64 `json:"rate,omitempty"`
}

type TaxRateStatusEnum string

const (
	TaxRateStatusEnumUnknown  TaxRateStatusEnum = "Unknown"
	TaxRateStatusEnumActive   TaxRateStatusEnum = "Active"
	TaxRateStatusEnumArchived TaxRateStatusEnum = "Archived"
)

// TaxRate
// > View the coverage for tax rates in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=taxRates" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// Accounting systems typically store a set of taxes and associated rates within the accounting package. This means that users don't have to look up or remember the rates for each type of tax. For example: Applying the tax "UK sales VAT" to line items of an invoice adds the correct rate of 20%.
//
// ### Tax components
//
// In some cases, a tax is made up of multiple sub taxes, often called _components_ of the tax.  For example: You may have an item that is charged a tax rate called "City import tax (8%)" that has two components:
//
// - A city tax of 5%.
// - An import tax of 3%.
//
// > **Effective tax rates**
// > Where there are multiple components of a tax, each component may be calculated on the original amount and added together. Alternatively, one tax may be calculated on the sub-total of the original amount plus another tax, which is referred to as _compounding_. When there is compounding, the effective tax rate is the rate that, if applied to the original amount, would result in the total amount of tax with compounding.
// >
// > **Example:**
// > A tax has two components. Both components have a rate of 10%, and one component is compound. In this case, there is a total tax rate of 20% but an effective tax rate of 21%. [Also see _Compound tax example_](#section-compound-tax-example).
// > - For QuickBooks Online, Codat doesn't use compound rates. Instead, the calculated effective tax rate for each component is shown. This means that the effective and total rates are the same because the total tax rate is a sum of the component rates.
type TaxRate struct {
	Code               *string                   `json:"code,omitempty"`
	Components         []TaxRateComponents       `json:"components,omitempty"`
	EffectiveTaxRate   *float64                  `json:"effectiveTaxRate,omitempty"`
	ID                 *string                   `json:"id,omitempty"`
	Metadata           *Metadata                 `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                `json:"modifiedDate,omitempty"`
	Name               *string                   `json:"name,omitempty"`
	SourceModifiedDate *time.Time                `json:"sourceModifiedDate,omitempty"`
	Status             *TaxRateStatusEnum        `json:"status,omitempty"`
	TotalTaxRate       *float64                  `json:"totalTaxRate,omitempty"`
	ValidDatatypeLinks []ValidDatatypeLinksitems `json:"validDatatypeLinks,omitempty"`
}
