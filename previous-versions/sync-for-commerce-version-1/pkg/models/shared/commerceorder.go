// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CommerceOrderSupplementalData - Supplemental data is additional data you can include in our standard data types.
//
// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
type CommerceOrderSupplementalData struct {
	Content map[string]map[string]interface{} `json:"content,omitempty"`
}

func (o *CommerceOrderSupplementalData) GetContent() map[string]map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Content
}

// CommerceOrder - Orders contain the transaction details for all products sold by the company, and include details of any payments, service charges, or refunds related to each order. You can use data from the Orders endpoints to calculate key metrics, such as gross sales values and monthly recurring revenue (MRR).
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-orders) for this data type.
type CommerceOrder struct {
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	ClosedDate *string `json:"closedDate,omitempty"`
	// The Codat country property is returned as it was provided in the underlying platform by the company without any formatting on our part.
	//
	// Depending on the platform the value of this property will either be an <a href="https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes" target="_blank">ISO 3166</a> code (2-alpha or 3-alpha) or free-form text returned as a string name in our model.
	//
	// For POST operations against platforms that demand a specific format for the country code, we have documented accepted values in the [options](https://docs.codat.io/sync-for-commerce-v1-api#/operations/get-companies-companyId-connections-connectionId-push) endpoint.
	Country *string `json:"country,omitempty"`
	// In Codat's data model, dates and times are represented using the <a class="external" href="https://en.wikipedia.org/wiki/ISO_8601" target="_blank">ISO 8601 standard</a>. Date and time fields are formatted as strings; for example:
	//
	// ```
	// 2020-10-08T22:40:50Z
	// 2021-01-01T00:00:00
	// ```
	//
	//
	//
	// When syncing data that contains `DateTime` fields from Codat, make sure you support the following cases when reading time information:
	//
	// - Coordinated Universal Time (UTC): `2021-11-15T06:00:00Z`
	// - Unqualified local time: `2021-11-15T01:00:00`
	// - UTC time offsets: `2021-11-15T01:00:00-05:00`
	//
	// > Time zones
	// >
	// > Not all dates from Codat will contain information about time zones.
	// > Where it is not available from the underlying platform, Codat will return these as times local to the business whose data has been synced.
	CreatedDate *string `json:"createdDate,omitempty"`
	Currency    *string `json:"currency,omitempty"`
	// Reference to the customer that placed the order.
	CustomerRef *CommerceCustomerRef `json:"customerRef,omitempty"`
	// A unique, persistent identifier for this record
	ID string `json:"id"`
	// Reference to the geographic location where the order was placed.
	LocationRef    *LocationRef    `json:"locationRef,omitempty"`
	ModifiedDate   *string         `json:"modifiedDate,omitempty"`
	OrderLineItems []OrderLineItem `json:"orderLineItems,omitempty"`
	// Friendly reference for the order in the commerce or point of sale platform.
	OrderNumber        *string         `json:"orderNumber,omitempty"`
	Payments           []PaymentRef    `json:"payments,omitempty"`
	ServiceCharges     []ServiceCharge `json:"serviceCharges,omitempty"`
	SourceModifiedDate *string         `json:"sourceModifiedDate,omitempty"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting platform. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *CommerceOrderSupplementalData `json:"supplementalData,omitempty"`
	// Total amount of the order, including tax, net of any discounts and refunds.
	TotalAmount *float64 `json:"totalAmount,omitempty"`
	// Total amount of discount applied to the order.
	TotalDiscount *float64 `json:"totalDiscount,omitempty"`
	// Extra amount added to a bill.
	TotalGratuity *float64 `json:"totalGratuity,omitempty"`
	// Total amount refunded issued by a merchant on an order (always a negative value).
	TotalRefund *float64 `json:"totalRefund,omitempty"`
	// Total amount of tax applied to the order.
	TotalTaxAmount *float64 `json:"totalTaxAmount,omitempty"`
}

func (o *CommerceOrder) GetClosedDate() *string {
	if o == nil {
		return nil
	}
	return o.ClosedDate
}

func (o *CommerceOrder) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *CommerceOrder) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *CommerceOrder) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CommerceOrder) GetCustomerRef() *CommerceCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *CommerceOrder) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CommerceOrder) GetLocationRef() *LocationRef {
	if o == nil {
		return nil
	}
	return o.LocationRef
}

func (o *CommerceOrder) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *CommerceOrder) GetOrderLineItems() []OrderLineItem {
	if o == nil {
		return nil
	}
	return o.OrderLineItems
}

func (o *CommerceOrder) GetOrderNumber() *string {
	if o == nil {
		return nil
	}
	return o.OrderNumber
}

func (o *CommerceOrder) GetPayments() []PaymentRef {
	if o == nil {
		return nil
	}
	return o.Payments
}

func (o *CommerceOrder) GetServiceCharges() []ServiceCharge {
	if o == nil {
		return nil
	}
	return o.ServiceCharges
}

func (o *CommerceOrder) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *CommerceOrder) GetSupplementalData() *CommerceOrderSupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *CommerceOrder) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *CommerceOrder) GetTotalDiscount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalDiscount
}

func (o *CommerceOrder) GetTotalGratuity() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalGratuity
}

func (o *CommerceOrder) GetTotalRefund() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalRefund
}

func (o *CommerceOrder) GetTotalTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalTaxAmount
}
