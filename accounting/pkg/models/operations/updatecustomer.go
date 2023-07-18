// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/accounting/pkg/models/shared"
	"net/http"
)

type UpdateCustomerRequest struct {
	Customer     *shared.Customer `request:"mediaType=application/json"`
	CompanyID    string           `pathParam:"style=simple,explode=false,name=companyId"`
	ConnectionID string           `pathParam:"style=simple,explode=false,name=connectionId"`
	CustomerID   string           `pathParam:"style=simple,explode=false,name=customerId"`
	// When updating data in the destination platform Codat checks the `sourceModifiedDate` against the `lastupdated` date from the accounting platform, if they're different Codat will return an error suggesting you should initiate another pull of the data. If this is set to `true` then the update will override this check.
	ForceUpdate      *bool `queryParam:"style=form,explode=true,name=forceUpdate"`
	TimeoutInMinutes *int  `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *UpdateCustomerRequest) GetCustomer() *shared.Customer {
	if o == nil {
		return nil
	}
	return o.Customer
}

func (o *UpdateCustomerRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *UpdateCustomerRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *UpdateCustomerRequest) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *UpdateCustomerRequest) GetForceUpdate() *bool {
	if o == nil {
		return nil
	}
	return o.ForceUpdate
}

func (o *UpdateCustomerRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type UpdateCustomerResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	UpdateCustomerResponse *shared.UpdateCustomerResponse
	// The request made is not valid.
	Schema *shared.Schema
}

func (o *UpdateCustomerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCustomerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCustomerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCustomerResponse) GetUpdateCustomerResponse() *shared.UpdateCustomerResponse {
	if o == nil {
		return nil
	}
	return o.UpdateCustomerResponse
}

func (o *UpdateCustomerResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
