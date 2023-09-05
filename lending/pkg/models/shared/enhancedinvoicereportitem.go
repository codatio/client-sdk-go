// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EnhancedInvoiceReportItem struct {
	AmountDue *float64 `json:"amountDue,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	Currency    *string             `json:"currency,omitempty"`
	CustomerRef *LendingCustomerRef `json:"customerRef,omitempty"`
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
	// ID of the invoice, which may be a GUID but it may be something else depending on the accounting platdform
	ID            *string `json:"id,omitempty"`
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
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
	IssueDate    *string `json:"issueDate,omitempty"`
	ModifiedDate *string `json:"modifiedDate,omitempty"`
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
	PaidOnDate         *string   `json:"paidOnDate,omitempty"`
	Payments           []Payment `json:"payments,omitempty"`
	SourceModifiedDate *string   `json:"sourceModifiedDate,omitempty"`
	// Current state of the invoice:
	//
	// - `Draft` - Invoice hasn't been submitted to the supplier. It may be in a pending state or is scheduled for future submission, for example by email.
	// - `Submitted` - Invoice is no longer a draft. It has been processed and, or, sent to the customer. In this state, it will impact the ledger. It also has no payments made against it (amountDue == totalAmount).
	// - `PartiallyPaid` - The balance paid against the invoice is positive, but less than the total invoice amount (0 < amountDue < totalAmount).
	// - `Paid` - Invoice is paid in full. This includes if the invoice has been credited or overpaid (amountDue == 0).
	// - `Void` - An invoice can become Void when it's deleted, refunded, written off, or cancelled. A voided invoice may still be PartiallyPaid, and so all outstanding amounts on voided invoices are removed from the accounts receivable account.
	Status      *InvoiceStatus `json:"status,omitempty"`
	TotalAmount *float64       `json:"totalAmount,omitempty"`
}

func (o *EnhancedInvoiceReportItem) GetAmountDue() *float64 {
	if o == nil {
		return nil
	}
	return o.AmountDue
}

func (o *EnhancedInvoiceReportItem) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *EnhancedInvoiceReportItem) GetCustomerRef() *LendingCustomerRef {
	if o == nil {
		return nil
	}
	return o.CustomerRef
}

func (o *EnhancedInvoiceReportItem) GetDueDate() *string {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *EnhancedInvoiceReportItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EnhancedInvoiceReportItem) GetInvoiceNumber() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceNumber
}

func (o *EnhancedInvoiceReportItem) GetIssueDate() *string {
	if o == nil {
		return nil
	}
	return o.IssueDate
}

func (o *EnhancedInvoiceReportItem) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *EnhancedInvoiceReportItem) GetPaidOnDate() *string {
	if o == nil {
		return nil
	}
	return o.PaidOnDate
}

func (o *EnhancedInvoiceReportItem) GetPayments() []Payment {
	if o == nil {
		return nil
	}
	return o.Payments
}

func (o *EnhancedInvoiceReportItem) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *EnhancedInvoiceReportItem) GetStatus() *InvoiceStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *EnhancedInvoiceReportItem) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}