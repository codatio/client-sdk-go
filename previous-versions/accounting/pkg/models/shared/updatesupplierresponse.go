// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateSupplierResponse struct {
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
	CompletedOnUtc *string `json:"completedOnUtc,omitempty"`
	// > View the coverage for suppliers in the <a className="external" href="https://knowledge.codat.io/supported-features/accounting?view=tab-by-data-type&dataType=suppliers" target="_blank">Data coverage explorer</a>.
	//
	// ## Overview
	//
	// From the **Suppliers** endpoints, you can retrieve a list of [all the suppliers for a company](https://docs.codat.io/accounting-api#/operations/list-suppliers). Suppliers' data links to accounts payable [bills](https://docs.codat.io/accounting-api#/schemas/Bill).
	Data *Supplier `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// Available Data types
	DataType     *DataType `json:"dataType,omitempty"`
	ErrorMessage *string   `json:"errorMessage,omitempty"`
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
	Status           PushOperationStatus `json:"status"`
	StatusCode       int64               `json:"statusCode"`
	TimeoutInMinutes *int                `json:"timeoutInMinutes,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TimeoutInSeconds *int `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *Validation `json:"validation,omitempty"`
}

func (o *UpdateSupplierResponse) GetChanges() []PushOperationChange {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *UpdateSupplierResponse) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateSupplierResponse) GetCompletedOnUtc() *string {
	if o == nil {
		return nil
	}
	return o.CompletedOnUtc
}

func (o *UpdateSupplierResponse) GetData() *Supplier {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UpdateSupplierResponse) GetDataConnectionKey() string {
	if o == nil {
		return ""
	}
	return o.DataConnectionKey
}

func (o *UpdateSupplierResponse) GetDataType() *DataType {
	if o == nil {
		return nil
	}
	return o.DataType
}

func (o *UpdateSupplierResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateSupplierResponse) GetPushOperationKey() string {
	if o == nil {
		return ""
	}
	return o.PushOperationKey
}

func (o *UpdateSupplierResponse) GetRequestedOnUtc() string {
	if o == nil {
		return ""
	}
	return o.RequestedOnUtc
}

func (o *UpdateSupplierResponse) GetStatus() PushOperationStatus {
	if o == nil {
		return PushOperationStatus("")
	}
	return o.Status
}

func (o *UpdateSupplierResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSupplierResponse) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

func (o *UpdateSupplierResponse) GetTimeoutInSeconds() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInSeconds
}

func (o *UpdateSupplierResponse) GetValidation() *Validation {
	if o == nil {
		return nil
	}
	return o.Validation
}
