// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateBankAccountTransaction struct {
	Amount  *float64 `json:"amount,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
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
	Date        *string `json:"date,omitempty"`
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
}

func (o *CreateBankAccountTransaction) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CreateBankAccountTransaction) GetBalance() *float64 {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *CreateBankAccountTransaction) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *CreateBankAccountTransaction) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateBankAccountTransaction) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}