package operations

import (
	"net/http"
	"time"
)

type GetSupplierPathParams struct {
	CompanyID  string `pathParam:"style=simple,explode=false,name=companyId"`
	SupplierID string `pathParam:"style=simple,explode=false,name=supplierId"`
}

type GetSupplierRequest struct {
	PathParams GetSupplierPathParams
}

type GetSupplierSourceModifiedDateAddressesTypeEnum string

const (
	GetSupplierSourceModifiedDateAddressesTypeEnumUnknown  GetSupplierSourceModifiedDateAddressesTypeEnum = "Unknown"
	GetSupplierSourceModifiedDateAddressesTypeEnumBilling  GetSupplierSourceModifiedDateAddressesTypeEnum = "Billing"
	GetSupplierSourceModifiedDateAddressesTypeEnumDelivery GetSupplierSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type GetSupplierSourceModifiedDateAddresses struct {
	City       *string                                        `json:"city,omitempty"`
	Country    *string                                        `json:"country,omitempty"`
	Line1      *string                                        `json:"line1,omitempty"`
	Line2      *string                                        `json:"line2,omitempty"`
	PostalCode *string                                        `json:"postalCode,omitempty"`
	Region     *string                                        `json:"region,omitempty"`
	Type       GetSupplierSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type GetSupplierSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type GetSupplierSourceModifiedDateStatusEnum string

const (
	GetSupplierSourceModifiedDateStatusEnumUnknown  GetSupplierSourceModifiedDateStatusEnum = "Unknown"
	GetSupplierSourceModifiedDateStatusEnumActive   GetSupplierSourceModifiedDateStatusEnum = "Active"
	GetSupplierSourceModifiedDateStatusEnumArchived GetSupplierSourceModifiedDateStatusEnum = "Archived"
)

type GetSupplierSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// GetSupplierSourceModifiedDate
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type GetSupplierSourceModifiedDate struct {
	Addresses          []GetSupplierSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                        `json:"contactName,omitempty"`
	DefaultCurrency    *string                                        `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                        `json:"emailAddress,omitempty"`
	ID                 *string                                        `json:"id,omitempty"`
	Metadata           *GetSupplierSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                     `json:"modifiedDate,omitempty"`
	Phone              *string                                        `json:"phone,omitempty"`
	RegistrationNumber *string                                        `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                     `json:"sourceModifiedDate,omitempty"`
	Status             GetSupplierSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *GetSupplierSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	SupplierName       *string                                        `json:"supplierName,omitempty"`
	TaxNumber          *string                                        `json:"taxNumber,omitempty"`
}

type GetSupplierResponse struct {
	ContentType        string
	SourceModifiedDate *GetSupplierSourceModifiedDate
	StatusCode         int
	RawResponse        *http.Response
}
