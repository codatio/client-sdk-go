// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Payment - Payments contain details of all payments made by customers to a company, including: amounts, currency used, payment method, payment provider, and payment status.
//
// Refunds are recorded as separate, negative payments. Note that a refund can only occur in relation to a payment that has been completed (i.e. has a status of `Paid`). When a customer cancels an order _before_ a payment has been completed, the payment shows as `Cancelled`.
//
// You can use data from the Payments endpoints to calculate key metrics, such as gross sales and monthly recurring revenue (MRR).
//
// Explore our [data coverage](https://knowledge.codat.io/supported-features/commerce?view=tab-by-data-type&dataType=commerce-payments) for this data type.
type Payment struct {
	// Payment Amount (including gratuity)
	Amount *float64 `json:"amount,omitempty"`
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
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency *string `json:"currency,omitempty"`
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
	DueDate *string `json:"dueDate,omitempty"`
	// A unique, persistent identifier for this record
	ID               string      `json:"id"`
	ModifiedDate     *string     `json:"modifiedDate,omitempty"`
	PaymentMethodRef interface{} `json:"paymentMethodRef,omitempty"`
	// Service provider of the payment, if applicable.
	PaymentProvider    *string `json:"paymentProvider,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of the payment.
	Status *PaymentStatus `json:"status,omitempty"`
}

func (o *Payment) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *Payment) GetCreatedDate() *string {
	if o == nil {
		return nil
	}
	return o.CreatedDate
}

func (o *Payment) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Payment) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *Payment) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Payment) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Payment) GetPaymentMethodRef() interface{} {
	if o == nil {
		return nil
	}
	return o.PaymentMethodRef
}

func (o *Payment) GetPaymentProvider() *string {
	if o == nil {
		return nil
	}
	return o.PaymentProvider
}

func (o *Payment) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *Payment) GetStatus() *PaymentStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
