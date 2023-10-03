// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/models/shared"
	"github.com/codatio/client-sdk-go/sync-for-expenses/v3/pkg/utils"
	"net/http"
)

type CreateCustomerRequest struct {
	Customer *shared.Customer `request:"mediaType=application/json"`
	// Allow a sync upon push completion.
	AllowSyncOnPushComplete *bool `default:"true" queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (c CreateCustomerRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCustomerRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCustomerRequest) GetCustomer() *shared.Customer {
	if o == nil {
		return nil
	}
	return o.Customer
}

func (o *CreateCustomerRequest) GetAllowSyncOnPushComplete() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSyncOnPushComplete
}

func (o *CreateCustomerRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateCustomerRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateCustomerRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateCustomerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	CreateCustomerResponse *shared.CreateCustomerResponse
	// The request made is not valid.
	ErrorMessage *shared.ErrorMessage
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCustomerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCustomerResponse) GetCreateCustomerResponse() *shared.CreateCustomerResponse {
	if o == nil {
		return nil
	}
	return o.CreateCustomerResponse
}

func (o *CreateCustomerResponse) GetErrorMessage() *shared.ErrorMessage {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *CreateCustomerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCustomerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
