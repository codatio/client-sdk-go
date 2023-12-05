// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Contact struct {
	Address *AccountingAddress `json:"address,omitempty"`
	// Email of a contact for a customer.
	Email *string `json:"email,omitempty"`
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
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// Name of a contact for a customer.
	Name *string `json:"name,omitempty"`
	// An array of Phone numbers.
	Phone []PhoneNumber `json:"phone,omitempty"`
	// Status of customer.
	Status CustomerStatus `json:"status"`
}

func (o *Contact) GetAddress() *AccountingAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Contact) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *Contact) GetModifiedDate() *string {
	if o == nil {
		return nil
	}
	return o.ModifiedDate
}

func (o *Contact) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Contact) GetPhone() []PhoneNumber {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *Contact) GetStatus() CustomerStatus {
	if o == nil {
		return CustomerStatus("")
	}
	return o.Status
}
