// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v2/pkg/utils"
	"net/http"
)

type UpdateSupplierRequest struct {
	Supplier *shared.Supplier `request:"mediaType=application/json"`
	// Allow a sync upon push completion.
	AllowSyncOnPushComplete *bool `default:"true" queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// When updating data in the destination platform Codat checks the `sourceModifiedDate` against the `lastupdated` date from the accounting platform, if they're different Codat will return an error suggesting you should initiate another pull of the data. If this is set to `true` then the update will override this check.
	ForceUpdate *bool `default:"false" queryParam:"style=form,explode=true,name=forceUpdate"`
	// Unique identifier for a supplier.
	SupplierID string `pathParam:"style=simple,explode=false,name=supplierId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (u UpdateSupplierRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateSupplierRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateSupplierRequest) GetSupplier() *shared.Supplier {
	if o == nil {
		return nil
	}
	return o.Supplier
}

func (o *UpdateSupplierRequest) GetAllowSyncOnPushComplete() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSyncOnPushComplete
}

func (o *UpdateSupplierRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateSupplierRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateSupplierRequest) GetForceUpdate() *bool {
	if o == nil {
		return nil
	}
	return o.ForceUpdate
}

func (o *UpdateSupplierRequest) GetSupplierID() string {
	if o == nil {
		return ""
	}
	return o.SupplierID
}

func (o *UpdateSupplierRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type UpdateSupplierResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	UpdateSupplierResponse *shared.UpdateSupplierResponse
}

func (o *UpdateSupplierResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSupplierResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *UpdateSupplierResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSupplierResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSupplierResponse) GetUpdateSupplierResponse() *shared.UpdateSupplierResponse {
	if o == nil {
		return nil
	}
	return o.UpdateSupplierResponse
}
