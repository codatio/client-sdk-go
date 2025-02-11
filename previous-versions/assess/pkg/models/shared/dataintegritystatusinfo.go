// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DataIntegrityStatusInfo struct {
	// The current status of the most recently run matching algorithm.
	CurrentStatus *IntegrityStatus `json:"currentStatus,omitempty"`
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
	LastMatched *string `json:"lastMatched,omitempty"`
	// Detailed explanation supporting the status value.
	StatusMessage *string `json:"statusMessage,omitempty"`
}

func (o *DataIntegrityStatusInfo) GetCurrentStatus() *IntegrityStatus {
	if o == nil {
		return nil
	}
	return o.CurrentStatus
}

func (o *DataIntegrityStatusInfo) GetLastMatched() *string {
	if o == nil {
		return nil
	}
	return o.LastMatched
}

func (o *DataIntegrityStatusInfo) GetStatusMessage() *string {
	if o == nil {
		return nil
	}
	return o.StatusMessage
}
