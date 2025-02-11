// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateCustomerResponseAccountingCustomer - > View the coverage for customers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=customers" target="_blank">Data coverage explorer</a>.
//
// ## Overview
//
// A customer is a person or organisation that buys goods or services. From the Customers endpoints, you can retrieve a [list of all the customers of a company](https://api.codat.io/swagger/index.html#/Customers/get_companies__companyId__data_customers).
//
// Customers' data links to accounts receivable [invoices](https://docs.codat.io/accounting-api#/schemas/Invoice).
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type UpdateCustomerResponseAccountingCustomer struct {
	// An array of Addresses.
	Addresses []Items `json:"addresses,omitempty"`
	// Name of the main contact for the identified customer.
	ContactName *string `json:"contactName,omitempty"`
	// An array of Contacts.
	Contacts []Contact `json:"contacts,omitempty"`
	// Name of the customer as recorded in the accounting system, typically the company name.
	CustomerName *string `json:"customerName,omitempty"`
	// The currency data type in Codat is the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code, e.g. _GBP_.
	//
	// ## Unknown currencies
	//
	// In line with the ISO 4217 specification, the code _XXX_ is used when the data source does not return a currency for a transaction.
	//
	// There are only a very small number of edge cases where this currency code is returned by the Codat system.
	DefaultCurrency *string `json:"defaultCurrency,omitempty"`
	// Email address the customer can be contacted by.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Identifier for the customer, unique to the company in the accounting software.
	ID           *string   `json:"id,omitempty"`
	Metadata     *Metadata `json:"metadata,omitempty"`
	ModifiedDate *string   `json:"modifiedDate,omitempty"`
	// Phone number the customer can be contacted by.
	Phone *string `json:"phone,omitempty"`
	// Company number. In the UK, this is typically the Companies House company registration number.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	SourceModifiedDate *string `json:"sourceModifiedDate,omitempty"`
	// Status of customer.
	Status CustomerStatus `json:"status"`
	// Supplemental data is additional data you can include in our standard data types.
	//
	// It is referenced as a configured dynamic key value pair that is unique to the accounting software. [Learn more](https://docs.codat.io/using-the-api/supplemental-data/overview) about supplemental data.
	SupplementalData *SupplementalData `json:"supplementalData,omitempty"`
	// Company tax number.
	TaxNumber *string `json:"taxNumber,omitempty"`
}

func (o *UpdateCustomerResponseAccountingCustomer) GetAddresses() []Items {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *UpdateCustomerResponseAccountingCustomer) GetContactName() *string {
	if o == nil {
		return nil
	}
	return o.ContactName
}

func (o *UpdateCustomerResponseAccountingCustomer) GetContacts() []Contact {
	if o == nil {
		return nil
	}
	return o.Contacts
}

func (o *UpdateCustomerResponseAccountingCustomer) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *UpdateCustomerResponseAccountingCustomer) GetDefaultCurrency() *string {
	if o == nil {
		return nil
	}
	return o.DefaultCurrency
}

func (o *UpdateCustomerResponseAccountingCustomer) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *UpdateCustomerResponseAccountingCustomer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateCustomerResponseAccountingCustomer) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *UpdateCustomerResponseAccountingCustomer) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *UpdateCustomerResponseAccountingCustomer) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *UpdateCustomerResponseAccountingCustomer) GetRegistrationNumber() *string {
	if o == nil {
		return nil
	}
	return o.RegistrationNumber
}

func (o *UpdateCustomerResponseAccountingCustomer) GetSourceModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.SourceModifiedDate
}

func (o *UpdateCustomerResponseAccountingCustomer) GetStatus() CustomerStatus {
	if o == nil {
		return CustomerStatus("")
	}
	return o.Status
}

func (o *UpdateCustomerResponseAccountingCustomer) GetSupplementalData() *SupplementalData {
	if o == nil {
		return nil
	}
	return o.SupplementalData
}

func (o *UpdateCustomerResponseAccountingCustomer) GetTaxNumber() *string {
	if o == nil {
		return nil
	}
	return o.TaxNumber
}

type UpdateCustomerResponse struct {
	// Contains a single entry that communicates which record has changed and the manner in which it changed.
	Changes []PushOperationChange `json:"changes,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID string `json:"companyId"`
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
	CompletedOnUtc *string                                   `json:"completedOnUtc,omitempty"`
	Data           *UpdateCustomerResponseAccountingCustomer `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// Available data types
	DataType *DataType `json:"dataType,omitempty"`
	// A message about the error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// A unique identifier generated by Codat to represent this single push operation. This identifier can be used to track the status of the push, and should be persisted.
	PushOperationKey string `json:"pushOperationKey"`
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
	RequestedOnUtc string `json:"requestedOnUtc"`
	// The current status of the push operation.
	Status PushOperationStatus `json:"status"`
	// Push status code.
	StatusCode int64 `json:"statusCode"`
	// Number of minutes the push operation must complete within before it times out.
	TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
	// Number of seconds the push operation must complete within before it times out.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *Validation `json:"validation,omitempty"`
}

func (o *UpdateCustomerResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *UpdateCustomerResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateCustomerResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *UpdateCustomerResponse) GetData() *UpdateCustomerResponseAccountingCustomer {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UpdateCustomerResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *UpdateCustomerResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *UpdateCustomerResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateCustomerResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *UpdateCustomerResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *UpdateCustomerResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *UpdateCustomerResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCustomerResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *UpdateCustomerResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *UpdateCustomerResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}
