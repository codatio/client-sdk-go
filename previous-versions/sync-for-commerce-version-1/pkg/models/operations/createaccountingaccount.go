// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/sync-for-commerce-version-1/pkg/models/shared"
	"net/http"
)

type CreateAccountingAccountRequest struct {
	AccountPrototype *shared.AccountPrototype `request:"mediaType=application/json"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (o *CreateAccountingAccountRequest) GetAccountPrototype() *shared.AccountPrototype {
	if o == nil {
		return nil
	}
	return o.AccountPrototype
}

func (o *CreateAccountingAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateAccountingAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateAccountingAccountRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateAccountingAccountResponse struct {
	// Success
	AccountingCreateAccountResponse *shared.AccountingCreateAccountResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAccountingAccountResponse) GetAccountingCreateAccountResponse() *shared.AccountingCreateAccountResponse {
	if o == nil {
		return nil
	}
	return o.AccountingCreateAccountResponse
}

func (o *CreateAccountingAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAccountingAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAccountingAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
