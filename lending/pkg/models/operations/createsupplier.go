// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/lending/v6/pkg/models/shared"
	"github.com/codatio/client-sdk-go/lending/v6/pkg/utils"
	"net/http"
)

type CreateSupplierRequest struct {
	AccountingSupplier *shared.AccountingSupplier `request:"mediaType=application/json"`
	// Allow a sync upon push completion.
	AllowSyncOnPushComplete *bool `default:"true" queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (c CreateSupplierRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateSupplierRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateSupplierRequest) GetAccountingSupplier() *shared.AccountingSupplier {
	if o == nil {
		return nil
	}
	return o.AccountingSupplier
}

func (o *CreateSupplierRequest) GetAllowSyncOnPushComplete() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSyncOnPushComplete
}

func (o *CreateSupplierRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateSupplierRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateSupplierRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateSupplierResponse struct {
	// Success
	AccountingCreateSupplierResponse *shared.AccountingCreateSupplierResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateSupplierResponse) GetAccountingCreateSupplierResponse() *shared.AccountingCreateSupplierResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreateSupplierResponse
}

func (o *CreateSupplierResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSupplierResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSupplierResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
