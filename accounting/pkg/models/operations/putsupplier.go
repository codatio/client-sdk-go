// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type PutSupplierRequest struct {
	Supplier     *shared.Supplier `request:"mediaType=application/json"`
	CompanyID    string           `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string           `pathParam:"style=simple,explode=false,name=connectionId"`
	// When updating data in the destination platform Codat checks the `sourceModifiedDate` against the `lastupdated` date from the accounting platform, if they're different Codat will return an error suggesting you should initiate another pull of the data. If this is set to `true` then the update will override this check.
	ForceUpdate *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	// Unique identifier for a supplier
	SupplierID       string `pathParam:"style=simple,explode=false,name=supplierId"`
	TimeoutInMinutes *int   `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

// PutSupplier200ApplicationJSON - Success
type PutSupplier200ApplicationJSON struct {
	Changes []shared.PushOperationChange `json:"changes,omitempty"`
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
	Data *shared.Supplier `json:"data,omitempty"`
	// Unique identifier for a company's data connection.
	DataConnectionKey string `json:"dataConnectionKey"`
	// Available Data types
	DataType     *shared.DataTypeEnum `json:"dataType,omitempty"`
	ErrorMessage *string              `json:"errorMessage,omitempty"`
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
	// The status of the push operation.
	Status           shared.PushOperationStatusEnum `json:"status"`
	StatusCode       int64                          `json:"statusCode"`
	TimeoutInMinutes *int                           `json:"timeoutInMinutes,omitempty"`
	TimeoutInSeconds *int                           `json:"timeoutInSeconds,omitempty"`
	// A human-readable object describing validation decisions Codat has made when pushing data into the platform. If a push has failed because of validation errors, they will be detailed here.
	Validation *shared.Validation `json:"validation,omitempty"`
}

type PutSupplierResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	PutSupplier200ApplicationJSONObject *PutSupplier200ApplicationJSON
}
