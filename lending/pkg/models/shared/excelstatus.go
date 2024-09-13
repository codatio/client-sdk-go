// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ExcelStatus struct {
	// Error details in case the report generation request was unsuccessful.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The file size in Bytes is populated upon successful generation of the report.
	FileSize *int64 `json:"fileSize,omitempty"`
	// When true, the request was successful and the report is being generated. If false, the request was unsuccessful and the report is not being generated.
	InProgress *bool `json:"inProgress,omitempty"`
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
	LastGenerated *string `json:"lastGenerated,omitempty"`
	// A unique ID generated for this request.
	LastInvocationID *string `json:"lastInvocationId,omitempty"`
	// The date and time of when a successful request was queued for the most recent report.
	Queued *string `json:"queued,omitempty"`
	// The type of the report requested in the query string.
	ReportType *ExcelReportTypes `json:"reportType,omitempty"`
	// True if the requested report was successfully queued and false if the requested report was not able to be queued.
	Success *bool `json:"success,omitempty"`
}

func (o *ExcelStatus) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ExcelStatus) GetFileSize() *int64 {
	if o == nil {
		return nil
	}
	return o.FileSize
}

func (o *ExcelStatus) GetInProgress() *bool {
	if o == nil {
		return nil
	}
	return o.InProgress
}

func (o *ExcelStatus) GetLastGenerated() *string {
	if o == nil {
		return nil
	}
	return o.LastGenerated
}

func (o *ExcelStatus) GetLastInvocationID() *string {
	if o == nil {
		return nil
	}
	return o.LastInvocationID
}

func (o *ExcelStatus) GetQueued() *string {
	if o == nil {
		return nil
	}
	return o.Queued
}

func (o *ExcelStatus) GetReportType() *ExcelReportTypes {
	if o == nil {
		return nil
	}
	return o.ReportType
}

func (o *ExcelStatus) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}
