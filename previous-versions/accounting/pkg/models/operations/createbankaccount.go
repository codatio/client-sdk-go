// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/models/shared"
	"github.com/codatio/client-sdk-go/previous-versions/accounting/pkg/utils"
	"net/http"
)

type CreateBankAccountRequest struct {
	BankAccount *shared.BankAccount `request:"mediaType=application/json"`
	// Allow a sync upon push completion.
	AllowSyncOnPushComplete *bool `default:"true" queryParam:"style=form,explode=true,name=allowSyncOnPushComplete"`
	// Unique identifier for a company.
	CompanyID string `pathParam:"style=simple,explode=false,name=companyId"`
	// Unique identifier for a connection.
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionId"`
	// Time limit for the push operation to complete before it is timed out.
	TimeoutInMinutes *int `queryParam:"style=form,explode=true,name=timeoutInMinutes"`
}

func (c CreateBankAccountRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateBankAccountRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateBankAccountRequest) GetBankAccount() *shared.BankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccount
}

func (o *CreateBankAccountRequest) GetAllowSyncOnPushComplete() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSyncOnPushComplete
}

func (o *CreateBankAccountRequest) GetCompanyID() string {
	if o == nil {
		return ""
	}
	return o.CompanyID
}

func (o *CreateBankAccountRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *CreateBankAccountRequest) GetTimeoutInMinutes() *int {
	if o == nil {
		return nil
	}
	return o.TimeoutInMinutes
}

type CreateBankAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	CreateBankAccountResponse *shared.CreateBankAccountResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateBankAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateBankAccountResponse) GetCreateBankAccountResponse() *shared.CreateBankAccountResponse {
	if o == nil {
		return nil
	}
	return o.CreateBankAccountResponse
}

func (o *CreateBankAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateBankAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
