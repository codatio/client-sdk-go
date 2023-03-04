package operations

import (
	"net/http"
	"time"
)

type GetTaxRatePathParams struct {
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	TaxRateID string `pathParam:"style=simple,explode=false,name=taxRateId"`
}

type GetTaxRateRequest struct {
	PathParams GetTaxRatePathParams
}

type GetTaxRate200ApplicationJSONComponents struct {
	IsCompound bool     `json:"isCompound"`
	Name       *string  `json:"name,omitempty"`
	Rate       *float64 `json:"rate,omitempty"`
}

type GetTaxRate200ApplicationJSONMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetTaxRate200ApplicationJSONStatusEnum string

const (
	GetTaxRate200ApplicationJSONStatusEnumUnknown  GetTaxRate200ApplicationJSONStatusEnum = "Unknown"
	GetTaxRate200ApplicationJSONStatusEnumActive   GetTaxRate200ApplicationJSONStatusEnum = "Active"
	GetTaxRate200ApplicationJSONStatusEnumArchived GetTaxRate200ApplicationJSONStatusEnum = "Archived"
)

// GetTaxRate200ApplicationJSONValidDataTypeLinks
// When querying Codat's data model, some data types return `validDatatypeLinks` metadata in the JSON response. This indicates where that object can be used as a reference—a _valid link_—when creating or updating other data.
//
// For example, `validDatatypeLinks` might indicate the following references:
//
// - Which tax rates are valid to use on the line item of a bill.
// - Which items can be used when creating an invoice.
//
// You can use `validDatatypeLinks` to present your SMB customers with only valid choices when selecting objects from a list, for example.
//
// ## `validDatatypeLinks` example
//
// The following example uses the `Accounting.Accounts` data type. It shows that, on the linked integration, this account is valid as the account on a payment or bill payment; and as the account referenced on the line item of a direct income or direct cost. Because there is no valid link to Invoices or Bills, using this account on those data types will result in an error.
//
// ```json validDatatypeLinks for an account
//
//	{
//	            "id": "bd9e85e0-0478-433d-ae9f-0b3c4f04bfe4",
//	            "nominalCode": "090",
//	            "name": "Business Bank Account",
//	            #...
//	            "validDatatypeLinks": [
//	                {
//	                    "property": "Id",
//	                    "links": [
//	                        "Payment.AccountRef.Id",
//	                        "BillPayment.AccountRef.Id",
//	                        "DirectIncome.LineItems.AccountRef.Id",
//	                        "DirectCost.LineItems.AccountRef.Id"
//	                    ]
//	                }
//	            ]
//	        }
//
// ```
//
// ## Support for `validDatatypeLinks`
//
// Codat currently supports `validDatatypeLinks` for some data types on our Xero, QuickBooks Online, QuickBooks Desktop, Exact (NL), and Sage Business Cloud integrations.
//
// If you'd like us to extend support to more data types or integrations, suggest or vote for this on our <a href="https://portal.productboard.com/codat/5-product-roadmap">Product Roadmap</a>.
type GetTaxRate200ApplicationJSONValidDataTypeLinks struct {
	Links    []string `json:"links,omitempty"`
	Property *string  `json:"property,omitempty"`
}

// GetTaxRate200ApplicationJSON
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
type GetTaxRate200ApplicationJSON struct {
	Code               *string                                          `json:"code,omitempty"`
	Components         []GetTaxRate200ApplicationJSONComponents         `json:"components,omitempty"`
	EffectiveTaxRate   *float64                                         `json:"effectiveTaxRate,omitempty"`
	ID                 *string                                          `json:"id,omitempty"`
	Metadata           *GetTaxRate200ApplicationJSONMetadata            `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                       `json:"modifiedDate,omitempty"`
	Name               *string                                          `json:"name,omitempty"`
	SourceModifiedDate *time.Time                                       `json:"sourceModifiedDate,omitempty"`
	Status             *GetTaxRate200ApplicationJSONStatusEnum          `json:"status,omitempty"`
	TotalTaxRate       *float64                                         `json:"totalTaxRate,omitempty"`
	ValidDatatypeLinks []GetTaxRate200ApplicationJSONValidDataTypeLinks `json:"validDatatypeLinks,omitempty"`
}

type GetTaxRateResponse struct {
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
	GetTaxRate200ApplicationJSONObject *GetTaxRate200ApplicationJSON
}
