// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SyncSummary - Success
type SyncSummary struct {
	// Unique identifier for the sync in Codat.
	CommerceSyncID *string `json:"commerceSyncId,omitempty"`
	// Unique identifier for your SMB in Codat.
	CompanyID *string `json:"companyId,omitempty"`
	// Array of containing objects data connection information for the company.
	DataConnections []Connection `json:"dataConnections,omitempty"`
	// Boolean indicator for data being pushed during a sync operation.
	DataPushed *bool `json:"dataPushed,omitempty"`
	// Friendly error message for the sync operation.
	ErrorMessage     *string    `json:"errorMessage,omitempty"`
	SyncDateRangeUtc *DateRange `json:"syncDateRangeUtc,omitempty"`
	// Exception message for the sync operation.
	SyncExceptionMessage *string `json:"syncExceptionMessage,omitempty"`
	// Status of the sync of the company data. This is linked to status code.
	SyncStatus *string `json:"syncStatus,omitempty"`
	// Numerical status code sync of the company data.
	SyncStatusCode *int `json:"syncStatusCode,omitempty"`
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
	SyncUtc *string `json:"syncUtc,omitempty"`
}
