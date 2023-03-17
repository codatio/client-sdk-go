package operations

import (
	"net/http"
	"time"
)

type ListSuppliersRequest struct {
	CompanyID string  `pathParam:"style=simple,explode=false,name=companyId"`
	OrderBy   *string `queryParam:"style=form,explode=true,name=orderBy"`
	Page      int     `queryParam:"style=form,explode=true,name=page"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=pageSize"`
	Query     *string `queryParam:"style=form,explode=true,name=query"`
}

type ListSuppliersLinksLinksCurrent struct {
	Href string `json:"href"`
}

type ListSuppliersLinksLinksNext struct {
	Href *string `json:"href,omitempty"`
}

type ListSuppliersLinksLinksPrevious struct {
	Href *string `json:"href,omitempty"`
}

type ListSuppliersLinksLinksSelf struct {
	Href string `json:"href"`
}

type ListSuppliersLinksLinks struct {
	Current  ListSuppliersLinksLinksCurrent   `json:"current"`
	Next     *ListSuppliersLinksLinksNext     `json:"next,omitempty"`
	Previous *ListSuppliersLinksLinksPrevious `json:"previous,omitempty"`
	Self     ListSuppliersLinksLinksSelf      `json:"self"`
}

type ListSuppliersLinksSourceModifiedDateAddressesTypeEnum string

const (
	ListSuppliersLinksSourceModifiedDateAddressesTypeEnumUnknown  ListSuppliersLinksSourceModifiedDateAddressesTypeEnum = "Unknown"
	ListSuppliersLinksSourceModifiedDateAddressesTypeEnumBilling  ListSuppliersLinksSourceModifiedDateAddressesTypeEnum = "Billing"
	ListSuppliersLinksSourceModifiedDateAddressesTypeEnumDelivery ListSuppliersLinksSourceModifiedDateAddressesTypeEnum = "Delivery"
)

type ListSuppliersLinksSourceModifiedDateAddresses struct {
	City       *string                                               `json:"city,omitempty"`
	Country    *string                                               `json:"country,omitempty"`
	Line1      *string                                               `json:"line1,omitempty"`
	Line2      *string                                               `json:"line2,omitempty"`
	PostalCode *string                                               `json:"postalCode,omitempty"`
	Region     *string                                               `json:"region,omitempty"`
	Type       ListSuppliersLinksSourceModifiedDateAddressesTypeEnum `json:"type"`
}

type ListSuppliersLinksSourceModifiedDateMetadata struct {
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

type ListSuppliersLinksSourceModifiedDateStatusEnum string

const (
	ListSuppliersLinksSourceModifiedDateStatusEnumUnknown  ListSuppliersLinksSourceModifiedDateStatusEnum = "Unknown"
	ListSuppliersLinksSourceModifiedDateStatusEnumActive   ListSuppliersLinksSourceModifiedDateStatusEnum = "Active"
	ListSuppliersLinksSourceModifiedDateStatusEnumArchived ListSuppliersLinksSourceModifiedDateStatusEnum = "Archived"
)

// ListSuppliersLinksSourceModifiedDateSupplementalData
// Reference to a configured dynamic key value pair that is unique to the accounting platform. This feature is in private beta, contact us if you would like to learn more.
type ListSuppliersLinksSourceModifiedDateSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

// ListSuppliersLinksSourceModifiedDate
// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://api.codat.io/swagger/index.html#/Suppliers/get_companies__companyId__data_suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
type ListSuppliersLinksSourceModifiedDate struct {
	Addresses          []ListSuppliersLinksSourceModifiedDateAddresses       `json:"addresses,omitempty"`
	ContactName        *string                                               `json:"contactName,omitempty"`
	DefaultCurrency    *string                                               `json:"defaultCurrency,omitempty"`
	EmailAddress       *string                                               `json:"emailAddress,omitempty"`
	ID                 *string                                               `json:"id,omitempty"`
	Metadata           *ListSuppliersLinksSourceModifiedDateMetadata         `json:"metadata,omitempty"`
	ModifiedDate       *time.Time                                            `json:"modifiedDate,omitempty"`
	Phone              *string                                               `json:"phone,omitempty"`
	RegistrationNumber *string                                               `json:"registrationNumber,omitempty"`
	SourceModifiedDate *time.Time                                            `json:"sourceModifiedDate,omitempty"`
	Status             ListSuppliersLinksSourceModifiedDateStatusEnum        `json:"status"`
	SupplementalData   *ListSuppliersLinksSourceModifiedDateSupplementalData `json:"supplementalData,omitempty"`
	SupplierName       *string                                               `json:"supplierName,omitempty"`
	TaxNumber          *string                                               `json:"taxNumber,omitempty"`
}

// ListSuppliersLinks
// Codat's Paging Model
type ListSuppliersLinks struct {
	Links        ListSuppliersLinksLinks                `json:"_links"`
	PageNumber   int64                                  `json:"pageNumber"`
	PageSize     int64                                  `json:"pageSize"`
	Results      []ListSuppliersLinksSourceModifiedDate `json:"results,omitempty"`
	TotalResults int64                                  `json:"totalResults"`
}

type ListSuppliersResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Links       *ListSuppliersLinks
}
